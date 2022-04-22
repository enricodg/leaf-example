package serviceParameter

import (
	"context"
	"errors"
	"github.com/enricodg/leaf-example/internal/inbound/model"
	ucModel "github.com/enricodg/leaf-example/internal/usecases/model"
	"github.com/labstack/echo/v4"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	leafValidatorV10 "github.com/paulusrobin/leaf-utilities/validator/integrations/v10"
	"net/http"
	"net/http/httptest"
	"strings"
)

func (s *ControllerSuite) TestController_Create() {
	var (
		method  = http.MethodPost
		path    = "/"
		request = model.ServiceParameterCreateRequest{
			Variable:    "test",
			Value:       "test",
			Description: "test",
		}
		serviceParameter = ucModel.ServiceParameterUpsertResponse{
			ID:          1,
			Variable:    "test",
			Value:       "test",
			Description: "test",
		}
	)

	tests := []struct {
		name            string
		request         interface{}
		echoContextFunc func(eCtx echo.Context) echo.Context
		expectFunc      func(s *ControllerSuite, ctx context.Context)
		wantErr         error
		wantStatusCode  int
		wantResponse    string
	}{
		{
			name:    "Success Case - [Inbound-HTTP-ServiceParameter] Create",
			request: request,
			expectFunc: func(s *ControllerSuite, ctx context.Context) {
				s.serviceParameterUseCase.EXPECT().
					Create(ctx, request.ToUCModel()).
					Return(serviceParameter, nil)
			},
			wantErr:        nil,
			wantStatusCode: http.StatusOK,
			wantResponse:   `{"data":{"id":1,"variable":"test","value":"test","description":"test"}}` + "\n",
		},
		{
			name:           "Failed Case - [Inbound-HTTP-ServiceParameter] Create - Empty Request Body",
			request:        nil,
			wantErr:        nil,
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   `{"error":{"message":"Key: 'ServiceParameterCreateRequest.Variable' Error:Field validation for 'Variable' failed on the 'required' tag","code":400,"errors":[{"field":"Variable","reason":"Key: 'ServiceParameterCreateRequest.Variable' Error:Field validation for 'Variable' failed on the 'required' tag"},{"field":"Value","reason":"Key: 'ServiceParameterCreateRequest.Value' Error:Field validation for 'Value' failed on the 'required' tag"},{"field":"Description","reason":"Key: 'ServiceParameterCreateRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag"}]}}` + "\n",
		},
		{
			name:    "Failed Case - [Inbound-HTTP-ServiceParameter] Create - Error From UseCase",
			request: request,
			expectFunc: func(s *ControllerSuite, ctx context.Context) {
				s.serviceParameterUseCase.EXPECT().
					Create(ctx, request.ToUCModel()).
					Return(ucModel.ServiceParameterUpsertResponse{}, errors.New("connection timeout"))
			},
			wantStatusCode: http.StatusInternalServerError,
			wantResponse:   `{"error":{"message":"connection timeout","code":500}}` + "\n",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			requestByte, _ := json.Marshal(tt.request)

			req := httptest.NewRequest(method, path, strings.NewReader(string(requestByte)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			resRecorder := httptest.NewRecorder()

			newEcho := echo.New()
			newEcho.Validator, _ = leafValidatorV10.New()
			echoContext := newEcho.NewContext(req, resRecorder)

			if tt.echoContextFunc != nil {
				echoContext = tt.echoContextFunc(echoContext)
			}

			// Assign EXPECT() func
			if tt.expectFunc != nil {
				tt.expectFunc(s, echoContext.Request().Context())
			}

			// ----- Main Logic -----
			err := s.controller.Create(echoContext)

			// ----- Assertions -----
			s.Equal(tt.wantErr, err)
			s.Equal(tt.wantStatusCode, resRecorder.Result().StatusCode)
			s.Equal(tt.wantResponse, resRecorder.Body.String())
		})
	}
}
