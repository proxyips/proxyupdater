package echoserver

import (
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func init() {
	e.HideBanner = false
	e.Debug = true
	e.Logger.Info()
}
