package appErr

type Error struct {
	Status  int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func New(status int, message string) *Error {
	return &Error{
		Status:  status,
		Message: message,
	}
}

var (
	ErrInvalidRequest     = New(400, "invalid request")
	ErrMissingField       = New(400, "missing required field")
	ErrInvalidID          = New(400, "invalid id")
	ErrInvalidDate        = New(400, "invalid date format")
	ErrInvalidURL         = New(400, "invalid url")
	ErrInvalidAmount      = New(400, "invalid amount")
	ErrInvalidQuantity    = New(400, "invalid quantity")
	ErrInvalidPagination  = New(400, "invalid pagination parameters")
	ErrInvalidSortField   = New(400, "invalid sort field")
	ErrInvalidFilterValue = New(400, "invalid filter value")
	ErrInvalidJSON        = New(400, "invalid json body")
	ErrInvalidContentType = New(400, "invalid content type")
	ErrRequestBodyEmpty   = New(400, "request body is empty")
	ErrRequestTooLarge    = New(400, "request body too large")
	ErrInvalidCoordinates = New(400, "invalid coordinates")
	ErrInvalidCurrency    = New(400, "invalid currency code")
	ErrInvalidCountryCode = New(400, "invalid country code")
	ErrInvalidLanguage    = New(400, "invalid language code")

	ErrInvalidEmail     = New(400, "invalid email address")
	ErrInvalidPassword  = New(400, "invalid password")
	ErrPasswordTooShort = New(400, "password is too short")
	ErrPasswordTooWeak  = New(400, "password is too weak")

	ErrInvalidPhone  = New(400, "invalid phone number")
	ErrInvalidName   = New(400, "invalid name")
	ErrInvalidRole   = New(400, "invalid role")
	ErrInvalidStatus = New(400, "invalid status")

	ErrFileRequired    = New(400, "file is required")
	ErrFileTooLarge    = New(400, "file too large")
	ErrInvalidFileType = New(400, "invalid file type")

	ErrOTPRequired = New(400, "otp is required")
	ErrOTPInvalid  = New(400, "invalid otp")
	ErrOTPNotSent  = New(400, "otp not sent")

	// Auth / Token (401)
	ErrSessionExpired      = New(401, "session has expired")
	ErrSessionInvalid      = New(401, "invalid session")
	ErrRefreshTokenExpired = New(401, "refresh token expired")
	ErrRefreshTokenInvalid = New(401, "invalid refresh token")
	ErrInvalidAPIKey       = New(401, "invalid api key")
	ErrAPIKeyExpired       = New(401, "api key expired")
	ErrMFARequired         = New(401, "multi-factor authentication required")
	ErrMFAInvalid          = New(401, "invalid mfa code")

	ErrUnauthorized       = New(401, "unauthorized")
	ErrTokenMissing       = New(401, "token missing")
	ErrTokenInvalid       = New(401, "invalid token")
	ErrTokenExpired       = New(401, "token expired")
	ErrTokenRevoked       = New(401, "token revoked")
	ErrAccountNotVerified = New(401, "account is not verified")

	ErrForbidden       = New(403, "forbidden")
	ErrAccountDisabled = New(403, "account is disabled")

	ErrUserNotFound    = New(404, "user not found")
	ErrProductNotFound = New(404, "product not found")

	ErrUserAlreadyExists = New(409, "user already exists")
	ErrEmailExists       = New(409, "email already exists")
	ErrOTPAlreadyUsed    = New(409, "otp already used")

	// Forbidden (403)
	ErrAccountSuspended  = New(403, "account is suspended")
	ErrAccountLocked     = New(403, "account is locked")
	ErrEmailNotVerified  = New(403, "email is not verified")
	ErrPhoneNotVerified  = New(403, "phone is not verified")
	ErrInsufficientPerms = New(403, "insufficient permissions")
	ErrOwnershipRequired = New(403, "you do not own this resource")
	ErrIPBlocked         = New(403, "access denied from this ip address")
	ErrRegionBlocked     = New(403, "service not available in your region")

	// Not Found (404)
	ErrRouteNotFound    = New(404, "route not found")
	ErrOrderNotFound    = New(404, "order not found")
	ErrCategoryNotFound = New(404, "category not found")
	ErrAddressNotFound  = New(404, "address not found")
	ErrCartNotFound     = New(404, "cart not found")
	ErrSessionNotFound = New(404, "session not found")
	ErrTokenNotFound   = New(404, "token not found")
	ErrFileNotFound    = New(404, "file not found")
	ErrRecordNotFound  = New(404, "record not found")

	// Conflict (409)
	ErrPhoneExists           = New(409, "phone number already exists")
	ErrUsernameExists        = New(409, "username already exists")
	ErrAlreadyVerified       = New(409, "account is already verified")
	ErrAlreadyEnabled        = New(409, "account is already enabled")
	ErrAlreadyDisabled       = New(409, "account is already disabled")
	ErrDuplicateEntry        = New(409, "duplicate entry")
	ErrOrderAlreadyPaid      = New(409, "order has already been paid")
	ErrOrderAlreadyCancelled = New(409, "order has already been cancelled")

	// Gone (410)
	ErrOTPExpired  = New(410, "otp expired")
	ErrLinkExpired = New(410, "link has expired")
	ErrTokenGone   = New(410, "token is no longer valid")

	// Unprocessable Entity (422)
	ErrUnprocessable     = New(422, "unprocessable entity")
	ErrBusinessLogic     = New(422, "business rule violation")
	ErrInsufficientStock = New(422, "insufficient stock")
	ErrInsufficientFunds = New(422, "insufficient funds")
	ErrCartEmpty         = New(422, "cart is empty")
	ErrMaxItemsExceeded  = New(422, "maximum items limit exceeded")
	ErrSelfAction        = New(422, "cannot perform this action on yourself")

	// Rate Limit (429)
	ErrTooManyRequests   = New(429, "too many requests")
	ErrOTPLimitReached   = New(429, "otp request limit reached")
	ErrOTPCooldown       = New(429, "please wait before requesting a new otp")
	ErrLoginLimitReached = New(429, "too many login attempts, please try again later")
	ErrSMSLimitReached   = New(429, "sms request limit reached")

	// Server Errors (500)
	ErrInternalServer = New(500, "internal server error")

	ErrDBInsertFailed = New(500, "failed to create record")
	ErrDBUpdateFailed = New(500, "failed to update record")
	ErrDBDeleteFailed = New(500, "failed to delete record")
	ErrDBQueryFailed  = New(500, "failed to fetch record")

	ErrEmailSendFailed = New(500, "failed to send email")

	// Server Errors (500)
	ErrDBConnectionFailed  = New(500, "database connection failed")
	ErrDBTransactionFailed = New(500, "database transaction failed")
	ErrCacheFailed         = New(500, "cache operation failed")
	ErrHashFailed          = New(500, "failed to hash data")
	ErrEncryptionFailed    = New(500, "encryption failed")
	ErrDecryptionFailed    = New(500, "decryption failed")
	ErrJWTSignFailed       = New(500, "failed to sign token")
	ErrFileSaveFailed      = New(500, "failed to save file")
	ErrFileDeleteFailed    = New(500, "failed to delete file")
	ErrSMSSendFailed       = New(500, "failed to send sms")
	ErrPushNotifFailed     = New(500, "failed to send push notification")
	ErrQueueFailed         = New(500, "failed to enqueue task")
	ErrExternalService     = New(500, "external service error")
	ErrTimeout             = New(500, "operation timed out")
	ErrConfigMissing       = New(500, "missing required configuration")

	// Not Implemented (501)
	ErrNotImplemented = New(501, "feature not yet implemented")

	// Gateway Errors (502)
	ErrBadGateway     = New(502, "bad gateway")
	ErrGatewayTimeout = New(504, "gateway timeout")

	// Service Unavailable (503)
	ErrServiceUnavailable = New(503, "service unavailable")
	ErrDBUnavailable      = New(503, "database service unavailable")
	ErrCacheUnavailable   = New(503, "cache service unavailable")
	ErrStorageUnavailable = New(503, "storage service unavailable")
	ErrMailUnavailable    = New(503, "mail service unavailable")
	ErrPaymentUnavailable = New(503, "payment service unavailable")
	ErrSMSUnavailable     = New(503, "sms service unavailable")
)
