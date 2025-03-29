package config

import "os"

type Config struct {
	ServerPort string
	JWTSecret  string
}

func LoadConfig() *Config {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "my-very-secret-jwt-key"
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "3000"
	}

	return &Config{
		JWTSecret:  jwtSecret,
		ServerPort: serverPort,
	}
}
