package selectors

type content struct {
	prefix Name
	suffix Name
}

func createContent() Content {
	return createContentInternally(nil, nil)
}

func createContentWithPrefix(
	prefix Name,
) Content {
	return createContentInternally(prefix, nil)
}

func createContentWithSuffix(
	suffix Name,
) Content {
	return createContentInternally(nil, suffix)
}

func createContentWithPrefixAndSuffix(
	prefix Name,
	suffix Name,
) Content {
	return createContentInternally(prefix, suffix)
}

func createContentInternally(
	prefix Name,
	suffix Name,
) Content {
	out := content{
		prefix: prefix,
		suffix: suffix,
	}

	return &out
}

// HasPrefix returns true if there is a prefix, false otherwise
func (obj *content) HasPrefix() bool {
	return obj.prefix != nil
}

// Prefix returns the prefix, if any
func (obj *content) Prefix() Name {
	return obj.prefix
}

// HasSuffix returns true if there is a suffix, false otherwise
func (obj *content) HasSuffix() bool {
	return obj.suffix != nil
}

// Suffix returns the suffix, if any
func (obj *content) Suffix() Name {
	return obj.suffix
}
