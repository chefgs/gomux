//API for getting SCM event payload requests
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

/* func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
} */

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1/").Subrouter()
	api.HandleFunc("", get).Methods("GET")
	api.HandleFunc("", post).Methods("POST")
	/* 	api.HandleFunc("", put).Methods(http.MethodPut)
	   	api.HandleFunc("", delete).Methods(http.MethodDelete) */

	http.ListenAndServe(":9080", router)
}
