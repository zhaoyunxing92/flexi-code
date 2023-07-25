package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	load, err := Load("../../conf/application.yaml")
	assert.Nil(t, err)
	assert.NotNil(t, load)
}
