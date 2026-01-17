package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"chat-service/test/unit/fixtures"
	"chat-service/test/unit/mocks"
)

func TestCreateMessage(t *testing.T) {
	tests := []struct {
		requestBody    interface{}
		mockSetup      func(*mocks.MockService)
		name           string
		chatID         string
		expectedStatus int
	}{
		{
			name:        "Success",
			chatID:      "1",
			requestBody: fixtures.GetCreateMessageRequest("success"),
			mockSetup: func(m *mocks.MockService) {
				m.On("CreateMessage", mock.Anything, mock.AnythingOfType("*models.Message")).
					Return(nil)
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "Invalid chatID",
			chatID:         "invalid",
			requestBody:    fixtures.GetCreateMessageRequest("hello"),
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid JSON",
			chatID:         "1",
			requestBody:    "invalid json",
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Validation Error",
			chatID:         "1",
			requestBody:    fixtures.GetCreateMessageRequest(""),
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "Service Error",
			chatID:      "1",
			requestBody: fixtures.GetCreateMessageRequest("hello"),
			mockSetup: func(m *mocks.MockService) {
				m.On("CreateMessage", mock.Anything, mock.AnythingOfType("*models.Message")).
					Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				h, mockSvc := setupHandlerTest()
				test.mockSetup(mockSvc)

				var body io.Reader
				switch v := test.requestBody.(type) {
				case string:
					body = strings.NewReader(v)
				default:
					b, err := json.Marshal(v)
					require.NoError(t, err)
					body = bytes.NewReader(b)
				}

				req := httptest.NewRequest(
					http.MethodPost,
					"/chats/"+test.chatID+"/messages",
					body,
				)
				req.Header.Set("Content-Type", "application/json")
				req.SetPathValue("chatID", test.chatID)

				w := httptest.NewRecorder()
				h.CreateMessage(w, req)

				assert.Equal(t, test.expectedStatus, w.Code)
				mockSvc.AssertExpectations(t)

				if test.expectedStatus == http.StatusBadRequest {
					mockSvc.AssertNotCalled(t, "CreateMessage", mock.Anything, mock.Anything)
				}
			},
		)
	}
}
