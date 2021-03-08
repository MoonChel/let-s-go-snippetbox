package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr      string
	StaticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Parsing the runtime configuration settings for the application;
	// Establishing the dependencies for the handlers; and
	// Running the HTTP server.

	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static folder")
	flag.Parse()

	app := &application{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	// As a rule of thumb, you should avoid using the Panic() and Fatal()
	// variations outside of your main() function —
	// it’s good practice to return errors instead, and only panic or exit directly from main().
	app.infoLog.Printf("Starting server on %s", cfg.Addr)
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)
}
