package echoserver

import (
	"fmt"
	"net"
	"os"
	"github.com/labstack/echo/v4"
	"github.com/proxyips/proxyupdater/lib/src/proxy"
)

func check(ip string) bool {
	trial := net.ParseIP(ip)
	if os.Getenv("nmpw3xEcjogyzbfA4uqr") == "zj9shEHitCdFwvmarygo" {
		fmt.Println(trial)
		fmt.Println(trial.To4())
	}

	if trial.To4() == nil {
		//fmt.Printf("%v is not an IPv4 address\n", trial)
		return false
	}
	return true
}

type TLMSMSServer struct {
	Port      string
	IPAddress string
}

func (t TLMSMSServer) TLMServer() {

	e.GET("/logs/", logHandler)
	if t.IPAddress != "" {
		x := check(t.IPAddress)
		if x {
			startStr := fmt.Sprintf("%v:%v", t.IPAddress, t.Port)
			e.Start(startStr)

		} else {
			startStr := fmt.Sprintf("%v:%v", "0.0.0.0", t.Port)
			e.Start(startStr)
		}
	} else {
		startStr := fmt.Sprintf("%v:%v", "localhost", t.Port)
		e.Start(startStr)
	}

}

func logHandler(c echo.Context)  (err error) {
	var x = proxy.BandwidthMeter{
		Host: "http://127.0.0.1:8010",
		Path: "/opt/gobetween/data",
	}
	data, err := x.Echo()
	if err != nil {
		return c.JSONBlob(301, []byte(`{"status": "failed"}`))

	}
return c.JSONBlob(200, data)
}