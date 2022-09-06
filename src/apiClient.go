package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
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

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Conf string `json:"conference"`
	Div  string `json:"division"`
	Year int    `json:"est_year"`
}

type ApiClient struct {
	r        *resty.Client
	pt       *provider.T
	Response *resty.Response
}

func (a *ApiClient) GetTeams(filters map[string]string) ([]Team, error) {
	u := url.URL{Path: "teams"}
	q := u.Query()
	for k, v := range filters {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	p := *a.pt
	finalUrl := u.JoinPath().String()
	fmt.Println(finalUrl)
	p.WithNewStep("Send request to 'teams' endpoint", func(sCtx provider.StepCtx) {}, allure.NewParameter("path and query", finalUrl))
	_resp, _ := a.r.R().Get(finalUrl)
	a.Response = _resp

	if _resp.StatusCode() != 200 {
		var m map[string]string
		json.Unmarshal(_resp.Body(), &m)

		return nil, errors.New(m["error"])
	}

	var teams []Team
	json.Unmarshal(_resp.Body(), &teams)

	return teams, nil
}

func NewApiClient(pt *provider.T, r *resty.Client) ApiClient {
	return ApiClient{r, pt, nil}
}
