package models

func (checkbox *Checkbox) OnAfterUpdate(
	stage *StageStruct,
	stageCheckbox, frontCheckbox *Checkbox) {

	checkbox.ValueBool = frontCheckbox.ValueBool

	if checkbox.Proxy != nil {
		checkbox.Proxy.Updated()
	}
}
