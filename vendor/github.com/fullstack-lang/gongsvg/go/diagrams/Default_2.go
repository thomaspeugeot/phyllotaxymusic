package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongsvg/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Default_2 models.StageStruct
var ___dummy__Time_Default_2 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Default_2 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Default_2 map[string]any = map[string]any{
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
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default_2"] = Default_2Injection
// }

// Default_2Injection will stage objects of database "Default_2"
func Default_2Injection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default_2 := (&models.Classdiagram{Name: `Default_2`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape

	// Declarations of staged instances of Link

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default_2.Name = `Default_2`
	__Classdiagram__000000_Default_2.IsInDrawMode = false

	// Setup of pointers
}
