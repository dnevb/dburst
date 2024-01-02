package provider

import (
	"flag"
	"os"
)

type Config struct {
	Host   string
	Path   string
	Secret []byte
}

func CreateConfig() *Config {
	var host, path, secret string

	flag.StringVar(&host, "host", os.Getenv("DB_HOST"), "")
	flag.StringVar(&path, "path", os.Getenv("DB_PATH"), "")
	flag.StringVar(&secret, "secret", os.Getenv("DB_SECRET"), "")

	flag.Parse()

	return &Config{
		Host:   host,
		Path:   path,
		Secret: []byte(secret),
	}
}
