package application

import (
	"PubNotification/src/notification/application/repositories"
	"PubNotification/src/notification/domain"
	"PubNotification/src/notification/domain/entities"
	"log"
)

type CreateAsignature struct {
	asignatureRepo      domain.INotification
	serviceNotification repositories.IMessageService
}

func NewCreateAsignature(asignatureRepo domain.INotification, serviceNotification repositories.IMessageService) *CreateAsignature {
	return &CreateAsignature{
		asignatureRepo:      asignatureRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *CreateAsignature) Execute(asignature entities.Notification) error {
	err := c.asignatureRepo.Send(asignature)
	if err != nil {
		return err
	}

	err = c.serviceNotification.PublishEvent("AsignatureCreated", asignature)
	if err != nil {
		log.Printf("Error notificando sobre la asignatura creada: %v", err)
		return err
	}

	return nil
}
