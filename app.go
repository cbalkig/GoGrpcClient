package main

import (
	"log"
	"net/http"

	"WebServiceDockerDemo/grpc"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	text := grpcCaller.Call()
	w.Write([]byte(`{"message": "` + text + `"}`))
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}