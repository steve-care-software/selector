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
func (app *application) Compile(script string) (selectors.Selector, error) {
	return app.adapter.ToSelector(script)
}

// Execute executes a selector on validation result
func (app *application) Execute(selector selectors.Selector, result results.Result) ([]byte, error) {
	if !result.Token().IsSuccess() {
		return nil, errors.New("the selector cannot extract result tokens because the result is invalid")
	}

	list := selector.List()
	token := result.Token()
	return app.elementsOnToken(list, token)
}

func (app *application) elementsOnToken(elements []selectors.Element, token results.Token) ([]byte, error) {
	output := []byte{}
	for _, oneElement := range elements {
		bytes, err := app.elementOnToken(oneElement, token)
		if err != nil {
			return nil, err
		}

		if bytes != nil {
			output = append(output, bytes...)
		}
	}

	return output, nil
}

func (app *application) elementOnToken(element selectors.Element, token results.Token) ([]byte, error) {
	if element.IsName() {
		name := element.Name()
		return app.nameInsOnToken(name, token)
	}

	anyElement := element.Any()
	return app.anyElementOnToken(anyElement, token)
}

func (app *application) nameInsOnToken(nameIns selectors.Name, token results.Token) ([]byte, error) {
	name := nameIns.Name()
	path := []string{}
	if nameIns.HasInsideNames() {
		path = nameIns.InsideNames()
	}

	path = append(path, name)
	isFound, bytes, err := app.nameOnToken(path, token)
	if err != nil {
		return nil, err
	}

	if isFound && nameIns.IsSelected() {
		return bytes, nil
	}

	return nil, nil
}

func (app *application) nameOnToken(path []string, token results.Token) (bool, []byte, error) {
	block := token.Block()
	if !block.IsSuccess() {
		str := fmt.Sprintf("the block's token (name: %s) is NOT successful and therefore its value cannot be extracted", token.Name())
		return false, nil, errors.New(str)
	}

	if len(path) <= 0 {
		return false, nil, nil
	}

	name := path[0]
	isFound := token.Name() == name
	currentPath := path
	if isFound {
		currentPath = path[1:]
	}

	output := []byte{}
	lines := block.List()
	for _, oneLine := range lines {
		if !oneLine.IsSuccess() {
			continue
		}

		elementsWithCardinality := oneLine.Elements()
		for _, oneElementWithCardinality := range elementsWithCardinality {
			if !oneElementWithCardinality.IsSuccess() {
				continue
			}

			if !oneElementWithCardinality.HasMatches() {
				continue
			}

			matches := oneElementWithCardinality.Matches()
			for _, oneElement := range matches {
				if oneElement.IsValue() {
					pValue := oneElement.Value()
					output = append(output, *pValue)
					continue
				}

				if oneElement.IsToken() {
					elementToken := oneElement.Token()
					isTokenFound, tokenBytes, err := app.nameOnToken(currentPath, elementToken)
					if err != nil {
						return false, nil, err
					}

					if tokenBytes == nil {
						continue
					}

					output = append(output, tokenBytes...)
					if isTokenFound {
						isFound = true
						break
					}

					continue
				}
			}

			if isFound && len(output) > 0 {
				break
			}
		}
	}

	if isFound {
		return isFound, output, nil
	}

	return false, nil, nil
}

func (app *application) anyElementOnToken(anyElement selectors.AnyElement, token results.Token) ([]byte, error) {
	block := token.Block()
	input := block.Input()
	index := block.Discovered()
	prefixName := anyElement.Prefix()
	prefix, err := app.nameInsOnToken(prefixName, token)
	if err != nil {
		return nil, err
	}

	stop := false
	data := input[index:]
	output := input[index:]
	for idx := range data {
		if bytes.HasPrefix(output, prefix) {
			stop = true
		}

		output = data[idx+1:]
		if stop {
			break
		}
	}

	return output, nil
}
