package config

import (
	"log"
	"os"
)

type Config struct {
	ServerAddr   string
	TemporalHost string
	TaskQueue    string
}

func Load() *Config {
	conf := &Config{
		ServerAddr:   getEnv("SERVER_ADDR", ":8765"),
		TemporalHost: getEnv("TEMPORAL_ADDRESS", "localhost:7233"),
		TaskQueue:    getEnv("TEMPORAL_TASK_QUEUE", "default-task-queue"),
	}
	log.Printf("Config loaded: %+v", conf)
	return conf
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
