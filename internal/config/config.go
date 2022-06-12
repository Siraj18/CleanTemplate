package config

type Config struct {
	Address string
}

//TODO сменить на yaml или переменные окружения
func NewConfig() *Config {
	return &Config{
		Address: ":8080",
	}
}
