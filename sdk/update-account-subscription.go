package sdk

import "fmt"

type updateAccountSubscriptionAttributes struct {
	CostPackage        bool `json:"costPackage"`
	RealTimeMonitoring bool `json:"hasRealTimeMonitoring"`
	SecurityPackage    bool `json:"securityPackage"`
}

type updateAccountSubscriptionData struct {
	Attributes updateAccountSubscriptionAttributes `json:"attributes"`
}

func (client Client) UpdateAccountSubscription(id string, cost bool, realtime bool, security bool) error {
	requestPayload := struct {
		Data updateAccountSubscriptionData `json:"data"`
	}{}
	requestPayload.Data.Attributes.CostPackage = cost
	requestPayload.Data.Attributes.RealTimeMonitoring = realtime
	requestPayload.Data.Attributes.SecurityPackage = security

	return client.genericPatch(fmt.Sprintf("/accounts/%s/subscription", id), &requestPayload)
}
