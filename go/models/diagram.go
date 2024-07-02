package models

type Diagram struct {
	Name string
	N    int
	M    int

	// DiamondAngle is the angle in degree of the diamond at the origin 0,0
	DiamondAngle float64

	OriginX     float64
	OriginY     float64
	DiamondSide float64
}
