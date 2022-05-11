package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_Handlers_UpdateGet_Metrics(t *testing.T) {
	updRoute := "/update"

	testSet := []struct {
		name       string
		value      string
		wantStatus int
	}{
		{
			name:       "common correct case",
			//value: ,
			wantStatus: 200,
		},
		{
			name:       "common correct case counter",
			wantStatus: 200,
		},
		{
			name:       "common correct case counter",
			wantStatus: 200,
		},
		{
			name:       "missed metric name",
			wantStatus: 404,
		},
		{
			name:       "missed metric value",
			wantStatus: 404,
		},
		{
			name:       "wrong type",
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
		req, _ := http.NewRequest(http.MethodPost, updRoute, nil)
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
