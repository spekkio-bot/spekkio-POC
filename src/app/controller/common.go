package controller

import (
	"encoding/json"
	"net/http"

	"github.com/park-junha/agw"
	"github.com/spekkio-bot/spekkio/src/app/model"
)

// GRAPHQL_API is the URL to GitHub's GraphQL API
const GRAPHQL_API = "https://api.github.com/graphql"

// LABEL_PREVIEW_HEADER is an Accept header value added to an HTTP request that allows access to GitHub's GraphQL Label Preview APIs
const LABEL_PREVIEW_HEADER = "application/vnd.github.bane-preview+json"

const spekkio400 = "What do you want?"
const spekkio500 = "GRRRR... That was most embarrassing!"

func sendJson(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		agw.WriteResponse(w, []byte(err.Error()), false)
		return
	}
	w.WriteHeader(status)
	agw.WriteResponse(w, []byte(response), false)
}

func send400(w http.ResponseWriter, err error) {
	res := model.Error{
		Message: spekkio400,
		Error:   err.Error(),
	}
	sendJson(w, http.StatusBadRequest, res)
}

func send500(w http.ResponseWriter, err error) {
	res := model.Error{
		Message: spekkio500,
		Error:   err.Error(),
	}
	sendJson(w, http.StatusInternalServerError, res)
}
