package src

import (
	"encoding/json"
	"fmt"
	"net/url"

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

func GetTeams(c *resty.Client, filters map[string]string) *resty.Response {
	rr := url.URL{Path: "teams"}
	q := rr.Query()

	if len(filters) > 0 {
		for k, v := range filters {
			q.Set(k, v)
		}
	}
	rr.RawQuery = q.Encode()
	r, err := c.R().Get(rr.JoinPath().String())
	if err != nil {
		fmt.Println(err)
	}
	return r
}

type RegStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Conf string `json:"conference"`
	Div  string `json:"division"`
	Year int    `json:"est_year"`
}

// type ApiClient struct {
// 	r        *resty.Client
// 	pt       *provider.T
// 	Response *resty.Response
// }

// func (a *ApiClient) GetTeams(filters map[string]string) []Team {
// 	u := url.URL{Path: "teams"}
// 	q := u.Query()
// 	for k, v := range filters {
// 		q.Set(k, v)
// 	}

// 	u.RawQuery = q.Encode()
// 	a.pt.WithNewParameters(" Request to 'teams'",
// 		func(p provider.StpCtx) {}, allure.Parameter("path", filters))
// 	_resp := r.R().Get(u.Parsed())
// 	a.Response = _resp
// 	var teams []Team
// 	json.Unmarshall(_resp.Body(), &teams)
// 	return teams
// }
