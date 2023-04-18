package xerrors

import (
	"errors"
)

var (
	InvalidBodyRequestStatus = errors.New("INVALID_BODY_OR_PARAM_REQUEST")
	InvalidBodyRequestCode   = 4001

	InvalidCredentialsStatus = errors.New("INVALID_CREDENTIALS")
	InvalidCredentialsCode   = 4002

	InternalErrorStatus = errors.New("SERVER_ERROR")
	InternalErrorCode   = 5001
)
