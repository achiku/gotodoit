package estc

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// constants
const (
	SuccessStatusCode = 0
	ErrorStatusCode   = 1
)

// Client api client
type Client struct {
	client *http.Client
	config *Config
	logger *log.Logger
}

// Config api client config
type Config struct {
	BaseEndpoint string `toml:"base_endpoint"`
	APIKey       string `toml:"api_key"`
	APISecret    string `toml:"api_secret"`
	Debug        bool   `toml:"debug"`
}

// NewClient creates api client
func NewClient(cfg *Config, c *http.Client, logger *log.Logger) *Client {
	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}
	return &Client{
		client: c,
		config: cfg,
		logger: logger,
	}
}

func (c *Client) createSig(nonce int64, url, payload string) string {
	sig := hmac.New(sha256.New, []byte(c.config.APISecret))
	message := fmt.Sprintf("%d%s%s", nonce, url, payload)
	sig.Write([]byte(message))
	return hex.EncodeToString(sig.Sum(nil))
}

func (c *Client) call(
	ctx context.Context, method, pathStr string, request interface{}, response interface{}) error {
	payload, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}
	if c.config.Debug {
		c.logger.Printf("request: %s", payload)
	}

	endpoint := fmt.Sprintf("%s%s", c.config.BaseEndpoint, pathStr)
	req, err := http.NewRequest(method, endpoint, strings.NewReader(string(payload)))
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}
	req.WithContext(ctx)

	nonce := time.Now().UnixNano()
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("ACCESS-KEY", c.config.APIKey)
	req.Header.Add("ACCESS-NONCE", fmt.Sprintf("%d", nonce))
	req.Header.Add("ACCESS-SIGNATURE", c.createSig(nonce, endpoint, string(payload)))

	res, err := c.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to request")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			return errors.Wrap(err, "httputil.DumpResponse failed")
		}
		return errors.Errorf(
			"status code: %d, body: %s", res.StatusCode, dump)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(response); err != nil {
		return errors.Wrap(err, "failed to decode response")
	}

	if c.config.Debug {
		b, err := json.Marshal(response)
		if err != nil {
			return err
		}
		c.logger.Printf("response: %s", b)
	}

	return nil
}

// Task request
type Task struct {
	Name       string `json:"name"`
	Difficulty int    `json:"difficulty"`
}

// ETCResponse estimated time to complete
type ETCResponse struct {
	StatusCode      int `json:"statusCode"`
	Message         int `json:"message,omitempty"`
	Time            int `json:"time,omitempty"`
	ConfidenceLevel int `json:"confidenceLevel,omitempty"`
}

// EstimateTimeToComplete estimate time to complete
func (c *Client) EstimateTimeToComplete(ctx context.Context, req *Task) (*ETCResponse, error) {
	pathStr := "/v1/estimate-time"
	method := "GET"
	var res ETCResponse
	if err := c.call(ctx, method, pathStr, req, &res); err != nil {
		return nil, errors.Wrapf(err, "%s:%s failed", method, pathStr)
	}
	if res.StatusCode != SuccessStatusCode {
		return nil, errors.Errorf("StatusCode: %d", res.StatusCode)
	}
	return &res, nil
}

// UpdateUserRequest update ETC user request
type UpdateUserRequest struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

// UpdateUserResponse update ETC user response
type UpdateUserResponse struct {
	ID         string `json:"id"`
	UserName   string `json:"userName"`
	Email      string `json:"email"`
	StatusCode int    `json:"statusCode"`
}

// UpdateUser update ETC user info
func (c *Client) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error) {
	pathStr := "/v1/users"
	method := "PATCH"
	var res UpdateUserResponse
	if err := c.call(ctx, method, pathStr, req, &res); err != nil {
		return nil, errors.Wrapf(err, "%s:%s failed", method, pathStr)
	}
	if res.StatusCode != SuccessStatusCode {
		return nil, errors.Errorf("StatusCode: %d", res.StatusCode)
	}
	return &res, nil
}
