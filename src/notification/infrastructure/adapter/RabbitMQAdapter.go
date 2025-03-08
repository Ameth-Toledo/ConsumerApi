package adapter

import (
	"PubNotification/src/notification/domain"
	"PubNotification/src/notification/domain/entities"
	"github.com/goccy/go-json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitMQAdapter struct {
	conn      *amqp.Connection
	ch        *amqp.Channel
	queueName string
}

// Asegura que RabbitMQAdapter implemente la interfaz INotification
var _ domain.INotification = (*RabbitMQAdapter)(nil)

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	conn, err := amqp.Dial("amqp://toledo:12345@35.170.134.124:5672/")
	if err != nil {
		log.Println("Error al conectar a RabbitMQ:", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Error al abrir el canal:", err)
		return nil, err
	}

	queueName := "notification"
	_, err = ch.QueueDeclare(
		queueName, // Nombre de la cola
		true,      // Duradero
		false,     // Auto-borrar
		false,     // Exclusivo
		false,     // Sin esperar
		nil,       // Argumentos
	)
	if err != nil {
		log.Println("Error al declarar la cola:", err)
		return nil, err
	}

	err = ch.Confirm(false)
	if err != nil {
		log.Println("Error al habilitar las confirmaciones de mensajes:", err)
		return nil, err
	}

	return &RabbitMQAdapter{
		conn:      conn,
		ch:        ch,
		queueName: queueName,
	}, nil
}

func (r *RabbitMQAdapter) Send(asignature entities.Notification) error {
	log.Println("Enviando notificación a través de RabbitMQ")

	err := r.PublishEvent("Notification", asignature)
	if err != nil {
		log.Printf("Error al enviar la notificación: %v", err)
		return err
	}

	log.Println("Notificación enviada con éxito")
	return nil
}

func (r *RabbitMQAdapter) PublishEvent(eventType string, asignature entities.Notification) error {
	body, err := json.Marshal(asignature)
	if err != nil {
		log.Println("Error al convertir el evento a JSON:", err)
		return err
	}

	ack, nack := r.ch.NotifyConfirm(make(chan uint64, 1), make(chan uint64, 1))

	err = r.ch.Publish(
		"",          // Intercambio
		r.queueName, // Usamos el nombre de la cola del adaptador, asegurando que publicamos en la cola correcta
		true,        // Obligatorio
		false,       // Inmediato
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Println("Error al enviar el mensaje a RabbitMQ:", err)
		return err
	}

	select {
	case <-ack:
		log.Println("Mensaje confirmado por RabbitMQ")
	case <-nack:
		log.Println("El mensaje no fue confirmado")
	}

	log.Println("Evento publicado:", eventType)
	return nil
}

func (r *RabbitMQAdapter) Close() {
	if r.ch != nil {
		r.ch.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
