package tests

import (
	"encoding/json"
	beetest "github.com/astaxie/beego/testing"
	"io/ioutil"
	"testing"
)

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

func TestShort(t *testing.T) {
	request := beetest.Post("/shorten")
	request.Param("longurl", "http://www.beego.me/")

	response, _ := request.Response()

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	var s ShortResult
	json.Unmarshal(content, &s)
	if s.UrlShort == "" {
		t.Fatal("shorturl is empty")
	}
}
