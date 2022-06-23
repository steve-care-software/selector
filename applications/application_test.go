package applications

import (
	"bytes"
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

	application := NewApplication()
	selectorIns, err := application.Compile(selector)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	bytes, err := NewApplication().Execute(selectorIns, result)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(bytes) != 2 {
		t.Errorf("%d bytes were expected, %d returned", 2, len(bytes))
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

	application := NewApplication()
	selectorIns, err := application.Compile(selector)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBytes, err := NewApplication().Execute(selectorIns, result)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := [][]byte{
		[]byte(" < 5"),
		[]byte{},
	}

	if len(retBytes) != len(expected) {
		t.Errorf("%d elements were expected, %d returned", len(expected), len(retBytes))
		return
	}

	for idx, data := range retBytes {
		if bytes.Compare(data, expected[idx]) != 0 {
			t.Errorf("%v bytes  were expected, %v returned at index: %d", expected[idx], data, idx)
			return
		}
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

	application := NewApplication()
	selectorIns, err := application.Compile(selector)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBytes, err := NewApplication().Execute(selectorIns, result)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := [][]byte{
		[]byte(" 5"),
	}

	if len(retBytes) != len(expected) {
		t.Errorf("%d elements were expected, %d returned", len(expected), len(retBytes))
		return
	}

	for idx, data := range retBytes {
		if bytes.Compare(data, expected[idx]) != 0 {
			t.Errorf("%v bytes  were expected, %v returned at index: %d", expected[idx], data, idx)
			return
		}
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

	data := []byte("(( 5 < 5 ))")
	result, err := validator.NewApplication().Execute(schema, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	selector := `
		+ @rootToken @rootToken .smallerThan *
	`

	application := NewApplication()
	selectorIns, err := application.Compile(selector)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBytes, err := NewApplication().Execute(selectorIns, result)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := [][]byte{
		[]byte(" 5 ))"),
	}

	if len(retBytes) != len(expected) {
		t.Errorf("%d elements were expected, %d returned", len(expected), len(retBytes))
		return
	}

	for idx, data := range retBytes {
		if bytes.Compare(data, expected[idx]) != 0 {
			t.Errorf("%v bytes  were expected, %v returned at index: %d", expected[idx], data, idx)
			return
		}
	}
}

func TestSelector_withBytes_isSuccess(t *testing.T) {
	schema := `
		%rootToken;
		-space;
		-endOfLine;

		rootToken: .bytes
				 | .byte
				 ;

		bytes: .openSquareBracket .byteWithSemiColon[1,] .closeSquareBracket;
		byteWithSemiColon: .byte .semiColon;
		byte: .dollar .number[1,3];

		number: .zero
			  | .one
			  | .two
			  | .three
			  | .four
			  | .five
			  | .six
			  | .seven
			  | .height
			  | .nine
			  ;

		openSquareBracket: $91;
		closeSquareBracket: $93;
		semiColon: $59;
		dollar: $36;
		zero: $48;
		one: $49;
		two: $50;
		three: $51;
		four: $52;
		five: $53;
		six: $54;
		seven: $55;
		height: $56;
		nine: $57;
		space: $32;
		endOfLine: $10;
	`

	data := []byte("[$100; $20; $30;]")
	result, err := validator.NewApplication().Execute(schema, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	selector := `
		+ @rootToken @bytes @byteWithSemiColon @byte .number
	`

	application := NewApplication()
	selectorIns, err := application.Compile(selector)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBytes, err := NewApplication().Execute(selectorIns, result)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := [][]byte{
		[]byte("100"),
		[]byte("20"),
		[]byte("30"),
	}

	if len(retBytes) != len(expected) {
		t.Errorf("%d elements were expected, %d returned", len(expected), len(retBytes))
		return
	}

	for idx, data := range retBytes {
		if bytes.Compare(data, expected[idx]) != 0 {
			t.Errorf("%v bytes  were expected, %v returned at index: %d", expected[idx], data, idx)
			return
		}
	}
}

func TestSelector_withByte_isSuccess(t *testing.T) {
	schema := `
		%rootToken;
		-space;
		-endOfLine;

		rootToken: .bytes
				 | .byte
				 ;

		bytes: .openSquareBracket .byteWithSemiColon[1,] .closeSquareBracket;
		byteWithSemiColon: .byte .semiColon;
		byte: .dollar .number[1,3];

		number: .zero
			  | .one
			  | .two
			  | .three
			  | .four
			  | .five
			  | .six
			  | .seven
			  | .height
			  | .nine
			  ;

		openSquareBracket: $91;
		closeSquareBracket: $93;
		semiColon: $59;
		dollar: $36;
		zero: $48;
		one: $49;
		two: $50;
		three: $51;
		four: $52;
		five: $53;
		six: $54;
		seven: $55;
		height: $56;
		nine: $57;
		space: $32;
		endOfLine: $10;
	`

	data := []byte("$100")
	result, err := validator.NewApplication().Execute(schema, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	selector := `
		+ @rootToken @byte .number
	`

	application := NewApplication()
	selectorIns, err := application.Compile(selector)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBytes, err := NewApplication().Execute(selectorIns, result)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := [][]byte{
		[]byte("100"),
	}

	for idx, data := range retBytes {
		if bytes.Compare(data, expected[idx]) != 0 {
			t.Errorf("%v bytes  were expected, %v returned at index: %d", expected[idx], data, idx)
			return
		}
	}
}
