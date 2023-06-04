package main

import (
	"net/http"
	"strconv"
	"time"
)

func longRequestHandler(w http.ResponseWriter, r *http.Request) {
	timeoutStr := r.FormValue("sleep")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil || timeout < 0 {
		timeout = 0
	}
	if timeout > 10 {
		timeout = 10
	}

	if timeout != 0 {
		time.Sleep(time.Duration(timeout) * time.Second)
	}

	renderTemplate(w, "behaviours/long-request", TemplateInput{Title: "Long Request Handler"})
}
