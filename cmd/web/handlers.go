package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	db "vladimir.chernenko/snippetbox/pkg/db"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	templates := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(templates...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	snippet := &db.SnippetModel{
		Title:   "My First Snippet",
		Content: "My first snippet content",
	}

	result := app.dbPool.Create(snippet)

	if result.Error != nil {
		app.serverError(w, result.Error)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", snippet.ID), http.StatusSeeOther)
}
