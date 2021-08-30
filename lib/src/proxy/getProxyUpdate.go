package proxy

import (
	"time"
	"github.com/go-resty/resty/v2"
	"encoding/json"
	"fmt"
)
const UserAgent = "proxyUpdater/service 2020-08"
func GetProxyUpdate()  (rsp []byte, err error) {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent": UserAgent,
	})
	client.SetTimeout(2 * time.Second)
	Uri := "https://raw.githubusercontent.com/proxyips/proxylist/main/proxylist.json"
	resp, err := client.R().
		EnableTrace().
		Get(Uri)

	return resp.Body(), nil
}
//PublicProxyJson JSON discovery is useful for custom setups when you have your own service registry implementation that can provide backends list in JSON format. gobetween will make HTTP query
func PublicProxyJson()  (proxies DiscoverProxies) {
	dat, err := GetProxyUpdate()
	if err != nil {
		return
	}
	var rawProxyList PublicProxies
	json.Unmarshal(dat, &rawProxyList)
	for _, v := range rawProxyList {
		if v.Alive == 1{
			if v.Reliability > 55 {
				var pprx  = DiscoverProxy{
					Host: v.IP,
					Port: fmt.Sprintf("%v",v.Port),
				}
				proxies = append(proxies, pprx)
			}
		}
	}
	return
}
