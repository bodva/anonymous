package httpjson

import (
	"testing"
)

func TestGetJson(t *testing.T) {
	var method string

	r := Response{}

	method = "api.test"
	err := GetJson(method, &r)
	if err != nil {
		t.Error("Internal GetJson error")
	}
	if !r.Status {
		t.Errorf("Method %s response with error: %s", method, r.Error)
	}
}

func TestPostJson(t *testing.T) {
	var method string

	r := Response{}

	method = "api.test"
	params := []Param{}
	err := PostJson(method, params, &r)
	if err != nil {
		t.Error("Internal PostJson error")
	}
	if !r.Status {
		t.Errorf("Method %s response with error: %s", method, r.Error)
	}
}
