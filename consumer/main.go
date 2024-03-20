package main
import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	channel, err := conn.Channel()
	defer channel.Close()

	queue, err := channel.QueueDeclare( "hello", false, false, false, false, nil,)

	messages, err := channel.Consume( queue.Name, "", true, false, false, false, nil,)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			fmt.Printf("Received a message: %s\n", message.Body)
		}
	}()

	fmt.Println("Waiting for messages. To exit press CTRL+C")
	<-forever
}
