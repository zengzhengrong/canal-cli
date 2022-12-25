package config

import "os"

var Conf *Config

type CanalConfig struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	ListenReg   string `json:"listen_db"`
	Destination string `json:"destination"`
}

type Config struct {
	Canal *CanalConfig
}

func init() {
	Conf = NewConfig()
}

func NewConfig() *Config {
	canal := &CanalConfig{
		Host:        os.Getenv("CANAL_HOST"),
		Port:        os.Getenv("CANAL_HOST_PORT"),
		ListenReg:   os.Getenv("CANAL_LISTEN_REG"),
		Destination: os.Getenv("CANAL_DESTINATION"),
	}
	return &Config{
		Canal: canal,
	}
}
