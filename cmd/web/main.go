package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"
	db "vladimir.chernenko/snippetbox/pkg/db"
)

type Config struct {
	Addr        string
	StaticDir   string
	DatabaseDSN string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	dbPool   *gorm.DB
}

func main() {
	// Parsing the runtime configuration settings for the application;
	// Establishing the dependencies for the handlers; and
	// Running the HTTP server.

	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static folder")
	flag.StringVar(&cfg.DatabaseDSN, "db-dsn", "postgresql://guest:guest@127.0.0.1:6000/snippetbox", "DSN for database")
	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// db init
	dbPool, err := db.OpenDB(cfg.DatabaseDSN)
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		dbPool:   dbPool,
	}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(cfg.StaticDir),
	}

	// As a rule of thumb, you should avoid using the Panic() and Fatal()
	// variations outside of your main() function —
	// it’s good practice to return errors instead, and only panic or exit directly from main().
	infoLog.Printf("Starting server on %s", cfg.Addr)

	err = srv.ListenAndServe()

	errorLog.Fatal(err)
}
