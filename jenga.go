package jenga

import (
	"errors"
	"github.com/gats/jenga-go/utilities"
	"github.com/gats/jenga-go/models"
	"encoding/json"
	"net/http"
)

const (
	productionBoxBaseURL = ""
	sandBoxBaseURL = "https://sandbox.jengahq.io/"
	sandBoxAccountsBase = "account-test/v2/"
	sandBoxTransactionBase = "transaction-test/v2/"
	sandBoxCustomerBase = "customer-test/v2/"
	sanboxTokenBase = "identity-test/v2/token"
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

func GetAccessToken(apikey, username, password, environment string) (string, error) {
	baseUrl, err := getBaseURL(environment)
	if err != nil {
		return "", err
	}
	resp := utilities.GenerateToken(apikey, username, password, baseUrl) 
	defer resp.Body.Close()
	var ts models.TokenResponse
	var te models.ErrorResponse
	if resp.StatusCode == http.StatusOK {
		if err :=json.NewDecoder(resp.Body).Decode(&ts); err != nil {
			return "", err
		}
		return string(mt.AccessToken), nil
	}
	if err :=json.NewDecoder(resp.Body).Decode(&te); err != nil {
		return "", err
	}
	return "", errors.New(te.Message)
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