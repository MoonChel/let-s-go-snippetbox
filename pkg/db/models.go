package db

import "time"

type SnippetModel struct {
	ID        uint
	Title     string `gorm:"not null;size:100"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SnippetModel) TableName() string {
	return "snippet"
}
