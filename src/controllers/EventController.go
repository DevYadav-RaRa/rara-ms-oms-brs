package controllers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/services"
	routing "github.com/qiangxue/fasthttp-routing"
)

type EventController struct {
	AppCtx framework.Framework
}

func (c *EventController) Track(ctx *routing.Context) error {
	var resp services.EventInsertResponse
	var indexName = strings.TrimSpace(string(ctx.QueryArgs().Peek("index")))
	var eventType = strings.TrimSpace(string(ctx.QueryArgs().Peek("type")))
	var eventSource = strings.TrimSpace(string(ctx.QueryArgs().Peek("source")))
	var encodedPayload = strings.TrimSpace(string(ctx.QueryArgs().Peek("payload")))
	decodedValue, errEncode := url.QueryUnescape(encodedPayload)

	if errEncode != nil {
		resp = services.EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     "Error while decoding",
		}
	}
	var payload = decodedValue

	if len(indexName) == 0 {
		resp = services.EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     "Please provide the index name",
		}
	} else if len(eventType) == 0 {
		resp = services.EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     "Please provide the event type",
		}
	} else if len(eventSource) == 0 {
		resp = services.EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     "Please provide the event source",
		}
	}

	var data map[string]interface{} = make(map[string]interface{})
	err := json.Unmarshal([]byte(payload), &data)

	if err != nil {
		data = map[string]interface{}{"data": payload}
	}
	switch eventType + "|" + indexName {
	case "received|notifications", "received|testing":
		resp = services.TrackReceivedStatus(c.AppCtx, indexName+"_status_map", eventType, data)
		break
	case "clicked|notifications", "clicked|testing":
		resp = services.TrackClickedStatus(c.AppCtx, indexName+"_status_map", eventType, data)
		break
	default:
		resp = services.InsertEventInIndex(c.AppCtx, indexName, eventType, eventSource, data)
		break
	}

	return buildJsonResponse(resp, ctx)

}

func (c *EventController) TrackPost(ctx *routing.Context) error {
	var resp services.EventInsertResponse
	var body = strings.TrimSpace(string(ctx.PostBody()))
	var data map[string]interface{} = make(map[string]interface{})
	err := json.Unmarshal([]byte(body), &data)

	if err != nil {
		resp = services.EventInsertResponse{
			Index:     "body",
			EventType: "",
			Status:    false,
			Error:     "Please provide body",
		}
	}

	var indexName, indexErr = data["index"].(string)
	var eventType, eventTypeErr = data["type"].(string)
	var eventSource, eventSourceErr = data["source"].(string)
	var payload, payloadErr = data["payload"].(map[string]interface{})

	if !indexErr {
		resp = services.EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     "Please provide the index name",
		}
	} else if !eventTypeErr {
		resp = services.EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     "Please provide the event type",
		}
	} else if !eventSourceErr {
		resp = services.EventInsertResponse{
			Index:     indexName,
			EventType: eventType,
			Status:    false,
			Error:     "Please provide the event source",
		}
	}

	if !payloadErr {
		data = map[string]interface{}{"data": payload}
	} else {
		data = payload
	}
	switch eventType + "|" + indexName {
	case "received|notifications", "received|testing":
		resp = services.TrackReceivedStatus(c.AppCtx, indexName+"_status_map", eventType, data)
		break
	case "clicked|notifications", "clicked|testing":
		resp = services.TrackClickedStatus(c.AppCtx, indexName+"_status_map", eventType, data)
		break
	default:
		resp = services.InsertEventInIndex(c.AppCtx, indexName, eventType, eventSource, data)
		break
	}

	return buildJsonResponse(resp, ctx)
}

// private
func buildJsonResponse(resp services.EventInsertResponse, ctx *routing.Context) error {
	bytes, err := json.Marshal(resp)
	if err != nil {
		ctx.SetStatusCode(500)
		return err
	}
	if resp.Status == false {
		ctx.SetStatusCode(400)
	} else {
		ctx.SetStatusCode(201)
	}
	ctx.SetContentType("application/json")
	fmt.Fprintf(ctx, "%+v", string(bytes))
	return nil
}
