package serviceParameter

import (
	"github.com/enricodg/leaf-example/internal/outbound"
	"github.com/enricodg/leaf-example/internal/outbound/repositories"
	"github.com/enricodg/leaf-example/internal/usecases/shared"
	mock_repository_serviceParameter "github.com/enricodg/leaf-example/mocks/outbound/repositories/serviceParameter"
	mock_leafLogger "github.com/enricodg/leaf-example/mocks/resource/logger"
	mock_leafSql "github.com/enricodg/leaf-example/mocks/resource/mysql"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
	"github.com/enricodg/leaf-example/pkg/resource/injection"
	"github.com/golang/mock/gomock"
	leafTime "github.com/paulusrobin/leaf-utilities/time"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type UseCaseSuite struct {
	suite.Suite
	*require.Assertions
	ctrl   *gomock.Controller
	logger *mock_leafLogger.MockLogger
	mysql  *mock_leafSql.MockORM

	serviceParameterRepository *mock_repository_serviceParameter.MockRepository

	outbound outbound.Outbound
	resource pkgResource.Resource
	shared   shared.UseCase
	useCase  UseCase
}

func TestUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}

func (s *UseCaseSuite) SetupTest() {
	now := time.Now()
	leafTime.Mock(now)

	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	s.logger = mock_leafLogger.NewMockLogger(s.ctrl)
	s.mysql = mock_leafSql.NewMockORM(s.ctrl)

	s.serviceParameterRepository = mock_repository_serviceParameter.NewMockRepository(s.ctrl)

	s.outbound = outbound.Outbound{
		Repositories: repositories.Repositories{
			ServiceParameter: s.serviceParameterRepository,
		},
	}

	s.resource = pkgResource.Resource{
		DatabaseSQL: injection.SQL{
			MySQL: s.mysql,
		},
		Log: s.logger,
	}
	s.shared = shared.New(s.outbound, s.resource)
	s.useCase = New(s.outbound, s.shared, s.resource)
}
