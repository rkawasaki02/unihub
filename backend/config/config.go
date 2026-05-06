package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port    int
	NASRoot string
}

func Load() *Config {
	port := 8080
	if p := os.Getenv("UNIHUB_PORT"); p != "" {
		if v, err := strconv.Atoi(p); err == nil {
			port = v
		}
	}

	nasRoot := os.Getenv("UNIHUB_NAS_ROOT")
	if nasRoot == "" {
		nasRoot = os.Getenv("HOME") + "/nas"
	}

	return &Config{
		Port:    port,
		NASRoot: nasRoot,
	}
}
