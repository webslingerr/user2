package config

type Config struct {
	UserFileName string
}

func Load() Config {
	cfg := Config{}

	cfg.UserFileName = "./data/user.json"

	return cfg
}
