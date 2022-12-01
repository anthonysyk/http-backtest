package httpbacktest

import "github.com/go-resty/resty/v2"

func compareStatusCodes(responseA, responseB *resty.Response, result *Result) {
	result.EnvironmentDetailsA.StatusCodes[responseA.StatusCode()]++
	result.EnvironmentDetailsB.StatusCodes[responseB.StatusCode()]++
	if responseA.StatusCode() == responseB.StatusCode() {
		result.StatusMatched++
	} else {
		result.StatusNoMatched++
	}
}
