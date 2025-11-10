package config

import "os"

type Config struct {
	ServerAddr   string
	TemporalHost string
	TaskQueue    string
}

func Load() *Config {
	return &Config{
		ServerAddr:   getEnv("SERVER_ADDR", ":8765"),
		TemporalHost: getEnv("TEMPORAL_HOST", "localhost:7233"),
		TaskQueue:    getEnv("TEMPORAL_TASK_QUEUE", "default-task-queue"),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
