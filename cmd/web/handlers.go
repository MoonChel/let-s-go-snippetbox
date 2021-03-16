package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	db "vladimir.chernenko/snippetbox/pkg/db"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	var snippets []db.SnippetModel

	result := app.dbPool.Where("expires > ?", time.Now().UTC()).Order("created_at").Limit(10).Find(&snippets)
	if result.Error != nil {
		app.serverError(w, result.Error)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Snippets: &snippets})
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.URL.Query().Get(":id"), 10, 64)

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
	} else if snippet.ID == 0 {
		app.notFound(w)
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{Snippet: snippet})
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	snippet := &db.SnippetModel{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
		Expires: time.Now().AddDate(0, 0, expires).UTC(),
	}

	result := app.dbPool.Create(snippet)

	if result.Error != nil {
		app.serverError(w, result.Error)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", snippet.ID), http.StatusSeeOther)
}
