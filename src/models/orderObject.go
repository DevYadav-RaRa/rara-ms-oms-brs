package models

import (
	"fmt"
	"log"
	"time"

	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

const db_name = "oms"

const collection_name = "orders"

type OrderObject struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id"`
	UploadId        string             `json:"uploadId" bson:"uploadId"`
	TenantToken     string             `json:"tenantToken" bson:"tenantToken"`
	BusinessDetails BusinessDetails    `json:"businessDetails" bson:"businessDetails"`
	OrderDetails    OrderDetails       `json:"orderDetails" bson:"orderDetails"`
	PickupDetails   PickupDetails      `json:"pickupDetails" bson:"pickupDetails"`
	DropOffDetails  DropOffDetails     `json:"dropoffDetails" bson:"dropoffDetails"`
	PackageDetails  PackageDetails     `json:"packageDetails" bson:"packageDetails"`
	PaymentDetails  PaymentDetails     `json:"paymentDetails" bson:"paymentDetails"`
	Pieces          []Piece            `json:"pieces" bson:"pieces"`
	Webhook         Webhook            `json:"webhook" bson:"webhook"`
}

type Order struct {
	TrackingId     string         `json:"trackingId" bson:"trackingId"`
	PickupDetails  PickupDetails  `json:"pickupDetails" bson:"pickupDetails"`
	DropOffDetails DropOffDetails `json:"dropoffDetails" bson:"dropoffDetails"`
	PackageDetails PackageDetails `json:"packageDetails" bson:"packageDetails"`
	PaymentDetails PaymentDetails `json:"paymentDetails" bson:"paymentDetails"`
	Pieces         []Piece        `json:"pieces" bson:"pieces"`
}

type BusinessDetails struct {
	BusinessName string `json:"businessName" bson:"businessName"`
	AccNo        string `json:"accountNumber" bson:"accountNumber"`
	ServiceType  string `json:"serviceType" bson:"serviceType"`
	ServiceId    string `json:"serviceId" bson:"serviceId"`
}

type OrderDetails struct {
	Status                string                `json:"status" bson:"status"`
	TrackingId            string                `json:"trackingId" bson:"trackingId"`
	PieceId               string                `json:"pieceId" bson:"pieceId"`
	DeliveryFee           float64               `json:"deliveryFee" bson:"deliveryFee"`
	Amount                float64               `json:"amount" bson:"amount"`
	OrderDeliveryDetails  OrderDeliveryDetails  `json:"orderDeliveryDetails" bson:"orderDeliveryDetails"`
	OrderDimensionDetails OrderDimensionDetails `json:"orderDimensionDetails" bson:"orderDimensionDetails"`
}

type OrderDeliveryDetails struct {
	OrderDate       string  `json:"orderDate" bson:"orderDate"`
	PickupDate      string  `json:"pickupDate" bson:"pickupDate"`
	OrderDistance   float64 `json:"orderDistance" bson:"orderDistance"`
	Linehaul        bool    `json:"linehaul" bson:"linehaul"`
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
	Unit   string  `json:"unit" bson:"unit"`
}

type PickupDetails struct {
	PickupInchargeDetails PersonalDetails `json:"pickupIncharge" bson:"pickupIncharge"`
	LocationDetails       LocationDetails `json:"locationDetails" bson:"locationDetails"`
	ExpectedPuDateAndTime string          `json:"expectedPuDateandTime" bson:"expectedPuDateandTime"`
	Slot                  string          `json:"slot" bson:"slot"`
	PuNote                string          `json:"puNote" bson:"puNote"`
}

type DropOffDetails struct {
	RecipientDetails PersonalDetails `json:"recipientDetails" bson:"recipientDetails"`
	LocationDetails  LocationDetails `json:"locationDetails" bson:"locationDetails"`
	ReqDlTime        string          `json:"reqDlTime" bson:"reqDlTime"`
	DlNote           string          `json:"dlNote" bson:"dlNote"`
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
	OrderId        primitive.ObjectID `json:"orderId,omitempty" bson:"orderId,omitempty"`
	PieceId        string             `json:"pieceId" bson:"pieceId"`
	Weight         float64            `json:"weight" bson:"weight"`
	Dimensions     Dimensions         `json:"dimensions" bson:"dimensions"`
	VolWeight      float64            `json:"volWeight" bson:"volWeight"`
	BillableWeight float64            `json:"billableWeight" bson:"billableWeight"`
	WeightIndex    float64            `json:"weightIndex" bson:"weightIndex"`
	Price          float64            `json:"price" bson:"price"`
}

func (obj *IamRequest) GetIamAuthentication(header string) IamResponse {
	var reqObj IamRequest
	reqObj.TenantToken = obj.TenantToken
	reqObj.BusinessDetails = obj.BusinessDetails

	// Call Iam using Request struct and headers
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

	obj.Id = primitive.NewObjectID()

	for i := range obj.Pieces {
		obj.Pieces[i].OrderId = obj.Id
	}

	doc, err := toDoc(obj)
	if err != nil {
		appCtx.Error(err)
		return err, false
	}

	_, err = db.Collection(collection_name).InsertOne(ctx, doc)

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
	res := db.Collection(collection_name).FindOne(ctx, bson.M{"_id": obj.Id})
	res.Decode(&data)
	return data, true
}

func (obj *OrderObject) FindByTrackingId() (OrderObject, bool) {
	err, ctx, appCtx, db, cancel := getDbContext(db_name, 15*time.Second)
	defer cancel()
	var data OrderObject
	if err != nil {
		appCtx.Error(err)
		return data, false
	}

	res := db.Collection(collection_name).FindOne(ctx, bson.M{"order.trackingId": obj.OrderDetails.TrackingId})
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
	res, err := db.Collection(collection_name).UpdateOne(ctx, bson.M{"_id": obj.Id}, bson.M{"$set": doc})
	if err != nil {
		appCtx.Error(err)
		return false, err
	}
	fmt.Println(res.ModifiedCount)
	return true, err
}

func FindBy(filters map[string]string, page int64, limit int64) ([]OrderObject, paginate.PaginationData, bool) {
	err, ctx, appCtx, db, cancel := getDbContext(db_name, 15*time.Second)
	defer cancel()
	var mm []OrderObject
	if err != nil {
		appCtx.Error(err)
		return nil, paginate.PaginationData{}, true
	}
	log.Println("I am here")

	// fil := bson.M{
	// 	"uploadId":            "124",
	// 	"orderDetails.status": "NEW",
	// }

	projection := bson.M{}

	collection := db.Collection(collection_name)
	paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Select(projection).Filter(filters).Decode(&mm).Find()
	if err != nil {
		log.Println("Error getting paginated data: ", err)
		return nil, paginate.PaginationData{}, true
	}
	//log.Println("Order data: ", mm)
	if len(mm) == 0 {
		return nil, paginate.PaginationData{}, true
	}
	return mm, paginatedData.Pagination, false

}

func FetchOrdersFromDB(uploadId string, page int64, limit int64) ([]OrderObject, paginate.PaginationData, bool) {
	var filters = make(map[string]string)
	filters["orderDetails.status"] = "NEW"
	filters["uploadId"] = uploadId
	log.Println("filters: ", filters)
	return FindBy(filters, page, limit)

}
