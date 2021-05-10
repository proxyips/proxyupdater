package proxy

import (
	"testing"
	"fmt"
	"encoding/json"
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
