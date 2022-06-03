package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
	routing "github.com/qiangxue/fasthttp-routing"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	appCtx := framework.GetCurrentAppContext()

	appCtx.Router.Get("/OrderObject/accnt-<id:[^/]*>", func(c *routing.Context) error {

		x := models.OrderObject{}
		var err error
		x.OrderDetails.OrderId, err = primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			fmt.Fprintf(c, "Not a valid id")
		}
		b, err := json.Marshal(&x)
		fmt.Fprintf(c, string(b))
		return nil
	})

	appCtx.Router.Post("/OrderObject/accnt-<object:[^/]*>", func(c *routing.Context) error {

		in := []byte(c.Param("id"))

		var object1 models.OrderObject
		err := json.Unmarshal(in, &object1)
		if err != nil {
			fmt.Fprintf(c, "Not a valid object")
		}
		object1.Save()
		return nil
	})
}
