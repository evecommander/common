package utils

import (
	"testing"
)

func TestCrypto(t *testing.T) {
	p := BcryptEncode("123qwe")
	valid := BcryptCheck("123qwe", p)

	if !valid {
		t.Fatal()
	}

	valid = BcryptCheck("qwe123", "$2a$10$8YFqO7jCoOSx9t6q5uWwv.TEyr61HfRHSmdn/2wo1MjOuot.MyG6i")

	if !valid {
		t.Failed()
	}
}
