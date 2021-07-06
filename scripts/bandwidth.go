package scripts

import "github.com/proxyips/proxyupdater/lib/src/proxy"

func BandwidthLog(host, path string)  (err error) {
	var x = proxy.BandwidthMeter{
		Host: host,
		Path: path,
	}
	return x.Get()
} 