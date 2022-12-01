package httpbacktest

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/ratelimit"
)

type BodyFormat string

const (
	BodyFormatJSON = "json"
	BodyFormatXML  = "xml"
	BodyFormatHTML = "html"
	BodyFormatCSV  = "csv"
)

type Environment struct {
	Name    string
	Headers map[string]string
}

type Client struct {
	http *resty.Client
	rl   ratelimit.Limiter
	envs []Environment
	body BodyFormat
}

func NewHttpBacktestClient(requestPerSecond int) *Client {
	return &Client{
		http: resty.New(),
		rl:   ratelimit.New(requestPerSecond),
	}
}

type Result struct {
	Name                 string             `json:"name"`
	TotalRequests        int                `json:"totalRequests"`
	UniqueURLs           int                `json:"uniqueURLs"`
	StatusMatched        int                `json:"statusMatched"`
	StatusNoMatched      int                `json:"statusNoMatched"`
	StatusCodeSimilarity string             `json:"statusCodeSimilarity"`
	BodyMatched          int                `json:"bodyMatched"`
	BodyNoMatched        int                `json:"bodyNoMatched"`
	BodySimilarity       string             `json:"bodySimilarity"`
	BodyEquivalent       map[string]int     `json:"bodyEquivalent"`
	EnvironmentDetailsA  EnvironmentDetails `json:"environmentDetailsA"`
	EnvironmentDetailsB  EnvironmentDetails `json:"environmentDetailsB"`
}

type EnvironmentDetails struct {
	Errors      []error     `json:"errors,omitempty"`
	StatusCodes map[int]int `json:"stagingStatusCodes"`
}
