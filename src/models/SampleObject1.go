package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
// SampleObject1_DB define db in which model exists
const SampleObject1_DB = "bms"


// // SampleObject1_NAME name of collection in mongo

const SampleObject1_COLLECTION_NAME = "batches"
type SampleObject1 struct{
	Id primitive.ObjectID
	Sample_field_1 string
	Sample_filed_2 float64
	Sample_filed_3 []SampleObject2
}
type SampleObject2 struct{
	Sample_field_4 string
	Sample_filed_5 int
	Sample_filed_6 []float64
}



func (obj *SampleObject1) Insert() (error, bool) {
		err, ctx, appCtx, db, cancel := getDbContext(SampleObject1_DB, 10*time.Second)
		defer cancel()
		if err != nil {
			appCtx.Error(err)
			return err, false
		}
		// // as new doc, id is needed
		obj.Id = primitive.NewObjectID()
		
		doc, err := toDoc(obj)
		if err != nil {
			appCtx.Error(err)
			return err, false
		}
		res, err := db.Collection(SampleObject1_COLLECTION_NAME).InsertOne(ctx, doc)
		if err != nil {
			appCtx.Error(err)
			return err, false
		}
	
		obj.Id = res.InsertedID.(primitive.ObjectID)
	
		return nil, true
	}


func (obj *SampleObject1) FindById() (interface{}, bool) {
	err, ctx, appCtx, db, cancel := getDbContext(SampleObject1_DB, 15*time.Second)
	defer cancel()
	if err != nil {
		appCtx.Error(err)
		return nil, false
	}

	var data map[string]interface{} = make(map[string]interface{})
	res := db.Collection(SampleObject1_COLLECTION_NAME).FindOne(ctx, bson.M{"_id": obj.Id})
	res.Decode(&data)
	return data, true
}

//

//

func (obj *SampleObject1) Save() (bool, error) {
	err, ctx, appCtx, db, cancel := getDbContext(SampleObject1_DB, 5*time.Second)

	defer cancel()
	doc, err := toDoc(obj)
	if err != nil {
		appCtx.Error(err)
		return false, err
	}
	res, err := db.Collection(SampleObject1_COLLECTION_NAME).UpdateOne(ctx, bson.M{"_id": obj.Id}, bson.M{"$set": doc})
	if err != nil {
		appCtx.Error(err)
		return false, err
	}
	fmt.Println(res.ModifiedCount)
	return true, err
}

func (obj *SampleObject1) FromJSONString(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), obj)
	if err != nil {
		framework.GetCurrentAppContext().Error(err)
		return err
	}
	return nil
}
