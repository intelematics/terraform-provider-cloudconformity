package sdk

import "fmt"

type getAccountAccessSettingsAttributes struct {
	Configuration getAccountAccessSettingsConfiguration `json:"configuration"`
}

type getAccountAccessSettingsConfiguration struct {
	ExternalId string `json:"externalId"`
	RoleArn    string `json:"roleArn"`
}

func (client Client) GetAccountAccessSettings(id string) (*AccountAccessSettings, error) {
	results := struct {
		Attributes getAccountAccessSettingsAttributes `json:"attributes"`
	}{}

	err := client.genericGet(fmt.Sprintf("accounts/%s/access", id), &results)
	if err != nil {
		return nil, err
	}

	result := AccountAccessSettings{
		ExternalId: results.Attributes.Configuration.ExternalId,
		RoleArn:    results.Attributes.Configuration.RoleArn,
	}

	return &result, nil
}
