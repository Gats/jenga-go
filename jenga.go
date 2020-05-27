package jenga

import (
	"errors"
	"github.com/gats/jenga-go/utilities"
	// "github.com/gats/jenga-go/models"
	"encoding/json"
	// "net/http"
	"io/ioutil"
)

const (
	productionBoxBaseURL = ""
	sandBoxBaseURL = "https://uat.jengahq.io/"
	sandBoxAccountsEndpoint = "account-test/v2/"
	sandBoxTransactionEndpoint = "transaction-test/v2/"
	sandBoxCustomerEndpoint = "customer-test/v2/"
	sandboxTokenEndpoint = "identity/v2/token"
)

type Jenga interface {
	GetAccountBalance()
}

type JengaImpl struct {
	BaseURL string
	Token	string
	KeyPath	string
}

func NewJenga(accessToken, environment, keyPath string) (*JengaImpl, error) {
	baseRL, err := getBaseURL(environment)
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
	baseUrl, err := getBaseURL(environment)
	if err != nil {
		return nil, err
	}
	resp := utilities.GenerateToken(apikey, username, password, baseUrl, sandboxTokenEndpoint) 
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
	var baseUrl string
	if environment == "dev" {
		baseUrl = sandBoxBaseURL
	} else if environment == "prod" {
		baseUrl = productionBoxBaseURL
	} else {
		err := errors.New("Environment must either be 'dev' or 'prod'. case sensitive ")
		return "", err
	}
	return baseUrl, nil
}