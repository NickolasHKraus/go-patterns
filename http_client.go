package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Making HTTP requests via an HTTP client.
//
// HTTP requests should leverage the HTTP client provided by the net/http
// package.
//
// See: https://pkg.go.dev/net/http
//
// This example offers a "belt and suspenders" method for defining a robust
// HTTP client that is flexible and easily testable.

// Though requests made by an HTTP client can be tested using the standard
// library httptest package, more often than not, applications making HTTP
// requests will do more than simply return the response from the server.
//
// Using an interface, we can define a set of method signatures for making HTTP
// requests to an external API and handling the response data. This allows for
// the flexibility to mock certain functions when testing.
type FooAPI interface {
	GetFooURL(string, string, string) string
	GetFooData(string) (*FooResponseData, error)
}

// HTTPClient is a common interface that specifies the method signatures on the
// http.Client struct that are to be mocked.
//
// See: net/http/client.go
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

// FooAPIClient is an application-specific HTTP client. It implements the
// FooAPI interface. It also holds an http.Client, which can be overridden for
// testing purposes.
type FooAPIClient struct {
	Client HTTPClient
}

// In almost all cases, an external API will return a response as JSON:
//
//	>The MIME media type for JSON text is application/json. The default
//	 encoding is UTF-8. (Source: RFC 4627)
//
//	 See: https://www.ietf.org/rfc/rfc4627.txt
//
// Likewise, the application should have some inkling about the structure of
// the returned data. For this reason, it is almost always recommended to use
// a struct for unmarshaling a JSON response. If the response is not
// unmarshalled into a struct, you will need to use type assertion when
// accessing an interface value's underlying concrete value.
type FooResponseData struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
	Type  string `json:"type"`
}

func (f FooAPIClient) GetFooURL(scheme string, host string, path string) string {
	return (&url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}).String()
}

func (f FooAPIClient) GetFooData(url string) (*FooResponseData, error) {
	resp, err := f.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %w", err)
	}
	var r FooResponseData
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("unable to unmarshal response JSON: %w", err)
	}
	return &r, nil
}

// This method demonstrates how a function might be passed a value that
// implements the FooAPI interface. A type implementing the FooAPI interface
// can be easily mocked and passed to the function when testing.
func DoSomething(f FooAPI) (string, *FooResponseData) {
	url := f.GetFooURL("https", "api.foo.com", "/v1/users/0")
	resp, err := f.GetFooData(url)
	if err != nil {
		fmt.Printf("An error occurred: %s", err)
	}
	return url, resp
}

// Example main() function.
//
// See: https://go.dev/ref/spec#Program_execution
//
// func main() {
// 	f := FooAPIClient{
// 		Client: &http.Client{},
// 	}
// 	DoSomething(f)
// }

// Example initializer function.
//
// func InitFooAPIClient() *FooAPIClient {
// 	return &FooAPIClient{
// 		Client: &http.Client{},
// 	}
// }
