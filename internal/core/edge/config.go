package edge

type CaddyConfig struct {
	// Admin   AdminConfig   `json:"admin"`
	Apps AppsConfig `json:"apps"`
	// Logging LoggingConfig `json:"logging"`
	// Storage StorageConfig `json:"storage"`
}

type AppsConfig struct {
	HTTP HTTPConfig `json:"http,omitempty"`
	// PKI  PKIConfig  `json:"pki"`
	// TLS  TLSConfig  `json:"tls"`
}

type HTTPConfig struct {
	Servers map[string]ServerConfig
}

type ServerConfig struct {
	Routes []RoutesConfig `json:"routes,omitempty"`
}

type RoutesConfig struct {
	Match  []MatchConfig `json:"match,omitempty"`
	Handle []struct {
		Handler   string `json:"handler,omitempty"`
		Upstreams []struct {
			Dial string `json:"dial,omitempty"`
		} `json:"upstreams"`
	} `json:"handle"`
}

type MatchConfig struct {
	Host string `json:"host,omitempty"`
}
