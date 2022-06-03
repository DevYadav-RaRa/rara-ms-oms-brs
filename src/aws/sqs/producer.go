package sqs

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/aws"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
)

func Produce(obj models.OrderObject) string {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Create a session instance.
	ses, err := aws.New(aws.Config{
		Region: os.Getenv("AWS_REGION"),
		ID:     os.Getenv("AWS_KEY"),
		Secret: os.Getenv("AWS_SECRET"),
	})
	if err != nil {
		log.Fatalln(err)
	}

	// creating sqs client
	SQSclient := NewSQS(ses, 5*time.Second)

	var attributes []Attribute = make([]Attribute, 0)
	bat, _ := json.Marshal(obj)
	atb := Attribute{
		Key:   "batch",
		Value: string(bat),
		Type:  "String",
	}

	attributes = append(attributes, atb)

	// making send request
	req := SendRequest{
		QueueURL:   os.Getenv("AWS_PRODUCER_QUEUE"),
		Body:       "BATCH SENT TO MONOLITH SERVER",
		Attributes: attributes,
	}

	res, err := SQSclient.Send(ctx, &req)
	if err != nil {
		log.Fatalln(err)
	}
	return res
}
