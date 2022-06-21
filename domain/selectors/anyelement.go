package selectors

type anyElement struct {
	isSelected bool
	prefix Name
}

func createAnyElement(
	isSelected bool,
	prefix Name,
) AnyElement {
	out := anyElement{
		isSelected: isSelected,
		prefix:    prefix,
	}

	return &out
}

// IsSelected returns true if selected, false otherwise
func (obj *anyElement) IsSelected() bool {
	return obj.isSelected
}

// Prefix returns the prefix
func (obj *anyElement) Prefix() Name {
	return obj.prefix
}
