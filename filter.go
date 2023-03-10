package filter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type Config struct {
	Params map[string]string `json:"params:omitempty"`
}

func CreateConfig() *Config {
	return &Config{
		Params: make(map[string]string),
	}
}

type Filter struct {
	next   http.Handler
	params map[string]string
	name   string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Params) == 0 {
		return nil, fmt.Errorf("params cannot be empty")
	}
	return &Filter{
		params: config.Params,
		next:   next,
		name:   name,
	}, nil
}

func (m *Filter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()
	for key, value := range m.params {
		param := values.Get(key)
		if param == value {
			http.Error(rw, errors.New("parameter error").Error(), http.StatusForbidden)
			return
		}
	}
	m.next.ServeHTTP(rw, req)
}
