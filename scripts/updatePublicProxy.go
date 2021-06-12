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
	var baseDefault Base
	json.Unmarshal([]byte(betweenbase), &baseDefault)
	baseDefault.Bind = newProxy.Bind
	baseDefault.Discovery = newProxy.Discovery
	baseDefault.Access = newProxy.Access
	baseDefault.Discovery.StaticList = proxiesDat
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

var betweenbase = `
{
    "max_connections": 50,
    "client_idle_timeout": "130s",
    "backend_idle_timeout": "0",
    "backend_connection_timeout": "0",
    "bind": "45.56.96.249:8091",
    "protocol": "tcp",
    "balance": "roundrobin",
    "access": {
        "default": "deny",
        "rules": [
            "allow 24.34.73.8",
            "allow 78.138.0.108"
        ]
    },
    "discovery": {
        "kind": "static",
        "failpolicy": "setempty",
        "interval": "5s",
        "timeout": "2s",
        "static_list": []
    },
    "healthcheck": {
        "kind": "exec",
        "interval": "11m",
        "passes": 1,
        "fails": 1,
        "timeout": "2s",
        "initial_status": "healthy",
        "exec_command": "/opt/bin/proxytester.sh",
        "exec_expected_positive_output": "1",
        "exec_expected_negative_output": ""
    }
}`

type Base struct {
	Access struct {
		Default string   `json:"default"`
		Rules   []string `json:"rules"`
	} `json:"access"`
	BackendConnectionTimeout string `json:"backend_connection_timeout"`
	BackendIdleTimeout       string `json:"backend_idle_timeout"`
	Balance                  string `json:"balance"`
	Bind                     string `json:"bind"`
	ClientIdleTimeout        string `json:"client_idle_timeout"`
	Discovery                struct {
		Failpolicy string        `json:"failpolicy"`
		Interval   string        `json:"interval"`
		Kind       string        `json:"kind"`
		StaticList []string `json:"static_list"`
		Timeout    string        `json:"timeout"`
	} `json:"discovery"`
	Healthcheck struct {
		ExecCommand                string `json:"exec_command"`
		ExecExpectedNegativeOutput string `json:"exec_expected_negative_output"`
		ExecExpectedPositiveOutput string `json:"exec_expected_positive_output"`
		Fails                      int64  `json:"fails"`
		InitialStatus              string `json:"initial_status"`
		Interval                   string `json:"interval"`
		Kind                       string `json:"kind"`
		Passes                     int64  `json:"passes"`
		Timeout                    string `json:"timeout"`
	} `json:"healthcheck"`
	MaxConnections int64  `json:"max_connections"`
	Protocol       string `json:"protocol"`
}