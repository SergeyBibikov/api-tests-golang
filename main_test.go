package main

import (
	"encoding/json"
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
}
func (s *SmokeSuite) TestIsReady() {
	resp, err := s.client.R().Get("/ready")
	if err != nil {
		s.T().Errorf("The service is not ready %s", err)
	}
	var respB Readiness
	json.Unmarshal(resp.Body(), &respB)

	s.Equal(200, resp.StatusCode())
	s.Equal(respB.Status, "ready")
}

func TestSmoke(t *testing.T) {
	suite.Run(t, new(SmokeSuite))
}
