// go test
// go test -fuzz=Fuzz
// go test -fuzz=FuzzEnvInt

package main

import (
	"testing"
)

func TestEnv(t *testing.T) {
	t.Setenv("STRING_VALUE", "1337")
	result := mustParseEnv("STRING_VALUE", "8080", false)
	if result != "1337" {
		t.Errorf("reading env given")
	}

	t.Setenv("STRING_VALUE_DEFAULT", "")
	result = mustParseEnv("STRING_VALUE_DEFAULT", "8080", false)
	if result != "8080" {
		t.Errorf("reading env default")
	}

	_, err := parseEnv("VALUE_NOT_SET", "8080", true)
	if err == nil {
		t.Errorf("reading env not set")
	}

	t.Setenv("VALUE_SET_EMPTY", "")
	result, err = parseEnv("VALUE_SET_EMPTY", "8080", true)
	if err == nil {
		t.Errorf("reading env set empty %q", result)
	}

	t.Setenv("INT_VALUE", "1338")
	result_int := mustParseEnv("INT_VALUE", 80, false)
	if result_int != 1338 {
		t.Errorf("reading env given int")
	}

	t.Setenv("BOOL_VALUE", "true")
	result_false := mustParseEnv("BOOL_VALUE", false, false)
	if result_false != true {
		t.Errorf("reading env given bool")
	}

}

func FuzzEnvString(f *testing.F) {
	f.Setenv("FUZZ_VALUE", "stage")
	f.Add("stage")
	f.Fuzz(func(t *testing.T, a string) {
		out, err := parseEnv("FUZZ_VALUE", "local", false)
		if err != nil && out != a {
			t.Errorf("%q %v", out, err)
		}
	})
}

func FuzzEnvInt(f *testing.F) {
	f.Setenv("FUZZ_VALUE", "1337")
	f.Add(1337)
	f.Fuzz(func(t *testing.T, a int) {
		out, err := parseEnv("FUZZ_VALUE", 1337, false)
		if err != nil && out != a {
			t.Errorf("%q %v", out, err)
		}
	})
}

func FuzzEnvBool(f *testing.F) {
	f.Setenv("FUZZ_VALUE", "true")
	f.Add("true")
	f.Fuzz(func(t *testing.T, a string) {
		out, err := parseEnv("FUZZ_VALUE", "false", false)
		if err != nil && out != a {
			t.Errorf("%q %v", out, err)
		}
	})
}
