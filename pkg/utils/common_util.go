package utils

import "go-gin-jwt-authorization-example/pkg/constant"

func GetRedisKey(id string) string {
	return JoinString(constant.JWT_REDIS_KEY_SEPARATOR, constant.JWT_REDIS_KEY_PREFIX, id)
}
