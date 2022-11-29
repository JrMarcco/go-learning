package service

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type svcTestSuite struct {
	suite.Suite
}

func TestSvc(t *testing.T) {
	suite.Run(t, &svcTestSuite{})
}

func (s *svcTestSuite) SetupSuite() {
	Setup()
}

func (s *svcTestSuite) TearDownSuite() {
	defer func() {
		TearDown()
	}()
}
