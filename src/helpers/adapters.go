package helpers

import "github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"

type Mapper interface {
	OrderObjectMapper() models.OrderObject
	OrderPieceMapper() models.Piece
}

type ApiPayload models.ApiPayload

type CsvPayload models.CsvPayload

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

func (obj *CsvPayload) OrderObjectMapper() models.OrderObject {
	var ordObj models.OrderObject

	ordObj.Order.TrackingId = obj.TrackingId

	ordObj.Order.PickupDetails.SenderDetails.Name = obj.SenderName
	ordObj.Order.PickupDetails.SenderDetails.PhoneNo = obj.SenderPhone
	ordObj.Order.PickupDetails.SenderDetails.Email = obj.SenderEmail

	ordObj.Order.PickupDetails.LocationDetails.Name = obj.SenderLocationName
	ordObj.Order.PickupDetails.LocationDetails.Address = obj.SenderAddress
	ordObj.Order.PickupDetails.LocationDetails.SubDistrict = obj.SenderSubDistrict
	ordObj.Order.PickupDetails.LocationDetails.District = obj.SenderDistrict
	ordObj.Order.PickupDetails.LocationDetails.City = obj.SenderCity
	ordObj.Order.PickupDetails.LocationDetails.PostalCode = obj.SenderPostalCode
	ordObj.Order.PickupDetails.LocationDetails.LatLng.Lat = obj.SenderLat
	ordObj.Order.PickupDetails.LocationDetails.LatLng.Lng = obj.SenderLng
	ordObj.Order.PickupDetails.LocationDetails.Type = obj.SenderAddressType
	ordObj.Order.PickupDetails.LocationDetails.Province = obj.SenderProvince

	ordObj.Order.PickupDetails.ExpectedPuDateAndTime = obj.ExpectedPuDateAndTime
	ordObj.Order.PickupDetails.PkgIncName = obj.PackageInchargeName
	ordObj.Order.PickupDetails.Slot = obj.Slot
	ordObj.Order.PickupDetails.PuNote.Message = obj.SenderNote

	ordObj.Order.DropOffDetails.RecipientDetails.PhoneNo = obj.RecipientPhone
	ordObj.Order.DropOffDetails.RecipientDetails.Name = obj.RecipientName
	ordObj.Order.DropOffDetails.RecipientDetails.Email = obj.RecipientEmail

	ordObj.Order.DropOffDetails.LocationDetails.Name = obj.RecipientLocationName
	ordObj.Order.DropOffDetails.LocationDetails.Address = obj.RecipientAddress
	ordObj.Order.DropOffDetails.LocationDetails.SubDistrict = obj.RecipientSubDistrict
	ordObj.Order.DropOffDetails.LocationDetails.District = obj.RecipientDistrict
	ordObj.Order.DropOffDetails.LocationDetails.City = obj.RecipientCity
	ordObj.Order.DropOffDetails.LocationDetails.PostalCode = obj.RecipientPostalCode
	ordObj.Order.DropOffDetails.LocationDetails.LatLng.Lat = obj.RecipientLat
	ordObj.Order.DropOffDetails.LocationDetails.LatLng.Lng = obj.RecipientLng
	ordObj.Order.DropOffDetails.LocationDetails.Type = obj.RecipientAddressType
	ordObj.Order.DropOffDetails.LocationDetails.Province = obj.RecipientProvince

	ordObj.Order.DropOffDetails.ReqDlTime = obj.ReqDlTime
	ordObj.Order.DropOffDetails.DlNote.Message = obj.RecipientNote

	ordObj.Order.PackageDetails.Size = obj.PackageSize
	ordObj.Order.PackageDetails.Description = obj.PackageDescription
	ordObj.Order.PackageDetails.Value = obj.PackageValue
	ordObj.Order.PackageDetails.NoOfPieces = obj.NumberOfPiece
	ordObj.Order.PackageDetails.Dimensions.Length = obj.PackageLength
	ordObj.Order.PackageDetails.Dimensions.Width = obj.PackageWidth
	ordObj.Order.PackageDetails.Dimensions.Height = obj.PackageHeight
	ordObj.Order.PackageDetails.VolWeight = obj.PackageVolWeight
	ordObj.Order.PackageDetails.BillableWeight = obj.PackageBillableWeight
	ordObj.Order.PackageDetails.WeightIndex = obj.PackageWeightIndex

	ordObj.Order.PaymentDetails.Method = obj.PaymentMethod
	ordObj.Order.PaymentDetails.Price = obj.PackagePrice

	ordPie := obj.OrderPieceMapper()

	ordObj.Order.Items = append(ordObj.Order.Items, ordPie)

	return ordObj
}

func (obj *CsvPayload) OrderPieceMapper() models.Piece {
	var ordPie models.Piece

	ordPie.PieceId = obj.PieceId
	ordPie.Weight = obj.PieceWeight

	ordPie.Dimensions.Length = obj.PieceLength
	ordPie.Dimensions.Width = obj.PieceWidth
	ordPie.Dimensions.Height = obj.PieceHeight

	ordPie.VolWeight = obj.PieceVolWeight
	ordPie.BillableWeight = obj.PieceBillableWeight
	ordPie.WeightIndex = obj.PieceWeightIndex
	ordPie.Price = obj.PiecePrice

	return ordPie
}
