package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

var errNotFound = errors.New("entity not found")

func (client Client) getUrl(path string) string {
	return fmt.Sprintf("https://%s-api.cloudconformity.com/v1/%s", client.region, path)
}

func (client Client) addHeaders(request *http.Request) {
	request.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("ApiKey %s", client.apiKey)},
		"Content-Type":  {"application/vnd.api+json"},
	}
}

func (client Client) genericPatch(path string, myInput interface{}) error {
	return client.genericRequest("PATCH", path, myInput, nil)
}

func (client Client) genericGet(path string, myOutput interface{}) error {
	return client.genericRequest("GET", path, nil, myOutput)
}

func (client Client) genericDelete(path string) error {
	return client.genericRequest("DELETE", path, nil, nil)
}

func (client Client) genericPost(path string, input interface{}, output interface{}) error {
	return client.genericRequest("POST", path, input, output)
}

func (client Client) genericRequest(requestType string, path string, input interface{}, output interface{}) error {
	var req *http.Request
	var err error
	if input != nil {
		var payload []byte
		payload, err = json.Marshal(input)
		if err != nil {
			return err
		}
		reader := bytes.NewReader(payload)
		req, err = http.NewRequest(requestType, client.getUrl(path), reader)
	} else {
		req, err = http.NewRequest(requestType, client.getUrl(path), nil)
	}

	if err != nil {
		return err
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
		return err
	}

	// Save a copy of this response for debugging.
	responseBody, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(responseBody))

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		if resp.StatusCode == 404 {
			return errNotFound
		}
		return errors.New(http.StatusText(resp.StatusCode))
	}

	if output != nil {
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

		err = json.Unmarshal(body, output)
		if err != nil {
			return err
		}
	}

	return nil
}
