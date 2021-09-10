package scripts

import (
	"github.com/Jeffail/gabs/v2"
	"fmt"
	"github.com/proxyips/proxyupdater/lib/src/proxy"
)

func PublicProxies()  {
	dat := proxy.PublicProxyJson()
	fmt.Println(gabs.Wrap(dat).String())
 return
}
func PublicCountryProxies(country string)  {
	dat := proxy.PublicProxyUsJson(country)
	fmt.Println(gabs.Wrap(dat).StringIndent("", "    "))
 return
}
