package tests

import (
	"encoding/json"
	"testing"

	"github.com/SergeyBibikov/api-tests-golang/src"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type TeamsSuite struct {
	BaseSuite
}

func (ts *TeamsSuite) TestTeamsQty(t provider.T) {
	t.Story("Positive")

	r := src.GetTeams(ts.client, nil)

	var tr map[string][]src.Team
	json.Unmarshal(r.Body(), &tr)

	t.Assert().Equal(len(tr["results"]), 30)
}

func (ts *TeamsSuite) TestNameFilterDoesntAllowOtherFilters(t provider.T) {
	t.Story("Negative")

	m := make(map[string]string)
	m["name"] = "Los Angeles Lakers"
	m["conference"] = "West"
	r := src.GetTeams(ts.client, m)
	resp := src.ResponseBodyToMap(r.Body())

	expectedMsg := "if name filter is present, other filters are not allowed"
	t.Assert().Equal(400, r.StatusCode())
	t.Assert().Equal(expectedMsg, resp["error"])
}

func TestTeams(t *testing.T) {
	suite.RunSuite(t, new(TeamsSuite))
}
