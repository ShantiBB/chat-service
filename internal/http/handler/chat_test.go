package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"chat-service/internal/config"
	"chat-service/internal/lib/utils/consts"
	"chat-service/test/unit/fixtures"
	"chat-service/test/unit/mocks"
)

func getTestConfig() *config.Config {
	return &config.Config{
		Server: config.ServerConfig{
			Context: config.ContextConfig{
				TimeOut: 5 * time.Second,
			},
		},
	}
}

func setupHandlerTest() (*Handler, *mocks.MockService) {
	mockSvc := new(mocks.MockService)
	cfg := getTestConfig()

	h := &Handler{
		Svc: mockSvc,
		Cfg: cfg,
	}

	return h, mockSvc
}

func TestCreateChat(t *testing.T) {
	tests := []struct {
		requestBody    interface{}
		mockSetup      func(*mocks.MockService)
		name           string
		expectedStatus int
	}{
		{
			name:        "Success",
			requestBody: fixtures.GetCreateChatRequest("Test Chat"),
			mockSetup: func(m *mocks.MockService) {
				m.On(
					"CreateChat", mock.Anything, mock.AnythingOfType("*models.Chat"),
				).Return(nil)
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "Invalid JSON",
			requestBody:    "invalid json",
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Validation Error - Empty Name",
			requestBody:    fixtures.GetCreateChatRequest(""),
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "Service Error",
			requestBody: fixtures.GetCreateChatRequest("Test Chat"),
			mockSetup: func(m *mocks.MockService) {
				m.On("CreateChat", mock.Anything, mock.AnythingOfType("*models.Chat")).
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

				var body []byte
				if str, ok := test.requestBody.(string); ok {
					body = []byte(str)
				} else {
					body, _ = json.Marshal(test.requestBody)
				}

				req := httptest.NewRequest(http.MethodPost, "/chats", bytes.NewBuffer(body))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()

				h.CreateChat(w, req)

				assert.Equal(t, test.expectedStatus, w.Code)
				mockSvc.AssertExpectations(t)

				if test.expectedStatus == http.StatusBadRequest {
					mockSvc.AssertNotCalled(
						t, "CreateChat", mock.Anything, mock.AnythingOfType("*models.Chat"),
					)
				}
			},
		)
	}
}

func TestGetChatByID(t *testing.T) {
	tests := []struct {
		urlParams      map[string]string
		mockSetup      func(*mocks.MockService)
		name           string
		queryParams    string
		expectedStatus int
	}{
		{
			name:        "Success",
			urlParams:   map[string]string{"id": "1"},
			queryParams: "?limit=10",
			mockSetup: func(m *mocks.MockService) {
				expectedChat := fixtures.GetTestChat(1, "Test Chat")
				m.On("GetChatByID", mock.Anything, uint(1), 10).Return(expectedChat, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid ID",
			urlParams:      map[string]string{"id": "invalid"},
			queryParams:    "",
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "Not Found",
			urlParams:   map[string]string{"id": "999"},
			queryParams: "?limit=10",
			mockSetup: func(m *mocks.MockService) {
				m.On("GetChatByID", mock.Anything, uint(999), 10).
					Return(nil, consts.ErrChatNotFound)
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Invalid Limit",
			urlParams:      map[string]string{"id": "1"},
			queryParams:    "?limit=invalid",
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "Service Error",
			urlParams:   map[string]string{"id": "1"},
			queryParams: "?limit=10",
			mockSetup: func(m *mocks.MockService) {
				m.On("GetChatByID", mock.Anything, uint(1), 10).
					Return(nil, errors.New(consts.MsgInternal))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				h, mockSvc := setupHandlerTest()
				test.mockSetup(mockSvc)

				req := httptest.NewRequest(http.MethodGet, "/chats/"+test.queryParams, nil)
				req.SetPathValue("id", test.urlParams["id"])
				w := httptest.NewRecorder()

				h.GetChatByID(w, req)

				assert.Equal(t, test.expectedStatus, w.Code)
				mockSvc.AssertExpectations(t)

				if test.expectedStatus == http.StatusBadRequest {
					mockSvc.AssertNotCalled(
						t, "CreateChat", mock.Anything, mock.AnythingOfType("*models.Chat"),
					)
				}
			},
		)
	}
}

func TestDeleteChatByID(t *testing.T) {
	tests := []struct {
		urlParams      map[string]string
		mockSetup      func(*mocks.MockService)
		name           string
		expectedStatus int
	}{
		{
			name:      "Success",
			urlParams: map[string]string{"id": "1"},
			mockSetup: func(m *mocks.MockService) {
				m.On("DeleteChatByID", mock.Anything, uint(1)).Return(nil)
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "Invalid ID",
			urlParams:      map[string]string{"id": "invalid"},
			mockSetup:      func(m *mocks.MockService) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:      "Not Found",
			urlParams: map[string]string{"id": "999"},
			mockSetup: func(m *mocks.MockService) {
				m.On("DeleteChatByID", mock.Anything, uint(999)).Return(consts.ErrChatNotFound)
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name:      "Service Error",
			urlParams: map[string]string{"id": "1"},
			mockSetup: func(m *mocks.MockService) {
				m.On("DeleteChatByID", mock.Anything, uint(1)).Return(errors.New(consts.MsgInternal))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				h, mockSvc := setupHandlerTest()
				test.mockSetup(mockSvc)

				req := httptest.NewRequest(http.MethodGet, "/chats/", nil)
				req.SetPathValue("id", test.urlParams["id"])
				w := httptest.NewRecorder()

				h.DeleteChatByID(w, req)

				assert.Equal(t, test.expectedStatus, w.Code)
				mockSvc.AssertExpectations(t)

				if test.expectedStatus == http.StatusBadRequest {
					mockSvc.AssertNotCalled(
						t, "CreateChat", mock.Anything, mock.AnythingOfType("*models.Chat"),
					)
				}
			},
		)
	}
}
