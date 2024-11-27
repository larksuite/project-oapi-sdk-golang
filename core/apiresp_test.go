package core

import "testing"

func TestAPIResp_Error_String(t *testing.T) {
	err := CodeError{
		ErrCode: 10001,
		ErrMsg:  "error message",
		Err: struct {
			Code  int    `json:"code,omitempty"`
			Msg   string `json:"msg,omitempty"`
			LogID string `json:"log_id,omitempty"`
		}{
			Code:  10002,
			Msg:   "inner error message",
			LogID: "20210101000001",
		},
	}
	t.Logf(err.String())
}
