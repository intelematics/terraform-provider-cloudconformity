package sdk

import "fmt"

type updateAccountSubscriptionAttributes struct {
	CostPackage bool `json:"costPackage"`
	RealTimeMonitoring bool `json:"hasRealTimeMonitoring"`
	SecurityPackage bool `json:"securityPackage"`
}

func (client Client) UpdateAccountSubscription(id string, cost bool, realtime bool, security bool) error {
	requestPayload := struct {
		Data updateAccountSubscriptionAttributes `json:"data"`
	}{}
	requestPayload.Data.CostPackage = cost
	requestPayload.Data.RealTimeMonitoring = realtime
	requestPayload.Data.SecurityPackage = security

	return client.genericPatch(fmt.Sprintf("/accounts/%s/subscription", id), &requestPayload)
}