package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type createAccountAccessKeys struct {
	RoleArn    string `json:"roleArn"`
	ExternalId string `json:"externalId"`
}

type createAccountAccess struct {
	Keys createAccountAccessKeys `json:"keys"`
}

type createAccountAttributes struct {
	Name        string              `json:"name"`
	Environment string              `json:"environment"`
	Access      createAccountAccess `json:"access"`
	CostPackage bool `json:"costPackage"`
	HasRealTimeMonitoring bool `json:"hasRealTimeMonitoring"`
}

type createAccountData struct {
	Type        string                  `json:"type"`
	Attributres createAccountAttributes `json:"attributes"`
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

	// Save a copy of this request for debugging.
	requestReq, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestReq))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Save a copy of this response for debugging.
	requestDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Unable to create account: %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode)))
	}

	// https://old.reddit.com/r/golang/comments/3735so/do_we_have_to_check_for_errors_when_we_call_close/
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	return nil, nil
}
