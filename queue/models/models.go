package models

import (
	"fmt"

	"github.com/streadway/amqp"
)

type submissionData struct {
	Id   int
	Link string
}

type environmentData struct {
	Id   int
	Link string
}

type testData struct {
	Link string
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
func NewQueueObj(submissionID int, submissionLink string, environmentID int, environmentLink, testLink string) *QueueObj {
	return &QueueObj{
		Submission:  submissionData{Id: submissionID, Link: submissionLink},
		Environment: environmentData{Id: environmentID, Link: environmentLink},
		TestData:    testData{Link: testLink},
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
