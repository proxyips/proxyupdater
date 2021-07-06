package proxy

import (
	"encoding/json"
	"github.com/proxyips/proxyupdater/lib/src/rest"
	"fmt"
	"io/ioutil"
)

type ApiAccess struct {
	BasicAuth struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	} `json:"basic_auth"`
	Bind    string      `json:"bind"`
}

func GetApiSettings(host string)  (gobetween GobetweenDump, err error) {
	var mg = rest.Manager{
		Uri: fmt.Sprintf("%vdump?format=json", host),
		UserName: "9m3quhdw7aozps4Ekiyetbxncrvjg",
		Password: "L4e7LTBKvr7vnXTnJjBJ",
	}
	dat, err := mg.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dat))
	json.Unmarshal(dat, &gobetween)
 return
}
func GetApiFileSettings(filename string)  (api ApiAccess, err error) {
	dat, err := ioutil.ReadFile(filename)
	var gobetween rest.Gobetween
	json.Unmarshal(dat, &gobetween)
	api.BasicAuth = gobetween.API.BasicAuth
	api.Bind = gobetween.API.Bind
 return
}

