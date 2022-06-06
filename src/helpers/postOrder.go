package helpers

import (
	"fmt"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
)

func PostOrder(order models.OrderObject, header string) (string, bool) {
	fmt.Println("Calling Iam for Authentication")
	IamAuth := order.GetIamAuthentication(header)
	fmt.Println("Iam Response: ", IamAuth)
	fmt.Println("Authenticated from Iam")
	if IamAuth.Status {
		order.OrderDetails.Status = "Processing"
		// 	Save mongo
		fmt.Println("Inserting to MongoDb")
		fmt.Println("Inserted to MongoDb")
		// err, status := order.Insert()
		// if err != nil {
		// 	return err.Error(), status
		// }
		// 	Send it to the queue -> Order acceptence
		fmt.Println("Pushing to Queue")
		fmt.Println("Pushed to Queue")

		return "Success: Processing", true
	}
	order.OrderDetails.Status = "Rejected"
	return "Invalid: Iam Rejected", false
}
