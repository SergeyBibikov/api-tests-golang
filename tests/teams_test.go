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
	t.Story("Positive")

	client := src.NewApiClient(&t, ts.client)
	teams, err := client.GetTeams(nil)
	t.Assert().Nil(err)
	t.Assert().Equal(len(teams), 30)
}

func (ts *TeamsSuite) TestNameFilter(t provider.T) {
	t.Story("Positive")

	m := make(map[string]string)
	m["name"] = "Denver Nuggets"

	client := src.NewApiClient(&t, ts.client)
	teams, err := client.GetTeams(m)
	r := client.Response

	t.Assert().Nil(err)
	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Equal(1, len(teams))
	t.Assert().Equal("West", teams[0].Conf)
	t.Assert().Equal("Northwest", teams[0].Div)
	t.Assert().Equal(1974, teams[0].Year)
}

func (ts *TeamsSuite) TestConferenceFilter(t provider.T) {

	testCases := []struct {
		conf string
	}{
		{conf: "East"},
		{conf: "West"},
	}
	for _, tC := range testCases {
		t.Story("Positive")
		tc := tC
		t.Run(tc.conf, func(t provider.T) {
			m := make(map[string]string)
			m["conference"] = tc.conf

			client := src.NewApiClient(&t, ts.client)
			teams, err := client.GetTeams(m)

			t.Assert().Nil(err)
			t.Assert().Equal(200, client.Response.StatusCode())
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

	client := src.NewApiClient(&t, ts.client)
	teams, err := client.GetTeams(m)

	t.Assert().Nil(err)
	t.Assert().Equal(200, client.Response.StatusCode())
	t.Assert().Equal(1, len(teams))
	t.Assert().Equal("Dallas Mavericks", teams[0].Name)
}

func (ts *TeamsSuite) TestNameFilterDoesntAllowOtherFilters(t provider.T) {
	t.Story("Negative")

	m := make(map[string]string)
	m["name"] = "Los Angeles Lakers"
	m["conference"] = "West"

	client := src.NewApiClient(&t, ts.client)
	teams, err := client.GetTeams(m)

	expectedMsg := "if name filter is present, other filters are not allowed"
	t.Assert().Nil(teams)
	t.Assert().Equal(400, client.Response.StatusCode())
	t.Assert().Equal(expectedMsg, err.Error())
}

func TestTeams(t *testing.T) {
	suite.RunSuite(t, new(TeamsSuite))
}
