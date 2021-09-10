package proxy

import (
	"testing"
	"fmt"
)

func TestGetApiSettings(t *testing.T) {
	dat, err := GetApiSettings("https://Ga4k-jvbn-qZmW.proxyips.cloud/")
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(dat.API)

}

