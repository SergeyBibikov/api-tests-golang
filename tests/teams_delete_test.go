package tests

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type TeamsDeleteSuite struct {
	BaseSuite
}

func (ts *TeamsSuite) TestPositiveCases_Wrapper(t provider.T) {}
func (ts *TeamsSuite) TestNegativeCases_Wrapper(t provider.T) {}

func TestTeamsDelete(t *testing.T) {
	suite.RunSuite(t, new(TeamsSuite))
}
