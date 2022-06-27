package applications

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/steve-care-software/selector/domain/selectors"
	"github.com/steve-care-software/validator/domain/results"
)

type application struct {
	adapter selectors.Adapter
}

func createApplication(
	adapter selectors.Adapter,
) Application {
	out := application{
		adapter: adapter,
	}

	return &out
}

// Compile compiles a selector
func (app *application) Compile(script string) (selectors.Selector, []byte, error) {
	return app.adapter.ToSelector(script)
}

// Execute executes a selector on validation result
func (app *application) Execute(selector selectors.Selector, result results.Result) ([][]byte, error) {
	if !result.Token().IsSuccess() {
		return nil, errors.New("the selector cannot extract result tokens because the result is invalid")
	}

	token := result.Token()
	return app.selectorOnToken(selector, token)
}

func (app *application) selectorOnToken(selector selectors.Selector, token results.Token) ([][]byte, error) {
	if selector.IsName() {
		name := selector.Name()
		return app.nameInsOnToken(name, token)
	}

	anyName := selector.Any()
	return app.anyNameOnToken(anyName, token)
}

func (app *application) nameInsOnToken(nameIns selectors.Name, token results.Token) ([][]byte, error) {
	name := nameIns.Name()
	path := []string{}
	if nameIns.HasInsideNames() {
		path = nameIns.InsideNames()
	}

	path = append(path, name)
	bytes, _, err := app.nameOnToken(path, token)
	if err != nil {
		return nil, err
	}

	if nameIns.IsSelected() {
		return bytes, nil
	}

	return nil, nil
}

func (app *application) nameOnToken(path []string, token results.Token) ([][]byte, bool, error) {
	block := token.Block()
	if !block.IsSuccess() {
		str := fmt.Sprintf("the block's token (name: %s) is NOT successful and therefore its value cannot be extracted", token.Name())
		return nil, false, errors.New(str)
	}

	if len(path) <= 0 {
		return nil, false, errors.New("the path is mandatory in order to retrieve the token's value, none provided")
	}

	name := path[0]
	currentPath := path
	if token.Name() == name {
		currentPath = path[1:]
	}

	isSameLine := false
	output := [][]byte{}
	lines := block.List()
	for _, oneLine := range lines {
		if !oneLine.IsSuccess() {
			continue
		}

		elements := oneLine.Elements()
		for _, oneElementWithCardinality := range elements {
			if !oneElementWithCardinality.IsSuccess() {
				continue
			}

			if !oneElementWithCardinality.HasMatches() {
				continue
			}

			data := []byte{}
			matches := oneElementWithCardinality.Matches()
			for _, oneElement := range matches {
				if oneElement.IsValue() {
					if len(currentPath) > 0 {
						continue
					}

					pValue := oneElement.Value()
					data = append(data, *pValue)
					isSameLine = true
					continue
				}

				if oneElement.IsToken() {
					elementToken := oneElement.Token()
					if len(currentPath) > 0 {
						tokenBytes, tokenIsSameLine, err := app.nameOnToken(currentPath, elementToken)
						if err != nil {
							return nil, false, err
						}

						if len(tokenBytes) <= 0 {
							continue
						}

						if tokenIsSameLine {
							for _, oneLine := range tokenBytes {
								data = append(data, oneLine...)
							}

							continue
						}

						output = append(output, tokenBytes...)
						continue
					}

					elementBlock := elementToken.Block()
					index := elementBlock.Discovered()
					input := elementBlock.Input()
					remaining := elementBlock.Remaining()
					amount := len(input) - len(remaining)
					data = append(data, input[index:amount]...)
					isSameLine = true
					continue
				}
			}

			if len(data) <= 0 {
				continue
			}

			output = append(output, data)
		}
	}

	return output, isSameLine, nil
}

func (app *application) anyNameOnToken(anyElement selectors.Name, token results.Token) ([][]byte, error) {
	block := token.Block()
	input := block.Input()
	index := block.Discovered()
	prefixes, err := app.nameInsOnToken(anyElement, token)
	if err != nil {
		return nil, err
	}

	data := input[index:]
	list := [][]byte{}
	for _, onePrefix := range prefixes {
		element := data
		for idx := range data {
			element = data[idx:]
			if bytes.HasPrefix(element, onePrefix) {
				break
			}
		}

		value := []byte{}
		index := len(data) - len(element) + 1
		if len(data) >= index {
			value = data[index:]
		}

		list = append(list, value)
		data = value
	}

	return list, nil
}

func (app *application) getIndexes(amount uint8, min uint8, pMax *uint8) (int, int) {
	max := amount - 1
	if pMax == nil {
		return int(min), int(max)
	}

	if max > *pMax {
		max = *pMax
	}

	if max < min {
		max = min
	}

	return int(min), int(max)
}
