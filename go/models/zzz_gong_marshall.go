// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage{{Identifiers}}

	// insertion point for initialization of values{{ValueInitializers}}

	// insertion point for setup of pointers{{PointersInitializers}}
}`

const GongIdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)` /* */

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const MetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}`

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *Stage) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Println(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)

	res, err := stage.MarshallToString(modelsPackageName, packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}

	fmt.Fprintln(file, res)
}

// MarshallToString marshall the stage content into a string
func (stage *Stage) MarshallToString(modelsPackageName, packageName string) (res string, err error) {

	res = marshallRes
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	axisOrdered := []*Axis{}
	for axis := range stage.Axiss {
		axisOrdered = append(axisOrdered, axis)
	}
	sort.Slice(axisOrdered[:], func(i, j int) bool {
		axisi := axisOrdered[i]
		axisj := axisOrdered[j]
		axisi_order, oki := stage.AxisMap_Staged_Order[axisi]
		axisj_order, okj := stage.AxisMap_Staged_Order[axisj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return axisi_order < axisj_order
	})
	if len(axisOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, axis := range axisOrdered {

		identifiersDecl += axis.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += axis.GongMarshallField(stage, "Name")
		initializerStatements += axis.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += axis.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += axis.GongMarshallField(stage, "AngleDegree")
		initializerStatements += axis.GongMarshallField(stage, "Length")
		initializerStatements += axis.GongMarshallField(stage, "CenterX")
		initializerStatements += axis.GongMarshallField(stage, "CenterY")
		initializerStatements += axis.GongMarshallField(stage, "EndX")
		initializerStatements += axis.GongMarshallField(stage, "EndY")
		initializerStatements += axis.GongMarshallField(stage, "Color")
		initializerStatements += axis.GongMarshallField(stage, "FillOpacity")
		initializerStatements += axis.GongMarshallField(stage, "Stroke")
		initializerStatements += axis.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += axis.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += axis.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += axis.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += axis.GongMarshallField(stage, "Transform")
	}

	axisgridOrdered := []*AxisGrid{}
	for axisgrid := range stage.AxisGrids {
		axisgridOrdered = append(axisgridOrdered, axisgrid)
	}
	sort.Slice(axisgridOrdered[:], func(i, j int) bool {
		axisgridi := axisgridOrdered[i]
		axisgridj := axisgridOrdered[j]
		axisgridi_order, oki := stage.AxisGridMap_Staged_Order[axisgridi]
		axisgridj_order, okj := stage.AxisGridMap_Staged_Order[axisgridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return axisgridi_order < axisgridj_order
	})
	if len(axisgridOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, axisgrid := range axisgridOrdered {

		identifiersDecl += axisgrid.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += axisgrid.GongMarshallField(stage, "Name")
		pointersInitializesStatements += axisgrid.GongMarshallField(stage, "Reference")
		initializerStatements += axisgrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += axisgrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += axisgrid.GongMarshallField(stage, "Axiss")
	}

	bezierOrdered := []*Bezier{}
	for bezier := range stage.Beziers {
		bezierOrdered = append(bezierOrdered, bezier)
	}
	sort.Slice(bezierOrdered[:], func(i, j int) bool {
		bezieri := bezierOrdered[i]
		bezierj := bezierOrdered[j]
		bezieri_order, oki := stage.BezierMap_Staged_Order[bezieri]
		bezierj_order, okj := stage.BezierMap_Staged_Order[bezierj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return bezieri_order < bezierj_order
	})
	if len(bezierOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, bezier := range bezierOrdered {

		identifiersDecl += bezier.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += bezier.GongMarshallField(stage, "Name")
		initializerStatements += bezier.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += bezier.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += bezier.GongMarshallField(stage, "StartX")
		initializerStatements += bezier.GongMarshallField(stage, "StartY")
		initializerStatements += bezier.GongMarshallField(stage, "ControlPointStartX")
		initializerStatements += bezier.GongMarshallField(stage, "ControlPointStartY")
		initializerStatements += bezier.GongMarshallField(stage, "EndX")
		initializerStatements += bezier.GongMarshallField(stage, "EndY")
		initializerStatements += bezier.GongMarshallField(stage, "ControlPointEndX")
		initializerStatements += bezier.GongMarshallField(stage, "ControlPointEndY")
		initializerStatements += bezier.GongMarshallField(stage, "Color")
		initializerStatements += bezier.GongMarshallField(stage, "FillOpacity")
		initializerStatements += bezier.GongMarshallField(stage, "Stroke")
		initializerStatements += bezier.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += bezier.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += bezier.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += bezier.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += bezier.GongMarshallField(stage, "Transform")
	}

	beziergridOrdered := []*BezierGrid{}
	for beziergrid := range stage.BezierGrids {
		beziergridOrdered = append(beziergridOrdered, beziergrid)
	}
	sort.Slice(beziergridOrdered[:], func(i, j int) bool {
		beziergridi := beziergridOrdered[i]
		beziergridj := beziergridOrdered[j]
		beziergridi_order, oki := stage.BezierGridMap_Staged_Order[beziergridi]
		beziergridj_order, okj := stage.BezierGridMap_Staged_Order[beziergridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return beziergridi_order < beziergridj_order
	})
	if len(beziergridOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, beziergrid := range beziergridOrdered {

		identifiersDecl += beziergrid.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += beziergrid.GongMarshallField(stage, "Name")
		pointersInitializesStatements += beziergrid.GongMarshallField(stage, "Reference")
		initializerStatements += beziergrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += beziergrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += beziergrid.GongMarshallField(stage, "Beziers")
	}

	beziergridstackOrdered := []*BezierGridStack{}
	for beziergridstack := range stage.BezierGridStacks {
		beziergridstackOrdered = append(beziergridstackOrdered, beziergridstack)
	}
	sort.Slice(beziergridstackOrdered[:], func(i, j int) bool {
		beziergridstacki := beziergridstackOrdered[i]
		beziergridstackj := beziergridstackOrdered[j]
		beziergridstacki_order, oki := stage.BezierGridStackMap_Staged_Order[beziergridstacki]
		beziergridstackj_order, okj := stage.BezierGridStackMap_Staged_Order[beziergridstackj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return beziergridstacki_order < beziergridstackj_order
	})
	if len(beziergridstackOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, beziergridstack := range beziergridstackOrdered {

		identifiersDecl += beziergridstack.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += beziergridstack.GongMarshallField(stage, "Name")
		initializerStatements += beziergridstack.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += beziergridstack.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += beziergridstack.GongMarshallField(stage, "BezierGrids")
	}

	chapterOrdered := []*Chapter{}
	for chapter := range stage.Chapters {
		chapterOrdered = append(chapterOrdered, chapter)
	}
	sort.Slice(chapterOrdered[:], func(i, j int) bool {
		chapteri := chapterOrdered[i]
		chapterj := chapterOrdered[j]
		chapteri_order, oki := stage.ChapterMap_Staged_Order[chapteri]
		chapterj_order, okj := stage.ChapterMap_Staged_Order[chapterj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return chapteri_order < chapterj_order
	})
	if len(chapterOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, chapter := range chapterOrdered {

		identifiersDecl += chapter.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += chapter.GongMarshallField(stage, "Name")
		initializerStatements += chapter.GongMarshallField(stage, "MardownContent")
	}

	circleOrdered := []*Circle{}
	for circle := range stage.Circles {
		circleOrdered = append(circleOrdered, circle)
	}
	sort.Slice(circleOrdered[:], func(i, j int) bool {
		circlei := circleOrdered[i]
		circlej := circleOrdered[j]
		circlei_order, oki := stage.CircleMap_Staged_Order[circlei]
		circlej_order, okj := stage.CircleMap_Staged_Order[circlej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return circlei_order < circlej_order
	})
	if len(circleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, circle := range circleOrdered {

		identifiersDecl += circle.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += circle.GongMarshallField(stage, "Name")
		initializerStatements += circle.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += circle.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += circle.GongMarshallField(stage, "CenterX")
		initializerStatements += circle.GongMarshallField(stage, "CenterY")
		initializerStatements += circle.GongMarshallField(stage, "HasBespokeRadius")
		initializerStatements += circle.GongMarshallField(stage, "BespopkeRadius")
		initializerStatements += circle.GongMarshallField(stage, "Color")
		initializerStatements += circle.GongMarshallField(stage, "FillOpacity")
		initializerStatements += circle.GongMarshallField(stage, "Stroke")
		initializerStatements += circle.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += circle.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += circle.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += circle.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += circle.GongMarshallField(stage, "Transform")
		initializerStatements += circle.GongMarshallField(stage, "Pitch")
		initializerStatements += circle.GongMarshallField(stage, "ShowName")
		initializerStatements += circle.GongMarshallField(stage, "BeatNb")
	}

	circlegridOrdered := []*CircleGrid{}
	for circlegrid := range stage.CircleGrids {
		circlegridOrdered = append(circlegridOrdered, circlegrid)
	}
	sort.Slice(circlegridOrdered[:], func(i, j int) bool {
		circlegridi := circlegridOrdered[i]
		circlegridj := circlegridOrdered[j]
		circlegridi_order, oki := stage.CircleGridMap_Staged_Order[circlegridi]
		circlegridj_order, okj := stage.CircleGridMap_Staged_Order[circlegridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return circlegridi_order < circlegridj_order
	})
	if len(circlegridOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, circlegrid := range circlegridOrdered {

		identifiersDecl += circlegrid.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += circlegrid.GongMarshallField(stage, "Name")
		pointersInitializesStatements += circlegrid.GongMarshallField(stage, "Reference")
		initializerStatements += circlegrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += circlegrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += circlegrid.GongMarshallField(stage, "Circles")
	}

	contentOrdered := []*Content{}
	for content := range stage.Contents {
		contentOrdered = append(contentOrdered, content)
	}
	sort.Slice(contentOrdered[:], func(i, j int) bool {
		contenti := contentOrdered[i]
		contentj := contentOrdered[j]
		contenti_order, oki := stage.ContentMap_Staged_Order[contenti]
		contentj_order, okj := stage.ContentMap_Staged_Order[contentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return contenti_order < contentj_order
	})
	if len(contentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, content := range contentOrdered {

		identifiersDecl += content.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += content.GongMarshallField(stage, "Name")
		initializerStatements += content.GongMarshallField(stage, "MardownContent")
		initializerStatements += content.GongMarshallField(stage, "ContentPath")
		initializerStatements += content.GongMarshallField(stage, "OutputPath")
		initializerStatements += content.GongMarshallField(stage, "LayoutPath")
		initializerStatements += content.GongMarshallField(stage, "StaticPath")
		initializerStatements += content.GongMarshallField(stage, "Target")
		pointersInitializesStatements += content.GongMarshallField(stage, "Chapters")
	}

	exporttomusicxmlOrdered := []*ExportToMusicxml{}
	for exporttomusicxml := range stage.ExportToMusicxmls {
		exporttomusicxmlOrdered = append(exporttomusicxmlOrdered, exporttomusicxml)
	}
	sort.Slice(exporttomusicxmlOrdered[:], func(i, j int) bool {
		exporttomusicxmli := exporttomusicxmlOrdered[i]
		exporttomusicxmlj := exporttomusicxmlOrdered[j]
		exporttomusicxmli_order, oki := stage.ExportToMusicxmlMap_Staged_Order[exporttomusicxmli]
		exporttomusicxmlj_order, okj := stage.ExportToMusicxmlMap_Staged_Order[exporttomusicxmlj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return exporttomusicxmli_order < exporttomusicxmlj_order
	})
	if len(exporttomusicxmlOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, exporttomusicxml := range exporttomusicxmlOrdered {

		identifiersDecl += exporttomusicxml.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += exporttomusicxml.GongMarshallField(stage, "Name")
		pointersInitializesStatements += exporttomusicxml.GongMarshallField(stage, "Parameter")
	}

	frontcurveOrdered := []*FrontCurve{}
	for frontcurve := range stage.FrontCurves {
		frontcurveOrdered = append(frontcurveOrdered, frontcurve)
	}
	sort.Slice(frontcurveOrdered[:], func(i, j int) bool {
		frontcurvei := frontcurveOrdered[i]
		frontcurvej := frontcurveOrdered[j]
		frontcurvei_order, oki := stage.FrontCurveMap_Staged_Order[frontcurvei]
		frontcurvej_order, okj := stage.FrontCurveMap_Staged_Order[frontcurvej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return frontcurvei_order < frontcurvej_order
	})
	if len(frontcurveOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, frontcurve := range frontcurveOrdered {

		identifiersDecl += frontcurve.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += frontcurve.GongMarshallField(stage, "Name")
		initializerStatements += frontcurve.GongMarshallField(stage, "Path")
	}

	frontcurvestackOrdered := []*FrontCurveStack{}
	for frontcurvestack := range stage.FrontCurveStacks {
		frontcurvestackOrdered = append(frontcurvestackOrdered, frontcurvestack)
	}
	sort.Slice(frontcurvestackOrdered[:], func(i, j int) bool {
		frontcurvestacki := frontcurvestackOrdered[i]
		frontcurvestackj := frontcurvestackOrdered[j]
		frontcurvestacki_order, oki := stage.FrontCurveStackMap_Staged_Order[frontcurvestacki]
		frontcurvestackj_order, okj := stage.FrontCurveStackMap_Staged_Order[frontcurvestackj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return frontcurvestacki_order < frontcurvestackj_order
	})
	if len(frontcurvestackOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, frontcurvestack := range frontcurvestackOrdered {

		identifiersDecl += frontcurvestack.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += frontcurvestack.GongMarshallField(stage, "Name")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += frontcurvestack.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += frontcurvestack.GongMarshallField(stage, "FrontCurves")
		pointersInitializesStatements += frontcurvestack.GongMarshallField(stage, "SpiralCircles")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "Color")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "FillOpacity")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "Stroke")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "Transform")
	}

	horizontalaxisOrdered := []*HorizontalAxis{}
	for horizontalaxis := range stage.HorizontalAxiss {
		horizontalaxisOrdered = append(horizontalaxisOrdered, horizontalaxis)
	}
	sort.Slice(horizontalaxisOrdered[:], func(i, j int) bool {
		horizontalaxisi := horizontalaxisOrdered[i]
		horizontalaxisj := horizontalaxisOrdered[j]
		horizontalaxisi_order, oki := stage.HorizontalAxisMap_Staged_Order[horizontalaxisi]
		horizontalaxisj_order, okj := stage.HorizontalAxisMap_Staged_Order[horizontalaxisj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return horizontalaxisi_order < horizontalaxisj_order
	})
	if len(horizontalaxisOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, horizontalaxis := range horizontalaxisOrdered {

		identifiersDecl += horizontalaxis.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += horizontalaxis.GongMarshallField(stage, "Name")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += horizontalaxis.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "AxisHandleBorderLength")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "Axis_Length")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "Color")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "FillOpacity")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "Stroke")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "Transform")
	}

	keyOrdered := []*Key{}
	for key := range stage.Keys {
		keyOrdered = append(keyOrdered, key)
	}
	sort.Slice(keyOrdered[:], func(i, j int) bool {
		keyi := keyOrdered[i]
		keyj := keyOrdered[j]
		keyi_order, oki := stage.KeyMap_Staged_Order[keyi]
		keyj_order, okj := stage.KeyMap_Staged_Order[keyj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return keyi_order < keyj_order
	})
	if len(keyOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, key := range keyOrdered {

		identifiersDecl += key.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += key.GongMarshallField(stage, "Name")
		initializerStatements += key.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += key.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += key.GongMarshallField(stage, "Path")
		initializerStatements += key.GongMarshallField(stage, "Color")
		initializerStatements += key.GongMarshallField(stage, "FillOpacity")
		initializerStatements += key.GongMarshallField(stage, "Stroke")
		initializerStatements += key.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += key.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += key.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += key.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += key.GongMarshallField(stage, "Transform")
	}

	parameterOrdered := []*Parameter{}
	for parameter := range stage.Parameters {
		parameterOrdered = append(parameterOrdered, parameter)
	}
	sort.Slice(parameterOrdered[:], func(i, j int) bool {
		parameteri := parameterOrdered[i]
		parameterj := parameterOrdered[j]
		parameteri_order, oki := stage.ParameterMap_Staged_Order[parameteri]
		parameterj_order, okj := stage.ParameterMap_Staged_Order[parameterj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return parameteri_order < parameterj_order
	})
	if len(parameterOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, parameter := range parameterOrdered {

		identifiersDecl += parameter.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += parameter.GongMarshallField(stage, "Name")
		initializerStatements += parameter.GongMarshallField(stage, "BackendColor")
		initializerStatements += parameter.GongMarshallField(stage, "MinuteColor")
		initializerStatements += parameter.GongMarshallField(stage, "HourColor")
		initializerStatements += parameter.GongMarshallField(stage, "N")
		initializerStatements += parameter.GongMarshallField(stage, "M")
		initializerStatements += parameter.GongMarshallField(stage, "Z")
		initializerStatements += parameter.GongMarshallField(stage, "InsideAngle")
		initializerStatements += parameter.GongMarshallField(stage, "ShiftToNearestCircle")
		initializerStatements += parameter.GongMarshallField(stage, "SideLength")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "InitialRhombus")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "InitialCircle")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "InitialRhombusGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "InitialCircleGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "InitialAxis")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "RotatedAxis")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "RotatedRhombus")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "RotatedRhombusGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "RotatedCircleGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "NextRhombus")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "NextCircle")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowingRhombusGridSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowingRhombusGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowingCircleGridSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowingCircleGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowingCircleGridLeftSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowingCircleGridLeft")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "ConstructionAxis")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "ConstructionAxisGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "ConstructionCircle")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "ConstructionCircleGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurve")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveShiftedRightSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveShiftedRight")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveNextSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveNext")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveNextShiftedRightSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveNextShiftedRight")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveStack")
		initializerStatements += parameter.GongMarshallField(stage, "StackWidth")
		initializerStatements += parameter.GongMarshallField(stage, "NbShitRight")
		initializerStatements += parameter.GongMarshallField(stage, "StackHeight")
		initializerStatements += parameter.GongMarshallField(stage, "BezierControlLengthRatio")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralRhombusGridSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralRhombusGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralCircleSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralCircleGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralCircleFullGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralConstructionOuterLineSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralConstructionInnerLineSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralConstructionOuterLineGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralConstructionInnerLineGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralConstructionCircleGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralConstructionOuterLineFullGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralBezierSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralBezierGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralBezierFullGrid")
		initializerStatements += parameter.GongMarshallField(stage, "SpiralBezierStrength")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "FrontCurveStack")
		initializerStatements += parameter.GongMarshallField(stage, "NbInterpolationPoints")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "Fkey")
		initializerStatements += parameter.GongMarshallField(stage, "FkeySizeRatio")
		initializerStatements += parameter.GongMarshallField(stage, "FkeyOriginRelativeX")
		initializerStatements += parameter.GongMarshallField(stage, "FkeyOriginRelativeY")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "PitchLines")
		initializerStatements += parameter.GongMarshallField(stage, "PitchHeight")
		initializerStatements += parameter.GongMarshallField(stage, "NbPitchLines")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "BeatLines")
		initializerStatements += parameter.GongMarshallField(stage, "BeatLinesHeightRatio")
		initializerStatements += parameter.GongMarshallField(stage, "NbBeatLines")
		initializerStatements += parameter.GongMarshallField(stage, "NbOfBeatsInTheme")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "FirstVoice")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "FirstVoiceShiftedRigth")
		initializerStatements += parameter.GongMarshallField(stage, "FirstVoiceShiftX")
		initializerStatements += parameter.GongMarshallField(stage, "FirstVoiceShiftY")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SecondVoice")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SecondVoiceShiftedRight")
		initializerStatements += parameter.GongMarshallField(stage, "PitchDifference")
		initializerStatements += parameter.GongMarshallField(stage, "BeatsPerSecond")
		initializerStatements += parameter.GongMarshallField(stage, "Level")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "FirstVoiceNotes")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "FirstVoiceNotesShiftedRight")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SecondVoiceNotes")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SecondVoiceNotesShiftedRight")
		initializerStatements += parameter.GongMarshallField(stage, "IsMinor")
		initializerStatements += parameter.GongMarshallField(stage, "ThemeBinaryEncoding")
		initializerStatements += parameter.GongMarshallField(stage, "OriginX")
		initializerStatements += parameter.GongMarshallField(stage, "OriginY")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "HorizontalAxis")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "VerticalAxis")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralOrigin")
		initializerStatements += parameter.GongMarshallField(stage, "SpiralOriginX")
		initializerStatements += parameter.GongMarshallField(stage, "SpiralOriginY")
		initializerStatements += parameter.GongMarshallField(stage, "OriginCrossWidth")
		initializerStatements += parameter.GongMarshallField(stage, "SpiralRadiusRatio")
		initializerStatements += parameter.GongMarshallField(stage, "ShowSpiralBezierConstruct")
		initializerStatements += parameter.GongMarshallField(stage, "ShowInterpolationPoints")
		initializerStatements += parameter.GongMarshallField(stage, "ActualBeatsTemporalShift")
		initializerStatements += parameter.GongMarshallField(stage, "PathToStaticFiles")
		initializerStatements += parameter.GongMarshallField(stage, "PathToGeneratedSVG")
		initializerStatements += parameter.GongMarshallField(stage, "PathToGeneratedScore")
	}

	rhombusOrdered := []*Rhombus{}
	for rhombus := range stage.Rhombuss {
		rhombusOrdered = append(rhombusOrdered, rhombus)
	}
	sort.Slice(rhombusOrdered[:], func(i, j int) bool {
		rhombusi := rhombusOrdered[i]
		rhombusj := rhombusOrdered[j]
		rhombusi_order, oki := stage.RhombusMap_Staged_Order[rhombusi]
		rhombusj_order, okj := stage.RhombusMap_Staged_Order[rhombusj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return rhombusi_order < rhombusj_order
	})
	if len(rhombusOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, rhombus := range rhombusOrdered {

		identifiersDecl += rhombus.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += rhombus.GongMarshallField(stage, "Name")
		initializerStatements += rhombus.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += rhombus.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += rhombus.GongMarshallField(stage, "CenterX")
		initializerStatements += rhombus.GongMarshallField(stage, "CenterY")
		initializerStatements += rhombus.GongMarshallField(stage, "SideLength")
		initializerStatements += rhombus.GongMarshallField(stage, "AngleDegree")
		initializerStatements += rhombus.GongMarshallField(stage, "InsideAngle")
		initializerStatements += rhombus.GongMarshallField(stage, "Color")
		initializerStatements += rhombus.GongMarshallField(stage, "FillOpacity")
		initializerStatements += rhombus.GongMarshallField(stage, "Stroke")
		initializerStatements += rhombus.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += rhombus.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += rhombus.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += rhombus.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += rhombus.GongMarshallField(stage, "Transform")
	}

	rhombusgridOrdered := []*RhombusGrid{}
	for rhombusgrid := range stage.RhombusGrids {
		rhombusgridOrdered = append(rhombusgridOrdered, rhombusgrid)
	}
	sort.Slice(rhombusgridOrdered[:], func(i, j int) bool {
		rhombusgridi := rhombusgridOrdered[i]
		rhombusgridj := rhombusgridOrdered[j]
		rhombusgridi_order, oki := stage.RhombusGridMap_Staged_Order[rhombusgridi]
		rhombusgridj_order, okj := stage.RhombusGridMap_Staged_Order[rhombusgridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return rhombusgridi_order < rhombusgridj_order
	})
	if len(rhombusgridOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, rhombusgrid := range rhombusgridOrdered {

		identifiersDecl += rhombusgrid.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += rhombusgrid.GongMarshallField(stage, "Name")
		pointersInitializesStatements += rhombusgrid.GongMarshallField(stage, "Reference")
		initializerStatements += rhombusgrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += rhombusgrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += rhombusgrid.GongMarshallField(stage, "Rhombuses")
	}

	shapecategoryOrdered := []*ShapeCategory{}
	for shapecategory := range stage.ShapeCategorys {
		shapecategoryOrdered = append(shapecategoryOrdered, shapecategory)
	}
	sort.Slice(shapecategoryOrdered[:], func(i, j int) bool {
		shapecategoryi := shapecategoryOrdered[i]
		shapecategoryj := shapecategoryOrdered[j]
		shapecategoryi_order, oki := stage.ShapeCategoryMap_Staged_Order[shapecategoryi]
		shapecategoryj_order, okj := stage.ShapeCategoryMap_Staged_Order[shapecategoryj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return shapecategoryi_order < shapecategoryj_order
	})
	if len(shapecategoryOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, shapecategory := range shapecategoryOrdered {

		identifiersDecl += shapecategory.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += shapecategory.GongMarshallField(stage, "Name")
		initializerStatements += shapecategory.GongMarshallField(stage, "IsExpanded")
	}

	spiralbezierOrdered := []*SpiralBezier{}
	for spiralbezier := range stage.SpiralBeziers {
		spiralbezierOrdered = append(spiralbezierOrdered, spiralbezier)
	}
	sort.Slice(spiralbezierOrdered[:], func(i, j int) bool {
		spiralbezieri := spiralbezierOrdered[i]
		spiralbezierj := spiralbezierOrdered[j]
		spiralbezieri_order, oki := stage.SpiralBezierMap_Staged_Order[spiralbezieri]
		spiralbezierj_order, okj := stage.SpiralBezierMap_Staged_Order[spiralbezierj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralbezieri_order < spiralbezierj_order
	})
	if len(spiralbezierOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spiralbezier := range spiralbezierOrdered {

		identifiersDecl += spiralbezier.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += spiralbezier.GongMarshallField(stage, "Name")
		initializerStatements += spiralbezier.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralbezier.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += spiralbezier.GongMarshallField(stage, "StartX")
		initializerStatements += spiralbezier.GongMarshallField(stage, "StartY")
		initializerStatements += spiralbezier.GongMarshallField(stage, "ControlPointStartX")
		initializerStatements += spiralbezier.GongMarshallField(stage, "ControlPointStartY")
		initializerStatements += spiralbezier.GongMarshallField(stage, "EndX")
		initializerStatements += spiralbezier.GongMarshallField(stage, "EndY")
		initializerStatements += spiralbezier.GongMarshallField(stage, "ControlPointEndX")
		initializerStatements += spiralbezier.GongMarshallField(stage, "ControlPointEndY")
		initializerStatements += spiralbezier.GongMarshallField(stage, "Color")
		initializerStatements += spiralbezier.GongMarshallField(stage, "FillOpacity")
		initializerStatements += spiralbezier.GongMarshallField(stage, "Stroke")
		initializerStatements += spiralbezier.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += spiralbezier.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += spiralbezier.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += spiralbezier.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += spiralbezier.GongMarshallField(stage, "Transform")
	}

	spiralbeziergridOrdered := []*SpiralBezierGrid{}
	for spiralbeziergrid := range stage.SpiralBezierGrids {
		spiralbeziergridOrdered = append(spiralbeziergridOrdered, spiralbeziergrid)
	}
	sort.Slice(spiralbeziergridOrdered[:], func(i, j int) bool {
		spiralbeziergridi := spiralbeziergridOrdered[i]
		spiralbeziergridj := spiralbeziergridOrdered[j]
		spiralbeziergridi_order, oki := stage.SpiralBezierGridMap_Staged_Order[spiralbeziergridi]
		spiralbeziergridj_order, okj := stage.SpiralBezierGridMap_Staged_Order[spiralbeziergridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralbeziergridi_order < spiralbeziergridj_order
	})
	if len(spiralbeziergridOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spiralbeziergrid := range spiralbeziergridOrdered {

		identifiersDecl += spiralbeziergrid.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += spiralbeziergrid.GongMarshallField(stage, "Name")
		initializerStatements += spiralbeziergrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralbeziergrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += spiralbeziergrid.GongMarshallField(stage, "SpiralBeziers")
	}

	spiralcircleOrdered := []*SpiralCircle{}
	for spiralcircle := range stage.SpiralCircles {
		spiralcircleOrdered = append(spiralcircleOrdered, spiralcircle)
	}
	sort.Slice(spiralcircleOrdered[:], func(i, j int) bool {
		spiralcirclei := spiralcircleOrdered[i]
		spiralcirclej := spiralcircleOrdered[j]
		spiralcirclei_order, oki := stage.SpiralCircleMap_Staged_Order[spiralcirclei]
		spiralcirclej_order, okj := stage.SpiralCircleMap_Staged_Order[spiralcirclej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralcirclei_order < spiralcirclej_order
	})
	if len(spiralcircleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spiralcircle := range spiralcircleOrdered {

		identifiersDecl += spiralcircle.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += spiralcircle.GongMarshallField(stage, "Name")
		initializerStatements += spiralcircle.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralcircle.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += spiralcircle.GongMarshallField(stage, "CenterX")
		initializerStatements += spiralcircle.GongMarshallField(stage, "CenterY")
		initializerStatements += spiralcircle.GongMarshallField(stage, "HasBespokeRadius")
		initializerStatements += spiralcircle.GongMarshallField(stage, "BespopkeRadius")
		initializerStatements += spiralcircle.GongMarshallField(stage, "Color")
		initializerStatements += spiralcircle.GongMarshallField(stage, "FillOpacity")
		initializerStatements += spiralcircle.GongMarshallField(stage, "Stroke")
		initializerStatements += spiralcircle.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += spiralcircle.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += spiralcircle.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += spiralcircle.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += spiralcircle.GongMarshallField(stage, "Transform")
		initializerStatements += spiralcircle.GongMarshallField(stage, "Pitch")
		initializerStatements += spiralcircle.GongMarshallField(stage, "ShowName")
		initializerStatements += spiralcircle.GongMarshallField(stage, "BeatNb")
		initializerStatements += spiralcircle.GongMarshallField(stage, "Path")
	}

	spiralcirclegridOrdered := []*SpiralCircleGrid{}
	for spiralcirclegrid := range stage.SpiralCircleGrids {
		spiralcirclegridOrdered = append(spiralcirclegridOrdered, spiralcirclegrid)
	}
	sort.Slice(spiralcirclegridOrdered[:], func(i, j int) bool {
		spiralcirclegridi := spiralcirclegridOrdered[i]
		spiralcirclegridj := spiralcirclegridOrdered[j]
		spiralcirclegridi_order, oki := stage.SpiralCircleGridMap_Staged_Order[spiralcirclegridi]
		spiralcirclegridj_order, okj := stage.SpiralCircleGridMap_Staged_Order[spiralcirclegridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralcirclegridi_order < spiralcirclegridj_order
	})
	if len(spiralcirclegridOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spiralcirclegrid := range spiralcirclegridOrdered {

		identifiersDecl += spiralcirclegrid.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += spiralcirclegrid.GongMarshallField(stage, "Name")
		initializerStatements += spiralcirclegrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralcirclegrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += spiralcirclegrid.GongMarshallField(stage, "SpiralRhombusGrid")
		pointersInitializesStatements += spiralcirclegrid.GongMarshallField(stage, "SpiralCircles")
	}

	spirallineOrdered := []*SpiralLine{}
	for spiralline := range stage.SpiralLines {
		spirallineOrdered = append(spirallineOrdered, spiralline)
	}
	sort.Slice(spirallineOrdered[:], func(i, j int) bool {
		spirallinei := spirallineOrdered[i]
		spirallinej := spirallineOrdered[j]
		spirallinei_order, oki := stage.SpiralLineMap_Staged_Order[spirallinei]
		spirallinej_order, okj := stage.SpiralLineMap_Staged_Order[spirallinej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spirallinei_order < spirallinej_order
	})
	if len(spirallineOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spiralline := range spirallineOrdered {

		identifiersDecl += spiralline.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += spiralline.GongMarshallField(stage, "Name")
		initializerStatements += spiralline.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralline.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += spiralline.GongMarshallField(stage, "StartX")
		initializerStatements += spiralline.GongMarshallField(stage, "EndX")
		initializerStatements += spiralline.GongMarshallField(stage, "StartY")
		initializerStatements += spiralline.GongMarshallField(stage, "EndY")
		initializerStatements += spiralline.GongMarshallField(stage, "Color")
		initializerStatements += spiralline.GongMarshallField(stage, "FillOpacity")
		initializerStatements += spiralline.GongMarshallField(stage, "Stroke")
		initializerStatements += spiralline.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += spiralline.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += spiralline.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += spiralline.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += spiralline.GongMarshallField(stage, "Transform")
	}

	spirallinegridOrdered := []*SpiralLineGrid{}
	for spirallinegrid := range stage.SpiralLineGrids {
		spirallinegridOrdered = append(spirallinegridOrdered, spirallinegrid)
	}
	sort.Slice(spirallinegridOrdered[:], func(i, j int) bool {
		spirallinegridi := spirallinegridOrdered[i]
		spirallinegridj := spirallinegridOrdered[j]
		spirallinegridi_order, oki := stage.SpiralLineGridMap_Staged_Order[spirallinegridi]
		spirallinegridj_order, okj := stage.SpiralLineGridMap_Staged_Order[spirallinegridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spirallinegridi_order < spirallinegridj_order
	})
	if len(spirallinegridOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spirallinegrid := range spirallinegridOrdered {

		identifiersDecl += spirallinegrid.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += spirallinegrid.GongMarshallField(stage, "Name")
		initializerStatements += spirallinegrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spirallinegrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += spirallinegrid.GongMarshallField(stage, "SpiralLines")
	}

	spiraloriginOrdered := []*SpiralOrigin{}
	for spiralorigin := range stage.SpiralOrigins {
		spiraloriginOrdered = append(spiraloriginOrdered, spiralorigin)
	}
	sort.Slice(spiraloriginOrdered[:], func(i, j int) bool {
		spiralorigini := spiraloriginOrdered[i]
		spiraloriginj := spiraloriginOrdered[j]
		spiralorigini_order, oki := stage.SpiralOriginMap_Staged_Order[spiralorigini]
		spiraloriginj_order, okj := stage.SpiralOriginMap_Staged_Order[spiraloriginj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralorigini_order < spiraloriginj_order
	})
	if len(spiraloriginOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spiralorigin := range spiraloriginOrdered {

		identifiersDecl += spiralorigin.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += spiralorigin.GongMarshallField(stage, "Name")
		initializerStatements += spiralorigin.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralorigin.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += spiralorigin.GongMarshallField(stage, "Color")
		initializerStatements += spiralorigin.GongMarshallField(stage, "FillOpacity")
		initializerStatements += spiralorigin.GongMarshallField(stage, "Stroke")
		initializerStatements += spiralorigin.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += spiralorigin.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += spiralorigin.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += spiralorigin.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += spiralorigin.GongMarshallField(stage, "Transform")
	}

	spiralrhombusOrdered := []*SpiralRhombus{}
	for spiralrhombus := range stage.SpiralRhombuss {
		spiralrhombusOrdered = append(spiralrhombusOrdered, spiralrhombus)
	}
	sort.Slice(spiralrhombusOrdered[:], func(i, j int) bool {
		spiralrhombusi := spiralrhombusOrdered[i]
		spiralrhombusj := spiralrhombusOrdered[j]
		spiralrhombusi_order, oki := stage.SpiralRhombusMap_Staged_Order[spiralrhombusi]
		spiralrhombusj_order, okj := stage.SpiralRhombusMap_Staged_Order[spiralrhombusj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralrhombusi_order < spiralrhombusj_order
	})
	if len(spiralrhombusOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spiralrhombus := range spiralrhombusOrdered {

		identifiersDecl += spiralrhombus.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Name")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralrhombus.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "X_r0")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Y_r0")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "X_r1")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Y_r1")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "X_r2")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Y_r2")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "X_r3")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Y_r3")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Color")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "FillOpacity")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Stroke")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Transform")
	}

	spiralrhombusgridOrdered := []*SpiralRhombusGrid{}
	for spiralrhombusgrid := range stage.SpiralRhombusGrids {
		spiralrhombusgridOrdered = append(spiralrhombusgridOrdered, spiralrhombusgrid)
	}
	sort.Slice(spiralrhombusgridOrdered[:], func(i, j int) bool {
		spiralrhombusgridi := spiralrhombusgridOrdered[i]
		spiralrhombusgridj := spiralrhombusgridOrdered[j]
		spiralrhombusgridi_order, oki := stage.SpiralRhombusGridMap_Staged_Order[spiralrhombusgridi]
		spiralrhombusgridj_order, okj := stage.SpiralRhombusGridMap_Staged_Order[spiralrhombusgridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralrhombusgridi_order < spiralrhombusgridj_order
	})
	if len(spiralrhombusgridOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, spiralrhombusgrid := range spiralrhombusgridOrdered {

		identifiersDecl += spiralrhombusgrid.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += spiralrhombusgrid.GongMarshallField(stage, "Name")
		initializerStatements += spiralrhombusgrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralrhombusgrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += spiralrhombusgrid.GongMarshallField(stage, "SpiralRhombuses")
	}

	verticalaxisOrdered := []*VerticalAxis{}
	for verticalaxis := range stage.VerticalAxiss {
		verticalaxisOrdered = append(verticalaxisOrdered, verticalaxis)
	}
	sort.Slice(verticalaxisOrdered[:], func(i, j int) bool {
		verticalaxisi := verticalaxisOrdered[i]
		verticalaxisj := verticalaxisOrdered[j]
		verticalaxisi_order, oki := stage.VerticalAxisMap_Staged_Order[verticalaxisi]
		verticalaxisj_order, okj := stage.VerticalAxisMap_Staged_Order[verticalaxisj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return verticalaxisi_order < verticalaxisj_order
	})
	if len(verticalaxisOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, verticalaxis := range verticalaxisOrdered {

		identifiersDecl += verticalaxis.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += verticalaxis.GongMarshallField(stage, "Name")
		initializerStatements += verticalaxis.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += verticalaxis.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += verticalaxis.GongMarshallField(stage, "AxisHandleBorderLength")
		initializerStatements += verticalaxis.GongMarshallField(stage, "Axis_Length")
		initializerStatements += verticalaxis.GongMarshallField(stage, "Color")
		initializerStatements += verticalaxis.GongMarshallField(stage, "FillOpacity")
		initializerStatements += verticalaxis.GongMarshallField(stage, "Stroke")
		initializerStatements += verticalaxis.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += verticalaxis.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += verticalaxis.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += verticalaxis.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += verticalaxis.GongMarshallField(stage, "Transform")
	}

	// insertion initialization of objects to stage
	for _, axis := range axisOrdered {
		_ = axis
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, axisgrid := range axisgridOrdered {
		_ = axisgrid
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, bezier := range bezierOrdered {
		_ = bezier
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, beziergrid := range beziergridOrdered {
		_ = beziergrid
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, beziergridstack := range beziergridstackOrdered {
		_ = beziergridstack
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, chapter := range chapterOrdered {
		_ = chapter
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, circle := range circleOrdered {
		_ = circle
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, circlegrid := range circlegridOrdered {
		_ = circlegrid
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, content := range contentOrdered {
		_ = content
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, exporttomusicxml := range exporttomusicxmlOrdered {
		_ = exporttomusicxml
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, frontcurve := range frontcurveOrdered {
		_ = frontcurve
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, frontcurvestack := range frontcurvestackOrdered {
		_ = frontcurvestack
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, horizontalaxis := range horizontalaxisOrdered {
		_ = horizontalaxis
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, key := range keyOrdered {
		_ = key
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, parameter := range parameterOrdered {
		_ = parameter
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, rhombus := range rhombusOrdered {
		_ = rhombus
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, rhombusgrid := range rhombusgridOrdered {
		_ = rhombusgrid
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, shapecategory := range shapecategoryOrdered {
		_ = shapecategory
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spiralbezier := range spiralbezierOrdered {
		_ = spiralbezier
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spiralbeziergrid := range spiralbeziergridOrdered {
		_ = spiralbeziergrid
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spiralcircle := range spiralcircleOrdered {
		_ = spiralcircle
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spiralcirclegrid := range spiralcirclegridOrdered {
		_ = spiralcirclegrid
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spiralline := range spirallineOrdered {
		_ = spiralline
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spirallinegrid := range spirallinegridOrdered {
		_ = spirallinegrid
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spiralorigin := range spiraloriginOrdered {
		_ = spiralorigin
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spiralrhombus := range spiralrhombusOrdered {
		_ = spiralrhombus
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spiralrhombusgrid := range spiralrhombusgridOrdered {
		_ = spiralrhombusgrid
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, verticalaxis := range verticalaxisOrdered {
		_ = verticalaxis
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}
	return
}

// insertion point for marshall field methods
func (axis *Axis) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(axis.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", axis.IsDisplayed))
	case "AngleDegree":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AngleDegree")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.AngleDegree))
	case "Length":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Length")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.Length))
	case "CenterX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CenterX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.CenterX))
	case "CenterY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CenterY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.CenterY))
	case "EndX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.EndX))
	case "EndY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.EndY))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(axis.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(axis.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(axis.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(axis.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(axis.Transform))

	case "ShapeCategory":
		if axis.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", axis.ShapeCategory.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Axis", fieldName)
	}
	return
}

func (axisgrid *AxisGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axisgrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(axisgrid.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axisgrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", axisgrid.IsDisplayed))

	case "Reference":
		if axisgrid.Reference != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", axisgrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Reference")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", axisgrid.Reference.GongGetIdentifier(stage))
		}
	case "ShapeCategory":
		if axisgrid.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", axisgrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", axisgrid.ShapeCategory.GongGetIdentifier(stage))
		}
	case "Axiss":
		for _, _axis := range axisgrid.Axiss {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", axisgrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Axiss")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _axis.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AxisGrid", fieldName)
	}
	return
}

func (bezier *Bezier) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(bezier.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", bezier.IsDisplayed))
	case "StartX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.StartX))
	case "StartY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.StartY))
	case "ControlPointStartX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControlPointStartX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.ControlPointStartX))
	case "ControlPointStartY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControlPointStartY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.ControlPointStartY))
	case "EndX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.EndX))
	case "EndY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.EndY))
	case "ControlPointEndX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControlPointEndX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.ControlPointEndX))
	case "ControlPointEndY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControlPointEndY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.ControlPointEndY))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(bezier.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(bezier.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(bezier.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(bezier.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(bezier.Transform))

	case "ShapeCategory":
		if bezier.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", bezier.ShapeCategory.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Bezier", fieldName)
	}
	return
}

func (beziergrid *BezierGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", beziergrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(beziergrid.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", beziergrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", beziergrid.IsDisplayed))

	case "Reference":
		if beziergrid.Reference != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", beziergrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Reference")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", beziergrid.Reference.GongGetIdentifier(stage))
		}
	case "ShapeCategory":
		if beziergrid.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", beziergrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", beziergrid.ShapeCategory.GongGetIdentifier(stage))
		}
	case "Beziers":
		for _, _bezier := range beziergrid.Beziers {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", beziergrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Beziers")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _bezier.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct BezierGrid", fieldName)
	}
	return
}

func (beziergridstack *BezierGridStack) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", beziergridstack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(beziergridstack.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", beziergridstack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", beziergridstack.IsDisplayed))

	case "ShapeCategory":
		if beziergridstack.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", beziergridstack.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", beziergridstack.ShapeCategory.GongGetIdentifier(stage))
		}
	case "BezierGrids":
		for _, _beziergrid := range beziergridstack.BezierGrids {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", beziergridstack.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "BezierGrids")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _beziergrid.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct BezierGridStack", fieldName)
	}
	return
}

func (chapter *Chapter) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", chapter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(chapter.Name))
	case "MardownContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", chapter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MardownContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(chapter.MardownContent))

	default:
		log.Panicf("Unknown field %s for Gongstruct Chapter", fieldName)
	}
	return
}

func (circle *Circle) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circle.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", circle.IsDisplayed))
	case "CenterX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CenterX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.CenterX))
	case "CenterY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CenterY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.CenterY))
	case "HasBespokeRadius":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasBespokeRadius")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", circle.HasBespokeRadius))
	case "BespopkeRadius":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BespopkeRadius")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.BespopkeRadius))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circle.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circle.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circle.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circle.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circle.Transform))
	case "Pitch":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Pitch")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", circle.Pitch))
	case "ShowName":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShowName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", circle.ShowName))
	case "BeatNb":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BeatNb")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", circle.BeatNb))

	case "ShapeCategory":
		if circle.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", circle.ShapeCategory.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Circle", fieldName)
	}
	return
}

func (circlegrid *CircleGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circlegrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circlegrid.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circlegrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", circlegrid.IsDisplayed))

	case "Reference":
		if circlegrid.Reference != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", circlegrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Reference")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", circlegrid.Reference.GongGetIdentifier(stage))
		}
	case "ShapeCategory":
		if circlegrid.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", circlegrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", circlegrid.ShapeCategory.GongGetIdentifier(stage))
		}
	case "Circles":
		for _, _circle := range circlegrid.Circles {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", circlegrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Circles")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _circle.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CircleGrid", fieldName)
	}
	return
}

func (content *Content) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.Name))
	case "MardownContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MardownContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.MardownContent))
	case "ContentPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ContentPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.ContentPath))
	case "OutputPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OutputPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.OutputPath))
	case "LayoutPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.LayoutPath))
	case "StaticPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StaticPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.StaticPath))
	case "Target":
		if content.Target != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Target")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+content.Target.ToCodeString())
		}

	case "Chapters":
		for _, _chapter := range content.Chapters {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", content.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Chapters")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _chapter.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Content", fieldName)
	}
	return
}

func (exporttomusicxml *ExportToMusicxml) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", exporttomusicxml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(exporttomusicxml.Name))

	case "Parameter":
		if exporttomusicxml.Parameter != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", exporttomusicxml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Parameter")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", exporttomusicxml.Parameter.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ExportToMusicxml", fieldName)
	}
	return
}

func (frontcurve *FrontCurve) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurve.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(frontcurve.Name))
	case "Path":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurve.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Path")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(frontcurve.Path))

	default:
		log.Panicf("Unknown field %s for Gongstruct FrontCurve", fieldName)
	}
	return
}

func (frontcurvestack *FrontCurveStack) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(frontcurvestack.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", frontcurvestack.IsDisplayed))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(frontcurvestack.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", frontcurvestack.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(frontcurvestack.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", frontcurvestack.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", frontcurvestack.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(frontcurvestack.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(frontcurvestack.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(frontcurvestack.Transform))

	case "ShapeCategory":
		if frontcurvestack.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", frontcurvestack.ShapeCategory.GongGetIdentifier(stage))
		}
	case "FrontCurves":
		for _, _frontcurve := range frontcurvestack.FrontCurves {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "FrontCurves")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _frontcurve.GongGetIdentifier(stage))
			res += tmp
		}
	case "SpiralCircles":
		for _, _spiralcircle := range frontcurvestack.SpiralCircles {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SpiralCircles")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spiralcircle.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FrontCurveStack", fieldName)
	}
	return
}

func (horizontalaxis *HorizontalAxis) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(horizontalaxis.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", horizontalaxis.IsDisplayed))
	case "AxisHandleBorderLength":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AxisHandleBorderLength")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", horizontalaxis.AxisHandleBorderLength))
	case "Axis_Length":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Axis_Length")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", horizontalaxis.Axis_Length))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(horizontalaxis.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", horizontalaxis.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(horizontalaxis.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", horizontalaxis.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", horizontalaxis.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(horizontalaxis.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(horizontalaxis.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(horizontalaxis.Transform))

	case "ShapeCategory":
		if horizontalaxis.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", horizontalaxis.ShapeCategory.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct HorizontalAxis", fieldName)
	}
	return
}

func (key *Key) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(key.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", key.IsDisplayed))
	case "Path":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Path")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(key.Path))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(key.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", key.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(key.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", key.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", key.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(key.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(key.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(key.Transform))

	case "ShapeCategory":
		if key.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", key.ShapeCategory.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Key", fieldName)
	}
	return
}

func (parameter *Parameter) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(parameter.Name))
	case "BackendColor":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BackendColor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(parameter.BackendColor))
	case "MinuteColor":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MinuteColor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(parameter.MinuteColor))
	case "HourColor":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HourColor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(parameter.HourColor))
	case "N":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "N")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.N))
	case "M":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "M")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.M))
	case "Z":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Z")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.Z))
	case "InsideAngle":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InsideAngle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.InsideAngle))
	case "ShiftToNearestCircle":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShiftToNearestCircle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.ShiftToNearestCircle))
	case "SideLength":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SideLength")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.SideLength))
	case "StackWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.StackWidth))
	case "NbShitRight":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbShitRight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.NbShitRight))
	case "StackHeight":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackHeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.StackHeight))
	case "BezierControlLengthRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BezierControlLengthRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.BezierControlLengthRatio))
	case "SpiralBezierStrength":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralBezierStrength")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.SpiralBezierStrength))
	case "NbInterpolationPoints":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbInterpolationPoints")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.NbInterpolationPoints))
	case "FkeySizeRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FkeySizeRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.FkeySizeRatio))
	case "FkeyOriginRelativeX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FkeyOriginRelativeX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.FkeyOriginRelativeX))
	case "FkeyOriginRelativeY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FkeyOriginRelativeY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.FkeyOriginRelativeY))
	case "PitchHeight":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PitchHeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.PitchHeight))
	case "NbPitchLines":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbPitchLines")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.NbPitchLines))
	case "BeatLinesHeightRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BeatLinesHeightRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.BeatLinesHeightRatio))
	case "NbBeatLines":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbBeatLines")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.NbBeatLines))
	case "NbOfBeatsInTheme":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbOfBeatsInTheme")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.NbOfBeatsInTheme))
	case "FirstVoiceShiftX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoiceShiftX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.FirstVoiceShiftX))
	case "FirstVoiceShiftY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoiceShiftY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.FirstVoiceShiftY))
	case "PitchDifference":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PitchDifference")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.PitchDifference))
	case "BeatsPerSecond":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BeatsPerSecond")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.BeatsPerSecond))
	case "Level":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Level")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.Level))
	case "IsMinor":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsMinor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", parameter.IsMinor))
	case "ThemeBinaryEncoding":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ThemeBinaryEncoding")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.ThemeBinaryEncoding))
	case "OriginX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OriginX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.OriginX))
	case "OriginY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OriginY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.OriginY))
	case "SpiralOriginX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralOriginX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.SpiralOriginX))
	case "SpiralOriginY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralOriginY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.SpiralOriginY))
	case "OriginCrossWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OriginCrossWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.OriginCrossWidth))
	case "SpiralRadiusRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralRadiusRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.SpiralRadiusRatio))
	case "ShowSpiralBezierConstruct":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShowSpiralBezierConstruct")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", parameter.ShowSpiralBezierConstruct))
	case "ShowInterpolationPoints":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShowInterpolationPoints")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", parameter.ShowInterpolationPoints))
	case "ActualBeatsTemporalShift":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ActualBeatsTemporalShift")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.ActualBeatsTemporalShift))
	case "PathToStaticFiles":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PathToStaticFiles")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(parameter.PathToStaticFiles))
	case "PathToGeneratedSVG":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PathToGeneratedSVG")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(parameter.PathToGeneratedSVG))
	case "PathToGeneratedScore":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PathToGeneratedScore")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(parameter.PathToGeneratedScore))

	case "InitialRhombus":
		if parameter.InitialRhombus != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialRhombus")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.InitialRhombus.GongGetIdentifier(stage))
		}
	case "InitialCircle":
		if parameter.InitialCircle != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialCircle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.InitialCircle.GongGetIdentifier(stage))
		}
	case "InitialRhombusGrid":
		if parameter.InitialRhombusGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.InitialRhombusGrid.GongGetIdentifier(stage))
		}
	case "InitialCircleGrid":
		if parameter.InitialCircleGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.InitialCircleGrid.GongGetIdentifier(stage))
		}
	case "InitialAxis":
		if parameter.InitialAxis != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.InitialAxis.GongGetIdentifier(stage))
		}
	case "RotatedAxis":
		if parameter.RotatedAxis != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RotatedAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.RotatedAxis.GongGetIdentifier(stage))
		}
	case "RotatedRhombus":
		if parameter.RotatedRhombus != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RotatedRhombus")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.RotatedRhombus.GongGetIdentifier(stage))
		}
	case "RotatedRhombusGrid":
		if parameter.RotatedRhombusGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RotatedRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.RotatedRhombusGrid.GongGetIdentifier(stage))
		}
	case "RotatedCircleGrid":
		if parameter.RotatedCircleGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RotatedCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.RotatedCircleGrid.GongGetIdentifier(stage))
		}
	case "NextRhombus":
		if parameter.NextRhombus != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NextRhombus")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.NextRhombus.GongGetIdentifier(stage))
		}
	case "NextCircle":
		if parameter.NextCircle != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NextCircle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.NextCircle.GongGetIdentifier(stage))
		}
	case "GrowingRhombusGridSeed":
		if parameter.GrowingRhombusGridSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingRhombusGridSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowingRhombusGridSeed.GongGetIdentifier(stage))
		}
	case "GrowingRhombusGrid":
		if parameter.GrowingRhombusGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowingRhombusGrid.GongGetIdentifier(stage))
		}
	case "GrowingCircleGridSeed":
		if parameter.GrowingCircleGridSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingCircleGridSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowingCircleGridSeed.GongGetIdentifier(stage))
		}
	case "GrowingCircleGrid":
		if parameter.GrowingCircleGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowingCircleGrid.GongGetIdentifier(stage))
		}
	case "GrowingCircleGridLeftSeed":
		if parameter.GrowingCircleGridLeftSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingCircleGridLeftSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowingCircleGridLeftSeed.GongGetIdentifier(stage))
		}
	case "GrowingCircleGridLeft":
		if parameter.GrowingCircleGridLeft != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingCircleGridLeft")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowingCircleGridLeft.GongGetIdentifier(stage))
		}
	case "ConstructionAxis":
		if parameter.ConstructionAxis != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConstructionAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.ConstructionAxis.GongGetIdentifier(stage))
		}
	case "ConstructionAxisGrid":
		if parameter.ConstructionAxisGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConstructionAxisGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.ConstructionAxisGrid.GongGetIdentifier(stage))
		}
	case "ConstructionCircle":
		if parameter.ConstructionCircle != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConstructionCircle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.ConstructionCircle.GongGetIdentifier(stage))
		}
	case "ConstructionCircleGrid":
		if parameter.ConstructionCircleGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConstructionCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.ConstructionCircleGrid.GongGetIdentifier(stage))
		}
	case "GrowthCurveSeed":
		if parameter.GrowthCurveSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveSeed.GongGetIdentifier(stage))
		}
	case "GrowthCurve":
		if parameter.GrowthCurve != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurve")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurve.GongGetIdentifier(stage))
		}
	case "GrowthCurveShiftedRightSeed":
		if parameter.GrowthCurveShiftedRightSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveShiftedRightSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveShiftedRightSeed.GongGetIdentifier(stage))
		}
	case "GrowthCurveShiftedRight":
		if parameter.GrowthCurveShiftedRight != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveShiftedRight.GongGetIdentifier(stage))
		}
	case "GrowthCurveNextSeed":
		if parameter.GrowthCurveNextSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveNextSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveNextSeed.GongGetIdentifier(stage))
		}
	case "GrowthCurveNext":
		if parameter.GrowthCurveNext != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveNext")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveNext.GongGetIdentifier(stage))
		}
	case "GrowthCurveNextShiftedRightSeed":
		if parameter.GrowthCurveNextShiftedRightSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveNextShiftedRightSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveNextShiftedRightSeed.GongGetIdentifier(stage))
		}
	case "GrowthCurveNextShiftedRight":
		if parameter.GrowthCurveNextShiftedRight != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveNextShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveNextShiftedRight.GongGetIdentifier(stage))
		}
	case "GrowthCurveStack":
		if parameter.GrowthCurveStack != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveStack")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveStack.GongGetIdentifier(stage))
		}
	case "SpiralRhombusGridSeed":
		if parameter.SpiralRhombusGridSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralRhombusGridSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralRhombusGridSeed.GongGetIdentifier(stage))
		}
	case "SpiralRhombusGrid":
		if parameter.SpiralRhombusGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralRhombusGrid.GongGetIdentifier(stage))
		}
	case "SpiralCircleSeed":
		if parameter.SpiralCircleSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralCircleSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralCircleSeed.GongGetIdentifier(stage))
		}
	case "SpiralCircleGrid":
		if parameter.SpiralCircleGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralCircleGrid.GongGetIdentifier(stage))
		}
	case "SpiralCircleFullGrid":
		if parameter.SpiralCircleFullGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralCircleFullGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralCircleFullGrid.GongGetIdentifier(stage))
		}
	case "SpiralConstructionOuterLineSeed":
		if parameter.SpiralConstructionOuterLineSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionOuterLineSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralConstructionOuterLineSeed.GongGetIdentifier(stage))
		}
	case "SpiralConstructionInnerLineSeed":
		if parameter.SpiralConstructionInnerLineSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionInnerLineSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralConstructionInnerLineSeed.GongGetIdentifier(stage))
		}
	case "SpiralConstructionOuterLineGrid":
		if parameter.SpiralConstructionOuterLineGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionOuterLineGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralConstructionOuterLineGrid.GongGetIdentifier(stage))
		}
	case "SpiralConstructionInnerLineGrid":
		if parameter.SpiralConstructionInnerLineGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionInnerLineGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralConstructionInnerLineGrid.GongGetIdentifier(stage))
		}
	case "SpiralConstructionCircleGrid":
		if parameter.SpiralConstructionCircleGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralConstructionCircleGrid.GongGetIdentifier(stage))
		}
	case "SpiralConstructionOuterLineFullGrid":
		if parameter.SpiralConstructionOuterLineFullGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionOuterLineFullGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralConstructionOuterLineFullGrid.GongGetIdentifier(stage))
		}
	case "SpiralBezierSeed":
		if parameter.SpiralBezierSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralBezierSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralBezierSeed.GongGetIdentifier(stage))
		}
	case "SpiralBezierGrid":
		if parameter.SpiralBezierGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralBezierGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralBezierGrid.GongGetIdentifier(stage))
		}
	case "SpiralBezierFullGrid":
		if parameter.SpiralBezierFullGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralBezierFullGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralBezierFullGrid.GongGetIdentifier(stage))
		}
	case "FrontCurveStack":
		if parameter.FrontCurveStack != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FrontCurveStack")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.FrontCurveStack.GongGetIdentifier(stage))
		}
	case "Fkey":
		if parameter.Fkey != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Fkey")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.Fkey.GongGetIdentifier(stage))
		}
	case "PitchLines":
		if parameter.PitchLines != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PitchLines")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.PitchLines.GongGetIdentifier(stage))
		}
	case "BeatLines":
		if parameter.BeatLines != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BeatLines")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.BeatLines.GongGetIdentifier(stage))
		}
	case "FirstVoice":
		if parameter.FirstVoice != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoice")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.FirstVoice.GongGetIdentifier(stage))
		}
	case "FirstVoiceShiftedRigth":
		if parameter.FirstVoiceShiftedRigth != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoiceShiftedRigth")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.FirstVoiceShiftedRigth.GongGetIdentifier(stage))
		}
	case "SecondVoice":
		if parameter.SecondVoice != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondVoice")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SecondVoice.GongGetIdentifier(stage))
		}
	case "SecondVoiceShiftedRight":
		if parameter.SecondVoiceShiftedRight != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondVoiceShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SecondVoiceShiftedRight.GongGetIdentifier(stage))
		}
	case "FirstVoiceNotes":
		if parameter.FirstVoiceNotes != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoiceNotes")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.FirstVoiceNotes.GongGetIdentifier(stage))
		}
	case "FirstVoiceNotesShiftedRight":
		if parameter.FirstVoiceNotesShiftedRight != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoiceNotesShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.FirstVoiceNotesShiftedRight.GongGetIdentifier(stage))
		}
	case "SecondVoiceNotes":
		if parameter.SecondVoiceNotes != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondVoiceNotes")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SecondVoiceNotes.GongGetIdentifier(stage))
		}
	case "SecondVoiceNotesShiftedRight":
		if parameter.SecondVoiceNotesShiftedRight != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondVoiceNotesShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SecondVoiceNotesShiftedRight.GongGetIdentifier(stage))
		}
	case "HorizontalAxis":
		if parameter.HorizontalAxis != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HorizontalAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.HorizontalAxis.GongGetIdentifier(stage))
		}
	case "VerticalAxis":
		if parameter.VerticalAxis != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "VerticalAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.VerticalAxis.GongGetIdentifier(stage))
		}
	case "SpiralOrigin":
		if parameter.SpiralOrigin != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralOrigin")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralOrigin.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Parameter", fieldName)
	}
	return
}

func (rhombus *Rhombus) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rhombus.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rhombus.IsDisplayed))
	case "CenterX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CenterX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.CenterX))
	case "CenterY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CenterY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.CenterY))
	case "SideLength":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SideLength")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.SideLength))
	case "AngleDegree":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AngleDegree")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.AngleDegree))
	case "InsideAngle":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InsideAngle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.InsideAngle))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rhombus.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rhombus.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rhombus.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rhombus.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rhombus.Transform))

	case "ShapeCategory":
		if rhombus.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", rhombus.ShapeCategory.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Rhombus", fieldName)
	}
	return
}

func (rhombusgrid *RhombusGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombusgrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rhombusgrid.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombusgrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rhombusgrid.IsDisplayed))

	case "Reference":
		if rhombusgrid.Reference != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rhombusgrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Reference")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", rhombusgrid.Reference.GongGetIdentifier(stage))
		}
	case "ShapeCategory":
		if rhombusgrid.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rhombusgrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", rhombusgrid.ShapeCategory.GongGetIdentifier(stage))
		}
	case "Rhombuses":
		for _, _rhombus := range rhombusgrid.Rhombuses {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", rhombusgrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Rhombuses")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _rhombus.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RhombusGrid", fieldName)
	}
	return
}

func (shapecategory *ShapeCategory) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", shapecategory.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(shapecategory.Name))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", shapecategory.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", shapecategory.IsExpanded))

	default:
		log.Panicf("Unknown field %s for Gongstruct ShapeCategory", fieldName)
	}
	return
}

func (spiralbezier *SpiralBezier) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralbezier.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralbezier.IsDisplayed))
	case "StartX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.StartX))
	case "StartY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.StartY))
	case "ControlPointStartX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControlPointStartX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.ControlPointStartX))
	case "ControlPointStartY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControlPointStartY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.ControlPointStartY))
	case "EndX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.EndX))
	case "EndY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.EndY))
	case "ControlPointEndX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControlPointEndX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.ControlPointEndX))
	case "ControlPointEndY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControlPointEndY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.ControlPointEndY))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralbezier.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralbezier.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralbezier.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralbezier.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralbezier.Transform))

	case "ShapeCategory":
		if spiralbezier.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralbezier.ShapeCategory.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpiralBezier", fieldName)
	}
	return
}

func (spiralbeziergrid *SpiralBezierGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbeziergrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralbeziergrid.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbeziergrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralbeziergrid.IsDisplayed))

	case "ShapeCategory":
		if spiralbeziergrid.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralbeziergrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralbeziergrid.ShapeCategory.GongGetIdentifier(stage))
		}
	case "SpiralBeziers":
		for _, _spiralbezier := range spiralbeziergrid.SpiralBeziers {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", spiralbeziergrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SpiralBeziers")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spiralbezier.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpiralBezierGrid", fieldName)
	}
	return
}

func (spiralcircle *SpiralCircle) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralcircle.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralcircle.IsDisplayed))
	case "CenterX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CenterX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.CenterX))
	case "CenterY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CenterY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.CenterY))
	case "HasBespokeRadius":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasBespokeRadius")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralcircle.HasBespokeRadius))
	case "BespopkeRadius":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BespopkeRadius")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.BespopkeRadius))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralcircle.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralcircle.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralcircle.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralcircle.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralcircle.Transform))
	case "Pitch":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Pitch")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", spiralcircle.Pitch))
	case "ShowName":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShowName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralcircle.ShowName))
	case "BeatNb":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BeatNb")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", spiralcircle.BeatNb))
	case "Path":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Path")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralcircle.Path))

	case "ShapeCategory":
		if spiralcircle.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralcircle.ShapeCategory.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpiralCircle", fieldName)
	}
	return
}

func (spiralcirclegrid *SpiralCircleGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcirclegrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralcirclegrid.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcirclegrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralcirclegrid.IsDisplayed))

	case "ShapeCategory":
		if spiralcirclegrid.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralcirclegrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralcirclegrid.ShapeCategory.GongGetIdentifier(stage))
		}
	case "SpiralRhombusGrid":
		if spiralcirclegrid.SpiralRhombusGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralcirclegrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralcirclegrid.SpiralRhombusGrid.GongGetIdentifier(stage))
		}
	case "SpiralCircles":
		for _, _spiralcircle := range spiralcirclegrid.SpiralCircles {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", spiralcirclegrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SpiralCircles")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spiralcircle.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpiralCircleGrid", fieldName)
	}
	return
}

func (spiralline *SpiralLine) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralline.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralline.IsDisplayed))
	case "StartX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.StartX))
	case "EndX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.EndX))
	case "StartY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.StartY))
	case "EndY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.EndY))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralline.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralline.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralline.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralline.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralline.Transform))

	case "ShapeCategory":
		if spiralline.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralline.ShapeCategory.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpiralLine", fieldName)
	}
	return
}

func (spirallinegrid *SpiralLineGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spirallinegrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spirallinegrid.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spirallinegrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spirallinegrid.IsDisplayed))

	case "ShapeCategory":
		if spirallinegrid.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spirallinegrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spirallinegrid.ShapeCategory.GongGetIdentifier(stage))
		}
	case "SpiralLines":
		for _, _spiralline := range spirallinegrid.SpiralLines {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", spirallinegrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SpiralLines")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spiralline.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpiralLineGrid", fieldName)
	}
	return
}

func (spiralorigin *SpiralOrigin) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralorigin.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralorigin.IsDisplayed))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralorigin.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralorigin.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralorigin.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralorigin.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralorigin.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralorigin.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralorigin.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralorigin.Transform))

	case "ShapeCategory":
		if spiralorigin.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralorigin.ShapeCategory.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpiralOrigin", fieldName)
	}
	return
}

func (spiralrhombus *SpiralRhombus) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralrhombus.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralrhombus.IsDisplayed))
	case "X_r0":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X_r0")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.X_r0))
	case "Y_r0":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y_r0")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.Y_r0))
	case "X_r1":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X_r1")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.X_r1))
	case "Y_r1":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y_r1")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.Y_r1))
	case "X_r2":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X_r2")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.X_r2))
	case "Y_r2":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y_r2")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.Y_r2))
	case "X_r3":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X_r3")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.X_r3))
	case "Y_r3":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y_r3")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.Y_r3))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralrhombus.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralrhombus.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralrhombus.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralrhombus.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralrhombus.Transform))

	case "ShapeCategory":
		if spiralrhombus.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralrhombus.ShapeCategory.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpiralRhombus", fieldName)
	}
	return
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombusgrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spiralrhombusgrid.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombusgrid.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralrhombusgrid.IsDisplayed))

	case "ShapeCategory":
		if spiralrhombusgrid.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombusgrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralrhombusgrid.ShapeCategory.GongGetIdentifier(stage))
		}
	case "SpiralRhombuses":
		for _, _spiralrhombus := range spiralrhombusgrid.SpiralRhombuses {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", spiralrhombusgrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SpiralRhombuses")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spiralrhombus.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SpiralRhombusGrid", fieldName)
	}
	return
}

func (verticalaxis *VerticalAxis) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(verticalaxis.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", verticalaxis.IsDisplayed))
	case "AxisHandleBorderLength":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AxisHandleBorderLength")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", verticalaxis.AxisHandleBorderLength))
	case "Axis_Length":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Axis_Length")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", verticalaxis.Axis_Length))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(verticalaxis.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", verticalaxis.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(verticalaxis.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", verticalaxis.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", verticalaxis.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(verticalaxis.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(verticalaxis.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(verticalaxis.Transform))

	case "ShapeCategory":
		if verticalaxis.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", verticalaxis.ShapeCategory.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct VerticalAxis", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (axis *Axis) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += axis.GongMarshallField(stage, "Name")
		initializerStatements += axis.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += axis.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += axis.GongMarshallField(stage, "AngleDegree")
		initializerStatements += axis.GongMarshallField(stage, "Length")
		initializerStatements += axis.GongMarshallField(stage, "CenterX")
		initializerStatements += axis.GongMarshallField(stage, "CenterY")
		initializerStatements += axis.GongMarshallField(stage, "EndX")
		initializerStatements += axis.GongMarshallField(stage, "EndY")
		initializerStatements += axis.GongMarshallField(stage, "Color")
		initializerStatements += axis.GongMarshallField(stage, "FillOpacity")
		initializerStatements += axis.GongMarshallField(stage, "Stroke")
		initializerStatements += axis.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += axis.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += axis.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += axis.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += axis.GongMarshallField(stage, "Transform")
	}
	return
}
func (axisgrid *AxisGrid) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += axisgrid.GongMarshallField(stage, "Name")
		pointersInitializesStatements += axisgrid.GongMarshallField(stage, "Reference")
		initializerStatements += axisgrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += axisgrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += axisgrid.GongMarshallField(stage, "Axiss")
	}
	return
}
func (bezier *Bezier) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += bezier.GongMarshallField(stage, "Name")
		initializerStatements += bezier.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += bezier.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += bezier.GongMarshallField(stage, "StartX")
		initializerStatements += bezier.GongMarshallField(stage, "StartY")
		initializerStatements += bezier.GongMarshallField(stage, "ControlPointStartX")
		initializerStatements += bezier.GongMarshallField(stage, "ControlPointStartY")
		initializerStatements += bezier.GongMarshallField(stage, "EndX")
		initializerStatements += bezier.GongMarshallField(stage, "EndY")
		initializerStatements += bezier.GongMarshallField(stage, "ControlPointEndX")
		initializerStatements += bezier.GongMarshallField(stage, "ControlPointEndY")
		initializerStatements += bezier.GongMarshallField(stage, "Color")
		initializerStatements += bezier.GongMarshallField(stage, "FillOpacity")
		initializerStatements += bezier.GongMarshallField(stage, "Stroke")
		initializerStatements += bezier.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += bezier.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += bezier.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += bezier.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += bezier.GongMarshallField(stage, "Transform")
	}
	return
}
func (beziergrid *BezierGrid) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += beziergrid.GongMarshallField(stage, "Name")
		pointersInitializesStatements += beziergrid.GongMarshallField(stage, "Reference")
		initializerStatements += beziergrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += beziergrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += beziergrid.GongMarshallField(stage, "Beziers")
	}
	return
}
func (beziergridstack *BezierGridStack) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += beziergridstack.GongMarshallField(stage, "Name")
		initializerStatements += beziergridstack.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += beziergridstack.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += beziergridstack.GongMarshallField(stage, "BezierGrids")
	}
	return
}
func (chapter *Chapter) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += chapter.GongMarshallField(stage, "Name")
		initializerStatements += chapter.GongMarshallField(stage, "MardownContent")
	}
	return
}
func (circle *Circle) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += circle.GongMarshallField(stage, "Name")
		initializerStatements += circle.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += circle.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += circle.GongMarshallField(stage, "CenterX")
		initializerStatements += circle.GongMarshallField(stage, "CenterY")
		initializerStatements += circle.GongMarshallField(stage, "HasBespokeRadius")
		initializerStatements += circle.GongMarshallField(stage, "BespopkeRadius")
		initializerStatements += circle.GongMarshallField(stage, "Color")
		initializerStatements += circle.GongMarshallField(stage, "FillOpacity")
		initializerStatements += circle.GongMarshallField(stage, "Stroke")
		initializerStatements += circle.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += circle.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += circle.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += circle.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += circle.GongMarshallField(stage, "Transform")
		initializerStatements += circle.GongMarshallField(stage, "Pitch")
		initializerStatements += circle.GongMarshallField(stage, "ShowName")
		initializerStatements += circle.GongMarshallField(stage, "BeatNb")
	}
	return
}
func (circlegrid *CircleGrid) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += circlegrid.GongMarshallField(stage, "Name")
		pointersInitializesStatements += circlegrid.GongMarshallField(stage, "Reference")
		initializerStatements += circlegrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += circlegrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += circlegrid.GongMarshallField(stage, "Circles")
	}
	return
}
func (content *Content) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += content.GongMarshallField(stage, "Name")
		initializerStatements += content.GongMarshallField(stage, "MardownContent")
		initializerStatements += content.GongMarshallField(stage, "ContentPath")
		initializerStatements += content.GongMarshallField(stage, "OutputPath")
		initializerStatements += content.GongMarshallField(stage, "LayoutPath")
		initializerStatements += content.GongMarshallField(stage, "StaticPath")
		initializerStatements += content.GongMarshallField(stage, "Target")
		pointersInitializesStatements += content.GongMarshallField(stage, "Chapters")
	}
	return
}
func (exporttomusicxml *ExportToMusicxml) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += exporttomusicxml.GongMarshallField(stage, "Name")
		pointersInitializesStatements += exporttomusicxml.GongMarshallField(stage, "Parameter")
	}
	return
}
func (frontcurve *FrontCurve) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += frontcurve.GongMarshallField(stage, "Name")
		initializerStatements += frontcurve.GongMarshallField(stage, "Path")
	}
	return
}
func (frontcurvestack *FrontCurveStack) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += frontcurvestack.GongMarshallField(stage, "Name")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += frontcurvestack.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += frontcurvestack.GongMarshallField(stage, "FrontCurves")
		pointersInitializesStatements += frontcurvestack.GongMarshallField(stage, "SpiralCircles")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "Color")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "FillOpacity")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "Stroke")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += frontcurvestack.GongMarshallField(stage, "Transform")
	}
	return
}
func (horizontalaxis *HorizontalAxis) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += horizontalaxis.GongMarshallField(stage, "Name")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += horizontalaxis.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "AxisHandleBorderLength")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "Axis_Length")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "Color")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "FillOpacity")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "Stroke")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += horizontalaxis.GongMarshallField(stage, "Transform")
	}
	return
}
func (key *Key) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += key.GongMarshallField(stage, "Name")
		initializerStatements += key.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += key.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += key.GongMarshallField(stage, "Path")
		initializerStatements += key.GongMarshallField(stage, "Color")
		initializerStatements += key.GongMarshallField(stage, "FillOpacity")
		initializerStatements += key.GongMarshallField(stage, "Stroke")
		initializerStatements += key.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += key.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += key.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += key.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += key.GongMarshallField(stage, "Transform")
	}
	return
}
func (parameter *Parameter) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += parameter.GongMarshallField(stage, "Name")
		initializerStatements += parameter.GongMarshallField(stage, "BackendColor")
		initializerStatements += parameter.GongMarshallField(stage, "MinuteColor")
		initializerStatements += parameter.GongMarshallField(stage, "HourColor")
		initializerStatements += parameter.GongMarshallField(stage, "N")
		initializerStatements += parameter.GongMarshallField(stage, "M")
		initializerStatements += parameter.GongMarshallField(stage, "Z")
		initializerStatements += parameter.GongMarshallField(stage, "InsideAngle")
		initializerStatements += parameter.GongMarshallField(stage, "ShiftToNearestCircle")
		initializerStatements += parameter.GongMarshallField(stage, "SideLength")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "InitialRhombus")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "InitialCircle")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "InitialRhombusGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "InitialCircleGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "InitialAxis")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "RotatedAxis")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "RotatedRhombus")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "RotatedRhombusGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "RotatedCircleGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "NextRhombus")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "NextCircle")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowingRhombusGridSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowingRhombusGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowingCircleGridSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowingCircleGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowingCircleGridLeftSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowingCircleGridLeft")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "ConstructionAxis")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "ConstructionAxisGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "ConstructionCircle")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "ConstructionCircleGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurve")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveShiftedRightSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveShiftedRight")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveNextSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveNext")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveNextShiftedRightSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveNextShiftedRight")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "GrowthCurveStack")
		initializerStatements += parameter.GongMarshallField(stage, "StackWidth")
		initializerStatements += parameter.GongMarshallField(stage, "NbShitRight")
		initializerStatements += parameter.GongMarshallField(stage, "StackHeight")
		initializerStatements += parameter.GongMarshallField(stage, "BezierControlLengthRatio")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralRhombusGridSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralRhombusGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralCircleSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralCircleGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralCircleFullGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralConstructionOuterLineSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralConstructionInnerLineSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralConstructionOuterLineGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralConstructionInnerLineGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralConstructionCircleGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralConstructionOuterLineFullGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralBezierSeed")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralBezierGrid")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralBezierFullGrid")
		initializerStatements += parameter.GongMarshallField(stage, "SpiralBezierStrength")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "FrontCurveStack")
		initializerStatements += parameter.GongMarshallField(stage, "NbInterpolationPoints")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "Fkey")
		initializerStatements += parameter.GongMarshallField(stage, "FkeySizeRatio")
		initializerStatements += parameter.GongMarshallField(stage, "FkeyOriginRelativeX")
		initializerStatements += parameter.GongMarshallField(stage, "FkeyOriginRelativeY")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "PitchLines")
		initializerStatements += parameter.GongMarshallField(stage, "PitchHeight")
		initializerStatements += parameter.GongMarshallField(stage, "NbPitchLines")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "BeatLines")
		initializerStatements += parameter.GongMarshallField(stage, "BeatLinesHeightRatio")
		initializerStatements += parameter.GongMarshallField(stage, "NbBeatLines")
		initializerStatements += parameter.GongMarshallField(stage, "NbOfBeatsInTheme")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "FirstVoice")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "FirstVoiceShiftedRigth")
		initializerStatements += parameter.GongMarshallField(stage, "FirstVoiceShiftX")
		initializerStatements += parameter.GongMarshallField(stage, "FirstVoiceShiftY")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SecondVoice")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SecondVoiceShiftedRight")
		initializerStatements += parameter.GongMarshallField(stage, "PitchDifference")
		initializerStatements += parameter.GongMarshallField(stage, "BeatsPerSecond")
		initializerStatements += parameter.GongMarshallField(stage, "Level")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "FirstVoiceNotes")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "FirstVoiceNotesShiftedRight")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SecondVoiceNotes")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SecondVoiceNotesShiftedRight")
		initializerStatements += parameter.GongMarshallField(stage, "IsMinor")
		initializerStatements += parameter.GongMarshallField(stage, "ThemeBinaryEncoding")
		initializerStatements += parameter.GongMarshallField(stage, "OriginX")
		initializerStatements += parameter.GongMarshallField(stage, "OriginY")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "HorizontalAxis")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "VerticalAxis")
		pointersInitializesStatements += parameter.GongMarshallField(stage, "SpiralOrigin")
		initializerStatements += parameter.GongMarshallField(stage, "SpiralOriginX")
		initializerStatements += parameter.GongMarshallField(stage, "SpiralOriginY")
		initializerStatements += parameter.GongMarshallField(stage, "OriginCrossWidth")
		initializerStatements += parameter.GongMarshallField(stage, "SpiralRadiusRatio")
		initializerStatements += parameter.GongMarshallField(stage, "ShowSpiralBezierConstruct")
		initializerStatements += parameter.GongMarshallField(stage, "ShowInterpolationPoints")
		initializerStatements += parameter.GongMarshallField(stage, "ActualBeatsTemporalShift")
		initializerStatements += parameter.GongMarshallField(stage, "PathToStaticFiles")
		initializerStatements += parameter.GongMarshallField(stage, "PathToGeneratedSVG")
		initializerStatements += parameter.GongMarshallField(stage, "PathToGeneratedScore")
	}
	return
}
func (rhombus *Rhombus) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += rhombus.GongMarshallField(stage, "Name")
		initializerStatements += rhombus.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += rhombus.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += rhombus.GongMarshallField(stage, "CenterX")
		initializerStatements += rhombus.GongMarshallField(stage, "CenterY")
		initializerStatements += rhombus.GongMarshallField(stage, "SideLength")
		initializerStatements += rhombus.GongMarshallField(stage, "AngleDegree")
		initializerStatements += rhombus.GongMarshallField(stage, "InsideAngle")
		initializerStatements += rhombus.GongMarshallField(stage, "Color")
		initializerStatements += rhombus.GongMarshallField(stage, "FillOpacity")
		initializerStatements += rhombus.GongMarshallField(stage, "Stroke")
		initializerStatements += rhombus.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += rhombus.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += rhombus.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += rhombus.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += rhombus.GongMarshallField(stage, "Transform")
	}
	return
}
func (rhombusgrid *RhombusGrid) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += rhombusgrid.GongMarshallField(stage, "Name")
		pointersInitializesStatements += rhombusgrid.GongMarshallField(stage, "Reference")
		initializerStatements += rhombusgrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += rhombusgrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += rhombusgrid.GongMarshallField(stage, "Rhombuses")
	}
	return
}
func (shapecategory *ShapeCategory) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += shapecategory.GongMarshallField(stage, "Name")
		initializerStatements += shapecategory.GongMarshallField(stage, "IsExpanded")
	}
	return
}
func (spiralbezier *SpiralBezier) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += spiralbezier.GongMarshallField(stage, "Name")
		initializerStatements += spiralbezier.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralbezier.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += spiralbezier.GongMarshallField(stage, "StartX")
		initializerStatements += spiralbezier.GongMarshallField(stage, "StartY")
		initializerStatements += spiralbezier.GongMarshallField(stage, "ControlPointStartX")
		initializerStatements += spiralbezier.GongMarshallField(stage, "ControlPointStartY")
		initializerStatements += spiralbezier.GongMarshallField(stage, "EndX")
		initializerStatements += spiralbezier.GongMarshallField(stage, "EndY")
		initializerStatements += spiralbezier.GongMarshallField(stage, "ControlPointEndX")
		initializerStatements += spiralbezier.GongMarshallField(stage, "ControlPointEndY")
		initializerStatements += spiralbezier.GongMarshallField(stage, "Color")
		initializerStatements += spiralbezier.GongMarshallField(stage, "FillOpacity")
		initializerStatements += spiralbezier.GongMarshallField(stage, "Stroke")
		initializerStatements += spiralbezier.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += spiralbezier.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += spiralbezier.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += spiralbezier.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += spiralbezier.GongMarshallField(stage, "Transform")
	}
	return
}
func (spiralbeziergrid *SpiralBezierGrid) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += spiralbeziergrid.GongMarshallField(stage, "Name")
		initializerStatements += spiralbeziergrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralbeziergrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += spiralbeziergrid.GongMarshallField(stage, "SpiralBeziers")
	}
	return
}
func (spiralcircle *SpiralCircle) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += spiralcircle.GongMarshallField(stage, "Name")
		initializerStatements += spiralcircle.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralcircle.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += spiralcircle.GongMarshallField(stage, "CenterX")
		initializerStatements += spiralcircle.GongMarshallField(stage, "CenterY")
		initializerStatements += spiralcircle.GongMarshallField(stage, "HasBespokeRadius")
		initializerStatements += spiralcircle.GongMarshallField(stage, "BespopkeRadius")
		initializerStatements += spiralcircle.GongMarshallField(stage, "Color")
		initializerStatements += spiralcircle.GongMarshallField(stage, "FillOpacity")
		initializerStatements += spiralcircle.GongMarshallField(stage, "Stroke")
		initializerStatements += spiralcircle.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += spiralcircle.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += spiralcircle.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += spiralcircle.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += spiralcircle.GongMarshallField(stage, "Transform")
		initializerStatements += spiralcircle.GongMarshallField(stage, "Pitch")
		initializerStatements += spiralcircle.GongMarshallField(stage, "ShowName")
		initializerStatements += spiralcircle.GongMarshallField(stage, "BeatNb")
		initializerStatements += spiralcircle.GongMarshallField(stage, "Path")
	}
	return
}
func (spiralcirclegrid *SpiralCircleGrid) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += spiralcirclegrid.GongMarshallField(stage, "Name")
		initializerStatements += spiralcirclegrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralcirclegrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += spiralcirclegrid.GongMarshallField(stage, "SpiralRhombusGrid")
		pointersInitializesStatements += spiralcirclegrid.GongMarshallField(stage, "SpiralCircles")
	}
	return
}
func (spiralline *SpiralLine) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += spiralline.GongMarshallField(stage, "Name")
		initializerStatements += spiralline.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralline.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += spiralline.GongMarshallField(stage, "StartX")
		initializerStatements += spiralline.GongMarshallField(stage, "EndX")
		initializerStatements += spiralline.GongMarshallField(stage, "StartY")
		initializerStatements += spiralline.GongMarshallField(stage, "EndY")
		initializerStatements += spiralline.GongMarshallField(stage, "Color")
		initializerStatements += spiralline.GongMarshallField(stage, "FillOpacity")
		initializerStatements += spiralline.GongMarshallField(stage, "Stroke")
		initializerStatements += spiralline.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += spiralline.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += spiralline.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += spiralline.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += spiralline.GongMarshallField(stage, "Transform")
	}
	return
}
func (spirallinegrid *SpiralLineGrid) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += spirallinegrid.GongMarshallField(stage, "Name")
		initializerStatements += spirallinegrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spirallinegrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += spirallinegrid.GongMarshallField(stage, "SpiralLines")
	}
	return
}
func (spiralorigin *SpiralOrigin) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += spiralorigin.GongMarshallField(stage, "Name")
		initializerStatements += spiralorigin.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralorigin.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += spiralorigin.GongMarshallField(stage, "Color")
		initializerStatements += spiralorigin.GongMarshallField(stage, "FillOpacity")
		initializerStatements += spiralorigin.GongMarshallField(stage, "Stroke")
		initializerStatements += spiralorigin.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += spiralorigin.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += spiralorigin.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += spiralorigin.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += spiralorigin.GongMarshallField(stage, "Transform")
	}
	return
}
func (spiralrhombus *SpiralRhombus) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Name")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralrhombus.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "X_r0")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Y_r0")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "X_r1")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Y_r1")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "X_r2")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Y_r2")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "X_r3")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Y_r3")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Color")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "FillOpacity")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Stroke")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += spiralrhombus.GongMarshallField(stage, "Transform")
	}
	return
}
func (spiralrhombusgrid *SpiralRhombusGrid) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += spiralrhombusgrid.GongMarshallField(stage, "Name")
		initializerStatements += spiralrhombusgrid.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += spiralrhombusgrid.GongMarshallField(stage, "ShapeCategory")
		pointersInitializesStatements += spiralrhombusgrid.GongMarshallField(stage, "SpiralRhombuses")
	}
	return
}
func (verticalaxis *VerticalAxis) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += verticalaxis.GongMarshallField(stage, "Name")
		initializerStatements += verticalaxis.GongMarshallField(stage, "IsDisplayed")
		pointersInitializesStatements += verticalaxis.GongMarshallField(stage, "ShapeCategory")
		initializerStatements += verticalaxis.GongMarshallField(stage, "AxisHandleBorderLength")
		initializerStatements += verticalaxis.GongMarshallField(stage, "Axis_Length")
		initializerStatements += verticalaxis.GongMarshallField(stage, "Color")
		initializerStatements += verticalaxis.GongMarshallField(stage, "FillOpacity")
		initializerStatements += verticalaxis.GongMarshallField(stage, "Stroke")
		initializerStatements += verticalaxis.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += verticalaxis.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += verticalaxis.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += verticalaxis.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += verticalaxis.GongMarshallField(stage, "Transform")
	}
	return
}
