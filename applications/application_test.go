package applications

import (
	"testing"

	validator "github.com/steve-care-software/validator/applications"
)

func TestSelector_withInsideNames_withName_isSuccess(t *testing.T) {
	schema := `
		%rootToken;
		-space;
		-endOfLine;

		rootToken : .five .smallerThan .five
				  ;

		openParenthesis: $40;
		closeParenthesis: $41;
		five: $53;
		smallerThan: $60;
		space: $32;
		endOfLine: $10;
	`

	data := []byte("5 < 5")
	result, err := validator.NewApplication().Execute(schema, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	selector := `
		+ @rootToken .five
	`

	bytes, err := NewApplication().Execute(selector, result)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(bytes) != 1 {
		t.Errorf("%d bytes were expected, %d returned", 1, len(bytes))
		return
	}
}

func TestSelector_withAnyElement_withPrefix_withSuffix_isSelected_isSuccess(t *testing.T) {
	schema := `
		%rootToken;
		-space;
		-endOfLine;

		rootToken : .five .smallerThan .five
				  ;

		openParenthesis: $40;
		closeParenthesis: $41;
		five: $53;
		smallerThan: $60;
		space: $32;
		endOfLine: $10;
	`

	data := []byte("5 < 5")
	result, err := validator.NewApplication().Execute(schema, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	selector := `
		+ @rootToken .five *
	`

	bytes, err := NewApplication().Execute(selector, result)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := " < 5"
	if string(bytes) != expected {
		t.Errorf("'%s' was expected, '%s' returned", expected, string(bytes))
		return
	}
}

func TestSelector_afterSmallerThan_withAnyElement_withPrefix_withSuffix_isSelected_isSuccess(t *testing.T) {
	schema := `
		%rootToken;
		-space;
		-endOfLine;

		rootToken : .five .smallerThan .five
				  ;

		openParenthesis: $40;
		closeParenthesis: $41;
		five: $53;
		smallerThan: $60;
		space: $32;
		endOfLine: $10;
	`

	data := []byte("5 < 5")
	result, err := validator.NewApplication().Execute(schema, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	selector := `
		+ @rootToken .smallerThan *
	`

	bytes, err := NewApplication().Execute(selector, result)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := " 5"
	if string(bytes) != expected {
		t.Errorf("'%s' was expected, '%s' returned", expected, string(bytes))
		return
	}
}

func TestSelector_withRecursiveToken_afterSmallerThan_withAnyElement_withPrefix_withSuffix_isSelected_isSuccess(t *testing.T) {
	schema := `
		%rootToken;
		-space;
		-endOfLine;

		rootToken : .openParenthesis .rootToken .closeParenthesis
				  | .five .smallerThan .five
				  ;

		openParenthesis: $40;
		closeParenthesis: $41;
		five: $53;
		smallerThan: $60;
		space: $32;
		endOfLine: $10;
	`

	data := []byte("( 5 < 5 )")
	result, err := validator.NewApplication().Execute(schema, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	selector := `
		+ @rootToken @rootToken .smallerThan *
	`

	bytes, err := NewApplication().Execute(selector, result)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := " 5 )"
	if string(bytes) != expected {
		t.Errorf("'%s' was expected, '%s' returned", expected, string(bytes))
		return
	}
}
