package qlite

type Config struct {
	Addr		string	`json:"addr"`
	Password	string	`json:"password,omitempty"`
	Cap			int		`json:"cap,omitempty"`
}

func DefaultConfig() *Config {
	return &Config{
		Addr: "127.0.0.1:9810",
		Password: "admin",
		Cap: 1024,
	}
}
