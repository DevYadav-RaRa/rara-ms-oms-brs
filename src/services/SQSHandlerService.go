package services

import (
	"fmt"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/helpers"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
)

func OnSQSMessageOrderList(jsonBody string) error {
	framework.Logs("API order received by Queue")
	demoApi := models.ApiPayload{}
	err := demoApi.FromJSONString(jsonBody)
	if err != nil {
		fmt.Println("SHS error 1", err)
		return err
	}

	for i := range demoApi.Orders {
		var temp models.OrderObject
		temp.TenantToken = demoApi.TenantToken
		temp.BusinessDetails = demoApi.BusinessDetails

		temp.PickupDetails = demoApi.Orders[i].PickupDetails
		temp.DropOffDetails = demoApi.Orders[i].DropOffDetails
		temp.PackageDetails = demoApi.Orders[i].PackageDetails
		temp.PaymentDetails = demoApi.Orders[i].PaymentDetails
		temp.Pieces = demoApi.Orders[i].Pieces

		err, status := helpers.PostOrder(temp)

		if !status {
			return err
		}

		// var orderResp OrderStatus
		// orderResp.TrackingId = temp.OrderDetails.TrackingId
		// orderResp.Status = status
		// orderResp.Message = err.Error()

		// resp.OrderStatus = append(resp.OrderStatus, orderResp)

		fmt.Println("-------------------------------------------")
		fmt.Println("-------------------------------------------")
	}

	return nil
}
