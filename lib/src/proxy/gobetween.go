package proxy

import (
	"io/ioutil"
	"encoding/json"
	"github.com/proxyips/proxyupdater/lib/src/rest"
)

type ApiAccess struct {
	BasicAuth struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	} `json:"basic_auth"`
	Bind    string      `json:"bind"`
}

func GetApiSettings(file string)  (api ApiAccess, err error) {
	dat, err := ioutil.ReadFile(file)
	var gobetween rest.Gobetween
	json.Unmarshal(dat, &gobetween)
	api.BasicAuth = gobetween.API.BasicAuth
	api.Bind = gobetween.API.Bind
 return
}