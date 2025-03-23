package models

import (
	button "github.com/fullstack-lang/gong/lib/button/go/models"
)

func (parameter *Parameter) UpdateAndCommitButtonStage() {
	parameter.buttonStage.Reset()

	layout := new(button.Layout).Stage(parameter.buttonStage)

	group1 := new(button.Group).Stage(parameter.buttonStage)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	button := button.NewButton(
		// parameter is the target of the button. parameter implements interface method OnAfterUpdateButton()
		parameter,
		"Export to Musescore",
		"music_note",
		"Export to Musescore",
	)

	group1.Buttons = append(group1.Buttons, button)

	parameter.buttonStage.Commit()
}
