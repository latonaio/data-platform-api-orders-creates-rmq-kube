package handlers

type ErrorResponse struct {
	ResponseType string
	Name         string
	Message      string
}

const (
	typeMessage = "Message"
	typeError   = "Error"
)

const (
	InternalServerError = "INTERNAL_SERVER_ERROR"
)
