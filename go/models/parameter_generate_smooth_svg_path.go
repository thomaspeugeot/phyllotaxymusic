package models

import (
	"fmt"
	"strings"
)

// GenerateSmoothSVGPath generates an SVG path data string with smooth Bezier curves forming a closed loop.
// The coordinates are shifted by xc and yc.
func GenerateSmoothSVGPath(x []float64, y []float64, xc, yc float64) string {
	n := len(x)
	if n < 2 || n != len(y) {
		return ""
	}

	// Shift the coordinates by xc and yc
	xShifted := make([]float64, n)
	yShifted := make([]float64, n)
	for i := 0; i < n; i++ {
		xShifted[i] = x[i] + xc
		yShifted[i] = y[i] + yc
	}

	var path strings.Builder

	// Start the path at the first point
	path.WriteString(fmt.Sprintf("M%.2f %.2f ", xShifted[0], yShifted[0]))

	// Calculate control points and append Bezier curves
	for i := 0; i < n; i++ {
		// Indices for points
		p0 := (i - 1 + n) % n
		p1 := i
		p2 := (i + 1) % n
		p3 := (i + 2) % n

		// Calculate control points using Catmull-Rom to Bezier conversion
		cp1x := xShifted[p1] + (xShifted[p2]-xShifted[p0])/6.0
		cp1y := yShifted[p1] + (yShifted[p2]-yShifted[p0])/6.0
		cp2x := xShifted[p2] - (xShifted[p3]-xShifted[p1])/6.0
		cp2y := yShifted[p2] - (yShifted[p3]-yShifted[p1])/6.0

		// Append cubic Bezier curve command
		path.WriteString(fmt.Sprintf("C%.2f %.2f %.2f %.2f %.2f %.2f ",
			cp1x, cp1y, cp2x, cp2y, xShifted[p2], yShifted[p2]))
	}

	// Close the path
	path.WriteString("Z")

	return strings.TrimSpace(path.String())
}

func main() {
	x := []float64{100.0, 150.0, 200.0, 250.0, 200.0, 150.0}
	y := []float64{100.0, 50.0, 100.0, 150.0, 200.0, 150.0}
	xc := 10.0 // Shift in the x-direction
	yc := 20.0 // Shift in the y-direction

	svgPath := GenerateSmoothSVGPath(x, y, xc, yc)
	fmt.Println(svgPath)
}
