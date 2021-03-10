package db

import "time"

type SnippetModel struct {
	ID      uint
	Title   string    `gorm:"not null;size:100"`
	Content string    `gorm:"not null"`
	Created time.Time `gorm:"not null"`
	Updated time.Time `gorm:"not null"`
}

func (SnippetModel) TableName() string {
	return "snippet"
}
