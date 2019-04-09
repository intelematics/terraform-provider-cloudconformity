package sdk

type getExternalIdData struct {
	Id string `json:"id"`
}

func (client Client) GetExternalId() (string, error) {
	getExternalId := struct {
		Data getExternalIdData `json:"data"`
	}{}

	err := client.genericGet("organisation/external-id", &getExternalId)
	if err != nil {
		return "", err
	}

	return getExternalId.Data.Id, nil
}
