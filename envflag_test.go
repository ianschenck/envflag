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
