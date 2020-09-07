package controller

import (
	"net/http"
)

// NotFound is called when a non-existent resource is requested.
func NotFound(w http.ResponseWriter, r *http.Request) {
	send404(w)
}
