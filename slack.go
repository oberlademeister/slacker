package slacker

import (
	"bytes"
	"net/http"
)

// SendBuffer can be used as a Reader that the
type SendBuffer struct {
	webhookURL      string
	useFlushString  bool
	flushString     string
	hideFlushString bool
	alwaysFlush     bool
	buffer          bytes.Buffer
	client          *http.Client
}

// NewSendBuffer creates a new SlackSender
func NewSendBuffer(webhookURL string, opts ...SBOption) *SendBuffer {
	ss := &SendBuffer{webhookURL: webhookURL, client: &http.Client{}}
	for _, opt := range opts {
		opt(ss)
	}
	return ss
}

// Write fulfills io.Writer
func (ss *SendBuffer) Write(p []byte) (n int, err error) {
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
func (ss *SendBuffer) Flush() error {
	return Send(ss.webhookURL, ss.buffer.String(), ss.client)
}
