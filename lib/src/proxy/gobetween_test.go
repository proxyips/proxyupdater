package proxy

import (
	"testing"
	"fmt"
)

func TestGetApiSettings(t *testing.T) {
	dat, err := GetApiSettings("http://prx.gcx.pw/")
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(dat.API)

}

