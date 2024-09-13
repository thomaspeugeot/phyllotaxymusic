package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongsvg/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram models.StageStruct
var ___dummy__Time_NewDiagram time.Time

// _ point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram ref_models.StageStruct

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
// 	InjectionGateway["NewDiagram"] = _
// }

// _ will stage objects of database "NewDiagram"
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CanHaveBottomHandle := (&models.Field{Name: `CanHaveBottomHandle`}).Stage(stage)
	__Field__000001_CanHaveLeftHandle := (&models.Field{Name: `CanHaveLeftHandle`}).Stage(stage)
	__Field__000002_CanHaveRightHandle := (&models.Field{Name: `CanHaveRightHandle`}).Stage(stage)
	__Field__000003_CanHaveTopHandle := (&models.Field{Name: `CanHaveTopHandle`}).Stage(stage)
	__Field__000004_CanMoveHorizontaly := (&models.Field{Name: `CanMoveHorizontaly`}).Stage(stage)
	__Field__000005_CanMoveVerticaly := (&models.Field{Name: `CanMoveVerticaly`}).Stage(stage)
	__Field__000006_Color := (&models.Field{Name: `Color`}).Stage(stage)
	__Field__000007_Content := (&models.Field{Name: `Content`}).Stage(stage)
	__Field__000008_CornerOffsetRatio := (&models.Field{Name: `CornerOffsetRatio`}).Stage(stage)
	__Field__000009_CornerRadius := (&models.Field{Name: `CornerRadius`}).Stage(stage)
	__Field__000010_Display := (&models.Field{Name: `Display`}).Stage(stage)
	__Field__000011_EndAnchorType := (&models.Field{Name: `EndAnchorType`}).Stage(stage)
	__Field__000012_EndArrowSize := (&models.Field{Name: `EndArrowSize`}).Stage(stage)
	__Field__000013_EndOrientation := (&models.Field{Name: `EndOrientation`}).Stage(stage)
	__Field__000014_EndRatio := (&models.Field{Name: `EndRatio`}).Stage(stage)
	__Field__000015_HasBottomHandle := (&models.Field{Name: `HasBottomHandle`}).Stage(stage)
	__Field__000016_HasLeftHandle := (&models.Field{Name: `HasLeftHandle`}).Stage(stage)
	__Field__000017_HasRightHandle := (&models.Field{Name: `HasRightHandle`}).Stage(stage)
	__Field__000018_HasTopHandle := (&models.Field{Name: `HasTopHandle`}).Stage(stage)
	__Field__000019_Height := (&models.Field{Name: `Height`}).Stage(stage)
	__Field__000020_IsSelectable := (&models.Field{Name: `IsSelectable`}).Stage(stage)
	__Field__000021_IsSelected := (&models.Field{Name: `IsSelected`}).Stage(stage)
	__Field__000022_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000023_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000024_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000025_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000026_RX := (&models.Field{Name: `RX`}).Stage(stage)
	__Field__000027_StartAnchorType := (&models.Field{Name: `StartAnchorType`}).Stage(stage)
	__Field__000028_StartOrientation := (&models.Field{Name: `StartOrientation`}).Stage(stage)
	__Field__000029_StartRatio := (&models.Field{Name: `StartRatio`}).Stage(stage)
	__Field__000030_Stroke := (&models.Field{Name: `Stroke`}).Stage(stage)
	__Field__000031_TargetAnchorPosition := (&models.Field{Name: `TargetAnchorPosition`}).Stage(stage)
	__Field__000032_Width := (&models.Field{Name: `Width`}).Stage(stage)
	__Field__000033_X := (&models.Field{Name: `X`}).Stage(stage)
	__Field__000034_X_Offset := (&models.Field{Name: `X_Offset`}).Stage(stage)
	__Field__000035_Y := (&models.Field{Name: `Y`}).Stage(stage)
	__Field__000036_Y_Offset := (&models.Field{Name: `Y_Offset`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_AnchoredText := (&models.GongStructShape{Name: `NewDiagram-AnchoredText`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_Layer := (&models.GongStructShape{Name: `NewDiagram-Layer`}).Stage(stage)
	__GongStructShape__000002_NewDiagram_Link := (&models.GongStructShape{Name: `NewDiagram-Link`}).Stage(stage)
	__GongStructShape__000003_NewDiagram_Rect := (&models.GongStructShape{Name: `NewDiagram-Rect`}).Stage(stage)
	__GongStructShape__000004_NewDiagram_RectAnchoredRect := (&models.GongStructShape{Name: `NewDiagram-RectAnchoredRect`}).Stage(stage)
	__GongStructShape__000005_NewDiagram_RectAnchoredText := (&models.GongStructShape{Name: `NewDiagram-RectAnchoredText`}).Stage(stage)
	__GongStructShape__000006_NewDiagram_RectLinkLink := (&models.GongStructShape{Name: `NewDiagram-RectLinkLink`}).Stage(stage)
	__GongStructShape__000007_NewDiagram_SVG := (&models.GongStructShape{Name: `NewDiagram-SVG`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_End := (&models.Link{Name: `End`}).Stage(stage)
	__Link__000001_End := (&models.Link{Name: `End`}).Stage(stage)
	__Link__000002_Layers := (&models.Link{Name: `Layers`}).Stage(stage)
	__Link__000003_Links := (&models.Link{Name: `Links`}).Stage(stage)
	__Link__000004_RectAnchoredRects := (&models.Link{Name: `RectAnchoredRects`}).Stage(stage)
	__Link__000005_RectAnchoredTexts := (&models.Link{Name: `RectAnchoredTexts`}).Stage(stage)
	__Link__000006_RectLinkLinks := (&models.Link{Name: `RectLinkLinks`}).Stage(stage)
	__Link__000007_Rects := (&models.Link{Name: `Rects`}).Stage(stage)
	__Link__000008_Start := (&models.Link{Name: `Start`}).Stage(stage)
	__Link__000009_Start := (&models.Link{Name: `Start`}).Stage(stage)
	__Link__000010_TextAtArrowEnd := (&models.Link{Name: `TextAtArrowEnd`}).Stage(stage)
	__Link__000011_TextAtArrowStart := (&models.Link{Name: `TextAtArrowStart`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_AnchoredText := (&models.Position{Name: `Pos-NewDiagram-AnchoredText`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_Layer := (&models.Position{Name: `Pos-NewDiagram-Layer`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_Link := (&models.Position{Name: `Pos-NewDiagram-Link`}).Stage(stage)
	__Position__000003_Pos_NewDiagram_Rect := (&models.Position{Name: `Pos-NewDiagram-Rect`}).Stage(stage)
	__Position__000004_Pos_NewDiagram_RectAnchoredRect := (&models.Position{Name: `Pos-NewDiagram-RectAnchoredRect`}).Stage(stage)
	__Position__000005_Pos_NewDiagram_RectAnchoredText := (&models.Position{Name: `Pos-NewDiagram-RectAnchoredText`}).Stage(stage)
	__Position__000006_Pos_NewDiagram_RectLinkLink := (&models.Position{Name: `Pos-NewDiagram-RectLinkLink`}).Stage(stage)
	__Position__000007_Pos_NewDiagram_SVG := (&models.Position{Name: `Pos-NewDiagram-SVG`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Link := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Link`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Rect`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_RectLinkLink := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-RectLinkLink`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_AnchoredText := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Link and NewDiagram-AnchoredText`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_AnchoredText := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Link and NewDiagram-AnchoredText`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Link and NewDiagram-Rect`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Link and NewDiagram-Rect`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredRect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Rect and NewDiagram-RectAnchoredRect`}).Stage(stage)
	__Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredText := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Rect and NewDiagram-RectAnchoredText`}).Stage(stage)
	__Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_RectLinkLink_and_NewDiagram_Link := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-RectLinkLink and NewDiagram-Link`}).Stage(stage)
	__Vertice__000010_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_RectLinkLink_and_NewDiagram_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-RectLinkLink and NewDiagram-Rect`}).Stage(stage)
	__Vertice__000011_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Layer := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Layer`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_CanHaveBottomHandle.Name = `CanHaveBottomHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.CanHaveBottomHandle]
	__Field__000000_CanHaveBottomHandle.Identifier = `ref_models.Rect.CanHaveBottomHandle`
	__Field__000000_CanHaveBottomHandle.FieldTypeAsString = ``
	__Field__000000_CanHaveBottomHandle.Structname = `Rect`
	__Field__000000_CanHaveBottomHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000001_CanHaveLeftHandle.Name = `CanHaveLeftHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.CanHaveLeftHandle]
	__Field__000001_CanHaveLeftHandle.Identifier = `ref_models.Rect.CanHaveLeftHandle`
	__Field__000001_CanHaveLeftHandle.FieldTypeAsString = ``
	__Field__000001_CanHaveLeftHandle.Structname = `Rect`
	__Field__000001_CanHaveLeftHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000002_CanHaveRightHandle.Name = `CanHaveRightHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.CanHaveRightHandle]
	__Field__000002_CanHaveRightHandle.Identifier = `ref_models.Rect.CanHaveRightHandle`
	__Field__000002_CanHaveRightHandle.FieldTypeAsString = ``
	__Field__000002_CanHaveRightHandle.Structname = `Rect`
	__Field__000002_CanHaveRightHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000003_CanHaveTopHandle.Name = `CanHaveTopHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.CanHaveTopHandle]
	__Field__000003_CanHaveTopHandle.Identifier = `ref_models.Rect.CanHaveTopHandle`
	__Field__000003_CanHaveTopHandle.FieldTypeAsString = ``
	__Field__000003_CanHaveTopHandle.Structname = `Rect`
	__Field__000003_CanHaveTopHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000004_CanMoveHorizontaly.Name = `CanMoveHorizontaly`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.CanMoveHorizontaly]
	__Field__000004_CanMoveHorizontaly.Identifier = `ref_models.Rect.CanMoveHorizontaly`
	__Field__000004_CanMoveHorizontaly.FieldTypeAsString = ``
	__Field__000004_CanMoveHorizontaly.Structname = `Rect`
	__Field__000004_CanMoveHorizontaly.Fieldtypename = `bool`

	// Field values setup
	__Field__000005_CanMoveVerticaly.Name = `CanMoveVerticaly`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.CanMoveVerticaly]
	__Field__000005_CanMoveVerticaly.Identifier = `ref_models.Rect.CanMoveVerticaly`
	__Field__000005_CanMoveVerticaly.FieldTypeAsString = ``
	__Field__000005_CanMoveVerticaly.Structname = `Rect`
	__Field__000005_CanMoveVerticaly.Fieldtypename = `bool`

	// Field values setup
	__Field__000006_Color.Name = `Color`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LinkAnchoredText.Color]
	__Field__000006_Color.Identifier = `ref_models.LinkAnchoredText.Color`
	__Field__000006_Color.FieldTypeAsString = ``
	__Field__000006_Color.Structname = `AnchoredText`
	__Field__000006_Color.Fieldtypename = `string`

	// Field values setup
	__Field__000007_Content.Name = `Content`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LinkAnchoredText.Content]
	__Field__000007_Content.Identifier = `ref_models.LinkAnchoredText.Content`
	__Field__000007_Content.FieldTypeAsString = ``
	__Field__000007_Content.Structname = `AnchoredText`
	__Field__000007_Content.Fieldtypename = `string`

	// Field values setup
	__Field__000008_CornerOffsetRatio.Name = `CornerOffsetRatio`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.CornerOffsetRatio]
	__Field__000008_CornerOffsetRatio.Identifier = `ref_models.Link.CornerOffsetRatio`
	__Field__000008_CornerOffsetRatio.FieldTypeAsString = ``
	__Field__000008_CornerOffsetRatio.Structname = `Link`
	__Field__000008_CornerOffsetRatio.Fieldtypename = `float64`

	// Field values setup
	__Field__000009_CornerRadius.Name = `CornerRadius`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.CornerRadius]
	__Field__000009_CornerRadius.Identifier = `ref_models.Link.CornerRadius`
	__Field__000009_CornerRadius.FieldTypeAsString = ``
	__Field__000009_CornerRadius.Structname = `Link`
	__Field__000009_CornerRadius.Fieldtypename = `float64`

	// Field values setup
	__Field__000010_Display.Name = `Display`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Display]
	__Field__000010_Display.Identifier = `ref_models.Layer.Display`
	__Field__000010_Display.FieldTypeAsString = ``
	__Field__000010_Display.Structname = `Layer`
	__Field__000010_Display.Fieldtypename = `bool`

	// Field values setup
	__Field__000011_EndAnchorType.Name = `EndAnchorType`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.EndAnchorType]
	__Field__000011_EndAnchorType.Identifier = `ref_models.Link.EndAnchorType`
	__Field__000011_EndAnchorType.FieldTypeAsString = ``
	__Field__000011_EndAnchorType.Structname = `Link`
	__Field__000011_EndAnchorType.Fieldtypename = `AnchorType`

	// Field values setup
	__Field__000012_EndArrowSize.Name = `EndArrowSize`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.EndArrowSize]
	__Field__000012_EndArrowSize.Identifier = `ref_models.Link.EndArrowSize`
	__Field__000012_EndArrowSize.FieldTypeAsString = ``
	__Field__000012_EndArrowSize.Structname = `Link`
	__Field__000012_EndArrowSize.Fieldtypename = `float64`

	// Field values setup
	__Field__000013_EndOrientation.Name = `EndOrientation`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.EndOrientation]
	__Field__000013_EndOrientation.Identifier = `ref_models.Link.EndOrientation`
	__Field__000013_EndOrientation.FieldTypeAsString = ``
	__Field__000013_EndOrientation.Structname = `Link`
	__Field__000013_EndOrientation.Fieldtypename = `OrientationType`

	// Field values setup
	__Field__000014_EndRatio.Name = `EndRatio`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.EndRatio]
	__Field__000014_EndRatio.Identifier = `ref_models.Link.EndRatio`
	__Field__000014_EndRatio.FieldTypeAsString = ``
	__Field__000014_EndRatio.Structname = `Link`
	__Field__000014_EndRatio.Fieldtypename = `float64`

	// Field values setup
	__Field__000015_HasBottomHandle.Name = `HasBottomHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.HasBottomHandle]
	__Field__000015_HasBottomHandle.Identifier = `ref_models.Rect.HasBottomHandle`
	__Field__000015_HasBottomHandle.FieldTypeAsString = ``
	__Field__000015_HasBottomHandle.Structname = `Rect`
	__Field__000015_HasBottomHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000016_HasLeftHandle.Name = `HasLeftHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.HasLeftHandle]
	__Field__000016_HasLeftHandle.Identifier = `ref_models.Rect.HasLeftHandle`
	__Field__000016_HasLeftHandle.FieldTypeAsString = ``
	__Field__000016_HasLeftHandle.Structname = `Rect`
	__Field__000016_HasLeftHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000017_HasRightHandle.Name = `HasRightHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.HasRightHandle]
	__Field__000017_HasRightHandle.Identifier = `ref_models.Rect.HasRightHandle`
	__Field__000017_HasRightHandle.FieldTypeAsString = ``
	__Field__000017_HasRightHandle.Structname = `Rect`
	__Field__000017_HasRightHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000018_HasTopHandle.Name = `HasTopHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.HasTopHandle]
	__Field__000018_HasTopHandle.Identifier = `ref_models.Rect.HasTopHandle`
	__Field__000018_HasTopHandle.FieldTypeAsString = ``
	__Field__000018_HasTopHandle.Structname = `Rect`
	__Field__000018_HasTopHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000019_Height.Name = `Height`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.Height]
	__Field__000019_Height.Identifier = `ref_models.Rect.Height`
	__Field__000019_Height.FieldTypeAsString = ``
	__Field__000019_Height.Structname = `Rect`
	__Field__000019_Height.Fieldtypename = `float64`

	// Field values setup
	__Field__000020_IsSelectable.Name = `IsSelectable`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.IsSelectable]
	__Field__000020_IsSelectable.Identifier = `ref_models.Rect.IsSelectable`
	__Field__000020_IsSelectable.FieldTypeAsString = ``
	__Field__000020_IsSelectable.Structname = `Rect`
	__Field__000020_IsSelectable.Fieldtypename = `bool`

	// Field values setup
	__Field__000021_IsSelected.Name = `IsSelected`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.IsSelected]
	__Field__000021_IsSelected.Identifier = `ref_models.Rect.IsSelected`
	__Field__000021_IsSelected.FieldTypeAsString = ``
	__Field__000021_IsSelected.Structname = `Rect`
	__Field__000021_IsSelected.Fieldtypename = `bool`

	// Field values setup
	__Field__000022_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG.Name]
	__Field__000022_Name.Identifier = `ref_models.SVG.Name`
	__Field__000022_Name.FieldTypeAsString = ``
	__Field__000022_Name.Structname = `SVG`
	__Field__000022_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000023_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LinkAnchoredText.Name]
	__Field__000023_Name.Identifier = `ref_models.LinkAnchoredText.Name`
	__Field__000023_Name.FieldTypeAsString = ``
	__Field__000023_Name.Structname = `AnchoredText`
	__Field__000023_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000024_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Name]
	__Field__000024_Name.Identifier = `ref_models.Layer.Name`
	__Field__000024_Name.FieldTypeAsString = ``
	__Field__000024_Name.Structname = `Layer`
	__Field__000024_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000025_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.Name]
	__Field__000025_Name.Identifier = `ref_models.Rect.Name`
	__Field__000025_Name.FieldTypeAsString = ``
	__Field__000025_Name.Structname = `Rect`
	__Field__000025_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000026_RX.Name = `RX`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.RX]
	__Field__000026_RX.Identifier = `ref_models.Rect.RX`
	__Field__000026_RX.FieldTypeAsString = ``
	__Field__000026_RX.Structname = `Rect`
	__Field__000026_RX.Fieldtypename = `float64`

	// Field values setup
	__Field__000027_StartAnchorType.Name = `StartAnchorType`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.StartAnchorType]
	__Field__000027_StartAnchorType.Identifier = `ref_models.Link.StartAnchorType`
	__Field__000027_StartAnchorType.FieldTypeAsString = ``
	__Field__000027_StartAnchorType.Structname = `Link`
	__Field__000027_StartAnchorType.Fieldtypename = `AnchorType`

	// Field values setup
	__Field__000028_StartOrientation.Name = `StartOrientation`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.StartOrientation]
	__Field__000028_StartOrientation.Identifier = `ref_models.Link.StartOrientation`
	__Field__000028_StartOrientation.FieldTypeAsString = ``
	__Field__000028_StartOrientation.Structname = `Link`
	__Field__000028_StartOrientation.Fieldtypename = `OrientationType`

	// Field values setup
	__Field__000029_StartRatio.Name = `StartRatio`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.StartRatio]
	__Field__000029_StartRatio.Identifier = `ref_models.Link.StartRatio`
	__Field__000029_StartRatio.FieldTypeAsString = ``
	__Field__000029_StartRatio.Structname = `Link`
	__Field__000029_StartRatio.Fieldtypename = `float64`

	// Field values setup
	__Field__000030_Stroke.Name = `Stroke`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LinkAnchoredText.Stroke]
	__Field__000030_Stroke.Identifier = `ref_models.LinkAnchoredText.Stroke`
	__Field__000030_Stroke.FieldTypeAsString = ``
	__Field__000030_Stroke.Structname = `AnchoredText`
	__Field__000030_Stroke.Fieldtypename = `string`

	// Field values setup
	__Field__000031_TargetAnchorPosition.Name = `TargetAnchorPosition`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectLinkLink.TargetAnchorPosition]
	__Field__000031_TargetAnchorPosition.Identifier = `ref_models.RectLinkLink.TargetAnchorPosition`
	__Field__000031_TargetAnchorPosition.FieldTypeAsString = ``
	__Field__000031_TargetAnchorPosition.Structname = `RectLinkLink`
	__Field__000031_TargetAnchorPosition.Fieldtypename = `float64`

	// Field values setup
	__Field__000032_Width.Name = `Width`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.Width]
	__Field__000032_Width.Identifier = `ref_models.Rect.Width`
	__Field__000032_Width.FieldTypeAsString = ``
	__Field__000032_Width.Structname = `Rect`
	__Field__000032_Width.Fieldtypename = `float64`

	// Field values setup
	__Field__000033_X.Name = `X`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.X]
	__Field__000033_X.Identifier = `ref_models.Rect.X`
	__Field__000033_X.FieldTypeAsString = ``
	__Field__000033_X.Structname = `Rect`
	__Field__000033_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000034_X_Offset.Name = `X_Offset`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LinkAnchoredText.X_Offset]
	__Field__000034_X_Offset.Identifier = `ref_models.LinkAnchoredText.X_Offset`
	__Field__000034_X_Offset.FieldTypeAsString = ``
	__Field__000034_X_Offset.Structname = `AnchoredText`
	__Field__000034_X_Offset.Fieldtypename = `float64`

	// Field values setup
	__Field__000035_Y.Name = `Y`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.Y]
	__Field__000035_Y.Identifier = `ref_models.Rect.Y`
	__Field__000035_Y.FieldTypeAsString = ``
	__Field__000035_Y.Structname = `Rect`
	__Field__000035_Y.Fieldtypename = `float64`

	// Field values setup
	__Field__000036_Y_Offset.Name = `Y_Offset`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LinkAnchoredText.Y_Offset]
	__Field__000036_Y_Offset.Identifier = `ref_models.LinkAnchoredText.Y_Offset`
	__Field__000036_Y_Offset.FieldTypeAsString = ``
	__Field__000036_Y_Offset.Structname = `AnchoredText`
	__Field__000036_Y_Offset.Fieldtypename = `float64`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_AnchoredText.Name = `NewDiagram-AnchoredText`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LinkAnchoredText]
	__GongStructShape__000000_NewDiagram_AnchoredText.Identifier = `ref_models.LinkAnchoredText`
	__GongStructShape__000000_NewDiagram_AnchoredText.ShowNbInstances = true
	__GongStructShape__000000_NewDiagram_AnchoredText.NbInstances = 0
	__GongStructShape__000000_NewDiagram_AnchoredText.Width = 240.000000
	__GongStructShape__000000_NewDiagram_AnchoredText.Height = 153.000000
	__GongStructShape__000000_NewDiagram_AnchoredText.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Layer.Name = `NewDiagram-Layer`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer]
	__GongStructShape__000001_NewDiagram_Layer.Identifier = `ref_models.Layer`
	__GongStructShape__000001_NewDiagram_Layer.ShowNbInstances = true
	__GongStructShape__000001_NewDiagram_Layer.NbInstances = 1
	__GongStructShape__000001_NewDiagram_Layer.Width = 138.000000
	__GongStructShape__000001_NewDiagram_Layer.Height = 605.000000
	__GongStructShape__000001_NewDiagram_Layer.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_Link.Name = `NewDiagram-Link`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link]
	__GongStructShape__000002_NewDiagram_Link.Identifier = `ref_models.Link`
	__GongStructShape__000002_NewDiagram_Link.ShowNbInstances = true
	__GongStructShape__000002_NewDiagram_Link.NbInstances = 0
	__GongStructShape__000002_NewDiagram_Link.Width = 330.000000
	__GongStructShape__000002_NewDiagram_Link.Height = 198.000000
	__GongStructShape__000002_NewDiagram_Link.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_Rect.Name = `NewDiagram-Rect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__GongStructShape__000003_NewDiagram_Rect.Identifier = `ref_models.Rect`
	__GongStructShape__000003_NewDiagram_Rect.ShowNbInstances = true
	__GongStructShape__000003_NewDiagram_Rect.NbInstances = 5
	__GongStructShape__000003_NewDiagram_Rect.Width = 316.000000
	__GongStructShape__000003_NewDiagram_Rect.Height = 330.000000
	__GongStructShape__000003_NewDiagram_Rect.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_NewDiagram_RectAnchoredRect.Name = `NewDiagram-RectAnchoredRect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectAnchoredRect]
	__GongStructShape__000004_NewDiagram_RectAnchoredRect.Identifier = `ref_models.RectAnchoredRect`
	__GongStructShape__000004_NewDiagram_RectAnchoredRect.ShowNbInstances = true
	__GongStructShape__000004_NewDiagram_RectAnchoredRect.NbInstances = 0
	__GongStructShape__000004_NewDiagram_RectAnchoredRect.Width = 240.000000
	__GongStructShape__000004_NewDiagram_RectAnchoredRect.Height = 63.000000
	__GongStructShape__000004_NewDiagram_RectAnchoredRect.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_NewDiagram_RectAnchoredText.Name = `NewDiagram-RectAnchoredText`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectAnchoredText]
	__GongStructShape__000005_NewDiagram_RectAnchoredText.Identifier = `ref_models.RectAnchoredText`
	__GongStructShape__000005_NewDiagram_RectAnchoredText.ShowNbInstances = true
	__GongStructShape__000005_NewDiagram_RectAnchoredText.NbInstances = 0
	__GongStructShape__000005_NewDiagram_RectAnchoredText.Width = 240.000000
	__GongStructShape__000005_NewDiagram_RectAnchoredText.Height = 63.000000
	__GongStructShape__000005_NewDiagram_RectAnchoredText.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_NewDiagram_RectLinkLink.Name = `NewDiagram-RectLinkLink`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectLinkLink]
	__GongStructShape__000006_NewDiagram_RectLinkLink.Identifier = `ref_models.RectLinkLink`
	__GongStructShape__000006_NewDiagram_RectLinkLink.ShowNbInstances = true
	__GongStructShape__000006_NewDiagram_RectLinkLink.NbInstances = 0
	__GongStructShape__000006_NewDiagram_RectLinkLink.Width = 240.000000
	__GongStructShape__000006_NewDiagram_RectLinkLink.Height = 78.000000
	__GongStructShape__000006_NewDiagram_RectLinkLink.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000007_NewDiagram_SVG.Name = `NewDiagram-SVG`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG]
	__GongStructShape__000007_NewDiagram_SVG.Identifier = `ref_models.SVG`
	__GongStructShape__000007_NewDiagram_SVG.ShowNbInstances = true
	__GongStructShape__000007_NewDiagram_SVG.NbInstances = 1
	__GongStructShape__000007_NewDiagram_SVG.Width = 139.000000
	__GongStructShape__000007_NewDiagram_SVG.Height = 63.000000
	__GongStructShape__000007_NewDiagram_SVG.IsSelected = false

	// Link values setup
	__Link__000000_End.Name = `End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.End]
	__Link__000000_End.Identifier = `ref_models.Link.End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000000_End.Fieldtypename = `ref_models.Rect`
	__Link__000000_End.FieldOffsetX = -37.000000
	__Link__000000_End.FieldOffsetY = 23.000000
	__Link__000000_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_End.TargetMultiplicityOffsetX = 14.000000
	__Link__000000_End.TargetMultiplicityOffsetY = 21.000000
	__Link__000000_End.SourceMultiplicity = models.MANY
	__Link__000000_End.SourceMultiplicityOffsetX = 11.000000
	__Link__000000_End.SourceMultiplicityOffsetY = -12.000000
	__Link__000000_End.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_End.StartRatio = 0.743838
	__Link__000000_End.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_End.EndRatio = 0.766550
	__Link__000000_End.CornerOffsetRatio = -0.434343

	// Link values setup
	__Link__000001_End.Name = `End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectLinkLink.End]
	__Link__000001_End.Identifier = `ref_models.RectLinkLink.End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link]
	__Link__000001_End.Fieldtypename = `ref_models.Link`
	__Link__000001_End.FieldOffsetX = -50.000000
	__Link__000001_End.FieldOffsetY = -16.000000
	__Link__000001_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_End.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_End.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_End.SourceMultiplicity = models.MANY
	__Link__000001_End.SourceMultiplicityOffsetX = -26.000000
	__Link__000001_End.SourceMultiplicityOffsetY = 22.000000
	__Link__000001_End.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_End.StartRatio = 0.718611
	__Link__000001_End.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_End.EndRatio = 0.227273
	__Link__000001_End.CornerOffsetRatio = 1.015873

	// Link values setup
	__Link__000002_Layers.Name = `Layers`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG.Layers]
	__Link__000002_Layers.Identifier = `ref_models.SVG.Layers`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer]
	__Link__000002_Layers.Fieldtypename = `ref_models.Layer`
	__Link__000002_Layers.FieldOffsetX = 8.000000
	__Link__000002_Layers.FieldOffsetY = -12.000000
	__Link__000002_Layers.TargetMultiplicity = models.MANY
	__Link__000002_Layers.TargetMultiplicityOffsetX = -26.000000
	__Link__000002_Layers.TargetMultiplicityOffsetY = -2.000000
	__Link__000002_Layers.SourceMultiplicity = models.ZERO_ONE
	__Link__000002_Layers.SourceMultiplicityOffsetX = -33.000000
	__Link__000002_Layers.SourceMultiplicityOffsetY = 17.000000
	__Link__000002_Layers.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_Layers.StartRatio = 0.379167
	__Link__000002_Layers.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_Layers.EndRatio = 0.416425
	__Link__000002_Layers.CornerOffsetRatio = 1.698413

	// Link values setup
	__Link__000003_Links.Name = `Links`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Links]
	__Link__000003_Links.Identifier = `ref_models.Layer.Links`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link]
	__Link__000003_Links.Fieldtypename = `ref_models.Link`
	__Link__000003_Links.FieldOffsetX = -52.000000
	__Link__000003_Links.FieldOffsetY = -17.000000
	__Link__000003_Links.TargetMultiplicity = models.MANY
	__Link__000003_Links.TargetMultiplicityOffsetX = -17.000000
	__Link__000003_Links.TargetMultiplicityOffsetY = 24.000000
	__Link__000003_Links.SourceMultiplicity = models.ZERO_ONE
	__Link__000003_Links.SourceMultiplicityOffsetX = 14.000000
	__Link__000003_Links.SourceMultiplicityOffsetY = 25.000000
	__Link__000003_Links.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_Links.StartRatio = 0.927273
	__Link__000003_Links.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_Links.EndRatio = 0.742424
	__Link__000003_Links.CornerOffsetRatio = 1.431111

	// Link values setup
	__Link__000004_RectAnchoredRects.Name = `RectAnchoredRects`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.RectAnchoredRects]
	__Link__000004_RectAnchoredRects.Identifier = `ref_models.Rect.RectAnchoredRects`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectAnchoredRect]
	__Link__000004_RectAnchoredRects.Fieldtypename = `ref_models.RectAnchoredRect`
	__Link__000004_RectAnchoredRects.FieldOffsetX = -135.000000
	__Link__000004_RectAnchoredRects.FieldOffsetY = -16.000000
	__Link__000004_RectAnchoredRects.TargetMultiplicity = models.MANY
	__Link__000004_RectAnchoredRects.TargetMultiplicityOffsetX = -30.000000
	__Link__000004_RectAnchoredRects.TargetMultiplicityOffsetY = 30.000000
	__Link__000004_RectAnchoredRects.SourceMultiplicity = models.ZERO_ONE
	__Link__000004_RectAnchoredRects.SourceMultiplicityOffsetX = 16.000000
	__Link__000004_RectAnchoredRects.SourceMultiplicityOffsetY = 26.000000
	__Link__000004_RectAnchoredRects.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_RectAnchoredRects.StartRatio = 0.342424
	__Link__000004_RectAnchoredRects.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_RectAnchoredRects.EndRatio = 0.444444
	__Link__000004_RectAnchoredRects.CornerOffsetRatio = 1.297778

	// Link values setup
	__Link__000005_RectAnchoredTexts.Name = `RectAnchoredTexts`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.RectAnchoredTexts]
	__Link__000005_RectAnchoredTexts.Identifier = `ref_models.Rect.RectAnchoredTexts`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectAnchoredText]
	__Link__000005_RectAnchoredTexts.Fieldtypename = `ref_models.RectAnchoredText`
	__Link__000005_RectAnchoredTexts.FieldOffsetX = -137.000000
	__Link__000005_RectAnchoredTexts.FieldOffsetY = -17.000000
	__Link__000005_RectAnchoredTexts.TargetMultiplicity = models.MANY
	__Link__000005_RectAnchoredTexts.TargetMultiplicityOffsetX = 180.000000
	__Link__000005_RectAnchoredTexts.TargetMultiplicityOffsetY = -8.000000
	__Link__000005_RectAnchoredTexts.SourceMultiplicity = models.ZERO_ONE
	__Link__000005_RectAnchoredTexts.SourceMultiplicityOffsetX = 15.000000
	__Link__000005_RectAnchoredTexts.SourceMultiplicityOffsetY = 26.000000
	__Link__000005_RectAnchoredTexts.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_RectAnchoredTexts.StartRatio = 0.075758
	__Link__000005_RectAnchoredTexts.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_RectAnchoredTexts.EndRatio = 0.460317
	__Link__000005_RectAnchoredTexts.CornerOffsetRatio = 1.264444

	// Link values setup
	__Link__000006_RectLinkLinks.Name = `RectLinkLinks`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.RectLinkLinks]
	__Link__000006_RectLinkLinks.Identifier = `ref_models.Layer.RectLinkLinks`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectLinkLink]
	__Link__000006_RectLinkLinks.Fieldtypename = `ref_models.RectLinkLink`
	__Link__000006_RectLinkLinks.FieldOffsetX = -98.000000
	__Link__000006_RectLinkLinks.FieldOffsetY = -12.000000
	__Link__000006_RectLinkLinks.TargetMultiplicity = models.MANY
	__Link__000006_RectLinkLinks.TargetMultiplicityOffsetX = -34.000000
	__Link__000006_RectLinkLinks.TargetMultiplicityOffsetY = 21.000000
	__Link__000006_RectLinkLinks.SourceMultiplicity = models.ZERO_ONE
	__Link__000006_RectLinkLinks.SourceMultiplicityOffsetX = 12.000000
	__Link__000006_RectLinkLinks.SourceMultiplicityOffsetY = 24.000000
	__Link__000006_RectLinkLinks.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_RectLinkLinks.StartRatio = 0.454545
	__Link__000006_RectLinkLinks.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_RectLinkLinks.EndRatio = 0.460317
	__Link__000006_RectLinkLinks.CornerOffsetRatio = 1.670833

	// Link values setup
	__Link__000007_Rects.Name = `Rects`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Rects]
	__Link__000007_Rects.Identifier = `ref_models.Layer.Rects`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000007_Rects.Fieldtypename = `ref_models.Rect`
	__Link__000007_Rects.FieldOffsetX = -56.000000
	__Link__000007_Rects.FieldOffsetY = -18.000000
	__Link__000007_Rects.TargetMultiplicity = models.MANY
	__Link__000007_Rects.TargetMultiplicityOffsetX = -31.000000
	__Link__000007_Rects.TargetMultiplicityOffsetY = 36.000000
	__Link__000007_Rects.SourceMultiplicity = models.ZERO_ONE
	__Link__000007_Rects.SourceMultiplicityOffsetX = 15.000000
	__Link__000007_Rects.SourceMultiplicityOffsetY = 28.000000
	__Link__000007_Rects.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_Rects.StartRatio = 0.076033
	__Link__000007_Rects.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_Rects.EndRatio = 0.630303
	__Link__000007_Rects.CornerOffsetRatio = 1.304167

	// Link values setup
	__Link__000008_Start.Name = `Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Start]
	__Link__000008_Start.Identifier = `ref_models.Link.Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000008_Start.Fieldtypename = `ref_models.Rect`
	__Link__000008_Start.FieldOffsetX = -55.000000
	__Link__000008_Start.FieldOffsetY = 24.000000
	__Link__000008_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000008_Start.TargetMultiplicityOffsetX = 12.000000
	__Link__000008_Start.TargetMultiplicityOffsetY = 25.000000
	__Link__000008_Start.SourceMultiplicity = models.MANY
	__Link__000008_Start.SourceMultiplicityOffsetX = 18.000000
	__Link__000008_Start.SourceMultiplicityOffsetY = -7.000000
	__Link__000008_Start.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000008_Start.StartRatio = 0.210505
	__Link__000008_Start.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000008_Start.EndRatio = 0.219181
	__Link__000008_Start.CornerOffsetRatio = -0.479798

	// Link values setup
	__Link__000009_Start.Name = `Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectLinkLink.Start]
	__Link__000009_Start.Identifier = `ref_models.RectLinkLink.Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000009_Start.Fieldtypename = `ref_models.Rect`
	__Link__000009_Start.FieldOffsetX = -49.000000
	__Link__000009_Start.FieldOffsetY = -13.000000
	__Link__000009_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000009_Start.TargetMultiplicityOffsetX = -43.000000
	__Link__000009_Start.TargetMultiplicityOffsetY = 26.000000
	__Link__000009_Start.SourceMultiplicity = models.MANY
	__Link__000009_Start.SourceMultiplicityOffsetX = 12.000000
	__Link__000009_Start.SourceMultiplicityOffsetY = -9.000000
	__Link__000009_Start.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000009_Start.StartRatio = 0.768611
	__Link__000009_Start.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000009_Start.EndRatio = 0.912121
	__Link__000009_Start.CornerOffsetRatio = -0.015873

	// Link values setup
	__Link__000010_TextAtArrowEnd.Name = `TextAtArrowEnd`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.TextAtArrowEnd]
	__Link__000010_TextAtArrowEnd.Identifier = `ref_models.Link.TextAtArrowEnd`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LinkAnchoredText]
	__Link__000010_TextAtArrowEnd.Fieldtypename = `ref_models.LinkAnchoredText`
	__Link__000010_TextAtArrowEnd.FieldOffsetX = -126.000000
	__Link__000010_TextAtArrowEnd.FieldOffsetY = -22.000000
	__Link__000010_TextAtArrowEnd.TargetMultiplicity = models.MANY
	__Link__000010_TextAtArrowEnd.TargetMultiplicityOffsetX = -31.000000
	__Link__000010_TextAtArrowEnd.TargetMultiplicityOffsetY = 20.000000
	__Link__000010_TextAtArrowEnd.SourceMultiplicity = models.ZERO_ONE
	__Link__000010_TextAtArrowEnd.SourceMultiplicityOffsetX = 10.000000
	__Link__000010_TextAtArrowEnd.SourceMultiplicityOffsetY = 20.000000
	__Link__000010_TextAtArrowEnd.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000010_TextAtArrowEnd.StartRatio = 0.661616
	__Link__000010_TextAtArrowEnd.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000010_TextAtArrowEnd.EndRatio = 0.862745
	__Link__000010_TextAtArrowEnd.CornerOffsetRatio = 1.316566

	// Link values setup
	__Link__000011_TextAtArrowStart.Name = `TextAtArrowStart`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.TextAtArrowStart]
	__Link__000011_TextAtArrowStart.Identifier = `ref_models.Link.TextAtArrowStart`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LinkAnchoredText]
	__Link__000011_TextAtArrowStart.Fieldtypename = `ref_models.LinkAnchoredText`
	__Link__000011_TextAtArrowStart.FieldOffsetX = -128.000000
	__Link__000011_TextAtArrowStart.FieldOffsetY = -16.000000
	__Link__000011_TextAtArrowStart.TargetMultiplicity = models.MANY
	__Link__000011_TextAtArrowStart.TargetMultiplicityOffsetX = -33.000000
	__Link__000011_TextAtArrowStart.TargetMultiplicityOffsetY = 23.000000
	__Link__000011_TextAtArrowStart.SourceMultiplicity = models.ZERO_ONE
	__Link__000011_TextAtArrowStart.SourceMultiplicityOffsetX = 11.000000
	__Link__000011_TextAtArrowStart.SourceMultiplicityOffsetY = 22.000000
	__Link__000011_TextAtArrowStart.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000011_TextAtArrowStart.StartRatio = 0.217172
	__Link__000011_TextAtArrowStart.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000011_TextAtArrowStart.EndRatio = 0.307190
	__Link__000011_TextAtArrowStart.CornerOffsetRatio = 1.204444

	// Position values setup
	__Position__000000_Pos_NewDiagram_AnchoredText.X = 1171.000000
	__Position__000000_Pos_NewDiagram_AnchoredText.Y = 619.000000
	__Position__000000_Pos_NewDiagram_AnchoredText.Name = `Pos-NewDiagram-AnchoredText`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Layer.X = 44.000000
	__Position__000001_Pos_NewDiagram_Layer.Y = 204.000000
	__Position__000001_Pos_NewDiagram_Layer.Name = `Pos-NewDiagram-Layer`

	// Position values setup
	__Position__000002_Pos_NewDiagram_Link.X = 604.000000
	__Position__000002_Pos_NewDiagram_Link.Y = 622.000000
	__Position__000002_Pos_NewDiagram_Link.Name = `Pos-NewDiagram-Link`

	// Position values setup
	__Position__000003_Pos_NewDiagram_Rect.X = 610.000000
	__Position__000003_Pos_NewDiagram_Rect.Y = 43.000000
	__Position__000003_Pos_NewDiagram_Rect.Name = `Pos-NewDiagram-Rect`

	// Position values setup
	__Position__000004_Pos_NewDiagram_RectAnchoredRect.X = 1146.000000
	__Position__000004_Pos_NewDiagram_RectAnchoredRect.Y = 122.000000
	__Position__000004_Pos_NewDiagram_RectAnchoredRect.Name = `Pos-NewDiagram-RectAnchoredRect`

	// Position values setup
	__Position__000005_Pos_NewDiagram_RectAnchoredText.X = 1141.000000
	__Position__000005_Pos_NewDiagram_RectAnchoredText.Y = 37.000000
	__Position__000005_Pos_NewDiagram_RectAnchoredText.Name = `Pos-NewDiagram-RectAnchoredText`

	// Position values setup
	__Position__000006_Pos_NewDiagram_RectLinkLink.X = 308.000000
	__Position__000006_Pos_NewDiagram_RectLinkLink.Y = 446.000000
	__Position__000006_Pos_NewDiagram_RectLinkLink.Name = `Pos-NewDiagram-RectLinkLink`

	// Position values setup
	__Position__000007_Pos_NewDiagram_SVG.X = 44.000000
	__Position__000007_Pos_NewDiagram_SVG.Y = 28.000000
	__Position__000007_Pos_NewDiagram_SVG.Name = `Pos-NewDiagram-SVG`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Link.X = 384.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Link.Y = 548.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Link.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Link`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Rect.X = 348.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Rect.Y = 684.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Rect.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Rect`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_RectLinkLink.X = 405.500000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_RectLinkLink.Y = 633.500000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_RectLinkLink.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-RectLinkLink`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_AnchoredText.X = 1408.500000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_AnchoredText.Y = 638.000000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_AnchoredText.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Link and NewDiagram-AnchoredText`

	// Vertice values setup
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_AnchoredText.X = 1408.500000
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_AnchoredText.Y = 638.000000
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_AnchoredText.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Link and NewDiagram-AnchoredText`

	// Vertice values setup
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect.X = 955.000000
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect.Y = 516.500000
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Link and NewDiagram-Rect`

	// Vertice values setup
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect.X = 965.000000
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect.Y = 576.500000
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Link and NewDiagram-Rect`

	// Vertice values setup
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredRect.X = 1025.000000
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredRect.Y = 744.000000
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredRect.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Rect and NewDiagram-RectAnchoredRect`

	// Vertice values setup
	__Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredText.X = 965.000000
	__Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredText.Y = 619.000000
	__Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredText.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Rect and NewDiagram-RectAnchoredText`

	// Vertice values setup
	__Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_RectLinkLink_and_NewDiagram_Link.X = 829.000000
	__Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_RectLinkLink_and_NewDiagram_Link.Y = 567.500000
	__Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_RectLinkLink_and_NewDiagram_Link.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-RectLinkLink and NewDiagram-Link`

	// Vertice values setup
	__Vertice__000010_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_RectLinkLink_and_NewDiagram_Rect.X = 832.000000
	__Vertice__000010_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_RectLinkLink_and_NewDiagram_Rect.Y = 278.000000
	__Vertice__000010_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_RectLinkLink_and_NewDiagram_Rect.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-RectLinkLink and NewDiagram-Rect`

	// Vertice values setup
	__Vertice__000011_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Layer.X = 53.000000
	__Vertice__000011_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Layer.Y = 284.000000
	__Vertice__000011_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Layer.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Layer`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_AnchoredText)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Layer)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_Link)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_Rect)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000004_NewDiagram_RectAnchoredRect)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000005_NewDiagram_RectAnchoredText)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000006_NewDiagram_RectLinkLink)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000007_NewDiagram_SVG)
	__GongStructShape__000000_NewDiagram_AnchoredText.Position = __Position__000000_Pos_NewDiagram_AnchoredText
	__GongStructShape__000000_NewDiagram_AnchoredText.Fields = append(__GongStructShape__000000_NewDiagram_AnchoredText.Fields, __Field__000023_Name)
	__GongStructShape__000000_NewDiagram_AnchoredText.Fields = append(__GongStructShape__000000_NewDiagram_AnchoredText.Fields, __Field__000007_Content)
	__GongStructShape__000000_NewDiagram_AnchoredText.Fields = append(__GongStructShape__000000_NewDiagram_AnchoredText.Fields, __Field__000034_X_Offset)
	__GongStructShape__000000_NewDiagram_AnchoredText.Fields = append(__GongStructShape__000000_NewDiagram_AnchoredText.Fields, __Field__000036_Y_Offset)
	__GongStructShape__000000_NewDiagram_AnchoredText.Fields = append(__GongStructShape__000000_NewDiagram_AnchoredText.Fields, __Field__000006_Color)
	__GongStructShape__000000_NewDiagram_AnchoredText.Fields = append(__GongStructShape__000000_NewDiagram_AnchoredText.Fields, __Field__000030_Stroke)
	__GongStructShape__000001_NewDiagram_Layer.Position = __Position__000001_Pos_NewDiagram_Layer
	__GongStructShape__000001_NewDiagram_Layer.Fields = append(__GongStructShape__000001_NewDiagram_Layer.Fields, __Field__000010_Display)
	__GongStructShape__000001_NewDiagram_Layer.Fields = append(__GongStructShape__000001_NewDiagram_Layer.Fields, __Field__000024_Name)
	__GongStructShape__000001_NewDiagram_Layer.Links = append(__GongStructShape__000001_NewDiagram_Layer.Links, __Link__000006_RectLinkLinks)
	__GongStructShape__000001_NewDiagram_Layer.Links = append(__GongStructShape__000001_NewDiagram_Layer.Links, __Link__000003_Links)
	__GongStructShape__000001_NewDiagram_Layer.Links = append(__GongStructShape__000001_NewDiagram_Layer.Links, __Link__000007_Rects)
	__GongStructShape__000002_NewDiagram_Link.Position = __Position__000002_Pos_NewDiagram_Link
	__GongStructShape__000002_NewDiagram_Link.Fields = append(__GongStructShape__000002_NewDiagram_Link.Fields, __Field__000027_StartAnchorType)
	__GongStructShape__000002_NewDiagram_Link.Fields = append(__GongStructShape__000002_NewDiagram_Link.Fields, __Field__000011_EndAnchorType)
	__GongStructShape__000002_NewDiagram_Link.Fields = append(__GongStructShape__000002_NewDiagram_Link.Fields, __Field__000028_StartOrientation)
	__GongStructShape__000002_NewDiagram_Link.Fields = append(__GongStructShape__000002_NewDiagram_Link.Fields, __Field__000029_StartRatio)
	__GongStructShape__000002_NewDiagram_Link.Fields = append(__GongStructShape__000002_NewDiagram_Link.Fields, __Field__000013_EndOrientation)
	__GongStructShape__000002_NewDiagram_Link.Fields = append(__GongStructShape__000002_NewDiagram_Link.Fields, __Field__000014_EndRatio)
	__GongStructShape__000002_NewDiagram_Link.Fields = append(__GongStructShape__000002_NewDiagram_Link.Fields, __Field__000008_CornerOffsetRatio)
	__GongStructShape__000002_NewDiagram_Link.Fields = append(__GongStructShape__000002_NewDiagram_Link.Fields, __Field__000009_CornerRadius)
	__GongStructShape__000002_NewDiagram_Link.Fields = append(__GongStructShape__000002_NewDiagram_Link.Fields, __Field__000012_EndArrowSize)
	__GongStructShape__000002_NewDiagram_Link.Links = append(__GongStructShape__000002_NewDiagram_Link.Links, __Link__000000_End)
	__GongStructShape__000002_NewDiagram_Link.Links = append(__GongStructShape__000002_NewDiagram_Link.Links, __Link__000008_Start)
	__GongStructShape__000002_NewDiagram_Link.Links = append(__GongStructShape__000002_NewDiagram_Link.Links, __Link__000010_TextAtArrowEnd)
	__GongStructShape__000002_NewDiagram_Link.Links = append(__GongStructShape__000002_NewDiagram_Link.Links, __Link__000011_TextAtArrowStart)
	__GongStructShape__000003_NewDiagram_Rect.Position = __Position__000003_Pos_NewDiagram_Rect
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000025_Name)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000033_X)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000035_Y)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000032_Width)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000019_Height)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000026_RX)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000020_IsSelectable)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000021_IsSelected)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000001_CanHaveLeftHandle)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000016_HasLeftHandle)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000002_CanHaveRightHandle)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000017_HasRightHandle)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000003_CanHaveTopHandle)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000018_HasTopHandle)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000000_CanHaveBottomHandle)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000015_HasBottomHandle)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000004_CanMoveHorizontaly)
	__GongStructShape__000003_NewDiagram_Rect.Fields = append(__GongStructShape__000003_NewDiagram_Rect.Fields, __Field__000005_CanMoveVerticaly)
	__GongStructShape__000003_NewDiagram_Rect.Links = append(__GongStructShape__000003_NewDiagram_Rect.Links, __Link__000004_RectAnchoredRects)
	__GongStructShape__000003_NewDiagram_Rect.Links = append(__GongStructShape__000003_NewDiagram_Rect.Links, __Link__000005_RectAnchoredTexts)
	__GongStructShape__000004_NewDiagram_RectAnchoredRect.Position = __Position__000004_Pos_NewDiagram_RectAnchoredRect
	__GongStructShape__000005_NewDiagram_RectAnchoredText.Position = __Position__000005_Pos_NewDiagram_RectAnchoredText
	__GongStructShape__000006_NewDiagram_RectLinkLink.Position = __Position__000006_Pos_NewDiagram_RectLinkLink
	__GongStructShape__000006_NewDiagram_RectLinkLink.Fields = append(__GongStructShape__000006_NewDiagram_RectLinkLink.Fields, __Field__000031_TargetAnchorPosition)
	__GongStructShape__000006_NewDiagram_RectLinkLink.Links = append(__GongStructShape__000006_NewDiagram_RectLinkLink.Links, __Link__000009_Start)
	__GongStructShape__000006_NewDiagram_RectLinkLink.Links = append(__GongStructShape__000006_NewDiagram_RectLinkLink.Links, __Link__000001_End)
	__GongStructShape__000007_NewDiagram_SVG.Position = __Position__000007_Pos_NewDiagram_SVG
	__GongStructShape__000007_NewDiagram_SVG.Fields = append(__GongStructShape__000007_NewDiagram_SVG.Fields, __Field__000022_Name)
	__GongStructShape__000007_NewDiagram_SVG.Links = append(__GongStructShape__000007_NewDiagram_SVG.Links, __Link__000002_Layers)
	__Link__000000_End.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect
	__Link__000001_End.Middlevertice = __Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_RectLinkLink_and_NewDiagram_Link
	__Link__000002_Layers.Middlevertice = __Vertice__000011_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Layer
	__Link__000003_Links.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Link
	__Link__000004_RectAnchoredRects.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredRect
	__Link__000005_RectAnchoredTexts.Middlevertice = __Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredText
	__Link__000006_RectLinkLinks.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_RectLinkLink
	__Link__000007_Rects.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Rect
	__Link__000008_Start.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect
	__Link__000009_Start.Middlevertice = __Vertice__000010_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_RectLinkLink_and_NewDiagram_Rect
	__Link__000010_TextAtArrowEnd.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_AnchoredText
	__Link__000011_TextAtArrowStart.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_AnchoredText
}
