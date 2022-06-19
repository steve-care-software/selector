package selectors

type name struct {
	isSelected  bool
	name        string
	insideNames []string
}

func createName(
	isSelected bool,
	name string,
) Name {
	return createNameInternally(isSelected, name, nil)
}

func createNameWithInsideNames(
	isSelected bool,
	name string,
	insideNames []string,
) Name {
	return createNameInternally(isSelected, name, insideNames)
}

func createNameInternally(
	isSelected bool,
	nameStr string,
	insideNames []string,
) Name {
	out := name{
		isSelected:  isSelected,
		name:        nameStr,
		insideNames: insideNames,
	}

	return &out
}

// IsSelected returns true if selected, false otherwise
func (obj *name) IsSelected() bool {
	return obj.isSelected
}

// Name returns the name
func (obj *name) Name() string {
	return obj.name
}

// HasInsideNames returns true if there is an insideNames, false otherwise
func (obj *name) HasInsideNames() bool {
	return obj.insideNames != nil
}

// InsideNames returns the insideNames, if any
func (obj *name) InsideNames() []string {
	return obj.insideNames
}
