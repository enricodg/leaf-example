package serviceParameter

import (
	"context"
	"fmt"
	obModel "github.com/enricodg/leaf-example/internal/outbound/model"
	ucModel "github.com/enricodg/leaf-example/internal/usecases/model"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
	"gorm.io/gorm"
)

func (s *UseCaseSuite) Test_ucServiceParameter_Create() {
	var (
		email   = "test@test.com"
		request = ucModel.ServiceParameterUpsertRequest{
			Variable:    "test",
			Value:       "test",
			Description: "test",
		}
		response = ucModel.ServiceParameterUpsertResponse{
			ID:          1,
			Variable:    "test",
			Value:       "test",
			Description: "test",
		}
		serviceParameter = obModel.ServiceParameter{
			Variable:    "test",
			Value:       "test",
			Description: "test",
			BaseAuditable: obModel.BaseAuditable[uint64]{
				CreatedBy: email,
			},
		}
		serviceParameterSaved = obModel.ServiceParameter{
			Variable:    "test",
			Value:       "test",
			Description: "test",
			BaseAuditable: obModel.BaseAuditable[uint64]{
				ID:        1,
				CreatedBy: email,
			},
		}
	)

	// test case
	tests := []struct {
		name         string
		ctxFunc      func() context.Context
		requestFunc  func() ucModel.ServiceParameterUpsertRequest
		expectFunc   func(s *UseCaseSuite, ctx context.Context, request ucModel.ServiceParameterUpsertRequest)
		wantResponse ucModel.ServiceParameterUpsertResponse
		wantError    error
	}{
		{
			name: "Success Case - [UseCase-ServiceParameter] Create",
			ctxFunc: func() context.Context {
				ctx := context.Background()
				builder, _ := leafMandatory.NewMandatoryBuilder()
				return leafMandatory.Context(ctx, builder.
					WithUser(uint64(1), email).
					Build())
			},
			requestFunc: func() ucModel.ServiceParameterUpsertRequest {
				return request
			},
			expectFunc: func(s *UseCaseSuite, ctx context.Context, request ucModel.ServiceParameterUpsertRequest) {
				_, ctx = tracer.StartSpanFromContext(ctx, tracingCreate)

				s.serviceParameterRepository.EXPECT().Create(ctx, s.mysql, serviceParameter).
					Return(serviceParameterSaved, nil)
			},
			wantResponse: response,
			wantError:    nil,
		},
		{
			name: "Failed Case - [UseCase-ServiceParameter] Create - Error",
			ctxFunc: func() context.Context {
				ctx := context.Background()
				builder, _ := leafMandatory.NewMandatoryBuilder()
				return leafMandatory.Context(ctx, builder.
					WithUser(uint64(1), email).
					Build())
			},
			requestFunc: func() ucModel.ServiceParameterUpsertRequest {
				return request
			},
			expectFunc: func(s *UseCaseSuite, ctx context.Context, request ucModel.ServiceParameterUpsertRequest) {
				_, ctx = tracer.StartSpanFromContext(ctx, tracingCreate)

				s.serviceParameterRepository.EXPECT().Create(ctx, s.mysql, serviceParameter).
					Return(obModel.ServiceParameter{}, gorm.ErrInvalidData)
				s.logger.EXPECT().Error(leafLogger.BuildMessage(ctx,
					fmt.Sprintf("%s error on create: %v", tag, gorm.ErrInvalidData)))
			},
			wantResponse: ucModel.ServiceParameterUpsertResponse{},
			wantError:    gorm.ErrInvalidData,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			var ctx = context.Background()
			if tt.ctxFunc != nil {
				ctx = tt.ctxFunc()
			}

			var request ucModel.ServiceParameterUpsertRequest
			if tt.requestFunc != nil {
				request = tt.requestFunc()
			}

			if tt.expectFunc != nil {
				tt.expectFunc(s, ctx, request)
			}

			// ----- Main Logic -----
			response, err := s.useCase.Create(ctx, request)

			// ----- Assertions -----
			s.Equal(tt.wantError, err)
			s.Equal(tt.wantResponse, response)
		})
	}
}
