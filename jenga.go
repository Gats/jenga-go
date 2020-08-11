package jenga

import (
	"errors"
	"github.com/gats/jenga-go/utilities"
	"encoding/json"
	"io/ioutil"
)

const (
	productionBaseURL = ""
	sandBoxBaseURL = "https://uat.jengahq.io/"
	sandBoxAccountsEndpoint = "account-test/v2/"
	sandBoxTransactionEndpoint = "transaction-test/v2/"
	sandBoxCustomerEndpoint = "customer-test/v2/"
	sandboxTokenEndpoint = "identity/v2/token"
)

// type Jenga interface {
// 	GetAccountBalance()
// }

type JengaImpl struct {
	BaseURL string
	Token	string
	KeyPath	string
}

func NewJenga(accessToken, environment, keyPath string) (*JengaImpl, error) {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return nil, err
	}
	return &JengaImpl{
		BaseURL: baseURL,
		Token:  accessToken,
		KeyPath: keyPath,
	}, nil
}

func GetAccessToken(apikey, username, password, environment string) (map[string]interface{}, error) {
	baseURL, err := getBaseURL(environment)
	if err != nil {
		return nil, err
	}
	resp := utilities.GenerateToken(apikey, username, password, baseURL, sandboxTokenEndpoint) 
	defer resp.Body.Close()
	resMap := make(map[string]interface{})
	body, _ := ioutil.ReadAll(resp.Body)
	err1 := json.Unmarshal([]byte(body), &resMap)
	if err1 != nil {
		return nil, err1
	}
	return resMap, nil
}

func getBaseURL(environment string) (string, error) {
	var baseURL string
	if environment == "dev" {
		baseURL = sandBoxBaseURL
	} else if environment == "prod" {
		baseURL = productionBaseURL
	} else {
		err := errors.New("Environment must either be 'dev' or 'prod'. case sensitive ")
		return "", err
	}
	return baseURL, nil
}