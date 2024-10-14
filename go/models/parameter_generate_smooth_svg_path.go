package models

import (
	"fmt"
	"math"
	"strings"
)

// GenerateSmoothSVGPath generates an SVG path data string with smooth Bezier curves forming a closed loop.
// The coordinates are shifted by xc and yc, then rotated around (cx, cy) by angle alpha (in degrees).
func GenerateSmoothSVGPath(x []float64, y []float64, xc, yc, cx, cy, alpha float64) string {
	n := len(x)
	if n < 2 || n != len(y) {
		return ""
	}

	// Precompute sine and cosine of the rotation angle
	alphaRad := alpha * math.Pi / 180.0
	cosAlpha := math.Cos(alphaRad)
	sinAlpha := math.Sin(alphaRad)

	// Shift and rotate the coordinates
	xTransformed := make([]float64, n)
	yTransformed := make([]float64, n)
	for i := 0; i < n; i++ {
		// Shift coordinates by xc and yc
		xShifted := x[i] + xc
		yShifted := y[i] + yc

		// Rotate around (cx, cy)
		xRotated := cosAlpha*(xShifted-cx) - sinAlpha*(yShifted-cy) + cx
		yRotated := sinAlpha*(xShifted-cx) + cosAlpha*(yShifted-cy) + cy

		xTransformed[i] = xRotated
		yTransformed[i] = yRotated
	}

	var path strings.Builder

	// Start the path at the first point
	path.WriteString(fmt.Sprintf("M%.2f %.2f ", xTransformed[0], yTransformed[0]))

	// Calculate control points and append Bezier curves
	for i := 0; i < n; i++ {
		// Indices for points
		p0 := (i - 1 + n) % n
		p1 := i
		p2 := (i + 1) % n
		p3 := (i + 2) % n

		// Calculate original control points using Catmull-Rom to Bezier conversion
		cp1x := x[p1] + (x[p2]-x[p0])/6.0
		cp1y := y[p1] + (y[p2]-y[p0])/6.0
		cp2x := x[p2] - (x[p3]-x[p1])/6.0
		cp2y := y[p2] - (y[p3]-y[p1])/6.0

		// Shift and rotate control points
		// First control point
		cp1xShifted := cp1x + xc
		cp1yShifted := cp1y + yc
		cp1xRotated := cosAlpha*(cp1xShifted-cx) - sinAlpha*(cp1yShifted-cy) + cx
		cp1yRotated := sinAlpha*(cp1xShifted-cx) + cosAlpha*(cp1yShifted-cy) + cy

		// Second control point
		cp2xShifted := cp2x + xc
		cp2yShifted := cp2y + yc
		cp2xRotated := cosAlpha*(cp2xShifted-cx) - sinAlpha*(cp2yShifted-cy) + cx
		cp2yRotated := sinAlpha*(cp2xShifted-cx) + cosAlpha*(cp2yShifted-cy) + cy

		// Append cubic Bezier curve command
		path.WriteString(fmt.Sprintf("C%.2f %.2f %.2f %.2f %.2f %.2f ",
			cp1xRotated, cp1yRotated, cp2xRotated, cp2yRotated, xTransformed[p2], yTransformed[p2]))
	}

	// Close the path
	path.WriteString("Z")

	return strings.TrimSpace(path.String())
}
