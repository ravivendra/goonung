package goonung

import "testing"

func TestPalindrom(t *testing.T) {
	texts := []string{}

	if _, err := Palindrom(texts); err.Error() == "Empty parameter" {
		t.Errorf("Empty parameter : %s", err.Error())
	}

	texts = []string{"Facebook"}

	if _, err := Palindrom(texts); err != nil {
		t.Errorf("No palindrom parameter : %s", err.Error())
	}

	texts = []string{"Kasur rusak"}

	if _, err := Palindrom(texts); err != nil {
		t.Errorf("Palindrom parameter : %s", err.Error())
	}
}
