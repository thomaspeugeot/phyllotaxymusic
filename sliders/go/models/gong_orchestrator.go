// generated code - do not edit
package models

// insertion point
// SliderOrchestrator
type SliderOrchestrator struct {
}

func (orchestrator *SliderOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedSlider, backRepoSlider *Slider) {

	stagedSlider.OnAfterUpdate(gongsvgStage, stagedSlider, backRepoSlider)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Slider:
		stage.OnAfterSliderUpdateCallback = new(SliderOrchestrator)

	}

}
