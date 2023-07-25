package tools

import (
	"testing"
)

func TestGenerateFromPassword(t *testing.T) {
	password := "12345678"
	pwd, _ := GeneratePassword(password)

	pass := ComparePassword(pwd, password)
	t.Log(pass)
}
