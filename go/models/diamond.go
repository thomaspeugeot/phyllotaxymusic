package models

// Rhombus is defined by its center, a side length, a direction (with an angle)
// and an opening angle along this direction
type Rhombus struct {
	Name string

	CenterX, CenterY float64

	SideLength float64

	// Angle of one arbitrary Axis of the Rhombus to the horizontal plan
	// in degress.
	Angle float64

	// InsideAngle is the angle inside the rhombus in the edge that
	// crosses the axis of the rhombus
	InsideAngle float64
}
