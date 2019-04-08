package sdk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type getExternalIdData struct {
	Id string `json:"id"`
}

func (client Client) GetExternalId() (string, error) {
	req, err := http.NewRequest("GET", client.getUrl("organisation/external-id"), nil)

	if err != nil {
		return "", err
	}

	client.addHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return "", errors.New(http.StatusText(resp.StatusCode))
	}

	// https://old.reddit.com/r/golang/comments/3735so/do_we_have_to_check_for_errors_when_we_call_close/
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	getExternalId := struct {
		Data getExternalIdData `json:"data"`
	}{}

	err = json.Unmarshal(body, &getExternalId)
	if err != nil {
		return "", err
	}

	return getExternalId.Data.Id, nil
}
