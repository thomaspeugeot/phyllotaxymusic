package models

import "fmt"

func (p *Parameter) ComputeNoteInfos(stage *StageStruct) {

	for i, noteInfo := range p.NoteInfos {
		noteInfo.Name = fmt.Sprintf("%d", i)
	}

	if len(p.NoteInfos) == p.NbBeatLinesPerCurve {
		return
	}

	// remove extra note infos
	if len(p.NoteInfos) > p.NbBeatLinesPerCurve {
		for _, ni := range p.NoteInfos[p.NbBeatLinesPerCurve:] {
			ni.Unstage(stage)
		}
		p.NoteInfos = p.NoteInfos[:p.NbBeatLinesPerCurve]
	}

}
