package main

import (
	"flag"
	"log"
	"os"

	"vladimir.chernenko/snippetbox/pkg/db"
)

type config struct {
	DatabaseDSN string
}

func main() {
	cfg := new(config)
	flag.StringVar(&cfg.DatabaseDSN, "db-dsn", "postgresql://guest:guest@127.0.0.1:6000/snippetbox", "DSN for database")
	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// db init
	dbPool, err := db.OpenDB(cfg.DatabaseDSN)

	if err != nil {
		errorLog.Fatal(err)
	}

	if err = dbPool.AutoMigrate(&db.SnippetModel{}); err != nil {
		errorLog.Fatal(err)
	}
}
