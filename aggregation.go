package httpbacktest

import "github.com/anthonysyk/http-backtest/format"

type Iteration struct {
	Name   string
	URLs   []string
	envA   Env
	envB   Env
	format format.Comparator
}

func Aggregate(results []Result) FinalResult {
	final := FinalResult{Details: results}
	final.ProcessTotal()
	return final
}

func (c *Client) RunWithAggregate(iterations []Iteration) FinalResult {
	var results []Result
	for _, i := range iterations {
		result := c.Run(i.Name, i.URLs, i.envA, i.envB, i.format)
		results = append(results, *result)
	}
	return Aggregate(results)
}
