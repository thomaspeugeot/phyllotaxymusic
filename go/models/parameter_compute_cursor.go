package models

func (p *Parameter) computeCursor() {

	p.Cursor.AngleDegree = 90
	p.Cursor.Length = 2000
	p.Cursor.CenterX = 0
	p.Cursor.CenterY = float64(p.FirstVoiceShiftY) * p.SideLength
}
