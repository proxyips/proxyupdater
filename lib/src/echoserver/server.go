package echoserver

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/middleware"
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

	// static files
	//path := "/Source/golang/src/sc.tpnfc.us/oyoshi/tlmtext"
	path, _ := os.Getwd()
	path = filepath.Join(path, "frontend", "deploy")
	path = strings.TrimPrefix(path, "C:")
	path = strings.Replace(path, "\\", "/", 100)
	path = strings.Replace(path, "lib/src/echoserver/", "", 1)
	if os.Getenv("nmpw3xEcjogyzbfA4uqr") == "zj9shEHitCdFwvmarygo" {
		fmt.Println(path)

	}
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   path,
		Browse: true,
		HTML5:  true,
		Index:  "index.html",
	}))
	//e.Static("/main.dart.js", path + "app/published/web/main.dart.js")
	//e.Static("/static/", path + "app/published/web/static")
	//e.Static("/packages/", path + "app/published/web/packages")

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
