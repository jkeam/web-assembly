package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stealthrocket/net/wasip1" // needed to use that Listen instead of net.Listen
	"net/http"
	"time"
)

// Handler functions
func internalServerErrorHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	jsonBytes, errMarshal := json.Marshal(&ErrorResponse{Message: "Internal Server Error"})
	if errMarshal != nil {
		return
	}
	w.Write(jsonBytes)
}
func successHandler(w http.ResponseWriter, jsonBytes []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// Objects
type ErrorResponse struct {
	Message string `json:"message"`
}
type HelloResponse struct {
	Message string `json:"message"`
}
type SumRequest struct {
	Operand1 int64 `json:"operand1"`
	Operand2 int64 `json:"operand2"`
}
type SumResponse struct {
	Operand1 int64 `json:"operand1"`
	Operand2 int64 `json:"operand2"`
	Result   int64 `json:"result"`
}

// Home handler
type HomeHandler struct{}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	jsonBytes, errMarshal := json.Marshal(HelloResponse{Message: "Hi!"})
	if errMarshal != nil {
		internalServerErrorHandler(w)
		return
	}
	successHandler(w, jsonBytes)
}

// Add handler
type AddHandler struct{}

func sum(w http.ResponseWriter, r *http.Request) {
	var sumRequest SumRequest
	if err := json.NewDecoder(r.Body).Decode(&sumRequest); err != nil {
		internalServerErrorHandler(w)
		return
	}

	sumResponse := &SumResponse{
		Operand1: sumRequest.Operand1,
		Operand2: sumRequest.Operand2,
		Result:   sumRequest.Operand1 + sumRequest.Operand2,
	}
	jsonBytes, errMarshal := json.Marshal(sumResponse)
	if errMarshal != nil {
		internalServerErrorHandler(w)
		return
	}
	successHandler(w, jsonBytes)
}
func (h *AddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost:
		sum(w, r)
		return
	default:
		return
	}
}

func listenAndServe(port string, mux *http.ServeMux) {
	listener, err := wasip1.Listen("tcp4", port)
	if err != nil {
		panic(err)
		return
	}

	server := &http.Server{
		ReadHeaderTimeout: 3 * time.Second,
		IdleTimeout:       3 * time.Second,
		Handler:           mux,
	}
	err = server.Serve(listener)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
		return
	}
	fmt.Println("Server stopped listening")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &HomeHandler{})
	mux.Handle("/add", &AddHandler{})

	port := ":8080"
	fmt.Printf("Server started on %s", port)
	listenAndServe(port, mux)
}
