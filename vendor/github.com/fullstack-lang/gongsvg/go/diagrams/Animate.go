package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongsvg/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Animate models.StageStruct
var ___dummy__Time_Animate time.Time

// _ point for meta package dummy declaration
var ___dummy__ref_models_Animate ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.ANCHOR_BOTTOM": ref_models.ANCHOR_BOTTOM,

	"ref_models.ANCHOR_CENTER": ref_models.ANCHOR_CENTER,

	"ref_models.ANCHOR_LEFT": ref_models.ANCHOR_LEFT,

	"ref_models.ANCHOR_RIGHT": ref_models.ANCHOR_RIGHT,

	"ref_models.ANCHOR_TOP": ref_models.ANCHOR_TOP,

	"ref_models.Aliceblue": ref_models.Aliceblue,

	"ref_models.AnchorType": ref_models.AnchorType(""),

	"ref_models.Animate": &(ref_models.Animate{}),

	"ref_models.Animate.AttributeName": (ref_models.Animate{}).AttributeName,

	"ref_models.Animate.Dur": (ref_models.Animate{}).Dur,

	"ref_models.Animate.Name": (ref_models.Animate{}).Name,

	"ref_models.Animate.RepeatCount": (ref_models.Animate{}).RepeatCount,

	"ref_models.Animate.Values": (ref_models.Animate{}).Values,

	"ref_models.Antiquewhite": ref_models.Antiquewhite,

	"ref_models.Aqua": ref_models.Aqua,

	"ref_models.Aquamarine": ref_models.Aquamarine,

	"ref_models.Azure": ref_models.Azure,

	"ref_models.Beige": ref_models.Beige,

	"ref_models.Bisque": ref_models.Bisque,

	"ref_models.Black": ref_models.Black,

	"ref_models.Blanchedalmond": ref_models.Blanchedalmond,

	"ref_models.Blue": ref_models.Blue,

	"ref_models.Blueviolet": ref_models.Blueviolet,

	"ref_models.Brown": ref_models.Brown,

	"ref_models.Burlywood": ref_models.Burlywood,

	"ref_models.Cadetblue": ref_models.Cadetblue,

	"ref_models.Chartreuse": ref_models.Chartreuse,

	"ref_models.Chocolate": ref_models.Chocolate,

	"ref_models.Circle": &(ref_models.Circle{}),

	"ref_models.Circle.Animations": (ref_models.Circle{}).Animations,

	"ref_models.Circle.CX": (ref_models.Circle{}).CX,

	"ref_models.Circle.CY": (ref_models.Circle{}).CY,

	"ref_models.Circle.Color": (ref_models.Circle{}).Color,

	"ref_models.Circle.FillOpacity": (ref_models.Circle{}).FillOpacity,

	"ref_models.Circle.Name": (ref_models.Circle{}).Name,

	"ref_models.Circle.Radius": (ref_models.Circle{}).Radius,

	"ref_models.Circle.Stroke": (ref_models.Circle{}).Stroke,

	"ref_models.Circle.StrokeDashArray": (ref_models.Circle{}).StrokeDashArray,

	"ref_models.Circle.StrokeDashArrayWhenSelected": (ref_models.Circle{}).StrokeDashArrayWhenSelected,

	"ref_models.Circle.StrokeWidth": (ref_models.Circle{}).StrokeWidth,

	"ref_models.Circle.Transform": (ref_models.Circle{}).Transform,

	"ref_models.ColorType": ref_models.ColorType(""),

	"ref_models.Coral": ref_models.Coral,

	"ref_models.Cornflowerblue": ref_models.Cornflowerblue,

	"ref_models.Cornsilk": ref_models.Cornsilk,

	"ref_models.Crimson": ref_models.Crimson,

	"ref_models.Cyan": ref_models.Cyan,

	"ref_models.DRAWING_LINE": ref_models.DRAWING_LINK,

	"ref_models.Darkblue": ref_models.Darkblue,

	"ref_models.Darkcyan": ref_models.Darkcyan,

	"ref_models.Darkgoldenrod": ref_models.Darkgoldenrod,

	"ref_models.Darkgray": ref_models.Darkgray,

	"ref_models.Darkgreen": ref_models.Darkgreen,

	"ref_models.Darkgrey": ref_models.Darkgrey,

	"ref_models.Darkkhaki": ref_models.Darkkhaki,

	"ref_models.Darkmagenta": ref_models.Darkmagenta,

	"ref_models.Darkolivegreen": ref_models.Darkolivegreen,

	"ref_models.Darkorange": ref_models.Darkorange,

	"ref_models.Darkorchid": ref_models.Darkorchid,

	"ref_models.Darkred": ref_models.Darkred,

	"ref_models.Darksalmon": ref_models.Darksalmon,

	"ref_models.Darkseagreen": ref_models.Darkseagreen,

	"ref_models.Darkslateblue": ref_models.Darkslateblue,

	"ref_models.Darkslategray": ref_models.Darkslategray,

	"ref_models.Darkslategrey": ref_models.Darkslategrey,

	"ref_models.Darkturquoise": ref_models.Darkturquoise,

	"ref_models.Darkviolet": ref_models.Darkviolet,

	"ref_models.Deeppink": ref_models.Deeppink,

	"ref_models.Deepskyblue": ref_models.Deepskyblue,

	"ref_models.Dimgray": ref_models.Dimgray,

	"ref_models.Dimgrey": ref_models.Dimgrey,

	"ref_models.Dodgerblue": ref_models.Dodgerblue,

	"ref_models.DrawingState": ref_models.DrawingState(""),

	"ref_models.Ellipse": &(ref_models.Ellipse{}),

	"ref_models.Ellipse.Animates": (ref_models.Ellipse{}).Animates,

	"ref_models.Ellipse.CX": (ref_models.Ellipse{}).CX,

	"ref_models.Ellipse.CY": (ref_models.Ellipse{}).CY,

	"ref_models.Ellipse.Color": (ref_models.Ellipse{}).Color,

	"ref_models.Ellipse.FillOpacity": (ref_models.Ellipse{}).FillOpacity,

	"ref_models.Ellipse.Name": (ref_models.Ellipse{}).Name,

	"ref_models.Ellipse.RX": (ref_models.Ellipse{}).RX,

	"ref_models.Ellipse.RY": (ref_models.Ellipse{}).RY,

	"ref_models.Ellipse.Stroke": (ref_models.Ellipse{}).Stroke,

	"ref_models.Ellipse.StrokeDashArray": (ref_models.Ellipse{}).StrokeDashArray,

	"ref_models.Ellipse.StrokeDashArrayWhenSelected": (ref_models.Ellipse{}).StrokeDashArrayWhenSelected,

	"ref_models.Ellipse.StrokeWidth": (ref_models.Ellipse{}).StrokeWidth,

	"ref_models.Ellipse.Transform": (ref_models.Ellipse{}).Transform,

	"ref_models.Firebrick": ref_models.Firebrick,

	"ref_models.Floralwhite": ref_models.Floralwhite,

	"ref_models.Forestgreen": ref_models.Forestgreen,

	"ref_models.Fuchsia": ref_models.Fuchsia,

	"ref_models.Gainsboro": ref_models.Gainsboro,

	"ref_models.Ghostwhite": ref_models.Ghostwhite,

	"ref_models.Gold": ref_models.Gold,

	"ref_models.Goldenrod": ref_models.Goldenrod,

	"ref_models.Gray": ref_models.Gray,

	"ref_models.Green": ref_models.Green,

	"ref_models.Greenyellow": ref_models.Greenyellow,

	"ref_models.Grey": ref_models.Grey,

	"ref_models.Honeydew": ref_models.Honeydew,

	"ref_models.Hotpink": ref_models.Hotpink,

	"ref_models.Indianred": ref_models.Indianred,

	"ref_models.Indigo": ref_models.Indigo,

	"ref_models.Ivory": ref_models.Ivory,

	"ref_models.Khaki": ref_models.Khaki,

	"ref_models.LINK_TYPE_FLOATING_ORTHOGONAL": ref_models.LINK_TYPE_FLOATING_ORTHOGONAL,

	"ref_models.LINK_TYPE_LINE_WITH_CONTROL_POINTS": ref_models.LINK_TYPE_LINE_WITH_CONTROL_POINTS,

	"ref_models.Lavender": ref_models.Lavender,

	"ref_models.Lavenderblush": ref_models.Lavenderblush,

	"ref_models.Lawngreen": ref_models.Lawngreen,

	"ref_models.Layer": &(ref_models.Layer{}),

	"ref_models.Layer.Circles": (ref_models.Layer{}).Circles,

	"ref_models.Layer.Display": (ref_models.Layer{}).Display,

	"ref_models.Layer.Ellipses": (ref_models.Layer{}).Ellipses,

	"ref_models.Layer.Lines": (ref_models.Layer{}).Lines,

	"ref_models.Layer.Links": (ref_models.Layer{}).Links,

	"ref_models.Layer.Name": (ref_models.Layer{}).Name,

	"ref_models.Layer.Paths": (ref_models.Layer{}).Paths,

	"ref_models.Layer.Polygones": (ref_models.Layer{}).Polygones,

	"ref_models.Layer.Polylines": (ref_models.Layer{}).Polylines,

	"ref_models.Layer.RectLinkLinks": (ref_models.Layer{}).RectLinkLinks,

	"ref_models.Layer.Rects": (ref_models.Layer{}).Rects,

	"ref_models.Layer.Texts": (ref_models.Layer{}).Texts,

	"ref_models.Lemonchiffon": ref_models.Lemonchiffon,

	"ref_models.Lightblue": ref_models.Lightblue,

	"ref_models.Lightcoral": ref_models.Lightcoral,

	"ref_models.Lightcyan": ref_models.Lightcyan,

	"ref_models.Lightgoldenrodyellow": ref_models.Lightgoldenrodyellow,

	"ref_models.Lightgray": ref_models.Lightgray,

	"ref_models.Lightgreen": ref_models.Lightgreen,

	"ref_models.Lightgrey": ref_models.Lightgrey,

	"ref_models.Lightpink": ref_models.Lightpink,

	"ref_models.Lightsalmon": ref_models.Lightsalmon,

	"ref_models.Lightseagreen": ref_models.Lightseagreen,

	"ref_models.Lightskyblue": ref_models.Lightskyblue,

	"ref_models.Lightslategray": ref_models.Lightslategray,

	"ref_models.Lightslategrey": ref_models.Lightslategrey,

	"ref_models.Lightsteelblue": ref_models.Lightsteelblue,

	"ref_models.Lightyellow": ref_models.Lightyellow,

	"ref_models.Lime": ref_models.Lime,

	"ref_models.Limegreen": ref_models.Limegreen,

	"ref_models.Line": &(ref_models.Line{}),

	"ref_models.Line.Animates": (ref_models.Line{}).Animates,

	"ref_models.Line.Color": (ref_models.Line{}).Color,

	"ref_models.Line.FillOpacity": (ref_models.Line{}).FillOpacity,

	"ref_models.Line.MouseClickX": (ref_models.Line{}).MouseClickX,

	"ref_models.Line.MouseClickY": (ref_models.Line{}).MouseClickY,

	"ref_models.Line.Name": (ref_models.Line{}).Name,

	"ref_models.Line.Stroke": (ref_models.Line{}).Stroke,

	"ref_models.Line.StrokeDashArray": (ref_models.Line{}).StrokeDashArray,

	"ref_models.Line.StrokeDashArrayWhenSelected": (ref_models.Line{}).StrokeDashArrayWhenSelected,

	"ref_models.Line.StrokeWidth": (ref_models.Line{}).StrokeWidth,

	"ref_models.Line.Transform": (ref_models.Line{}).Transform,

	"ref_models.Line.X1": (ref_models.Line{}).X1,

	"ref_models.Line.X2": (ref_models.Line{}).X2,

	"ref_models.Line.Y1": (ref_models.Line{}).Y1,

	"ref_models.Line.Y2": (ref_models.Line{}).Y2,

	"ref_models.Linen": ref_models.Linen,

	"ref_models.Link": &(ref_models.Link{}),

	"ref_models.Link.Color": (ref_models.Link{}).Color,

	"ref_models.Link.ControlPoints": (ref_models.Link{}).ControlPoints,

	"ref_models.Link.CornerOffsetRatio": (ref_models.Link{}).CornerOffsetRatio,

	"ref_models.Link.CornerRadius": (ref_models.Link{}).CornerRadius,

	"ref_models.Link.End": (ref_models.Link{}).End,

	"ref_models.Link.EndAnchorType": (ref_models.Link{}).EndAnchorType,

	"ref_models.Link.EndArrowSize": (ref_models.Link{}).EndArrowSize,

	"ref_models.Link.EndOrientation": (ref_models.Link{}).EndOrientation,

	"ref_models.Link.EndRatio": (ref_models.Link{}).EndRatio,

	"ref_models.Link.FillOpacity": (ref_models.Link{}).FillOpacity,

	"ref_models.Link.HasEndArrow": (ref_models.Link{}).HasEndArrow,

	"ref_models.Link.Name": (ref_models.Link{}).Name,

	"ref_models.Link.Start": (ref_models.Link{}).Start,

	"ref_models.Link.StartAnchorType": (ref_models.Link{}).StartAnchorType,

	"ref_models.Link.StartOrientation": (ref_models.Link{}).StartOrientation,

	"ref_models.Link.StartRatio": (ref_models.Link{}).StartRatio,

	"ref_models.Link.Stroke": (ref_models.Link{}).Stroke,

	"ref_models.Link.StrokeDashArray": (ref_models.Link{}).StrokeDashArray,

	"ref_models.Link.StrokeDashArrayWhenSelected": (ref_models.Link{}).StrokeDashArrayWhenSelected,

	"ref_models.Link.StrokeWidth": (ref_models.Link{}).StrokeWidth,

	"ref_models.Link.TextAtArrowEnd": (ref_models.Link{}).TextAtArrowEnd,

	"ref_models.Link.TextAtArrowStart": (ref_models.Link{}).TextAtArrowStart,

	"ref_models.Link.Transform": (ref_models.Link{}).Transform,

	"ref_models.Link.Type": (ref_models.Link{}).Type,

	"ref_models.LinkAnchoredText": &(ref_models.LinkAnchoredText{}),

	"ref_models.LinkAnchoredText.Animates": (ref_models.LinkAnchoredText{}).Animates,

	"ref_models.LinkAnchoredText.Color": (ref_models.LinkAnchoredText{}).Color,

	"ref_models.LinkAnchoredText.Content": (ref_models.LinkAnchoredText{}).Content,

	"ref_models.LinkAnchoredText.FillOpacity": (ref_models.LinkAnchoredText{}).FillOpacity,

	"ref_models.LinkAnchoredText.FontWeight": (ref_models.LinkAnchoredText{}).FontWeight,

	"ref_models.LinkAnchoredText.Name": (ref_models.LinkAnchoredText{}).Name,

	"ref_models.LinkAnchoredText.Stroke": (ref_models.LinkAnchoredText{}).Stroke,

	"ref_models.LinkAnchoredText.StrokeDashArray": (ref_models.LinkAnchoredText{}).StrokeDashArray,

	"ref_models.LinkAnchoredText.StrokeDashArrayWhenSelected": (ref_models.LinkAnchoredText{}).StrokeDashArrayWhenSelected,

	"ref_models.LinkAnchoredText.StrokeWidth": (ref_models.LinkAnchoredText{}).StrokeWidth,

	"ref_models.LinkAnchoredText.Transform": (ref_models.LinkAnchoredText{}).Transform,

	"ref_models.LinkAnchoredText.X_Offset": (ref_models.LinkAnchoredText{}).X_Offset,

	"ref_models.LinkAnchoredText.Y_Offset": (ref_models.LinkAnchoredText{}).Y_Offset,

	"ref_models.LinkType": ref_models.LinkType(""),

	"ref_models.Magenta": ref_models.Magenta,

	"ref_models.Maroon": ref_models.Maroon,

	"ref_models.Mediumaquamarine": ref_models.Mediumaquamarine,

	"ref_models.Mediumblue": ref_models.Mediumblue,

	"ref_models.Mediumorchid": ref_models.Mediumorchid,

	"ref_models.Mediumpurple": ref_models.Mediumpurple,

	"ref_models.Mediumseagreen": ref_models.Mediumseagreen,

	"ref_models.Mediumslateblue": ref_models.Mediumslateblue,

	"ref_models.Mediumspringgreen": ref_models.Mediumspringgreen,

	"ref_models.Mediumturquoise": ref_models.Mediumturquoise,

	"ref_models.Mediumvioletred": ref_models.Mediumvioletred,

	"ref_models.Midnightblue": ref_models.Midnightblue,

	"ref_models.Mintcream": ref_models.Mintcream,

	"ref_models.Mistyrose": ref_models.Mistyrose,

	"ref_models.Moccasin": ref_models.Moccasin,

	"ref_models.NOT_DRAWING_LINE": ref_models.NOT_DRAWING_LINK,

	"ref_models.Navajowhite": ref_models.Navajowhite,

	"ref_models.Navy": ref_models.Navy,

	"ref_models.ORIENTATION_HORIZONTAL": ref_models.ORIENTATION_HORIZONTAL,

	"ref_models.ORIENTATION_VERTICAL": ref_models.ORIENTATION_VERTICAL,

	"ref_models.Oldlace": ref_models.Oldlace,

	"ref_models.Olive": ref_models.Olive,

	"ref_models.Olivedrab": ref_models.Olivedrab,

	"ref_models.Orange": ref_models.Orange,

	"ref_models.Orangered": ref_models.Orangered,

	"ref_models.Orchid": ref_models.Orchid,

	"ref_models.OrientationType": ref_models.OrientationType(""),

	"ref_models.POSITION_ON_ARROW_END": ref_models.POSITION_ON_ARROW_END,

	"ref_models.POSITION_ON_ARROW_START": ref_models.POSITION_ON_ARROW_START,

	"ref_models.Palegoldenrod": ref_models.Palegoldenrod,

	"ref_models.Palegreen": ref_models.Palegreen,

	"ref_models.Paleturquoise": ref_models.Paleturquoise,

	"ref_models.Palevioletred": ref_models.Palevioletred,

	"ref_models.Papayawhip": ref_models.Papayawhip,

	"ref_models.Path": &(ref_models.Path{}),

	"ref_models.Path.Animates": (ref_models.Path{}).Animates,

	"ref_models.Path.Color": (ref_models.Path{}).Color,

	"ref_models.Path.Definition": (ref_models.Path{}).Definition,

	"ref_models.Path.FillOpacity": (ref_models.Path{}).FillOpacity,

	"ref_models.Path.Name": (ref_models.Path{}).Name,

	"ref_models.Path.Stroke": (ref_models.Path{}).Stroke,

	"ref_models.Path.StrokeDashArray": (ref_models.Path{}).StrokeDashArray,

	"ref_models.Path.StrokeDashArrayWhenSelected": (ref_models.Path{}).StrokeDashArrayWhenSelected,

	"ref_models.Path.StrokeWidth": (ref_models.Path{}).StrokeWidth,

	"ref_models.Path.Transform": (ref_models.Path{}).Transform,

	"ref_models.Peachpuff": ref_models.Peachpuff,

	"ref_models.Peru": ref_models.Peru,

	"ref_models.Pink": ref_models.Pink,

	"ref_models.Plum": ref_models.Plum,

	"ref_models.Point": &(ref_models.Point{}),

	"ref_models.Point.Name": (ref_models.Point{}).Name,

	"ref_models.Point.X": (ref_models.Point{}).X,

	"ref_models.Point.Y": (ref_models.Point{}).Y,

	"ref_models.Polygone": &(ref_models.Polygone{}),

	"ref_models.Polygone.Animates": (ref_models.Polygone{}).Animates,

	"ref_models.Polygone.Color": (ref_models.Polygone{}).Color,

	"ref_models.Polygone.FillOpacity": (ref_models.Polygone{}).FillOpacity,

	"ref_models.Polygone.Name": (ref_models.Polygone{}).Name,

	"ref_models.Polygone.Points": (ref_models.Polygone{}).Points,

	"ref_models.Polygone.Stroke": (ref_models.Polygone{}).Stroke,

	"ref_models.Polygone.StrokeDashArray": (ref_models.Polygone{}).StrokeDashArray,

	"ref_models.Polygone.StrokeDashArrayWhenSelected": (ref_models.Polygone{}).StrokeDashArrayWhenSelected,

	"ref_models.Polygone.StrokeWidth": (ref_models.Polygone{}).StrokeWidth,

	"ref_models.Polygone.Transform": (ref_models.Polygone{}).Transform,

	"ref_models.Polyline": &(ref_models.Polyline{}),

	"ref_models.Polyline.Animates": (ref_models.Polyline{}).Animates,

	"ref_models.Polyline.Color": (ref_models.Polyline{}).Color,

	"ref_models.Polyline.FillOpacity": (ref_models.Polyline{}).FillOpacity,

	"ref_models.Polyline.Name": (ref_models.Polyline{}).Name,

	"ref_models.Polyline.Points": (ref_models.Polyline{}).Points,

	"ref_models.Polyline.Stroke": (ref_models.Polyline{}).Stroke,

	"ref_models.Polyline.StrokeDashArray": (ref_models.Polyline{}).StrokeDashArray,

	"ref_models.Polyline.StrokeDashArrayWhenSelected": (ref_models.Polyline{}).StrokeDashArrayWhenSelected,

	"ref_models.Polyline.StrokeWidth": (ref_models.Polyline{}).StrokeWidth,

	"ref_models.Polyline.Transform": (ref_models.Polyline{}).Transform,

	"ref_models.PositionOnArrowType": ref_models.PositionOnArrowType(""),

	"ref_models.Powderblue": ref_models.Powderblue,

	"ref_models.Purple": ref_models.Purple,

	"ref_models.RECT_ANCHOR_BOTTOM": ref_models.RECT_BOTTOM,

	"ref_models.RECT_ANCHOR_BOTTOM_LEFT": ref_models.RECT_BOTTOM_LEFT,

	"ref_models.RECT_ANCHOR_BOTTOM_RIGHT": ref_models.RECT_BOTTOM_RIGHT,

	"ref_models.RECT_ANCHOR_CENTER": ref_models.RECT_CENTER,

	"ref_models.RECT_ANCHOR_LEFT": ref_models.RECT_LEFT,

	"ref_models.RECT_ANCHOR_RIGHT": ref_models.RECT_RIGHT,

	"ref_models.RECT_ANCHOR_TOP": ref_models.RECT_TOP,

	"ref_models.RECT_ANCHOR_TOP_LEFT": ref_models.RECT_TOP_LEFT,

	"ref_models.RECT_ANCHOR_TOP_RIGHT": ref_models.RECT_TOP_RIGHT,

	"ref_models.Rect": &(ref_models.Rect{}),

	"ref_models.Rect.Animations": (ref_models.Rect{}).Animations,

	"ref_models.Rect.CanHaveBottomHandle": (ref_models.Rect{}).CanHaveBottomHandle,

	"ref_models.Rect.CanHaveLeftHandle": (ref_models.Rect{}).CanHaveLeftHandle,

	"ref_models.Rect.CanHaveRightHandle": (ref_models.Rect{}).CanHaveRightHandle,

	"ref_models.Rect.CanHaveTopHandle": (ref_models.Rect{}).CanHaveTopHandle,

	"ref_models.Rect.CanMoveHorizontaly": (ref_models.Rect{}).CanMoveHorizontaly,

	"ref_models.Rect.CanMoveVerticaly": (ref_models.Rect{}).CanMoveVerticaly,

	"ref_models.Rect.Color": (ref_models.Rect{}).Color,

	"ref_models.Rect.FillOpacity": (ref_models.Rect{}).FillOpacity,

	"ref_models.Rect.HasBottomHandle": (ref_models.Rect{}).HasBottomHandle,

	"ref_models.Rect.HasLeftHandle": (ref_models.Rect{}).HasLeftHandle,

	"ref_models.Rect.HasRightHandle": (ref_models.Rect{}).HasRightHandle,

	"ref_models.Rect.HasTopHandle": (ref_models.Rect{}).HasTopHandle,

	"ref_models.Rect.Height": (ref_models.Rect{}).Height,

	"ref_models.Rect.IsSelectable": (ref_models.Rect{}).IsSelectable,

	"ref_models.Rect.IsSelected": (ref_models.Rect{}).IsSelected,

	"ref_models.Rect.Name": (ref_models.Rect{}).Name,

	"ref_models.Rect.RX": (ref_models.Rect{}).RX,

	"ref_models.Rect.RectAnchoredRects": (ref_models.Rect{}).RectAnchoredRects,

	"ref_models.Rect.RectAnchoredTexts": (ref_models.Rect{}).RectAnchoredTexts,

	"ref_models.Rect.Stroke": (ref_models.Rect{}).Stroke,

	"ref_models.Rect.StrokeDashArray": (ref_models.Rect{}).StrokeDashArray,

	"ref_models.Rect.StrokeDashArrayWhenSelected": (ref_models.Rect{}).StrokeDashArrayWhenSelected,

	"ref_models.Rect.StrokeWidth": (ref_models.Rect{}).StrokeWidth,

	"ref_models.Rect.Transform": (ref_models.Rect{}).Transform,

	"ref_models.Rect.Width": (ref_models.Rect{}).Width,

	"ref_models.Rect.X": (ref_models.Rect{}).X,

	"ref_models.Rect.Y": (ref_models.Rect{}).Y,

	"ref_models.RectAnchorType": ref_models.RectAnchorType(""),

	"ref_models.RectAnchoredRect": &(ref_models.RectAnchoredRect{}),

	"ref_models.RectAnchoredRect.Color": (ref_models.RectAnchoredRect{}).Color,

	"ref_models.RectAnchoredRect.FillOpacity": (ref_models.RectAnchoredRect{}).FillOpacity,

	"ref_models.RectAnchoredRect.Height": (ref_models.RectAnchoredRect{}).Height,

	"ref_models.RectAnchoredRect.HeightFollowRect": (ref_models.RectAnchoredRect{}).HeightFollowRect,

	"ref_models.RectAnchoredRect.Name": (ref_models.RectAnchoredRect{}).Name,

	"ref_models.RectAnchoredRect.RX": (ref_models.RectAnchoredRect{}).RX,

	"ref_models.RectAnchoredRect.RectAnchorType": (ref_models.RectAnchoredRect{}).RectAnchorType,

	"ref_models.RectAnchoredRect.Stroke": (ref_models.RectAnchoredRect{}).Stroke,

	"ref_models.RectAnchoredRect.StrokeDashArray": (ref_models.RectAnchoredRect{}).StrokeDashArray,

	"ref_models.RectAnchoredRect.StrokeDashArrayWhenSelected": (ref_models.RectAnchoredRect{}).StrokeDashArrayWhenSelected,

	"ref_models.RectAnchoredRect.StrokeWidth": (ref_models.RectAnchoredRect{}).StrokeWidth,

	"ref_models.RectAnchoredRect.Transform": (ref_models.RectAnchoredRect{}).Transform,

	"ref_models.RectAnchoredRect.Width": (ref_models.RectAnchoredRect{}).Width,

	"ref_models.RectAnchoredRect.WidthFollowRect": (ref_models.RectAnchoredRect{}).WidthFollowRect,

	"ref_models.RectAnchoredRect.X": (ref_models.RectAnchoredRect{}).X,

	"ref_models.RectAnchoredRect.X_Offset": (ref_models.RectAnchoredRect{}).X_Offset,

	"ref_models.RectAnchoredRect.Y": (ref_models.RectAnchoredRect{}).Y,

	"ref_models.RectAnchoredRect.Y_Offset": (ref_models.RectAnchoredRect{}).Y_Offset,

	"ref_models.RectAnchoredText": &(ref_models.RectAnchoredText{}),

	"ref_models.RectAnchoredText.Animates": (ref_models.RectAnchoredText{}).Animates,

	"ref_models.RectAnchoredText.Color": (ref_models.RectAnchoredText{}).Color,

	"ref_models.RectAnchoredText.Content": (ref_models.RectAnchoredText{}).Content,

	"ref_models.RectAnchoredText.FillOpacity": (ref_models.RectAnchoredText{}).FillOpacity,

	"ref_models.RectAnchoredText.FontSize": (ref_models.RectAnchoredText{}).FontSize,

	"ref_models.RectAnchoredText.FontWeight": (ref_models.RectAnchoredText{}).FontWeight,

	"ref_models.RectAnchoredText.Name": (ref_models.RectAnchoredText{}).Name,

	"ref_models.RectAnchoredText.RectAnchorType": (ref_models.RectAnchoredText{}).RectAnchorType,

	"ref_models.RectAnchoredText.Stroke": (ref_models.RectAnchoredText{}).Stroke,

	"ref_models.RectAnchoredText.StrokeDashArray": (ref_models.RectAnchoredText{}).StrokeDashArray,

	"ref_models.RectAnchoredText.StrokeDashArrayWhenSelected": (ref_models.RectAnchoredText{}).StrokeDashArrayWhenSelected,

	"ref_models.RectAnchoredText.StrokeWidth": (ref_models.RectAnchoredText{}).StrokeWidth,

	"ref_models.RectAnchoredText.TextAnchorType": (ref_models.RectAnchoredText{}).TextAnchorType,

	"ref_models.RectAnchoredText.Transform": (ref_models.RectAnchoredText{}).Transform,

	"ref_models.RectAnchoredText.X_Offset": (ref_models.RectAnchoredText{}).X_Offset,

	"ref_models.RectAnchoredText.Y_Offset": (ref_models.RectAnchoredText{}).Y_Offset,

	"ref_models.RectLinkLink": &(ref_models.RectLinkLink{}),

	"ref_models.RectLinkLink.Color": (ref_models.RectLinkLink{}).Color,

	"ref_models.RectLinkLink.End": (ref_models.RectLinkLink{}).End,

	"ref_models.RectLinkLink.FillOpacity": (ref_models.RectLinkLink{}).FillOpacity,

	"ref_models.RectLinkLink.Name": (ref_models.RectLinkLink{}).Name,

	"ref_models.RectLinkLink.Start": (ref_models.RectLinkLink{}).Start,

	"ref_models.RectLinkLink.Stroke": (ref_models.RectLinkLink{}).Stroke,

	"ref_models.RectLinkLink.StrokeDashArray": (ref_models.RectLinkLink{}).StrokeDashArray,

	"ref_models.RectLinkLink.StrokeDashArrayWhenSelected": (ref_models.RectLinkLink{}).StrokeDashArrayWhenSelected,

	"ref_models.RectLinkLink.StrokeWidth": (ref_models.RectLinkLink{}).StrokeWidth,

	"ref_models.RectLinkLink.TargetAnchorPosition": (ref_models.RectLinkLink{}).TargetAnchorPosition,

	"ref_models.RectLinkLink.Transform": (ref_models.RectLinkLink{}).Transform,

	"ref_models.Red": ref_models.Red,

	"ref_models.Rosybrown": ref_models.Rosybrown,

	"ref_models.Royalblue": ref_models.Royalblue,

	"ref_models.SIDE_BOTTOM": ref_models.SIDE_BOTTOM,

	"ref_models.SIDE_LEFT": ref_models.SIDE_LEFT,

	"ref_models.SIDE_RIGHT": ref_models.SIDE_RIGHT,

	"ref_models.SIDE_TOP": ref_models.SIDE_TOP,

	"ref_models.SVG": &(ref_models.SVG{}),

	"ref_models.SVG.DrawingState": (ref_models.SVG{}).DrawingState,

	"ref_models.SVG.EndRect": (ref_models.SVG{}).EndRect,

	"ref_models.SVG.IsEditable": (ref_models.SVG{}).IsEditable,

	"ref_models.SVG.Layers": (ref_models.SVG{}).Layers,

	"ref_models.SVG.Name": (ref_models.SVG{}).Name,

	"ref_models.SVG.StartRect": (ref_models.SVG{}).StartRect,

	"ref_models.Saddlebrown": ref_models.Saddlebrown,

	"ref_models.Salmon": ref_models.Salmon,

	"ref_models.Sandybrown": ref_models.Sandybrown,

	"ref_models.Seagreen": ref_models.Seagreen,

	"ref_models.Seashell": ref_models.Seashell,

	"ref_models.SideType": ref_models.SideType(""),

	"ref_models.Sienna": ref_models.Sienna,

	"ref_models.Silver": ref_models.Silver,

	"ref_models.Skyblue": ref_models.Skyblue,

	"ref_models.Slateblue": ref_models.Slateblue,

	"ref_models.Slategray": ref_models.Slategray,

	"ref_models.Slategrey": ref_models.Slategrey,

	"ref_models.Snow": ref_models.Snow,

	"ref_models.Springgreen": ref_models.Springgreen,

	"ref_models.StackName": ref_models.StackName(""),

	"ref_models.StackNameDefault": ref_models.StackNameDefault,

	"ref_models.Steelblue": ref_models.Steelblue,

	"ref_models.TEXT_ANCHOR_CENTER": ref_models.TEXT_ANCHOR_CENTER,

	"ref_models.TEXT_ANCHOR_END": ref_models.TEXT_ANCHOR_END,

	"ref_models.TEXT_ANCHOR_START": ref_models.TEXT_ANCHOR_START,

	"ref_models.Tan": ref_models.Tan,

	"ref_models.Teal": ref_models.Teal,

	"ref_models.Text": &(ref_models.Text{}),

	"ref_models.Text.Animates": (ref_models.Text{}).Animates,

	"ref_models.Text.Color": (ref_models.Text{}).Color,

	"ref_models.Text.Content": (ref_models.Text{}).Content,

	"ref_models.Text.FillOpacity": (ref_models.Text{}).FillOpacity,

	"ref_models.Text.Name": (ref_models.Text{}).Name,

	"ref_models.Text.Stroke": (ref_models.Text{}).Stroke,

	"ref_models.Text.StrokeDashArray": (ref_models.Text{}).StrokeDashArray,

	"ref_models.Text.StrokeDashArrayWhenSelected": (ref_models.Text{}).StrokeDashArrayWhenSelected,

	"ref_models.Text.StrokeWidth": (ref_models.Text{}).StrokeWidth,

	"ref_models.Text.Transform": (ref_models.Text{}).Transform,

	"ref_models.Text.X": (ref_models.Text{}).X,

	"ref_models.Text.Y": (ref_models.Text{}).Y,

	"ref_models.TextAnchorType": ref_models.TextAnchorType(""),

	"ref_models.Thistle": ref_models.Thistle,

	"ref_models.Tomato": ref_models.Tomato,

	"ref_models.Turquoise": ref_models.Turquoise,

	"ref_models.Violet": ref_models.Violet,

	"ref_models.Wheat": ref_models.Wheat,

	"ref_models.White": ref_models.White,

	"ref_models.Whitesmoke": ref_models.Whitesmoke,

	"ref_models.Yellow": ref_models.Yellow,

	"ref_models.Yellowgreen": ref_models.Yellowgreen,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the _ gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Animate"] = _
// }

// _ will stage objects of database "Animate"
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Animate := (&models.Classdiagram{Name: `Animate`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_AttributeName := (&models.Field{Name: `AttributeName`}).Stage(stage)
	__Field__000001_Dur := (&models.Field{Name: `Dur`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000004_RepeatCount := (&models.Field{Name: `RepeatCount`}).Stage(stage)
	__Field__000005_Values := (&models.Field{Name: `Values`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Animate_Animate := (&models.GongStructShape{Name: `Animate-Animate`}).Stage(stage)
	__GongStructShape__000001_Animate_Circle := (&models.GongStructShape{Name: `Animate-Circle`}).Stage(stage)
	__GongStructShape__000002_Animate_Ellipse := (&models.GongStructShape{Name: `Animate-Ellipse`}).Stage(stage)
	__GongStructShape__000003_Animate_Line := (&models.GongStructShape{Name: `Animate-Line`}).Stage(stage)
	__GongStructShape__000004_Animate_Rect := (&models.GongStructShape{Name: `Animate-Rect`}).Stage(stage)
	__GongStructShape__000005_Animate_Text := (&models.GongStructShape{Name: `Animate-Text`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Animates := (&models.Link{Name: `Animates`}).Stage(stage)
	__Link__000001_Animates := (&models.Link{Name: `Animates`}).Stage(stage)
	__Link__000002_Animates := (&models.Link{Name: `Animates`}).Stage(stage)
	__Link__000003_Animations := (&models.Link{Name: `Animations`}).Stage(stage)
	__Link__000004_Animations := (&models.Link{Name: `Animations`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Animate_Animate := (&models.Position{Name: `Pos-Animate-Animate`}).Stage(stage)
	__Position__000001_Pos_Animate_Circle := (&models.Position{Name: `Pos-Animate-Circle`}).Stage(stage)
	__Position__000002_Pos_Animate_Ellipse := (&models.Position{Name: `Pos-Animate-Ellipse`}).Stage(stage)
	__Position__000003_Pos_Animate_Line := (&models.Position{Name: `Pos-Animate-Line`}).Stage(stage)
	__Position__000004_Pos_Animate_Rect := (&models.Position{Name: `Pos-Animate-Rect`}).Stage(stage)
	__Position__000005_Pos_Animate_Text := (&models.Position{Name: `Pos-Animate-Text`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Circle_and_Animate_Animate := (&models.Vertice{Name: `Verticle in class diagram Animate in middle between Animate-Circle and Animate-Animate`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Ellipse_and_Animate_Animate := (&models.Vertice{Name: `Verticle in class diagram Animate in middle between Animate-Ellipse and Animate-Animate`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Line_and_Animate_Animate := (&models.Vertice{Name: `Verticle in class diagram Animate in middle between Animate-Line and Animate-Animate`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Rect_and_Animate_Animate := (&models.Vertice{Name: `Verticle in class diagram Animate in middle between Animate-Rect and Animate-Animate`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Text_and_Animate_Animate := (&models.Vertice{Name: `Verticle in class diagram Animate in middle between Animate-Text and Animate-Animate`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Animate.Name = `Animate`
	__Classdiagram__000000_Animate.IsInDrawMode = true

	// Field values setup
	__Field__000000_AttributeName.Name = `AttributeName`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate.AttributeName]
	__Field__000000_AttributeName.Identifier = `ref_models.Animate.AttributeName`
	__Field__000000_AttributeName.FieldTypeAsString = ``
	__Field__000000_AttributeName.Structname = `Animate`
	__Field__000000_AttributeName.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Dur.Name = `Dur`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate.Dur]
	__Field__000001_Dur.Identifier = `ref_models.Animate.Dur`
	__Field__000001_Dur.FieldTypeAsString = ``
	__Field__000001_Dur.Structname = `Animate`
	__Field__000001_Dur.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.Name]
	__Field__000002_Name.Identifier = `ref_models.Rect.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Rect`
	__Field__000002_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate.Name]
	__Field__000003_Name.Identifier = `ref_models.Animate.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `Animate`
	__Field__000003_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000004_RepeatCount.Name = `RepeatCount`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate.RepeatCount]
	__Field__000004_RepeatCount.Identifier = `ref_models.Animate.RepeatCount`
	__Field__000004_RepeatCount.FieldTypeAsString = ``
	__Field__000004_RepeatCount.Structname = `Animate`
	__Field__000004_RepeatCount.Fieldtypename = `string`

	// Field values setup
	__Field__000005_Values.Name = `Values`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate.Values]
	__Field__000005_Values.Identifier = `ref_models.Animate.Values`
	__Field__000005_Values.FieldTypeAsString = ``
	__Field__000005_Values.Structname = `Animate`
	__Field__000005_Values.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_Animate_Animate.Name = `Animate-Animate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate]
	__GongStructShape__000000_Animate_Animate.Identifier = `ref_models.Animate`
	__GongStructShape__000000_Animate_Animate.ShowNbInstances = false
	__GongStructShape__000000_Animate_Animate.NbInstances = 0
	__GongStructShape__000000_Animate_Animate.Width = 240.000000
	__GongStructShape__000000_Animate_Animate.Height = 466.000000
	__GongStructShape__000000_Animate_Animate.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Animate_Circle.Name = `Animate-Circle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle]
	__GongStructShape__000001_Animate_Circle.Identifier = `ref_models.Circle`
	__GongStructShape__000001_Animate_Circle.ShowNbInstances = false
	__GongStructShape__000001_Animate_Circle.NbInstances = 0
	__GongStructShape__000001_Animate_Circle.Width = 240.000000
	__GongStructShape__000001_Animate_Circle.Height = 63.000000
	__GongStructShape__000001_Animate_Circle.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Animate_Ellipse.Name = `Animate-Ellipse`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Ellipse]
	__GongStructShape__000002_Animate_Ellipse.Identifier = `ref_models.Ellipse`
	__GongStructShape__000002_Animate_Ellipse.ShowNbInstances = false
	__GongStructShape__000002_Animate_Ellipse.NbInstances = 0
	__GongStructShape__000002_Animate_Ellipse.Width = 240.000000
	__GongStructShape__000002_Animate_Ellipse.Height = 63.000000
	__GongStructShape__000002_Animate_Ellipse.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_Animate_Line.Name = `Animate-Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__GongStructShape__000003_Animate_Line.Identifier = `ref_models.Line`
	__GongStructShape__000003_Animate_Line.ShowNbInstances = false
	__GongStructShape__000003_Animate_Line.NbInstances = 0
	__GongStructShape__000003_Animate_Line.Width = 240.000000
	__GongStructShape__000003_Animate_Line.Height = 63.000000
	__GongStructShape__000003_Animate_Line.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_Animate_Rect.Name = `Animate-Rect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__GongStructShape__000004_Animate_Rect.Identifier = `ref_models.Rect`
	__GongStructShape__000004_Animate_Rect.ShowNbInstances = false
	__GongStructShape__000004_Animate_Rect.NbInstances = 0
	__GongStructShape__000004_Animate_Rect.Width = 240.000000
	__GongStructShape__000004_Animate_Rect.Height = 78.000000
	__GongStructShape__000004_Animate_Rect.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_Animate_Text.Name = `Animate-Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text]
	__GongStructShape__000005_Animate_Text.Identifier = `ref_models.Text`
	__GongStructShape__000005_Animate_Text.ShowNbInstances = false
	__GongStructShape__000005_Animate_Text.NbInstances = 0
	__GongStructShape__000005_Animate_Text.Width = 240.000000
	__GongStructShape__000005_Animate_Text.Height = 63.000000
	__GongStructShape__000005_Animate_Text.IsSelected = false

	// Link values setup
	__Link__000000_Animates.Name = `Animates`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Ellipse.Animates]
	__Link__000000_Animates.Identifier = `ref_models.Ellipse.Animates`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate]
	__Link__000000_Animates.Fieldtypename = `ref_models.Animate`
	__Link__000000_Animates.FieldOffsetX = -89.000000
	__Link__000000_Animates.FieldOffsetY = -20.000000
	__Link__000000_Animates.TargetMultiplicity = models.MANY
	__Link__000000_Animates.TargetMultiplicityOffsetX = -43.000000
	__Link__000000_Animates.TargetMultiplicityOffsetY = 26.000000
	__Link__000000_Animates.SourceMultiplicity = models.ZERO_ONE
	__Link__000000_Animates.SourceMultiplicityOffsetX = 17.000000
	__Link__000000_Animates.SourceMultiplicityOffsetY = 23.000000
	__Link__000000_Animates.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Animates.StartRatio = 0.500000
	__Link__000000_Animates.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Animates.EndRatio = 0.896996
	__Link__000000_Animates.CornerOffsetRatio = 1.626042

	// Link values setup
	__Link__000001_Animates.Name = `Animates`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Animates]
	__Link__000001_Animates.Identifier = `ref_models.Line.Animates`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate]
	__Link__000001_Animates.Fieldtypename = `ref_models.Animate`
	__Link__000001_Animates.FieldOffsetX = -90.000000
	__Link__000001_Animates.FieldOffsetY = -18.000000
	__Link__000001_Animates.TargetMultiplicity = models.MANY
	__Link__000001_Animates.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_Animates.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_Animates.SourceMultiplicity = models.ZERO_ONE
	__Link__000001_Animates.SourceMultiplicityOffsetX = 18.000000
	__Link__000001_Animates.SourceMultiplicityOffsetY = 23.000000
	__Link__000001_Animates.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Animates.StartRatio = 0.500000
	__Link__000001_Animates.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Animates.EndRatio = 0.738197
	__Link__000001_Animates.CornerOffsetRatio = 1.555208

	// Link values setup
	__Link__000002_Animates.Name = `Animates`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text.Animates]
	__Link__000002_Animates.Identifier = `ref_models.Text.Animates`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate]
	__Link__000002_Animates.Fieldtypename = `ref_models.Animate`
	__Link__000002_Animates.FieldOffsetX = -95.000000
	__Link__000002_Animates.FieldOffsetY = -20.000000
	__Link__000002_Animates.TargetMultiplicity = models.MANY
	__Link__000002_Animates.TargetMultiplicityOffsetX = -45.000000
	__Link__000002_Animates.TargetMultiplicityOffsetY = 26.000000
	__Link__000002_Animates.SourceMultiplicity = models.ZERO_ONE
	__Link__000002_Animates.SourceMultiplicityOffsetX = 14.000000
	__Link__000002_Animates.SourceMultiplicityOffsetY = 24.000000
	__Link__000002_Animates.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Animates.StartRatio = 0.500000
	__Link__000002_Animates.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Animates.EndRatio = 0.360515
	__Link__000002_Animates.CornerOffsetRatio = 1.530208

	// Link values setup
	__Link__000003_Animations.Name = `Animations`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.Animations]
	__Link__000003_Animations.Identifier = `ref_models.Rect.Animations`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate]
	__Link__000003_Animations.Fieldtypename = `ref_models.Animate`
	__Link__000003_Animations.FieldOffsetX = -111.000000
	__Link__000003_Animations.FieldOffsetY = -21.000000
	__Link__000003_Animations.TargetMultiplicity = models.MANY
	__Link__000003_Animations.TargetMultiplicityOffsetX = -40.000000
	__Link__000003_Animations.TargetMultiplicityOffsetY = 36.000000
	__Link__000003_Animations.SourceMultiplicity = models.ZERO_ONE
	__Link__000003_Animations.SourceMultiplicityOffsetX = 15.000000
	__Link__000003_Animations.SourceMultiplicityOffsetY = 24.000000
	__Link__000003_Animations.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_Animations.StartRatio = 0.500000
	__Link__000003_Animations.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_Animations.EndRatio = 0.141631
	__Link__000003_Animations.CornerOffsetRatio = 1.521875

	// Link values setup
	__Link__000004_Animations.Name = `Animations`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.Animations]
	__Link__000004_Animations.Identifier = `ref_models.Circle.Animations`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate]
	__Link__000004_Animations.Fieldtypename = `ref_models.Animate`
	__Link__000004_Animations.FieldOffsetX = -104.000000
	__Link__000004_Animations.FieldOffsetY = -22.000000
	__Link__000004_Animations.TargetMultiplicity = models.MANY
	__Link__000004_Animations.TargetMultiplicityOffsetX = -44.000000
	__Link__000004_Animations.TargetMultiplicityOffsetY = 32.000000
	__Link__000004_Animations.SourceMultiplicity = models.ZERO_ONE
	__Link__000004_Animations.SourceMultiplicityOffsetX = 15.000000
	__Link__000004_Animations.SourceMultiplicityOffsetY = 20.000000
	__Link__000004_Animations.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_Animations.StartRatio = 0.500000
	__Link__000004_Animations.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_Animations.EndRatio = 0.549356
	__Link__000004_Animations.CornerOffsetRatio = 1.805208

	// Position values setup
	__Position__000000_Pos_Animate_Animate.X = 611.000000
	__Position__000000_Pos_Animate_Animate.Y = 139.000000
	__Position__000000_Pos_Animate_Animate.Name = `Pos-Animate-Animate`

	// Position values setup
	__Position__000001_Pos_Animate_Circle.X = 39.000000
	__Position__000001_Pos_Animate_Circle.Y = 364.000000
	__Position__000001_Pos_Animate_Circle.Name = `Pos-Animate-Circle`

	// Position values setup
	__Position__000002_Pos_Animate_Ellipse.X = 38.000000
	__Position__000002_Pos_Animate_Ellipse.Y = 528.000000
	__Position__000002_Pos_Animate_Ellipse.Name = `Pos-Animate-Ellipse`

	// Position values setup
	__Position__000003_Pos_Animate_Line.X = 39.000000
	__Position__000003_Pos_Animate_Line.Y = 447.000000
	__Position__000003_Pos_Animate_Line.Name = `Pos-Animate-Line`

	// Position values setup
	__Position__000004_Pos_Animate_Rect.X = 41.000000
	__Position__000004_Pos_Animate_Rect.Y = 169.000000
	__Position__000004_Pos_Animate_Rect.Name = `Pos-Animate-Rect`

	// Position values setup
	__Position__000005_Pos_Animate_Text.X = 36.000000
	__Position__000005_Pos_Animate_Text.Y = 271.000000
	__Position__000005_Pos_Animate_Text.Name = `Pos-Animate-Text`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Circle_and_Animate_Animate.X = 685.000000
	__Vertice__000000_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Circle_and_Animate_Animate.Y = 283.000000
	__Vertice__000000_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Circle_and_Animate_Animate.Name = `Verticle in class diagram Animate in middle between Animate-Circle and Animate-Animate`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Ellipse_and_Animate_Animate.X = 702.500000
	__Vertice__000001_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Ellipse_and_Animate_Animate.Y = 121.000000
	__Vertice__000001_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Ellipse_and_Animate_Animate.Name = `Verticle in class diagram Animate in middle between Animate-Ellipse and Animate-Animate`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Line_and_Animate_Animate.X = 684.500000
	__Vertice__000002_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Line_and_Animate_Animate.Y = 321.500000
	__Vertice__000002_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Line_and_Animate_Animate.Name = `Verticle in class diagram Animate in middle between Animate-Line and Animate-Animate`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Rect_and_Animate_Animate.X = 657.500000
	__Vertice__000003_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Rect_and_Animate_Animate.Y = 163.500000
	__Vertice__000003_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Rect_and_Animate_Animate.Name = `Verticle in class diagram Animate in middle between Animate-Rect and Animate-Animate`

	// Vertice values setup
	__Vertice__000004_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Text_and_Animate_Animate.X = 683.500000
	__Vertice__000004_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Text_and_Animate_Animate.Y = 236.500000
	__Vertice__000004_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Text_and_Animate_Animate.Name = `Verticle in class diagram Animate in middle between Animate-Text and Animate-Animate`

	// Setup of pointers
	__Classdiagram__000000_Animate.GongStructShapes = append(__Classdiagram__000000_Animate.GongStructShapes, __GongStructShape__000000_Animate_Animate)
	__Classdiagram__000000_Animate.GongStructShapes = append(__Classdiagram__000000_Animate.GongStructShapes, __GongStructShape__000004_Animate_Rect)
	__Classdiagram__000000_Animate.GongStructShapes = append(__Classdiagram__000000_Animate.GongStructShapes, __GongStructShape__000001_Animate_Circle)
	__Classdiagram__000000_Animate.GongStructShapes = append(__Classdiagram__000000_Animate.GongStructShapes, __GongStructShape__000002_Animate_Ellipse)
	__Classdiagram__000000_Animate.GongStructShapes = append(__Classdiagram__000000_Animate.GongStructShapes, __GongStructShape__000005_Animate_Text)
	__Classdiagram__000000_Animate.GongStructShapes = append(__Classdiagram__000000_Animate.GongStructShapes, __GongStructShape__000003_Animate_Line)
	__GongStructShape__000000_Animate_Animate.Position = __Position__000000_Pos_Animate_Animate
	__GongStructShape__000000_Animate_Animate.Fields = append(__GongStructShape__000000_Animate_Animate.Fields, __Field__000003_Name)
	__GongStructShape__000000_Animate_Animate.Fields = append(__GongStructShape__000000_Animate_Animate.Fields, __Field__000000_AttributeName)
	__GongStructShape__000000_Animate_Animate.Fields = append(__GongStructShape__000000_Animate_Animate.Fields, __Field__000005_Values)
	__GongStructShape__000000_Animate_Animate.Fields = append(__GongStructShape__000000_Animate_Animate.Fields, __Field__000001_Dur)
	__GongStructShape__000000_Animate_Animate.Fields = append(__GongStructShape__000000_Animate_Animate.Fields, __Field__000004_RepeatCount)
	__GongStructShape__000001_Animate_Circle.Position = __Position__000001_Pos_Animate_Circle
	__GongStructShape__000001_Animate_Circle.Links = append(__GongStructShape__000001_Animate_Circle.Links, __Link__000004_Animations)
	__GongStructShape__000002_Animate_Ellipse.Position = __Position__000002_Pos_Animate_Ellipse
	__GongStructShape__000002_Animate_Ellipse.Links = append(__GongStructShape__000002_Animate_Ellipse.Links, __Link__000000_Animates)
	__GongStructShape__000003_Animate_Line.Position = __Position__000003_Pos_Animate_Line
	__GongStructShape__000003_Animate_Line.Links = append(__GongStructShape__000003_Animate_Line.Links, __Link__000001_Animates)
	__GongStructShape__000004_Animate_Rect.Position = __Position__000004_Pos_Animate_Rect
	__GongStructShape__000004_Animate_Rect.Fields = append(__GongStructShape__000004_Animate_Rect.Fields, __Field__000002_Name)
	__GongStructShape__000004_Animate_Rect.Links = append(__GongStructShape__000004_Animate_Rect.Links, __Link__000003_Animations)
	__GongStructShape__000005_Animate_Text.Position = __Position__000005_Pos_Animate_Text
	__GongStructShape__000005_Animate_Text.Links = append(__GongStructShape__000005_Animate_Text.Links, __Link__000002_Animates)
	__Link__000000_Animates.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Ellipse_and_Animate_Animate
	__Link__000001_Animates.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Line_and_Animate_Animate
	__Link__000002_Animates.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Text_and_Animate_Animate
	__Link__000003_Animations.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Rect_and_Animate_Animate
	__Link__000004_Animations.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Animate_in_middle_between_Animate_Circle_and_Animate_Animate
}
