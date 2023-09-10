package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type task_details struct {
	Slack_name      string `json:"slack_name"`
	Current_day     string `json:"current_day"`
	Utc_time        string `json:"utc_time"`
	Track           string `json:"track"`
	Github_file_url string `json:"github_file_url"`
	Github_repo_url string `json:"github_repo_url"`
	Status_code     int    `json:"status_code "`
}

var logger log.Logger = *log.New(os.Stdout, "task-on", log.LstdFlags)

func (t *task_details) ToJson(w io.Writer) error { // Convert task_details to json

	logger.Println("ToJson function called")
	e := json.NewEncoder(w)
	return e.Encode(t)
}

func New_task_details() *task_details {
	logger.Println("New instance of task details")

	return &task_details{
		Slack_name:      "",
		Current_day:     "",
		Utc_time:        "",
		Track:           "",
		Github_file_url: "https://github.com/MustaphaAwwal/JSONEndpointWithGo/blob/master/main.go",
		Github_repo_url: "https://github.com/MustaphaAwwal/JSONEndpointWithGo",
		Status_code:     http.StatusOK,
	}

}

func (t *task_details) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//serveHTTP task details API

	rw.Header().Set("Content-Type", "application/json")

	currentTime := time.Now().UTC()
	currentDate := time.Now().Weekday().String()

	t.Utc_time = currentTime.Format("2006-01-02T15:04:05Z")
	t.Current_day = currentDate
	t.Slack_name = url.QueryEscape(r.URL.Query().Get("slack_name"))
	t.Track = url.QueryEscape(r.URL.Query().Get("track"))
	err := t.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to Marshal", http.StatusInternalServerError)
	}
	return
}

func main() {
	logger.Println("Starting server on port 80")
	task_one := New_task_details()
	fmt.Print(*task_one)

	sm := http.NewServeMux()

	sm.Handle("/api", task_one)

	http.ListenAndServe(":80", sm)
}
