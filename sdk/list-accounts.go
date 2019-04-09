package sdk

type listAccountAttributes struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
}

type listAccountOverview struct {
	Id         string                `json:"id"`
	Attributes listAccountAttributes `json:"attributes"`
}

func (client Client) ListAccounts() ([]AccountOverview, error) {
	listAccounts := struct {
		Data []listAccountOverview `json:"data"`
	}{}

	err := client.genericGet("accounts", &listAccounts)
	if err != nil {
		return nil, err
	}

	result := make([]AccountOverview, len(listAccounts.Data))
	for _, overview := range listAccounts.Data {
		result = append(result, AccountOverview{
			Id:          overview.Id,
			Environment: overview.Attributes.Environment,
			Name:        overview.Attributes.Name,
		})
	}

	return result, nil
}
