package main

import (
	"github.com/go-resty/resty/v2"
)

func main() {

}

func GetClient() *resty.Client {
	return resty.New().SetBaseURL("http://localhost:8080")
}
