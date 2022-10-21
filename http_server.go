package main

// Handling HTTP requests via an HTTP server.
//
// HTTP servers should use the implmentation provided by the net/http package.
//
// See: https://pkg.go.dev/net/http
//
// This example offers a pattern for building an idiomatic, testable HTTP
// server following the REST API pattern.
//
// In most cases, it is enough to simply use the default HTTP server
// implementation provided by the net/http package:
//
//   1. Define handler function:
//
//        func(http.ResponseWriter, *http.Request)
//
//   2. Register handler for a given pattern:
//
//        http.HandleFunc("/", func(http.ResponseWriter, *http.Request))
//
//   3. Listen and serve:
//
//        http.ListenAndServe(":1337", nil)
//
// This pattern uses the DefaultServeMux for the Handler under the hood.
//
// You can, however, create your own ServeMux and Server, as these types are
// exported from the package. The default configuration provides sensible
// defaults.

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
	Type  string `json:"type"`
}

// A Handler responds to an HTTP request. It is an interface that defines a
// single method ServeHTTP.
//
// ServeHTTP should write reply headers and data to the ResponseWriter and then
// return.
//
// Handlers are associated with a pattern (ex. /users) and registered with a
// ServeMux using the Handle() method. The ServerMux matches the URL of each
// incoming request against a list of registered patterns and calls the handler
// for the pattern that most closely matches the URL.
//
// In typical RESTful design, resources should be grouped by noun, that is,
// each resource should have its own handler. It should be noted that you can
// only have one handler per pattern.
//
// See: https://pkg.go.dev/net/http#Handler
type UserHandler struct{}

// A ResponseWriter interface is used by an HTTP handler to construct an HTTP
// response.
//
// A Request represents an HTTP request received by a server or to be sent by
// a client.
func (h UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&User{
		Login: "0x4e0x4b",
		ID:    1,
		Type:  "User",
	})
}

// Example main() function.
//
// I have found that if you want to define your handlers and/or multiplexer
// in their own package, then you should return the ServeMux NOT the Server.
//
// The ServeMux comprises the logic for handling incoming HTTP requests,
// whereas the Server is responsible for low-level network configuration (ex.
// managing listeners on the local network address).
//
// See: https://go.dev/ref/spec#Program_execution
// func main() {
// 	mux := http.NewServeMux()
// 	mux.Handle("/users", userHandler{})
// 	log.Fatal(http.ListenAndServe(":8080", mux))
// }
