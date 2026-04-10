package model

import (
	"fmt"
	"runtime"
)

type wrapError struct {
	err      *error
	nameAddr *string
	file     *string
	line     *int
	ok       *bool
}
type WrapError interface {
	WrapError() string
}

func (err *wrapError) WrapError() string {

	return fmt.Sprintf("Name: %s\nFile: %s\nLine: %d\nMessageL:  %s", *err.nameAddr, *err.file, *err.line, (*err.err).Error())
}

func NewError(err *error) WrapError {
	var fatherFunctionCaller = 1
	pc, file, line, ok := runtime.Caller(fatherFunctionCaller)
	if !ok {
		print("loi toi")
	}
	nameAddr := runtime.FuncForPC(pc).Name()
	return &wrapError{
		err:      err,
		nameAddr: new(nameAddr),
		file:     &file,
		line:     &line,
	}
}
