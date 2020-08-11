package utilities

import (
	"net/http"
	"net/url"
    "time"
    "strings"
    "sort"
    "log"
    "bytes"
)

func MakeGetRequest(signature, token, baseURL, path string) *http.Response {
    u, _ := url.ParseRequestURI(baseURL)
    u.Path = path
    URL := u.String()
    log.Println(URL)
    client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
    r, _ := http.NewRequest("GET", URL, nil)
    r.Header.Add("Authorization", "Bearer "+ token)
    r.Header.Add("signature", signature)
    log.Printf("%v", r.Header)
    resp, _ := client.Do(r)
	return resp
}

func MakePostRequest(signature, token, baseURL, path string, reqBody []byte) *http.Response {
    u, _ := url.ParseRequestURI(baseURL)
    u.Path = path
    URL := u.String()
    log.Println(URL)
    client := http.Client{
		Timeout: time.Second * 30, // Maximum of 30 secs
	}
	log.Println(string(reqBody))
	req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(reqBody))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Printf("%v", err)
	}
	return res
}

func GenerateToken(apiKey, username, password, baseUrl, path string) *http.Response {
    data := url.Values{}
    data.Set("username", username)
    data.Set("password", password)
    u, _ := url.ParseRequestURI(baseUrl)
    u.Path = path
    URL := u.String()
    client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
    r, _ := http.NewRequest("POST", URL, strings.NewReader(data.Encode()))
    r.Header.Add("Authorization", "Basic "+ apiKey)
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    resp, _ := client.Do(r)
	return resp
}

func IsAllowedCountryCode(code string) bool {
    allowed := []string{"KE", "UG", "TZ", "RW", "DRC", "SS"}
    sort.Strings(allowed)
    i := sort.SearchStrings(allowed, code)
    if i < len(allowed) && allowed[i] == code {
        return true
    }
    return false
}