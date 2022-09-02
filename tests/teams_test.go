package tests

import (
	"encoding/json"
	"testing"

	"github.com/SergeyBibikov/api-tests-golang/src"
	"github.com/ozontech/allure-go/pkg/allure"
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

func (ts *TeamsSuite) TestNameFilter(t provider.T) {
	t.Story("Positive")

	m := make(map[string]string)
	m["name"] = "Denver Nuggets"

	r := src.GetTeams(ts.client, m)

	var _t map[string][]src.Team
	json.Unmarshal(r.Body(), &_t)
	teams := _t["results"]

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Equal(1, len(teams))
	t.Assert().Equal("West", teams[0].Conf)
	t.Assert().Equal("Northwest", teams[0].Div)
	t.Assert().Equal(1974, teams[0].Year)
}

func (ts *TeamsSuite) TestConferenceFilter(t provider.T) {
	t.Story("Positive")

	testCases := []struct {
		conf string
	}{
		{conf: "East"},
		{conf: "West"},
	}
	for _, tC := range testCases {
		tc := tC
		t.Run(tc.conf, func(t provider.T) {
			m := make(map[string]string)
			m["conference"] = tc.conf

			t.WithNewStep("Send request", func(sCtx provider.StepCtx) {}, allure.NewParameter("body", m))
			r := src.GetTeams(ts.client, m)
			var _t map[string][]src.Team
			json.Unmarshal(r.Body(), &_t)
			teams := _t["results"]
			t.Assert().Equal(200, r.StatusCode())
			t.Assert().Equal(15, len(teams))
		})
	}

}

func (ts *TeamsSuite) TestAllFiltersButName(t provider.T) {
	t.Story("Positive")

	m := make(map[string]string)
	m["conference"] = "West"
	m["division"] = "Southwest"
	m["est_year"] = "1980"

	r := src.GetTeams(ts.client, m)

	var _t map[string][]src.Team
	json.Unmarshal(r.Body(), &_t)
	teams := _t["results"]

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Equal(1, len(teams))
	t.Assert().Equal("Dallas Mavericks", teams[0].Name)
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
