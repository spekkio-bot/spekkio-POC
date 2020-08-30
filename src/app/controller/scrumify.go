package controller

import (
	"net/http"

	"github.com/spekkio-bot/spekkio/src/app/model"
)

// Scrumify sets up a GitHub repository's Issues to facilitate scrum-driven development.
func Scrumify(w http.ResponseWriter, r *http.Request) {
	res := model.Ping{
		Message: "Request successful.",
	}

	sendJson(w, http.StatusOK, res)
}
