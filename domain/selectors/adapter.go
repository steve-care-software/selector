package selectors

import (
	"errors"
	"fmt"
)

type adapter struct {
	builder             Builder
	elementBuilder      ElementBuilder
	anyElementBuilder   AnyElementBuilder
	nameBuilder         NameBuilder
	anyByte             byte
	tokenNameByte       byte
	insideByte          byte
	selectByte          byte
	tokenNameCharacters []byte
	channelCharacters   []byte
}

func createAdapter(
	builder Builder,
	elementBuilder ElementBuilder,
	anyElementBuilder AnyElementBuilder,
	nameBuilder NameBuilder,
	anyByte byte,
	tokenNameByte byte,
	insideByte byte,
	selectByte byte,
	tokenNameCharacters []byte,
	channelCharacters []byte,
) Adapter {
	out := adapter{
		builder:             builder,
		elementBuilder:      elementBuilder,
		anyElementBuilder:   anyElementBuilder,
		nameBuilder:         nameBuilder,
		anyByte:             anyByte,
		tokenNameByte:       tokenNameByte,
		insideByte:          insideByte,
		selectByte:          selectByte,
		tokenNameCharacters: tokenNameCharacters,
		channelCharacters:   channelCharacters,
	}

	return &out
}

// ToScript converts a selector to script
func (app *adapter) ToScript(selector Selector) []byte {
	return nil
}

// ToSelector converts a script to selector
func (app *adapter) ToSelector(script []byte) (Selector, error) {
	// convert to bytes:
	bytes := []byte(script)

	// remove channel characters:
	remainingAfterChans := app.removeChannelCharacters(bytes)

	// retrieve the selector:
	return app.retrieveSelector(remainingAfterChans)
}

func (app *adapter) retrieveSelector(data []byte) (Selector, error) {
	list := []Element{}
	lastRemaining := data
	for {
		retElement, retRemaining, err := app.retrieveElement(lastRemaining)
		if err != nil {
			break
		}

		lastRemaining = retRemaining
		list = append(list, retElement)
	}

	return app.builder.Create().WithList(list).Now()
}

func (app *adapter) retrieveElement(data []byte) (Element, []byte, error) {
	anyElement, remainingAfterAny, err := app.retrieveAnyElement(data)
	if err == nil {
		return anyElement, remainingAfterAny, nil
	}

	name, remainingAfterName, err := app.retrieveElementName(data)
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.elementBuilder.Create().WithName(name).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remainingAfterName, nil
}

func (app *adapter) retrieveElementName(data []byte) (Name, []byte, error) {
	isSelected, remainingAfterIsSelected := app.elementIsSelected(data)
	insideNames, retAfterInsideNames := app.retrieveElementInsideNames(remainingAfterIsSelected)
	tokenName, retAfterTokenName, err := app.retrieveTokenName(retAfterInsideNames, app.tokenNameByte)
	if err != nil {
		return nil, nil, err
	}

	nameBuilder := app.nameBuilder.Create().WithName(tokenName)
	if insideNames != nil {
		nameBuilder.WithInsideNames(insideNames)
	}

	if isSelected {
		nameBuilder.IsSelected()
	}

	ins, err := nameBuilder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, retAfterTokenName, nil
}

func (app *adapter) retrieveElementInsideNames(data []byte) ([]string, []byte) {
	if len(data) <= 0 {
		return nil, data
	}

	remaining := data
	names := []string{}
	for {
		if remaining[0] != app.insideByte {
			break
		}

		tokenName, retAfterTokenName, err := app.retrieveTokenName(remaining, app.insideByte)
		if err != nil {
			break
		}

		names = append(names, tokenName)
		remaining = retAfterTokenName
	}

	return names, remaining
}

func (app *adapter) elementIsSelected(data []byte) (bool, []byte) {
	if len(data) <= 0 {
		return false, data
	}

	if data[0] == app.selectByte {
		return true, data[1:]
	}

	return false, data
}

func (app *adapter) retrieveAnyElement(data []byte) (Element, []byte, error) {
	prevVal := byte(0)
	isAny := false
	isSelected := false
	prefixData := []byte{}
	suffixData := []byte{}
	for idx, value := range data {
		if value == app.anyByte {
			isAny = true
			beforeIdx := idx
			increment := 1
			if prevVal == app.selectByte {
				isSelected = true
				beforeIdx = idx - 1
				increment = 2
			}

			prefixData = data[0:beforeIdx]
			suffixData = data[beforeIdx+increment:]
			break
		}

		prevVal = value
	}

	if !isAny {
		return nil, nil, errors.New("the given data does not represent an AnyElement instance")
	}

	anyElementBuilder := app.anyElementBuilder.Create()
	if isSelected {
		anyElementBuilder.IsSelected()
	}

	if len(prefixData) > 0 {
		prefix, _, err := app.retrieveElementName(prefixData)
		if err != nil {
			return nil, nil, err
		}

		anyElementBuilder.WithPrefix(prefix)
	}

	remainingAfterSuffix := suffixData
	if len(suffixData) > 0 {
		suffix, remaining, err := app.retrieveElementName(suffixData)
		if err != nil {
			return nil, nil, err
		}

		remainingAfterSuffix = remaining
		anyElementBuilder.WithSuffix(suffix)
	}

	any, err := anyElementBuilder.Now()
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.elementBuilder.Create().WithAny(any).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remainingAfterSuffix, nil
}

func (app *adapter) retrieveTokenName(data []byte, prefixByte byte) (string, []byte, error) {
	if len(data) < 1 {
		str := fmt.Sprintf("the tokenName was NOT expecting empty data")
		return "", nil, errors.New(str)
	}

	if data[0] == prefixByte {
		return app.fetchTokenName(data[1:])
	}

	str := fmt.Sprintf("the tokenName was expecting a prefix byte (%d), none provided", prefixByte)
	return "", nil, errors.New(str)
}

func (app *adapter) removeChannelCharacters(input []byte) []byte {
	output := []byte{}
	for _, oneInputByte := range input {
		if app.isBytePresent(oneInputByte, app.channelCharacters) {
			continue
		}

		output = append(output, oneInputByte)
	}

	return output
}

func (app *adapter) isBytePresent(value byte, data []byte) bool {
	isPresent := false
	for _, oneChanByte := range data {
		if value == oneChanByte {
			isPresent = true
			break
		}
	}

	return isPresent
}

func (app *adapter) fetchTokenName(input []byte) (string, []byte, error) {
	nameBytes := []byte{}
	for _, oneInputByte := range input {
		if !app.isBytePresent(oneInputByte, app.tokenNameCharacters) {
			break
		}

		nameBytes = append(nameBytes, oneInputByte)
	}

	if len(nameBytes) <= 0 {
		return "", nil, errors.New("the tokenName must contain at least 1 character, none provided")
	}

	return string(nameBytes), input[len(nameBytes):], nil
}
