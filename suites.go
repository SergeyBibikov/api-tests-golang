package main

import (
	"github.com/stretchr/testify/suite"

	"github.com/go-resty/resty/v2"
)

type BaseSuite struct {
	suite.Suite
	client *resty.Client
}

func (s *BaseSuite) SetupSuite() {
	s.client = resty.New().SetBaseURL("http://localhost:8080")
	_, err := s.client.R().Get("/ready")
	if err != nil {
		s.T().Errorf("The service is not ready %s", err)
	}
}
