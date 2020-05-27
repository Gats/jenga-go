package jenga

import (
	"errors"
	"log"

	"github.com/gats/jenga-go/utilities"
)


func (j *JengaImpl) GetAccountBalance(accInfo map[string]interface{}) error {
	ret := utilities.IsAllowedCountryCode(accInfo["countrycode"])
	if !ret {
		return errors.New("Country code not allowed")
	}
	data :=  accInfo["account"].(string) + accInfo["countrycode"].(string)
	sig, err := utilities.GenerateSignature(data, j.KeyPath)
	if err != nil {
		log.Printf("%v", err)
	}
}
