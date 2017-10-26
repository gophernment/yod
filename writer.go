package yod

import (
	"net/http"
)

const (
	writerNotSupportDataType = "writer not support data type"
)

type Writer struct {
	w http.ResponseWriter
}

func (w *Writer) SetHeader(key, value string) {
	w.w.Header().Set(key, value)
}

func (w *Writer) internalServerErrorString() error {
	w.w.WriteHeader(http.StatusInternalServerError)
	_, err := w.w.Write([]byte(writerNotSupportDataType))
	return err
}

func (w *Writer) OK(v interface{}) error {
	if s, ok := v.(string); ok {
		w.w.WriteHeader(http.StatusOK)
		_, err := w.w.Write([]byte(s))
		return err
	}

	return w.InternalServerError(writerNotSupportDataType)
}

func (w *Writer) InternalServerError(v interface{}) error {
	if s, ok := v.(string); ok {
		w.w.WriteHeader(http.StatusInternalServerError)
		_, err := w.w.Write([]byte(s))
		return err
	}

	return w.internalServerErrorString()
}

func (w *Writer) String(code int, s string) error {
	w.w.WriteHeader(code)
	_, err := w.w.Write([]byte(s))
	return err
}

func (w *Writer) Informational(code int, v interface{}) error {
	if s, ok := v.(string); ok {
		return w.String(code, s)
	}

	return w.internalServerErrorString()
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

type StringWriter = Writer

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
