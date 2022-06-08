package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const db_name = "oms"

const collection_name = "orders"

type OrderObject struct {
	TenantToken     string          `json:"tenantToken" bson:"tenantToken"`
	BusinessDetails BusinessDetails `json:"businessDetails" bson:"businessDetails"`
	OrderDetails    OrderDetails    `json:"orderDetails" bson:"orderDetails"`
	Order           Order           `json:"order" bson:"order"`
	Webhook         Webhook         `json:"webhook" bson:"webhook"`
}

type Order struct {
	TrackingId     string         `json:"trackingId" bson:"trackingId"`
	PickupDetails  PickupDetails  `json:"pickupDetails" bson:"pickupDetails"`
	DropOffDetails DropOffDetails `json:"dropoffDetails" bson:"dropoffDetails"`
	PackageDetails PackageDetails `json:"packageDetails" bson:"packageDetails"`
	PaymentDetails PaymentDetails `json:"paymentDetails" bson:"paymentDetails"`
	Items          []Piece        `json:"items" bson:"items"`
}

type BusinessDetails struct {
	BusinessName string `json:"businessName" bson:"businessName"`
	AccNo        string `json:"accountNumber" bson:"accountNumber"`
	ServiceType  string `json:"serviceType" bson:"serviceType"`
	ServiceId    string `json:"serviceId" bson:"serviceId"`
}

type OrderDetails struct {
	Status                string                `json:"status" bson:"status"`
	OrderId               primitive.ObjectID    `json:"orderId" bson:"orderId"`
	PieceId               string                `json:"pieceId" bson:"pieceId"`
	DeliveryFee           float64               `json:"deliveryFee" bson:"deliveryFee"`
	OrderDeliveryDetails  OrderDeliveryDetails  `json:"orderDeliveryDetails" bson:"orderDeliveryDetails"`
	OrderDimensionDetails OrderDimensionDetails `json:"orderDimensionDetails" bson:"orderDimensionDetails"`
	BatchDetails          BatchDetails          `json:"batchDetails" bson:"batchDetails"`
}

type OrderDeliveryDetails struct {
	OrderDate       string  `json:"orderDate" bson:"orderDate"`
	PickupDate      string  `json:"pickupDate" bson:"pickupDate"`
	DlSla           float64 `json:"dlSla" bson:"dlSla"`
	OrderDistance   float64 `json:"orderDistance" bson:"orderDistance"`
	Linehaul        string  `json:"linehaul" bson:"linehaul"`
	SpecialHandling string  `json:"specialHandling" bson:"specialHandling"`
}

type OrderDimensionDetails struct {
	Weight     float64    `json:"weight" bson:"weight"`
	Dimensions Dimensions `json:"dimensions" bson:"dimensions"`
	VolWeight  float64    `json:"volWeight" bson:"volWeight"`
}

type Dimensions struct {
	Length float64 `json:"length" bson:"length"`
	Width  float64 `json:"width" bson:"width"`
	Height float64 `json:"height" bson:"height"`
}

type BatchDetails struct {
	BatchCreationTime  string             `json:"batchCreationTime" bson:"batchCreationTime"`
	BatchId            primitive.ObjectID `json:"batchId" bson:"batchId"`
	OrderCancelledTime string             `json:"orderCancelledTime" bson:"orderCancelledTime"`
	OrderCancelReason  string             `json:"orderCancelReason" bson:"orderCancelReason"`
	OrderReturnedTime  string             `json:"orderReturnedTime" bson:"orderReturnedTime"`
}

type PickupDetails struct {
	SenderDetails         PersonalDetails `json:"senderDetails" bson:"senderDetails"`
	LocationDetails       LocationDetails `json:"locationDetails" bson:"locationDetails"`
	PkgIncName            string          `json:"packageInchargeName" bson:"packageInchargeName"`
	ExpectedPuDateAndTime string          `json:"expectedPuDate&Time" bson:"expectedPuDate&Time"`
	Slot                  string          `json:"slot" bson:"slot"`
	PuNote                Note            `json:"puNote" bson:"puNote"`
}

type DropOffDetails struct {
	RecipientDetails PersonalDetails `json:"recipientDetails" bson:"recipientDetails"`
	LocationDetails  LocationDetails `json:"locationDetails" bson:"locationDetails"`
	ReqDlTime        string          `json:"reqDlTime" bson:"reqDlTime"`
	DlNote           Note            `json:"dlNote" bson:"dlNote"`
}

type PersonalDetails struct {
	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email"`
	PhoneNo string `json:"phone" bson:"phone"`
}

type LocationDetails struct {
	Name        string   `json:"locationName" bson:"locationName"`
	Address     string   `json:"address" bson:"address"`
	SubDistrict string   `json:"subDistrict" bson:"subDistrict"`
	District    string   `json:"district" bson:"district"`
	City        string   `json:"city" bson:"city"`
	PostalCode  string   `json:"postalCode" bson:"postalCode"`
	LatLng      GeoPoint `json:"geoPoint" bson:"geoPoint"`
	Type        string   `json:"addressType" bson:"addressType"`
	Province    string   `json:"province" bson:"province"`
}

type GeoPoint struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}

type Note struct {
	Message string `json:"message" bson:"message"`
}

type PackageDetails struct {
	Size           string     `json:"packageSize" bson:"packageSize"`
	Description    string     `json:"packageDescription" bson:"packageDescription"`
	Value          string     `json:"packageValue" bson:"packageValue"`
	NoOfPieces     int        `json:"numberofPiece" bson:"numberofPiece"`
	Dimensions     Dimensions `json:"dimensions" bson:"dimensions"`
	VolWeight      float64    `json:"volWeight" bson:"volWeight"`
	BillableWeight float64    `json:"billableWeight" bson:"billableWeight"`
	WeightIndex    float64    `json:"weightIndex" bson:"weightIndex"`
}

type PaymentDetails struct {
	Method string  `json:"paymentMethod" bson:"paymentMethod"`
	Price  float64 `json:"price" bson:"price"`
}

type Piece struct {
	OrderId        primitive.ObjectID `json:"orderId" bson:"orderId"`
	PieceId        string             `json:"pieceId" bson:"pieceId"`
	Weight         float64            `json:"weight" bson:"weight"`
	Dimensions     Dimensions         `json:"dimensions" bson:"dimensions"`
	VolWeight      float64            `json:"volWeight" bson:"volWeight"`
	BillableWeight float64            `json:"billableWeight" bson:"billableWeight"`
	WeightIndex    float64            `json:"weightIndex" bson:"weightIndex"`
	Price          float64            `json:"price" bson:"price"`
}

type Webhook struct {
	Url     string  `json:"url" bson:"url"`
	Headers Headers `json:headers bson:headers`
	Payload string  `json:"payload" bson:"payload"`
}

type Headers struct {
	header string `json:"header" bson:"header"`
}

func (obj *IamRequest) GetIamAuthentication(header string) IamResponse {
	var reqObj IamRequest
	reqObj.TenantToken = obj.TenantToken
	reqObj.BusinessDetails = obj.BusinessDetails

	// Call IM using Request struct and headers
	// Get response in the form of IamResponse struct

	// ONLY FOR TESTING PURPOSES
	responseObj := reqObj.GetIamResponse(header)
	// ONLY FOR TESTING PURPOSES

	return responseObj
}

func (obj *OrderObject) Insert() (error, bool) {
	err, ctx, appCtx, db, cancel := getDbContext(db_name, 10*time.Second)
	defer cancel()
	if err != nil {
		appCtx.Error(err)
		return err, false
	}

	obj.OrderDetails.OrderId = primitive.NewObjectID()
	for i := range obj.Order.Items {
		obj.Order.Items[i].OrderId = obj.OrderDetails.OrderId
	}
	fmt.Println("OrderObject: ", obj)
	doc, err := toDoc(obj)
	if err != nil {
		appCtx.Error(err)
		return err, false
	}

	fmt.Println("collection_name: ", collection_name)
	fmt.Println("doc: ", doc)
	fmt.Println("ctx: ", ctx)
	res, err := db.Collection(collection_name).InsertOne(ctx, doc)
	fmt.Println("res: ", res)
	if err != nil {
		appCtx.Error(err)
		return err, false
	}
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
