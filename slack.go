package slacker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// SlackSender can be used as a Reader that the
type SlackSender struct {
	webhookURL      string
	useFlushString  bool
	flushString     string
	hideFlushString bool
	alwaysFlush     bool
	buffer          bytes.Buffer
	client          *http.Client
}

// NewSlackSender creates a new SlackSender
func NewSlackSender(webhookURL string, opts ...SlackSenderOption) *SlackSender {
	ss := &SlackSender{webhookURL: webhookURL, client: &http.Client{}}
	for _, opt := range opts {
		opt(ss)
	}
	return ss
}

// Write fulfills io.Writer
func (ss *SlackSender) Write(p []byte) (n int, err error) {
	if ss.alwaysFlush {
		n, err := ss.buffer.Write(p)
		if err != nil {
			return n, err
		}
		err = ss.Flush()
		return n, err
	}
	if ss.useFlushString && ss.flushString == string(p) {
		if !ss.hideFlushString {
			n, err := ss.buffer.Write(p)
			if err != nil {
				return n, err
			}
		}
		err = ss.Flush()
		return n, err
	}
	return ss.buffer.Write(p)

}

// Flush sends the buffer contents to slack
func (ss *SlackSender) Flush() error {
	bodyStruct := struct {
		Text string `json:"text"`
	}{
		Text: ss.buffer.String(),
	}
	body, err := json.Marshal(&bodyStruct)
	if err != nil {
		return errors.Wrap(err, "can't marshal body")
	}
	r := bytes.NewReader(body)
	req, err := http.NewRequest("POST", ss.webhookURL, r)
	if err != nil {
		return errors.Wrap(err, "failed to create POST request")
	}
	req.Header.Set("Content-type", "application/json")
	fmt.Println(req)
	res, err := ss.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send http request")
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("http response status %d (), not 200", res.StatusCode, res.Status)
	}
	return nil

}
