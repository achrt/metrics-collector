package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/achrt/metrics-collector/internal/domain/models"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func Test_Handlers_UpdateGet_Metrics(t *testing.T) {
	updRoute := "/update/"
	wrongType := "something"

	fl := float64(1)
	i := int64(50)

	testSet := []struct {
		name       string
		value      models.Metrics
		wantStatus int
	}{
		{
			name:       "common correct case",
			value:      models.Metrics{ID: "Mallocs", MType: models.TypeGauge, Value: &fl},
			wantStatus: 200,
		},
		{
			name:       "common correct case counter",
			value:      models.Metrics{ID: "Count", MType: models.TypeCounter, Delta: &i},
			wantStatus: 200,
		},
		{
			name:       "common correct case counter",
			value:      models.Metrics{ID: "Count", MType: models.TypeCounter, Delta: &i},
			wantStatus: 200,
		},
		{
			name:       "missed metric name",
			value:      models.Metrics{ID: "", MType: models.TypeCounter, Delta: &i},
			wantStatus: 400,
		},
		{
			name:       "missed metric value",
			value:      models.Metrics{ID: "Count", MType: models.TypeCounter},
			wantStatus: 400,
		},
		{
			name:       "wrong type",
			value:      models.Metrics{ID: "Mallocs", MType: wrongType, Value: &fl},
			wantStatus: 200,
		},
	}

	testGet := []struct {
		name       string
		uri        string
		body       models.Metrics
		wantValue  string
		wantStatus int
		wantErr    bool
	}{
		{
			name:       "common correct case",
			uri:        "/value/",
			body:       models.Metrics{ID: "Mallocs", MType: models.TypeGauge},
			wantValue:  `{"id":"Mallocs","type":"gauge","value":1}`,
			wantStatus: 200,
		},
		{
			name:       "common correct case counter",
			uri:        "/value/",
			body:       models.Metrics{ID: "Count", MType: models.TypeCounter},
			wantValue:  `{"id":"Count","type":"counter","delta":100}`,
			wantStatus: 200,
		},
		{
			name:       "wrong type",
			uri:        "/value/",
			body:       models.Metrics{ID: "Wrong", MType: wrongType},
			wantValue:  `err`,
			wantStatus: 404,
		},
	}

	for _, tt := range testSet {
		b, err := json.Marshal(tt.value)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, updRoute, bytes.NewBuffer(b))
		h.router.ServeHTTP(w, req)

		assert.Equal(t, tt.wantStatus, w.Code)
	}

	for _, tt := range testGet {
		b, err := json.Marshal(tt.body)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, tt.uri, bytes.NewBuffer(b))
		h.router.ServeHTTP(w, req)

		assert.Equal(t, tt.wantStatus, w.Code)
		assert.Equal(t, tt.wantValue, w.Body.String())
	}
}
