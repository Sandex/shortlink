package config

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"127.0.0.1:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"/api/shorten"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"/tmp/store.dat"`
}
