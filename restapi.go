package autoria

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	APIBaseURL = "https://developers.ria.com"
)

func (s *service) request(method, path string, params map[string]string, body, response any) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, APIBaseURL+path, buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("api_key", s.apikey)
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	case http.StatusForbidden:
		return ErrInvalidAPIKey
	case http.StatusUnauthorized:
		return ErrInvalidAPIKey
	case http.StatusTooManyRequests:
		return ErrTooManyRequests
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if s.debug {
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, responseBody, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("Body: %s\n", prettyJSON.String())
	}

	if response != nil {
		if err := json.Unmarshal(responseBody, &response); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
