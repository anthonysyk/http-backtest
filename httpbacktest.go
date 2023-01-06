package httpbacktest

import (
	"fmt"
	"github.com/anthonysyk/http-backtest/format"
	"github.com/dariubs/percent"
	"github.com/go-resty/resty/v2"
)

type Env struct {
	Host    string
	Headers map[string]string
}

func (c *Client) Run(name string, URLs []string, envA, envB Env, comparator format.Comparator) *Result {
	uniqueURLs := make(map[string]int)
	for _, URL := range URLs {
		uniqueURLs[URL]++
	}

	result := &Result{
		Name:           name,
		TotalRequests:  len(URLs),
		BodyEquivalent: make(map[string]int),
		EnvironmentDetailsA: EnvironmentDetails{
			StatusCodes: make(map[int]int),
		},
		EnvironmentDetailsB: EnvironmentDetails{
			StatusCodes: make(map[int]int),
		},
	}
	for k := range uniqueURLs {
		c.rl.Take()
		result.UniqueURLs++
		responseA, errA := c.Call(envA, k)
		if errA != nil {
			result.EnvironmentDetailsA.Errors = append(result.EnvironmentDetailsA.Errors, errA)
		}
		responseB, errB := c.Call(envB, k)
		if errB != nil {
			result.EnvironmentDetailsB.Errors = append(result.EnvironmentDetailsA.Errors, errB)
		}
		if responseA != nil && responseB != nil {
			compareStatusCodes(responseA, responseB, result)
			bodyMatched, comment := comparator.Compare(responseA, responseB)
			result.BodyEquivalent[comment]++
			if bodyMatched {
				result.BodyMatched++
			} else {
				result.BodyNoMatched++
			}
		}
	}

	result.StatusCodeSimilarity = fmt.Sprintf("%v%%", percent.PercentOf(result.StatusMatched, result.UniqueURLs))
	result.BodySimilarity = fmt.Sprintf("%v%%", percent.PercentOf(result.BodyMatched, result.UniqueURLs))
	return result
}

func (c *Client) Call(env Env, URL string) (*resty.Response, error) {
	c.rl.Take()
	res, err := c.http.R().SetHeaders(env.Headers).Get(fmt.Sprintf("%s%s", env.Host, URL))
	if err != nil {
		return nil, err
	}
	return res, nil
}

type FinalResult struct {
	TotalRequests         int      `json:"totalRequests"`
	TotalUniqueURLs       int      `json:"totalUniqueURLs"`
	TotalStatusMatched    int      `json:"totalStatusMatched"`
	TotalStatusNoMatched  int      `json:"totalStatusNoMatched"`
	TotalStatusSimilarity string   `json:"totalStatusSimilarity"`
	TotalBodyMatched      int      `json:"totalBodyMatched"`
	TotalBodyNoMatched    int      `json:"totalBodyNoMatched"`
	TotalBodySimilarity   string   `json:"totalBodySimilarity"`
	Details               []Result `json:"details"`
}

func (fr *FinalResult) ProcessTotal() {
	for _, d := range fr.Details {
		fr.TotalRequests += d.TotalRequests
		fr.TotalBodyNoMatched += d.BodyNoMatched
		fr.TotalBodyMatched += d.BodyMatched
		fr.TotalStatusMatched += d.StatusMatched
		fr.TotalStatusNoMatched += d.StatusNoMatched
		fr.TotalUniqueURLs += d.UniqueURLs
	}

	fr.TotalBodySimilarity = fmt.Sprintf("%v%%", percent.PercentOf(fr.TotalBodyMatched, fr.TotalUniqueURLs))
	fr.TotalStatusSimilarity = fmt.Sprintf("%v%%", percent.PercentOf(fr.TotalStatusMatched, fr.TotalUniqueURLs))
}
