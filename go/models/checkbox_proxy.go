package models

import (
	"log"

	m "github.com/thomaspeugeot/phyllotaxymusic/sliders/go/models"
)

// Generic checkbox creation function
func NewCheckbox(
	parameter *Parameter,
	name string,
	labelForTrue string,
	labelForFalse string,
	valueRef *bool,
) *m.Checkbox {
	checkbox := new(m.Checkbox).Stage(parameter.slidersStage)
	checkbox.Name = name
	checkbox.LabelForTrue = labelForTrue
	checkbox.LabelForFalse = labelForFalse
	checkbox.ValueBool = *valueRef

	proxy := NewCheckboxProxy(
		checkbox,
		valueRef,
		parameter,
	)

	checkbox.Proxy = proxy

	return checkbox
}

// CheckboxProxy is a generic proxy for both int and float64
type CheckboxProxy struct {
	checkbox  *m.Checkbox
	Value     *bool
	parameter *Parameter
}

// Updated handles updating values when the checkbox changes
func (proxy *CheckboxProxy) Updated() {
	*proxy.Value = proxy.checkbox.ValueBool

	log.Println(proxy.parameter.SideLength)

	proxy.parameter.UpdatePhyllotaxyStage()
	proxy.parameter.UpdateAndCommitCursorStage()
	proxy.parameter.UpdateAndCommitSVGStage()
	proxy.parameter.UpdateAndCommitToneStage()
	proxy.parameter.treeProxy.UpdateAndCommitTreeStage()
	proxy.parameter.CommitPhyllotaxymusicStage()
}

// NewCheckboxProxy creates a new proxy for a checkbox
func NewCheckboxProxy(
	checkbox *m.Checkbox,
	value *bool,
	parameter *Parameter,
) *CheckboxProxy {
	proxy := new(CheckboxProxy)
	proxy.checkbox = checkbox
	proxy.Value = value
	proxy.parameter = parameter
	return proxy
}
