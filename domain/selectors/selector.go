package selectors

type selector struct {
	name Name
	any  Name
}

func createSelectorWithName(
	name Name,
) Selector {
	return createSelectorInternally(name, nil)
}

func createSelectorWithAnySelector(
	any Name,
) Selector {
	return createSelectorInternally(nil, any)
}

func createSelectorInternally(
	name Name,
	any Name,
) Selector {
	out := selector{
		name: name,
		any:  any,
	}

	return &out
}

// IsName returns true if name, false otherwise
func (obj *selector) IsName() bool {
	return obj.name != nil
}

// Name returns the name, if any
func (obj *selector) Name() Name {
	return obj.name
}

// IsAny returns true if any, false otherwise
func (obj *selector) IsAny() bool {
	return obj.any != nil
}

// Any returns the anySelector, if any
func (obj *selector) Any() Name {
	return obj.any
}
