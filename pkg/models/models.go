package models

import (
	"errors"
	"time"
)

type Chat struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

type Message struct {
	Username string `json:"username"`
	ChatId   int    `json:"chat_id"`
	Message  string `json:"message"`
}

type UpdateChatInput struct {
	Title *string `json:"title"`
}

func (i UpdateChatInput) Validate() error {
	if i.Title == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
