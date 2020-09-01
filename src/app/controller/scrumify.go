package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spekkio-bot/spekkio/src/app/model"
	"github.com/spekkio-bot/spekkio/src/queries/graphql"
)

const SCRUMIFY_QUERY = "queries/sql/get_scrumify_labels.sql"

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

	mutations := make([]interface{}, len(labels))

	for id, label := range labels {
		mutationInputs := []gqlbuilder.MutationInput{}
		mutationInputs = append(mutationInputs, gqlbuilder.MutationInput{
			Key:   "repositoryId",
			Value: req.RepoID,
		})
		mutationInputs = append(mutationInputs, gqlbuilder.MutationInput{
			Key:   "name",
			Value: label.Name,
		})
		mutationInputs = append(mutationInputs, gqlbuilder.MutationInput{
			Key:   "color",
			Value: label.Color,
		})
		mutationInputs = append(mutationInputs, gqlbuilder.MutationInput{
			Key:   "description",
			Value: label.Description,
		})
		mutation := &gqlbuilder.Mutation{
			Name:   "createLabel",
			Alias:  fmt.Sprintf("cl%d", id),
			Inputs: mutationInputs,
			Return: []string{"clientMutationId"},
		}
		mutations[id] = *mutation
	}

	gql := &gqlbuilder.Operation{
		Name:       "Scrumify",
		Type:       "mutation",
		Operations: mutations,
	}

	gqlQuery, err := gql.Build()
	if err != nil {
		res := model.Error{
			Message: "GRRRR... That was most embarrassing!",
			Error:   err.Error(),
		}
		sendJson(w, http.StatusInternalServerError, res)
		return
	}

	var apiReq *http.Request
	var apiResp *http.Response
	apiClient := &http.Client{}
	tokenHeader := fmt.Sprintf("bearer %s", req.Token)

	apiReq, err = http.NewRequest("POST", GRAPHQL_API, gqlQuery)
	apiReq.Header.Add("Authorization", tokenHeader)
	apiReq.Header.Add("Accept", LABEL_PREVIEW_HEADER)
	apiReq.Header.Add("Content-Type", "application/json")

	apiResp, err = apiClient.Do(apiReq)
	if err != nil {
		res := model.Error{
			Message: "GRRRR... That was most embarrassing!",
			Error:   err.Error(),
		}
		sendJson(w, http.StatusInternalServerError, res)
		return
	}

	// TODO: handle various api responses
	fmt.Println(apiResp.Status)

	res := model.Ping{
		Message: "Ipso facto, meeny moe... MAGICO! Your repository was successfully scrumified!",
	}

	sendJson(w, http.StatusOK, res)
}
