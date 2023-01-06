package example

import (
	_ "embed"
	"fmt"
	"github.com/anthonysyk/http-backtest"
	"github.com/anthonysyk/http-backtest/format"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	servermock "github.com/anthonysyk/http-server-mock"
)

//go:embed server_a.yaml
var ServerA string

//go:embed server_b.yaml
var ServerB string

func TestHttpBacktest(t *testing.T) {
	client := httpbacktest.NewHttpBacktestClient(10)
	_ = client

	routes, err := servermock.GetRoutes(ServerA)
	assert.NoError(t, err)
	serverA, err := servermock.GenerateRouter(ServerA)
	assert.NoError(t, err)
	serverB, err := servermock.GenerateRouter(ServerB)
	assert.NoError(t, err)

	go http.ListenAndServe(":8081", serverA)
	go http.ListenAndServe(":8082", serverB)

	var urls []string
	for _, r := range routes {
		urls = append(urls, r.URL)
	}
	envA := httpbacktest.Env{Host: "http://localhost:8081"}
	envB := httpbacktest.Env{Host: "http://localhost:8082"}
	res := client.Run("movies-backtest", urls, envA, envB, format.JSON{})
	fmt.Println(res.JSON())
}
