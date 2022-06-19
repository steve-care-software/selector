package selectors

type anyElement struct {
	isSelected bool
	content    Content
}

func createAnyElement(
	isSelected bool,
	content Content,
) AnyElement {
	out := anyElement{
		isSelected: isSelected,
		content:    content,
	}

	return &out
}

// IsSelected returns true if selected, false otherwise
func (obj *anyElement) IsSelected() bool {
	return obj.isSelected
}

// Content returns the content
func (obj *anyElement) Content() Content {
	return obj.content
}
