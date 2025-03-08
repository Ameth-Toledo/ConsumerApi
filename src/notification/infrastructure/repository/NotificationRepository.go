package repository

import (
	"PubNotification/src/notification/domain/entities"
	"PubNotification/src/notification/infrastructure/adapter"
)

type NotificationRepository struct {
	rmqClient *adapter.RabbitMQAdapter
}

func NewNotificationRepository(rmqClient *adapter.RabbitMQAdapter) *NotificationRepository {
	return &NotificationRepository{rmqClient: rmqClient}
}

func (nr *NotificationRepository) Send(asignature entities.Notification) error {
	return nr.rmqClient.Send(asignature)
}
