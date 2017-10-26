package yod

import (
	"errors"
	"net/http"
)

type Writer struct {
	w http.ResponseWriter
}

func (w *Writer) SetHeader(key, value string) {
	w.w.Header().Set(key, value)
}

func (w *Writer) OK(v interface{}) error {
	if s, ok := v.(string); ok {
		_, err := w.w.Write([]byte(s))
		return err
	}
	return errors.New("writer not support content-type")
}

func (w *Writer) InternalServerError(v interface{}) error {
	return nil
}

func (w *Writer) Informational(code int, v interface{}) error {
	return nil
}

func (w *Writer) Successful(code int, v interface{}) error {
	return nil
}

func (w *Writer) Redirection(code int, v interface{}) error {
	return nil
}

func (w *Writer) ClientError(code int, v interface{}) error {
	return nil
}

func (w *Writer) ServerError(code int, v interface{}) error {
	return nil
}

type StringWriter struct {
	ResponseWriter
}

func (StringWriter) Value(code int, v interface{}) error {
	return nil
}

type JSONWriter struct {
	ResponseWriter
}

func (JSONWriter) Value(code int, v interface{}) error {
	return nil
}

type XMLWriter struct {
	ResponseWriter
}

func (XMLWriter) Value(code int, v interface{}) error {
	return nil
}
