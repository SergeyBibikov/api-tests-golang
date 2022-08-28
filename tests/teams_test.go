package tests

import (
	"testing"

	"github.com/SergeyBibikov/api-tests-golang/src"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type TeamsSuite struct {
	BaseSuite
}

func (ts *TeamsSuite) TestTeamsQty(t provider.T) {
	r := src.GetTeams(ts.client)

	t.Assert().Equal(len(r), 30)
}

func TestTeams(t *testing.T) {
	suite.RunSuite(t, new(TeamsSuite))
}
