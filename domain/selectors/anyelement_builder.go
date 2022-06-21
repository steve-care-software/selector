package selectors

import (
	"errors"
)

type anyElementBuilder struct {
	isSelected bool
	prefix     Name
}

func createAnyElementBuilder() AnyElementBuilder {
	out := anyElementBuilder{
		isSelected: false,
		prefix:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *anyElementBuilder) Create() AnyElementBuilder {
	return createAnyElementBuilder()
}

// IsSelected flags the builder as selected
func (app *anyElementBuilder) IsSelected() AnyElementBuilder {
	app.isSelected = true
	return app
}

// WithPrefix adds a prefix to the builder
func (app *anyElementBuilder) WithPrefix(prefix Name) AnyElementBuilder {
	app.prefix = prefix
	return app
}

// Now builds a new AnyElement instance
func (app *anyElementBuilder) Now() (AnyElement, error) {
	if app.prefix == nil {
		return nil, errors.New("the prefix is mandatory in order to build an AnyElement instance")
	}

	return createAnyElement(app.isSelected, app.prefix), nil
}
