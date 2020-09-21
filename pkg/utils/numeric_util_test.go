package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCastStringToInt(t *testing.T)  {
	value, err := CastStringToInt("5")
	assert.Equal(t, err, nil)
	assert.Equal(t, value, 5)
}
