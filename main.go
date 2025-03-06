package main

import (
	"gopkg.in/yaml.v2"
	"html/template"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type SlidePage struct {
	Pages []SlideData `yaml:"pages"`
}
type SlideData struct {
	Content string `yaml:"content"`
	Next    int
	HasNext bool
	Prev    int
	HasPrev bool
}

type Color struct {
	Name float64
}

type FormContent struct {
	Voornaam   string
	Achternaam string
	Eigenschap string
}

func handleSlide(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	if len(params) != 3 {
		return
	}

	page, err := strconv.Atoi(params[2])
	if err != nil || page < 1 {
		return
	}

	data, err := os.ReadFile("data/slides.yaml")
	if err != nil {
		log.Fatal("Error reading file")
	}

	var pages SlidePage
	err = yaml.Unmarshal(data, &pages)
	if err != nil {
		log.Fatal(err)
	}

	if len(pages.Pages) >= page {
		pageData := pages.Pages[page-1]
		// show previous page if current page isn't the first in the list
		pageData.Prev = page - 1
		pageData.HasPrev = page-1 > 0
		// show next page if current page isn't the last in the list
		pageData.Next = page + 1
		pageData.HasNext = len(pages.Pages) > page

		tmpl, err := template.ParseFiles("./templates/slide.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, pageData)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func handlesColor(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/color.html")
	if err != nil {
		log.Fatal(err)
	}
	randCol := math.Ceil(rand.Float64() * 7)

	err = tmpl.Execute(w, Color{Name: randCol})

}

func handleFormAnswer(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	content := FormContent{
		Voornaam:   query.Get("voornaam"),
		Achternaam: query.Get("achternaam"),
		Eigenschap: query.Get("eigenschap"),
	}
	tmpl, err := template.ParseFiles("./templates/formanswer.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, content)
}

func main() {
	// Serve files from the ./public directory
	fs := http.FileServer(http.Dir("./public"))

	// Handle all requests by serving the file server
	http.Handle("/", fs)
	http.HandleFunc("/slide/", handleSlide)
	http.HandleFunc("/color", handlesColor)
	http.HandleFunc("/form-answer", handleFormAnswer)

	// Start the server on port 8080
	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
