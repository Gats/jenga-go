package jenga

import (
	"errors"
	"fmt"
	// "log"
	"encoding/json"
	"io/ioutil"
	"github.com/gats/jenga-go/utilities"
)

// GetAccountBalance - fetches account balance from the Jenga HQ API
// Accepts account and country code and returns a map[string]interface{} and or an error
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
	uridata := fmt.Sprintf("accounts/balances/%s/%s", countryCode, account)
	endpoint := sandBoxAccountsEndpoint+uridata
	resp := utilities.MakeGetRequest(signature, j.Token, j.BaseURL, endpoint) 
	defer resp.Body.Close()
	resMap := make(map[string]interface{})
	body, _ := ioutil.ReadAll(resp.Body)
	err1 := json.Unmarshal([]byte(body), &resMap)
	if err1 != nil {
		return nil, err1
	}
	return resMap, nil
}

// GetMiniStatement - fetches the mini statement from the Jenga HQ API
// Accepts account and country code and returns a map[string]interface{} and or an error
func (j *JengaImpl) GetMiniStatement(account, countryCode string) (map[string]interface{}, error) {
	ret := utilities.IsAllowedCountryCode(countryCode)
	if !ret {
		return nil, errors.New("Country code not allowed")
	}
	data :=  account + countryCode
	signature, err := utilities.GenerateSignature(data, j.KeyPath)
	if err != nil {
		return nil, err
	}
	uridata := fmt.Sprintf("accounts/ministatement/%s/%s", countryCode, account)
	endpoint := sandBoxAccountsEndpoint+uridata
	resp := utilities.MakeGetRequest(signature, j.Token, j.BaseURL, endpoint) 
	defer resp.Body.Close()
	resMap := make(map[string]interface{})
	body, _ := ioutil.ReadAll(resp.Body)
	err1 := json.Unmarshal([]byte(body), &resMap)
	if err1 != nil {
		return nil, err1
	}
	return resMap, nil
}

// GetFullStatement - fetches the full statement from the Jenga HQ API
// Current implementation only accepts account, countryCode, from date, to date and the limit.
// From and To Dates should be in the ISO 8601 format i.e. YYYY-MM-DD 
// Other paramaters to be added later.
func (j *JengaImpl) GetFullStatement(account, countryCode, fromDate, toDate string, limit int) (map[string]interface{}, error) {
	ret := utilities.IsAllowedCountryCode(countryCode)
	if !ret {
		return nil, errors.New("Country code not allowed")
	}
	// Technical Debt - To add other parameters as provided by the spec.
	type FullStatement struct{
		CCode      string `json:"countryCode"`
		AccNUmber string	`json:"accountNumber"`
		FromDate string	`json:"fromDate"`
		ToDate string	`json:"toDate"`
		Limit int	`json:"limit"`
	}

	u := FullStatement{
		CCode: countryCode,
		AccNUmber: account,
		FromDate: fromDate,
		ToDate: toDate,
		Limit: limit,
	}
	msg, _ := json.Marshal(u)
	data :=  account + countryCode + toDate
	signature, err := utilities.GenerateSignature(data, j.KeyPath)
	if err != nil {
		return nil, err
	}
	uridata := "accounts/fullstatement"
	endpoint := sandBoxAccountsEndpoint+uridata
	resp := utilities.MakePostRequest(signature, j.Token, j.BaseURL, endpoint, msg) 
	defer resp.Body.Close()
	resMap := make(map[string]interface{})
	body, _ := ioutil.ReadAll(resp.Body)
	err1 := json.Unmarshal([]byte(body), &resMap)
	if err1 != nil {
		return nil, err1
	}
	return resMap, nil
}
// GetOpenAndCloseBalance - fetches the full statement from the Jenga HQ API
// Accepts account, countryCode and date.
// Date should be in the ISO 8601 format i.e. YYYY-MM-DD 
func (j *JengaImpl) GetOpenAndCloseBalance(account, countryCode, date string) (map[string]interface{}, error) {
	ret := utilities.IsAllowedCountryCode(countryCode)
	if !ret {
		return nil, errors.New("Country code not allowed")
	}
	type FullStatement struct{
		CCode      string `json:"countryCode"`
		AccNUmber string	`json:"accountId"`
		Date string	`json:"date"`
	}

	u := FullStatement{
		CCode: countryCode,
		AccNUmber: account,
		Date: date,
	}
	msg, _ := json.Marshal(u)
	data :=  account + countryCode + date
	signature, err := utilities.GenerateSignature(data, j.KeyPath)
	if err != nil {
		return nil, err
	}
	uridata := "accounts/accountbalance/query"
	endpoint := sandBoxAccountsEndpoint+uridata
	resp := utilities.MakePostRequest(signature, j.Token, j.BaseURL, endpoint, msg) 
	defer resp.Body.Close()
	resMap := make(map[string]interface{})
	body, _ := ioutil.ReadAll(resp.Body)
	err1 := json.Unmarshal([]byte(body), &resMap)
	if err1 != nil {
		return nil, err1
	}
	return resMap, nil
}

// GetAccountDetails - fetches the the details of an account (Account Inquiry)
// Accepts account and country code and returns a map[string]interface{} and or an error
func (j *JengaImpl) GetAccountDetails(account, countryCode string) (map[string]interface{}, error) {
	ret := utilities.IsAllowedCountryCode(countryCode)
	if !ret {
		return nil, errors.New("Country code not allowed")
	}
	data :=  account + countryCode
	signature, err := utilities.GenerateSignature(data, j.KeyPath)
	if err != nil {
		return nil, err
	}
	uridata := fmt.Sprintf("search/%s/%s", countryCode, account)
	endpoint := sandBoxAccountsEndpoint+uridata
	resp := utilities.MakeGetRequest(signature, j.Token, j.BaseURL, endpoint) 
	defer resp.Body.Close()
	resMap := make(map[string]interface{})
	body, _ := ioutil.ReadAll(resp.Body)
	err1 := json.Unmarshal([]byte(body), &resMap)
	if err1 != nil {
		return nil, err1
	}
	return resMap, nil
}
