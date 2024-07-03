package main

import (
	"testing"
)

func TestLogic1(t *testing.T) {
	ret := Wget()
	if ret != true {
		t.Error("fail")
	}
	t.Log("success")
}
