package restApis

import (
	"errors"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/conf"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	routing "github.com/qiangxue/fasthttp-routing"
)

func InitApis(appCtx framework.Framework) {

	appCtx.Router.Post("/", func(c *routing.Context) error {
		resp, _ := conf.ConsumeApiOrders(string(c.PostBody()))
		return errors.New(resp)
	})
}
