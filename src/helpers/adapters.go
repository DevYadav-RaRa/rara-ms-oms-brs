package helpers

import "github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"

type Mapper interface {
	OrderObjectMapper() models.OrderObject
	OrderPieceMapper() models.Piece
}

type ApiPayload models.ApiPayload

func (obj *ApiPayload) OrderObjectMapper() []models.OrderObject {
	var response []models.OrderObject

	for i := range obj.Orders {
		var ordObj models.OrderObject

		ordObj.TenantToken = obj.TenantToken
		ordObj.BusinessDetails = obj.BusinessDetails
		ordObj.PickupDetails = obj.Orders[i].PickupDetails
		ordObj.DropOffDetails = obj.Orders[i].DropOffDetails
		ordObj.PackageDetails = obj.Orders[i].PackageDetails
		ordObj.PaymentDetails = obj.Orders[i].PaymentDetails
		ordObj.Pieces = obj.Orders[i].Pieces

		response = append(response, ordObj)
	}

	return response
}
