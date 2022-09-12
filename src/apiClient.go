package src

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

const BASE_URL = "http://localhost:8080"

type Bodies interface {
	GetTokenRequest | RegStruct | RegisterResponse | Team
}

func getJson[V Bodies](body V) []byte {
	b, _ := json.Marshal(body)
	return b
}

type GetTokenRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type GetTokenResponse struct {
	Token      string `json:"token,omitempty"`
	Error      string `json:"error,omitempty"`
	StatusCode int
}

type RegStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterResponse struct {
	StatusCode int
	Message    string `json:"message,omitempty"`
	UserId     int    `json:"userId,omitempty"`
	Error      string `json:"error,omitempty"`
}

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Conf string `json:"conference"`
	Div  string `json:"division"`
	Year int    `json:"est_year"`
}

type TeamsResponse struct {
	Teams      []Team
	Error      string `json:"error,omitempty"`
	StatusCode int
}

// func ResponseBodyToMap(r []byte) map[string]interface{} {
// 	var resp map[string]interface{}
// 	json.Unmarshal(r, &resp)
// 	return resp
// }

// func GetToken(c *resty.Client, uname string, pass string) *resty.Response {

// 	body := make(map[string]string)
// 	if uname != "" {
// 		body["username"] = uname
// 	}
// 	if pass != "" {
// 		body["password"] = pass
// 	}

// 	req := c.R().SetBody(body)
// 	r, _ := req.Post("/token/get")
// 	return r
// }

func ValidateToken(c *resty.Client, token string) *resty.Response {
	req := c.R().SetBody(map[string]string{
		"token": token,
	})
	r, _ := req.Post("/token/validate")
	return r
}

type ApiClient struct {
	_url     *url.URL
	pt       *provider.T
	Response *http.Response
}

func (a *ApiClient) GetToken(gtb GetTokenRequest) GetTokenResponse {
	a._url.Path = "token/get"
	finalUrl := a._url.String()

	p := *a.pt
	p.WithNewStep("Send request to 'token/get'", func(sCtx provider.StepCtx) {}, allure.NewParameter("body", gtb))

	resp, body, err := post(finalUrl, getJson(gtb))
	if err != nil {
		return GetTokenResponse{Error: err.Error()}
	}
	a.Response = resp

	var gtr GetTokenResponse
	json.Unmarshal(body, &gtr)
	gtr.StatusCode = resp.StatusCode

	return gtr
}

func (a *ApiClient) GetTeams(filters map[string]string) TeamsResponse {
	u := a._url
	u.Path = "teams"
	q := u.Query()
	for k, v := range filters {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	p := *a.pt
	finalUrl := u.String()
	p.WithNewStep("Send request to '"+finalUrl+"'", func(sCtx provider.StepCtx) {})
	resp, body := get(finalUrl)
	a.Response = resp

	var tr TeamsResponse
	if resp.StatusCode != 200 {
		json.Unmarshal(body, &tr)
	} else {
		json.Unmarshal(body, &tr.Teams)
	}
	tr.StatusCode = resp.StatusCode

	return tr
}

func (a *ApiClient) Register(body RegStruct) RegisterResponse {

	prov := *a.pt
	prov.WithNewStep("Send request to 'register' endpoint", func(sCtx provider.StepCtx) {}, allure.NewParameter("body", body))

	a._url.Path = "register"
	resp, b, err := post(a._url.String(), getJson(body))
	if err != nil {
		return RegisterResponse{Error: err.Error()}
	}

	var r RegisterResponse
	json.Unmarshal(b, &r)
	r.StatusCode = resp.StatusCode
	return r
}

func NewApiClient(pt *provider.T) ApiClient {
	_url, _ := url.Parse(BASE_URL)

	return ApiClient{_url, pt, nil}
}

func get(url string) (*http.Response, []byte) {
	resp, _ := http.Get(url)
	b, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return resp, b
}

func post(url string, body []byte) (*http.Response, []byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}

	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, nil, err
	}

	return resp, b, nil
}
