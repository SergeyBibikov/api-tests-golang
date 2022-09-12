package tests

import (
	"testing"

	"github.com/SergeyBibikov/api-tests-golang/src"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type GetTokenSuite struct {
	BaseSuite
}

func (to *GetTokenSuite) BeforeEach(t provider.T) {
	t.Epic("Token tests")
	t.Feature("Get token")
}

func (to *GetTokenSuite) Test_SuccessfullGetToken_Wrapper(t provider.T) {
	testCases := []struct {
		testName      string
		username      string
		password      string
		expectedToken string
	}{
		{
			testName:      "Get admin token",
			username:      "Jack",
			password:      "JackPass",
			expectedToken: "Admin_token_Jack",
		},
		{
			testName:      "Get regular user token",
			username:      "Steve",
			password:      "StevePass",
			expectedToken: "Regular_token_Steve",
		},
		{
			testName:      "Get premium user token",
			username:      "Mike",
			password:      "MikePass",
			expectedToken: "Premium_token_Mike",
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.testName, func(t provider.T) {
			t.Story("Positive")

			client := src.NewApiClient(&t)
			resp := client.GetToken(src.GetTokenRequest{tC.username, tC.password})

			t.Assert().Equal(200, resp.StatusCode)
			t.Assert().Equal(tC.expectedToken, resp.Token)
		})
	}
}

func (to *GetTokenSuite) TestGetTokenWithWrongUsername(t provider.T) {
	t.Parallel()
	t.Story("Negative")

	r := src.GetToken(to.client, "Mike1", "MikePass")
	resp := src.ResponseBodyToMap(r.Body())

	t.Assert().Equal(400, r.StatusCode())
	t.Assert().Equal(resp["error"], "invalid username and/or password")
}

func (to *GetTokenSuite) TestGetTokenWithWrongPassword(t provider.T) {
	t.Parallel()
	t.Story("Negative")

	r := src.GetToken(to.client, "Mike", "MikePass1")
	resp := src.ResponseBodyToMap(r.Body())

	t.Assert().Equal(400, r.StatusCode())
	t.Assert().Equal(resp["error"], "invalid username and/or password")
}

func (to *GetTokenSuite) TestGetTokenWithoutPasswordInBody(t provider.T) {
	t.Parallel()
	t.Story("Negative")

	r := src.GetToken(to.client, "Mike", "")
	resp := src.ResponseBodyToMap(r.Body())

	t.Assert().Equal(400, r.StatusCode())
	t.Assert().Equal(resp["error"], "Password is a required field")
}

func (to *GetTokenSuite) TestGetTokenWithoutUsernameInBody(t provider.T) {
	t.Parallel()
	t.Story("Negative")

	r := src.GetToken(to.client, "", "MikePass1")
	resp := src.ResponseBodyToMap(r.Body())

	t.Assert().Equal(400, r.StatusCode())
	t.Assert().Equal(resp["error"], "Username is a required field")
}

func TestGetToken(t *testing.T) {
	suite.RunSuite(t, new(GetTokenSuite))
}
