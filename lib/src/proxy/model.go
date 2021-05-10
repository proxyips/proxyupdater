package proxy

import "github.com/proxyips/proxyupdater/lib/src/rest"

type PublicProxies []PublicProxy
type PublicProxy struct {
	Alive        int64    `json:"Alive"`
	Checked      int64    `json:"Checked"`
	Country      string   `json:"Country"`
	FirstSeen    int64    `json:"FirstSeen"`
	IP           string   `json:"Ip"`
	Port         int64    `json:"Port"`
	PrivacyLevel string   `json:"PrivacyLevel"`
	Reliability  float64  `json:"Reliability"`
	Speed        int64    `json:"Speed"`
	Tags         []string `json:"Tags"`
	TimesAlive   int64    `json:"TimesAlive"`
	TimesDead    int64    `json:"TimesDead"`
	Type         string   `json:"Type"`
}

type ApiSettings rest.GobetweenSettings
