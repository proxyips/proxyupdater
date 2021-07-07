package proxy

import (
	"testing"
	"fmt"
)

func TestBandwidthDump(t *testing.T) {
	dat, rst , err := bandwidthDump("https://prx.gcx.pw/")
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	for key, value := range dat.Servers {
		fmt.Println(key)
		fmt.Println(value)
	}
	fmt.Println(rst)
}
func TestBandwidthMeter(t *testing.T) {
	var x = BandwidthMeter{
		Host: "https://prx.gcx.pw",
		Path:"data",
	}
	err := x.Get()
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
}
func TestBandwidthMeterEcho(t *testing.T) {
	var x = BandwidthMeter{
		Host: "https://prx.gcx.pw",
		Path:"data",
	}
	dat, err := x.Echo()
	if err != nil {
		t.Name()
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(string(dat))
}
