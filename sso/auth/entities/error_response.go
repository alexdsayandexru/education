package entities

const (
	UnsupportedGrantType = "unsupported_grant_type"
	UnauthorizedClient   = "unauthorized_client"
	InvalidScope         = "invalid_scope"
	InvalidGrant         = "invalid_grant"
	InvalidClient        = "invalid_client"
	InvalidRequest       = "invalid_request"
	InvalidToken         = "invalid_token"
)

func NewErrorResponse(code int, err, desc string) *ErrorResponse {
	return &ErrorResponse{
		Code:             code,
		Error:            err,
		ErrorDescription: desc,
	}
}

type ErrorResponse struct {
	Code             int    `json:"-"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
