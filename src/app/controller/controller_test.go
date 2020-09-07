package controller

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Ping)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 200 {
		t.Errorf("NotFound returned wrong status code:\ngot %v\nwant %v\n", status, 200)
	}

	want := `{"message":"Request successful."}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("NotFound return unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}

func TestGetNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/nonexistentroute", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NotFound)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 404 {
		t.Errorf("NotFound returned wrong status code:\ngot %v\nwant %v\n", status, 404)
	}

	want := `{"message":"What do you want?","error":"resource not found."}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("NotFound return unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}

func TestInitGraphqlRequest(t *testing.T) {
	query := []byte("{}")
	body := bytes.NewBuffer(query)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "bearer 0123456789abcdef"
	req, err := initGraphqlRequest(body, headers)

	if err != nil {
		t.Errorf("initGraphqlRequest returned unexpected error: %v\n", err.Error())
	}

	if req.Method != "POST" {
		t.Errorf("initGraphqlRequest returned unexpected method:\ngot %v\nwant %v\n", req.Method, "POST")
	}

	reqBodyReadCloser, _ := req.GetBody()
	reqBody, _ := ioutil.ReadAll(reqBodyReadCloser)
	if string(reqBody) != "{}" {
		t.Errorf("initGraphqlRequest returned unexpected body:\ngot %v\nwant %v\n", string(reqBody), "{}")
	}

	if len(req.Header["Content-Type"]) != 1 {
		t.Errorf("initGraphqlRequest did not return correct number of headers: \"Content-Type\"\nreceived %v Content-Type headers\n", len(req.Header["Content-Type"]))
	} else if req.Header["Content-Type"][0] != "application/json" {
		t.Errorf("initGraphqlRequest did not return header: \"Content-Type: application/json\"\n")
	}

	if len(req.Header["Authorization"]) != 1 {
		t.Errorf("initGraphqlRequest did not return correct number of headers: \"Authorization\"\nreceived %v Authorization headers\n", len(req.Header["Authorization"]))
	} else if req.Header["Authorization"][0] != "bearer 0123456789abcdef" {
		t.Errorf("initGraphqlRequest did not return header: \"Authorization: bearer 0123456789abcdef\"\n")
	}

}

func TestSendJsonInvalidPayload(t *testing.T) {
	unmarshallable := func(invalid int) string {
		return "BAD"
	}

	rr := httptest.NewRecorder()
	sendJson(rr, 200, unmarshallable)

	if status := rr.Code; status != 500 {
		t.Errorf("sendJson returned wrong status code on unmarshallable input:\ngot %v\nwant %v\n", status, 500)
	}

	want := "json: unsupported type: func(int) string"
	got := rr.Body.String()
	if got != want {
		t.Errorf("sendJson return unexpected body on unmarshallable input:\ngot %v\nwant %v\n", got, want)
	}
}

func TestSend400(t *testing.T) {
	err := errors.New("400 error is expected")
	rr := httptest.NewRecorder()
	send400(rr, err)

	if status := rr.Code; status != 400 {
		t.Errorf("send400 returned wrong status code:\ngot %v\nwant %v\n", status, 400)
	}

	want := `{"message":"No cheating! I'm watching you!","error":"400 error is expected"}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("send400 returned unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}

func TestSend404(t *testing.T) {
	rr := httptest.NewRecorder()
	send404(rr)

	if status := rr.Code; status != 404 {
		t.Errorf("send404 returned wrong status code:\ngot %v\nwant %v\n", status, 404)
	}

	want := `{"message":"What do you want?","error":"resource not found."}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("send404 returned unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}

func TestSend500(t *testing.T) {
	err := errors.New("500 error is expected")
	rr := httptest.NewRecorder()
	send500(rr, err)

	if status := rr.Code; status != 500 {
		t.Errorf("send500 returned wrong status code:\ngot %v\nwant %v\n", status, 500)
	}

	want := `{"message":"GRRRR... That was most embarrassing!","error":"500 error is expected"}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("send500 returned unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}
