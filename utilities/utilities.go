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

func GenerateToken(apiKey, username, password, baseUrl, path string) *http.Response {
    data := url.Values{}
    data.Set("username", username)
    data.Set("password", password)
    log.Printf("%v\n", data)
    u, _ := url.ParseRequestURI(baseUrl)
    u.Path = path
    URL := u.String()
    log.Println(URL)
    client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
    r, _ := http.NewRequest("POST", URL, strings.NewReader(data.Encode()))
    r.Header.Add("Authorization", "Basic "+ apiKey)
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    log.Printf("%v\n", r.Header)
    resp, _ := client.Do(r)
	log.Printf("%v\n", resp)
	return resp
}