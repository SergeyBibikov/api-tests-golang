package tests

import (
	"testing"

	"github.com/SergeyBibikov/api-tests-golang/src"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type ValidateTokenSuite struct {
	BaseSuite
}

func (to *ValidateTokenSuite) BeforeEach(t provider.T) {
	t.Epic("Token tests")
	t.Feature("Validate token")
}

func (to *ValidateTokenSuite) Test_SuccessfulValidation_Wrapper(t provider.T) {
	testCases := []struct {
		testName      string
		tokenToVerify string
	}{
		{
			testName:      "Admin token",
			tokenToVerify: "Admin_token_Jack",
		},
		{
			testName:      "Regular user token",
			tokenToVerify: "Regular_token_Steve",
		},
		{
			testName:      "Premium user token",
			tokenToVerify: "Premium_token_Mike",
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.testName, func(t provider.T) {
			t.Parallel()
			t.Story("Positive")

			client := src.NewApiClient(&t)
			resp := client.ValidateToken(tC.tokenToVerify)

			t.Assert().Equal(200, resp.StatusCode)
			t.Assert().Empty(resp.Error)
		})
	}
}
func (to *ValidateTokenSuite) Test_FailedValidation_Wrapper(t provider.T) {
	testCases := []struct {
		testName           string
		tokenToVerify      string
		expectedStatusCode int
		expectedErrorText  string
	}{
		{
			testName:           "Invalid token format: one part",
			tokenToVerify:      "PremiumtokenMike",
			expectedErrorText:  "Incorrect token format. Proper format: role_token_username",
			expectedStatusCode: 400,
		},
		{
			testName:           "Invalid token format: two parts",
			tokenToVerify:      "Premium_token",
			expectedErrorText:  "Incorrect token format. Proper format: role_token_username",
			expectedStatusCode: 400,
		},
		{
			testName:           "Invalid token format: four parts",
			tokenToVerify:      "Premium_token_Mike_more",
			expectedErrorText:  "Incorrect token format. Proper format: role_token_username",
			expectedStatusCode: 400,
		},
		{
			testName:           "Non existing user",
			tokenToVerify:      "Premium_token_Bob",
			expectedErrorText:  "invalid username",
			expectedStatusCode: 401,
		},
		{
			testName:           "Incorrect user role",
			tokenToVerify:      "Premium_token_Jack",
			expectedErrorText:  "incorrect user role",
			expectedStatusCode: 401,
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.testName, func(t provider.T) {
			t.Parallel()
			t.Story("Negative")

			client := src.NewApiClient(&t)
			resp := client.ValidateToken(tC.tokenToVerify)

			t.Assert().Equal(tC.expectedStatusCode, resp.StatusCode)
			t.Assert().Equal(tC.expectedErrorText, resp.Error)
		})
	}

}

func TestValidateToken(t *testing.T) {
	suite.RunSuite(t, new(ValidateTokenSuite))
}
