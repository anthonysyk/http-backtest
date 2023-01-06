package format

import (
	"github.com/go-resty/resty/v2"
)

type Comparator interface {
	Compare(responseA, responseB *resty.Response) (bool, string)
}

type Result struct {
	BodyMatched   bool
	BodyNoMatched bool
	Comment       string
}
