package svg

import (
	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func drawHorizontalAxis(
	gongsvgStage *gongsvg_models.StageStruct,
	axisLayer *gongsvg_models.Layer,
	parameter *phylotaxymusic_models.Parameter,

) {
	// creation of 2 transparant rects, one at each ends of the vertical
	horizontalAxisLeftHandle := new(gongsvg_models.Rect).Stage(gongsvgStage)
	horizontalAxisLeftHandle.Name = "Horizontal axis left handle"
	axisLayer.Rects = append(axisLayer.Rects, horizontalAxisLeftHandle)
	horizontalAxisLeftHandle.X = parameter.OriginX - parameter.AxisHandleBorderLength
	horizontalAxisLeftHandle.Y = parameter.OriginY - parameter.AxisHandleBorderLength/2.0
	horizontalAxisLeftHandle.Width = parameter.AxisHandleBorderLength
	horizontalAxisLeftHandle.Height = parameter.AxisHandleBorderLength

	horizontalAxisRightHandle := new(gongsvg_models.Rect).Stage(gongsvgStage)
	horizontalAxisRightHandle.Name = "Horizontal axis rigth handle"
	axisLayer.Rects = append(axisLayer.Rects, horizontalAxisRightHandle)

	horizontalAxisRightHandle.X = parameter.OriginX + parameter.HorizontalAxis_Length
	horizontalAxisRightHandle.Y = parameter.OriginY - parameter.AxisHandleBorderLength/2.0
	horizontalAxisRightHandle.Width = parameter.AxisHandleBorderLength
	horizontalAxisRightHandle.Height = parameter.AxisHandleBorderLength

	horizontalAxis := new(gongsvg_models.Link).Stage(gongsvgStage)
	axisLayer.Links = append(axisLayer.Links, horizontalAxis)

	horizontalAxis.StrokeWidth = parameter.Axis_StrokeWidth
	horizontalAxis.StrokeOpacity = 1
	horizontalAxis.Name = "Horizontal Axis"
	horizontalAxis.Stroke = gongsvg_models.Black.ToString()
	horizontalAxis.Start = horizontalAxisLeftHandle
	horizontalAxis.End = horizontalAxisRightHandle

	horizontalAxis.HasStartArrow = false
	horizontalAxis.HasEndArrow = true

	horizontalAxis.CornerOffsetRatio = 2.0

	horizontalAxis.EndArrowSize = 8
	horizontalAxis.Type = gongsvg_models.LINK_TYPE_FLOATING_ORTHOGONAL

	horizontalAxis.StartOrientation = gongsvg_models.ORIENTATION_HORIZONTAL
	horizontalAxis.EndOrientation = gongsvg_models.ORIENTATION_HORIZONTAL

	horizontalAxis.StartRatio = 0.5
	horizontalAxis.EndRatio = 0.5
}
