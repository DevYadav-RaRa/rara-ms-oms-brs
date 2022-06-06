package conf

import "github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"

var MongoConnections []framework.MongoConnectionDescription = []framework.MongoConnectionDescription{
	{
		Name:        "oms",
		EnvVarName:  "MONGO_URL_BMS",
		Description: "Connects to 'oms' Mongo Database.",
		CanFail:     true,
	},
}
