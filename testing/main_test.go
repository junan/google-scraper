package main

import (
	"testing"
)

func BenchmarkBcrypt(b *testing.B) {
	Bcrypt()
}
