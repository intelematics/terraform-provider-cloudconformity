package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (client Client) getUrl(path string) string {
	return fmt.Sprintf("https://%s-api.cloudconformity.com/v1/%s", client.region, path)
}

func (client Client) addHeaders(request *http.Request) {
	request.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("ApiKey %s", client.apiKey)},
		"Content-Type":  {"application/vnd.api+json"},
	}
}

func (client Client) genericGet(path string, myOutput interface{}) error {
	req, err := http.NewRequest("GET", client.getUrl(path), nil)
	if err != nil {
		return err
	}
	client.addHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return errors.New(http.StatusText(resp.StatusCode))
	}
	// https://old.reddit.com/r/golang/comments/3735so/do_we_have_to_check_for_errors_when_we_call_close/
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, myOutput)
	if err != nil {
		return err
	}

	return nil
}
