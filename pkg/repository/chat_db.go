package repository

import (
	"errors"
	"github.com/rovoq19/GoChat/pkg/models"
	"gorm.io/gorm"
)

type ChatPostgres struct {
	db *gorm.DB
}

func NewChatPostgres(db *gorm.DB) *ChatPostgres {
	return &ChatPostgres{db: db}
}

func (r *ChatPostgres) Create(chat models.Chat) (int, error) {
	result := r.db.Create(&chat)
	if result.Error != nil {
		return 0, result.Error
	}
	return chat.Id, nil
}

func (r *ChatPostgres) GetAll() ([]models.Chat, error) {
	var chats []models.Chat
	if result := r.db.Find(&chats); result.Error != nil {
		return []models.Chat{}, result.Error
	}
	return chats, nil
}

func (r *ChatPostgres) GetById(chatId int) (models.Chat, error) {
	var chat models.Chat
	if result := r.db.Where("id = ?", chatId).Find(&chat); result.Error != nil {
		return models.Chat{}, result.Error
	}
	return chat, nil
}

func (r *ChatPostgres) Delete(chatId int) error {
	result := r.db.Where("id = ?", chatId).Delete(&models.Chat{})
	return result.Error
}

func (r *ChatPostgres) Update(chatId int, input models.UpdateChatInput) error {
	var chat models.Chat
	if input.Title != nil {
		if result := r.db.Where("id = ?", chatId).Find(&chat); result.Error != nil {
			return result.Error
		} else {
			chat.Title = *input.Title
			r.db.Save(&chat)
		}
	} else {
		return errors.New("empty title param")
	}
	return nil
}
