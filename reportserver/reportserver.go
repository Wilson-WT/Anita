package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/ibmboy19/Anita/reportserver/config"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handleGet(w, r)
	}
}

// Do Get process
func handleGet(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	parameters := strings.Split(url, "/")

	if len(parameters) != 3 || parameters[2] == "" {
		w.Write([]byte("Invalid url format! Please check your url.\n"))
		w.Write([]byte("For an example:\n"))
		w.Write([]byte("http://10.20.108.20:2001/jacky.wu/201603-201609"))
	} else {
		// Call sheet server's API
		sheetServerURL := config.ImportConfig().GetSheetServerURL()
		actionPath := fmt.Sprintf(sheetServerURL+"/%s/%s", parameters[1], parameters[2])
		response, err := http.Get(actionPath)
		defer response.Body.Close()
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			responseContent, err := ioutil.ReadAll(response.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.Write([]byte(responseContent))
			}
		}
	}
}

func handlerICon(w http.ResponseWriter, r *http.Request) {
	fname := path.Base(r.URL.Path)
	http.ServeFile(w, r, "./"+fname)
}

func main() {
	// Check all config settings are correct
	config.CheckAllConfig()

	http.HandleFunc("/favicon.ico", handlerICon)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":2001", nil)
}
