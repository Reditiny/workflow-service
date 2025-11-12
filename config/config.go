package config

import (
	"log"
	"os"
)

type Config struct {
	ServerAddr   string
	TemporalHost string
	Namespace    string
	APIKey       string
	TaskQueue    string
}

func Load() *Config {
	conf := &Config{
		ServerAddr:   getEnv("SERVER_PORT", ":8765"),
		TemporalHost: getEnv("TEMPORAL_ADDRESS", "localhost:7233"),
		Namespace:    getEnv("TEMPORAL_NAMESPACE", "default"),
		APIKey:       getEnv("TEMPORAL_API_KEY", ""),
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
