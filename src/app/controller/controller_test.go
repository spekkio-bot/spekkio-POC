package controller

import (
	"errors"
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

	want := `{"message":"No cheating! I'm watching you!","error":"resource not found."}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("NotFound return unexpected body:\ngot %v\nwant %v\n", got, want)
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
	err := errors.New("error is expected")
	rr := httptest.NewRecorder()
	send400(rr, err)

	if status := rr.Code; status != 400 {
		t.Errorf("send400 returned wrong status code:\ngot %v\nwant %v\n", status, 400)
	}

	want := `{"message":"What do you want?","error":"error is expected"}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("send400 returned unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}

func TestSend500(t *testing.T) {
	err := errors.New("error is expected")
	rr := httptest.NewRecorder()
	send500(rr, err)

	if status := rr.Code; status != 500 {
		t.Errorf("send500 returned wrong status code:\ngot %v\nwant %v\n", status, 500)
	}

	want := `{"message":"GRRRR... That was most embarrassing!","error":"error is expected"}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("send500 returned unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}
