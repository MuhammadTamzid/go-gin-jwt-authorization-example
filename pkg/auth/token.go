package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"log"
	"net/http"
	"go-gin-jwt-authorization-example/configs"
	"go-gin-jwt-authorization-example/pkg/app/app_response"
	"go-gin-jwt-authorization-example/pkg/constant"
	"go-gin-jwt-authorization-example/pkg/enums"
	"strings"
	"time"
)

const(
	HEADER_TYPE = "typ"
	JWT_VALUE = "JWT"
	TOKEN_PREFIX = "Bearer "
)

type TokenDto struct {
	AccessToken         string
	RefreshToken        string
	AccessTokenExpires  int64
	RefreshTokenExpires int64
}

type AccessDetails struct {
	TokenUuid string
	UserId    string
}

func GenerateToken(userId string) (*TokenDto, error) {
	tokenEnv := configs.EnvVariables.Token
	accessTokenString, accessTokenErr := generateJWTToken(userId, tokenEnv.AccessTimeExpired, enums.Access)
	if accessTokenErr != nil {
		return nil, accessTokenErr
	}

	refreshTokenString, refreshTokenErr := generateJWTToken(userId, tokenEnv.RefreshTimeExpired, enums.Refresh)
	if refreshTokenErr != nil {
		return nil, refreshTokenErr
	}

	tokenDto := &TokenDto{}
	tokenDto.AccessToken = accessTokenString
	tokenDto.RefreshToken = refreshTokenString
	tokenDto.AccessTokenExpires = tokenEnv.AccessTimeExpired
	tokenDto.RefreshTokenExpires = tokenEnv.RefreshTimeExpired

	return tokenDto, nil
}

func ParseToken(tokenString string, token *jwt.Token) (int, int, bool) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.EnvVariables.Token.Secret), nil
	})

	if err != nil {
		log.Println(err)
		if err == jwt.ErrSignatureInvalid {
			return  http.StatusUnauthorized, app_response.INVALID_TOKEN, true
		}
		return  http.StatusBadRequest, app_response.INVALID_TOKEN, true
	}

	if !parsedToken.Valid {
		return  http.StatusUnauthorized, app_response.INVALID_TOKEN, true
	}

	*token = *parsedToken
	return 0, 0, false
}

func GetTokenFromHeader(clientToken string, token *string) (int, int, bool) {
	if clientToken == "" {
		return  http.StatusUnauthorized, app_response.AUTHORIZATION_TOKEN_REQUIRED, true
	}

	extractedToken := strings.Split(clientToken, TOKEN_PREFIX)
	if len(extractedToken) != 2 {
		return http.StatusBadRequest, app_response.INCORRECT_AUTHORIZATION_TOKEN_FORMAT, true
	}

	*token = extractedToken[1]
	return 0, 0, false
}

func GetUserIDFromToken(token jwt.Token) string {
	claims := token.Claims.(jwt.MapClaims)
	return claims[constant.USER_ID].(string)
}

func generateJWTToken(userId string, expirationTime int64, tokenType byte) (string, error)  {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Header[HEADER_TYPE] = JWT_VALUE
	expirationTime = time.Now().Add(time.Duration(expirationTime)).Unix()

	claims := token.Claims.(jwt.MapClaims)
	claims[constant.ID] = uuid.NewV4().String()
	claims[constant.TYPE] = tokenType
	claims[constant.EXP] = expirationTime
	claims[constant.USER_ID] = userId

	tokenString, err := token.SignedString([]byte(configs.EnvVariables.Token.Secret))
	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	return tokenString, nil
}
