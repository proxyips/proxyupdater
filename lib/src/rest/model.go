package rest

type Server struct {
	Access struct {
		Default string   `json:"default"`
		Rules   []string `json:"rules"`
	} `json:"access"`
	BackendConnectionTimeout string      `json:"backend_connection_timeout"`
	BackendIdleTimeout       string      `json:"backend_idle_timeout"`
	BackendsTLS              interface{} `json:"backends_tls"`
	Balance                  string      `json:"balance"`
	Bind                     string      `json:"bind"`
	ClientIdleTimeout        string      `json:"client_idle_timeout"`
	Discovery                struct {
		Failpolicy string   `json:"failpolicy"`
		Interval   string   `json:"interval"`
		Kind       string   `json:"kind"`
		StaticList []string `json:"static_list"`
		Timeout    string   `json:"timeout"`
	} `json:"discovery"`
	Healthcheck struct {
		Fails         int64  `json:"fails"`
		InitialStatus string `json:"initial_status"`
		Interval      string `json:"interval"`
		Kind          string `json:"kind"`
		Passes        int64  `json:"passes"`
		Timeout       string `json:"timeout"`
	} `json:"healthcheck"`
	MaxConnections int64       `json:"max_connections"`
	Protocol       string      `json:"protocol"`
	ProxyProtocol  interface{} `json:"proxy_protocol"`
	Sni            interface{} `json:"sni"`
	TLS            interface{} `json:"tls"`
	Udp            interface{} `json:"udp"`
}

type GobetweenSettings struct {
	Acme interface{} `json:"acme"`
	API  struct {
		BasicAuth struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		} `json:"basic_auth"`
		Bind    string      `json:"bind"`
		Cors    bool        `json:"cors"`
		Enabled bool        `json:"enabled"`
		TLS     interface{} `json:"tls"`
	} `json:"api"`
	Defaults struct {
		BackendConnectionTimeout interface{} `json:"backend_connection_timeout"`
		BackendIdleTimeout       interface{} `json:"backend_idle_timeout"`
		ClientIdleTimeout        interface{} `json:"client_idle_timeout"`
		MaxConnections           interface{} `json:"max_connections"`
	} `json:"defaults"`
	Logging struct {
		Format string `json:"format"`
		Level  string `json:"level"`
		Output string `json:"output"`
	} `json:"logging"`
	Metrics struct {
		Bind    string `json:"bind"`
		Enabled bool   `json:"enabled"`
	} `json:"metrics"`
	Profiler interface{} `json:"profiler"`
	Servers []Server `json:"servers"`
}
type Gobetween struct {
	Acme interface{} `json:"acme"`
	API  struct {
		BasicAuth struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		} `json:"basic_auth"`
		Bind    string      `json:"bind"`
		Cors    bool        `json:"cors"`
		Enabled bool        `json:"enabled"`
		TLS     interface{} `json:"tls"`
	} `json:"api"`
	Defaults struct {
		BackendConnectionTimeout interface{} `json:"backend_connection_timeout"`
		BackendIdleTimeout       interface{} `json:"backend_idle_timeout"`
		ClientIdleTimeout        interface{} `json:"client_idle_timeout"`
		MaxConnections           interface{} `json:"max_connections"`
	} `json:"defaults"`
	Logging struct {
		Format string `json:"format"`
		Level  string `json:"level"`
		Output string `json:"output"`
	} `json:"logging"`
	Metrics struct {
		Bind    string `json:"bind"`
		Enabled bool   `json:"enabled"`
	} `json:"metrics"`
	Profiler interface{} `json:"profiler"`
	Servers interface{} `json:"servers"`
}
type Stats struct {
	ActiveConnections int64 `json:"active_connections"`
	Backends          []struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Priority int64  `json:"priority"`
		Stats    struct {
			ActiveConnections  int64 `json:"active_connections"`
			Discovered         bool  `json:"discovered"`
			Live               bool  `json:"live"`
			RefusedConnections int64 `json:"refused_connections"`
			Rx                 int64 `json:"rx"`
			RxSecond           int64 `json:"rx_second"`
			TotalConnections   int64 `json:"total_connections"`
			Tx                 int64 `json:"tx"`
			TxSecond           int64 `json:"tx_second"`
		} `json:"stats"`
		Weight int64 `json:"weight"`
	} `json:"backends"`
	RxSecond int64 `json:"rx_second"`
	RxTotal  int64 `json:"rx_total"`
	TxSecond int64 `json:"tx_second"`
	TxTotal  int64 `json:"tx_total"`
}