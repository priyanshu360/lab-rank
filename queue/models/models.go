package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type submissionData struct {
	id   uuid.UUID
	link string
}

type environmentData struct {
	id   uuid.UUID
	link string
}

type testData struct {
	link string
}

type QueueObj struct {
	Submission  submissionData
	Environment environmentData
	TestData    testData
}

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
	msgq <-chan amqp.Delivery
}

// Todo : dirty constructor
func NewQueueObj(submissionID uuid.UUID, submissionLink string, environmentID uuid.UUID, environmentLink, testLink string) *QueueObj {
	return &QueueObj{
		Submission:  submissionData{id: submissionID, link: submissionLink},
		Environment: environmentData{id: environmentID, link: environmentLink},
		TestData:    testData{link: testLink},
	}
}

func NewRabbitMQ(conn *amqp.Connection, ch *amqp.Channel, q amqp.Queue, msgq <-chan amqp.Delivery) *RabbitMQ {
	return &RabbitMQ{
		conn: conn,
		ch:   ch,
		q:    q,
		msgq: msgq,
	}
}

func (r *RabbitMQ) Close() {
	r.ch.Close()
	r.conn.Close()
}

func (r *RabbitMQ) Publish(body []byte) error {

	err := r.ch.Publish(
		"",       // exchange
		r.q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("Failed to publish a message: %s", err)
	}

	return nil
}

func (r *RabbitMQ) GetMsgCh() <-chan amqp.Delivery {
	return r.msgq
}
