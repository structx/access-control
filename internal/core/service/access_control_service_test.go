package service_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/trevatk/anastasia/internal/core/service"
)

type AccessControlSuite struct {
	suite.Suite
	ac *service.AccessControl
}

func (suite *AccessControlSuite) SetupTest() {}

func TestAccessControlSuite(t *testing.T) {
	suite.Run(t, new(AccessControlSuite))
}
