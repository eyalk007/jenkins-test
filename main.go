package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port    int    `yaml:"port"`
	Message string `yaml:"message"`
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting Jenkins Test Application")

	config := Config{
		Port:    8080,
		Message: "Hello from Jenkins Test!",
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("Error marshaling config: %v", err)
	}
	fmt.Printf("Config:\n%s\n", string(data))

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": config.Message,
			"status":  "healthy",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	addr := fmt.Sprintf(":%d", config.Port)
	log.Infof("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

