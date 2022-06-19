package selectors

import "errors"

type nameBuilder struct {
	isSelected  bool
	name        string
	insideNames []string
}

func createNameBuilder() NameBuilder {
	out := nameBuilder{
		isSelected:  false,
		name:        "",
		insideNames: nil,
	}

	return &out
}

// Create initializes the builder
func (app *nameBuilder) Create() NameBuilder {
	return createNameBuilder()
}

// IsSelected flags the builder as selected
func (app *nameBuilder) IsSelected() NameBuilder {
	app.isSelected = true
	return app
}

// WithName adds a name to the builder
func (app *nameBuilder) WithName(name string) NameBuilder {
	app.name = name
	return app
}

// WithInsideNames add insideNames to the builder
func (app *nameBuilder) WithInsideNames(insideNames []string) NameBuilder {
	app.insideNames = insideNames
	return app
}

// Now builds a new Name instance
func (app *nameBuilder) Now() (Name, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a NameWithDelimiter instance")
	}

	if app.insideNames != nil && len(app.insideNames) <= 0 {
		app.insideNames = nil
	}

	if app.insideNames != nil {
		return createNameWithInsideNames(app.isSelected, app.name, app.insideNames), nil
	}

	return createName(app.isSelected, app.name), nil
}
