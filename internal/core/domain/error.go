package domain

type CustomErr struct {
	code string
	message string
	details map[string]interface{}
}

func (e CustomErr) Code() string {
	return e.code
}

func (e CustomErr) Error() string {
	return e.message
}

func (e CustomErr) Details() map[string]interface{} {
	return e.details
}

func CustomError(code string, message string, details map[string]interface{}) *CustomErr {
	return &CustomErr{code, message, details}
}

const (
	NotFound = "not-found"
	Unknown = "unknown"
	InvalidRequest = "invalid-request"
)

var (
	ErrNotFound = &CustomErr{NotFound, "No resource found", nil}
	ErrUnknown = &CustomErr{Unknown, "Unknown error occurred. Please contact support", nil}
	ErrInvalidLoginDetails = &CustomErr{InvalidRequest, "Invalid login details", nil}
)