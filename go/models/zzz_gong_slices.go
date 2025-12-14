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
	stage.BezierGrid_Beziers_reverseMap = make(map[*Bezier]*BezierGrid)
	for beziergrid := range stage.BezierGrids {
		_ = beziergrid
		for _, _bezier := range beziergrid.Beziers {
			stage.BezierGrid_Beziers_reverseMap[_bezier] = beziergrid
		}
	}

	// Compute reverse map for named struct BezierGridStack
	// insertion point per field
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
	stage.CircleGrid_Circles_reverseMap = make(map[*Circle]*CircleGrid)
	for circlegrid := range stage.CircleGrids {
		_ = circlegrid
		for _, _circle := range circlegrid.Circles {
			stage.CircleGrid_Circles_reverseMap[_circle] = circlegrid
		}
	}

	// Compute reverse map for named struct Content
	// insertion point per field
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
	stage.FrontCurveStack_FrontCurves_reverseMap = make(map[*FrontCurve]*FrontCurveStack)
	for frontcurvestack := range stage.FrontCurveStacks {
		_ = frontcurvestack
		for _, _frontcurve := range frontcurvestack.FrontCurves {
			stage.FrontCurveStack_FrontCurves_reverseMap[_frontcurve] = frontcurvestack
		}
	}
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

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Axiss {
		res = append(res, instance)
	}

	for instance := range stage.AxisGrids {
		res = append(res, instance)
	}

	for instance := range stage.Beziers {
		res = append(res, instance)
	}

	for instance := range stage.BezierGrids {
		res = append(res, instance)
	}

	for instance := range stage.BezierGridStacks {
		res = append(res, instance)
	}

	for instance := range stage.Chapters {
		res = append(res, instance)
	}

	for instance := range stage.Circles {
		res = append(res, instance)
	}

	for instance := range stage.CircleGrids {
		res = append(res, instance)
	}

	for instance := range stage.Contents {
		res = append(res, instance)
	}

	for instance := range stage.ExportToMusicxmls {
		res = append(res, instance)
	}

	for instance := range stage.FrontCurves {
		res = append(res, instance)
	}

	for instance := range stage.FrontCurveStacks {
		res = append(res, instance)
	}

	for instance := range stage.HorizontalAxiss {
		res = append(res, instance)
	}

	for instance := range stage.Keys {
		res = append(res, instance)
	}

	for instance := range stage.Parameters {
		res = append(res, instance)
	}

	for instance := range stage.Rhombuss {
		res = append(res, instance)
	}

	for instance := range stage.RhombusGrids {
		res = append(res, instance)
	}

	for instance := range stage.ShapeCategorys {
		res = append(res, instance)
	}

	for instance := range stage.SpiralBeziers {
		res = append(res, instance)
	}

	for instance := range stage.SpiralBezierGrids {
		res = append(res, instance)
	}

	for instance := range stage.SpiralCircles {
		res = append(res, instance)
	}

	for instance := range stage.SpiralCircleGrids {
		res = append(res, instance)
	}

	for instance := range stage.SpiralLines {
		res = append(res, instance)
	}

	for instance := range stage.SpiralLineGrids {
		res = append(res, instance)
	}

	for instance := range stage.SpiralOrigins {
		res = append(res, instance)
	}

	for instance := range stage.SpiralRhombuss {
		res = append(res, instance)
	}

	for instance := range stage.SpiralRhombusGrids {
		res = append(res, instance)
	}

	for instance := range stage.VerticalAxiss {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (axis *Axis) GongCopy() GongstructIF {
	newInstance := *axis
	return &newInstance
}

func (axisgrid *AxisGrid) GongCopy() GongstructIF {
	newInstance := *axisgrid
	return &newInstance
}

func (bezier *Bezier) GongCopy() GongstructIF {
	newInstance := *bezier
	return &newInstance
}

func (beziergrid *BezierGrid) GongCopy() GongstructIF {
	newInstance := *beziergrid
	return &newInstance
}

func (beziergridstack *BezierGridStack) GongCopy() GongstructIF {
	newInstance := *beziergridstack
	return &newInstance
}

func (chapter *Chapter) GongCopy() GongstructIF {
	newInstance := *chapter
	return &newInstance
}

func (circle *Circle) GongCopy() GongstructIF {
	newInstance := *circle
	return &newInstance
}

func (circlegrid *CircleGrid) GongCopy() GongstructIF {
	newInstance := *circlegrid
	return &newInstance
}

func (content *Content) GongCopy() GongstructIF {
	newInstance := *content
	return &newInstance
}

func (exporttomusicxml *ExportToMusicxml) GongCopy() GongstructIF {
	newInstance := *exporttomusicxml
	return &newInstance
}

func (frontcurve *FrontCurve) GongCopy() GongstructIF {
	newInstance := *frontcurve
	return &newInstance
}

func (frontcurvestack *FrontCurveStack) GongCopy() GongstructIF {
	newInstance := *frontcurvestack
	return &newInstance
}

func (horizontalaxis *HorizontalAxis) GongCopy() GongstructIF {
	newInstance := *horizontalaxis
	return &newInstance
}

func (key *Key) GongCopy() GongstructIF {
	newInstance := *key
	return &newInstance
}

func (parameter *Parameter) GongCopy() GongstructIF {
	newInstance := *parameter
	return &newInstance
}

func (rhombus *Rhombus) GongCopy() GongstructIF {
	newInstance := *rhombus
	return &newInstance
}

func (rhombusgrid *RhombusGrid) GongCopy() GongstructIF {
	newInstance := *rhombusgrid
	return &newInstance
}

func (shapecategory *ShapeCategory) GongCopy() GongstructIF {
	newInstance := *shapecategory
	return &newInstance
}

func (spiralbezier *SpiralBezier) GongCopy() GongstructIF {
	newInstance := *spiralbezier
	return &newInstance
}

func (spiralbeziergrid *SpiralBezierGrid) GongCopy() GongstructIF {
	newInstance := *spiralbeziergrid
	return &newInstance
}

func (spiralcircle *SpiralCircle) GongCopy() GongstructIF {
	newInstance := *spiralcircle
	return &newInstance
}

func (spiralcirclegrid *SpiralCircleGrid) GongCopy() GongstructIF {
	newInstance := *spiralcirclegrid
	return &newInstance
}

func (spiralline *SpiralLine) GongCopy() GongstructIF {
	newInstance := *spiralline
	return &newInstance
}

func (spirallinegrid *SpiralLineGrid) GongCopy() GongstructIF {
	newInstance := *spirallinegrid
	return &newInstance
}

func (spiralorigin *SpiralOrigin) GongCopy() GongstructIF {
	newInstance := *spiralorigin
	return &newInstance
}

func (spiralrhombus *SpiralRhombus) GongCopy() GongstructIF {
	newInstance := *spiralrhombus
	return &newInstance
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongCopy() GongstructIF {
	newInstance := *spiralrhombusgrid
	return &newInstance
}

func (verticalaxis *VerticalAxis) GongCopy() GongstructIF {
	newInstance := *verticalaxis
	return &newInstance
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
	stage.new = make(map[GongstructIF]struct{})
	stage.modified = make(map[GongstructIF]struct{})
	stage.deleted = make(map[GongstructIF]struct{})
}
