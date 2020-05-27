package utilities

import (
	"net/http"
	"net/url"
    "time"
    "strings"
    // "log"
)

func MakeRequest() {

}

func GenerateToken(apiKey, username, password, baseUrl, path string) *http.Response {
    data := url.Values{}
    data.Set("username", username)
    data.Set("password", password)
    u, _ := url.ParseRequestURI(baseUrl)
    u.Path = path
    URL := u.String()
    // log.Println(URL)
    client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
    r, _ := http.NewRequest("POST", URL, strings.NewReader(data.Encode()))
    r.Header.Add("Authorization", "Basic "+ apiKey)
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    // log.Printf("%v\n", r.Header)
    resp, _ := client.Do(r)
	// log.Printf("%v\n", resp)
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