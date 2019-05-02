package sdk

type createAccountAccessKeys struct {
	RoleArn    string `json:"roleArn"`
	ExternalId string `json:"externalId"`
}

type createAccountAccess struct {
	Keys createAccountAccessKeys `json:"keys"`
}

type createAccountAttributes struct {
	Name                  string              `json:"name"`
	Environment           string              `json:"environment"`
	Access                createAccountAccess `json:"access"`
	CostPackage           bool                `json:"costPackage"`
	HasRealTimeMonitoring bool                `json:"hasRealTimeMonitoring"`
	SecurityPackage       bool                `json:"securityPackage"`
}

type createAccountData struct {
	Type       string                  `json:"type"`
	Attributes createAccountAttributes `json:"attributes"`
}

type createAccountResponse struct {
	Id string `json:"id"`
}

func (client Client) CreateAccount(request CreateAccountRequest) (string, error) {
	payload := struct {
		Data createAccountData `json:"data"`
	}{}
	payload.Data.Type = "account"
	payload.Data.Attributes.Access.Keys.ExternalId = request.ExternalId
	payload.Data.Attributes.Access.Keys.RoleArn = request.Role
	payload.Data.Attributes.Environment = request.Environment
	payload.Data.Attributes.Name = request.Name
	payload.Data.Attributes.CostPackage = request.CostPackage
	payload.Data.Attributes.SecurityPackage = request.SecurityPackage
	payload.Data.Attributes.HasRealTimeMonitoring = request.HasRealTimeMonitoring

	output := struct {
		Data createAccountResponse `json:"data"`
	}{}

	err := client.genericPost("accounts", &payload, &output)
	if err != nil {
		return "", err
	}
	return output.Data.Id, nil
}
