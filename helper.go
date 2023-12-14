package gemini

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func post[T any](hc *http.Client, url string, body interface{}) (rst T, err error) {
	bs, err := json.Marshal(body)
	if err != nil {
		return
	}
	rsp, err := hc.Post(url, "application/json", bytes.NewReader(bs))
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	bs, err = io.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, &rst)
	return
}

func get[T any](hc *http.Client, url string) (rst T, err error) {
	rsp, err := hc.Get(url)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	bs, err := io.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, &rst)
	return
}
