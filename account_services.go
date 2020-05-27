package jenga

import (
	"errors"
	"log"

	"github.com/gats/jenga-go/utilities"
)


func (j *JengaImpl) GetAccountBalance(account, countryCode string) error {
	ret := utilities.IsAllowedCountryCode(countryCode)
	if !ret {
		return errors.New("Country code not allowed")
	}
	data :=  account + countryCode
	sig, err := utilities.GenerateSignature(data, j.KeyPath)
	if err != nil {
		log.Printf("%v", err)
	}
	log.Println(sig)
	return nil
}
