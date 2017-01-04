package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGETWithCorrectUrlFormat(t *testing.T) {
	req, err := http.NewRequest("GET", "/timmy.tsai/201609-201609", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("GETWithCorrectUrlFormat FAIL.")
	}

	isContain := strings.Contains(rr.Body.String(), "Invalid url format! Please check your url.")
	if isContain == false {
		t.Log("GETWithCorrectUrlFormat PASS.")
	} else {
		t.Error("GETWithCorrectUrlFormat FAIL.")
	}
}

func TestGETWithNoAccount(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("GETWithNoAccount FAIL.")
	}

	isContain := strings.Contains(rr.Body.String(), "Invalid url format! Please check your url.")
	if isContain == true {
		t.Log("GETWithNoAccount PASS.")
	} else {
		t.Error("GETWithNoAccount FAIL.")
	}
}

func TestGETWithNoPeriod(t *testing.T) {
	req, err := http.NewRequest("GET", "/timmy.tsai/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("GETWithNoDate FAIL.")
	}

	isContain := strings.Contains(rr.Body.String(), "Invalid url format! Please check your url.")
	if isContain == true {
		t.Log("GETWithNoDate PASS.")
	} else {
		t.Error("GETWithNoDate FAIL.")
	}
}

func TestGETWithInvalidPeriod(t *testing.T) {
	req, err := http.NewRequest("GET", "/timmy.tsai/201607-201603", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("GETWithNoDate FAIL.")
	}

	isContain := strings.Contains(rr.Body.String(), "Period is invalid")
	if isContain == true {
		t.Log("GETWithNoDate PASS.")
	} else {
		t.Error("GETWithNoDate FAIL.")
	}
}

func TestPOSTWithInvalidJsonFormat(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader("test"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("POSTWithInvalidJsonFormat FAIL.")
	}

	isContain := strings.Contains(rr.Body.String(), "Unable to decode body to json.")
	if isContain == true {
		t.Log("POSTWithInvalidJsonFormat PASS.")
	} else {
		t.Error("POSTWithInvalidJsonFormat FAIL.")
	}
}

func TestPOSTWithNoDateTime(t *testing.T) {
	var jsonStr = []byte(`{"user": "timmy","project": "AFU","content": "hello world","hours": "1","finishDate": "20161031"}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("POSTWithNoDateTime FAIL.")
	}

	isContain := strings.Contains(rr.Body.String(), "Unable to get value of dateTime.")
	if isContain == true {
		t.Log("POSTWithNoDateTime PASS.")
	} else {
		t.Error("POSTWithNoDateTime FAIL.")
	}
}

func TestPOSTWithNoUser(t *testing.T) {
	var jsonStr = []byte(`{"dateTime": "20161031","project": "AFU","content": "hello world","hours": "1","finishDate": "20161031"}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("POSTWithNoUser FAIL.")
	}

	isContain := strings.Contains(rr.Body.String(), "Unable to get value of user.")
	if isContain == true {
		t.Log("POSTWithNoUser PASS.")
	} else {
		t.Error("POSTWithNoUser FAIL.")
	}
}

func TestPOSTWithNoProject(t *testing.T) {
	var jsonStr = []byte(`{"dateTime": "20161031","user": "timmy","content": "hello world","hours": "1","finishDate": "20161031"}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("POSTWithNoProject FAIL.")
	}

	isContain := strings.Contains(rr.Body.String(), "Unable to get value of project.")
	if isContain == true {
		t.Log("POSTWithNoProject PASS.")
	} else {
		t.Error("POSTWithNoProject FAIL.")
	}
}

func TestPOSTWithNoContent(t *testing.T) {
	var jsonStr = []byte(`{"dateTime": "20161031","user": "timmy","project": "timmy","hours": "1","finishDate": "20161031"}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("POSTWithNoContent FAIL.")
	}

	isContain := strings.Contains(rr.Body.String(), "Unable to get value of content.")
	if isContain == true {
		t.Log("POSTWithNoContent PASS.")
	} else {
		t.Error("POSTWithNoContent FAIL.")
	}
}

func TestPOSTWithNoHours(t *testing.T) {
	var jsonStr = []byte(`{"dateTime": "20161031","user": "timmy","project": "AFU","content": "hello world","finishDate": "20161031"}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("POSTWithNoHours FAIL.")
	}

	isContain := strings.Contains(rr.Body.String(), "Unable to get value of hours.")
	if isContain == true {
		t.Log("POSTWithNoHours PASS.")
	} else {
		t.Error("POSTWithNoHours FAIL.")
	}
}

func TestPOSTWithNoFinishDate(t *testing.T) {
	var jsonStr = []byte(`{"dateTime": "20161031","user": "timmy","project": "AFU","content": "hello world","hours": "1"}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("POSTWithNoFinishDate FAIL.")
	}

	isContain := strings.Contains(rr.Body.String(), "Unable to get value of finishDate.")
	if isContain == true {
		t.Log("POSTWithNoFinishDate PASS.")
	} else {
		t.Error("POSTWithNoFinishDate FAIL.")
	}
}
