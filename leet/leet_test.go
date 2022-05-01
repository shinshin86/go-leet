package leet

import "testing"

func TestLeet(t *testing.T) {
	leet1 := Leet("hello")
	if leet1 == "hello" {
		t.Errorf("leet1 is equal hello. want not equal")
	}

	leet2 := Leet("hello1234567890")
	if leet2 == "hello1234567890" {
		t.Errorf("leet2 is equal hello1234567890. want not equal")
	}

	leet3 := Leet("1234567890")
	if leet3 != "1234567890" {
		t.Errorf("leet3 is not equal 1234567890. want equal")
	}
}
