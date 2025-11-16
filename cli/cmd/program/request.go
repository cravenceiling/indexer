package program

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cravenceiling/indexer/cli/internal/document"
)

type Credentials struct {
	User     string
	Password string
}

type HttpRequest struct {
	Creds   Credentials
	BaseURL string
	Port    string
	Index   string
	Type    string
}

type IndexDocument struct {
	Index string `json:"_index"`
}

type IndexAction struct {
	Index IndexDocument `json:"index"`
}

type Payload struct {
	Index        string              `json:"index"`
	DocumentData []document.Document `json:"records"`
}

var client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     30 * time.Second,
	},
}

// Upload
func Upload(re HttpRequest, payload *bytes.Buffer) error {
	u := fmt.Sprintf("http://%s:%s/api/_bulk", re.BaseURL, re.Port)
	req, err := http.NewRequest("POST", u, payload)
	if err != nil {
		return err
	}

	req.SetBasicAuth(re.Creds.User, re.Creds.Password)
	req.Header.Set("Content-Type", "application/x-ndjson")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := getBodyResponse(res)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("status code: %d - %s\n", res.StatusCode, body)
	}

	payload.Reset()

	return nil
}

func getBodyResponse(res *http.Response) (string, error) {
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
