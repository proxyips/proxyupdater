package echoserver

import (
	"gitlab.com/labstack/echo"
)

var e = echo.New()

func init() {
	e.HideBanner = false
	e.Debug = true
	e.Logger.Info()
}
