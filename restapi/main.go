package restapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/julienschmidt/httprouter"
	"github.com/noptics/focal/config"
	"github.com/noptics/focal/registrygrpc"
	"github.com/noptics/golog"
)

type RESTServer struct {
	l       golog.Logger
	hs      http.Server
	config  config.Store
	errChan chan error
}

// RESTError is the standard structure for rest api errors
type RESTError struct {
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
	Details     string `json:"details"`
}

type RouteReply struct {
	Code  int
	Error *RESTError
	Data  interface{}
}

type RouteHandler func(*http.Request, httprouter.Params) *RouteReply

func NewRestServer(errChan chan error, l golog.Logger, config config.Store) *RESTServer {
	rs := &RESTServer{
		l:       l,
		errChan: errChan,
	}

	port, _ := config.String(config.RESTPort)

	rs.hs = http.Server{Addr: ":" + port, Handler: rs.Router()}

	l.Infow("starting rest server", "port", port)
	go func() {
		if err := rs.hs.ListenAndServe(); err != nil {
			rs.errChan <- err
		}
	}()

	return rs
}

func (rs *RESTServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return rs.hs.Shutdown(ctx)
}

func (rs *RESTServer) Status(r *http.Request, ps httprouter.Params) *RouteReply {
	return &RouteReply{Code: 200, Data: map[string]string{
		"go":      runtime.Version(),
		"version": os.Getenv("VERSION"),
		"commit":  os.Getenv("COMMIT")}}
}

func setHeaders(w http.ResponseWriter, rh http.Header) {
	allowMethod := "GET, POST, PUT, DELETE, OPTIONS"
	allowHeaders := "Content-Type"
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Allow", allowMethod)
	w.Header().Set("Access-Control-Allow-Methods", allowMethod)
	w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

	o := rh.Get("Origin")
	if o == "" {
		o = "*"
	}
	w.Header().Set("Access-Control-Allow-Origin", o)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
}

func writeReply(status int, body []byte, w http.ResponseWriter) {
	w.WriteHeader(status)
	if len(body) != 0 {
		w.Write(body)
	}
}

func wrapRoute(rh RouteHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		response := rh(r, ps)

		setHeaders(w, r.Header)
		var body []byte
		var err error

		if response.Error != nil {
			body, err = json.Marshal(response.Error)
		} else if response.Data != nil {
			body, err = json.Marshal(response.Data)
		}

		if err != nil {
			body = []byte(fmt.Sprintf(`{"message": "error marshalling response", "description":"unable to json encode data", "details":"%s"`, err.Error()))
			response.Code = 500
		}

		w.WriteHeader(response.Code)
		if len(body) != 0 {
			w.Write(body)
		}
	}
}

func newReplyError(message, description, details string, status int) *RouteReply {
	return &RouteReply{
		Code: status,
		Error: &RESTError{
			Message:     message,
			Description: description,
			Details:     details,
		},
	}
}

func (rs *RESTServer) Router() *httprouter.Router {
	r := httprouter.New()

	r.HandleOPTIONS = true
	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w, r.Header)
		w.WriteHeader(http.StatusNoContent)
	})

	r.POST("/:cluster/:channel/files", wrapRoute(rs.SaveFiles))
	// r.POST("/:cluster/:channel/message", wrapRoute(rs.SetMessage))
	// r.GET("/:cluster/:channel", wrapRoute(rs.GetChannel))
	// r.POST("/:cluster/:channel", wrapRoute(rs.SetChannelData))
	// r.GET("/:cluster", wrapRoute(rs.GetChannels))
	r.GET("/", wrapRoute(rs.Status))

	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(404)
	})

	return r
}

func (rs *RESTServer) SaveFiles(r *http.Request, ps httprouter.Params) *RouteReply {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return newReplyError("request error", "unble to read POST body", err.Error(), 500)
	}

	sfr := &registrygrpc.SaveFilesRequest{}
	err = jsonpb.Unmarshal(bytes.NewBuffer(body), sfr)
	if err != nil {
		return newReplyError("unable to save files", "error parsing POST body", err.Error(), 400)
	}

	if len(sfr.Channel) == 0 {
		sfr.Channel = ps.ByName("channel")
	}

	if len(sfr.ClusterID) == 0 {
		sfr.ClusterID = ps.ByName("cluster")
	}

	if len(sfr.ClusterID) == 0 || len(sfr.Channel) == 0 {
		return newReplyError("unable to save files", "invalid parameters provided", "must provide channel and cluster", 400)
	}

	rc, _ := rs.config.Get(config.Registry)

	_, err = rc.(registrygrpc.ProtoRegistryClient).SaveFiles(context.Background(), sfr)
	if err != nil {
		return newReplyError("unable to save files", "registry service error", err.Error(), 400)
	}

	return &RouteReply{Code: 200}
}
