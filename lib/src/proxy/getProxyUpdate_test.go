package proxy

import (
	"testing"
	"fmt"
	"encoding/json"
	"github.com/Jeffail/gabs/v2"
)


func TestGetProxyUpdate(t *testing.T) {
	dat, err := GetProxyUpdate()
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	//fmt.Println(string(dat))
	var rawProxyList PublicProxies
	json.Unmarshal(dat, &rawProxyList)
	fmt.Println(rawProxyList)
}
func TestPublicProxyJson(t *testing.T) {
	dat := PublicProxyJson()
	fmt.Println(gabs.Wrap(dat).String())
}
func TestPublicProxyUsJson(t *testing.T) {
	dat := PublicProxyUsJson("us")
	fmt.Println(gabs.Wrap(dat).String())
}
