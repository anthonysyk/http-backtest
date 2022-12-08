package example

import (
	"fmt"
	"github.com/anthonysyk/http-backtest"
	"github.com/go-resty/resty/v2"
	"testing"
)

func TestHttpBacktest(t *testing.T) {
	client := httpbacktest.NewHttpBacktestClient(10)
	_ = client

	serverA := runServerWithSeed("123")
	serverB := runServerWithSeed("456")

	fmt.Println(serverA.URL)
	fmt.Println(serverB.URL)

	res, err := resty.New().R().Get(fmt.Sprintf("%s/search?query=test", serverA.URL))
	fmt.Println(res.String(), err)
}
