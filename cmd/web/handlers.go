package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	db "vladimir.chernenko/snippetbox/pkg/db"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// templates := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// _, err := template.ParseFiles(templates...)

	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	var snippets []db.SnippetModel

	result := app.dbPool.Where("expires > ?", time.Now().UTC()).Order("created_at").Limit(10).Take(&snippets)
	if result.Error != nil {
		app.serverError(w, result.Error)
		return
	}

	fmt.Fprintf(w, "%v", snippets)

	// err = ts.Execute(w, snippets)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet := &db.SnippetModel{}
	result := app.dbPool.Where(
		"id = ? AND expires > ?", id, time.Now().UTC(),
	).First(snippet)

	if result.Error != nil {
		app.serverError(w, result.Error)
		return
	}

	fmt.Fprintf(w, "%v", snippet)
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
		Expires: time.Now().AddDate(0, 0, 7).UTC(),
	}

	result := app.dbPool.Create(snippet)

	if result.Error != nil {
		app.serverError(w, result.Error)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", snippet.ID), http.StatusSeeOther)
}
