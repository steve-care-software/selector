package selectors

// NewAdapter creates a new selector adapter instance
func NewAdapter() Adapter {
	selectorBuilder := NewBuilder()
	nameBuilder := NewNameBuilder()
	anyByte := []byte("*")[0]
	tokenNameByte := []byte(".")[0]
	insideByte := []byte("@")[0]
	selectByte := []byte("+")[0]
	tokenNameCharacters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	channelCharacters := []byte{
		[]byte("\t")[0],
		[]byte("\n")[0],
		[]byte("\r")[0],
		[]byte(" ")[0],
	}

	return createAdapter(
		selectorBuilder,
		nameBuilder,
		anyByte,
		tokenNameByte,
		insideByte,
		selectByte,
		tokenNameCharacters,
		channelCharacters,
	)
}

// NewBuilder creates a new selector builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewNameBuilder creates a new name builder
func NewNameBuilder() NameBuilder {
	return createNameBuilder()
}

// Adapter represents the selector adapter
type Adapter interface {
	ToScript(selector Selector) []byte
	ToSelector(script string) (Selector, []byte, error)
}

// Builder represents a selector builder
type Builder interface {
	Create() Builder
	WithName(name Name) Builder
	WithAny(any Name) Builder
	Now() (Selector, error)
}

// Selector represents a selector element
type Selector interface {
	IsName() bool
	Name() Name
	IsAny() bool
	Any() Name
}

// NameBuilder represents a name builder
type NameBuilder interface {
	Create() NameBuilder
	IsSelected() NameBuilder
	WithName(name string) NameBuilder
	WithInsideNames(insideNames []string) NameBuilder
	Now() (Name, error)
}

// Name represents a name
type Name interface {
	Name() string
	IsSelected() bool
	HasInsideNames() bool
	InsideNames() []string
}
