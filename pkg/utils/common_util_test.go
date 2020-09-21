package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRedisKey(t *testing.T) {
	id := "1"
	assert.Equal(t, GetRedisKey(id), "jwt_token:"+id)
}
