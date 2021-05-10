package scripts

import (
	"testing"
	"fmt"
)

func TestUpdateProxyList(t *testing.T) {
	dat, err := updateProxyList("us")
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(dat)
}
func TestUpdateGobetween(t *testing.T) {
	err := UpdateGobetween("data/dump.json", "")
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
}
