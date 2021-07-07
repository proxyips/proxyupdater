package scripts

import "github.com/proxyips/proxyupdater/lib/src/echoserver"

func StartApiService(ipaddress, port string) (err error) {
	var x = echoserver.TLMSMSServer{
		IPAddress: ipaddress,
		Port:      port,
	}
	x.TLMServer()
	return
}

