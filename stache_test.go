package main

import "testing"

func TestParseContextStrings(t *testing.T) {
	s := parseContextStrings([]string{"test=one=two"})
	if s["test"] != "one=two" {
		t.Fail()
	}

	s = parseContextStrings([]string{"test=one", "test=two"})
	if s["test"] != "two" {
		t.Fail()
	}

	s = parseContextStrings([]string{"a=one", "b=two"})
	if s["a"] != "one" && s["b"] != "two" {
		t.Fail()
	}
}
