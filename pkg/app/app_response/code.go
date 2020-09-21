package app_response

const (
	SUCCESS                              = 200
	INTERNAL_SERVER_ERROR                = 500
	INVALID_DATA                         = 400
	FAIL_TO_REGISTER_USER                = 10001
	USER_NOT_FOUND                       = 10002
	INVALID_USER_EMAIL_OR_PASSWORD       = 10003
	AUTHORIZATION_TOKEN_REQUIRED         = 10004
	INCORRECT_AUTHORIZATION_TOKEN_FORMAT = 10005
	INVALID_TOKRN_SIGNATURE              = 10006
	INVALID_TOKEN                        = 10007
	TOKEN_NOT_EXISTS                     = 10008
	COURSE_NOT_FOUND                     = 10009
	NOT_AUTHORIZED                       = 10010
	ENROLLED_COURSE_NOT_FOUND            = 10011
	ALREADY_REGISTER_IN_THIS_COURSE      = 10012
	NOT_ENROLLED      					 = 10013
)
