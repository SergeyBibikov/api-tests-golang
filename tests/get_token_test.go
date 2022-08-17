package tests

import (
	"fmt"
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
func (to *GetTokenSuite) TestGetAdminToken(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	username := "Jack"
	password := "JackPass"

	r := src.GetToken(to.client, username, password)
	resp := src.ResponseBodyToMap(r.Body())

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Equal(resp["token"], fmt.Sprintf("Admin_token_%s", username))
}

func (to *GetTokenSuite) TestGetRegularUserToken(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	username := "Steve"
	password := "StevePass"

	r := src.GetToken(to.client, username, password)
	resp := src.ResponseBodyToMap(r.Body())

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Equal(resp["token"], fmt.Sprintf("Regular_token_%s", username))
}

func (to *GetTokenSuite) TestGetPremiumUserToken(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	username := "Mike"
	password := "MikePass"

	r := src.GetToken(to.client, username, password)
	resp := src.ResponseBodyToMap(r.Body())

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Equal(resp["token"], fmt.Sprintf("Premium_token_%s", username))
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
