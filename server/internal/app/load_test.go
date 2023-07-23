package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	load, err := Load("../../conf/chat.yaml")
	assert.Nil(t, err)
	assert.NotNil(t, load)
}
