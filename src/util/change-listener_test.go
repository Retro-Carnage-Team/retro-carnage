package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeHandlerShouldHandleMatchingProperty(t *testing.T) {
	const propertyName = "name"
	var listener = ChangeListener{Callback: func(value interface{}, property string) {}, PropertyNames: []string{propertyName}}
	var result = listener.handlesProperty(propertyName)

	assert.Equal(t, true, result)
}

func TestChangeHandlerShouldHandleMatchingPropertyFromList(t *testing.T) {
	const propertyName = "name"
	var listener = ChangeListener{Callback: func(value interface{}, property string) {}, PropertyNames: []string{propertyName, "name2"}}
	var result = listener.handlesProperty(propertyName)

	assert.Equal(t, true, result)
}

func TestChangeHandlerShouldNotHandleOtherProperties(t *testing.T) {
	const propertyName = "name"
	var listener = ChangeListener{Callback: func(value interface{}, property string) {}, PropertyNames: []string{propertyName}}
	var result = listener.handlesProperty("not-name")

	assert.Equal(t, false, result)
}

func TestChangeHandlerShouldHandleAllPropertiesWhenNoneIsSpecified(t *testing.T) {
	var listener = ChangeListener{Callback: func(value interface{}, property string) {}, PropertyNames: []string{}}
	var result = listener.handlesProperty("name")

	assert.Equal(t, true, result)
}

func TestChangeHandlerShouldCallSpecifiedCallbackCorrectly(t *testing.T) {
	const propertyNameValue = "property-name"
	const newValue = 42
	var called = false

	var callback = func(value interface{}, propertyName string) {
		called = true
		assert.Equal(t, newValue, value)
		assert.Equal(t, propertyNameValue, propertyName)
	}

	var filteredListener = ChangeListener{Callback: callback, PropertyNames: []string{propertyNameValue}}
	filteredListener.Call(newValue, propertyNameValue)
	assert.Equal(t, true, called)

	called = false

	var unfilteredListener = ChangeListener{Callback: callback, PropertyNames: []string{}}
	unfilteredListener.Call(newValue, propertyNameValue)
	assert.Equal(t, true, called)
}
