package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_Handlers_UpdateGet(t *testing.T) {

	testSet := []struct {
		name       string
		uri        string
		wantStatus int
		wantErr    bool
	}{
		{
			name:       "common correct case",
			uri:        "/update/gauge/Mallocs/300",
			wantStatus: 200,
		},
		{
			name:       "common correct case counter",
			uri:        "/update/counter/Test/300",
			wantStatus: 200,
		},
		{
			name:       "common correct case counter",
			uri:        "/update/counter/Test/700",
			wantStatus: 200,
		},
		{
			name:       "missed metric name",
			uri:        "/update/gauge/",
			wantStatus: 404,
		},
		{
			name:       "missed metric value",
			uri:        "/update/gauge/Mallocs",
			wantStatus: 404,
		},
		{
			name:       "wrong type",
			uri:        "/update/something/",
			wantStatus: 404,
		},
	}

	testGet := []struct {
		name       string
		uri        string
		wantValue  string
		wantStatus int
		wantErr    bool
	}{
		{
			name:       "common correct case",
			uri:        "/value/gauge/Mallocs",
			wantValue:  "300",
			wantStatus: 200,
		},
		{
			name:       "common correct case",
			uri:        "/value/counter/Test",
			wantValue:  "1000",
			wantStatus: 200,
		},
		{
			name:       "wrong type",
			uri:        "/value/something/Mallocs",
			wantValue:  "Not Found",
			wantStatus: 404,
		},
	}

	for _, tt := range testSet {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, tt.uri, nil)
		h.router.ServeHTTP(w, req)

		assert.Equal(t, tt.wantStatus, w.Code)
	}

	for _, tt := range testGet {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, tt.uri, nil)
		h.router.ServeHTTP(w, req)

		assert.Equal(t, tt.wantStatus, w.Code)
		assert.Equal(t, tt.wantValue, w.Body.String())
	}
}
