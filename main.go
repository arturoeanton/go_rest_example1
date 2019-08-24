package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Person is struct of test
type Person struct {
	Name string `xml:"name" json:"name"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}
func reqJSONName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Person{Name: vars["name"]})
}
func reqHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "<html><head><title>test</title></head><body><h1>Titulo</h1></body></html>")
}
func reqJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Person{Name: "tim"})
}
func reqXML(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	xml.NewEncoder(w).Encode(Person{Name: "tim"})
}

// only for demo
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s [%s]\n", r.RemoteAddr, r.Method, r.URL, time.Since(start).String())
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthCheckHandler)
	r.HandleFunc("/", reqHTML)
	r.HandleFunc("/html", reqHTML)
	r.HandleFunc("/json", reqJSON)
	r.HandleFunc("/json/{name}", reqJSONName)
	r.HandleFunc("/xml", reqXML).Methods(http.MethodGet)
	r.Use(middleware)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
