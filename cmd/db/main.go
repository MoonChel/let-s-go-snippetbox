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

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// db init
	dbPool, err := db.OpenDB(cfg.DatabaseDSN)

	if err != nil {
		errorLog.Fatal(err)
	}

	if err = db.MigrateDB(dbPool); err != nil {
		errorLog.Fatal(err)
	}

	infoLog.Println("Migrated successfully")
}
