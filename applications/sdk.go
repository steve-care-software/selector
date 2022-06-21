package applications

import (
	"github.com/steve-care-software/validator/domain/results"
	"github.com/steve-care-software/selector/domain/selectors"
)

// NewApplication creates a new application instance
func NewApplication() Application {
	adapter := selectors.NewAdapter()
	return createApplication(adapter)
}

// Application represents the selector application
type Application interface {
    Execute(selector string, result results.Result) ([]byte, error)
}
