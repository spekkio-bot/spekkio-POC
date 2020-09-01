package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/spekkio-bot/spekkio/src/app/model"
)

const SCRUMIFY_QUERY = "queries/get_scrumify_labels.sql"

// Scrumify sets up a GitHub repository's Issues to facilitate scrum-driven development.
func Scrumify(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var req model.ScrumifyRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res := model.Error{
			Message: "What do you want?",
			Error:   err.Error(),
		}
		sendJson(w, http.StatusBadRequest, res)
		return
	}

	if req.RepoID == "" {
		res := model.Error{
			Message: "No cheating!",
			Error:   "request body is missing repo_id property",
		}
		sendJson(w, http.StatusBadRequest, res)
		return
	}

	if req.Token == "" {
		res := model.Error{
			Message: "No cheating!",
			Error:   "request body is missing token property",
		}
		sendJson(w, http.StatusBadRequest, res)
		return
	}

	var query string
	query, err = getSqlFrom(SCRUMIFY_QUERY)
	if err != nil {
		res := model.Error{
			Message: "GRRRR... That was most embarrassing!",
			Error:   err.Error(),
		}
		sendJson(w, http.StatusInternalServerError, res)
		return
	}

	var results *sql.Rows
	results, err = db.Query(query)

	if err != nil {
		res := model.Error{
			Message: "GRRRR... That was most embarrassing!",
			Error:   err.Error(),
		}
		sendJson(w, http.StatusInternalServerError, res)
		return
	}
	defer results.Close()

	var labels []model.ScrumifyQueryResponse
	for results.Next() {
		var label model.ScrumifyQueryResponse
		err = results.Scan(&label.ID,
			&label.Name,
			&label.Color,
			&label.Description)
		if err != nil {
			res := model.Error{
				Message: "GRRRR... That was most embarrassing!",
				Error:   err.Error(),
			}
			sendJson(w, http.StatusInternalServerError, res)
			return
		}
		labels = append(labels, label)
	}

	res := model.Ping{
		Message: "Ipso facto, meeny moe... MAGICO! Your repository was successfully scrumified!",
	}

	sendJson(w, http.StatusOK, res)
}
