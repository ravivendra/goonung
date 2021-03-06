package goonung

import "testing"

func TestCalculate(t *testing.T) {
	if _, err := Calculate(""); err != nil {
		t.Errorf("Input by string argument : %s", err.Error())
	}

	if _, err := Calculate(100.01); err != nil {
		t.Errorf("Input by float argument : %s", err.Error())
	}

	if _, err := Calculate(1234567890123456); err != nil {
		t.Errorf("Input by not zero integer argument : %s", err.Error())
	}

	if _, err := Calculate(0); err != nil {
		t.Errorf("Input by zero integer argument : %s", err.Error())
	}
}

func TestIsValidNumber(t *testing.T) {
	if ok := IsValidNumber(0); ok != true {
		t.Errorf("Invalid number for Luhn calculation")
	}
}
