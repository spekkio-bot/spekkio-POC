package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/spekkio-bot/spekkio/src/app/model"
)

// Scrumify sets up a GitHub repository's Issues to facilitate scrum-driven development.
func Scrumify(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var req model.ScrumifyRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res := model.Error{
			Message: "What do you want?",
			Error: err.Error(),
		}
		sendJson(w, http.StatusBadRequest, res)
		return
	}

	if req.RepoID == "" {
		res := model.Error{
			Message: "No cheating!",
			Error: "request body is missing repo_id property",
		}
		sendJson(w, http.StatusBadRequest, res)
		return
	}

	if req.Token == "" {
		res := model.Error{
			Message: "No cheating!",
			Error: "request body is missing token property",
		}
		sendJson(w, http.StatusBadRequest, res)
		return
	}

	res := model.Ping{
		Message: "Ipso facto, meeny moe... MAGICO! Your repository was successfully scrumified!",
	}

	sendJson(w, http.StatusOK, res)
}
