package voicetext

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

const endpoint = "https://api.voicetext.jp/v1/tts"

// NewRequest creates new request.
func NewRequest(apiKey, text, speaker string) (*http.Request, error) {
	vs := url.Values{}
	vs.Set("text", text)
	vs.Set("speaker", speaker)
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(vs.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(apiKey, "")
	return req, nil
}

// Text2Speech writes wav binary.
func Text2Speech(dst io.Writer, apiKey, text, speaker string) error {
	req, err := NewRequest(apiKey, text, speaker)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	_, err = io.Copy(dst, resp.Body)
	return err
}
