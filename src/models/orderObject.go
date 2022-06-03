package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const db_name = "bms"

const collection_name = "batches"

type OrderObject struct {
	TenantToken     string
	BusinessDetails BusinessDetails
	OrderDetails    OrderDetails
	Order           Order
	Webhook         Webhook
}

type Order struct {
	TrackingId     string
	PickupDetails  PickupDetails
	DropOffDetails DropOffDetails
	PackageDetails PackageDetails
	PaymentDetails PaymentDetails
	Items          []Piece
}

type BusinessDetails struct {
	BusinessName string
	AccNo        string
	ServiceType  string
	ServiceId    string
}

type OrderDetails struct {
	Status                string
	OrderId               primitive.ObjectID
	PieceId               string
	DeliveryFee           float64
	OrderDeliveryDetails  OrderDeliveryDetails
	OrderDimensionDetails OrderDimensionDetails
	BatchDetails          BatchDetails
}

type OrderDeliveryDetails struct {
	OrderDate       string
	PickupDate      string
	DlSla           float64
	OrderDistance   float64
	Linehaul        string
	SpecialHandling string
}

type OrderDimensionDetails struct {
	Weight     float64
	Dimensions Dimensions
	VolWeight  float64
}

type Dimensions struct {
	Length float64
	Width  float64
	Height float64
}

type BatchDetails struct {
	BatchCreationTime  string
	BatchId            primitive.ObjectID
	OrderCancelledTime string
	OrderCancelReason  string
	OrderReturnedTime  string
}

type PickupDetails struct {
	SenderDetails         PersonalDetails
	LocationDetails       LocationDetails
	PkgIncName            string
	ExpectedPuDateAndTime string
	Slot                  string
	PuNote                Note
}

type DropOffDetails struct {
	RecipientDetails PersonalDetails
	LocationDetails  LocationDetails
	ReqDlTime        string
	DlNote           Note
}

type PersonalDetails struct {
	Name    string
	Email   string
	PhoneNo string
}

type LocationDetails struct {
	Name        string
	Address     string
	SubDistrict string
	District    string
	City        string
	PostalCode  string
	LatLng      GeoPoint
	Type        string
	Province    string
}

type GeoPoint struct {
	Lat float64
	Lng float64
}

type Note struct {
	Message string
}

type PackageDetails struct {
	Size           string
	Description    string
	Value          string
	NoOfPieces     int
	Dimensions     Dimensions
	VolWeight      float64
	BillableWeight float64
	WeightIndex    float64
}

type PaymentDetails struct {
	Method string
	Price  float64
}

type Piece struct {
	OrderId        primitive.ObjectID
	PieceId        string
	Weight         float64
	Dimensions     Dimensions
	VolWeight      float64
	BillableWeight float64
	WeightIndex    float64
	Price          float64
}

type Webhook struct {
	Url     string
	Header  Header
	Payload string
}

type Header struct {
}

func (obj *OrderObject) Insert() (error, bool) {
	err, ctx, appCtx, db, cancel := getDbContext(db_name, 10*time.Second)
	defer cancel()
	if err != nil {
		appCtx.Error(err)
		return err, false
	}

	// obj.OrderDetails.OrderId = primitive.NewObjectID()

	doc, err := toDoc(obj)
	if err != nil {
		appCtx.Error(err)
		return err, false
	}
	res, err := db.Collection(collection_name).InsertOne(ctx, doc)
	if err != nil {
		appCtx.Error(err)
		return err, false
	}
	fmt.Println(res)
	// obj.OrderDetails.OrderId = res.InsertedID.(primitive.ObjectID)

	return nil, true
}

func (obj *OrderObject) FindById() (interface{}, bool) {
	err, ctx, appCtx, db, cancel := getDbContext(db_name, 15*time.Second)
	defer cancel()
	if err != nil {
		appCtx.Error(err)
		return nil, false
	}

	var data map[string]interface{} = make(map[string]interface{})
	res := db.Collection(collection_name).FindOne(ctx, bson.M{"_id": obj.OrderDetails.OrderId})
	res.Decode(&data)
	return data, true
}

func (obj *OrderObject) Save() (bool, error) {
	err, ctx, appCtx, db, cancel := getDbContext(db_name, 5*time.Second)

	defer cancel()
	doc, err := toDoc(obj)
	if err != nil {
		appCtx.Error(err)
		return false, err
	}
	res, err := db.Collection(collection_name).UpdateOne(ctx, bson.M{"_id": obj.OrderDetails.OrderId}, bson.M{"$set": doc})
	if err != nil {
		appCtx.Error(err)
		return false, err
	}
	fmt.Println(res.ModifiedCount)
	return true, err
}
