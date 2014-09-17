package main

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"github.com/amqp-simplelogger/config"
	"github.com/amqp-simplelogger/message"
	"github.com/amqp-simplelogger/rabbitmq/workQueue"
	"github.com/streadway/amqp"
	"labix.org/v2/mgo"
	"runtime"
	"time"
)

var mongoSession *mgo.Session

type mongoLog struct {
	Message     string
	Timestamp   time.Time
	Type        int32
	FromService string
}

func mqCallback(delivery amqp.Delivery) {
	msg := new(message.LogMessage)
	proto.Unmarshal(delivery.Body, msg)
	fmt.Println(*msg.Message)
	switch msg.GetType() {
	case message.CRITICAL:
		break
	case message.ERROR:
		break
	case message.WARNING:
		break
	case message.NOTIFICATION:
		break
	}
	handleError(msg)
}

// store log message to mongodb
func handleError(msg *message.LogMessage) {
	fmt.Printf("%s %s: %s\n", time.Unix(msg.GetTimestamp(), 0).String(), msg.GetFromService(), msg.GetMessage())
	// insert into mongo db
	c := mongoSession.DB("log").C("log")
	c.Insert(&mongoLog{
		msg.GetMessage(),
		time.Unix(msg.GetTimestamp(), 0),
		msg.GetType(),
		msg.GetFromService(),
	})
}

func Init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var err error
	mongoSession, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
}

func main() {
	Init()
	defer mongoSession.Close()
	consumer := workQueue.NewConsumer()
	err := consumer.Init(config.RABBITMQ_URL, config.RABBITMQ_QUEUE, runtime.NumCPU(), mqCallback)
	defer consumer.Close()
	fmt.Println(err)
}
