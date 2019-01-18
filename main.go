package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/apex/gateway"
)

// ContentType is the Content-Type header set in responses.
const ContentType = "application/json; charset=utf8"

// Message contains a simple message response.
type Message struct {
	Message string `json:"message"`
}

// Messages used by http.HandlerFunc functions.
var (
	WelcomeMessage = Message{"Welcome to the AWeSome SAM golang example API!"}
	HelloMessage   = Message{"Hello, world!"}
	GoodbyeMessage = Message{"Goodbye, world!"}
)

// RootHandler is a http.HandlerFunc for the / path.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(WelcomeMessage)
}

// Functions handling the voting
func VoteHandlerP1(w http.ResponseWriter, r *http.Request) {

}

func VoteHandlerP2(w http.ResponseWriter, r *http.Request) {

}

func VoteHandlerP3(w http.ResponseWriter, r *http.Request) {

}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {

}

func VersionHandler(w http.ResponseWriter, r *http.Request) {

}

// RegisterRoutes registers the API's routes.
func RegisterRoutes() {
	http.Handle("/", h(RootHandler))
	http.Handle("/p1", h(VoteHandlerP1))
	http.Handle("/p2", h(VoteHandlerP2))
	http.Handle("/p3", h(VoteHandlerP3))
	http.Handle("/results", h(ResultsHandler))
	http.Handle("/version", h(VersionHandler))
}

// h wraps a http.HandlerFunc and adds common headers.
func h(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		next.ServeHTTP(w, r)
	})
}

func main() {
	RegisterRoutes()
	log.Fatal(gateway.ListenAndServe(":3000", nil))
}
