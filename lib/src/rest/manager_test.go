package rest

import (
	"testing"
	"fmt"
	"encoding/json"
)
var host = "http://192.53.121.65:8080/"
func TestManagerGet(t *testing.T) {
	var mg = Manager{
		Uri: fmt.Sprintf("%vdump?format=json", host),
		UserName: "9m3quhdw7aozps4Ekiyetbxncrvjg",
		Password: "L4e7LTBKvr7vnXTnJjBJ",
	 }
 	dat, err := mg.Get()
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	var data Gobetween

	json.Unmarshal(dat, &data)
 	sDat, err := json.Marshal(data.Servers)
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	var servers = make(map[string]Server)
	json.Unmarshal(sDat, &servers)
	fmt.Println(servers["public_proxies"].Discovery.StaticList)

	mg.Uri = fmt.Sprintf("%v/servers/public_proxies", host)
	pDat, err := json.Marshal(servers["public_proxies"])
	mg.Payload = pDat
	mg.Delete()
	prsp, err := mg.Post()
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(string(prsp))
}
