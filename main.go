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

// Results contains the results structure.
type Results struct {
	Pid    string `json:"pid"`
	Result int    `json:"result"`
}

// Used by the main application
var (
	Version string
	Build   string
)

// Messages used by http.HandlerFunc functions.
var (
	HomeMessage    = Message{"Masterbuilder serverless API"}
	HelloMessage   = Message{"Hello, world!"}
	GoodbyeMessage = Message{"Goodbye, world!"}
)

// RootHandler is a http.HandlerFunc for the / path.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(HomeMessage)
}

// VoteHandlerP1 handlers the vote for participant #1
func VoteHandlerP1(w http.ResponseWriter, r *http.Request) {
	// Get p1 detail from the DynamoDB table and increment count then write the result
	// var TotalP1 = 0
	// TotalP1++
	// Don't forget error management (discard if cannot write, but for production app do somne more sophisticated stuff)
	ResponseMessage := Message{"Added vote for participant P1"}
	json.NewEncoder(w).Encode(ResponseMessage)
}

// VoteHandlerP2 handlers the vote for participant #2
func VoteHandlerP2(w http.ResponseWriter, r *http.Request) {
	// Get p2 detail from the DynamoDB table and increment count then write the result
	// var TotalP2 = 0
	// TotalP2++
	ResponseMessage := Message{"Added vote for participant P2"}
	json.NewEncoder(w).Encode(ResponseMessage)
}

// VoteHandlerP3 handlers the vote for participant #3
func VoteHandlerP3(w http.ResponseWriter, r *http.Request) {
	// Get p3 detail from the DynamoDB table and increment count then write the result
	// var TotalP3 = 0
	// TotalP3++
	ResponseMessage := Message{"Added vote for participant P3"}
	json.NewEncoder(w).Encode(ResponseMessage)
}

// ResultsHandler returns the results of the votes for all participants
func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	// Get results for all participants and return to the caller
	var TotalP1, TotalP2, TotalP3 int
	// For offline dev, let's set some values
	TotalP1 = 5403
	TotalP2 = 6603
	TotalP3 = 10
	var RetResults []Results
	RetResults = append(RetResults, Results{"Participant 1", TotalP1})
	RetResults = append(RetResults, Results{"Participant 2", TotalP2})
	RetResults = append(RetResults, Results{"Participant 3", TotalP3})
	json.NewEncoder(w).Encode(RetResults)
}

// VersionHandler returns the version of the api
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	var BMsg = Build[:5]
	var VersionMessage = "Version: " + Version + " - Build: " + BMsg
	json.NewEncoder(w).Encode(VersionMessage)
}

// ClearHandler clears the dynamobdb table
func ClearHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: clear the DynamoDB table
}

// RegisterRoutes registers the API's routes.
func RegisterRoutes() {
	http.Handle("/", h(RootHandler))
	http.Handle("/p1", h(VoteHandlerP1))
	http.Handle("/p2", h(VoteHandlerP2))
	http.Handle("/p3", h(VoteHandlerP3))
	http.Handle("/results", h(ResultsHandler))
	http.Handle("/version", h(VersionHandler))
	http.Handle("/cleartable", h(ClearHandler))
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
