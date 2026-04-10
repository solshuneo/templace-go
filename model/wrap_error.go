package model

import (
	"fmt"
	"runtime"
)

type wrapError struct {
	msg      string
	nameAddr string
	file     string
	line     int
}
type WrapError interface {
	String() string
}

func (err *wrapError) String() string {

	return fmt.Sprintf("--------DEBUGGER-----------------\n"+
		"Name: %s\nFile: %s\nLine: %d\nMessageL:  %s\n"+
		"--------DEBUGGER-----------------\n", err.nameAddr, err.file, err.line, err.msg)
}

func NewError(err error) WrapError {
	var fatherFunctionCaller = 1
	pc, file, line, ok := runtime.Caller(fatherFunctionCaller)
	if !ok {
		panic("Panic in NewError")
	}
	nameAddr := runtime.FuncForPC(pc).Name()
	return &wrapError{
		msg:      err.Error(),
		nameAddr: nameAddr,
		file:     file,
		line:     line,
	}
}
