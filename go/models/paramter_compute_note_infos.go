package models

import "fmt"

func (p *Parameter) ComputeNoteInfos(stage *StageStruct) {

	for i, noteInfo := range p.NoteInfos {
		noteInfo.Name = fmt.Sprintf("%d", i)
	}

	if len(p.NoteInfos) == p.NbOfBeatsInTheme {
		return
	}

	// remove extra note infos
	if len(p.NoteInfos) > p.NbOfBeatsInTheme {
		for _, ni := range p.NoteInfos[p.NbOfBeatsInTheme:] {
			ni.Unstage(stage)
		}
		p.NoteInfos = p.NoteInfos[:p.NbOfBeatsInTheme]
	}

}
