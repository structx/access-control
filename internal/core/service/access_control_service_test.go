package service_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/structx/access-control/internal/core/service"
)

type AccessControlSuite struct {
	suite.Suite
	_ *service.AccessControl
}

func (suite *AccessControlSuite) SetupTest() {}

func TestAccessControlSuite(t *testing.T) {
	suite.Run(t, new(AccessControlSuite))
}
