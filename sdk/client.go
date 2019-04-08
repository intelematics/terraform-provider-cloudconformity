package sdk

import (
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

func (client Client) getUrl(path string) string {
	return fmt.Sprintf("https://%s-api.cloudconformity.com/v1/%s", client.region, path)
}

func (client Client) addHeaders(request *http.Request) {
	request.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("ApiKey %s", client.apiKey)},
		"Content-Type":  {"application/vnd.api+json"},
	}
}

func (client Client) DeleteAccount(id string) error {
	return nil
}

func NewClient(apiKey string, region string) *Client {
	return &Client{apiKey: apiKey, region: region}
}
