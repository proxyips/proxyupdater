package proxy

import (
	"time"
	"github.com/go-resty/resty/v2"
)
const UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
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
