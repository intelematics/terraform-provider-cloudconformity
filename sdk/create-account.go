package sdk

import (
	"fmt"
	"time"
)

type createAccountAccessKeys struct {
	RoleArn    string `json:"roleArn"`
	ExternalId string `json:"externalId"`
}

type createAccountAccess struct {
	Keys createAccountAccessKeys `json:"keys"`
}

type createAccountAttributes struct {
	Name                  string              `json:"name"`
	Environment           string              `json:"environment"`
	Access                createAccountAccess `json:"access"`
	CostPackage           bool                `json:"costPackage"`
	HasRealTimeMonitoring bool                `json:"hasRealTimeMonitoring"`
	SecurityPackage       bool                `json:"securityPackage"`
}

type createAccountData struct {
	Type        string                  `json:"type"`
	Attributres createAccountAttributes `json:"attributes"`
}

type createAccountResponse struct {
	Id string `json:"id"`
}

func (client Client) CreateAccount(request CreateAccountRequest) (string, error) {
	payload := struct {
		Data createAccountData `json:"data"`
	}{}
	payload.Data.Type = "account"
	payload.Data.Attributres.Access.Keys.ExternalId = request.ExternalId
	payload.Data.Attributres.Access.Keys.RoleArn = request.Role
	payload.Data.Attributres.Environment = request.Environment
	payload.Data.Attributres.Name = request.Name
	payload.Data.Attributres.CostPackage = request.CostPackage
	payload.Data.Attributres.SecurityPackage = request.SecurityPackage
	payload.Data.Attributres.HasRealTimeMonitoring = request.HasRealTimeMonitoring

	output := struct {
		Data createAccountResponse `json:"data"`
	}{}

	retries := request.Retries

	for {
		fmt.Printf("Trying to create Account %d",retries)
		err := client.genericPost("accounts", &payload, &output)
		if err == nil {
			break
		}
		if err != nil && retries == 0 {
			return "", err
		}
		retries -= 1
		time.Sleep(2 * time.Second)
	}

	return output.Data.Id, nil
}
