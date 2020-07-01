package main

import (
	languageId "GoGrpcClient/grpc/proto/clients/language_id"
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	text := getParameter(w, r, "text")
	if text != "" {
		response := languageId.Call(text)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "` + response + `"}`))
	}
}

func getParameter(w http.ResponseWriter, r *http.Request, param string) string {
	keys, ok := r.URL.Query()[param]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "Missing parameter"}`))
		return ""
	}

	key := keys[0]
	return string(key)
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
