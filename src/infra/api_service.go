package infra

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ApiService interface {
	Call(method, url string, body, model any) error
}

type ApiServiceImpl struct {
	ApiKey string
}

func (s ApiServiceImpl) Call(method, url string, body, model any) error {
	postBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(postBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.ApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	err = s.readResponse(res, model)
	if err != nil {
		return err
	}

	return nil
}

func (s *ApiServiceImpl) readResponse(res *http.Response, model any) error {
	// TODO: Replace any for a custom interface
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, model)
	if err != nil {
		return err
	}

	return nil
}
