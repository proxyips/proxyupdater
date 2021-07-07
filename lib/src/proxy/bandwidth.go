package proxy

import (
	"github.com/proxyips/proxyupdater/lib/src/rest"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"time"
)

func bandwidthDump(hostname string)  (gobetween GobetweenDump, mg *rest.Manager, err error) {
	var username, password string
	if hostname == ""{
		hostname = "http://127.0.0.1:8010/"
	} else {
		username = "9m3quhdw7aozps4Ekiyetbxncrvjg"
		password = "L4e7LTBKvr7vnXTnJjBJ"
	}
	mg = &rest.Manager{
		Uri: fmt.Sprintf("%vdump?format=json", hostname),
		UserName: username,
		Password: password,
	}

	dat, err := mg.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dat))
	json.Unmarshal(dat, &gobetween)
 return
}

type BandwidthMeter struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

func (b BandwidthMeter) Get() (err error) {
	dat, rst, err := bandwidthDump(b.Host)
	if err != nil {
		fmt.Println(err)
	}
	var ingressPoints = IngressPoints{
		Time: time.Now(),
	}

	for key, value := range dat.Servers {
		var ingressStats = IngressStats{
			Name: key,
			Server: value,
		}

		rst.Uri = fmt.Sprintf("%vservers/%v/stats",b.Host, key )
		d, err := rst.Get()
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(d, &ingressStats.Point)
		ingressPoints.Ingress = append(ingressPoints.Ingress, ingressStats)

	}
	datIngressPoints, err := json.Marshal(ingressPoints)
	ioutil.WriteFile(fmt.Sprintf("%v/bandwidth_%v.json",b.Path,ingressPoints.Time.Unix()), datIngressPoints, 0644)
	ioutil.WriteFile(fmt.Sprintf("%v/bandwidth.json",b.Path), datIngressPoints, 0644)
	return
}
func (b BandwidthMeter) Echo() (data []byte,  err error) {
	dat, rst, err := bandwidthDump(b.Host)
	if err != nil {
		fmt.Println(err)
	}
	var ingressPoints = IngressPoints{
		Time: time.Now(),
	}

	for key, value := range dat.Servers {
		var ingressStats = IngressStats{
			Name: key,
			Server: value,
		}

		rst.Uri = fmt.Sprintf("%vservers/%v/stats",b.Host, key )
		d, err := rst.Get()
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(d, &ingressStats.Point)
		ingressPoints.Ingress = append(ingressPoints.Ingress, ingressStats)

	}
	datIngressPoints, err := json.Marshal(ingressPoints)
	ioutil.WriteFile(fmt.Sprintf("%v/bandwidth_%v.json",b.Path,ingressPoints.Time.Unix()), datIngressPoints, 0644)
	ioutil.WriteFile(fmt.Sprintf("%v/bandwidth.json",b.Path), datIngressPoints, 0644)
	return datIngressPoints, err
}
