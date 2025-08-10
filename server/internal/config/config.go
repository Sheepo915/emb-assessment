package config

import (
	"flag"
	"os"
)

type Config struct {
	Addr     string
	DummyAPI string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ParseFlag() {
	flag.StringVar(&c.Addr, "addr", os.Getenv("ADDR"), "Address of the server")
	flag.StringVar(&c.DummyAPI, "dummy_api", os.Getenv("DUMMY_API"), "API to the dummy data")
	flag.Parse()
}
