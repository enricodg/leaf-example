package serviceParameter

import (
	"github.com/enricodg/leaf-example/internal/usecases"
	mock_usecase_serviceParameter "github.com/enricodg/leaf-example/mocks/usecases/serviceParameter"
	"github.com/golang/mock/gomock"
	leafTime "github.com/paulusrobin/leaf-utilities/time"
	leafTranslatorValidator "github.com/paulusrobin/leaf-utilities/translator/validator"
	leafValidatorV10 "github.com/paulusrobin/leaf-utilities/validator/integrations/v10"
	leafValidator "github.com/paulusrobin/leaf-utilities/validator/validator"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ControllerSuite struct {
	suite.Suite
	*require.Assertions
	ctrl                    *gomock.Controller
	useCases                usecases.UseCase
	serviceParameterUseCase *mock_usecase_serviceParameter.MockUseCase
	leafValidator           leafValidator.Validator

	controller Controller
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerSuite))
}

func (s *ControllerSuite) SetupTest() {
	leafTime.Mock(time.Date(2021, time.December, 24, 0, 0, 0, 0, time.Local))

	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	s.serviceParameterUseCase = mock_usecase_serviceParameter.NewMockUseCase(s.ctrl)
	s.useCases = usecases.UseCase{
		ServiceParameter: s.serviceParameterUseCase,
	}

	validator, e := leafValidatorV10.New(leafValidatorV10.WithTranslator(leafTranslatorValidator.GetTranslator()))
	if e == nil {
		s.leafValidator = validator
	}

	s.controller = Controller{
		UseCase: s.useCases,
	}
}

func (s *ControllerSuite) TearDownTest() {
	s.ctrl.Finish()
	leafTime.ResetMock()
}
