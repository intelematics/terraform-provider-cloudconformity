package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)


type CreateAccountRequest struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
	Role        string `json:"role"`
	ExternalId  string `json:"external-id"`
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
