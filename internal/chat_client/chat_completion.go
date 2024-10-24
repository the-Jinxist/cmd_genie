package chat_client

import (
	"net/http"
	"runtime"
)

type ChatCompletion struct {
	client *client
}

// newSubscription returns a new Subscription.
func newChatCompletion(client *client) *ChatCompletion {
	return &ChatCompletion{
		client: client,
	}
}

type CompletionRequest struct {
	Prompt string `json:"prompt"`
	Os     string `json:"os"`
}

type CompletionResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (c ChatCompletion) GetChatCompletion(prompt string) (*CompletionResponse, error) {

	os := runtime.GOOS

	body := CompletionRequest{
		Prompt: prompt,
		Os:     os,
	}
	res := new(CompletionResponse)
	err := c.client.req(http.MethodPost, "user/cmd_completion", body, res)

	return res, err
}
