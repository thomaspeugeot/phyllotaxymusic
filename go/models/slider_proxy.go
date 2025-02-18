package models

import (
	"log"

	m "github.com/thomaspeugeot/phyllotaxymusic/sliders/go/models"
)

// Define a Number constraint for our generic types
type Number interface {
	~int | ~float64
}

// Generic slider creation function
func NewSlider[T Number](
	parameter *Parameter,
	name string,
	min T,
	max T,
	step T,
	valueRef *T,
) *m.Slider {
	slider := new(m.Slider).Stage(parameter.slidersStage)
	slider.Name = name

	// Set appropriate values based on type
	switch any(min).(type) {
	case float64:
		slider.IsFloat64 = true
		slider.MinFloat64 = any(min).(float64)
		slider.MaxFloat64 = any(max).(float64)
		slider.StepFloat64 = any(step).(float64)
		slider.ValueFloat64 = any(*valueRef).(float64)
	case int:
		slider.IsInt = true
		slider.MinInt = any(min).(int)
		slider.MaxInt = any(max).(int)
		slider.StepInt = any(step).(int)
		slider.ValueInt = any(*valueRef).(int)
	}

	proxy := NewSliderProxy(
		slider,
		valueRef,
		parameter,
	)

	slider.Proxy = proxy

	return slider
}

// SliderProxy is a generic proxy for both int and float64
type SliderProxy[T Number] struct {
	slider    *m.Slider
	Value     *T
	parameter *Parameter
}

// Updated handles updating values when the slider changes
func (proxy *SliderProxy[T]) Updated() {
	// Update the value based on its type
	switch value := any(proxy.Value).(type) {
	case *float64:
		*value = proxy.slider.ValueFloat64
	case *int:
		*value = proxy.slider.ValueInt
	}

	log.Println(proxy.parameter.SideLength)

	proxy.parameter.UpdatePhyllotaxyStage()
	proxy.parameter.UpdateAndCommitCursorStage()
	proxy.parameter.UpdateAndCommitSVGStage()
	proxy.parameter.UpdateAndCommitToneStage()
	proxy.parameter.treeProxy.UpdateAndCommitTreeStage()
	proxy.parameter.CommitPhyllotaxymusicStage()
}

// NewSliderProxy creates a new proxy for a slider
func NewSliderProxy[T Number](
	slider *m.Slider,
	value *T,
	parameter *Parameter,
) *SliderProxy[T] {
	proxy := new(SliderProxy[T])
	proxy.slider = slider
	proxy.Value = value
	proxy.parameter = parameter
	return proxy
}
