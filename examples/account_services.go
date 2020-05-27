package examples

import(
	"github.com/gats/jenga-go"
	"log"
)


func AccountBalance(accessToken, env, countryCode, account, keyPath string) (map[string]interface{}, error){
	jg, err := jenga.NewJenga(accessToken, env, keyPath)
	if err != nil {
		log.Printf("%v", err)
	}
	return jg.GetAccountBalance(account, countryCode)
}




