package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) keyParam() string {
	return "?key=" + c.key
}

func (c *Client) url(model, action string) string {
	if action == "" {
		return c.baseUrl + model
	}
	return c.baseUrl + model + ":" + action
}

func (c *Client) post(url string, body interface{}) (rst []byte, err error) {
	return c.simpleReq("POST", url, body)
}

func (c *Client) newReq(method string, url string, body any) (req *http.Request, err error) {
	var reader io.Reader = nil
	if body != nil {
		var bs []byte
		bs, err = json.Marshal(body)
		if err != nil {
			return
		}
		reader = bytes.NewReader(bs)
	}

	if c.auth == AuthByUrlQuery {
		url = url + c.keyParam()
	}
	req, err = http.NewRequest(method, url, reader)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if c.auth == AuthByHttpHeader {
		req.Header.Set("Authorization", "Bearer "+c.key)
	}
	return
}

func (c *Client) handleReq(req *http.Request) (rst []byte, err error) {
	rsp, err := c.hc.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	rst, err = io.ReadAll(rsp.Body)
	return
}

func (c *Client) simpleReq(method string, url string, body any) (rst []byte, err error) {
	req, err := c.newReq(method, url, body)
	if err != nil {
		return
	}
	return c.handleReq(req)
}

func (c *Client) get(url string) (rst []byte, err error) {
	return c.simpleReq("GET", url, nil)
}

func unjson[T any](bs []byte, e error) (rst T, err error) {
	if e != nil {
		err = e
		return
	}
	err = json.Unmarshal(bs, &rst)
	return
}

func jout(v any) {
	bs, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(bs))
}
