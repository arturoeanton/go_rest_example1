package main

import (
	"encoding/json"
	"encoding/xml"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/arturoeanton/go_rest_example1/pkg/middleware"
	"github.com/arturoeanton/go_rest_example1/pkg/mime"
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
	t := template.Must(template.ParseFiles("templates/template.templ"))
	_ = t.Execute(w, Person{Name: "tim"})
}
func reqJSON(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Person{Name: "tim"})
}
func reqXML(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)

	xml.NewEncoder(w).Encode(Person{Name: "tim"})
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/site/").Handler(http.StripPrefix("/site/", http.FileServer(http.Dir("site/"))))
	r.PathPrefix("/webassembly/").Handler(http.StripPrefix("/webassembly/", http.FileServer(http.Dir("webassembly/"))))

	r.HandleFunc("/health", mime.Produces("json", healthCheckHandler))
	r.HandleFunc("/", mime.Produces("html", reqHTML))
	r.HandleFunc("/html", mime.Produces("html", reqHTML))
	r.HandleFunc("/json", mime.Produces("json", reqJSON))
	r.HandleFunc("/json/{name}", mime.Produces("json", reqJSONName))
	r.HandleFunc("/xml", mime.Produces("xml", reqXML)).Methods(http.MethodGet)
	r.Use(middleware.Log)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
