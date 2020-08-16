package controller

import (
	"net/http"

	"github.com/spekkio-bot/spekkio/src/app/model"
)

// Ping is called when a request is made to the index route of the app.
func Ping(w http.ResponseWriter, r *http.Request) {
	res := model.Ping{
		Message: "Request successful.",
	}

	sendJson(w, http.StatusOK, res)
}
