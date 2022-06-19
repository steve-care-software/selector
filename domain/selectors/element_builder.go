package selectors

import "errors"

type elementBuilder struct {
	name Name
	any  AnyElement
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		name: nil,
		any:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithName adds a name to the builder
func (app *elementBuilder) WithName(name Name) ElementBuilder {
	app.name = name
	return app
}

// WithAny adds an anyElement to the builder
func (app *elementBuilder) WithAny(any AnyElement) ElementBuilder {
	app.any = any
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.name != nil {
		return createElementWithName(app.name), nil
	}

	if app.any != nil {
		return createElementWithAnyElement(app.any), nil
	}

	return nil, errors.New("the Element is invalid")
}
