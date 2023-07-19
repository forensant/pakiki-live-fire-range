package main

import "net/http"

type IDORUserDetails struct {
	LoggedIn bool
	Username string
}

func idorHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		http.Redirect(w, r, "/auth/idor?id=1", http.StatusFound)
		return
	}

	userDetails := &IDORUserDetails{
		LoggedIn: false,
		Username: "",
	}

	if id == "1" {
		userDetails.LoggedIn = true
		userDetails.Username = "Jill"
	} else if id == "9" {
		userDetails.LoggedIn = true
		userDetails.Username = "Jack"
	} else if id == "21" {
		userDetails.LoggedIn = true
		userDetails.Username = "Joe"
	}

	renderTemplate(w, "auth/idor", TemplateInput{
		Title:      "Insecure Direct Object Reference",
		Additional: userDetails,
	})
}
