package sdk

type Client struct {
	apiKey string
	region string
}

type Account struct {
	Id                    string
	Name                  string
	Environment           string
	HasRealTimeMonitoring bool
	SecurityPackage       bool
	CostPackage           bool
}

type CreateAccountRequest struct {
	Name                  string
	Environment           string
	Role                  string
	ExternalId            string
	HasRealTimeMonitoring bool
	CostPackage           bool
	SecurityPackage       bool
}

type AccountOverview struct {
	Id          string
	Name        string
	Environment string
}

type AccountAccessSettings struct {
	RoleArn    string
	ExternalId string
}

func NewClient(apiKey string, region string) *Client {
	return &Client{apiKey: apiKey, region: region}
}
