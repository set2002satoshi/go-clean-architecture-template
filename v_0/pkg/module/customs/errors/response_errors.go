package errors

import (
	"fmt"
)

type IError interface {
	SetErr(err error)
}

func Response(err error, r IError) IError {
	cErr := NewCustomError()
	cErr = Combine(cErr, err)
	m := ToMap(cErr)
	for k, v := range m {
		m[k] = fmt.Sprintf("%s:%s", k, v)
	}
	r.SetErr(cErr)
	return r
}
