package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

// Template data we will pass into the templates
type TemplateData struct {
	Days   string
	Footer string
}

// The cached templates
var templates *template.Template

func getTimeSinceWedding() string {
	location, err := time.LoadLocation("America/Detroit")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load location properly, returning nil.  Error: %s", err)
	}
	t := time.Date(2023, time.January, 21, 12, 0, 0, 0, location)
	elapsedTime := time.Since(t)
	return fmt.Sprintf("It's been %d days since we got married!",int(elapsedTime.Hours() / 24))
}

func getFooter() string {
	n := time.Now()
	return fmt.Sprintf("Blanchard %d",n.Year())
}

// Main page handler.
func handler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Days:   getTimeSinceWedding(),
		Footer: getFooter(),
	}
	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Printf("Something happened:\n %s", err.Error())
		fmt.Fprintf(w, "Failed to load template content properly.")
	}
}

func loadTemplates() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	loadTemplates()
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
