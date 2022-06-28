package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Common method
func getDbContext(db string, timeout time.Duration) (error, context.Context, *framework.Framework, *mongo.Database, context.CancelFunc) {
	appCtx := framework.GetCurrentAppContext()
	// fmt.Println(appCtx.Mongo, "errrrrr check ")
	if _, ok := appCtx.Mongo[db]; !ok {
		return errors.New(fmt.Sprintf("%+v", errors.New("database not initialized"))), nil, appCtx, nil, nil
	}
	var connection framework.MongoConnection = appCtx.Mongo[db]
	if !connection.Connected() {
		return errors.New(fmt.Sprintf("%+v", errors.New("database not connected"))), nil, appCtx, nil, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return nil, ctx, appCtx, connection.DB(), cancel
}

func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal([]byte(data), &doc)
	// fmt.Println(doc, "here check")
	return
}
