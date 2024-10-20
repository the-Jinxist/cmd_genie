package chat_client

import (
	"fmt"
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

type GeminiRequest struct {
	Contents []Contents `json:"contents"`
}
type Parts struct {
	Text string `json:"text"`
}
type Contents struct {
	Parts []Parts `json:"parts"`
}

type GeminiResponse struct {
	Candidates    []Candidates  `json:"candidates"`
	UsageMetadata UsageMetadata `json:"usageMetadata"`
}

type Content struct {
	Parts []Parts `json:"parts"`
	Role  string  `json:"role"`
}
type SafetyRatings struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}
type Candidates struct {
	Content       Content         `json:"content"`
	FinishReason  string          `json:"finishReason"`
	Index         int             `json:"index"`
	SafetyRatings []SafetyRatings `json:"safetyRatings"`
}
type UsageMetadata struct {
	PromptTokenCount     int `json:"promptTokenCount"`
	CandidatesTokenCount int `json:"candidatesTokenCount"`
	TotalTokenCount      int `json:"totalTokenCount"`
}

func (c ChatCompletion) GetChatCompletion(prompt string) (*GeminiResponse, error) {

	os := runtime.GOOS
	primer := fmt.Sprintf("You're a CLI savant with experience in naviagting CLIs on all operating system. Please provide succintly the cli command related to what the prompt needs. The OS is %s", os)

	url := fmt.Sprintf("v1beta/models/gemini-1.5-flash-latest:generateContent?key=%s", c.client.APIKey)
	body := GeminiRequest{
		Contents: []Contents{
			{
				Parts: []Parts{
					{
						Text: prompt,
					},
					{
						Text: primer,
					},
				},
			},
		},
	}
	res := new(GeminiResponse)
	err := c.client.req(http.MethodPost, url, body, res)

	return res, err
}
