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

func (to *RegistrationSuite) TestSuccessfulRegistration(t provider.T) {
	t.Parallel()
	t.Story("Positive")

	u := src.GetRandomString(5)
	p := fmt.Sprintf("As%s", src.GetRandomString(6))
	e := fmt.Sprintf("%s@gmail.com", src.GetRandomString(5))

	client := src.NewApiClient(&t)
	resp := client.Register(src.RegStruct{
		Username: u,
		Password: p,
		Email:    e})

	c, _ := src.NewDbClient()
	m := make(map[string]string)
	m["id"] = fmt.Sprint(resp.UserId)
	users := c.GetUsers(m)

	t.Assert().Equal(201, resp.StatusCode)
	t.Assert().Equal("", resp.Error)
	t.Assert().Equal("user created", resp.Message)
	t.Assert().Equal(u, users[0].Username)

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

			client := src.NewApiClient(&t)
			resp := client.Register(b)

			t.Assert().Equal(resp.Message, "")
			t.Assert().Equal(400, resp.StatusCode)
			t.Assert().Equal("Username, password and email are required", resp.Error)
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

			client := src.NewApiClient(&t)
			resp := client.Register(b)

			t.Assert().Equal(resp.Message, "")
			t.Assert().Equal(400, resp.StatusCode)
			t.Assert().Equal("The email has an invalid format", resp.Error)
		})
	}
}

func TestRegistration(t *testing.T) {
	suite.RunSuite(t, new(RegistrationSuite))
}
