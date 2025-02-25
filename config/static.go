package config

// status codes
const (
	ErrorCodeInvalidURL        = "INVALID_URL"
	ErrorCodeInvalidJSON       = "INVALID_JSON"
	ErrorCodeInternal          = "INTERNAL"
	ErrorCodeAlreadyExists     = "ALREADY_EXISTS"
	ErrorCodeNotFound          = "NOT_FOUND"
	ErrorBadRequest            = "BAD_REQUEST"
	ErrorCodeUnauthorized      = "UNAUTHORIZED"
	ErrorCodeForbidden         = "FORBIDDEN"
	ErrorCodeTokenExpired      = "TOKEN_EXPIRED"
	ErrorCodeNotApproved       = "NOT_APPROVED"
	ErrorCodeIncorrectPassword = "INCORRECT_PASSWORD"
	ErrorCodeInvalidToken      = "INVALID_TOKEN"
	ErrorCodeSessionLimit      = "SESSION_LIMIT"
	ErrorCodeInvalidCode       = "INVALID_CODE"
	StatusSuccess              = "SUCCESS"
	Status2faRequired          = "2FA_REQUIRED"
)

// grpc context timeout
const (
	GrpcContextTimeout = 15
)

// token types
const (
	AccessToken    = "access"
	RefreshToken   = "refresh"
	Verfication    = "verification"
	ForgetPassword = "forget_password"
	ResetPassword  = "reset_password"
)

// openapis
var (
	OpenApis = map[string]bool{}
)
