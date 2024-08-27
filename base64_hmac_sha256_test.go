package goonung

import (
	"testing"
)

func TestCalculateBase64Hmac256(t *testing.T) {
	if _, err := CalculateBase64Hmac256("Ravi Vendra R", "github"); err != nil {
		t.Errorf("Message or secret is broken : %s", err.Error())
	}
}
