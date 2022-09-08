package tests

import (
	"fmt"
	"testing"

	"github.com/SergeyBibikov/api-tests-golang/src"
	"github.com/ozontech/allure-go/pkg/allure"
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
// TODO: replace Register func in all tests
func (to *RegistrationSuite) TestSuccessfulRegistration(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	u := src.GetRandomString(5)
	p := fmt.Sprintf("As%s", src.GetRandomString(6))
	e := fmt.Sprintf("%s@gmail.com", src.GetRandomString(5))

	client := src.NewApiClient(&t, to.client)
	message, err := client.Register(src.RegStruct{
		Username: u,
		Password: p,
		Email:    e})

	t.Assert().Equal(201, client.Response.StatusCode())
	t.Assert().Nil(err)
	t.Assert().Equal("user created", message)
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
		tc := tc
		t.Run(tc.testName, func(t provider.T) {

			t.Parallel()
			t.Story("Negative")

			b := src.RegStruct{
				Username: tc.username,
				Password: tc.password,
				Email:    tc.email}

			t.WithNewStep("Send request", func(sCtx provider.StepCtx) {}, allure.NewParameter("body", b))

			r := src.Register(to.client, b)
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
		tc := tc
		t.Run(tc.testName, func(t provider.T) {

			t.Parallel()
			t.Story("Negative")

			b := src.RegStruct{
				Username: src.GetRandomString(6),
				Password: src.GetRandomString(8),
				Email:    tc.email}
			client := src.NewApiClient(&t, to.client)
			message, error := client.Register(b)

			t.Assert().Equal(message, "")
			t.Assert().Equal(400, client.Response.StatusCode())
			t.Assert().Equal("The email has an invalid format", error.Error())
		})
	}
}

func TestRegistration(t *testing.T) {
	suite.RunSuite(t, new(RegistrationSuite))
}
