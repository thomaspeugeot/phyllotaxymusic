// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Axis
	// insertion point per field

	// Compute reverse map for named struct AxisGrid
	// insertion point per field
	clear(stage.AxisGrid_Axiss_reverseMap)
	stage.AxisGrid_Axiss_reverseMap = make(map[*Axis]*AxisGrid)
	for axisgrid := range stage.AxisGrids {
		_ = axisgrid
		for _, _axis := range axisgrid.Axiss {
			stage.AxisGrid_Axiss_reverseMap[_axis] = axisgrid
		}
	}

	// Compute reverse map for named struct Bezier
	// insertion point per field

	// Compute reverse map for named struct BezierGrid
	// insertion point per field
	clear(stage.BezierGrid_Beziers_reverseMap)
	stage.BezierGrid_Beziers_reverseMap = make(map[*Bezier]*BezierGrid)
	for beziergrid := range stage.BezierGrids {
		_ = beziergrid
		for _, _bezier := range beziergrid.Beziers {
			stage.BezierGrid_Beziers_reverseMap[_bezier] = beziergrid
		}
	}

	// Compute reverse map for named struct BezierGridStack
	// insertion point per field
	clear(stage.BezierGridStack_BezierGrids_reverseMap)
	stage.BezierGridStack_BezierGrids_reverseMap = make(map[*BezierGrid]*BezierGridStack)
	for beziergridstack := range stage.BezierGridStacks {
		_ = beziergridstack
		for _, _beziergrid := range beziergridstack.BezierGrids {
			stage.BezierGridStack_BezierGrids_reverseMap[_beziergrid] = beziergridstack
		}
	}

	// Compute reverse map for named struct Chapter
	// insertion point per field

	// Compute reverse map for named struct Circle
	// insertion point per field

	// Compute reverse map for named struct CircleGrid
	// insertion point per field
	clear(stage.CircleGrid_Circles_reverseMap)
	stage.CircleGrid_Circles_reverseMap = make(map[*Circle]*CircleGrid)
	for circlegrid := range stage.CircleGrids {
		_ = circlegrid
		for _, _circle := range circlegrid.Circles {
			stage.CircleGrid_Circles_reverseMap[_circle] = circlegrid
		}
	}

	// Compute reverse map for named struct Content
	// insertion point per field
	clear(stage.Content_Chapters_reverseMap)
	stage.Content_Chapters_reverseMap = make(map[*Chapter]*Content)
	for content := range stage.Contents {
		_ = content
		for _, _chapter := range content.Chapters {
			stage.Content_Chapters_reverseMap[_chapter] = content
		}
	}

	// Compute reverse map for named struct ExportToMusicxml
	// insertion point per field

	// Compute reverse map for named struct FrontCurve
	// insertion point per field

	// Compute reverse map for named struct FrontCurveStack
	// insertion point per field
	clear(stage.FrontCurveStack_FrontCurves_reverseMap)
	stage.FrontCurveStack_FrontCurves_reverseMap = make(map[*FrontCurve]*FrontCurveStack)
	for frontcurvestack := range stage.FrontCurveStacks {
		_ = frontcurvestack
		for _, _frontcurve := range frontcurvestack.FrontCurves {
			stage.FrontCurveStack_FrontCurves_reverseMap[_frontcurve] = frontcurvestack
		}
	}
	clear(stage.FrontCurveStack_SpiralCircles_reverseMap)
	stage.FrontCurveStack_SpiralCircles_reverseMap = make(map[*SpiralCircle]*FrontCurveStack)
	for frontcurvestack := range stage.FrontCurveStacks {
		_ = frontcurvestack
		for _, _spiralcircle := range frontcurvestack.SpiralCircles {
			stage.FrontCurveStack_SpiralCircles_reverseMap[_spiralcircle] = frontcurvestack
		}
	}

	// Compute reverse map for named struct HorizontalAxis
	// insertion point per field

	// Compute reverse map for named struct Key
	// insertion point per field

	// Compute reverse map for named struct Parameter
	// insertion point per field

	// Compute reverse map for named struct Rhombus
	// insertion point per field

	// Compute reverse map for named struct RhombusGrid
	// insertion point per field
	clear(stage.RhombusGrid_Rhombuses_reverseMap)
	stage.RhombusGrid_Rhombuses_reverseMap = make(map[*Rhombus]*RhombusGrid)
	for rhombusgrid := range stage.RhombusGrids {
		_ = rhombusgrid
		for _, _rhombus := range rhombusgrid.Rhombuses {
			stage.RhombusGrid_Rhombuses_reverseMap[_rhombus] = rhombusgrid
		}
	}

	// Compute reverse map for named struct ShapeCategory
	// insertion point per field

	// Compute reverse map for named struct SpiralBezier
	// insertion point per field

	// Compute reverse map for named struct SpiralBezierGrid
	// insertion point per field
	clear(stage.SpiralBezierGrid_SpiralBeziers_reverseMap)
	stage.SpiralBezierGrid_SpiralBeziers_reverseMap = make(map[*SpiralBezier]*SpiralBezierGrid)
	for spiralbeziergrid := range stage.SpiralBezierGrids {
		_ = spiralbeziergrid
		for _, _spiralbezier := range spiralbeziergrid.SpiralBeziers {
			stage.SpiralBezierGrid_SpiralBeziers_reverseMap[_spiralbezier] = spiralbeziergrid
		}
	}

	// Compute reverse map for named struct SpiralCircle
	// insertion point per field

	// Compute reverse map for named struct SpiralCircleGrid
	// insertion point per field
	clear(stage.SpiralCircleGrid_SpiralCircles_reverseMap)
	stage.SpiralCircleGrid_SpiralCircles_reverseMap = make(map[*SpiralCircle]*SpiralCircleGrid)
	for spiralcirclegrid := range stage.SpiralCircleGrids {
		_ = spiralcirclegrid
		for _, _spiralcircle := range spiralcirclegrid.SpiralCircles {
			stage.SpiralCircleGrid_SpiralCircles_reverseMap[_spiralcircle] = spiralcirclegrid
		}
	}

	// Compute reverse map for named struct SpiralLine
	// insertion point per field

	// Compute reverse map for named struct SpiralLineGrid
	// insertion point per field
	clear(stage.SpiralLineGrid_SpiralLines_reverseMap)
	stage.SpiralLineGrid_SpiralLines_reverseMap = make(map[*SpiralLine]*SpiralLineGrid)
	for spirallinegrid := range stage.SpiralLineGrids {
		_ = spirallinegrid
		for _, _spiralline := range spirallinegrid.SpiralLines {
			stage.SpiralLineGrid_SpiralLines_reverseMap[_spiralline] = spirallinegrid
		}
	}

	// Compute reverse map for named struct SpiralOrigin
	// insertion point per field

	// Compute reverse map for named struct SpiralRhombus
	// insertion point per field

	// Compute reverse map for named struct SpiralRhombusGrid
	// insertion point per field
	clear(stage.SpiralRhombusGrid_SpiralRhombuses_reverseMap)
	stage.SpiralRhombusGrid_SpiralRhombuses_reverseMap = make(map[*SpiralRhombus]*SpiralRhombusGrid)
	for spiralrhombusgrid := range stage.SpiralRhombusGrids {
		_ = spiralrhombusgrid
		for _, _spiralrhombus := range spiralrhombusgrid.SpiralRhombuses {
			stage.SpiralRhombusGrid_SpiralRhombuses_reverseMap[_spiralrhombus] = spiralrhombusgrid
		}
	}

	// Compute reverse map for named struct VerticalAxis
	// insertion point per field

}
