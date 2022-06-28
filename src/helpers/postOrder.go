package helpers

import (
	"errors"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
)

func PostOrder(order models.OrderObject) (error, bool) {
	order.OrderDetails.Status = "Processing"

	framework.Logs("Inserting to MongoDb")

	// Dont's save same order again and return
	errFind, _ := order.FindByTrackingId() // and status != failed (another status in order object for order placement either success or failed)
	if errFind.OrderDetails.TrackingId != "" {
		return errors.New("order processing"), false
	}

	err, status := order.Insert()
	if err != nil {
		return err, status
	}

	framework.Logs("Inserted to MongoDb")

	// 	Send it to the queue -> Order acceptence
	return errors.New("success: processing"), true
}
