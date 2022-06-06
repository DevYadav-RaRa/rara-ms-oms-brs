package conf

import (
	"fmt"
	"io/ioutil"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/helpers"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
)

func Testing() {
	b, err := ioutil.ReadFile("./src/conf/api.json")
	if err != nil {
		fmt.Print("Bootstrap error: ", err)
		return
	}
	// fmt.Println(string(b))
	var demoApi models.ApiPayload
	demoApi.FromJSONString(string(b))
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------------------------------------")
	for i := range demoApi.Orders {
		var temp models.OrderObject
		temp.TenantToken = demoApi.TenantToken
		temp.BusinessDetails = demoApi.BusinessDetails
		temp.Order = demoApi.Orders[i]
		status, resp := helpers.PostOrder(temp, "Business Header")
		fmt.Println(status, " :: ", resp)
		fmt.Println("-------------------------------------------")
		fmt.Println("-------------------------------------------")
	}
}

func Bootstrap(appCtx framework.Framework) {
	fmt.Println("Running Bootstrap...")
	Testing()
	fmt.Println("App is ready!")
}
