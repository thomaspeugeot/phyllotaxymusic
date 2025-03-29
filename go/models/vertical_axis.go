package models

import (
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type VerticalAxis struct {
	Name string

	Shape
	AxisHandleBorderLength float64
	Axis_Length            float64

	Presentation
}

func (verticalAxis *VerticalAxis) Draw(
	gongsvgStage *gongsvg_models.Stage,
	layer *gongsvg_models.Layer,
	parameter *Parameter,

) {
	// creation of 2 transparant rects, one at each ends of the vertical
	verticalAxisLeftHandle := new(gongsvg_models.Rect).Stage(gongsvgStage)
	verticalAxisLeftHandle.Name = "Vertical axis bottom handle"
	layer.Rects = append(layer.Rects, verticalAxisLeftHandle)
	verticalAxisLeftHandle.X = parameter.OriginX - verticalAxis.AxisHandleBorderLength/2.0
	verticalAxisLeftHandle.Y = parameter.OriginY - verticalAxis.AxisHandleBorderLength
	verticalAxisLeftHandle.Width = verticalAxis.AxisHandleBorderLength
	verticalAxisLeftHandle.Height = verticalAxis.AxisHandleBorderLength

	verticalAxisRightHandle := new(gongsvg_models.Rect).Stage(gongsvgStage)
	verticalAxisRightHandle.Name = "Vertical axis top handle"
	layer.Rects = append(layer.Rects, verticalAxisRightHandle)

	verticalAxisRightHandle.X = parameter.OriginX - verticalAxis.AxisHandleBorderLength/2.0
	verticalAxisRightHandle.Y = parameter.OriginY - verticalAxis.Axis_Length
	verticalAxisRightHandle.Width = verticalAxis.AxisHandleBorderLength
	verticalAxisRightHandle.Height = verticalAxis.AxisHandleBorderLength

	verticalAxisLine := new(gongsvg_models.Link).Stage(gongsvgStage)
	layer.Links = append(layer.Links, verticalAxisLine)

	verticalAxisLine.StrokeWidth = verticalAxis.StrokeWidth
	verticalAxisLine.StrokeOpacity = 1
	verticalAxisLine.Name = "Vertical Axis"
	verticalAxisLine.Stroke = gongsvg_models.Black.ToString()
	verticalAxisLine.Start = verticalAxisLeftHandle
	verticalAxisLine.End = verticalAxisRightHandle

	verticalAxisLine.HasStartArrow = false
	verticalAxisLine.HasEndArrow = true

	verticalAxisLine.CornerOffsetRatio = 2.0

	verticalAxisLine.EndArrowSize = 8
	verticalAxisLine.Type = gongsvg_models.LINK_TYPE_FLOATING_ORTHOGONAL

	verticalAxisLine.StartOrientation = gongsvg_models.ORIENTATION_VERTICAL
	verticalAxisLine.EndOrientation = gongsvg_models.ORIENTATION_VERTICAL

	verticalAxisLine.StartRatio = 0.5
	verticalAxisLine.EndRatio = 0.5
}
