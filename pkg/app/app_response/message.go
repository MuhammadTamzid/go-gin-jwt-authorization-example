package app_response

var MsgFlags = map[int]string{
	SUCCESS:                              "Success",
	INTERNAL_SERVER_ERROR:                "Internal server error",
	INVALID_DATA:                         "Bad request",
	FAIL_TO_REGISTER_USER:                "Fail to register user",
	USER_NOT_FOUND:                       "This user is not exist",
	INVALID_USER_EMAIL_OR_PASSWORD:       "Invalid user email or password",
	AUTHORIZATION_TOKEN_REQUIRED:         "Authorization Token is required",
	INCORRECT_AUTHORIZATION_TOKEN_FORMAT: "Incorrect Format of Authorization Token",
	INVALID_TOKRN_SIGNATURE:              "Invalid Token Signature",
	INVALID_TOKEN:                        "Invalid token",
	TOKEN_NOT_EXISTS:                     "Token not exists",
	COURSE_NOT_FOUND:                     "Course not found",
	NOT_AUTHORIZED:                       "Don't have permission",
	ENROLLED_COURSE_NOT_FOUND:            "Enrolled course not found",
	ALREADY_REGISTER_IN_THIS_COURSE:      "Already register in this course",
	NOT_ENROLLED:						  "Not enrolled",
}

// Get error message based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[INTERNAL_SERVER_ERROR]
}
