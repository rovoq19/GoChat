package service

import (
	"encoding/json"
	"github.com/rovoq19/GoChat/cmd/ws"
	"github.com/rovoq19/GoChat/pkg/models"
	"github.com/rovoq19/GoChat/pkg/repository"
	"time"
)

type ChatService struct {
	repo repository.Chat
	hub  *ws.Hub
}

func NewChatService(repo repository.Chat, hub *ws.Hub) *ChatService {
	return &ChatService{repo: repo, hub: hub}
}

func (s *ChatService) Create(chat models.Chat) (int, error) {
	chat.CreatedAt = time.Now().Local()
	return s.repo.Create(chat)
}

func (s *ChatService) GetAll() ([]models.Chat, error) {
	return s.repo.GetAll()
}

func (s *ChatService) GetById(chatId int) (models.Chat, error) {
	return s.repo.GetById(chatId)
}

func (s *ChatService) Delete(chatId int) error {
	return s.repo.Delete(chatId)
}

func (s *ChatService) Update(chatId int, input models.UpdateChatInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(chatId, input)
}

func (s *ChatService) SendMessage(m models.Message) error {
	message, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return s.hub.Publish(message)
}
