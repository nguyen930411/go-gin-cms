package app

import (
	"github.com/astaxie/beego/validation"

	"github.com/nguyen930411/go-gin-cms/pkg/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
