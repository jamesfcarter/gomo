package gomo

import "testing"

func TestNewClientCredentials(t *testing.T) {
	creds := NewClientCredentials("abc", "def")
	if creds.clientID != "abc" {
		t.Error("Incorrect client id")
	}
	if creds.clientSecret != "def" {
		t.Error("Incorrect client secret")
	}
	if creds.grantType() != clientCredentialsGrantType {
		t.Error("Incorrect grant type")
	}
}

func TestNewImplicitCredentials(t *testing.T) {
	creds := NewImplicitCredentials("abc")
	if creds.clientID != "abc" {
		t.Error("Incorrect client id")
	}
	if creds.grantType() != implicitGrantType {
		t.Error("Incorrect grant type")
	}
}
