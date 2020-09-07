package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/spekkio-bot/spekkio/src/app/model"
	"github.com/spekkio-bot/spekkio/src/queries/graphql"
	"github.com/spekkio-bot/spekkio/src/queries/sql"
)

// Scrumify sets up a GitHub repository's Issues to facilitate scrum-driven development.
func Scrumify(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var req model.ScrumifyRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		send400(w, err)
		return
	}

	if req.RepoID == "" {
		send400(w, errors.New("request body is missing repo_id property"))
		return
	}

	if req.Token == "" {
		send400(w, errors.New("request body is missing token property"))
		return
	}

	var query string
	scrumifyQueryProps := &sqlbuilder.SelectQueryProps{
		BaseTable: "ScrumifyLabels",
		Columns: []sqlbuilder.Column{
			{
				Name:  "id",
				Alias: "",
			},
			{
				Name:  "name",
				Alias: "",
			},
			{
				Name:  "color",
				Alias: "",
			},
			{
				Name:  "description",
				Alias: "",
			},
		},
	}
	query, err = scrumifyQueryProps.BuildQuery()
	if err != nil {
		send500(w, err)
		return
	}

	var results *sql.Rows
	results, err = db.Query(query)

	if err != nil {
		send500(w, err)
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

	mutations := []gqlbuilder.Mutation{}

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
		mutations = append(mutations, *mutation)
	}

	gql := &gqlbuilder.Mutations{
		Name:      "Scrumify",
		Mutations: mutations,
	}

	gqlQuery, err := gql.Build()
	if err != nil {
		send500(w, err)
		return
	}

	var apiReq *http.Request
	var apiResp *http.Response
	apiClient := &http.Client{}
	headers := make(map[string][]string)
	headers["Authorization"] = []string{
		fmt.Sprintf("bearer %s", req.Token),
	}
	headers["Accept"] = []string{
		LABEL_PREVIEW_HEADER,
	}
	headers["Content-Type"] = []string{
		"application/json",
	}

	apiReq, err = initGraphqlRequest(gqlQuery, headers)
	if err != nil {
		send500(w, err)
		return
	}

	apiResp, err = apiClient.Do(apiReq)
	if err != nil {
		send500(w, err)
		return
	}

	// TODO: handle various api responses
	fmt.Println(apiResp.Status)

	res := model.Ping{
		Message: "Ipso facto, meeny moe... MAGICO! Your repository was successfully scrumified!",
	}

	sendJson(w, http.StatusOK, res)
}
