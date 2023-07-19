package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	ans := 2

	if ans != 0 {
		t.Errorf("Should fail: %d", ans)
	}
}
