package main

import (
	"github.com/upb-code-labs/tests-microservice/infrastructure/rabbitmq"
)

func main() {
	// Setup rabbitmq connection
	rabbitmq.ConnectToRabbitMQ()
	defer rabbitmq.CloseRabbitMQConnection()

	// Start listening to submissions queue
	submissionsQueueManager := rabbitmq.GetSubmissionQueueManager()
	submissionsQueueManager.ListenForSubmissions()

	// Block forever
	var forever chan bool
	<-forever
}
