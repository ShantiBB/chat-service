package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"chat-service/internal/lib/utils/consts"
	"chat-service/internal/repository/models"
	"chat-service/test/unit/fixtures"
	"chat-service/test/unit/mocks"
)

func TestCreateChat(t *testing.T) {
	tests := []struct {
		chat      *models.Chat
		mockSetup func(repo *mocks.MockRepository)
		name      string
	}{
		{
			name: "Success",
			chat: fixtures.GetTestChat(uint(1), "Test Chat"),
			mockSetup: func(repo *mocks.MockRepository) {
				repo.On("InsertChat", mock.Anything, mock.AnythingOfType("*models.Chat")).
					Return(nil)
			},
		},
		{
			name: "Title trim space",
			chat: fixtures.GetTestChat(uint(1), "    Test Chat   "),
			mockSetup: func(repo *mocks.MockRepository) {
				repo.On("InsertChat", mock.Anything, mock.AnythingOfType("*models.Chat")).
					Return(nil)
			},
		},
		{
			name: "Repository error",
			chat: fixtures.GetTestChat(uint(1), "Test Chat"),
			mockSetup: func(repo *mocks.MockRepository) {
				repo.On("InsertChat", mock.Anything, mock.AnythingOfType("*models.Chat")).
					Return(errors.New(consts.MsgInternal))
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				mockRepo := mocks.NewMockRepository(t)
				tt.mockSetup(mockRepo)

				service := &Service{repo: mockRepo}
				err := service.CreateChat(context.Background(), tt.chat)

				if err != nil {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, "Test Chat", tt.chat.Title)
				}
			},
		)
	}
}

func TestGetChatByID(t *testing.T) {
	tests := []struct {
		mockSetup func(repo *mocks.MockRepository)
		name      string
		chatID    uint
		limit     int
	}{
		{
			name:   "Success",
			chatID: 1,
			limit:  10,
			mockSetup: func(repo *mocks.MockRepository) {
				chat := fixtures.GetTestChat(uint(1), "Test Chat")
				messages := []*models.Message{
					fixtures.GetTestMessage(1, 1, "Message 1"),
					fixtures.GetTestMessage(2, 2, "Message 2"),
				}

				repo.On("SelectChat", mock.Anything, uint(1)).
					Return(chat, nil)
				repo.On("SelectMessages", mock.Anything, uint(1), 10).
					Return(messages, nil)
			},
		},
		{
			name:   "Chat not found",
			chatID: 999,
			limit:  10,
			mockSetup: func(repo *mocks.MockRepository) {
				repo.On("SelectChat", mock.Anything, uint(999)).
					Return((*models.Chat)(nil), consts.ErrChatNotFound)
			},
		},
		{
			name:   "Empty messages",
			chatID: 1,
			limit:  10,
			mockSetup: func(repo *mocks.MockRepository) {
				chat := fixtures.GetTestChat(uint(1), "Empty Chat")

				repo.On("SelectChat", mock.Anything, uint(1)).
					Return(chat, nil)
				repo.On("SelectMessages", mock.Anything, uint(1), 10).
					Return([]*models.Message{}, nil)
			},
		},
		{
			name:   "Repository error",
			chatID: 1,
			limit:  10,
			mockSetup: func(repo *mocks.MockRepository) {
				chat := fixtures.GetTestChat(uint(1), "Test Chat")

				repo.On("SelectChat", mock.Anything, uint(1)).
					Return(chat, nil)
				repo.On("SelectMessages", mock.Anything, uint(1), 10).
					Return(([]*models.Message)(nil), errors.New(consts.MsgInternal))
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				mockRepo := mocks.NewMockRepository(t)
				tt.mockSetup(mockRepo)

				service := &Service{repo: mockRepo}
				chat, err := service.GetChatByID(context.Background(), tt.chatID, tt.limit)

				if err != nil {
					assert.Error(t, err)
					assert.Nil(t, chat)
				} else {
					assert.NoError(t, err)
					assert.NotNil(t, chat)
					assert.Equal(t, tt.chatID, chat.ID)
				}
			},
		)
	}
}

func TestDeleteChatByID(t *testing.T) {
	tests := []struct {
		mockSetup func(repo *mocks.MockRepository)
		name      string
		chatID    uint
	}{
		{
			name:   "Success",
			chatID: 1,
			mockSetup: func(repo *mocks.MockRepository) {
				repo.On("DeleteChat", mock.Anything, uint(1)).
					Return(nil)
			},
		},
		{
			name:   "Chat not found",
			chatID: 999,
			mockSetup: func(repo *mocks.MockRepository) {
				repo.On("DeleteChat", mock.Anything, uint(999)).
					Return(consts.ErrChatNotFound)
			},
		},
		{
			name:   "Repository error",
			chatID: 1,
			mockSetup: func(repo *mocks.MockRepository) {
				repo.On("DeleteChat", mock.Anything, uint(1)).
					Return(errors.New(consts.MsgInternal))
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				mockRepo := mocks.NewMockRepository(t)
				tt.mockSetup(mockRepo)

				service := &Service{repo: mockRepo}
				err := service.DeleteChatByID(context.Background(), tt.chatID)

				if err != nil {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			},
		)
	}
}
