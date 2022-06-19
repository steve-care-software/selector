package applications

import (
	"github.com/steve-care-software/validator/domain/results"
)

// Application represents the selector application
type Application interface {
    Execute(selector string, result results.Result) ([]byte, error)
}
