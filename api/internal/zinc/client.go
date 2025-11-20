package zinc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/cravenceiling/indexer/api/internal/config"
	_ "github.com/joho/godotenv/autoload"
)

// ZincSearch Client
type Client struct {
	Host     string
	Port     string
	Username string
	Password string
	http     *http.Client
}

// NewClient
func NewClient() *Client {
	return &Client{
		Host:     config.GetEnv("ZINC_HOST", "localhost"),
		Port:     config.GetEnv("ZINC_PORT", "4080"),
		Username: config.GetEnv("ZINC_USERNAME", "admin"),
		Password: config.GetEnv("ZINC_PASSWORD", "Complexpass#123"),
		http: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) PingDB() error {
	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", c.Host, c.Port),
		Path:   "/version",
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil
	}

	req.SetBasicAuth(c.Username, c.Password)
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("unexpected status code connecting to Zinc: %d", res.StatusCode)
	}

	return nil
}

// DoZincRequest
func (c *Client) DoZincRequest(r *http.Request, query matchQuery) (*ZincResponse, error) {
	index := r.URL.Query().Get("index")
	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", c.Host, c.Port),
		Path:   fmt.Sprintf("/api/%s/_search", index),
	}

	body, err := json.Marshal(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("query: ", string(body))

	req, err := http.NewRequest("POST", url.String(), bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36",
	)

	resp, err := c.http.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	zr := &ZincResponse{}
	err = json.NewDecoder(resp.Body).Decode(&zr)
	if err != nil {
		log.Println("error decoding json response: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	return zr, nil
}
