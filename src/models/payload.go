package models

import (
	"encoding/json"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
)

type ApiPayload struct {
	TenantToken     string
	BusinessDetails BusinessDetails
	Orders          []Order
}

type CsvPayload struct {
	TrackingId            string
	SenderName            string
	SenderPhone           string
	SenderEmail           string
	SenderLocationName    string
	SenderAddress         string
	SenderSubDistrict     string
	SenderDistrict        string
	SenderCity            string
	SenderPostalCode      string
	SenderLat             float64
	SenderLng             float64
	SenderAddressType     string
	SenderProvince        string
	PackageInchargeName   string
	ExpectedPuDateAndTime string
	Slot                  string
	SenderNote            string
	RecipientName         string
	RecipientPhone        string
	RecipientEmail        string
	RecipientLocationName string
	RecipientAddress      string
	RecipientSubDistrict  string
	RecipientDistrict     string
	RecipientCity         string
	RecipientPostalCode   string
	RecipientLat          float64
	RecipientLng          float64
	RecipientAddressType  string
	RecipientProvince     string
	ReqDlTime             string
	RecipientNote         string
	PackageSize           string
	PackageDescription    string
	PackageValue          string
	PackageLength         float64
	PackageWidth          float64
	PackageHeight         float64
	PackageVolWeight      float64
	PackageBillableWeight float64
	PackageWeightIndex    float64
	PaymentMethod         string
	PackagePrice          float64
	NumberOfPiece         int
	PieceId               string
	PieceWeight           float64
	PieceLength           float64
	PieceWidth            float64
	PieceHeight           float64
	PieceVolWeight        float64
	PieceBillableWeight   float64
	PieceWeightIndex      float64
	PiecePrice            float64
}

func (obj *ApiPayload) FromJSONString(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), obj)
	if err != nil {
		framework.GetCurrentAppContext().Error(err)
		return err
	}
	return nil
}

func (obj *CsvPayload) FromJSONString(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), obj)
	if err != nil {
		framework.GetCurrentAppContext().Error(err)
		return err
	}
	return nil
}

func (obj *OrderObject) FromJSONString(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), obj)
	if err != nil {
		framework.GetCurrentAppContext().Error(err)
		return err
	}
	return nil
}
