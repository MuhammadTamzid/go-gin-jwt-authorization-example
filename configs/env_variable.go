package configs

import (
	"os"
	"strconv"
)

var EnvVariables Env

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Redis struct {
	Host     string
	Port     int
	Password string
}

type Token struct {
	Secret             string
	AccessTimeExpired  int64
	RefreshTimeExpired int64
}

type Env struct {
	Database Database
	Redis    Redis
	Token    Token
	Port     int
}

func InitEnv() {
	EnvVariables = Env{
		Database: Database{
			Host:     GetEnv("DB_HOST", ""),
			Port:     GetEnvAsInt("DB_PORT", 3306),
			User:     GetEnv("DB_USER", ""),
			Password: GetEnv("DB_PASSWORD", ""),
			DBName:   GetEnv("DB_NAME", ""),
		},
		Redis: Redis{
			Host:     GetEnv("REDIS_HOST", ""),
			Port:     GetEnvAsInt("REDIS_PORT", 6379),
			Password: GetEnv("REDIS_PASSWORD", ""),
		},
		Token: Token{
			Secret:             GetEnv("TOKEN_SECRET", ""),
			AccessTimeExpired:  GetEnvAsInt64("ACCESS_TIME_EXPIRED", 900000000000),
			RefreshTimeExpired: GetEnvAsInt64("REFRESH_TIME_EXPIRED", 604800000000000),
		},
		Port: GetEnvAsInt("PORT", 4000),
	}
}

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func GetEnvAsInt64(name string, defaultVal int64) int64 {
	valueStr := GetEnv(name, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}

	return defaultVal
}
