package controller

import (
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
