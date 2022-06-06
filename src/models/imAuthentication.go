package models

type ImResponse struct {
	Status             bool
	BusinessValidation BusinessValidation
	Webhook            Webhook
}

type BusinessValidation struct {
	Success string
	Failure string
}

type Request struct {
	TenantToken     string
	BusinessDetails BusinessDetails
}

// ONLY FOR TESTING PURPOSES

func (requestObject *Request) GetImResponse(header string) ImResponse {
	var respObj ImResponse
	respObj.BusinessValidation.Success = "successUrl"
	respObj.BusinessValidation.Failure = "FailureUrl"
	respObj.Webhook.Header.header = header
	respObj.Webhook.Url = "webhookUrl"
	respObj.Webhook.Payload = "payloadString"

	// now := time.Now()
	// nowNano := now.Nanosecond()
	// fmt.Println(nowNano)
	// respObj.Status = now.UnixNano()%2 == 0

	respObj.Status = true

	return respObj
}

// ONLY FOR TESTING PURPOSES
