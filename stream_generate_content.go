package gemini

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Limit-LAB/go-gemini/models"
	"io"
	"strings"
)

func (c *Client) GenerateContentStream(model models.GeminiModel, req *models.GenerateContentRequest) (*GenerateContentStreamer, error) {
	for _, content := range req.Contents {
		err := validateGenerateContentRequest(model, content)
		if err != nil {
			return nil, err
		}
	}
	url := c.url(string(model), "streamGenerateContent")
	httpReq, err := c.newReq("POST", url, req)
	if err != nil {
		return nil, err
	}
	resp, err := c.hc.Do(httpReq)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		defer resp.Body.Close()
		var bs []byte
		bs, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var eResp []models.ErrorResponse
		err = json.Unmarshal(bs, &eResp)
		if err != nil {
			return nil, err
		}
		if len(eResp) > 0 {
			return nil, errors.New(eResp[0].Error.Message)
		}
	}
	return newStreamScanner(resp.Body), nil
}

type GenerateContentStreamer struct {
	buf *bufio.Scanner
	raw io.ReadCloser
}

func (r *GenerateContentStreamer) Close() error {
	return r.raw.Close()
}

func (r *GenerateContentStreamer) Receive() (models.GenerateContentResponse, error) {
	if !r.buf.Scan() {
		err := r.buf.Err()
		if err == nil {
			err = io.EOF
		}
		return models.GenerateContentResponse{}, err
	}
	txt := r.buf.Text()
	// remove head '[' and tail ']'
	txt = strings.TrimLeft(txt, "[,\r\n")
	txt = strings.TrimRight(txt, "],\r\n")
	var res models.GenerateContentResponse
	err := json.Unmarshal([]byte(txt), &res)
	return res, err
}

func newStreamScanner(eventStream io.ReadCloser) *GenerateContentStreamer {
	scanner := bufio.NewScanner(eventStream)

	scanner.Buffer(make([]byte, 4096), 4096)

	split := func(data []byte, atEOF bool) (int, []byte, error) {
		if i := bytes.Index(data, []byte("}\n,\r\n")); i >= 0 {
			return i + 5, data[0 : i+1], nil
		}
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
	scanner.Split(split)

	return &GenerateContentStreamer{
		buf: scanner,
		raw: eventStream,
	}
}
