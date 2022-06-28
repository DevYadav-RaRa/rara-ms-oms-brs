package models

import (
	"encoding/json"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
)

type ApiPayload struct {
	TenantToken     string          `json:"tenantToken" bson:"tenantToken"`
	BusinessDetails BusinessDetails `json:"businessDetails" bson:"businessDetails"`
	Orders          []Order         `json:"orders" bson:"orders"`
}

func (obj *OrderObject) FromJSONString(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), obj)
	if err != nil {
		framework.GetCurrentAppContext().Error(err)
		return err
	}
	return nil
}

func (obj *ApiPayload) FromJSONString(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), obj)
	if err != nil {
		framework.GetCurrentAppContext().Error(err)
		return err
	}
	return nil
}
