// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type MenuCategory struct {
	Name  string
	Items []MenuItem
}

type MenuItem struct {
	Name string
	Path string
}

type TemplateInput struct {
	Title          string
	NavbarTitle    string
	MenuCategories []MenuCategory
	Additional     interface{}
}

var menu = []MenuCategory{
	{
		Name: "Behaviours",
		Items: []MenuItem{
			{
				Name: "Long Request",
				Path: "/behaviours/long-request",
			},
		},
	},
	{
		Name: "Reconissance",
		Items: []MenuItem{
			{
				Name: "Overview",
				Path: "/recon",
			},
		},
	},
	{
		Name: "Authorisation",
		Items: []MenuItem{
			{
				Name: "Privilege Escalation",
				Path: "/auth/privesc/overview",
			},
			{
				Name: "Insecure Direct Object Reference",
				Path: "/auth/idor",
			},
		},
	},
	{
		Name: "Injection",
		Items: []MenuItem{
			{
				Name: "XSS",
				Path: "/injection/xss",
			},
			{
				Name: "SQL Injection",
				Path: "/injection/sqli",
			},
		},
	},
	{
		Name: "Cryptography",
		Items: []MenuItem{
			{
				Name: "ECB Padding Oracle",
				Path: "/not-implemented",
			},
			{
				Name: "CBC Padding Oracle",
				Path: "/not-implemented",
			},
		},
	},
}

func main() {
	log.Print("starting server...")
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/assets/", fileServer)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/robots.txt", robotsHandler)
	http.HandleFunc("/sitemap.xml", sitemapHandler)
	http.HandleFunc("/recon", reconHandler)
	http.HandleFunc("/recon/Ysiewc58rC", reconSecretPage1Handler)
	http.HandleFunc("/recon/NVCwG68UI4", reconSecretPage2Handler)
	http.HandleFunc("/auth/idor", idorHandler)
	http.HandleFunc("/auth/privesc/overview", privescOverview)
	http.HandleFunc("/auth/privesc/login", privescAuthorisationLogin)
	http.HandleFunc("/auth/privesc/menu", privescMenu)
	http.HandleFunc("/auth/privesc/users", privescPage)
	http.HandleFunc("/auth/privesc/post", privescPage)
	http.HandleFunc("/injection/xss", xssHandler)
	http.HandleFunc("/injection/sqli", sqliHandler)
	http.HandleFunc("/not-implemented", notImplementedHandler)
	http.HandleFunc("/behaviours/long-request", longRequestHandler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, templateName string, input TemplateInput) {
	input.MenuCategories = menu
	if input.Title == "" {
		input.Title = "Pākiki Proxy — Live Fire Range"
	} else {
		input.Title = input.Title + " — Pākiki Proxy Live Fire Range"
	}

	paths := []string{
		"templates/" + templateName + ".tmpl",
		"templates/header.tmpl",
		"templates/footer.tmpl",
	}

	tmpl, err := template.ParseFiles(paths...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", TemplateInput{})
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "not-implemented", TemplateInput{Title: "Not Implemented"})
}
