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
	// Enviar la notificación usando asignatureRepo
	err := c.asignatureRepo.Send(asignature)
	if err != nil {
		return err
	}

	// Usar el serviceNotification para publicar el evento
	err = c.serviceNotification.PublishEvent("AsignatureCreated", asignature) // Usa 'asignature' en lugar de 'created'
	if err != nil {
		log.Printf("Error notificando sobre la asignatura creada: %v", err)
		return err
	}

	return nil
}
