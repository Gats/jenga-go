package jenga

import (
	"log"
	"github.com/gats/jenga-go/utilities"
)


func (j *JengaImpl) GetAccountBalance(accInfo map[string]interface{}) {
	log.Printf("%v\n", j)
	data :=  accInfo["account"].(string) + accInfo["countrycode"].(string)
	log.Println(data)
	sig, err := utilities.GenerateSignature(data, j.KeyPath)
	if err != nil {
		log.Printf("%v", err)
	}
	log.Println(sig)
}