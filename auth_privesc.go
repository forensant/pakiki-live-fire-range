package main

import (
	"net/http"
	"strings"
)

const privescAdminCookie string = "GOFrKAb8Rk"
const privescUserCookie string = "3liaT9cdHx"

func privescOverview(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "auth/privesc/overview", TemplateInput{
		Title: "Privilege Escalation",
	})
}

func privescAuthorisationLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("username") == "admin" && r.FormValue("password") == "password" {
			cookie := http.Cookie{
				Name: "privesc_auth",
				Value: privescAdminCookie,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/auth/privesc/menu", http.StatusFound)
		} else if r.FormValue("username") == "reguser" && r.FormValue("password") == "password" {
			cookie := http.Cookie{
				Name: "privesc_auth",
				Value: privescUserCookie,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/auth/privesc/menu", http.StatusFound)
		} else {
			cookie := http.Cookie{
				Name: "privesc_auth",
				Value: "",
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/auth/privesc/overview", http.StatusFound)
		}
	}
}

type PrivescUserDetails struct {
	User  string
	Admin bool
}

func privescMenu(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("privesc_auth")
	if err != nil || (cookie.Value != privescAdminCookie && cookie.Value != privescUserCookie) {
		// redirect to the login page
		http.Redirect(w, r, "/auth/privesc/login", http.StatusFound)
		return
	}

	userDetails := &PrivescUserDetails{
		User: "Administrator",
		Admin: true,
	}

	if cookie.Value != privescAdminCookie {
		userDetails.User = "Regular User"
		userDetails.Admin = false
	}

	renderTemplate(w, "auth/privesc/menu", TemplateInput{
		Title: "Privilege Escalation - Menu",
		Additional: userDetails,
	})
}

func privescPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("privesc_auth")
	granted := true
	if err == nil && cookie.Value == privescAdminCookie {
		// access granted
	} else if err == nil && cookie.Value == privescUserCookie {
		if strings.Contains(r.URL.Path, "/auth/privesc/post") {
			if r.FormValue("id") != "1" {
				granted = false
			}
		}
	} else {
		granted = false
	}

	if !granted {
		w.WriteHeader(http.StatusForbidden)
		renderTemplate(w, "auth/privesc/denied", TemplateInput{
			Title: "Privilege Escalation - Access Denied",
		})
	} else {
		renderTemplate(w, "auth/privesc/granted", TemplateInput{
			Title: "Privilege Escalation - Access Granted",
		})
	}
}
