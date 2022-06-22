package conf

import (
	"fmt"
	"log"
	"os"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/aws/sqs"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/services"
	awsqs "github.com/aws/aws-sdk-go/service/sqs"
)

type ApiResponse struct {
	Message     string        `json:"message" bson:"message"`
	OrderStatus []OrderStatus `json:"ordersStatus" bson:"ordersStatus"`
}

type OrderStatus struct {
	TrackingId string `json:"trackingId" bson:"trackingId"`
	Message    string `json:"message" bson:"message"`
	Status     bool   `json:"status" bson:"status"`
}

func startSQSConsumer(appCtx framework.Framework) {
	go sqs.CreateContinuousConsumer(
		os.Getenv("AWS_KEY"),
		os.Getenv("AWS_SECRET"),
		os.Getenv("AWS_REGION"),
		os.Getenv("OMS_BRS_ORDER_QUEUE"),
		4,
		1,
		sqs.SyncConsumer,
		func(msg *awsqs.Message) error {
			log.Println(*msg.Body)
			err := services.OnSQSMessageOrderList(*msg.Body)
			return err
		},
	)
	appCtx.Info("SQS Consumers initialized.")
}

func ConsumeApiOrders(apiOrder string) {
	var resp ApiResponse
	var demoApi models.ApiPayload

	err := demoApi.FromJSONString(apiOrder)
	if err != nil {
		resp.Message = "Invalid API Request Body, Error: " + err.Error()
		return
	}

	fmt.Println("-------------------------------------------")
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------------------------------------")

	framework.Logs("Calling Iam for Authentication")
	var req models.IamRequest
	req.TenantToken = demoApi.TenantToken
	req.BusinessDetails = demoApi.BusinessDetails
	IamAuth := req.GetIamAuthentication("BusinessHeader")

	framework.Logs("Iam Response: ")
	fmt.Println(IamAuth)

	if !IamAuth.Status {
		resp.Message = "Rejected from Iam"
		return
	}

	framework.Logs("Authenticated from Iam")
	resp.Message = "Orders Processing"

	framework.Logs("Pushing to Queue")
	res := sqs.Produce(demoApi)
	fmt.Print("SQS Response: ")
	framework.Logs(res)
	framework.Logs("Pushed to Queue")
	return
}

func Bootstrap(appCtx framework.Framework) {
	framework.Logs("Running Bootstrap...")
	startSQSConsumer(appCtx)
	framework.Logs("App is ready!")
}
