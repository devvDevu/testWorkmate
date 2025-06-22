package http_config

type HttpConfig struct {
	Addr     string `yaml:"addr"`
	UseHttps bool   `yaml:"use_https"`
}

func (h *HttpConfig) GetAddr() string {
	return h.Addr
}

func (h *HttpConfig) GetUseHttps() bool {
	return h.UseHttps
}
