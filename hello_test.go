package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	TestObj := hello()
	TestSam := "hello Wen"
	if TestObj != TestSam {
		t.Errorf("want %s, got: %s", TestSam, TestObj)
	}
}
