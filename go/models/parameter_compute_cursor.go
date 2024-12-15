package models

func (p *Parameter) computeCursor() {

	p.Cursor.AngleDegree = 90
	p.Cursor.Length = 2000
	p.Cursor.CenterX = 0
	p.Cursor.CenterY = float64(p.FirstVoiceShiftY) * p.SideLength
	p.Cursor.IsMoving = p.IsPlaying
	p.Cursor.StartX = p.Cursor.CenterX
	p.Cursor.EndX = p.SideLength * 2
	p.Cursor.DurationSeconds = 1 / p.Speed
}
