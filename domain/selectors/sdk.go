package selectors

// NewAdapter creates a new selector adapter instance
func NewAdapter() Adapter {
	selectorBuilder := NewBuilder()
	elementBuilder := NewElementBuilder()
	anyElementBuilder := NewAnyElementBuilder()
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
		elementBuilder,
		anyElementBuilder,
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

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewAnyElementBuilder creates a new anyElement builder
func NewAnyElementBuilder() AnyElementBuilder {
	return createAnyElementBuilder()
}

// NewNameBuilder creates a new name builder
func NewNameBuilder() NameBuilder {
	return createNameBuilder()
}

// Adapter represents the selector adapter
type Adapter interface {
	ToScript(selector Selector) []byte
	ToSelector(script string) (Selector, error)
}

// Builder represents a selector builder
type Builder interface {
	Create() Builder
	WithList(list []Element) Builder
	Now() (Selector, error)
}

// Selector represents a selector
type Selector interface {
	List() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithName(name Name) ElementBuilder
	WithAny(any AnyElement) ElementBuilder
	Now() (Element, error)
}

// Element represents a selector element
type Element interface {
	IsName() bool
	Name() Name
	IsAny() bool
	Any() AnyElement
}

// AnyElementBuilder represents an any element builder
type AnyElementBuilder interface {
	Create() AnyElementBuilder
	IsSelected() AnyElementBuilder
	WithPrefix(prefix Name) AnyElementBuilder
	Now() (AnyElement, error)
}

// AnyElement represents an any element
type AnyElement interface {
	IsSelected() bool
	Prefix() Name
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
