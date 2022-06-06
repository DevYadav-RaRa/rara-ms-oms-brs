package helpers

import (
	"fmt"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
)

func PostOrder(order models.OrderObject, header string) (string, bool) {
	IamAuth := order.GetIamAuthentication(header)

	if IamAuth.Status {
		order.OrderDetails.Status = "Processing"
		// 	Save mongo
		err, status := order.Insert()
		if err != nil {
			return err.Error(), status
		}
		// 	Send it to the queue -> Order acceptence
		fmt.Println("Pushing to Queue")
		fmt.Println("Pushed to Queue")

		return "Success: Processing", true
	}
	order.OrderDetails.Status = "Rejected"
	return "Invalid: Iam Rejected", false
}
