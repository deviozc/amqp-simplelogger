package workQueue

import (
	"github.com/streadway/amqp"
)

type Consumer struct {
	done   chan bool
	closed bool
}

func NewConsumer() *Consumer {
	done := make(chan bool)
	return &Consumer{done, false}
}

func worker(msgs <-chan amqp.Delivery, fn func(amqp.Delivery), workerId int) {
	for msg := range msgs {
		msg.Ack(false)
		fn(msg)
	}
}

// Use the default exchange
func (this *Consumer) Init(url string, queueName string, numWorkers int, fn func(amqp.Delivery)) error {
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	_, err = ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return err
	}
	msgs, err := ch.Consume(queueName, "", false, false, false, false, nil)

	func() {
		for i := 0; i < numWorkers; i++ {
			// msgs, _ := newConsumeChannel(conn, queueName)

			go worker(msgs, fn, i+1)
		}
	}()

	<-this.done
	return nil
}
func newConsumeChannel(conn *amqp.Connection, queueName string) (<-chan amqp.Delivery, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch.Consume(queueName, "", false, false, false, false, nil)
}

func (this *Consumer) Close() {
	if this.closed {
		return
	}
	this.closed = true
	this.done <- true
}
