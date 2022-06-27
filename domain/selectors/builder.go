package selectors

import "errors"

type builder struct {
	name Name
	any  Name
}

func createBuilder() Builder {
	out := builder{
		name: nil,
		any:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name Name) Builder {
	app.name = name
	return app
}

// WithAny adds an anySelector to the builder
func (app *builder) WithAny(any Name) Builder {
	app.any = any
	return app
}

// Now builds a new Selector instance
func (app *builder) Now() (Selector, error) {
	if app.name != nil {
		return createSelectorWithName(app.name), nil
	}

	if app.any != nil {
		return createSelectorWithAnySelector(app.any), nil
	}

	return nil, errors.New("the Selector is invalid")
}
