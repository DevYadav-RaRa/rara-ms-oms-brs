package helpers

import (
	"fmt"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
)

func PostOrder(order models.OrderObject) (string, bool) {
	order.OrderDetails.Status = "Processing"
	// 	Save mongo
	fmt.Println("Inserting to MongoDb")
	fmt.Println("Inserted to MongoDb")
	// Print struct order
	fmt.Println("Order Before Mongo: ", order)
	err, status := order.Insert()
	if err != nil {
		return err.Error(), status
	}
	// 	Send it to the queue -> Order acceptence
	fmt.Println("Pushing to Queue")
	fmt.Println("Pushed to Queue")

	return "Success: Processing", true
}
