package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/davyzhang/agw"
)

const GRAPHQL_API = "https://api.github.com/graphql"
const LABEL_PREVIEW_HEADER = "application/vnd.github.bane-preview+json"

func sendJson(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		agw.WriteResponse(w, []byte(err.Error()), false)
	}
	w.WriteHeader(status)
	agw.WriteResponse(w, []byte(response), false)
}

func getSqlFrom(file string) (string, error) {
	sql, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(sql), nil
}
