package main

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type TokenSuite struct {
	BaseSuite
}

func (to *TokenSuite) BeforeEach(t provider.T) {
	t.Epic("Token tests")
	t.Feature("Get token")
}
func (to *TokenSuite) TestGetAdminToken(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	r := getToken(to.client, "Jack", "JackPass")
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Equal(resp["token"], "Admin_token")
}

func (to *TokenSuite) TestGetRegularUserToken(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	r := getToken(to.client, "Steve", "StevePass")
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Equal(resp["token"], "Regular_user_token")
}

func (to *TokenSuite) TestGetPremiumUserToken(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	r := getToken(to.client, "Mike", "MikePass")
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Equal(resp["token"], "Premium_user_token")
}

func (to *TokenSuite) TestGetTokenWithWrongUsername(t provider.T) {
	t.Parallel()
	t.Story("Negative")

	r := getToken(to.client, "Mike1", "MikePass")
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(400, r.StatusCode())
	t.Assert().Equal(resp["error"], "invalid username and/or password")
}

func (to *TokenSuite) TestGetTokenWithWrongPassword(t provider.T) {
	t.Parallel()
	t.Story("Negative")

	r := getToken(to.client, "Mike", "MikePass1")
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(400, r.StatusCode())
	t.Assert().Equal(resp["error"], "invalid username and/or password")
}

func (to *TokenSuite) TestGetTokenWithoutPasswordInBody(t provider.T) {
	t.Parallel()
	t.Story("Negative")

	r := getToken(to.client, "Mike", "")
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(400, r.StatusCode())
	t.Assert().Equal(resp["error"], "Password is a required field")
}

func (to *TokenSuite) TestGetTokenWithoutUsernameInBody(t provider.T) {
	t.Parallel()
	t.Story("Negative")

	r := getToken(to.client, "", "MikePass1")
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(400, r.StatusCode())
	t.Assert().Equal(resp["error"], "Username is a required field")
}

// func (to *TokenSuite) TestCheckValidToken(t provider.T) {
// 	t.Parallel()

// 	r := checkToken(to.client, "access_t")
// 	resp := responseBodyToMap(r.Body())

// 	t.Assert().Equal(200, r.StatusCode())
// 	t.Assert().True(resp["valid"].(bool))
// }
// func (to *TokenSuite) TestCheckInvalidToken(t provider.T) {
// 	t.Parallel()

// 	r := checkToken(to.client, "access_tok")
// 	resp := responseBodyToMap(r.Body())

// 	t.Assert().Equal(400, r.StatusCode())
// 	t.Assert().False(resp["valid"].(bool))
// }

func TestTokens(t *testing.T) {
	suite.RunSuite(t, new(TokenSuite))
}
