package errors


type ErrorInfo struct {
	ErrCode string `json:"err_code"`
	ErrMsg string `json:"err_msg"`

}

type Errors struct {
	Errors []ErrorInfo `json:"errors"`
}