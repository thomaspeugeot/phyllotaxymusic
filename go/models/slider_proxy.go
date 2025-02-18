package models

import (
	"log"

	m "github.com/thomaspeugeot/phyllotaxymusic/sliders/go/models"
)

func (parameter *Parameter) newFloat64Slider(
	name string, // The label for the slider
	min float64,
	max float64,
	step float64,
	valueRef *float64,
) *m.Slider {

	slider := new(m.Slider).Stage(parameter.slidersStage)
	slider.Name = name
	slider.IsFloat64 = true
	slider.MinFloat64 = min
	slider.MaxFloat64 = max
	slider.StepFloat64 = step
	slider.ValueFloat64 = *valueRef

	slider.Proxy = NewSliderFloat64Proxy(
		slider,
		valueRef,
		parameter,
	)

	return slider
}

type SliderFloat64Proxy struct {
	slider    *m.Slider
	Value     *float64
	parameter *Parameter
}

func NewSliderFloat64Proxy(
	slider *m.Slider,
	value *float64,
	parameter *Parameter,
) (proxy *SliderFloat64Proxy) {
	proxy = new(SliderFloat64Proxy)
	proxy.slider = slider
	proxy.Value = value
	proxy.parameter = parameter
	return
}

func (proxy *SliderFloat64Proxy) Updated() {
	*proxy.Value = proxy.slider.ValueFloat64

	log.Println(proxy.parameter.SideLength)

	proxy.parameter.UpdatePhyllotaxyStage()
	proxy.parameter.UpdateAndCommitCursorStage()
	proxy.parameter.UpdateAndCommitSVGStage()
	proxy.parameter.UpdateAndCommitToneStage()
	proxy.parameter.treeProxy.UpdateAndCommitTreeStage()
	proxy.parameter.CommitPhyllotaxymusicStage()
}

func (parameter *Parameter) newIntSlider(
	name string, // The label for the slider
	min int,
	max int,
	step int,
	valueRef *int,
) *m.Slider {

	slider := new(m.Slider).Stage(parameter.slidersStage)
	slider.Name = name
	slider.IsInt = true
	slider.MinInt = min
	slider.MaxInt = max
	slider.StepInt = step
	slider.ValueInt = *valueRef

	slider.Proxy = NewSliderIntProxy(
		slider,
		valueRef,
		parameter,
	)

	return slider
}

type SliderIntProxy struct {
	slider    *m.Slider
	Value     *int
	parameter *Parameter
}

func NewSliderIntProxy(
	slider *m.Slider,
	value *int,
	parameter *Parameter,
) (proxy *SliderIntProxy) {
	proxy = new(SliderIntProxy)
	proxy.slider = slider
	proxy.Value = value
	proxy.parameter = parameter
	return
}

func (proxy *SliderIntProxy) Updated() {
	*proxy.Value = proxy.slider.ValueInt

	log.Println(proxy.parameter.SideLength)

	proxy.parameter.UpdatePhyllotaxyStage()
	proxy.parameter.UpdateAndCommitCursorStage()
	proxy.parameter.UpdateAndCommitSVGStage()
	proxy.parameter.UpdateAndCommitToneStage()
	proxy.parameter.treeProxy.UpdateAndCommitTreeStage()
	proxy.parameter.CommitPhyllotaxymusicStage()
}
