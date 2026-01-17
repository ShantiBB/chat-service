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

func TestCreateMessage(t *testing.T) {
	tests := []struct {
		message   *models.Message
		mockSetup func(repo *mocks.MockRepository)
		name      string
	}{
		{
			name:    "Success",
			message: fixtures.GetTestMessage(uint(1), uint(1), "Test Message"),
			mockSetup: func(repo *mocks.MockRepository) {
				repo.On("InsertMessage", mock.Anything, mock.AnythingOfType("*models.Message")).
					Return(nil)
			},
		},
		{
			name:    "Text trim space",
			message: fixtures.GetTestMessage(uint(1), uint(1), "    Test Message   "),
			mockSetup: func(repo *mocks.MockRepository) {
				repo.On("InsertMessage", mock.Anything, mock.AnythingOfType("*models.Message")).
					Return(nil)
			},
		},
		{
			name:    "Repository error",
			message: fixtures.GetTestMessage(uint(1), uint(1), "Test Message"),
			mockSetup: func(repo *mocks.MockRepository) {
				repo.On("InsertMessage", mock.Anything, mock.AnythingOfType("*models.Message")).
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
				err := service.CreateMessage(context.Background(), tt.message)

				if err != nil {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, "Test Message", tt.message.Text)
				}
			},
		)
	}
}
