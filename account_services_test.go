package jenga

import (
	"testing"
)

func TestGetAccountBalance(t *testing.T) {
	jg, err := NewJenga(accessToken, environment, keyPath)
	resp, err := jg.GetAccountBalance("12345678", "KE")
	if err != nil {
		t.Errorf("Didnt expect errors but got %v", err)
	}
	t.Logf("%v %v", resp, err) 
}

func TestGetMiniStatement(t *testing.T) {
	jg, err := NewJenga(accessToken, environment, keyPath)
	resp, err := jg.GetMiniStatement("12345678", "KE")
	if err != nil {
		t.Errorf("Didnt expect errors but got %v", err)
	}
	t.Logf("%v %v", resp, err) 
}

func TestGetFullStatement(t *testing.T) {
	jg, err := NewJenga(accessToken, environment, keyPath)
	resp, err := jg.GetFullStatement("12345678", "KE", "2019-12-01", "2020-03-31", 10)
	if err != nil {
		t.Errorf("Didnt expect errors but got %v", err)
	}
	t.Logf("%v %v", resp, err) 
}

func TestGetOpenAndCloseBalance(t *testing.T) {
	jg, err := NewJenga(accessToken, environment, keyPath)
	resp, err := jg.GetOpenAndCloseBalance("12345678", "KE", "2020-06-04")
	if err != nil {
		t.Errorf("Didnt expect errors but got %v", err)
	}
	t.Logf("%v %v", resp, err) 
}

func TestGetAccountDetails(t *testing.T) {
	jg, err := NewJenga(accessToken, environment, keyPath)
	resp, err := jg.GetAccountDetails("12345678", "KE")
	if err != nil {
		t.Errorf("Didnt expect errors but got %v", err)
	}
	t.Logf("%v %v", resp, err) 
}