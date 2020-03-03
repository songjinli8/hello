package hello

import "testing"

func TestHello(t *testing.T) {
	if got := Hello(); got == "你好，世。" {
		t.Errorf("Hello() = %q", got)
	}
}
