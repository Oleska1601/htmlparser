package controller_test

import (
	"errors"
	"htmlparser/internal/controller"
	controllerMocks "htmlparser/internal/controller/mocks"
	loggingMocks "htmlparser/internal/logging/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServer_GetDataHander(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		queryParams map[string]string
		setupMocks  func(usecase *controllerMocks.UsecaseInterface, l *loggingMocks.LoggerInterface)
		// Named input parameters for target function.
		expectedStatus int
	}{
		{
			name:        "test1",
			queryParams: map[string]string{},
			setupMocks: func(usecase *controllerMocks.UsecaseInterface, l *loggingMocks.LoggerInterface) {
				l.On("Error", "GetDataHander c.Param", mock.Anything, mock.Anything).Return()
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "test2",
			queryParams: map[string]string{
				"url": "https://google.com",
			},
			setupMocks: func(usecase *controllerMocks.UsecaseInterface, l *loggingMocks.LoggerInterface) {
				usecase.On("GetParsingDataV1", mock.Anything, mock.Anything).Return("", errors.New(""))
				l.On("Error", "GetDataHander s.u.GetParsingDataV1", mock.Anything, mock.Anything).Return()
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "test3",
			queryParams: map[string]string{
				"url": "https://google.com",
			},
			setupMocks: func(usecase *controllerMocks.UsecaseInterface, l *loggingMocks.LoggerInterface) {
				usecase.On("GetParsingDataV1", mock.Anything, mock.Anything).Return("", nil)
				l.On("Info", "GetDataHander", mock.Anything, mock.Anything).Return()
			},
			expectedStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			gin.SetMode(gin.TestMode)

			usecase := controllerMocks.NewUsecaseInterface(t)
			logger := loggingMocks.NewLoggerInterface(t)

			if tt.setupMocks != nil {
				tt.setupMocks(usecase, logger)
			}
			controllerTest := controller.New("0.0.0.0", 8081, usecase, logger)
			req, _ := http.NewRequest("GET", "/data", nil)
			q := req.URL.Query()
			for key, value := range tt.queryParams {
				q.Add(key, value)
			}
			req.URL.RawQuery = q.Encode()
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			controllerTest.GetDataHander(ctx)

			// Assertions
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
