package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

func sqliHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	additional := struct {
		AuthSuccessful bool
		AttemptMade    bool
	}{
		AuthSuccessful: false,
		AttemptMade:    false,
	}

	if username == "" && password == "" {
		renderTemplate(w, "injection/sqli", TemplateInput{
			Title:      "SQL Injection",
			Additional: additional,
		})
		return
	}

	additional.AttemptMade = true

	// otherwise create the database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table users (id integer not null primary key, username text, password text);
	insert into users values (1, 'admin', 'W9hIa54BEJWI');
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	// now query for the username/password
	sqlStmt = "select username, password from users where username='" + username + "' and password='" + password + "'"

	rows, err := db.Query(sqlStmt)
	if err == nil && rows.Next() {
		additional.AuthSuccessful = true
	}
	defer rows.Close()

	renderTemplate(w, "injection/sqli", TemplateInput{
		Title:      "SQL Injection",
		Additional: additional,
	})
}

func xssHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	templateInput := TemplateInput{
		Title:          "XSS - Pākiki Proxy Live Fire Range",
		MenuCategories: menu,
		Additional: struct {
			Name string
		}{
			Name: name,
		},
	}

	paths := []string{
		"templates/injection/xss.tmpl",
		"templates/header.tmpl",
		"templates/footer.tmpl",
	}

	tmpl, err := template.ParseFiles(paths...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, templateInput); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
