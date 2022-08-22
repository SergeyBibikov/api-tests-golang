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

func (to *RegistrationSuite) TestValidationOfEmptyFields(t provider.T) {

	var testCases = []struct {
		testName string
		username string
		password string
		email    string
	}{
		{"Empty username", "", "asdfaEfsefs", "sadffa@google.com"},
		{"Empty password", "Fdfasdfasdfas", "", "sadffa@google.com"},
		{"Empty email", "asdfasdfasF", "asdfaEfsefs", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t provider.T) {

			t.Parallel()
			t.Story("Negative")

			r := src.Register(to.client, src.RegStruct{
				Username: tc.username,
				Password: tc.password,
				Email:    tc.email})
			resp := src.ResponseBodyToMap(r.Body())

			t.Assert().Equal(400, r.StatusCode())
			t.Assert().Equal("Username, password and email are required", resp["error"])
		})
	}
}

func (to *RegistrationSuite) TestValidationEmailFormat(t provider.T) {

	var testCases = []struct {
		testName string
		email    string
	}{
		{"No username", "@google.com"},
		{"Russian letters in username", "ВасяSmith@google.com"},
		{"No @ sign", "Bobbygoogle.com"},
		{"No mail server", "Jackie@.com"},
		{"No domain", "Jackie@google."},
		{"One letter domain", "Jackie@google.c"},
		{"Number in domain", "Jackie@google.c1"},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t provider.T) {

			t.Parallel()
			t.Story("Negative")

			r := src.Register(to.client, src.RegStruct{
				Username: src.GetRandomString(6),
				Password: src.GetRandomString(8),
				Email:    tc.email})
			resp := src.ResponseBodyToMap(r.Body())

			t.Assert().Equal(400, r.StatusCode())
			t.Assert().Equal("The email has an invalid format", resp["error"])
		})
	}
}

func TestRegistration(t *testing.T) {
	suite.RunSuite(t, new(RegistrationSuite))
}
