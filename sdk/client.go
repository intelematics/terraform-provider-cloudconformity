package sdk

type Client struct {
	apiKey string
	region string
}

type Account struct {
	Id          string
	Name        string
	Environment string
	RoleArn     string
	ExternalId  string
}

type AccountOverview struct {
	Id          string
	Name        string
	Environment string
}

func (client Client) DeleteAccount(id string) error {
	return nil
}

func NewClient(apiKey string, region string) *Client {
	return &Client{apiKey: apiKey, region: region}
}
