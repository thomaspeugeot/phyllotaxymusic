package models

import (
	buttons "github.com/fullstack-lang/gong/lib/button/go/models"
)

func (parameter *Parameter) UpdateAndCommitButtonStage() {
	parameter.buttonStage.Reset()

	layout := new(buttons.Layout).Stage(parameter.buttonStage)

	group1 := new(buttons.Group).Stage(parameter.buttonStage)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	button := buttons.NewButton(parameter,
		"Export to Musescore",
		"music_note",
		"Export to Musescore",
	)

	group1.Buttons = append(group1.Buttons, button)

	parameter.buttonStage.Commit()
}
