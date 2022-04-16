package qlite

type Config struct {
	Addr		string	`json:"addr"`
	Password	string	`json:"password,omitempty"`
	Cap			int		`json:"cap,omitempty"`
}
