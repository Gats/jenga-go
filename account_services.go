package jenga

import (
	"errors"
	"log"

	"github.com/gats/jenga-go/utilities"
)


func (j *JengaImpl) GetAccountBalance(account, countryCode string) (map[string]interface{}, error) {
	ret := utilities.IsAllowedCountryCode(countryCode)
	if !ret {
		return nil, errors.New("Country code not allowed")
	}
	data :=  account + countryCode
	signature, err := utilities.GenerateSignature(data, j.KeyPath)
	if err != nil {
		return nil, err
	}
	uridata := "/accounts/balances/"+countryCode+"/"+account
	endpoint := sandBoxAccountsEndpoint+uridata
	resp := utilities.MakeGetRequest(signature, j.Token, j.BaseURL, endpoint) 
	log.Printf("%v", resp)
	return nil,nil
}
