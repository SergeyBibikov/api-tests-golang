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

func (to *GetTokenSuite) Test_Positive_GetToken_Wrapper(t provider.T) {
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
			resp := client.GetToken(src.GetTokenRequest{Username: tC.username, Password: tC.password})

			t.Assert().Equal(200, resp.StatusCode)
			t.Assert().Equal(tC.expectedToken, resp.Token)
		})
	}
}
func (to *GetTokenSuite) Test_Negative_GetToken_Wrapper(t provider.T) {
	testCases := []struct {
		testName      string
		username      string
		password      string
		expectedError string
	}{
		{
			testName:      "Wrong username",
			username:      "Mike1",
			password:      "MikePass",
			expectedError: "invalid username and/or password",
		},
		{
			testName:      "Wrong password",
			username:      "Mike",
			password:      "MikePass1",
			expectedError: "invalid username and/or password",
		},
		{
			testName:      "No password",
			username:      "Mike",
			password:      "",
			expectedError: "Password is a required field",
		},
		{
			testName:      "No username",
			username:      "",
			password:      "MikePass",
			expectedError: "Username is a required field",
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.testName, func(t provider.T) {
			t.Story("Negative")

			client := src.NewApiClient(&t)
			resp := client.GetToken(src.GetTokenRequest{Username: tC.username, Password: tC.password})

			t.Assert().Equal(400, resp.StatusCode)
			t.Assert().Equal(tC.expectedError, resp.Error)
		})
	}
}

func TestGetToken(t *testing.T) {
	suite.RunSuite(t, new(GetTokenSuite))
}
