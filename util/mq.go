package util

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const MQURL = "amqp://imoocuser:imoocuser@127.0.0.1:5672/imooc"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queueName string
	Exchange  string
	Key       string
	MQurl     string
}

func CreateRabbitMQ(queuename, exchange, key string) *RabbitMQ {
	mq := &RabbitMQ{queueName: queuename, Exchange: exchange, Key: key, MQurl: MQURL}
	var err error
	mq.conn, err = amqp.Dial(mq.MQurl)
	mq.IfFailed(err, "connection to mq failed")
	mq.channel, err = mq.conn.Channel()
	mq.IfFailed(err, "failed to get mq list")
	return mq
}

func (r *RabbitMQ) IfFailed(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
	}
}

func (r *RabbitMQ) SendMessageRouting(message []byte) {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.IfFailed(err, "failed to create ##a exchanger")
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func (r *RabbitMQ) ReceiveMessageRouteing() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.IfFailed(err, "failed to declare a exchanger"+"nge")
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.IfFailed(err, "failed to declare a queue")
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	message, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	list := make(chan bool)
	go func() {
		for d := range message {
			log.Printf("received a message:%s,", d.Body)
		}
	}()
	fmt.Println("press ctrl+c to quit")
	<-list
}
func ConsumerReceive() (msg string, isok bool) {
	return "0", true
}
