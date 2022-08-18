package src

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

func ResponseBodyToMap(r []byte) map[string]interface{} {
	var resp map[string]interface{}
	json.Unmarshal(r, &resp)
	return resp
}

func GetToken(c *resty.Client, uname string, pass string) *resty.Response {

	body := make(map[string]string)
	if uname != "" {
		body["username"] = uname
	}
	if pass != "" {
		body["password"] = pass
	}

	req := c.R().SetBody(body)
	r, _ := req.Post("/token/get")
	return r
}

func ValidateToken(c *resty.Client, token string) *resty.Response {
	req := c.R().SetBody(map[string]string{
		"token": token,
	})
	r, _ := req.Post("/token/validate")
	return r
}

func Register(c *resty.Client, body RegStruct) *resty.Response {
	req := c.R().SetBody(body)
	r, _ := req.Post("/register")
	return r
}

type RegStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
