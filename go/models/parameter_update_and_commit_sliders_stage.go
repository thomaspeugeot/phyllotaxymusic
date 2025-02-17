package models

import (
	m "github.com/thomaspeugeot/phyllotaxymusic/sliders/go/models"
)

func (p *Parameter) UpdateAndCommitSlidersStage() {
	p.slidersStage.Reset()

	layout := new(m.Layout).Stage(p.slidersStage)

	groupSideLength := new(m.Group).Stage(p.slidersStage)
	layout.Groups = append(layout.Groups, groupSideLength)

	{
		slider := new(m.Slider).Stage(p.slidersStage)
		slider.Name = "Side Length"
		slider.IsInt = true
		slider.MinInt = 5
		slider.MaxInt = 200
		slider.StepInt = 5
		groupSideLength.Sliders = append(groupSideLength.Sliders, slider)
	}

	p.slidersStage.Commit()
}
