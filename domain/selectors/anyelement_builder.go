package selectors

type anyElementBuilder struct {
	isSelected bool
	prefix     Name
	suffix     Name
}

func createAnyElementBuilder() AnyElementBuilder {
	out := anyElementBuilder{
		isSelected: false,
		prefix:     nil,
		suffix:     nil,
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

// WithSuffix adds a suffix to the builder
func (app *anyElementBuilder) WithSuffix(suffix Name) AnyElementBuilder {
	app.suffix = suffix
	return app
}

// Now builds a new AnyElement instance
func (app *anyElementBuilder) Now() (AnyElement, error) {
	if app.prefix != nil && app.suffix != nil {
		content := createContentWithPrefixAndSuffix(app.prefix, app.suffix)
		return createAnyElement(app.isSelected, content), nil
	}

	if app.prefix != nil {
		content := createContentWithPrefix(app.prefix)
		return createAnyElement(app.isSelected, content), nil
	}

	if app.suffix != nil {
		content := createContentWithSuffix(app.suffix)
		return createAnyElement(app.isSelected, content), nil
	}

	content := createContent()
	return createAnyElement(app.isSelected, content), nil
}
