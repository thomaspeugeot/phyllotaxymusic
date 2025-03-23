package models

func (stager *Stager) OnAfterUpdateNode() {
	stager.UpdateAllStages()
}

func (stager *Stager) UpdateAllStages() {
	stager.UpdatePhyllotaxyStage()
	stager.UpdateAndCommitCursorStage()
	stager.UpdateAndCommitSVGStage()
	stager.UpdateAndCommitToneStage()
	stager.UpdateAndCommitTreeStage()
	stager.UpdateAndCommitSlidersStage()
	stager.UpdateAndCommitButtonStage()
	stager.phyllotaxymusicStage.Commit()
}

func (stager *Stager) UpdateAllStagesButSliders() {
	stager.UpdatePhyllotaxyStage()
	stager.UpdateAndCommitCursorStage()
	stager.UpdateAndCommitSVGStage()
	stager.UpdateAndCommitToneStage()
	stager.UpdateAndCommitTreeStage()
	stager.UpdateAndCommitButtonStage()
	stager.phyllotaxymusicStage.Commit()
}
