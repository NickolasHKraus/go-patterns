package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetFooURL(t *testing.T) {
	t.Run("get foo url", func(t *testing.T) {
		c := FooAPIClient{}
		exp := (&url.URL{
			Scheme: "https",
			Host:   "api.foo.com",
			Path:   "/v1/users/0",
		}).String()
		ret := c.GetFooURL("https", "api.foo.com", "/v1/users/0")
		if ret != exp {
			t.Errorf("incorrect url.\nExpected: %s\nGot: %5s%s", exp, " ", ret)
		}
	})
}

type MockHTTPClient struct {
	// MockGetFunc is a field on the MockHTTPClient struct that holds the
	// function to be called by the `Get` method.
	MockGetFunc func(url string) (resp *http.Response, err error)
}

func (m MockHTTPClient) Get(url string) (resp *http.Response, err error) {
	return m.MockGetFunc(url)
}

func TestGetFooData(t *testing.T) {
	// Since we are attempting to test the scenario in which the client fails to
	// make a connection with the server, we forgo using the test server and
	// instead create a mock HTTP client, which simply returns an error for its
	// Get method.
	t.Run("connection refused", func(t *testing.T) {
		mockHTTPClient := &MockHTTPClient{
			MockGetFunc: func(url string) (resp *http.Response, err error) {
				return nil, fmt.Errorf("connection refused")
			},
		}
		f := FooAPIClient{
			Client: mockHTTPClient,
		}
		resp, err := f.GetFooData("")
		assert.Nil(t, resp)
		assert.EqualError(t, err, "connection refused")
	})
	// Define httptest test cases
	cases := []struct {
		Name       string
		StatusCode int
		Header     http.Header
		Body       string
		WantData   *FooResponseData
		WantErr    error
	}{
		{
			Name:       "unexpected status code",
			StatusCode: http.StatusInternalServerError,
			Header:     nil,
			Body:       "",
			WantData:   nil,
			WantErr:    fmt.Errorf("unexpected status code: 500"),
		},
		{
			Name:       "unable to read response body",
			StatusCode: http.StatusOK,
			// This is a cheeky way of causing io.ReadAll to return an error without
			// having to create a custom io.Reader.
			//
			// The handler is instructed that the body has a length of 1 byte,
			// however, the response body is empty.
			Header:   map[string][]string{"Content-Length": {"1"}},
			Body:     "",
			WantData: nil,
			WantErr:  fmt.Errorf("unable to read response body: unexpected EOF"),
		},
		{
			Name:       "unable to unmarshal response JSON",
			StatusCode: http.StatusOK,
			Header:     nil,
			Body:       "",
			WantData:   nil,
			WantErr:    fmt.Errorf("unable to unmarshal response JSON: unexpected end of JSON input"),
		},
		{
			Name:       "success",
			StatusCode: http.StatusOK,
			Header:     nil,
			Body: `{"login": "0x4e0x4b",
"id": 1,
"type": "User"
}`,
			WantData: &FooResponseData{
				Login: "0x4e0x4b",
				ID:    1,
				Type:  "User",
			},
			WantErr: nil,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			// A Server is an HTTP server listening on a system-chosen port on the
			// local loopback interface, for use in end-to-end HTTP tests.
			//
			// The URL of the HTTP server is of the form http://ipaddr:port with no
			// trailing slash.
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method != "GET" {
					t.Errorf("Expected request method 'GET', got: %s", req.Method)
				}
				for k, vs := range tc.Header {
					for _, v := range vs {
						w.Header().Set(k, v)
					}
				}
				w.WriteHeader(tc.StatusCode)
				w.Write([]byte(tc.Body))
			}))
			defer ts.Close()
			f := FooAPIClient{
				Client: &http.Client{},
			}
			// It should be noted that HTTP requests must be made to the URL of the
			// test server.
			data, err := f.GetFooData(ts.URL)
			if err != nil {
				assert.EqualError(t, err, tc.WantErr.Error())
			}
			assert.Equal(t, tc.WantData, data)
		})
	}
}

// See http_client_mock.go for mock FooAPI implmentation.
//
// This code is generated using gomock:
//
//	$ mockgen -package main -source ... -destination ...
//
// See: https://github.com/golang/mock
func TestDoSomething(t *testing.T) {
	t.Run("do something", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		// Assert that methods are invoked.
		defer ctrl.Finish()

		m := NewMockFooAPI(ctrl)

		// Assert that GetFooURL() is called with expected arguments.
		// Anything else will fail.
		expURL := "https://api.foo.com/v1/users/0"
		m.
			EXPECT().
			GetFooURL(gomock.Eq("https"), gomock.Eq("api.foo.com"), gomock.Eq("/v1/users/0")).
			Return(expURL)

		// Assert that GetFooData() is called with expected arguments.
		// Anything else will fail.
		expData := &FooResponseData{
			Login: "0x4e0x4b",
			ID:    1,
			Type:  "User",
		}
		m.
			EXPECT().
			GetFooData(gomock.Eq("https://api.foo.com/v1/users/0")).
			Return(expData, nil)

		retURL, retData := DoSomething(m)
		assert.Equal(t, expURL, retURL)
		assert.Equal(t, expData, retData)
	})
}
