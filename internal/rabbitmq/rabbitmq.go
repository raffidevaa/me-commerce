package rabbitmq

import (
	"context"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn *amqp.Connection
}

func NewRabbitMQConnection() *RabbitMQ {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	log.Println("Connected to RabbitMQ")

	return &RabbitMQ{Conn: conn}
}

func (r *RabbitMQ) Publish(queue string, body []byte) error {
	ch, err := r.Conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queue,
		true, // durable ✅
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return ch.PublishWithContext(
		ctx,
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (r *RabbitMQ) Consume(queue string, handler func([]byte) error) {
	ch, err := r.Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	_, err = ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		queue,
		"",
		false, // ❗ manual ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for msg := range msgs {
			err := handler(msg.Body)
			if err != nil {
				msg.Nack(false, true) // retry
				continue
			}
			msg.Ack(false)
		}
	}()
}
