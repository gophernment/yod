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

func TestWriterOKWithOtherTypeShouldError(t *testing.T) {
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

func TestWriterInternalServerErrorWithOtherTypeShouldError(t *testing.T) {
	recoder := httptest.NewRecorder()
	w := Writer{
		w: recoder,
	}

	err := w.InternalServerError(map[string]string{"status": "internal server error"})
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

func TestWriterInformationalExceptOnltStringType(t *testing.T) {
	recoder := httptest.NewRecorder()
	w := Writer{
		w: recoder,
	}

	w.Informational(http.StatusContinue, "continue")
	resp := recoder.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if string(body) != "continue" {
		t.Errorf("It should write plain text %q but was %q\n", "continue", string(body))
	}

	if recoder.Code != http.StatusContinue {
		t.Errorf("It should write status code as %d error but was %d\n", http.StatusContinue, recoder.Code)
	}

	recoder2 := httptest.NewRecorder()
	w2 := Writer{
		w: recoder2,
	}
	w2.Informational(http.StatusContinue, map[string]string{"status": "other type"})

	resp2 := recoder2.Result()
	body, _ = ioutil.ReadAll(resp2.Body)

	if string(body) != writerNotSupportDataType {
		t.Errorf("It should write plain text %q but was %q\n", writerNotSupportDataType, string(body))
	}

	// if recoder.Code != http.StatusInternalServerError {
	// 	t.Error("It should write status code as http internal server error but was", recoder.Code)
	// }
}
