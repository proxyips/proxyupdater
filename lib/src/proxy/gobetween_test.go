package proxy

import (
	"testing"
	"fmt"
)

func TestGetApiSettings(t *testing.T) {
	dat, err := GetApiSettings("data/dump.json")
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(dat)

}
