package sdk

import "fmt"

type getAccountAccessSettingsConfiguration struct {
	ExternalId string `json:"externalId"`
	RoleArn    string `json:"roleArn"`
}

func (client Client) GetAccountAccessSettings(id string) (*AccountAccessSettings, error) {
	results := struct {
		Attributes getAccountAccessSettingsConfiguration `json:"attributes"`
	}{}

	err := client.genericGet(fmt.Sprintf("accounts/%s/access", id), &results)
	if err != nil {
		return nil, err
	}

	result := AccountAccessSettings{
		ExternalId: results.Attributes.ExternalId,
		RoleArn:    results.Attributes.RoleArn,
	}

	return &result, nil
}
