package responses

const (
	// Success
	ErrorCodeSuccess      = 20001
	ErrorCodeParamInvalid = 20003
	ErrInvalidToken       = 30001
)

var message = map[int]string{
	ErrorCodeSuccess:      "Success",
	ErrorCodeParamInvalid: "Param is invalid",
	ErrInvalidToken:       "Token is invalid",
}
