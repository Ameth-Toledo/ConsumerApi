package infraestructure

import (
	"PubNotification/src/notification/application"
	"PubNotification/src/notification/application/repositories"
	"PubNotification/src/notification/infrastructure/adapter"
	"PubNotification/src/notification/infrastructure/controllers"
	"log"
)

type DependenciesAsignature struct {
	CreateAsignatureController *controllers.CreateAsignatureController
	RabbitMQAdapter            *adapter.RabbitMQAdapter
}

func InitAsignature() *DependenciesAsignature {
	rmqClient, err := adapter.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("Error creating RabbitMQ client: %v", err)
	}

	messageService := repositories.NewServiceNotification(rmqClient)

	createAsignatureUseCase := application.NewCreateAsignature(rmqClient, messageService)

	return &DependenciesAsignature{
		CreateAsignatureController: controllers.NewCreateAsignatureController(createAsignatureUseCase, rmqClient),
		RabbitMQAdapter:            rmqClient,
	}
}
