package jenga

import (
	"errors"
	"log"

	"github.com/gats/jenga-go/utilities"
)


func (j *JengaImpl) GetAccountBalance(account, countryCode string) error {
	ret := utilities.IsAllowedCountryCode(countryCode)
	if !ret {
		return nil, errors.New("Country code not allowed")
	}
	data :=  account + countryCode
	sig, err := utilities.GenerateSignature(data, j.KeyPath)
	if err != nil {
		return nil, err
	}
	return nil,nil
}
