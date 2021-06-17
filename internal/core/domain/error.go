package domain

type CustomErr struct {
	code string
	message string
}

func (e CustomErr) Code() string {
	return e.code
}

func (e CustomErr) Error() string {
	return e.message
}

func CustomError(code string, message string) CustomErr {
	return CustomErr{code, message}
}

const (
	NotFound = "not-found"
	Unknown = "unknown"
)

var (
	ErrNotFound = CustomErr{NotFound, "No resource found"}
	ErrUnknown = CustomErr{Unknown, "Unknown error occurred. Please contact support"}
)