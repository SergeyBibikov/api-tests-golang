package tests

import (
	"fmt"
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

// TODO: check user was added to the db
func (to *RegistrationSuite) TestSuccessfulRegistration(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	u := src.GetRandomString(5)
	p := fmt.Sprintf("As%s", src.GetRandomString(6))
	e := fmt.Sprintf("%s@gmail.com", src.GetRandomString(5))
	r := src.Register(to.client, src.RegStruct{
		Username: u,
		Password: p,
		Email:    e})
	resp := src.ResponseBodyToMap(r.Body())

	t.Assert().Equal(201, r.StatusCode())
	t.Assert().Nil(resp["error"])
	t.Assert().Equal("user created", resp["message"])
}

func (to *RegistrationSuite) TestValidationOfEmptyUsername(t provider.T) {
	t.Parallel()
	t.Story("Negative")

	p := fmt.Sprintf("As%s", src.GetRandomString(6))
	e := fmt.Sprintf("%s@gmail.com", src.GetRandomString(5))
	r := src.Register(to.client, src.RegStruct{
		Username: "",
		Password: p,
		Email:    e})
	resp := src.ResponseBodyToMap(r.Body())

	t.Assert().Equal(400, r.StatusCode())
	t.Assert().Equal("Username, password and email are required", resp["error"])
}

func TestRegistration(t *testing.T) {
	suite.RunSuite(t, new(RegistrationSuite))
}
