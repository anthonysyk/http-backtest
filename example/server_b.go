package example

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

func runServerWithSeed(seed string) *httptest.Server {
	mock := http.NewServeMux()
	server := httptest.NewServer(mock)

	mock.HandleFunc("/diverge/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		item := struct {
			ID   string `json:"id"`
			Seed string `json:"seed"`
		}{ID: mux.Vars(r)["id"]}

		b, _ := json.Marshal(item)
		_, _ = w.Write(b)
	})

	mock.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		if query == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(query))
	})

	mock.HandleFunc("/doc/{id}", func(w http.ResponseWriter, r *http.Request) {
		query, exists := mux.Vars(r)["id"]
		if !exists {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(query))
	})

	fmt.Println(server.URL)
	return server
}
