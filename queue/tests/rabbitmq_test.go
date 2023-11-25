package tests

import (
	"testing"
	"time"

	"github.com/priyanshu360/lab-rank/queue/internal"
	"github.com/stretchr/testify/assert"
)

func TestRabbitMQ_PublishConsume(t *testing.T) {
	queueName := "test_queue"

	// Initialize RabbitMQ publisher
	publisher, err := internal.InitRabbitMQPublisher(queueName)
	assert.NoError(t, err, "Error initializing RabbitMQ publisher")
	defer publisher.Close()

	// Initialize RabbitMQ consumer
	consumer, err := internal.InitRabbitMQConsumer(queueName)
	assert.NoError(t, err, "Error initializing RabbitMQ consumer")
	defer consumer.Close()

	// Publish a message
	message := []byte("Hello, RabbitMQ!")
	err = publisher.Publish(message)
	assert.NoError(t, err, "Error publishing message to RabbitMQ")

	// Wait for the consumer to receive the message
	msgq := consumer.GetMsgCh()
	select {
	case delivery := <-msgq:
		receivedMessage := string(delivery.Body)
		// log.Println(receivedMessage)
		assert.Equal(t, string(message), receivedMessage, "Received message does not match published message")
	case <-time.After(5 * time.Second):
		t.Fatal("Timed out waiting for message to be consumed")
	}
}
