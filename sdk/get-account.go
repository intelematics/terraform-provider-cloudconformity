package sdk

import (
	"fmt"
)

type accountData struct {
	Id         string            `json:"id"`
	Attributes accountAttributes `json:"attributes"`
}

type accountAttributes struct {
	Name                  string `json:"name"`
	Environment           string `json:"environment"`
	HasRealTimeMonitoring bool   `json:"has-real-time-monitoring"`
	SecurityPackage       bool   `json:"security-package"`
	CostPackage           bool   `json:"cost-package"`
}

func (client Client) GetAccount(id string) (*Account, error) {

	getAccounts := struct {
		Data accountData `json:"data"`
	}{}

	err := client.genericGet(fmt.Sprintf("accounts/%s", id), &getAccounts)

	account := Account{
		Name:                  getAccounts.Data.Attributes.Name,
		Id:                    getAccounts.Data.Id,
		Environment:           getAccounts.Data.Attributes.Environment,
		HasRealTimeMonitoring: getAccounts.Data.Attributes.HasRealTimeMonitoring,
		SecurityPackage:       getAccounts.Data.Attributes.SecurityPackage,
		CostPackage:           getAccounts.Data.Attributes.CostPackage,
	}

	return &account, err
}

func (client Client) DoesAccountExist(id string) (bool, error) {
	account, err := client.GetAccount(id)

	result := account != nil

	resultErr := err
	if err == errNotFound {
		resultErr = nil
	}

	return result, resultErr
}