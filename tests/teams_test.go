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

func (ts *TeamsSuite) TestAllTeamsQty(t provider.T) {
	t.Story("Positive")

	client := src.NewApiClient(&t)
	resp := client.GetTeams(nil)
	t.Assert().Equal(len(resp.Teams), 30)
}

func (ts *TeamsSuite) TestNameFilter(t provider.T) {
	t.Story("Positive")

	m := make(map[string]string)
	m["name"] = "Denver Nuggets"

	client := src.NewApiClient(&t)
	resp := client.GetTeams(m)
	r := client.Response

	t.Assert().Equal(200, r.StatusCode)
	t.Assert().Equal(1, len(resp.Teams))
	t.Assert().Equal("West", resp.Teams[0].Conf)
	t.Assert().Equal("Northwest", resp.Teams[0].Div)
	t.Assert().Equal(1974, resp.Teams[0].Year)
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

			client := src.NewApiClient(&t)
			resp := client.GetTeams(m)

			t.Assert().Equal(200, resp.StatusCode)
			t.Assert().Equal(15, len(resp.Teams))
		})
	}

}

func (ts *TeamsSuite) TestAllFiltersButName(t provider.T) {
	t.Story("Positive")

	m := make(map[string]string)
	m["conference"] = "West"
	m["division"] = "Southwest"
	m["est_year"] = "1980"

	client := src.NewApiClient(&t)
	resp := client.GetTeams(m)

	t.Assert().Equal(200, resp.StatusCode)
	t.Assert().Equal(1, len(resp.Teams))
	t.Assert().Equal("Dallas Mavericks", resp.Teams[0].Name)
}

func (ts *TeamsSuite) TestUnsupportedFilter(t provider.T) {
	t.Story("Negative")

	m := make(map[string]string)
	m["aname"] = "Denver Nuggets"

	client := src.NewApiClient(&t)
	resp := client.GetTeams(m)

	t.Assert().Equal(len(resp.Teams), 30)
}

func (ts *TeamsSuite) TestNameFilterDoesntAllowOtherFilters(t provider.T) {
	t.Story("Negative")

	m := make(map[string]string)
	m["name"] = "Los Angeles Lakers"
	m["conference"] = "West"

	client := src.NewApiClient(&t)
	resp := client.GetTeams(m)

	expectedMsg := "if name filter is present, other filters are not allowed"
	t.Assert().Nil(resp.Teams)
	t.Assert().Equal(400, resp.StatusCode)
	t.Assert().Equal(expectedMsg, resp.Error)
}

func (ts *TeamsSuite) Test_NoResultsMatchingFilter_Wrapper(t provider.T) {
	testCases := []struct {
		testName    string
		filterName  string
		filterValue string
	}{
		{
			testName:    "No team with name Los Angeles Likers",
			filterName:  "name",
			filterValue: "Los Angeles Likers",
		},
		{
			testName:    "No team with conference Weast",
			filterName:  "conference",
			filterValue: "Weast",
		},
		{
			testName:    "No team with division Northweast",
			filterName:  "division",
			filterValue: "Northweast",
		},
		{
			testName:    "No team with est_year = 2022",
			filterName:  "est_year",
			filterValue: "2022",
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.testName, func(t provider.T) {
			m := make(map[string]string)
			m[tC.filterName] = tC.filterValue

			client := src.NewApiClient(&t)
			resp := client.GetTeams(m)

			t.Assert().Equal(200, resp.StatusCode)
			t.Assert().Empty(resp.Teams)
		})
	}
}

func TestTeams(t *testing.T) {
	suite.RunSuite(t, new(TeamsSuite))
}
