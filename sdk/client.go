package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	apiKey string
	region string
}

type Account struct {
	Name string
	Id   string
}

type createAccountAccessKeys struct {
	RoleArn string `json:"roleArn"`
	ExternalId string `json:"externalId"`
}

type createAccountAccess struct {
	Keys createAccountAccessKeys `json:"keys"`
}

type createAccountAttributes struct {
	Name string `json:"name"`
	Environment string `json:"environment"`
	Access createAccountAccess `json:"access"`
}

type createAccountData struct {
	Type string `json:"type"`
	Attributres createAccountAttributes `json:"attributes"`
}

func (client Client) getUrl(path string) string {
	return fmt.Sprintf("https://%s-api.cloudconformity.com/v1/%s", client.region, path)
}

func (client Client) addHeaders(request *http.Request) {
	request.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("ApiKey %s", client.apiKey)},
		"Content-Type":  {"application/vnd.api+json"},
	}
}

func (client Client) CreateAccount(request CreateAccountRequest) (*Account, error) {
	createAccountRequest := struct {
		Data createAccountData `json:"data"`
	}{}
	createAccountRequest.Data.Type = "account"
	createAccountRequest.Data.Attributres.Access.Keys.ExternalId = request.ExternalId
	createAccountRequest.Data.Attributres.Access.Keys.RoleArn = request.Role
	createAccountRequest.Data.Attributres.Environment = request.Environment
	createAccountRequest.Data.Attributres.Name = request.Name

	requestBody, err := json.Marshal(&createAccountRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", client.getUrl("accounts"), bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	client.addHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// https://old.reddit.com/r/golang/comments/3735so/do_we_have_to_check_for_errors_when_we_call_close/
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	return nil, nil
}

func NewClient(apiKey string, region string) *Client {
	return &Client{apiKey: apiKey, region: region}
}
