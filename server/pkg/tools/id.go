package tools

import (
	"github.com/matoous/go-nanoid/v2"
)

var (
	alphabet  = "0123456789abcdefghijklmnopqrstuvwxyz"
	space     = "spc"
	account   = "acc"
	datasheet = "dts"
)

func CreateSpaceId() string {
	if id, err := gonanoid.Generate(alphabet, 10); err == nil {
		return space + id
	}
	return ""
}

func CreateDatasheetId() string {
	if id, err := gonanoid.Generate(alphabet, 10); err == nil {
		return datasheet + id
	}
	return ""
}

func CreateAccount() string {
	if id, err := gonanoid.Generate(alphabet, 10); err == nil {
		return account + id
	}
	return ""
}
