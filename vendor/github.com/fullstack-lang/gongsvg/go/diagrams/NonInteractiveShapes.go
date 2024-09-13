package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongsvg/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NonInteractiveShapes models.StageStruct
var ___dummy__Time_NonInteractiveShapes time.Time

// _ point for meta package dummy declaration
var ___dummy__ref_models_NonInteractiveShapes ref_models.StageStruct

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

	"ref_models.Link.HasStartArrow": (ref_models.Link{}).HasStartArrow,

	"ref_models.Link.Name": (ref_models.Link{}).Name,

	"ref_models.Link.Start": (ref_models.Link{}).Start,

	"ref_models.Link.StartAnchorType": (ref_models.Link{}).StartAnchorType,

	"ref_models.Link.StartArrowSize": (ref_models.Link{}).StartArrowSize,

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

	"ref_models.Rect.RectAnchoredPaths": (ref_models.Rect{}).RectAnchoredPaths,

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

	"ref_models.RectAnchoredPath": &(ref_models.RectAnchoredPath{}),

	"ref_models.RectAnchoredPath.Color": (ref_models.RectAnchoredPath{}).Color,

	"ref_models.RectAnchoredPath.Definition": (ref_models.RectAnchoredPath{}).Definition,

	"ref_models.RectAnchoredPath.FillOpacity": (ref_models.RectAnchoredPath{}).FillOpacity,

	"ref_models.RectAnchoredPath.Name": (ref_models.RectAnchoredPath{}).Name,

	"ref_models.RectAnchoredPath.RectAnchorType": (ref_models.RectAnchoredPath{}).RectAnchorType,

	"ref_models.RectAnchoredPath.ScalePropotionnally": (ref_models.RectAnchoredPath{}).ScalePropotionnally,

	"ref_models.RectAnchoredPath.Stroke": (ref_models.RectAnchoredPath{}).Stroke,

	"ref_models.RectAnchoredPath.StrokeDashArray": (ref_models.RectAnchoredPath{}).StrokeDashArray,

	"ref_models.RectAnchoredPath.StrokeDashArrayWhenSelected": (ref_models.RectAnchoredPath{}).StrokeDashArrayWhenSelected,

	"ref_models.RectAnchoredPath.StrokeWidth": (ref_models.RectAnchoredPath{}).StrokeWidth,

	"ref_models.RectAnchoredPath.Transform": (ref_models.RectAnchoredPath{}).Transform,

	"ref_models.RectAnchoredPath.X_Offset": (ref_models.RectAnchoredPath{}).X_Offset,

	"ref_models.RectAnchoredPath.Y_Offset": (ref_models.RectAnchoredPath{}).Y_Offset,

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
// 	InjectionGateway["NonInteractiveShapes"] = _
// }

// _ will stage objects of database "NonInteractiveShapes"
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NonInteractiveShapes := (&models.Classdiagram{Name: `NonInteractiveShapes`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NonInteractiveShapes_Circle := (&models.GongStructShape{Name: `NonInteractiveShapes-Circle`}).Stage(stage)
	__GongStructShape__000001_NonInteractiveShapes_Ellipse := (&models.GongStructShape{Name: `NonInteractiveShapes-Ellipse`}).Stage(stage)
	__GongStructShape__000002_NonInteractiveShapes_Layer := (&models.GongStructShape{Name: `NonInteractiveShapes-Layer`}).Stage(stage)
	__GongStructShape__000003_NonInteractiveShapes_Line := (&models.GongStructShape{Name: `NonInteractiveShapes-Line`}).Stage(stage)
	__GongStructShape__000004_NonInteractiveShapes_Link := (&models.GongStructShape{Name: `NonInteractiveShapes-Link`}).Stage(stage)
	__GongStructShape__000005_NonInteractiveShapes_LinkAnchoredText := (&models.GongStructShape{Name: `NonInteractiveShapes-LinkAnchoredText`}).Stage(stage)
	__GongStructShape__000006_NonInteractiveShapes_Path := (&models.GongStructShape{Name: `NonInteractiveShapes-Path`}).Stage(stage)
	__GongStructShape__000007_NonInteractiveShapes_Point := (&models.GongStructShape{Name: `NonInteractiveShapes-Point`}).Stage(stage)
	__GongStructShape__000008_NonInteractiveShapes_Polygone := (&models.GongStructShape{Name: `NonInteractiveShapes-Polygone`}).Stage(stage)
	__GongStructShape__000009_NonInteractiveShapes_Polyline := (&models.GongStructShape{Name: `NonInteractiveShapes-Polyline`}).Stage(stage)
	__GongStructShape__000010_NonInteractiveShapes_Rect := (&models.GongStructShape{Name: `NonInteractiveShapes-Rect`}).Stage(stage)
	__GongStructShape__000011_NonInteractiveShapes_RectAnchoredPath := (&models.GongStructShape{Name: `NonInteractiveShapes-RectAnchoredPath`}).Stage(stage)
	__GongStructShape__000012_NonInteractiveShapes_RectAnchoredRect := (&models.GongStructShape{Name: `NonInteractiveShapes-RectAnchoredRect`}).Stage(stage)
	__GongStructShape__000013_NonInteractiveShapes_RectAnchoredText := (&models.GongStructShape{Name: `NonInteractiveShapes-RectAnchoredText`}).Stage(stage)
	__GongStructShape__000014_NonInteractiveShapes_RectLinkLink := (&models.GongStructShape{Name: `NonInteractiveShapes-RectLinkLink`}).Stage(stage)
	__GongStructShape__000015_NonInteractiveShapes_SVG := (&models.GongStructShape{Name: `NonInteractiveShapes-SVG`}).Stage(stage)
	__GongStructShape__000016_NonInteractiveShapes_Text := (&models.GongStructShape{Name: `NonInteractiveShapes-Text`}).Stage(stage)

	// Declarations of staged instances of Link

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NonInteractiveShapes_Circle := (&models.Position{Name: `Pos-NonInteractiveShapes-Circle`}).Stage(stage)
	__Position__000001_Pos_NonInteractiveShapes_Ellipse := (&models.Position{Name: `Pos-NonInteractiveShapes-Ellipse`}).Stage(stage)
	__Position__000002_Pos_NonInteractiveShapes_Layer := (&models.Position{Name: `Pos-NonInteractiveShapes-Layer`}).Stage(stage)
	__Position__000003_Pos_NonInteractiveShapes_Line := (&models.Position{Name: `Pos-NonInteractiveShapes-Line`}).Stage(stage)
	__Position__000004_Pos_NonInteractiveShapes_Link := (&models.Position{Name: `Pos-NonInteractiveShapes-Link`}).Stage(stage)
	__Position__000005_Pos_NonInteractiveShapes_LinkAnchoredText := (&models.Position{Name: `Pos-NonInteractiveShapes-LinkAnchoredText`}).Stage(stage)
	__Position__000006_Pos_NonInteractiveShapes_Path := (&models.Position{Name: `Pos-NonInteractiveShapes-Path`}).Stage(stage)
	__Position__000007_Pos_NonInteractiveShapes_Point := (&models.Position{Name: `Pos-NonInteractiveShapes-Point`}).Stage(stage)
	__Position__000008_Pos_NonInteractiveShapes_Polygone := (&models.Position{Name: `Pos-NonInteractiveShapes-Polygone`}).Stage(stage)
	__Position__000009_Pos_NonInteractiveShapes_Polyline := (&models.Position{Name: `Pos-NonInteractiveShapes-Polyline`}).Stage(stage)
	__Position__000010_Pos_NonInteractiveShapes_Rect := (&models.Position{Name: `Pos-NonInteractiveShapes-Rect`}).Stage(stage)
	__Position__000011_Pos_NonInteractiveShapes_RectAnchoredPath := (&models.Position{Name: `Pos-NonInteractiveShapes-RectAnchoredPath`}).Stage(stage)
	__Position__000012_Pos_NonInteractiveShapes_RectAnchoredRect := (&models.Position{Name: `Pos-NonInteractiveShapes-RectAnchoredRect`}).Stage(stage)
	__Position__000013_Pos_NonInteractiveShapes_RectAnchoredText := (&models.Position{Name: `Pos-NonInteractiveShapes-RectAnchoredText`}).Stage(stage)
	__Position__000014_Pos_NonInteractiveShapes_RectLinkLink := (&models.Position{Name: `Pos-NonInteractiveShapes-RectLinkLink`}).Stage(stage)
	__Position__000015_Pos_NonInteractiveShapes_SVG := (&models.Position{Name: `Pos-NonInteractiveShapes-SVG`}).Stage(stage)
	__Position__000016_Pos_NonInteractiveShapes_Text := (&models.Position{Name: `Pos-NonInteractiveShapes-Text`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NonInteractiveShapes.Name = `NonInteractiveShapes`
	__Classdiagram__000000_NonInteractiveShapes.IsInDrawMode = true

	// GongStructShape values setup
	__GongStructShape__000000_NonInteractiveShapes_Circle.Name = `NonInteractiveShapes-Circle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle]
	__GongStructShape__000000_NonInteractiveShapes_Circle.Identifier = `ref_models.Circle`
	__GongStructShape__000000_NonInteractiveShapes_Circle.ShowNbInstances = false
	__GongStructShape__000000_NonInteractiveShapes_Circle.NbInstances = 0
	__GongStructShape__000000_NonInteractiveShapes_Circle.Width = 240.000000
	__GongStructShape__000000_NonInteractiveShapes_Circle.Height = 63.000000
	__GongStructShape__000000_NonInteractiveShapes_Circle.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NonInteractiveShapes_Ellipse.Name = `NonInteractiveShapes-Ellipse`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Ellipse]
	__GongStructShape__000001_NonInteractiveShapes_Ellipse.Identifier = `ref_models.Ellipse`
	__GongStructShape__000001_NonInteractiveShapes_Ellipse.ShowNbInstances = false
	__GongStructShape__000001_NonInteractiveShapes_Ellipse.NbInstances = 0
	__GongStructShape__000001_NonInteractiveShapes_Ellipse.Width = 240.000000
	__GongStructShape__000001_NonInteractiveShapes_Ellipse.Height = 63.000000
	__GongStructShape__000001_NonInteractiveShapes_Ellipse.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NonInteractiveShapes_Layer.Name = `NonInteractiveShapes-Layer`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer]
	__GongStructShape__000002_NonInteractiveShapes_Layer.Identifier = `ref_models.Layer`
	__GongStructShape__000002_NonInteractiveShapes_Layer.ShowNbInstances = false
	__GongStructShape__000002_NonInteractiveShapes_Layer.NbInstances = 0
	__GongStructShape__000002_NonInteractiveShapes_Layer.Width = 240.000000
	__GongStructShape__000002_NonInteractiveShapes_Layer.Height = 63.000000
	__GongStructShape__000002_NonInteractiveShapes_Layer.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NonInteractiveShapes_Line.Name = `NonInteractiveShapes-Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__GongStructShape__000003_NonInteractiveShapes_Line.Identifier = `ref_models.Line`
	__GongStructShape__000003_NonInteractiveShapes_Line.ShowNbInstances = false
	__GongStructShape__000003_NonInteractiveShapes_Line.NbInstances = 0
	__GongStructShape__000003_NonInteractiveShapes_Line.Width = 240.000000
	__GongStructShape__000003_NonInteractiveShapes_Line.Height = 63.000000
	__GongStructShape__000003_NonInteractiveShapes_Line.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_NonInteractiveShapes_Link.Name = `NonInteractiveShapes-Link`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link]
	__GongStructShape__000004_NonInteractiveShapes_Link.Identifier = `ref_models.Link`
	__GongStructShape__000004_NonInteractiveShapes_Link.ShowNbInstances = false
	__GongStructShape__000004_NonInteractiveShapes_Link.NbInstances = 0
	__GongStructShape__000004_NonInteractiveShapes_Link.Width = 240.000000
	__GongStructShape__000004_NonInteractiveShapes_Link.Height = 63.000000
	__GongStructShape__000004_NonInteractiveShapes_Link.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_NonInteractiveShapes_LinkAnchoredText.Name = `NonInteractiveShapes-LinkAnchoredText`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LinkAnchoredText]
	__GongStructShape__000005_NonInteractiveShapes_LinkAnchoredText.Identifier = `ref_models.LinkAnchoredText`
	__GongStructShape__000005_NonInteractiveShapes_LinkAnchoredText.ShowNbInstances = false
	__GongStructShape__000005_NonInteractiveShapes_LinkAnchoredText.NbInstances = 0
	__GongStructShape__000005_NonInteractiveShapes_LinkAnchoredText.Width = 240.000000
	__GongStructShape__000005_NonInteractiveShapes_LinkAnchoredText.Height = 63.000000
	__GongStructShape__000005_NonInteractiveShapes_LinkAnchoredText.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_NonInteractiveShapes_Path.Name = `NonInteractiveShapes-Path`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Path]
	__GongStructShape__000006_NonInteractiveShapes_Path.Identifier = `ref_models.Path`
	__GongStructShape__000006_NonInteractiveShapes_Path.ShowNbInstances = false
	__GongStructShape__000006_NonInteractiveShapes_Path.NbInstances = 0
	__GongStructShape__000006_NonInteractiveShapes_Path.Width = 240.000000
	__GongStructShape__000006_NonInteractiveShapes_Path.Height = 63.000000
	__GongStructShape__000006_NonInteractiveShapes_Path.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000007_NonInteractiveShapes_Point.Name = `NonInteractiveShapes-Point`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__GongStructShape__000007_NonInteractiveShapes_Point.Identifier = `ref_models.Point`
	__GongStructShape__000007_NonInteractiveShapes_Point.ShowNbInstances = false
	__GongStructShape__000007_NonInteractiveShapes_Point.NbInstances = 0
	__GongStructShape__000007_NonInteractiveShapes_Point.Width = 240.000000
	__GongStructShape__000007_NonInteractiveShapes_Point.Height = 63.000000
	__GongStructShape__000007_NonInteractiveShapes_Point.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000008_NonInteractiveShapes_Polygone.Name = `NonInteractiveShapes-Polygone`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Polygone]
	__GongStructShape__000008_NonInteractiveShapes_Polygone.Identifier = `ref_models.Polygone`
	__GongStructShape__000008_NonInteractiveShapes_Polygone.ShowNbInstances = false
	__GongStructShape__000008_NonInteractiveShapes_Polygone.NbInstances = 0
	__GongStructShape__000008_NonInteractiveShapes_Polygone.Width = 240.000000
	__GongStructShape__000008_NonInteractiveShapes_Polygone.Height = 63.000000
	__GongStructShape__000008_NonInteractiveShapes_Polygone.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000009_NonInteractiveShapes_Polyline.Name = `NonInteractiveShapes-Polyline`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Polyline]
	__GongStructShape__000009_NonInteractiveShapes_Polyline.Identifier = `ref_models.Polyline`
	__GongStructShape__000009_NonInteractiveShapes_Polyline.ShowNbInstances = false
	__GongStructShape__000009_NonInteractiveShapes_Polyline.NbInstances = 0
	__GongStructShape__000009_NonInteractiveShapes_Polyline.Width = 240.000000
	__GongStructShape__000009_NonInteractiveShapes_Polyline.Height = 63.000000
	__GongStructShape__000009_NonInteractiveShapes_Polyline.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000010_NonInteractiveShapes_Rect.Name = `NonInteractiveShapes-Rect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__GongStructShape__000010_NonInteractiveShapes_Rect.Identifier = `ref_models.Rect`
	__GongStructShape__000010_NonInteractiveShapes_Rect.ShowNbInstances = false
	__GongStructShape__000010_NonInteractiveShapes_Rect.NbInstances = 0
	__GongStructShape__000010_NonInteractiveShapes_Rect.Width = 240.000000
	__GongStructShape__000010_NonInteractiveShapes_Rect.Height = 63.000000
	__GongStructShape__000010_NonInteractiveShapes_Rect.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000011_NonInteractiveShapes_RectAnchoredPath.Name = `NonInteractiveShapes-RectAnchoredPath`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectAnchoredPath]
	__GongStructShape__000011_NonInteractiveShapes_RectAnchoredPath.Identifier = `ref_models.RectAnchoredPath`
	__GongStructShape__000011_NonInteractiveShapes_RectAnchoredPath.ShowNbInstances = false
	__GongStructShape__000011_NonInteractiveShapes_RectAnchoredPath.NbInstances = 0
	__GongStructShape__000011_NonInteractiveShapes_RectAnchoredPath.Width = 240.000000
	__GongStructShape__000011_NonInteractiveShapes_RectAnchoredPath.Height = 63.000000
	__GongStructShape__000011_NonInteractiveShapes_RectAnchoredPath.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000012_NonInteractiveShapes_RectAnchoredRect.Name = `NonInteractiveShapes-RectAnchoredRect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectAnchoredRect]
	__GongStructShape__000012_NonInteractiveShapes_RectAnchoredRect.Identifier = `ref_models.RectAnchoredRect`
	__GongStructShape__000012_NonInteractiveShapes_RectAnchoredRect.ShowNbInstances = false
	__GongStructShape__000012_NonInteractiveShapes_RectAnchoredRect.NbInstances = 0
	__GongStructShape__000012_NonInteractiveShapes_RectAnchoredRect.Width = 240.000000
	__GongStructShape__000012_NonInteractiveShapes_RectAnchoredRect.Height = 63.000000
	__GongStructShape__000012_NonInteractiveShapes_RectAnchoredRect.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000013_NonInteractiveShapes_RectAnchoredText.Name = `NonInteractiveShapes-RectAnchoredText`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectAnchoredText]
	__GongStructShape__000013_NonInteractiveShapes_RectAnchoredText.Identifier = `ref_models.RectAnchoredText`
	__GongStructShape__000013_NonInteractiveShapes_RectAnchoredText.ShowNbInstances = false
	__GongStructShape__000013_NonInteractiveShapes_RectAnchoredText.NbInstances = 0
	__GongStructShape__000013_NonInteractiveShapes_RectAnchoredText.Width = 240.000000
	__GongStructShape__000013_NonInteractiveShapes_RectAnchoredText.Height = 63.000000
	__GongStructShape__000013_NonInteractiveShapes_RectAnchoredText.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000014_NonInteractiveShapes_RectLinkLink.Name = `NonInteractiveShapes-RectLinkLink`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectLinkLink]
	__GongStructShape__000014_NonInteractiveShapes_RectLinkLink.Identifier = `ref_models.RectLinkLink`
	__GongStructShape__000014_NonInteractiveShapes_RectLinkLink.ShowNbInstances = false
	__GongStructShape__000014_NonInteractiveShapes_RectLinkLink.NbInstances = 0
	__GongStructShape__000014_NonInteractiveShapes_RectLinkLink.Width = 240.000000
	__GongStructShape__000014_NonInteractiveShapes_RectLinkLink.Height = 63.000000
	__GongStructShape__000014_NonInteractiveShapes_RectLinkLink.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000015_NonInteractiveShapes_SVG.Name = `NonInteractiveShapes-SVG`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG]
	__GongStructShape__000015_NonInteractiveShapes_SVG.Identifier = `ref_models.SVG`
	__GongStructShape__000015_NonInteractiveShapes_SVG.ShowNbInstances = false
	__GongStructShape__000015_NonInteractiveShapes_SVG.NbInstances = 0
	__GongStructShape__000015_NonInteractiveShapes_SVG.Width = 240.000000
	__GongStructShape__000015_NonInteractiveShapes_SVG.Height = 63.000000
	__GongStructShape__000015_NonInteractiveShapes_SVG.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000016_NonInteractiveShapes_Text.Name = `NonInteractiveShapes-Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text]
	__GongStructShape__000016_NonInteractiveShapes_Text.Identifier = `ref_models.Text`
	__GongStructShape__000016_NonInteractiveShapes_Text.ShowNbInstances = false
	__GongStructShape__000016_NonInteractiveShapes_Text.NbInstances = 0
	__GongStructShape__000016_NonInteractiveShapes_Text.Width = 240.000000
	__GongStructShape__000016_NonInteractiveShapes_Text.Height = 63.000000
	__GongStructShape__000016_NonInteractiveShapes_Text.IsSelected = false

	// Position values setup
	__Position__000000_Pos_NonInteractiveShapes_Circle.X = 4.000000
	__Position__000000_Pos_NonInteractiveShapes_Circle.Y = 52.000000
	__Position__000000_Pos_NonInteractiveShapes_Circle.Name = `Pos-NonInteractiveShapes-Circle`

	// Position values setup
	__Position__000001_Pos_NonInteractiveShapes_Ellipse.X = 8.000000
	__Position__000001_Pos_NonInteractiveShapes_Ellipse.Y = 211.000000
	__Position__000001_Pos_NonInteractiveShapes_Ellipse.Name = `Pos-NonInteractiveShapes-Ellipse`

	// Position values setup
	__Position__000002_Pos_NonInteractiveShapes_Layer.X = -13.000000
	__Position__000002_Pos_NonInteractiveShapes_Layer.Y = 338.000000
	__Position__000002_Pos_NonInteractiveShapes_Layer.Name = `Pos-NonInteractiveShapes-Layer`

	// Position values setup
	__Position__000003_Pos_NonInteractiveShapes_Line.X = 244.000000
	__Position__000003_Pos_NonInteractiveShapes_Line.Y = 48.000000
	__Position__000003_Pos_NonInteractiveShapes_Line.Name = `Pos-NonInteractiveShapes-Line`

	// Position values setup
	__Position__000004_Pos_NonInteractiveShapes_Link.X = 828.000000
	__Position__000004_Pos_NonInteractiveShapes_Link.Y = 51.000000
	__Position__000004_Pos_NonInteractiveShapes_Link.Name = `Pos-NonInteractiveShapes-Link`

	// Position values setup
	__Position__000005_Pos_NonInteractiveShapes_LinkAnchoredText.X = 203.000000
	__Position__000005_Pos_NonInteractiveShapes_LinkAnchoredText.Y = 532.000000
	__Position__000005_Pos_NonInteractiveShapes_LinkAnchoredText.Name = `Pos-NonInteractiveShapes-LinkAnchoredText`

	// Position values setup
	__Position__000006_Pos_NonInteractiveShapes_Path.X = 521.000000
	__Position__000006_Pos_NonInteractiveShapes_Path.Y = 513.000000
	__Position__000006_Pos_NonInteractiveShapes_Path.Name = `Pos-NonInteractiveShapes-Path`

	// Position values setup
	__Position__000007_Pos_NonInteractiveShapes_Point.X = 810.000000
	__Position__000007_Pos_NonInteractiveShapes_Point.Y = 406.000000
	__Position__000007_Pos_NonInteractiveShapes_Point.Name = `Pos-NonInteractiveShapes-Point`

	// Position values setup
	__Position__000008_Pos_NonInteractiveShapes_Polygone.X = 537.000000
	__Position__000008_Pos_NonInteractiveShapes_Polygone.Y = 57.000000
	__Position__000008_Pos_NonInteractiveShapes_Polygone.Name = `Pos-NonInteractiveShapes-Polygone`

	// Position values setup
	__Position__000009_Pos_NonInteractiveShapes_Polyline.X = 824.000000
	__Position__000009_Pos_NonInteractiveShapes_Polyline.Y = 167.000000
	__Position__000009_Pos_NonInteractiveShapes_Polyline.Name = `Pos-NonInteractiveShapes-Polyline`

	// Position values setup
	__Position__000010_Pos_NonInteractiveShapes_Rect.X = 830.000000
	__Position__000010_Pos_NonInteractiveShapes_Rect.Y = 267.000000
	__Position__000010_Pos_NonInteractiveShapes_Rect.Name = `Pos-NonInteractiveShapes-Rect`

	// Position values setup
	__Position__000011_Pos_NonInteractiveShapes_RectAnchoredPath.X = 190.000000
	__Position__000011_Pos_NonInteractiveShapes_RectAnchoredPath.Y = 154.000000
	__Position__000011_Pos_NonInteractiveShapes_RectAnchoredPath.Name = `Pos-NonInteractiveShapes-RectAnchoredPath`

	// Position values setup
	__Position__000012_Pos_NonInteractiveShapes_RectAnchoredRect.X = 238.000000
	__Position__000012_Pos_NonInteractiveShapes_RectAnchoredRect.Y = 413.000000
	__Position__000012_Pos_NonInteractiveShapes_RectAnchoredRect.Name = `Pos-NonInteractiveShapes-RectAnchoredRect`

	// Position values setup
	__Position__000013_Pos_NonInteractiveShapes_RectAnchoredText.X = 518.000000
	__Position__000013_Pos_NonInteractiveShapes_RectAnchoredText.Y = 407.000000
	__Position__000013_Pos_NonInteractiveShapes_RectAnchoredText.Name = `Pos-NonInteractiveShapes-RectAnchoredText`

	// Position values setup
	__Position__000014_Pos_NonInteractiveShapes_RectLinkLink.X = 504.000000
	__Position__000014_Pos_NonInteractiveShapes_RectLinkLink.Y = 166.000000
	__Position__000014_Pos_NonInteractiveShapes_RectLinkLink.Name = `Pos-NonInteractiveShapes-RectLinkLink`

	// Position values setup
	__Position__000015_Pos_NonInteractiveShapes_SVG.X = 517.000000
	__Position__000015_Pos_NonInteractiveShapes_SVG.Y = 271.000000
	__Position__000015_Pos_NonInteractiveShapes_SVG.Name = `Pos-NonInteractiveShapes-SVG`

	// Position values setup
	__Position__000016_Pos_NonInteractiveShapes_Text.X = 190.000000
	__Position__000016_Pos_NonInteractiveShapes_Text.Y = 266.000000
	__Position__000016_Pos_NonInteractiveShapes_Text.Name = `Pos-NonInteractiveShapes-Text`

	// Setup of pointers
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000000_NonInteractiveShapes_Circle)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000001_NonInteractiveShapes_Ellipse)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000002_NonInteractiveShapes_Layer)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000003_NonInteractiveShapes_Line)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000004_NonInteractiveShapes_Link)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000005_NonInteractiveShapes_LinkAnchoredText)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000006_NonInteractiveShapes_Path)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000007_NonInteractiveShapes_Point)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000008_NonInteractiveShapes_Polygone)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000009_NonInteractiveShapes_Polyline)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000010_NonInteractiveShapes_Rect)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000011_NonInteractiveShapes_RectAnchoredPath)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000012_NonInteractiveShapes_RectAnchoredRect)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000013_NonInteractiveShapes_RectAnchoredText)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000014_NonInteractiveShapes_RectLinkLink)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000015_NonInteractiveShapes_SVG)
	__Classdiagram__000000_NonInteractiveShapes.GongStructShapes = append(__Classdiagram__000000_NonInteractiveShapes.GongStructShapes, __GongStructShape__000016_NonInteractiveShapes_Text)
	__GongStructShape__000000_NonInteractiveShapes_Circle.Position = __Position__000000_Pos_NonInteractiveShapes_Circle
	__GongStructShape__000001_NonInteractiveShapes_Ellipse.Position = __Position__000001_Pos_NonInteractiveShapes_Ellipse
	__GongStructShape__000002_NonInteractiveShapes_Layer.Position = __Position__000002_Pos_NonInteractiveShapes_Layer
	__GongStructShape__000003_NonInteractiveShapes_Line.Position = __Position__000003_Pos_NonInteractiveShapes_Line
	__GongStructShape__000004_NonInteractiveShapes_Link.Position = __Position__000004_Pos_NonInteractiveShapes_Link
	__GongStructShape__000005_NonInteractiveShapes_LinkAnchoredText.Position = __Position__000005_Pos_NonInteractiveShapes_LinkAnchoredText
	__GongStructShape__000006_NonInteractiveShapes_Path.Position = __Position__000006_Pos_NonInteractiveShapes_Path
	__GongStructShape__000007_NonInteractiveShapes_Point.Position = __Position__000007_Pos_NonInteractiveShapes_Point
	__GongStructShape__000008_NonInteractiveShapes_Polygone.Position = __Position__000008_Pos_NonInteractiveShapes_Polygone
	__GongStructShape__000009_NonInteractiveShapes_Polyline.Position = __Position__000009_Pos_NonInteractiveShapes_Polyline
	__GongStructShape__000010_NonInteractiveShapes_Rect.Position = __Position__000010_Pos_NonInteractiveShapes_Rect
	__GongStructShape__000011_NonInteractiveShapes_RectAnchoredPath.Position = __Position__000011_Pos_NonInteractiveShapes_RectAnchoredPath
	__GongStructShape__000012_NonInteractiveShapes_RectAnchoredRect.Position = __Position__000012_Pos_NonInteractiveShapes_RectAnchoredRect
	__GongStructShape__000013_NonInteractiveShapes_RectAnchoredText.Position = __Position__000013_Pos_NonInteractiveShapes_RectAnchoredText
	__GongStructShape__000014_NonInteractiveShapes_RectLinkLink.Position = __Position__000014_Pos_NonInteractiveShapes_RectLinkLink
	__GongStructShape__000015_NonInteractiveShapes_SVG.Position = __Position__000015_Pos_NonInteractiveShapes_SVG
	__GongStructShape__000016_NonInteractiveShapes_Text.Position = __Position__000016_Pos_NonInteractiveShapes_Text
}
