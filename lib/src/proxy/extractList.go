package proxy

import (
	"encoding/json"
	"strings"
	"fmt"
)

func ExtractProxyByCountry(data []byte, country string)  (proxies []string, err error) {
	var rawProxyList PublicProxies
	json.Unmarshal(data, &rawProxyList)
	//fmt.Println(rawProxyList)
	country = strings.ToUpper(country)
	for _, v := range rawProxyList {
		if v.Country == country{
			//fmt.Println(v)
			proxy := fmt.Sprintf(`%v:%v`, v.IP, v.Port)
			proxies = append(proxies, proxy)
		}
	}
 return
}
func ExtractProxy(data []byte)  (proxies []string, err error) {
	var rawProxyList PublicProxies
	json.Unmarshal(data, &rawProxyList)
	//fmt.Println(rawProxyList)
	for _, v := range rawProxyList {
		//fmt.Println(v)
		proxy := fmt.Sprintf(`%v:%v`, v.IP, v.Port)
		proxies = append(proxies, proxy)
	}
 return
}
