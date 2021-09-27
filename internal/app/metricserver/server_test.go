package metricserver

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_CustomMetrics(t *testing.T) {
	config := NewConfig()
	srv := newServer(config)

	testCases := []struct {
		name           string
		path           string
		expectedMetric string
		expectedCode   int
	}{
		{
			name:           "valid_files_count",
			path:           config.MetricsPath,
			expectedMetric: "files_count",
			expectedCode:   200,
		},
		{
			name:           "valid_files_size_total",
			path:           config.MetricsPath,
			expectedMetric: "files_size_total",
			expectedCode:   200,
		},
		{
			name:           "unvalid_path",
			path:           "/unvalid_path",
			expectedMetric: "",
			expectedCode:   404,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			req, _ := http.NewRequest(http.MethodGet, tc.path, b)
			srv.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tc.expectedMetric)
		})
	}
}
