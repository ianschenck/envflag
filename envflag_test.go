package envflag

import (
	"os"
	"testing"
)

func TestParseString(t *testing.T) {
	val := String("VALUE", "default", "set this value to a string")
	os.Setenv("VALUE", "somevalue")
	Parse()
	if *val != "somevalue" {
		t.Errorf("Got unexpected value: %s", *val)
	}
}

func TestParsePrefix(t *testing.T) {
	val := String("VALUE_WITH_PREFIX", "default", "set this value to a string")
	os.Setenv("ENV_PREFIXED_VALUE_WITH_PREFIX", "somevalue")
	ParsePrefix("ENV_PREFIXED_")
	if *val != "somevalue" {
		t.Errorf("Got unexpected value: %s", *val)
	}
}
