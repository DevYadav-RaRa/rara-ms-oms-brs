package restApis

import (
	"errors"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/conf"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/helpers"
	routing "github.com/qiangxue/fasthttp-routing"
)

func InitApis(appCtx framework.Framework) {

	appCtx.Router.Post("/orders", func(c *routing.Context) error {
		conf.ConsumeApiOrders(string(c.PostBody()))
		return errors.New(string("Done"))
	})

	appCtx.Router.Post("/orders/csv", func(c *routing.Context) error {
		return helpers.CsvMiddleware(c)
	})
}

// Code for checking csv file

// v, _ := c.FormFile("file")
// fmt.Println(v.Filename, v, e)
// u, _ := v.Open()
// csvLines, err := csv.NewReader(u).ReadAll()
// if err != nil {
// 	fmt.Println(err)
// }
// for i, line := range csvLines {
// 	emp := map[string]string{
// 		"Name":         line[0],
// 		"MicroService": line[1],
// 		"Progress":     line[2],
// 	}
// 	fmt.Println(i, ": ", emp["Name"]+" "+emp["MicroService"]+" "+emp["Progress"])
// }
