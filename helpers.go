package main

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

func responseBodyToMap(r []byte) map[string]interface{} {
	var resp map[string]interface{}
	json.Unmarshal(r, &resp)
	return resp
}

func getToken(c *resty.Client, uname string, pass string) *resty.Response {

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

// func checkToken(c *resty.Client, token string) *resty.Response {
// 	req := c.R().SetBody(map[string]string{
// 		"accessToken": token,
// 	})
// 	r, _ := req.Post("/token/check")
// 	return r
// }
