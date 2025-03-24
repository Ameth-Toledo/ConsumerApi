package controllers

import (
	"PubNotification/src/notification/application"
	"PubNotification/src/notification/domain"
	"PubNotification/src/notification/domain/entities"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CreateAsignatureController struct {
	useCase    *application.CreateAsignature
	asignature domain.INotification
}

func NewCreateAsignatureController(useCase *application.CreateAsignature, asignature domain.INotification) *CreateAsignatureController {
	return &CreateAsignatureController{useCase: useCase, asignature: asignature}
}

func (cs_a *CreateAsignatureController) Execute(c *gin.Context) {
	var asignature entities.Notification
	if err := c.ShouldBindJSON(&asignature); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	err := cs_a.useCase.Execute(asignature)
	if err != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear la asignatura"})
		return
	}

	go sendNotificationToNodeServer(asignature.Message)

	c.JSON(201, gin.H{"message": "Asignatura creada correctamente", "asignature": asignature})
}

func sendNotificationToNodeServer(message string) {
	url := "http://3.232.89.200:4004/send-notification"

	payload := map[string]string{
		"message": message,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error al convertir el mensaje a JSON: %v", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Error enviando solicitud POST al servidor Node.js: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error en la respuesta del servidor Node.js: %v", resp.Status)
		return
	}

	log.Println("Notificación enviada a través de WebSocket al servidor Node.js.")
}
