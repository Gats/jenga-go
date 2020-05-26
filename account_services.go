package jenga

import (
	"log"
)


func (j *JengaImpl) GetAccountBalance(accInfo map[string]interface{}) {
	log.Printf("%v\n", j)
	log.Printf("%v\n", accInfo)

}