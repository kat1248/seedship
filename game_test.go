package main

import "testing"

func TestStringInSlice(t *testing.T) {
	v := []string{"foo", "bar", "baz"}
	b := stringInSlice("foo", v)
	if !b {
		t.Errorf("stringInSlice didn't find foo")
	}
	b = stringInSlice("hello", v)
	if b {
		t.Errorf("stringInSlice found hello")
	}
}
