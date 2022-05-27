package conf

import "github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"

var MongoConnections []framework.MongoConnectionDescription = []framework.MongoConnectionDescription{
	{
		Name:        "bms",
		EnvVarName:  "MONGO_URL_BMS",
		Description: "Connects to 'bms' Mongo Database.",
		CanFail:     true,
	},
}
