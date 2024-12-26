package models

type Cursor struct {
	Name string

	// in degrees
	AngleDegree float64
	Length      float64

	CenterX, CenterY float64

	StartX, EndX    float64
	DurationSeconds float64
	IsMoving        bool

	Presentation

	IsPlaying bool

	// to receive notification that the music started
	notifyCh chan struct{}
}

func (cursor *Cursor) SetNotifyChannel(notifyCh chan struct{}) {
	cursor.notifyCh = notifyCh
}

func (cursor *Cursor) WaitForPlayNotifications(stage *StageStruct) {
	go func() {

		for range cursor.notifyCh {
			cursor.IsPlaying = true
			stage.Commit()
		}
	}()
}
