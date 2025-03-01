package models

import (
	m "github.com/fullstack-lang/gong/lib/slider/go/models"
)

func (parameter *Parameter) OnAfterUpdateSliderElement() {

	parameter.UpdateAllStagesButSliders()
}

func (parameter *Parameter) GetSliderStage() *m.StageStruct {
	return parameter.slidersStage
}
