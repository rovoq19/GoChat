package repository

import (
	"github.com/rovoq19/GoChat/pkg/models"
	"gorm.io/gorm"
)

type Chat interface {
	Create(list models.Chat) (int, error)
	GetAll() ([]models.Chat, error)
	GetById(chatId int) (models.Chat, error)
	Delete(chatId int) error
	Update(chatId int, input models.UpdateChatInput) error
}

type Repository struct {
	Chat
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Chat: NewChatPostgres(db),
	}
}
