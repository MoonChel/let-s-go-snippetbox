package main

import (
	db "vladimir.chernenko/snippetbox/pkg/db"
)

type templateData struct {
	Snippet  *db.SnippetModel
	Snippets *[]db.SnippetModel
}
