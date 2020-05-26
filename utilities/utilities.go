package utilities

import (
	"net/http"
	"net/url"
    "time"
    "strings"
    "log"
)

func MakeRequest() {

}

func GenerateToken(apiKey, username, password, baseUrl string) *http.Response {
    data := url.Values{}
    data.Set("username", username)
    data.Set("password", password)
    u, _ := url.ParseRequestURI(baseUrl)
    u.Path = "/identity-test/v2/token"
    endpoint := u.String()
    client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
    r, _ := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
    r.Header.Add("Authorization", "Basic "+ apiKey)
    resp, _ := client.Do(r)
	log.Printf("%v", resp)
	return resp
}