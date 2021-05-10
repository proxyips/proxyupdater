package scripts

import (
	"fmt"
	"github.com/proxyips/proxyupdater/lib/src/proxy"
	"encoding/json"
	"github.com/proxyips/proxyupdater/lib/src/rest"
	"regexp"
)

func updateProxyList(country string)  (proxies []string, err error) {
	d, err := proxy.GetProxyUpdate()
	if err != nil {
		fmt.Println("updateProxyList - GetProxyUpdate")
		fmt.Println(err)
		return
	}
	var proxyList []string
	if country == ""{
		proxies, err := proxy.ExtractProxy(d)
		if err != nil {

		}
		proxyList = proxies
	} else {
		proxies, err := proxy.ExtractProxyByCountry(d, country)
		if err != nil {

		}
		proxyList = proxies
	}
	if err != nil {
		fmt.Println("updateProxyList - ExtractProxy")
		fmt.Println(err)
		return
	}

 return proxyList, err
}

func UpdateGobetween(filename string, country string)  ( err error) {
	//TODO get info from gobetween.json
	var host string
	apiDat, err := proxy.GetApiSettings(filename)
	domain := regexp.MustCompile(`[aA-zZ]`)
	if domain.MatchString(apiDat.Bind){
		host = fmt.Sprintf("https://%v/", apiDat.Bind)
	} else {
		host = fmt.Sprintf("http://%v/", apiDat.Bind)
	}
	//fmt.Println(host)
	var mg = rest.Manager{
		Uri: fmt.Sprintf("%vdump?format=json", host),
		UserName: apiDat.BasicAuth.Login,
		Password: apiDat.BasicAuth.Password,
	}
	dat, err := mg.Get()
	if err != nil {
	return
	}
	var data rest.Gobetween

	json.Unmarshal(dat, &data)
	sDat, err := json.Marshal(data.Servers)
	if err != nil {
	return
	}
	var servers = make(map[string]rest.Server)
	json.Unmarshal(sDat, &servers)
	//fmt.Println(servers["public_proxies"].Discovery.StaticList)
	proxiesDat, err := updateProxyList(country)
	if err != nil {
		return
	}
	mg.Uri = fmt.Sprintf("%vservers/public_proxies", host)
	fmt.Println(mg.Uri)
	var newProxy rest.Server
	newProxy = servers["public_proxies"]
	newProxy.Discovery.StaticList = proxiesDat
	pDat, err := json.Marshal(newProxy)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(pDat))
	mg.Delete()
	mg.Payload = pDat
	_, err = mg.Post()
	if err != nil {
		return
	}
	//fmt.Println(string(prsp))

	return
}
var Countries = proxy.CountryList