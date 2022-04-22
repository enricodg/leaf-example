package serviceParameter

import (
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

type RepositorySuite struct {
	suite.Suite
	*require.Assertions
	ctrl       *gomock.Controller
	logger     *mock_leafLogger.MockLogger
	mysql      *mock_leafSql.MockORM
	resource   pkgResource.Resource
	repository Repository
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (s *RepositorySuite) SetupTest() {
	now := time.Now()
	leafTime.Mock(now)

	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	s.logger = mock_leafLogger.NewMockLogger(s.ctrl)
	s.mysql = mock_leafSql.NewMockORM(s.ctrl)
	s.resource = pkgResource.Resource{
		DatabaseSQL: injection.SQL{
			MySQL: s.mysql,
		},
		Log: s.logger,
	}
	s.repository = New(s.resource)
}

func (s *RepositorySuite) TearDownTest() {
	s.ctrl.Finish()
	leafTime.ResetMock()
}

func (s *RepositorySuite) TestRepository_New() {
	s.Equal(&repository{resource: s.resource}, New(s.resource))
}
