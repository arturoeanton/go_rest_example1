package main

import (
	"encoding/json"
	"encoding/xml"
	"html/template"
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
	io.WriteString(w, `{"alive": true}`)
}
func reqJSONName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	json.NewEncoder(w).Encode(Person{Name: vars["name"]})
}
func reqHTML(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template.templ"))
	_ = t.Execute(w, Person{Name: "tim"})
}
func reqJSON(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Person{Name: "tim"})
}
func reqXML(w http.ResponseWriter, r *http.Request) {
	xml.NewEncoder(w).Encode(Person{Name: "tim"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", produces("json", healthCheckHandler))
	r.HandleFunc("/", produces("html", reqHTML))
	r.HandleFunc("/html", produces("html", reqHTML))
	r.HandleFunc("/json", produces("json", reqJSON))
	r.HandleFunc("/json/{name}", produces("json", reqJSONName))
	r.HandleFunc("/xml", produces("xml", reqXML)).Methods(http.MethodGet)
	r.Use(middleware)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s [%s]\n", r.RemoteAddr, r.Method, r.URL, time.Since(start).String())
	})
}

func produces(format string, f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	value := format
	if format == "json" {
		value = "application/json"
	}
	if format == "xml" {
		value = "application/xml"
	}
	if format == "html" {
		value = "text/html; charset=utf-8"
	}
	if format == "text" {
		value = "text/plain; charset=utf-8"
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", value)
		w.WriteHeader(http.StatusOK)
		f(w, r)
	}
}
