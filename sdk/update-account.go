package sdk

import "fmt"

type updateAccountAttributes struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
}

type updateAccountData struct {
	Attributes updateAccountAttributes `json:"attributes"`
}

func (client Client) UpdateAccount(id string, name string, environment string) error {
	requestPayload := struct {
		Data updateAccountData `json:"data"`
	}{}
	requestPayload.Data.Attributes.Name = name
	requestPayload.Data.Attributes.Environment = environment

	return client.genericPatch(fmt.Sprintf("/accounts/%s", id), &requestPayload)
}
