package controller

import (
	"net/http"

	"github.com/spekkio-bot/spekkio/src/app/model"
)

// NotFound is called when a non-existent resource is requested.
func NotFound(w http.ResponseWriter, r *http.Request) {
	res := model.Error{
		Message: "Resource not found.",
	}

	sendJson(w, http.StatusNotFound, res)
}
