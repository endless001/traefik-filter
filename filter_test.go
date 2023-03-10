package filter_test

import (
	"context"
	"github.com/endless001/filter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFilter(t *testing.T) {
	cfg := filter.CreateConfig()
	cfg.Params["name"] = "lq"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := filter.New(ctx, next, cfg, "test-plugin")

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost?name=lq", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

}
