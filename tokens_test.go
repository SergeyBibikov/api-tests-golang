package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TokenSuite struct {
	BaseSuite
}

func (t *TokenSuite) TestCheckValidToken() {
	r := checkToken(t.client, "access_t")
	resp := responseBodyToMap(r.Body())

	t.Equal(200, r.StatusCode())
	t.True(resp["valid"].(bool))
}
func (t *TokenSuite) TestCheckInvalidToken() {
	r := checkToken(t.client, "access_tok")
	resp := responseBodyToMap(r.Body())

	t.Equal(400, r.StatusCode())
	t.False(resp["valid"].(bool))
}

func TestTokens(t *testing.T) {
	suite.Run(t, new(TokenSuite))
}
