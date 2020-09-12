package controller

import (
	"net/http"
)

// MethodNotAllowed is called when a disallowed resource method is requested.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	send405(w)
}
