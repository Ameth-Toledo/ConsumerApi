package repositories

import "PubNotification/src/notification/domain/entities"

type IMessageService interface {
	PublishEvent(eventType string, asignature entities.Notification) error
}
