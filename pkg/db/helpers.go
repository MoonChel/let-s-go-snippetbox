package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateDB(dbPool *gorm.DB) error {
	err := dbPool.AutoMigrate(&SnippetModel{})

	return err
}
