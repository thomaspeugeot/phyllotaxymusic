package models

import "log"

type Cursor struct {
	Name string

	// where it starts and end
	StartX, EndX float64
	Y1, Y2       float64

	DurationSeconds float64

	Presentation

	IsPlaying bool

	// to receive notification that the music started
	notifyCh chan bool
}

func (cursor *Cursor) SetNotifyChannel(notifyCh chan bool) {
	cursor.notifyCh = notifyCh
}

func (cursor *Cursor) WaitForPlayNotifications(stage *StageStruct) {
	go func() {

		for newOrder := range cursor.notifyCh {
			cursor.IsPlaying = newOrder
			log.Println("cursor received new order via channel", cursor.IsPlaying)
			stage.Commit()
		}
	}()
}
