package main

// import (
// 	"testing"

// 	"github.com/ozontech/allure-go/pkg/framework/provider"
// 	"github.com/ozontech/allure-go/pkg/framework/suite"
// )

// type SmokeSuite struct {
// 	BaseSuite
// }

// // TODO: Add tests for token validation and move token tests to a separate suite
// // TODO: Add users test
// func (s *SmokeSuite) TestAuthValidData(t provider.T) {

// 	r := getToken(s.client, "valid_user", "valid_password")
// 	tokens := responseBodyToMap(r.Body())

// 	t.Assert().Equal(201, r.StatusCode())
// 	t.Assert().True(tokens["success"].(bool))
// 	t.Assert().Equal(tokens["refreshToken"].(string), "ref_resh")
// 	t.Assert().Equal(tokens["accessToken"].(string), "access_t")
// }
// func (s *SmokeSuite) TestAuthWrongUsername(t provider.T) {

// 	r := getToken(s.client, "valid", "valid_password")
// 	tokens := responseBodyToMap(r.Body())

// 	t.Assert().Equal(401, r.StatusCode())
// 	t.Assert().False(tokens["success"].(bool))
// 	t.Assert().Empty(tokens["refreshToken"].(string))
// 	t.Assert().Empty(tokens["accessToken"].(string))
// }
// func (s *SmokeSuite) TestAuthWrongPassword(t provider.T) {

// 	r := getToken(s.client, "valid_user", "valid")
// 	tokens := responseBodyToMap(r.Body())

// 	t.Assert().Equal(401, r.StatusCode())
// 	t.Assert().False(tokens["success"].(bool))
// 	t.Assert().Empty(tokens["refreshToken"].(string))
// 	t.Assert().Empty(tokens["accessToken"].(string))
// }

// func TestSmoke(t *testing.T) {
// 	suite.RunSuite(t, new(SmokeSuite))
// }
