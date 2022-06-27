package selectors

type element struct {
	name Name
	any  Name
}

func createElementWithName(
	name Name,
) Element {
	return createElementInternally(name, nil)
}

func createElementWithAnyElement(
	any Name,
) Element {
	return createElementInternally(nil, any)
}

func createElementInternally(
	name Name,
	any Name,
) Element {
	out := element{
		name: name,
		any:  any,
	}

	return &out
}

// IsName returns true if name, false otherwise
func (obj *element) IsName() bool {
	return obj.name != nil
}

// Name returns the name, if any
func (obj *element) Name() Name {
	return obj.name
}

// IsAny returns true if any, false otherwise
func (obj *element) IsAny() bool {
	return obj.any != nil
}

// Any returns the anyElement, if any
func (obj *element) Any() Name {
	return obj.any
}
