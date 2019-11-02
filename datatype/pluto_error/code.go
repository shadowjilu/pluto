package pluto_error

import "net/http"

var (
	ServerError = NewPlutoError(http.StatusInternalServerError, 1000, "Server Error", nil)
	BadRequest  = NewPlutoError(http.StatusBadRequest, 1001, "Bad Request", nil)

	MailIsAlreadyRegister = NewPlutoError(http.StatusForbidden, 2001, "Mail is already been registered", nil)
	MailIsNotExsit        = NewPlutoError(http.StatusForbidden, 2002, "Mail is not exist", nil)
	MailIsNotVerified     = NewPlutoError(http.StatusForbidden, 2003, "Mail is not verified", nil)
	MailAlreadyVerified   = NewPlutoError(http.StatusBadRequest, 2004, "Mail is already verified", nil)

	InvalidPassword      = NewPlutoError(http.StatusForbidden, 3001, "Invalid Password", nil)
	InvalidRefreshToken  = NewPlutoError(http.StatusForbidden, 3002, "Invalid Refresh Token", nil)
	InvalidJWTToekn      = NewPlutoError(http.StatusForbidden, 3003, "Invalid JWT Token", nil)
	InvalidGoogleIDToken = NewPlutoError(http.StatusForbidden, 3004, "Invalid Google ID Token", nil)
	InvalidWechatCode    = NewPlutoError(http.StatusForbidden, 3005, "Invalid Wechat Code", nil)
	InvalidAvatarFormat  = NewPlutoError(http.StatusBadRequest, 3006, "Invalid Avatar Format", nil)

	JWTTokenExpired = NewPlutoError(http.StatusForbidden, 3007, "JWT Token Expired", nil)
)
