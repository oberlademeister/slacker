package slacker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Send sends the String to Slack
func Send(webhookURL string, message string, client *http.Client) error {
	bodyStruct := struct {
		Text string `json:"text"`
	}{
		Text: message,
	}
	body, err := json.Marshal(&bodyStruct)
	if err != nil {
		return errors.Wrap(err, "can't marshal body")
	}
	r := bytes.NewReader(body)
	req, err := http.NewRequest("POST", webhookURL, r)
	if err != nil {
		return errors.Wrap(err, "failed to create POST request")
	}
	req.Header.Set("Content-type", "application/json")
	if client == nil {
		client = &http.Client{}
	}
	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send http request")
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("http response status %d (), not 200", res.StatusCode, res.Status)
	}
	return nil
}
