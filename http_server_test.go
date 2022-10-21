package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// When testing an HTTP server, we need only test the handlers or handler
// functions.
func TestUserHandler(t *testing.T) {
	// Define httptest test cases
	//
	// For example tests, see:
	//   * https://go.dev/src/net/http/httptest/example_test.go
	tt := []struct {
		Name       string
		Method     string
		StatusCode int
		Body       string
	}{
		{
			Name:       "success",
			Method:     http.MethodGet,
			StatusCode: http.StatusOK,
			Body:       `{"login":"0x4e0x4b","id":1,"type":"User"}`,
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			// Create a new incoming server Request, suitable for passing to an
			// http.Handler for testing.
			req := httptest.NewRequest("GET", "/users/0", nil)
			// Create an initialized ResponseRecorder. A ResponseRecorder is an
			// implementation of http.ResponseWriter that records its mutations for
			// later inspection in tests.
			w := httptest.NewRecorder()
			// The ResponseRecorder and Request are passed to the Handler under test.
			UserHandler{}.ServeHTTP(w, req)
			resp := w.Result()
			body, err := io.ReadAll(resp.Body)
			assert.Nil(t, err)
			assert.Equal(t, tc.StatusCode, resp.StatusCode)
			assert.Equal(t, tc.Body, strings.TrimSpace(string(body)))
		})
	}
}
