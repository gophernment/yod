package yod

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriterOK(t *testing.T) {
	recoder := httptest.NewRecorder()
	w := Writer{
		w: recoder,
	}

	err := w.OK("ok")
	if err != nil {
		t.Error("It should not error", err)
		return
	}

	resp := recoder.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if string(body) != "ok" {
		t.Error("It should write plain text ok but was", string(body))
	}

	if recoder.Code != http.StatusOK {
		t.Error("It should write status code as http ok but was", recoder.Code)
	}
}

func TestWriterWithOtherTypeShouldError(t *testing.T) {
	recoder := httptest.NewRecorder()
	w := Writer{
		w: recoder,
	}

	err := w.OK(map[string]string{"status": "ok"})
	if err != nil {
		t.Error("It should not error")
		return
	}

	resp := recoder.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if string(body) != writerNotSupportDataType {
		t.Errorf("It should write plain text %q but was %q\n", writerNotSupportDataType, string(body))
	}

	if recoder.Code != http.StatusInternalServerError {
		t.Error("It should write status code as http internal server error but was", recoder.Code)
	}
}

func TestWriterInternalServerError(t *testing.T) {
	recoder := httptest.NewRecorder()
	w := Writer{
		w: recoder,
	}

	err := w.InternalServerError("internal server error")
	if err != nil {
		t.Error("It should not error", err)
		return
	}

	resp := recoder.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if string(body) != "internal server error" {
		t.Error("It should write plain text ok but was", string(body))
	}

	if recoder.Code != http.StatusInternalServerError {
		t.Error("It should write status code as http ok but was", recoder.Code)
	}
}
