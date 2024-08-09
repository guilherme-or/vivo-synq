package handler

type AppErr struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewAppErr(status int, code, message string) *AppErr {
	return &AppErr{Status: status, Code: code, Message: message}
}

func (e *AppErr) Error() string {
	return e.Message
}

var (
	ErrInvalidArgument  = NewAppErr(400, "INVALID_ARGUMENT", "Client specified an invalid argument, request body or query param")
	ErrPermissionDenied = NewAppErr(403, "PERMISSION_DENIED", "Authenticated user has no permission to access the requested resource")
	ErrNotFound         = NewAppErr(404, "NOT_FOUND", "A specified resource is not found")
	ErrInternal         = NewAppErr(500, "INTERNAL", "Server error")
	ErrTimeout          = NewAppErr(504, "TIMEOUT", "Request timeout exceeded. Try it later")
)
