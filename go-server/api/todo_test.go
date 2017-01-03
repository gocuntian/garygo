package api_test

import (
	"net/http"
	"testing"
)

func TestIndex(t *testing.T) {
	_, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("Did not return successful")
	}
}
