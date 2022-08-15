package main

import (
	"testing"

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

func (to *ValidateTokenSuite) TestValidAdminToken(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	r := validateToken(to.client, "Admin_token_Jack")
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Empty(resp)
}
func (to *ValidateTokenSuite) TestValidRegularUserToken(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	r := validateToken(to.client, "Regular_token_Steve")
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Empty(resp)
}

func (to *ValidateTokenSuite) TestValidPremiumUserToken(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	r := validateToken(to.client, "Premium_token_Mike")
	resp := responseBodyToMap(r.Body())

	t.Assert().Equal(200, r.StatusCode())
	t.Assert().Empty(resp)
}

func TestValidateToken(t *testing.T) {
	suite.RunSuite(t, new(ValidateTokenSuite))
}
