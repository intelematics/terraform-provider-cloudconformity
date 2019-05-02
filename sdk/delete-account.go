package sdk

import "fmt"

func (client Client) DeleteAccount(id string) error {
	return client.genericDelete(fmt.Sprintf("accounts/%s", id))
}
