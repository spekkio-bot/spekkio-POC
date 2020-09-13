package controller

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
		t.Errorf("Ping returned wrong status code:\ngot %v\nwant %v\n", status, 200)
	}

	want := `{"message":"Request successful."}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("Ping returned unexpected body:\ngot %v\nwant %v\n", got, want)
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
		t.Errorf("NotFound returned unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MethodNotAllowed)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 405 {
		t.Errorf("MethodNotAllowed returned wrong status code:\ngot %v\nwant %v\n", status, 405)
	}

	want := `{"message":"No cheating!","error":"method not allowed."}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("MethodNotAllowed returned unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}

func scrumifyTestWrapper(w http.ResponseWriter, r *http.Request) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "color", "description"}).
		AddRow(1, "Defect", "ff0000", "Something is not working").
		AddRow(2, "Story", "00ffff", "New feature")

	mock.ExpectQuery("^SELECT (.+) FROM ScrumifyLabels").
		WillReturnRows(rows)

	Scrumify(db, w, r)
}

func TestPostScrumify(t *testing.T) {
	reqBodyBytes := []byte("{\"repo_id\":\"123\",\"token\":\"321\"}")
	reqBody := bytes.NewBuffer(reqBodyBytes)

	req, err := http.NewRequest("POST", "/scrumify", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(scrumifyTestWrapper)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 401 {
		t.Errorf("Scrumify returned wrong status code:\ngot %v\nwant %v\n", status, 401)
	}

	//want := `{"message":"Ipso facto, meeny moe... MAGICO! Your repository was successfully scrumified!"}`
	want := `{"message":"No cheating!","error":"unauthorized."}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("Scrumify returned unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}

func TestPostScrumifyNoToken(t *testing.T) {
	reqBodyBytes := []byte("{\"repo_id\":\"notoken\"}")
	reqBody := bytes.NewBuffer(reqBodyBytes)

	req, err := http.NewRequest("POST", "/scrumify", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(scrumifyTestWrapper)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 400 {
		t.Errorf("Scrumify (no token) returned wrong status code:\ngot %v\nwant %v\n", status, 400)
	}

	want := `{"message":"No cheating! I'm watching you!","error":"request body is missing token property"}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("Scrumify (no token) returned unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}

func TestPostScrumifyNoRepoId(t *testing.T) {
	reqBodyBytes := []byte("{\"token\":\"norepoid\"}")
	reqBody := bytes.NewBuffer(reqBodyBytes)

	req, err := http.NewRequest("POST", "/scrumify", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(scrumifyTestWrapper)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 400 {
		t.Errorf("Scrumify (no repo id) returned wrong status code:\ngot %v\nwant %v\n", status, 400)
	}

	want := `{"message":"No cheating! I'm watching you!","error":"request body is missing repo_id property"}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("Scrumify (no repo id) returned unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}

func TestPostScrumifyEmpty(t *testing.T) {
	req, err := http.NewRequest("POST", "/scrumify", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(scrumifyTestWrapper)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 400 {
		t.Errorf("Scrumify (empty body) returned wrong status code:\ngot %v\nwant %v\n", status, 400)
	}

	want := `{"message":"No cheating! I'm watching you!","error":"request body is empty"}`
	got := rr.Body.String()
	if got != want {
		t.Errorf("Scrumify (empty body) returned unexpected body:\ngot %v\nwant %v\n", got, want)
	}
}

func TestInitGraphqlRequest(t *testing.T) {
	query := []byte("{}")
	body := bytes.NewBuffer(query)
	headers := make(map[string][]string)
	headers["Content-Type"] = []string{
		"application/json",
	}
	headers["Cache-Control"] = []string{
		"no-cache",
		"no-store",
		"must-revalidate",
	}
	headers["Authorization"] = []string{
		"bearer 0123456789abcdef",
	}
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

	if len(req.Header["Cache-Control"]) != 3 {
		t.Errorf("initGraphqlRequest did not return correct number of headers: \"Cache-Control\"\nreceived %v Cache-Control headers\n", len(req.Header["Cache-Control"]))
	} else {
		for i, expectedHeader := range headers["Cache-Control"] {
			if req.Header["Cache-Control"][i] != expectedHeader {
				t.Errorf("initGraphqlRequest did not return header: \"Cache-Control: %v\"\ngot \"Cache-Control: %v\"\n", expectedHeader, req.Header["Cache-Control"][i])
			}
		}
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
