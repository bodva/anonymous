package httpjson

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const baseUrl = "https://slack.com/api/"

type Param struct {
	Key   string
	Value string
}

type Response struct {
	Status bool   `json:"ok"`
	Error  string `json:"error"`
}

func PostJson(method string, params []Param, target interface{}) error {
	form := url.Values{}
	for _, param := range params {
		form.Add(param.Key, param.Value)
	}

	urlStr := baseUrl + method
	res, err := http.Post(urlStr, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

func GetJson(method string, target interface{}) error {
	urlStr := baseUrl + method
	res, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
