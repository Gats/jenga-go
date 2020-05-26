package jenga

import (
	"errors"
	"github.com/gats/jenga-go/utilities"
	"github.com/gats/jenga-go/models"
	"encoding/json"
	"net/http"
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
}

func NewJenga(accessToken, environment string) (*JengaImpl, error) {
	baseUrl, err := getBaseURL(environment)
	if err != nil {
		return nil, err
	}
	return &JengaImpl{
		BaseURL: baseUrl,
		Token:  accessToken,
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
		panic(err1)
	}
	// var ts models.TokenResponse
	// var te models.ErrorResponse
	// if resp.StatusCode == http.StatusOK {
	// 	if err :=json.NewDecoder(resp.Body).Decode(&ts); err != nil {
	// 		return "", err
	// 	}
	// 	return map[string]interface{} {
	// 		"access_token": ts.AccessToken,
	// 		"expires_in": 
	// 	}, nil
	// }
	// if err :=json.NewDecoder(resp.Body).Decode(&te); err != nil {
	// 	return "", err
	// }
	// return "", errors.New(te.Message)
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