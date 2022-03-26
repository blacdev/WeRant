package config

import (
	"fmt"
	"github.com/blacdev/werant/env"
	"os"
)

func GetServerAddress() string {
	port := os.Getenv(env.AppPort)
	if port == "" {
		port = "8080"
	}
	return fmt.Sprintf(":%s", port)
}
