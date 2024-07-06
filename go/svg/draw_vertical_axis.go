package svg

import (
	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func drawVerticalAxis(
	gongsvgStage *gongsvg_models.StageStruct,
	axisLayer *gongsvg_models.Layer,
	verticalAxis *phylotaxymusic_models.VerticalAxis,

) {
	// creation of 2 transparant rects, one at each ends of the vertical
	verticalAxisLeftHandle := new(gongsvg_models.Rect).Stage(gongsvgStage)
	verticalAxisLeftHandle.Name = "Vertical axis bottom handle"
	axisLayer.Rects = append(axisLayer.Rects, verticalAxisLeftHandle)
	verticalAxisLeftHandle.X = verticalAxis.OriginX - verticalAxis.AxisHandleBorderLength/2.0
	verticalAxisLeftHandle.Y = verticalAxis.OriginY - verticalAxis.AxisHandleBorderLength
	verticalAxisLeftHandle.Width = verticalAxis.AxisHandleBorderLength
	verticalAxisLeftHandle.Height = verticalAxis.AxisHandleBorderLength

	verticalAxisRightHandle := new(gongsvg_models.Rect).Stage(gongsvgStage)
	verticalAxisRightHandle.Name = "Vertical axis top handle"
	axisLayer.Rects = append(axisLayer.Rects, verticalAxisRightHandle)

	verticalAxisRightHandle.X = verticalAxis.OriginX - verticalAxis.AxisHandleBorderLength/2.0
	verticalAxisRightHandle.Y = verticalAxis.OriginY - verticalAxis.Axis_Length
	verticalAxisRightHandle.Width = verticalAxis.AxisHandleBorderLength
	verticalAxisRightHandle.Height = verticalAxis.AxisHandleBorderLength

	verticalAxisLine := new(gongsvg_models.Link).Stage(gongsvgStage)
	axisLayer.Links = append(axisLayer.Links, verticalAxisLine)

	verticalAxisLine.StrokeWidth = verticalAxis.Axis_StrokeWidth
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
