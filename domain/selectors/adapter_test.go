package selectors

import (
	"testing"
)

func TestSelectorAdapter_isName_isNotSelected_Success(t *testing.T) {
	script := `
		.myToken
	`

	adapter := NewAdapter()
	selector, err := adapter.ToSelector(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	elements := selector.List()
	if len(elements) != 1 {
		t.Errorf("the selector was expecting %d elements, %d returned", 1, len(elements))
		return
	}

	if elements[0].IsAny() {
		t.Errorf("the element was expected to NOT contain an any")
		return
	}

	name := elements[0].Name()
	if name.Name() != "myToken" {
		t.Errorf("the name was expected to be '%s', '%s' returned", "myToken", name)
		return
	}

	if name.HasInsideNames() {
		t.Errorf("the name was expected to NOT contain inside names")
		return
	}

	if name.IsSelected() {
		t.Errorf("the name was NOT expecting to be selected")
		return
	}
}

func TestSelectorAdapter_isName_isSelected_Success(t *testing.T) {
	script := `
		+ .myToken
	`

	adapter := NewAdapter()
	selector, err := adapter.ToSelector(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	elements := selector.List()
	if len(elements) != 1 {
		t.Errorf("the selector was expecting %d elements, %d returned", 1, len(elements))
		return
	}

	if elements[0].IsAny() {
		t.Errorf("the element was expected to NOT contain an any")
		return
	}

	name := elements[0].Name()
	if name.Name() != "myToken" {
		t.Errorf("the name was expected to be '%s', '%s' returned", "myToken", name)
		return
	}

	if name.HasInsideNames() {
		t.Errorf("the name was expected to NOT contain inside names")
		return
	}

	if !name.IsSelected() {
		t.Errorf("the name was expected to be selected")
		return
	}
}

func TestSelectorAdapter_isName_isSelected_withInsideNames_Success(t *testing.T) {
	script := `
		+ @firstInside @secondInside .myToken
	`

	adapter := NewAdapter()
	selector, err := adapter.ToSelector(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	elements := selector.List()
	if len(elements) != 1 {
		t.Errorf("the selector was expecting %d elements, %d returned", 1, len(elements))
		return
	}

	if elements[0].IsAny() {
		t.Errorf("the element was expected to NOT contain an any")
		return
	}

	name := elements[0].Name()
	if name.Name() != "myToken" {
		t.Errorf("the name was expected to be '%s', '%s' returned", "myToken", name)
		return
	}

	if !name.HasInsideNames() {
		t.Errorf("the name was expected to contain inside names")
		return
	}

	if !name.IsSelected() {
		t.Errorf("the name was expected to be selected")
		return
	}

	insideNames := name.InsideNames()
	if len(insideNames) != 2 {
		t.Errorf("%d inside names were expected, %d returned", 2, len(insideNames))
		return
	}

	if insideNames[0] != "firstInside" {
		t.Errorf("the first insideName was expected to be '%s', '%s' returned", "firstInside", insideNames[0])
		return
	}

	if insideNames[1] != "secondInside" {
		t.Errorf("the first insideName was expected to be '%s', '%s' returned", "secondInside", insideNames[1])
		return
	}
}

func TestSelectorAdapter_isAny_withPrefix_isSelect_Success(t *testing.T) {
	script := `
		@firstInside @secondInside .myToken +*
	`

	adapter := NewAdapter()
	selector, err := adapter.ToSelector(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	elements := selector.List()
	if len(elements) != 1 {
		t.Errorf("the selector was expecting %d elements, %d returned", 1, len(elements))
		return
	}

	if !elements[0].IsAny() {
		t.Errorf("the element was expected to contain an Any")
		return
	}

	if elements[0].IsName() {
		t.Errorf("the element was expected to NOT contain a name")
		return
	}

	any := elements[0].Any()
	if !any.IsSelected() {
		t.Errorf("the any was expecting to be selected")
		return
	}

	prefix := any.Prefix().Name()
	if prefix != "myToken" {
		t.Errorf("the prefix was expected to be '%s', '%s' returned", "myToken", prefix)
		return
	}
}

func TestSelectorAdapter_isAny_withPrefix_isNOTSelect_Success(t *testing.T) {
	script := `
		@firstInside @secondInside .myToken *
	`

	adapter := NewAdapter()
	selector, err := adapter.ToSelector(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	elements := selector.List()
	if len(elements) != 1 {
		t.Errorf("the selector was expecting %d elements, %d returned", 1, len(elements))
		return
	}

	if !elements[0].IsAny() {
		t.Errorf("the element was expected to contain an Any")
		return
	}

	if elements[0].IsName() {
		t.Errorf("the element was expected to NOT contain a name")
		return
	}

	any := elements[0].Any()
	if any.IsSelected() {
		t.Errorf("the any was expecting to NOT be selected")
		return
	}

	prefix := any.Prefix().Name()
	if prefix != "myToken" {
		t.Errorf("the prefix was expected to be '%s', '%s' returned", "myToken", prefix)
		return
	}
}
