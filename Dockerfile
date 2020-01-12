FROM golang:1.13 as Builder

RUN mkdir -p /go/src/github.com/noptics/focal
ADD . /go/src/github.com/noptics/focal

WORKDIR /go/src/github.com/noptics/focal

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o focal

FROM alpine:3.9

RUN apk add --no-cache curl bash ca-certificates

COPY --from=builder /go/src/github.com/noptics/focal/focal /focal

CMD ["/focal"]