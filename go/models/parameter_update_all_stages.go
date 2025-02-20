package models

import (
	m "github.com/thomaspeugeot/phyllotaxymusic/sliders/go/models"
)

func (parameter *Parameter) OnAfterUpdateSliderElement() {

	parameter.UpdateAllStagesButSliders()
}

func (parameter *Parameter) GetSliderStage() *m.StageStruct {
	return parameter.slidersStage	
}
