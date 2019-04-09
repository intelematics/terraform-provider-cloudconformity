package sdk

import (
	"fmt"
)

type CreateAccountRequest struct {
	Name                  string
	Environment           string
	Role                  string
	ExternalId            string
	HasRealTimeMonitoring bool
	CostPackage           bool
}

type accountData struct {
	Id         string            `json:"id"`
	Attributes accountAttributes `json:"attributes"`
}

type accountAttributes struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
}

func (client Client) GetAccount(id string) (*Account, error) {

	getAccounts := struct {
		Data accountData `json:"data"`
	}{}

	err := client.genericGet(fmt.Sprintf("accounts/%s", id), &getAccounts)

	account := Account{
		Name:        getAccounts.Data.Attributes.Name,
		Id:          getAccounts.Data.Id,
		Environment: getAccounts.Data.Attributes.Environment,
	}

	return &account, err
}
