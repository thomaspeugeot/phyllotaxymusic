package models

import (
	m "github.com/thomaspeugeot/phyllotaxymusic/buttons/go/models"
)

func (parameter *Parameter) UpdateAndCommitButtonsStage() {
	parameter.buttonsStage.Reset()

	layout := new(m.Layout).Stage(parameter.buttonsStage)

	group1 := new(m.Group).Stage(parameter.buttonsStage)
	group1.Percentage = 25
	layout.Groups = append(layout.Groups, group1)

	button := new(m.Button).Stage(parameter.buttonsStage)
	button.Icon = "music_note"
	button.Label = "Export to Musescore"
	button.Name = "Export to Musescore"

	group1.Buttons = append(group1.Buttons, button)

	parameter.buttonsStage.Commit()
}
