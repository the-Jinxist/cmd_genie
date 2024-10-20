package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Client struct {
	APIKey     string
	HttpClient *http.Client
	log        *slog.Logger
	BaseUrl    *url.URL

	ChatCompletion *ChatCompletion

	// Transaction      *Transaction
	// TransactionSplit *TransactionSplit
	// Plan             *Plans
	// Subscription     *Subscription
}

const BASE_URL = "https://generativelanguage.googleapis.com/"

func NewClient(apiKey string) *Client {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	parsedUrl, _ := url.Parse(BASE_URL)
	c := &Client{APIKey: apiKey, HttpClient: httpClient, log: logger, BaseUrl: parsedUrl}
	c.ChatCompletion = newChatCompletion(c)
	return c
}

func (c Client) req(method string, url string, body interface{}, response interface{}) error {

	var reqBody *bytes.Buffer
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return err
		}

		reqBody = bytes.NewBuffer(buf)
	}

	reqUrl, _ := c.BaseUrl.Parse(url)
	req, _ := http.NewRequest(method, reqUrl.String(), reqBody)

	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error processing request - %+v", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf("error while unmarshalling the response bytes %+v ", err)
	}

	return nil
}
