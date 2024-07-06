package svg

import (
	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func drawHorizontalAxis(
	gongsvgStage *gongsvg_models.StageStruct,
	axisLayer *gongsvg_models.Layer,
	horizontalAxis *phylotaxymusic_models.HorizontalAxis,

) {
	// creation of 2 transparant rects, one at each ends of the vertical
	horizontalAxisLeftHandle := new(gongsvg_models.Rect).Stage(gongsvgStage)
	horizontalAxisLeftHandle.Name = "Horizontal axis left handle"
	axisLayer.Rects = append(axisLayer.Rects, horizontalAxisLeftHandle)
	horizontalAxisLeftHandle.X = horizontalAxis.OriginX - horizontalAxis.AxisHandleBorderLength
	horizontalAxisLeftHandle.Y = horizontalAxis.OriginY - horizontalAxis.AxisHandleBorderLength/2.0
	horizontalAxisLeftHandle.Width = horizontalAxis.AxisHandleBorderLength
	horizontalAxisLeftHandle.Height = horizontalAxis.AxisHandleBorderLength

	horizontalAxisRightHandle := new(gongsvg_models.Rect).Stage(gongsvgStage)
	horizontalAxisRightHandle.Name = "Horizontal axis rigth handle"
	axisLayer.Rects = append(axisLayer.Rects, horizontalAxisRightHandle)

	horizontalAxisRightHandle.X = horizontalAxis.OriginX + horizontalAxis.HorizontalAxis_Length
	horizontalAxisRightHandle.Y = horizontalAxis.OriginY - horizontalAxis.AxisHandleBorderLength/2.0
	horizontalAxisRightHandle.Width = horizontalAxis.AxisHandleBorderLength
	horizontalAxisRightHandle.Height = horizontalAxis.AxisHandleBorderLength

	horizontalAxisLine := new(gongsvg_models.Link).Stage(gongsvgStage)
	axisLayer.Links = append(axisLayer.Links, horizontalAxisLine)

	horizontalAxisLine.StrokeWidth = horizontalAxis.Axis_StrokeWidth
	horizontalAxisLine.StrokeOpacity = 1
	horizontalAxisLine.Name = "Horizontal Axis"
	horizontalAxisLine.Stroke = gongsvg_models.Black.ToString()
	horizontalAxisLine.Start = horizontalAxisLeftHandle
	horizontalAxisLine.End = horizontalAxisRightHandle

	horizontalAxisLine.HasStartArrow = false
	horizontalAxisLine.HasEndArrow = true

	horizontalAxisLine.CornerOffsetRatio = 2.0

	horizontalAxisLine.EndArrowSize = 8
	horizontalAxisLine.Type = gongsvg_models.LINK_TYPE_FLOATING_ORTHOGONAL

	horizontalAxisLine.StartOrientation = gongsvg_models.ORIENTATION_HORIZONTAL
	horizontalAxisLine.EndOrientation = gongsvg_models.ORIENTATION_HORIZONTAL

	horizontalAxisLine.StartRatio = 0.5
	horizontalAxisLine.EndRatio = 0.5
}
