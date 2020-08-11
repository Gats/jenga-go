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

func GetMiniStatement(accessToken, env, countryCode, account, keyPath string) (map[string]interface{}, error){
	jg, err := jenga.NewJenga(accessToken, env, keyPath)
	if err != nil {
		log.Printf("%v", err)
	}
	return jg.GetMiniStatement(account, countryCode)
}

func GetFullStatement(accessToken, env, countryCode, account, fromDate, toDate, keyPath string, limit int) (map[string]interface{}, error){
	jg, err := jenga.NewJenga(accessToken, env, keyPath)
	if err != nil {
		log.Printf("%v", err)
	}
	return jg.GetFullStatement(account, countryCode, fromDate, toDate, limit)
}

func GetOpenAndCloseBalance(accessToken, env, countryCode, account, keyPath, date string) (map[string]interface{}, error){
	jg, err := jenga.NewJenga(accessToken, env, keyPath)
	if err != nil {
		log.Printf("%v", err)
	}
	return jg.GetOpenAndCloseBalance(account, countryCode, date)
}

func GetAccountDetails(accessToken, env, countryCode, account, keyPath string) (map[string]interface{}, error){
	jg, err := jenga.NewJenga(accessToken, env, keyPath)
	if err != nil {
		log.Printf("%v", err)
	}
	return jg.GetAccountDetails(account, countryCode)
}




