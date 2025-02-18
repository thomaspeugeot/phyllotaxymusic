package models

import (
	"log"

	m "github.com/thomaspeugeot/phyllotaxymusic/sliders/go/models"
)

func (parameter *Parameter) UpdateAndCommitSlidersStage() {
	parameter.slidersStage.Reset()

	layout := new(m.Layout).Stage(parameter.slidersStage)

	groupSideLength := new(m.Group).Stage(parameter.slidersStage)
	groupSideLength.Size = 60
	layout.Groups = append(layout.Groups, groupSideLength)

	groupSideLength.Sliders = append(
		groupSideLength.Sliders,
		parameter.newFloat64Slider(
			"Side Length",
			5,   // Min
			200, // Max
			5,   // Step
			&parameter.SideLength,
		),
	)

	groupSideLength.Sliders = append(
		groupSideLength.Sliders,
		parameter.newFloat64Slider(
			"Origin X",
			-100, // Min
			600,  // Max
			10,   // Step
			&parameter.OriginX,
		),
	)

	groupInsideAngle := new(m.Group).Stage(parameter.slidersStage)
	groupInsideAngle.Size = 60
	layout.Groups = append(layout.Groups, groupInsideAngle)

	parameter.slidersStage.Commit()
}

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

	slider.Proxy = NewSliderProxy(
		slider,
		valueRef,
		parameter,
	)

	return slider
}

type SliderProxy struct {
	slider    *m.Slider
	Value     *float64
	parameter *Parameter
}

func NewSliderProxy(
	slider *m.Slider,
	value *float64,
	parameter *Parameter,
) (proxy *SliderProxy) {
	proxy = new(SliderProxy)
	proxy.slider = slider
	proxy.Value = value
	proxy.parameter = parameter
	return
}

func (proxy *SliderProxy) Updated() {
	*proxy.Value = proxy.slider.ValueFloat64

	log.Println(proxy.parameter.SideLength)

	proxy.parameter.UpdatePhyllotaxyStage()
	proxy.parameter.UpdateAndCommitCursorStage()
	proxy.parameter.UpdateAndCommitSVGStage()
	proxy.parameter.UpdateAndCommitToneStage()
	proxy.parameter.treeProxy.UpdateAndCommitTreeStage()
	proxy.parameter.CommitPhyllotaxymusicStage()
}
