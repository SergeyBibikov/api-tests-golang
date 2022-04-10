package main

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/go-resty/resty/v2"
)

type SmokeSuite struct {
	suite.Suite
	client *resty.Client
}

func (s *SmokeSuite) SetupSuite() {
	s.client = resty.New().SetBaseURL("http://localhost:8080")
	_, err := s.client.R().Get("/ready")
	if err != nil {
		s.T().Errorf("The service is not ready %s", err)
	}
}

//TODO: Add tests for token validation and move token tests to a separate suite
func (s *SmokeSuite) TestAuthValidData() {

	r := getToken(s.client, "valid_user", "valid_password")
	tokens := responseBodyToMap(r.Body())

	s.Equal(201, r.StatusCode())
	s.True(tokens["success"].(bool))
	s.Equal(tokens["refreshToken"].(string), "ref_resh")
	s.Equal(tokens["accessToken"].(string), "access_t")
}
func (s *SmokeSuite) TestAuthWrongUsername() {

	r := getToken(s.client, "valid", "valid_password")
	tokens := responseBodyToMap(r.Body())

	s.Equal(401, r.StatusCode())
	s.False(tokens["success"].(bool))
	s.Empty(tokens["refreshToken"].(string))
	s.Empty(tokens["accessToken"].(string))
}
func (s *SmokeSuite) TestAuthWrongPassword() {

	r := getToken(s.client, "valid_user", "valid")
	tokens := responseBodyToMap(r.Body())

	s.Equal(401, r.StatusCode())
	s.False(tokens["success"].(bool))
	s.Empty(tokens["refreshToken"].(string))
	s.Empty(tokens["accessToken"].(string))
}

func TestSmoke(t *testing.T) {
	suite.Run(t, new(SmokeSuite))
}
