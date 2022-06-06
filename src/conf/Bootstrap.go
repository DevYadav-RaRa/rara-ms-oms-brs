package conf

import (
	"fmt"
	"io/ioutil"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/helpers"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
)

func TestBatchCreation() {
	b, err := ioutil.ReadFile("./payload.json") // just pass the file name
	if err != nil {
		fmt.Print("Bootstrap error 1", err)
		return
	}
	fmt.Println(string(b), "Hello")
	fmt.Println(string(b))

}

func Bootstrap(appCtx framework.Framework) {
	fmt.Println("Running Bootstrap...")

	// test if S2 Geo is available else exit with error
	// testS2GeometryOnBoot(appCtx)
	// start listening to queue
	//startSQSConsumer(appCtx)
	// testSQSProducer()
	// inserting demo data..
	// insertDefault(appCtx)
	//TestBatchCreation()
	var demo models.OrderObject
	err, status := helpers.PostOrder(demo, "Business headers")
	if !status {
		fmt.Println("MongoDb Error :", err)
	}
	fmt.Println("App is ready!")
}
