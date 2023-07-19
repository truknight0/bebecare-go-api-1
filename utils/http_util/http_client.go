package http_util

import (
	"bytes"
	"net/http"
)

func Post(url string, content []byte, header map[string]string) (*http.Response, error) {
	body := bytes.NewBuffer(content)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
