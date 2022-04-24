package main

import (
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/go-resty/resty/v2"
)

type BaseSuite struct {
	suite.Suite
	client *resty.Client
}

func (s *BaseSuite) SetupSuite() {
	s.client = resty.New().SetBaseURL("http://localhost:8080")
	var err error
	var r *resty.Response
	for i := 0; i < 20; i++ {
		r, err = s.client.R().Get("/ready")
		if r.StatusCode() == 200 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	if err != nil {
		s.T().Errorf("The service was not available for 2 seconds %s", err)
	}
}
