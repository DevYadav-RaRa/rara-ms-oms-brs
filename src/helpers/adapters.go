package helpers

import "github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"

type Mapper interface {
	OrderObjectMapper() models.OrderObject
	OrderPieceMapper() models.Piece
}

type ApiPayload models.ApiPayload

// type CsvPayload models.CsvPayload

func (obj *ApiPayload) OrderObjectMapper() []models.OrderObject {
	var response []models.OrderObject

	for i := range obj.Orders {
		var ordObj models.OrderObject

		ordObj.TenantToken = obj.TenantToken
		ordObj.BusinessDetails = obj.BusinessDetails
		ordObj.Order = obj.Orders[i]

		response = append(response, ordObj)
	}

	return response
}

// func (obj *CsvPayload) OrderObjectMapper() models.OrderObject {
// 	var ordObj models.OrderObject
// 	return ordObj
// }

// func (obj *CsvPayload) OrderPieceMapper() models.Piece {
// 	var ordPie models.Piece
// 	return ordPie
// }
