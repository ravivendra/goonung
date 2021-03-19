package goonung

import "testing"

func TestIsArray(t *testing.T) {
	if _, err := IsArray("Ravi Vendra R"); err != nil {
		t.Errorf("Input by invalid array argument : %s", err.Error())
	}
}

func TestIsBool(t *testing.T) {
	if _, err := IsBool(100.10); err != nil {
		t.Errorf("Input by invalid bool argument : %s", err.Error())
	}
}

func TestIsComplex(t *testing.T) {
	if _, err := IsComplex(100); err != nil {
		t.Errorf("Input by invalid complex argument : %s", err.Error())
	}
}

func TestIsFloat(t *testing.T) {
	if _, err := IsFloat(100); err != nil {
		t.Errorf("Input by invalid float argument : %s", err.Error())
	}
}

func TestIsInt(t *testing.T) {
	if _, err := IsInt("Ravi Vendra R"); err != nil {
		t.Errorf("Input by invalid integer argument : %s", err.Error())
	}
}

func TestIsString(t *testing.T) {
	if _, err := IsString(100.10); err != nil {
		t.Errorf("Input by invalid string argument : %s", err.Error())
	}
}

func TestIsSlice(t *testing.T) {
	if _, err := IsSlice(100.10); err != nil {
		t.Errorf("Input by invalid slice argument : %s", err.Error())
	}
}

func TestIsStruct(t *testing.T) {
	if _, err := IsStruct(100.10); err != nil {
		t.Errorf("Input by invalid struct argument : %s", err.Error())
	}
}

func TestIsUint(t *testing.T) {
	if _, err := IsUint(100.10); err != nil {
		t.Errorf("Input by invalid uint argument : %s", err.Error())
	}
}

func TestIsUintPtr(t *testing.T) {
	if _, err := IsUintPtr(100.10); err != nil {
		t.Errorf("Input by invalid uint pointer argument : %s", err.Error())
	}
}
