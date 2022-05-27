package sqs

import (
	"context"
	"fmt"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/aws"
	"log"
	"time"
)

func CreateContinuousConsumer(awsKey string, awsSecret string, awsRegion string, queueUrl string, maxWorker int, maxMessage int, consumerType ConsumerType, onMessageCallback SQSOnMessageCallback) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Create a session instance.
	ses, err := aws.New(aws.Config{
		Region: awsRegion,
		ID:     awsKey,
		Secret: awsSecret,
	})
	if err != nil {
		log.Fatalln(err)
	}
	// log
	fmt.Println("------------------CONSUMER------------------")
	fmt.Println("QUEUE                            :", queueUrl)
	fmt.Println("NUM WORKERS                      :", maxWorker)
	fmt.Println("NUM Messages To be picked at once:", maxMessage)
	fmt.Println("Consumer Type                    :", consumerType)
	fmt.Println("--------------------------------------------")
	// Instantiate client.
	client := NewSQS(ses, time.Second*5)
	// Instantiate consumer and start consuming.
	NewConsumer(client, ConsumerConfig{
		Type:      consumerType,
		QueueURL:  queueUrl,
		MaxWorker: maxWorker,  // 2
		MaxMsg:    maxMessage, // 10
	}, onMessageCallback).Start(ctx)
}
