package httpbacktest

import (
	"bytes"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/nsf/jsondiff"
	"net/http"
)

func compareJSON(responseA, responseB *resty.Response, result *Result) {
	if responseA.StatusCode() == http.StatusOK && responseB.StatusCode() == http.StatusOK {
		bodyA := new(bytes.Buffer)
		_ = json.Compact(bodyA, responseA.Body())

		bodyB := new(bytes.Buffer)
		_ = json.Compact(bodyB, responseB.Body())

		opts := jsondiff.DefaultJSONOptions()
		diff, _ := jsondiff.Compare(bodyA.Bytes(), bodyB.Bytes(), &opts)

		result.BodyEquivalent[diff.String()]++

		if diff == jsondiff.FullMatch {
			result.BodyMatched++
		} else {
			result.BodyNoMatched++
		}
	} else {
		result.BodyMatched++
	}
}
