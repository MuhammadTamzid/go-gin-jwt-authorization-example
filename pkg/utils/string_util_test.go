package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoinString(t *testing.T) {
	assert.Equal(t, JoinString(":", "str1", "str2"), "str1:str2")
	assert.Equal(t, JoinString(":", "str1"), "str1")
}
