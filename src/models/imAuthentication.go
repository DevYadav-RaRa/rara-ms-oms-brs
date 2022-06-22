package models

type IamResponse struct {
	Status             bool               `json:"status" bson:"status"`
	BusinessValidation BusinessValidation `json:"businessValidation" bson:"businessValidation"`
	Webhook            []Webhook          `json:"webhook" bson:"webhook"`
}

type BusinessValidation struct {
	Success string `json:"success" bson:"success"`
	Failure string `json:"failure" bson:"failure"`
}

type IamRequest struct {
	TenantToken     string          `json:"tenantToken" bson:"tenantToken"`
	BusinessDetails BusinessDetails `json:"businessDetails" bson:"businessDetails"`
}

type Webhook struct {
	Purpose       string    `json:"purpose" bson:"purpose"`
	RequestMethod string    `json:"requestMethod" bson:"requestMethod"`
	Url           string    `json:"url" bson:"url"`
	Headers       []Headers `json:"headers" bson:"headers"`
	Payload       string    `json:"payload" bson:"payload"`
}

type Headers struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}

// ONLY FOR TESTING PURPOSES

func (requestObject *IamRequest) GetIamResponse(header string) IamResponse {
	var respObj IamResponse
	respObj.BusinessValidation.Success = "SuccessUrl"
	respObj.BusinessValidation.Failure = "FailureUrl"

	var wbhk Webhook
	wbhk.Purpose = "OrderUpdate"
	wbhk.RequestMethod = "POST"
	wbhk.Headers = append(wbhk.Headers, Headers{Key: "Key1", Value: header}, Headers{Key: "Key2", Value: header})
	wbhk.Url = "https://endh5z9mbnve4.x.pipedream.net/"
	wbhk.Payload = "Tmpl_01"

	respObj.Webhook = append(respObj.Webhook, wbhk)

	respObj.Status = true

	return respObj
}

// ONLY FOR TESTING PURPOSES
