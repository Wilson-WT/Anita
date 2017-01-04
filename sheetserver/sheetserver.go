package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	simplejson "github.com/bitly/go-simplejson"

	"github.com/ibmboy19/Anita/sheetserver/config"
	"github.com/ibmboy19/Anita/sheetserver/util"
)

func handler(w http.ResponseWriter, r *http.Request) {
	requestInfo := fmt.Sprintf("%s %s %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto)
	log.Println(requestInfo)
	if r.Method == "GET" {
		handleGet(w, r)
	}
	if r.Method == "POST" {
		handlePost(w, r)
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
		reports, err := util.GetWorkingReports(parameters[1], parameters[2])
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			arrangedReports := util.ArrangeData(reports, parameters[2])
			w.Write([]byte(arrangedReports))
		}
	}
}

// Do Post process
func handlePost(w http.ResponseWriter, r *http.Request) {
	jsonContent, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to decode body to json."))
		return
	}

	dateTime, err := jsonContent.Get("dateTime").String()
	fmt.Println(dateTime)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to get value of dateTime."))
		return
	}

	user, err := jsonContent.Get("user").String()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to get value of user."))
		return
	}

	project, err := jsonContent.Get("project").String()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to get value of project."))
		return
	}

	content, err := jsonContent.Get("content").String()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to get value of content."))
		return
	}

	hours, err := jsonContent.Get("hours").String()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to get value of hours."))
		return
	}

	finishDate, err := jsonContent.Get("finishDate").String()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to get value of finishDate."))
		return
	}

	statusCode := util.Append(dateTime, user, project, content, hours, finishDate)
	if statusCode != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Fail to append record."))
		return
	}
}

func main() {
	// Auth Google sheet
	util.RetrieveToken()

	// Check all config settings are correct
	config.CheckAllConfig()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}
