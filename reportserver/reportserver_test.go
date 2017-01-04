package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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

func TestGETWithNoDate(t *testing.T) {
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

func TestGETFavicon(t *testing.T) {
	req, err := http.NewRequest("GET", "/favicon.ico", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerICon)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("GETFavicon FAIL.")
	}
}
