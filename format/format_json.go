package format

import (
	"bytes"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/nsf/jsondiff"
	"net/http"
)

type JSON struct{}

func (f JSON) Compare(responseA, responseB *resty.Response) (bool, string) {
	if responseA.StatusCode() == http.StatusOK && responseB.StatusCode() == http.StatusOK {
		bodyA := new(bytes.Buffer)
		_ = json.Compact(bodyA, responseA.Body())

		bodyB := new(bytes.Buffer)
		_ = json.Compact(bodyB, responseB.Body())

		opts := jsondiff.DefaultJSONOptions()
		diff, _ := jsondiff.Compare(bodyA.Bytes(), bodyB.Bytes(), &opts)
		return diff == jsondiff.FullMatch || diff == jsondiff.BothArgsAreInvalidJson, diff.String()
	}
	return true, "status code different"
}
