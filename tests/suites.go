package tests

import (
	"time"

	"github.com/SergeyBibikov/api-tests-golang/src"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type BaseSuite struct {
	suite.Suite
}

func (s *BaseSuite) BeforeAll(t provider.T) {
	client := src.NewApiClient(&t)
	for i := 0; i < 20; i++ {
		ok := client.Ready()
		if ok {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	t.Fatal("The service was not available for 2 seconds")
}
