package proxy

import (
	"testing"
	"fmt"
)

func TestExtractProxyByCountry(t *testing.T) {
	//d, err := ioutil.ReadFile("data/proxylist.json")
	d, err := GetProxyUpdate()
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	dat, err := ExtractProxyByCountry(d, "us")
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(dat)
}
func TestExtractProxy(t *testing.T) {
	//d, err := ioutil.ReadFile("data/proxylist.json")
	d, err := GetProxyUpdate()
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	dat, err := ExtractProxy(d)
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(dat)
}
