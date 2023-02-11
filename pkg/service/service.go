package service

import (
	"github.com/rovoq19/GoChat/cmd/ws"
	"github.com/rovoq19/GoChat/pkg/models"
	"github.com/rovoq19/GoChat/pkg/repository"
)

type Chat interface {
	Create(chat models.Chat) (int, error)
	GetAll() ([]models.Chat, error)
	GetById(chatId int) (models.Chat, error)
	Delete(chatId int) error
	Update(chatId int, input models.UpdateChatInput) error
	SendMessage(m models.Message) error
}

type Service struct {
	Chat
}

func NewService(repos *repository.Repository, hub *ws.Hub) *Service {
	return &Service{
		Chat: NewChatService(repos.Chat, hub),
	}
}
