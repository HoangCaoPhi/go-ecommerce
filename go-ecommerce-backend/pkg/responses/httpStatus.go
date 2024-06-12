package responses

const (
	// Success
	ErrorCodeSuccess      = 20001
	ErrorCodeParamInvalid = 20003
)

var message = map[int]string{
	ErrorCodeSuccess:      "Success",
	ErrorCodeParamInvalid: "Param is invalid",
}
