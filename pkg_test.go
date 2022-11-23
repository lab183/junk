package junk

import (
	"testing"
)

func Test_Junk(t *testing.T) {
	resp := Junk()
	if resp == "" {
		t.Errorf("Expected something other than %s", resp)
		t.Fail()
		return
	}
}
