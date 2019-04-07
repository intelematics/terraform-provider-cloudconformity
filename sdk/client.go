package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

type accountData struct {
	Id         string            `json:"id"`
	Attributes accountAttributes `json:"attributes"`
}

type accountAttributes struct {
	Name string `json:"name"`
}

func (client Client) GetAccount(id string) (*Account, error) {

	req, err := http.NewRequest("GET", client.getUrl(fmt.Sprintf("accounts/%s", id)), nil)
	if err != nil {
		return nil, err
	}
	client.addHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}
	// https://old.reddit.com/r/golang/comments/3735so/do_we_have_to_check_for_errors_when_we_call_close/
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getAccounts := struct {
		Data accountData `json:"data"`
	}{}

	err = json.Unmarshal(body, &getAccounts)
	if err != nil {
		return nil, err
	}

	account := Account{
		Name: getAccounts.Data.Attributes.Name,
		Id:   getAccounts.Data.Id,
	}

	return &account, err
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

func NewClient(apiKey string, region string) *Client {
	return &Client{apiKey: apiKey, region: region}
}
