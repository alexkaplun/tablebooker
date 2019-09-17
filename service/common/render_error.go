package common

import (
	"errors"
	"net/http"
)

type RenderError struct {
	err error
}

func NewRenderError() RenderError {
	return RenderError{}
}

func (RenderError) WithWrapError(err error, text string) RenderError {
	return RenderError{
		err: errors.New(text + " " + err.Error()),
	}
}

func (RenderError) WithError(err error) RenderError {
	return RenderError{
		err: err,
	}
}

func (re RenderError) Render(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}
