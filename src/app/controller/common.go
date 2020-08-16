package controller

import (
	"encoding/json"
	"net/http"

	"github.com/davyzhang/agw"
)

func sendJson(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		agw.WriteResponse(w, []byte(err.Error()), false)
	}
	w.WriteHeader(status)
	agw.WriteResponse(w, []byte(response), false)
}
