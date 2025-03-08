package main

import (
	infraestructure "PubNotification/src/notification/infrastructure"
	"PubNotification/src/notification/infrastructure/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	dependencies := infraestructure.InitAsignature()
	defer dependencies.RabbitMQAdapter.Close()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	routes.ConfigureRoutesAsignature(r, dependencies.CreateAsignatureController)

	if err := r.Run(":8088"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
