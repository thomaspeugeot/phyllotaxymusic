package models

import (
	"fmt"
	"math"
	"strings"
)

// GenerateSmoothSVGPath generates an SVG path data string with smooth Bezier curves forming a closed loop.
// The coordinates are shifted by xc and yc, rotated around (cx, cy) by angle alpha (in degrees),
// and then offset outward from (cx, cy) by the specified offset distance.
func GenerateSmoothSVGPath(x []float64, y []float64, xc, yc, cx, cy, alpha, offset float64) string {
	n := len(x)
	if n < 2 || n != len(y) {
		return ""
	}

	// Precompute sine and cosine of the rotation angle
	alphaRad := alpha * math.Pi / 180.0
	cosAlpha := math.Cos(alphaRad)
	sinAlpha := math.Sin(alphaRad)

	// Shift, rotate, and offset the coordinates
	xTransformed := make([]float64, n)
	yTransformed := make([]float64, n)
	for i := 0; i < n; i++ {
		// Shift coordinates by xc and yc
		xShifted := x[i] + xc
		yShifted := y[i] + yc

		// Rotate around (cx, cy)
		xRotated := cosAlpha*(xShifted-cx) - sinAlpha*(yShifted-cy) + cx
		yRotated := sinAlpha*(xShifted-cx) + cosAlpha*(yShifted-cy) + cy

		// Compute the vector from (cx, cy) to the rotated point
		dx := xRotated - cx
		dy := yRotated - cy
		length := math.Hypot(dx, dy)

		// Normalize the vector and apply the offset
		if length != 0 {
			dx /= length
			dy /= length
		}
		xOffset := xRotated + offset*dx
		yOffset := yRotated + offset*dy

		xTransformed[i] = xOffset
		yTransformed[i] = yOffset
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

		// Offset first control point
		dx1 := cp1xRotated - cx
		dy1 := cp1yRotated - cy
		length1 := math.Hypot(dx1, dy1)
		if length1 != 0 {
			dx1 /= length1
			dy1 /= length1
		}
		cp1xOffset := cp1xRotated + offset*dx1
		cp1yOffset := cp1yRotated + offset*dy1

		// Second control point
		cp2xShifted := cp2x + xc
		cp2yShifted := cp2y + yc
		cp2xRotated := cosAlpha*(cp2xShifted-cx) - sinAlpha*(cp2yShifted-cy) + cx
		cp2yRotated := sinAlpha*(cp2xShifted-cx) + cosAlpha*(cp2yShifted-cy) + cy

		// Offset second control point
		dx2 := cp2xRotated - cx
		dy2 := cp2yRotated - cy
		length2 := math.Hypot(dx2, dy2)
		if length2 != 0 {
			dx2 /= length2
			dy2 /= length2
		}
		cp2xOffset := cp2xRotated + offset*dx2
		cp2yOffset := cp2yRotated + offset*dy2

		// Append cubic Bezier curve command
		path.WriteString(fmt.Sprintf("C%.2f %.2f %.2f %.2f %.2f %.2f ",
			cp1xOffset, cp1yOffset, cp2xOffset, cp2yOffset, xTransformed[p2], yTransformed[p2]))
	}

	// Close the path
	path.WriteString("Z")

	return strings.TrimSpace(path.String())
}
