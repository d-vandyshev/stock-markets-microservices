package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"github.com/streadway/amqp"
)

var conf Config

func init() {
	conf.SetFromEnvOrDie()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("ERROR! %s: %s", msg, err)
	}
}

func main() {
	// log.Println("conf.DataPath:", conf.DataPath)
	log.Println("conf.RabbitmqUrl:", conf.RabbitmqUrl)

	// Get list of files
	stocks, err := ioutil.ReadDir(conf.DataPath)
	failOnError(err, "Can't read data directory "+conf.DataPath)

	for _, stock := range stocks {
		if !stock.IsDir() {
			continue
		}
		log.Println(stock.Name())
		filesDirPath := filepath.Join(conf.DataPath, stock.Name(), "csv")
		filenames, err := ioutil.ReadDir(filesDirPath)
		failOnError(err, "Can't read source files directory "+filesDirPath)

		log.Println(len(filenames))
	}
	
	// Connect to RabbitMQ
	conn, err := amqp.Dial(conf.RabbitmqUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "RabbitMQ: failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("data-sources", false,	false, false, false, nil)
	failOnError(err, "RabbitMQ: Failed to declare a queue")
	
	body := "Hello World!"
	err = ch.Publish("", q.Name, false, false, 
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "RabbitMQ: Failed to publish a message")
}
