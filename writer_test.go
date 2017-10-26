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
