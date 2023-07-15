package main

import (
	"net/http"
	"fmt")

func reconHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "recon/overview", TemplateInput{
		Title:          "Reconnaissance",
	})
}

func reconSecretPage1Handler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "recon/secret-page1", TemplateInput{
		Title: "Secret Page 1",
	})
}

func reconSecretPage2Handler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "recon/secret-page2", TemplateInput{
		Title: "Secret Page 2",
	})
	fmt.Fprintf(w, "<!-- Note to self, the password is: 2ns9mNKe2N -->")
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "User-agent: *\nDisallow: /recon/Ysiewc58rC\n\nSitemap: https://livefirerange.pakikiproxy.com/sitemap.xml")
}

func sitemapHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprintf(w, `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	<url>
		<loc>https://livefirerange.pakikiproxy.com/recon/NVCwG68UI4</loc>
		<lastmod>2020-08-01T00:00:00+00:00</lastmod>
	</url>
</urlset>`)
}

