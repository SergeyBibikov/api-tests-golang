package tests

import (
	"testing"

	"github.com/SergeyBibikov/api-tests-golang/src"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type RegistrationSuite struct {
	BaseSuite
}

func (to *RegistrationSuite) BeforeEach(t provider.T) {
	t.Epic("Registration tests")
}

func (to *RegistrationSuite) TestSuccessfulRegistration(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	r := src.Register(to.client, src.RegStruct{
		Username: "Alex",
		Password: "AaaaAaaa",
		Email:    "bob@gmail.com"})
	resp := src.ResponseBodyToMap(r.Body())

	t.Assert().Equal(201, r.StatusCode())
	t.Assert().Nil(resp["error"])
	t.Assert().Equal("user created", resp["message"])
}

func TestRegistration(t *testing.T) {
	suite.RunSuite(t, new(RegistrationSuite))
}
