package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateMetricHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name string
		url  string
		want want
	}{
		{
			name: "Success update gauge (Int)",
			url:  "/update/gauge/testName/300",
			want: want{
				code:        200,
				response:    "",
				contentType: "text/plain; charset=utf-8",
			},
		}, {
			name: "Success update gauge (Float)",
			url:  "/update/gauge/testName/300.30",
			want: want{
				code:        200,
				response:    "",
				contentType: "text/plain; charset=utf-8",
			},
		}, {
			name: "Success update counter",
			url:  "/update/counter/testName/100",
			want: want{
				code:        200,
				response:    "",
				contentType: "text/plain; charset=utf-8",
			},
		}, {
			name: "Empty metric type",
			url:  "/update/",
			want: want{
				code:        404,
				response:    "Request without metric type\n",
				contentType: "text/plain; charset=utf-8",
			},
		}, {
			name: "No metric type",
			url:  "/update",
			want: want{
				code:        404,
				response:    "Request without metric type\n",
				contentType: "text/plain; charset=utf-8",
			},
		}, {
			name: "Empty metric name",
			url:  "/update/counter/",
			want: want{
				code:        404,
				response:    "Request without metric name\n",
				contentType: "text/plain; charset=utf-8",
			},
		}, {
			name: "No metric name",
			url:  "/update/counter",
			want: want{
				code:        404,
				response:    "Request without metric name\n",
				contentType: "text/plain; charset=utf-8",
			},
		}, {
			name: "Empty metric value",
			url:  "/update/counter/testName/",
			want: want{
				code:        404,
				response:    "Request without metric value\n",
				contentType: "text/plain; charset=utf-8",
			},
		}, {
			name: "No metric value",
			url:  "/update/counter/testName",
			want: want{
				code:        404,
				response:    "Request without metric value\n",
				contentType: "text/plain; charset=utf-8",
			},
		}, {
			name: "More parameters",
			url:  "/update/counter/testName/300/more",
			want: want{
				code:        400,
				response:    "Invalid request\n",
				contentType: "text/plain; charset=utf-8",
			},
		}, {
			name: "Invalid metric type",
			url:  "/update/invalidMetricType/testName/300",
			want: want{
				code:        400,
				response:    "Invalid metric type\n",
				contentType: "text/plain; charset=utf-8",
			},
		}, {
			name: "Invalid metric value",
			url:  "/update/counter/testName/notNumber",
			want: want{
				code:        400,
				response:    "Invalid metric value\n",
				contentType: "text/plain; charset=utf-8",
			},
		}, {
			name: "Invalid metric value (float to counter)",
			url:  "/update/counter/testName/303.30",
			want: want{
				code:        400,
				response:    "Invalid metric value\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, test.url, nil)
			w := httptest.NewRecorder()

			UpdateMetricHandler(w, r)

			res := w.Result()
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Equal(t, test.want.code, res.StatusCode)
			assert.Equal(t, test.want.response, string(resBody))
			assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
