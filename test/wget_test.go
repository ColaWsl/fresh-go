package test

import (
	"github.com/ColaWsl/fresh-go/something"
	"testing"
)

func TestLogic1(t *testing.T) {
	ret := something.Wget()
	if ret != true {
		t.Error("fail")
	}
	t.Log("success")
}
