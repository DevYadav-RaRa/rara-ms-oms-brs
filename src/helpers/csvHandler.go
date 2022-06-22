package helpers

import (
	"errors"
	"fmt"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/aws/s3"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
	routing "github.com/qiangxue/fasthttp-routing"
)

// type headers struct {
// 	TenantToken  string `json:"Tenanttoken" bson:"Tenanttoken"`
// 	BusinessName string `json:"Businessname" bson:"Businessname"`
// 	AccNo        string `json:"Accountnumber" bson:"Accountnumber"`
// 	ServiceType  string `json:"Servicetype" bson:"Servicetype"`
// 	ServiceId    string `json:"Serviceid" bson:"Serviceid"`
// }

func CsvMiddleware(c *routing.Context) error {

	framework.Logs("Calling Iam for Authentication")
	framework.Logs(c.Request.Header.String())
	// header := headers{}
	// err := json.Unmarshal([]byte(c.Request.Header.String()), &header)
	// if err != nil {
	// 	framework.Logs(err.Error())
	// 	return err
	// }

	req := models.IamRequest{}
	req.TenantToken = string(c.Request.Header.Peek("Tenanttoken"))
	req.BusinessDetails.AccNo = string(c.Request.Header.Peek("Accountnumber"))
	req.BusinessDetails.BusinessName = string(c.Request.Header.Peek("Businessname"))
	req.BusinessDetails.ServiceId = string(c.Request.Header.Peek("Serviceid"))
	req.BusinessDetails.ServiceType = string(c.Request.Header.Peek("Servicetype"))

	fmt.Println(req)

	IamAuth := req.GetIamAuthentication("BusinessHeader")
	framework.Logs("Iam Response: ")
	fmt.Println(IamAuth)
	if !IamAuth.Status {
		return errors.New("rejected from iam")
	}
	framework.Logs("Authenticated from Iam")

	return errors.New(s3.GetPresignedUrl())
}
