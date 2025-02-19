package models

import (
	m "github.com/thomaspeugeot/phyllotaxymusic/sliders/go/models"
)

func (parameter *Parameter) OnAfterUpdateSliderElement() {

	parameter.UpdatePhyllotaxyStage()
	parameter.UpdateAndCommitCursorStage()
	parameter.UpdateAndCommitSVGStage()
	parameter.UpdateAndCommitToneStage()
	parameter.treeProxy.UpdateAndCommitTreeStage()
	parameter.CommitPhyllotaxymusicStage()
}

func (parameter *Parameter) GetSliderStage() *m.StageStruct {
	return parameter.slidersStage
}
