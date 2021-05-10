package rest

import (
	"time"
	"github.com/go-resty/resty/v2"
)

const UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
type Manager struct {
	Payload []byte `json:"payload"`
	Uri string `json:"uri"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (m Manager) Post()  (rsp []byte, err error) {
	client := resty.New()
	if m.UserName != ""{
		client.SetBasicAuth(m.UserName, m.Password)
	}
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent": UserAgent,
	})
	client.SetTimeout(3 * time.Minute)
	resp, err := client.R().
		EnableTrace().
		SetBody(m.Payload).
		Post(m.Uri)

	return resp.Body(), nil
}
func (m Manager) Get()  (rsp []byte, err error) {
	client := resty.New()
	if m.UserName != ""{
		client.SetBasicAuth(m.UserName, m.Password)
	}
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent": UserAgent,
	})
	client.SetTimeout(2 * time.Second)
	resp, err := client.R().
		EnableTrace().
		Get(m.Uri)

	return resp.Body(), nil
}
func (m Manager) Delete()  (rsp []byte, err error) {
	client := resty.New()
	if m.UserName != ""{
		client.SetBasicAuth(m.UserName, m.Password)
	}
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent": UserAgent,
	})
	client.SetTimeout(8 * time.Second)
	resp, err := client.R().
		EnableTrace().
		Delete(m.Uri)

	return resp.Body(), nil
}

