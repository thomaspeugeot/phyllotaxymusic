package models

func (stager *Stager) OnAfterUpdateSliderElement() {

	stager.UpdateAllStagesButSliders()
}
