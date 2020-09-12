package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/davyzhang/agw"
	"github.com/spekkio-bot/spekkio/src/app/model"
)

// GRAPHQL_API is the URL to GitHub's GraphQL API
const GRAPHQL_API = "https://api.github.com/graphql"

// LABEL_PREVIEW_HEADER is an Accept header value added to an HTTP request that allows access to GitHub's GraphQL Label Preview APIs
const LABEL_PREVIEW_HEADER = "application/vnd.github.bane-preview+json"

const spekkio400 = "No cheating! I'm watching you!"
const spekkio404 = "What do you want?"
const spekkio405 = "No cheating!"
const spekkio500 = "GRRRR... That was most embarrassing!"

func initGraphqlRequest(query io.Reader, headers map[string][]string) (*http.Request, error) {
	req, err := http.NewRequest("POST", GRAPHQL_API, query)
	if err != nil {
		return nil, err
	}
	for header, headerValues := range headers {
		for _, headerValue := range headerValues {
			req.Header.Add(header, headerValue)
		}
	}
	return req, nil
}

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

func send404(w http.ResponseWriter) {
	res := model.Error{
		Message: spekkio404,
		Error:   "resource not found.",
	}
	sendJson(w, http.StatusNotFound, res)
}

func send405(w http.ResponseWriter) {
	res := model.Error{
		Message: spekkio405,
		Error:   "method not allowed.",
	}
	sendJson(w, http.StatusMethodNotAllowed, res)
}

func send500(w http.ResponseWriter, err error) {
	res := model.Error{
		Message: spekkio500,
		Error:   err.Error(),
	}
	sendJson(w, http.StatusInternalServerError, res)
}
