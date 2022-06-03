package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventInsertResponse struct {
	Index     string      `json:"index"`
	EventType string      `json:"type"`
	Status    bool        `json:"status"`
	Error     string      `json:"error"`
	Id        interface{} `json:"id"`
}

func InsertEventInIndex(appCtx framework.Framework, indexName string, eventType string, eventSource string, payload map[string]interface{}) EventInsertResponse {
	if _, ok := appCtx.Mongo["events"]; !ok {
		return EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     fmt.Sprintf("%+v", errors.New("database not initialized")),
		}
	}
	var connection framework.MongoConnection = appCtx.Mongo["events"]
	if !connection.Connected() {
		return EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     fmt.Sprintf("%+v", errors.New("database not connected")),
		}
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	res, err := connection.DB().Collection(indexName).InsertOne(ctx, bson.M{
		"eventType": eventType,
		"source":    eventSource,
		"payload":   payload,
		"ts":        time.Now().Unix(),
	})

	if err != nil {
		return EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     fmt.Sprintf("%+v", err),
		}
	}
	return EventInsertResponse{
		Index:     indexName,
		EventType: eventType,
		Status:    true,
		Error:     "",
		Id:        res.InsertedID,
	}
}

func TrackReceivedStatus(appCtx framework.Framework, indexName string, eventType string, payload map[string]interface{}) EventInsertResponse {
	fmt.Println(payload)
	if _, ok := appCtx.Mongo["events"]; !ok {
		return EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     fmt.Sprintf("%+v", errors.New("database not initialized")),
		}
	}
	var connection framework.MongoConnection = appCtx.Mongo["events"]
	if !connection.Connected() {
		return EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     fmt.Sprintf("%+v", errors.New("database not connected")),
		}
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	//var eventMetadata map[string]interface{} = make(map[string]interface{})
	//err := json.Unmarshal([]byte(payload["eventMetadata"].(string)), &eventMetadata)

	if payload["parentNotificationId"] == nil {
		return EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     fmt.Sprintf("%+v", errors.New("parentNotificationId is required")),
		}
	}

	payload["ts"] = time.Now().Unix()
	payload["clicked"] = false

	res, err := connection.DB().Collection(indexName).InsertOne(ctx, bson.M(payload))

	if err != nil {
		fmt.Println(err)
		return EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     fmt.Sprintf("%+v", err),
		}
	}
	return EventInsertResponse{
		Index:     indexName,
		EventType: eventType,
		Status:    true,
		Error:     "",
		Id:        res.InsertedID,
	}
}

func TrackClickedStatus(appCtx framework.Framework, indexName string, eventType string, payload map[string]interface{}) EventInsertResponse {
	if _, ok := appCtx.Mongo["events"]; !ok {
		return EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     fmt.Sprintf("%+v", errors.New("database not initialized")),
		}
	}
	var connection framework.MongoConnection = appCtx.Mongo["events"]
	if !connection.Connected() {
		return EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     fmt.Sprintf("%+v", errors.New("database not connected")),
		}
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if payload["parentNotificationId"] == nil {
		return EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     fmt.Sprintf("%+v", errors.New("parentNotificationId is required")),
		}
	}

	filter := bson.M{
		"parentNotificationId": payload["parentNotificationId"],
		"userId":               payload["userId"],
		"orgId":                payload["orgId"],
	}

	payload["ts"] = time.Now().Unix()
	updateQuery := bson.M{
		"$set": bson.M{
			"clicked":   true,
			"clickedTs": time.Now().Unix(),
		},
		"$setOnInsert": bson.M(payload),
	}

	opts := options.Update().SetUpsert(true)

	res, err := connection.DB().Collection(indexName).UpdateOne(ctx, filter, updateQuery, opts)

	if err != nil {
		return EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     fmt.Sprintf("%+v", err),
		}
	}
	return EventInsertResponse{
		Index:     indexName,
		EventType: eventType,
		Status:    true,
		Error:     "",
		Id:        res,
	}
}
