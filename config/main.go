package config

import (
	"context"
	"strconv"
)

// key is the local type we use to store the config in the context.Context package
type key int

const config key = 0 // Key for storing a store in a context.Context
const (
	Registry  = "reg"
	NatsProxy = "nprox"
	Streamer  = "streamer"
	RESTPort  = "rport"
)

// Store is a generic repository for application configuartion information
type Store interface {
	Set(key string, data interface{})
	Get(key string) (interface{}, bool)
	String(key string) (string, bool)
	Int(key string) (int, bool)
}

// FromContext pulls a store out of the provided context if it exists
func FromContext(c context.Context) Store {
	s := c.Value(config)
	if s != nil {
		return s.(Store)
	}

	return nil
}

// ToContext will save the store into the returned context
func ToContext(c context.Context, s Store) context.Context {
	return context.WithValue(c, config, s)
}

// baseConfig implements the store interface
type baseConfig struct {
	options map[string]interface{}
}

// Returns a new baseConfig which implements the Store interface
func New() Store {
	return &baseConfig{options: map[string]interface{}{}}
}

func (b *baseConfig) Set(key string, data interface{}) {
	b.options[key] = data
}

func (b *baseConfig) Get(key string) (interface{}, bool) {
	k, ok := b.options[key]
	return k, ok
}

func (b *baseConfig) String(key string) (string, bool) {
	if s, ok := b.options[key]; ok {
		if sstring, ok := s.(string); ok {
			return sstring, ok
		} else {
			return "", false
		}
	} else {
		return "", false
	}
}

func (b *baseConfig) Int(key string) (int, bool) {
	if i, ok := b.options[key]; ok {
		if iint, ok := i.(int); ok {
			return iint, ok
		}
		if ist, ok := i.(string); ok {
			iint, err := strconv.Atoi(ist)
			if err == nil {
				return iint, true
			}
		}
	}
	return -1, false
}
