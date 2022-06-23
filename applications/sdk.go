package applications

import (
	"github.com/steve-care-software/selector/domain/selectors"
	"github.com/steve-care-software/validator/domain/results"
)

// NewApplication creates a new application instance
func NewApplication() Application {
	adapter := selectors.NewAdapter()
	return createApplication(adapter)
}

// Application represents the selector application
type Application interface {
	Compile(script string) (selectors.Selector, error)
	Execute(selector selectors.Selector, result results.Result) ([][]byte, error)
}
