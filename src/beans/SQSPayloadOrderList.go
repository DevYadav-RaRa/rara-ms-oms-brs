package beans

import (
	"encoding/json"
	"fmt"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
)

type PayloadOrderList struct {
	Type   string  `json:"type"`
	Status string  `json:"status"`
	Orders []Order `json:"orders"`
}

type Order struct {
	OrderDetails struct {
		Id        string  `json:"id"`
		CityId    string  `json:"cityId"`
		Type      string  `json:"type"`
		Sla       float64 `json:"sla"`
		Size      float64 `json:"size"`
		Weight    string  `json:"weight"`
		Kecemetan struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"kecemetan"`
	} `json:"orderDetails"`
	OrderGeo struct {
		Lat     string `json:"lat"`
		Lng     string `json:"lng"`
		PinCode string `json:"pinCode"`
		GeoHash string `json:"geoHash"`
	} `json:"orderGeo"`
	OriginDetails struct {
		Type          string `json:"type"`
		Id            string `json:"id"`
		CapacityIndex string `json:"capacityIndex"`
		WeightIndex   string `json:"weightIndex"`
		Lat           string `json:"lat"`
		Lng           string `json:"lng"`
	} `json:"originDetails"`
}

func (payloadOrderList *PayloadOrderList) FromJSONString(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), payloadOrderList)
	fmt.Println("CHECKING")
	if err != nil {
		framework.GetCurrentAppContext().Error(err)
		return err
	}

	return nil
}
