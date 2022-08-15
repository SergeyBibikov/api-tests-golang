package main

import (
	"fmt"
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

	username := "Jack"
	password := "JackPass"

	r := getToken(to.client, username, password)
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Equal(resp["token"], fmt.Sprintf("Admin_token_%s", username))
}

func (to *TokenSuite) TestGetRegularUserToken(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	username := "Steve"
	password := "StevePass"

	r := getToken(to.client, username, password)
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Equal(resp["token"], fmt.Sprintf("Regular_user_token_%s", username))
}

func (to *TokenSuite) TestGetPremiumUserToken(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	username := "Mike"
	password := "MikePass"

	r := getToken(to.client, username, password)
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Equal(resp["token"], fmt.Sprintf("Premium_user_token_%s", username))
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
