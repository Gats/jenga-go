package jenga

import (
	"testing"
	"reflect"
)

const(
	accessToken="12345tryryryryryr"
	environment="dev"
	keyPath="/path/to/privatekey"
	prodBaseURL = ""
	devBaseURL = "https://uat.jengahq.io/"

)
func TestNewJenga(t *testing.T) {
	j, err := NewJenga(accessToken, environment, keyPath)
	if !reflect.DeepEqual(j.Token, accessToken) {
		t.Errorf("Expected %s but got %s", accessToken, j.Token)
	}
	if !reflect.DeepEqual(j.KeyPath, keyPath) {
		t.Errorf("Expected %s but got %s", keyPath, j.KeyPath)
	}
	if environment == "dev" {
		if !reflect.DeepEqual(j.BaseURL, sandBoxBaseURL) {
			t.Errorf("Expected %s but got %s", sandBoxBaseURL, j.BaseURL)
		}
	} else if environment == "prod" {
		if !reflect.DeepEqual(j.BaseURL, productionBaseURL) {
			t.Errorf("Expected %s but got %s", productionBaseURL, j.BaseURL)
		}
	} else {
		if err == nil {
			t.Error("Expected an error but did not")
		}
	}
}

func TestGetAccessToken(t *testing.T) {
	// Test for invalid parameters
	resp, err := GetAccessToken("12345678", "12345", "rwttetet", "dev")
	t.Logf("%v %v", resp, err)
}
