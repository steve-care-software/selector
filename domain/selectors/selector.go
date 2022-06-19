package selectors

type selector struct {
	list []Element
}

func createSelector(
	list []Element,
) Selector {
	out := selector{
		list: list,
	}

	return &out
}

// List returns the list of element
func (obj *selector) List() []Element {
	return obj.list
}
