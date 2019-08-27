package mime

import "net/http"

// Produces set default content-type
func Produces(format string, f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
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
		f(w, r)
	}
}
