package serviceParameter

import (
	"context"
	"fmt"
	"github.com/enricodg/leaf-example/internal/outbound/model"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
	"gorm.io/gorm"
)

func (s *RepositorySuite) Test_repository_Create() {
	var (
		email            = "test@test.com"
		serviceParameter = model.ServiceParameter{
			Variable:    "test",
			Value:       "test",
			Description: "test",
			BaseAuditable: model.BaseAuditable[uint64]{
				CreatedBy: email,
			},
		}
	)

	// test case
	tests := []struct {
		name         string
		ctxFunc      func() context.Context
		requestFunc  func() model.ServiceParameter
		expectFunc   func(s *RepositorySuite, ctx context.Context, serviceParameter model.ServiceParameter)
		wantResponse model.ServiceParameter
		wantError    error
	}{
		{
			name: "Success Case - [Outbound-Repository-ServiceParameter] Create",
			ctxFunc: func() context.Context {
				return context.Background()
			},
			requestFunc: func() model.ServiceParameter {
				return serviceParameter
			},
			expectFunc: func(s *RepositorySuite, ctx context.Context, serviceParameter model.ServiceParameter) {
				_, ctx = tracer.StartSpanFromContext(ctx, tracingCreate)

				s.mysql.EXPECT().WithContext(ctx).Return(s.mysql)
				s.mysql.EXPECT().Table(model.ServiceParameter{}.TableName()).Return(s.mysql)
				s.mysql.EXPECT().Create(ctx, &serviceParameter).Return(s.mysql)
				s.mysql.EXPECT().Error().Return(nil)
			},
			wantResponse: serviceParameter,
			wantError:    nil,
		},
		{
			name: "Failed Case - [Outbound-Repository-ServiceParameter] Create - Error Saving",
			ctxFunc: func() context.Context {
				return context.Background()
			},
			requestFunc: func() model.ServiceParameter {
				return serviceParameter
			},
			expectFunc: func(s *RepositorySuite, ctx context.Context, serviceParameter model.ServiceParameter) {
				_, ctx = tracer.StartSpanFromContext(ctx, tracingCreate)

				s.mysql.EXPECT().WithContext(ctx).Return(s.mysql)
				s.mysql.EXPECT().Table(model.ServiceParameter{}.TableName()).Return(s.mysql)
				s.mysql.EXPECT().Create(ctx, &serviceParameter).Return(s.mysql)
				s.mysql.EXPECT().Error().Return(gorm.ErrUnsupportedDriver)
				s.logger.EXPECT().Error(leafLogger.BuildMessage(ctx,
					fmt.Sprintf("%s error create: %v", tag, gorm.ErrUnsupportedDriver),
					leafLogger.WithAttr("serviceParameter", serviceParameter)))
			},
			wantResponse: model.ServiceParameter{},
			wantError:    gorm.ErrUnsupportedDriver,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			var ctx = context.Background()
			if tt.ctxFunc != nil {
				ctx = tt.ctxFunc()
			}

			var request model.ServiceParameter
			if tt.requestFunc != nil {
				request = tt.requestFunc()
			}

			// Assign EXPECT() func
			if tt.expectFunc != nil {
				tt.expectFunc(s, ctx, tt.requestFunc())
			}

			// ----- Main Logic -----
			response, err := s.repository.Create(ctx, s.mysql, request)

			// ----- Assertions -----
			s.Equal(tt.wantError, err)
			s.Equal(tt.wantResponse, response)
		})
	}
}
