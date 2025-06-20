// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
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

	const __write__local_time = "{{LocalTimeStamp}}"
	const __write__utc_time__ = "{{UTCTimeStamp}}"

	const __commitId__ = "{{CommitId}}"

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

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
		log.Fatalln(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_Axis_Identifiers := make(map[*Axis]string)
	_ = map_Axis_Identifiers

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
	for idx, axis := range axisOrdered {

		id = generatesIdentifier("Axis", idx, axis.Name)
		map_Axis_Identifiers[axis] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Axis")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", axis.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(axis.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", axis.IsDisplayed))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AngleDegree")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.AngleDegree))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Length")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.Length))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CenterX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.CenterX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CenterY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.CenterY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.EndX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.EndY))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(axis.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(axis.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(axis.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(axis.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(axis.Transform))
		initializerStatements += setValueField

	}

	map_AxisGrid_Identifiers := make(map[*AxisGrid]string)
	_ = map_AxisGrid_Identifiers

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
	for idx, axisgrid := range axisgridOrdered {

		id = generatesIdentifier("AxisGrid", idx, axisgrid.Name)
		map_AxisGrid_Identifiers[axisgrid] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AxisGrid")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", axisgrid.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(axisgrid.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", axisgrid.IsDisplayed))
		initializerStatements += setValueField

	}

	map_Bezier_Identifiers := make(map[*Bezier]string)
	_ = map_Bezier_Identifiers

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
	for idx, bezier := range bezierOrdered {

		id = generatesIdentifier("Bezier", idx, bezier.Name)
		map_Bezier_Identifiers[bezier] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bezier")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bezier.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bezier.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", bezier.IsDisplayed))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.StartX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.StartY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ControlPointStartX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.ControlPointStartX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ControlPointStartY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.ControlPointStartY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.EndX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.EndY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ControlPointEndX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.ControlPointEndX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ControlPointEndY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.ControlPointEndY))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bezier.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bezier.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bezier.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bezier.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bezier.Transform))
		initializerStatements += setValueField

	}

	map_BezierGrid_Identifiers := make(map[*BezierGrid]string)
	_ = map_BezierGrid_Identifiers

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
	for idx, beziergrid := range beziergridOrdered {

		id = generatesIdentifier("BezierGrid", idx, beziergrid.Name)
		map_BezierGrid_Identifiers[beziergrid] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BezierGrid")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", beziergrid.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(beziergrid.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", beziergrid.IsDisplayed))
		initializerStatements += setValueField

	}

	map_BezierGridStack_Identifiers := make(map[*BezierGridStack]string)
	_ = map_BezierGridStack_Identifiers

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
	for idx, beziergridstack := range beziergridstackOrdered {

		id = generatesIdentifier("BezierGridStack", idx, beziergridstack.Name)
		map_BezierGridStack_Identifiers[beziergridstack] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BezierGridStack")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", beziergridstack.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(beziergridstack.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", beziergridstack.IsDisplayed))
		initializerStatements += setValueField

	}

	map_Chapter_Identifiers := make(map[*Chapter]string)
	_ = map_Chapter_Identifiers

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
	for idx, chapter := range chapterOrdered {

		id = generatesIdentifier("Chapter", idx, chapter.Name)
		map_Chapter_Identifiers[chapter] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Chapter")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", chapter.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(chapter.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MardownContent")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(chapter.MardownContent))
		initializerStatements += setValueField

	}

	map_Circle_Identifiers := make(map[*Circle]string)
	_ = map_Circle_Identifiers

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
	for idx, circle := range circleOrdered {

		id = generatesIdentifier("Circle", idx, circle.Name)
		map_Circle_Identifiers[circle] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Circle")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", circle.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", circle.IsDisplayed))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CenterX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.CenterX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CenterY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.CenterY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasBespokeRadius")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", circle.HasBespokeRadius))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BespopkeRadius")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.BespopkeRadius))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.Transform))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Pitch")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", circle.Pitch))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", circle.ShowName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BeatNb")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", circle.BeatNb))
		initializerStatements += setValueField

	}

	map_CircleGrid_Identifiers := make(map[*CircleGrid]string)
	_ = map_CircleGrid_Identifiers

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
	for idx, circlegrid := range circlegridOrdered {

		id = generatesIdentifier("CircleGrid", idx, circlegrid.Name)
		map_CircleGrid_Identifiers[circlegrid] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CircleGrid")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", circlegrid.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circlegrid.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", circlegrid.IsDisplayed))
		initializerStatements += setValueField

	}

	map_Content_Identifiers := make(map[*Content]string)
	_ = map_Content_Identifiers

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
	for idx, content := range contentOrdered {

		id = generatesIdentifier("Content", idx, content.Name)
		map_Content_Identifiers[content] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Content")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", content.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MardownContent")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.MardownContent))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ContentPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.ContentPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OutputPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.OutputPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LayoutPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.LayoutPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StaticPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.StaticPath))
		initializerStatements += setValueField

		if content.Target != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Target")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+content.Target.ToCodeString())
			initializerStatements += setValueField
		}

	}

	map_ExportToMusicxml_Identifiers := make(map[*ExportToMusicxml]string)
	_ = map_ExportToMusicxml_Identifiers

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
	for idx, exporttomusicxml := range exporttomusicxmlOrdered {

		id = generatesIdentifier("ExportToMusicxml", idx, exporttomusicxml.Name)
		map_ExportToMusicxml_Identifiers[exporttomusicxml] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ExportToMusicxml")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", exporttomusicxml.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(exporttomusicxml.Name))
		initializerStatements += setValueField

	}

	map_FrontCurve_Identifiers := make(map[*FrontCurve]string)
	_ = map_FrontCurve_Identifiers

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
	for idx, frontcurve := range frontcurveOrdered {

		id = generatesIdentifier("FrontCurve", idx, frontcurve.Name)
		map_FrontCurve_Identifiers[frontcurve] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FrontCurve")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", frontcurve.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(frontcurve.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Path")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(frontcurve.Path))
		initializerStatements += setValueField

	}

	map_FrontCurveStack_Identifiers := make(map[*FrontCurveStack]string)
	_ = map_FrontCurveStack_Identifiers

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
	for idx, frontcurvestack := range frontcurvestackOrdered {

		id = generatesIdentifier("FrontCurveStack", idx, frontcurvestack.Name)
		map_FrontCurveStack_Identifiers[frontcurvestack] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FrontCurveStack")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", frontcurvestack.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(frontcurvestack.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", frontcurvestack.IsDisplayed))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(frontcurvestack.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", frontcurvestack.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(frontcurvestack.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", frontcurvestack.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", frontcurvestack.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(frontcurvestack.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(frontcurvestack.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(frontcurvestack.Transform))
		initializerStatements += setValueField

	}

	map_HorizontalAxis_Identifiers := make(map[*HorizontalAxis]string)
	_ = map_HorizontalAxis_Identifiers

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
	for idx, horizontalaxis := range horizontalaxisOrdered {

		id = generatesIdentifier("HorizontalAxis", idx, horizontalaxis.Name)
		map_HorizontalAxis_Identifiers[horizontalaxis] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "HorizontalAxis")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", horizontalaxis.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(horizontalaxis.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", horizontalaxis.IsDisplayed))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AxisHandleBorderLength")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", horizontalaxis.AxisHandleBorderLength))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Axis_Length")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", horizontalaxis.Axis_Length))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(horizontalaxis.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", horizontalaxis.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(horizontalaxis.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", horizontalaxis.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", horizontalaxis.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(horizontalaxis.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(horizontalaxis.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(horizontalaxis.Transform))
		initializerStatements += setValueField

	}

	map_Key_Identifiers := make(map[*Key]string)
	_ = map_Key_Identifiers

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
	for idx, key := range keyOrdered {

		id = generatesIdentifier("Key", idx, key.Name)
		map_Key_Identifiers[key] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Key")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", key.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(key.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", key.IsDisplayed))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Path")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(key.Path))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(key.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", key.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(key.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", key.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", key.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(key.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(key.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(key.Transform))
		initializerStatements += setValueField

	}

	map_Parameter_Identifiers := make(map[*Parameter]string)
	_ = map_Parameter_Identifiers

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
	for idx, parameter := range parameterOrdered {

		id = generatesIdentifier("Parameter", idx, parameter.Name)
		map_Parameter_Identifiers[parameter] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Parameter")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", parameter.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(parameter.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BackendColor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(parameter.BackendColor))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinuteColor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(parameter.MinuteColor))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HourColor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(parameter.HourColor))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "N")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.N))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "M")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.M))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Z")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.Z))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InsideAngle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.InsideAngle))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShiftToNearestCircle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.ShiftToNearestCircle))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SideLength")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.SideLength))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.StackWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbShitRight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.NbShitRight))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.StackHeight))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BezierControlLengthRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.BezierControlLengthRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SpiralBezierStrength")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.SpiralBezierStrength))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbInterpolationPoints")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.NbInterpolationPoints))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FkeySizeRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.FkeySizeRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FkeyOriginRelativeX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.FkeyOriginRelativeX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FkeyOriginRelativeY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.FkeyOriginRelativeY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PitchHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.PitchHeight))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbPitchLines")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.NbPitchLines))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BeatLinesHeightRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.BeatLinesHeightRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbBeatLines")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.NbBeatLines))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbOfBeatsInTheme")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.NbOfBeatsInTheme))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FirstVoiceShiftX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.FirstVoiceShiftX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FirstVoiceShiftY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.FirstVoiceShiftY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PitchDifference")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.PitchDifference))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BeatsPerSecond")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.BeatsPerSecond))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Level")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.Level))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsMinor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", parameter.IsMinor))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ThemeBinaryEncoding")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.ThemeBinaryEncoding))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OriginX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.OriginX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OriginY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.OriginY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SpiralOriginX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.SpiralOriginX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SpiralOriginY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.SpiralOriginY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OriginCrossWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.OriginCrossWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SpiralRadiusRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.SpiralRadiusRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowSpiralBezierConstruct")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", parameter.ShowSpiralBezierConstruct))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowInterpolationPoints")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", parameter.ShowInterpolationPoints))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ActualBeatsTemporalShift")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", parameter.ActualBeatsTemporalShift))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PathToStaticFiles")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(parameter.PathToStaticFiles))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PathToGeneratedSVG")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(parameter.PathToGeneratedSVG))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PathToGeneratedScore")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(parameter.PathToGeneratedScore))
		initializerStatements += setValueField

	}

	map_Rhombus_Identifiers := make(map[*Rhombus]string)
	_ = map_Rhombus_Identifiers

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
	for idx, rhombus := range rhombusOrdered {

		id = generatesIdentifier("Rhombus", idx, rhombus.Name)
		map_Rhombus_Identifiers[rhombus] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rhombus")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rhombus.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rhombus.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rhombus.IsDisplayed))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CenterX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.CenterX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CenterY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.CenterY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SideLength")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.SideLength))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AngleDegree")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.AngleDegree))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InsideAngle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.InsideAngle))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rhombus.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rhombus.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rhombus.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rhombus.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rhombus.Transform))
		initializerStatements += setValueField

	}

	map_RhombusGrid_Identifiers := make(map[*RhombusGrid]string)
	_ = map_RhombusGrid_Identifiers

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
	for idx, rhombusgrid := range rhombusgridOrdered {

		id = generatesIdentifier("RhombusGrid", idx, rhombusgrid.Name)
		map_RhombusGrid_Identifiers[rhombusgrid] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RhombusGrid")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rhombusgrid.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rhombusgrid.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rhombusgrid.IsDisplayed))
		initializerStatements += setValueField

	}

	map_ShapeCategory_Identifiers := make(map[*ShapeCategory]string)
	_ = map_ShapeCategory_Identifiers

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
	for idx, shapecategory := range shapecategoryOrdered {

		id = generatesIdentifier("ShapeCategory", idx, shapecategory.Name)
		map_ShapeCategory_Identifiers[shapecategory] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShapeCategory")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", shapecategory.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(shapecategory.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", shapecategory.IsExpanded))
		initializerStatements += setValueField

	}

	map_SpiralBezier_Identifiers := make(map[*SpiralBezier]string)
	_ = map_SpiralBezier_Identifiers

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
	for idx, spiralbezier := range spiralbezierOrdered {

		id = generatesIdentifier("SpiralBezier", idx, spiralbezier.Name)
		map_SpiralBezier_Identifiers[spiralbezier] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralBezier")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralbezier.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralbezier.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralbezier.IsDisplayed))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.StartX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.StartY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ControlPointStartX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.ControlPointStartX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ControlPointStartY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.ControlPointStartY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.EndX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.EndY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ControlPointEndX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.ControlPointEndX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ControlPointEndY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.ControlPointEndY))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralbezier.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralbezier.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralbezier.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralbezier.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralbezier.Transform))
		initializerStatements += setValueField

	}

	map_SpiralBezierGrid_Identifiers := make(map[*SpiralBezierGrid]string)
	_ = map_SpiralBezierGrid_Identifiers

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
	for idx, spiralbeziergrid := range spiralbeziergridOrdered {

		id = generatesIdentifier("SpiralBezierGrid", idx, spiralbeziergrid.Name)
		map_SpiralBezierGrid_Identifiers[spiralbeziergrid] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralBezierGrid")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralbeziergrid.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralbeziergrid.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralbeziergrid.IsDisplayed))
		initializerStatements += setValueField

	}

	map_SpiralCircle_Identifiers := make(map[*SpiralCircle]string)
	_ = map_SpiralCircle_Identifiers

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
	for idx, spiralcircle := range spiralcircleOrdered {

		id = generatesIdentifier("SpiralCircle", idx, spiralcircle.Name)
		map_SpiralCircle_Identifiers[spiralcircle] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralCircle")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralcircle.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralcircle.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralcircle.IsDisplayed))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CenterX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.CenterX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CenterY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.CenterY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasBespokeRadius")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralcircle.HasBespokeRadius))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BespopkeRadius")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.BespopkeRadius))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralcircle.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralcircle.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralcircle.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralcircle.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralcircle.Transform))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Pitch")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", spiralcircle.Pitch))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralcircle.ShowName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BeatNb")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", spiralcircle.BeatNb))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Path")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralcircle.Path))
		initializerStatements += setValueField

	}

	map_SpiralCircleGrid_Identifiers := make(map[*SpiralCircleGrid]string)
	_ = map_SpiralCircleGrid_Identifiers

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
	for idx, spiralcirclegrid := range spiralcirclegridOrdered {

		id = generatesIdentifier("SpiralCircleGrid", idx, spiralcirclegrid.Name)
		map_SpiralCircleGrid_Identifiers[spiralcirclegrid] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralCircleGrid")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralcirclegrid.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralcirclegrid.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralcirclegrid.IsDisplayed))
		initializerStatements += setValueField

	}

	map_SpiralLine_Identifiers := make(map[*SpiralLine]string)
	_ = map_SpiralLine_Identifiers

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
	for idx, spiralline := range spirallineOrdered {

		id = generatesIdentifier("SpiralLine", idx, spiralline.Name)
		map_SpiralLine_Identifiers[spiralline] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralLine")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralline.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralline.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralline.IsDisplayed))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.StartX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.EndX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.StartY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.EndY))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralline.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralline.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralline.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralline.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralline.Transform))
		initializerStatements += setValueField

	}

	map_SpiralLineGrid_Identifiers := make(map[*SpiralLineGrid]string)
	_ = map_SpiralLineGrid_Identifiers

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
	for idx, spirallinegrid := range spirallinegridOrdered {

		id = generatesIdentifier("SpiralLineGrid", idx, spirallinegrid.Name)
		map_SpiralLineGrid_Identifiers[spirallinegrid] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralLineGrid")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spirallinegrid.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spirallinegrid.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spirallinegrid.IsDisplayed))
		initializerStatements += setValueField

	}

	map_SpiralOrigin_Identifiers := make(map[*SpiralOrigin]string)
	_ = map_SpiralOrigin_Identifiers

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
	for idx, spiralorigin := range spiraloriginOrdered {

		id = generatesIdentifier("SpiralOrigin", idx, spiralorigin.Name)
		map_SpiralOrigin_Identifiers[spiralorigin] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralOrigin")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralorigin.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralorigin.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralorigin.IsDisplayed))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralorigin.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralorigin.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralorigin.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralorigin.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralorigin.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralorigin.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralorigin.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralorigin.Transform))
		initializerStatements += setValueField

	}

	map_SpiralRhombus_Identifiers := make(map[*SpiralRhombus]string)
	_ = map_SpiralRhombus_Identifiers

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
	for idx, spiralrhombus := range spiralrhombusOrdered {

		id = generatesIdentifier("SpiralRhombus", idx, spiralrhombus.Name)
		map_SpiralRhombus_Identifiers[spiralrhombus] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralRhombus")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralrhombus.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralrhombus.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralrhombus.IsDisplayed))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X_r0")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.X_r0))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y_r0")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.Y_r0))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X_r1")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.X_r1))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y_r1")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.Y_r1))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X_r2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.X_r2))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y_r2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.Y_r2))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X_r3")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.X_r3))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y_r3")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.Y_r3))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralrhombus.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralrhombus.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralrhombus.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralrhombus.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralrhombus.Transform))
		initializerStatements += setValueField

	}

	map_SpiralRhombusGrid_Identifiers := make(map[*SpiralRhombusGrid]string)
	_ = map_SpiralRhombusGrid_Identifiers

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
	for idx, spiralrhombusgrid := range spiralrhombusgridOrdered {

		id = generatesIdentifier("SpiralRhombusGrid", idx, spiralrhombusgrid.Name)
		map_SpiralRhombusGrid_Identifiers[spiralrhombusgrid] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SpiralRhombusGrid")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spiralrhombusgrid.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spiralrhombusgrid.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralrhombusgrid.IsDisplayed))
		initializerStatements += setValueField

	}

	map_VerticalAxis_Identifiers := make(map[*VerticalAxis]string)
	_ = map_VerticalAxis_Identifiers

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
	for idx, verticalaxis := range verticalaxisOrdered {

		id = generatesIdentifier("VerticalAxis", idx, verticalaxis.Name)
		map_VerticalAxis_Identifiers[verticalaxis] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "VerticalAxis")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", verticalaxis.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(verticalaxis.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", verticalaxis.IsDisplayed))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AxisHandleBorderLength")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", verticalaxis.AxisHandleBorderLength))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Axis_Length")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", verticalaxis.Axis_Length))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(verticalaxis.Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", verticalaxis.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(verticalaxis.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", verticalaxis.StrokeOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", verticalaxis.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(verticalaxis.StrokeDashArray))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(verticalaxis.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(verticalaxis.Transform))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(axisOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Axis instances pointers"
	}
	for idx, axis := range axisOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Axis", idx, axis.Name)
		map_Axis_Identifiers[axis] = id

		// Initialisation of values
		if axis.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[axis.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(axisgridOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AxisGrid instances pointers"
	}
	for idx, axisgrid := range axisgridOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AxisGrid", idx, axisgrid.Name)
		map_AxisGrid_Identifiers[axisgrid] = id

		// Initialisation of values
		if axisgrid.Reference != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Reference")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Axis_Identifiers[axisgrid.Reference])
			pointersInitializesStatements += setPointerField
		}

		if axisgrid.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[axisgrid.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

		for _, _axis := range axisgrid.Axiss {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Axiss")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Axis_Identifiers[_axis])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(bezierOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Bezier instances pointers"
	}
	for idx, bezier := range bezierOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Bezier", idx, bezier.Name)
		map_Bezier_Identifiers[bezier] = id

		// Initialisation of values
		if bezier.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[bezier.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(beziergridOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of BezierGrid instances pointers"
	}
	for idx, beziergrid := range beziergridOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("BezierGrid", idx, beziergrid.Name)
		map_BezierGrid_Identifiers[beziergrid] = id

		// Initialisation of values
		if beziergrid.Reference != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Reference")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bezier_Identifiers[beziergrid.Reference])
			pointersInitializesStatements += setPointerField
		}

		if beziergrid.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[beziergrid.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

		for _, _bezier := range beziergrid.Beziers {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Beziers")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bezier_Identifiers[_bezier])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(beziergridstackOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of BezierGridStack instances pointers"
	}
	for idx, beziergridstack := range beziergridstackOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("BezierGridStack", idx, beziergridstack.Name)
		map_BezierGridStack_Identifiers[beziergridstack] = id

		// Initialisation of values
		if beziergridstack.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[beziergridstack.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

		for _, _beziergrid := range beziergridstack.BezierGrids {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "BezierGrids")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BezierGrid_Identifiers[_beziergrid])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(chapterOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Chapter instances pointers"
	}
	for idx, chapter := range chapterOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Chapter", idx, chapter.Name)
		map_Chapter_Identifiers[chapter] = id

		// Initialisation of values
	}

	if len(circleOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Circle instances pointers"
	}
	for idx, circle := range circleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Circle", idx, circle.Name)
		map_Circle_Identifiers[circle] = id

		// Initialisation of values
		if circle.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[circle.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(circlegridOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of CircleGrid instances pointers"
	}
	for idx, circlegrid := range circlegridOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("CircleGrid", idx, circlegrid.Name)
		map_CircleGrid_Identifiers[circlegrid] = id

		// Initialisation of values
		if circlegrid.Reference != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Reference")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Circle_Identifiers[circlegrid.Reference])
			pointersInitializesStatements += setPointerField
		}

		if circlegrid.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[circlegrid.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

		for _, _circle := range circlegrid.Circles {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Circles")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Circle_Identifiers[_circle])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(contentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Content instances pointers"
	}
	for idx, content := range contentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Content", idx, content.Name)
		map_Content_Identifiers[content] = id

		// Initialisation of values
		for _, _chapter := range content.Chapters {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Chapters")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Chapter_Identifiers[_chapter])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(exporttomusicxmlOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ExportToMusicxml instances pointers"
	}
	for idx, exporttomusicxml := range exporttomusicxmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ExportToMusicxml", idx, exporttomusicxml.Name)
		map_ExportToMusicxml_Identifiers[exporttomusicxml] = id

		// Initialisation of values
		if exporttomusicxml.Parameter != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Parameter")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Parameter_Identifiers[exporttomusicxml.Parameter])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(frontcurveOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of FrontCurve instances pointers"
	}
	for idx, frontcurve := range frontcurveOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FrontCurve", idx, frontcurve.Name)
		map_FrontCurve_Identifiers[frontcurve] = id

		// Initialisation of values
	}

	if len(frontcurvestackOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of FrontCurveStack instances pointers"
	}
	for idx, frontcurvestack := range frontcurvestackOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FrontCurveStack", idx, frontcurvestack.Name)
		map_FrontCurveStack_Identifiers[frontcurvestack] = id

		// Initialisation of values
		if frontcurvestack.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[frontcurvestack.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

		for _, _frontcurve := range frontcurvestack.FrontCurves {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FrontCurves")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FrontCurve_Identifiers[_frontcurve])
			pointersInitializesStatements += setPointerField
		}

		for _, _spiralcircle := range frontcurvestack.SpiralCircles {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralCircles")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralCircle_Identifiers[_spiralcircle])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(horizontalaxisOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of HorizontalAxis instances pointers"
	}
	for idx, horizontalaxis := range horizontalaxisOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("HorizontalAxis", idx, horizontalaxis.Name)
		map_HorizontalAxis_Identifiers[horizontalaxis] = id

		// Initialisation of values
		if horizontalaxis.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[horizontalaxis.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(keyOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Key instances pointers"
	}
	for idx, key := range keyOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Key", idx, key.Name)
		map_Key_Identifiers[key] = id

		// Initialisation of values
		if key.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[key.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(parameterOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Parameter instances pointers"
	}
	for idx, parameter := range parameterOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Parameter", idx, parameter.Name)
		map_Parameter_Identifiers[parameter] = id

		// Initialisation of values
		if parameter.InitialRhombus != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "InitialRhombus")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Rhombus_Identifiers[parameter.InitialRhombus])
			pointersInitializesStatements += setPointerField
		}

		if parameter.InitialCircle != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "InitialCircle")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Circle_Identifiers[parameter.InitialCircle])
			pointersInitializesStatements += setPointerField
		}

		if parameter.InitialRhombusGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "InitialRhombusGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RhombusGrid_Identifiers[parameter.InitialRhombusGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.InitialCircleGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "InitialCircleGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CircleGrid_Identifiers[parameter.InitialCircleGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.InitialAxis != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "InitialAxis")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Axis_Identifiers[parameter.InitialAxis])
			pointersInitializesStatements += setPointerField
		}

		if parameter.RotatedAxis != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RotatedAxis")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Axis_Identifiers[parameter.RotatedAxis])
			pointersInitializesStatements += setPointerField
		}

		if parameter.RotatedRhombus != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RotatedRhombus")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Rhombus_Identifiers[parameter.RotatedRhombus])
			pointersInitializesStatements += setPointerField
		}

		if parameter.RotatedRhombusGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RotatedRhombusGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RhombusGrid_Identifiers[parameter.RotatedRhombusGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.RotatedCircleGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RotatedCircleGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CircleGrid_Identifiers[parameter.RotatedCircleGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.NextRhombus != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "NextRhombus")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Rhombus_Identifiers[parameter.NextRhombus])
			pointersInitializesStatements += setPointerField
		}

		if parameter.NextCircle != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "NextCircle")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Circle_Identifiers[parameter.NextCircle])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowingRhombusGridSeed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowingRhombusGridSeed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Rhombus_Identifiers[parameter.GrowingRhombusGridSeed])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowingRhombusGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowingRhombusGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RhombusGrid_Identifiers[parameter.GrowingRhombusGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowingCircleGridSeed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowingCircleGridSeed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Circle_Identifiers[parameter.GrowingCircleGridSeed])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowingCircleGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowingCircleGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CircleGrid_Identifiers[parameter.GrowingCircleGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowingCircleGridLeftSeed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowingCircleGridLeftSeed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Circle_Identifiers[parameter.GrowingCircleGridLeftSeed])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowingCircleGridLeft != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowingCircleGridLeft")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CircleGrid_Identifiers[parameter.GrowingCircleGridLeft])
			pointersInitializesStatements += setPointerField
		}

		if parameter.ConstructionAxis != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ConstructionAxis")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Axis_Identifiers[parameter.ConstructionAxis])
			pointersInitializesStatements += setPointerField
		}

		if parameter.ConstructionAxisGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ConstructionAxisGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AxisGrid_Identifiers[parameter.ConstructionAxisGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.ConstructionCircle != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ConstructionCircle")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Circle_Identifiers[parameter.ConstructionCircle])
			pointersInitializesStatements += setPointerField
		}

		if parameter.ConstructionCircleGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ConstructionCircleGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CircleGrid_Identifiers[parameter.ConstructionCircleGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowthCurveSeed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowthCurveSeed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bezier_Identifiers[parameter.GrowthCurveSeed])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowthCurve != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowthCurve")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BezierGrid_Identifiers[parameter.GrowthCurve])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowthCurveShiftedRightSeed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowthCurveShiftedRightSeed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bezier_Identifiers[parameter.GrowthCurveShiftedRightSeed])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowthCurveShiftedRight != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowthCurveShiftedRight")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BezierGrid_Identifiers[parameter.GrowthCurveShiftedRight])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowthCurveNextSeed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowthCurveNextSeed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bezier_Identifiers[parameter.GrowthCurveNextSeed])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowthCurveNext != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowthCurveNext")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BezierGrid_Identifiers[parameter.GrowthCurveNext])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowthCurveNextShiftedRightSeed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowthCurveNextShiftedRightSeed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bezier_Identifiers[parameter.GrowthCurveNextShiftedRightSeed])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowthCurveNextShiftedRight != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowthCurveNextShiftedRight")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BezierGrid_Identifiers[parameter.GrowthCurveNextShiftedRight])
			pointersInitializesStatements += setPointerField
		}

		if parameter.GrowthCurveStack != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowthCurveStack")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BezierGridStack_Identifiers[parameter.GrowthCurveStack])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralRhombusGridSeed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralRhombusGridSeed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralRhombus_Identifiers[parameter.SpiralRhombusGridSeed])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralRhombusGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralRhombusGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralRhombusGrid_Identifiers[parameter.SpiralRhombusGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralCircleSeed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralCircleSeed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralCircle_Identifiers[parameter.SpiralCircleSeed])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralCircleGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralCircleGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralCircleGrid_Identifiers[parameter.SpiralCircleGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralCircleFullGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralCircleFullGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralCircleGrid_Identifiers[parameter.SpiralCircleFullGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralConstructionOuterLineSeed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralConstructionOuterLineSeed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralLine_Identifiers[parameter.SpiralConstructionOuterLineSeed])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralConstructionInnerLineSeed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralConstructionInnerLineSeed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralLine_Identifiers[parameter.SpiralConstructionInnerLineSeed])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralConstructionOuterLineGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralConstructionOuterLineGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralLineGrid_Identifiers[parameter.SpiralConstructionOuterLineGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralConstructionInnerLineGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralConstructionInnerLineGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralLineGrid_Identifiers[parameter.SpiralConstructionInnerLineGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralConstructionCircleGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralConstructionCircleGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralCircleGrid_Identifiers[parameter.SpiralConstructionCircleGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralConstructionOuterLineFullGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralConstructionOuterLineFullGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralLineGrid_Identifiers[parameter.SpiralConstructionOuterLineFullGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralBezierSeed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralBezierSeed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralBezier_Identifiers[parameter.SpiralBezierSeed])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralBezierGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralBezierGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralBezierGrid_Identifiers[parameter.SpiralBezierGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralBezierFullGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralBezierFullGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralBezierGrid_Identifiers[parameter.SpiralBezierFullGrid])
			pointersInitializesStatements += setPointerField
		}

		if parameter.FrontCurveStack != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FrontCurveStack")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FrontCurveStack_Identifiers[parameter.FrontCurveStack])
			pointersInitializesStatements += setPointerField
		}

		if parameter.Fkey != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fkey")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Key_Identifiers[parameter.Fkey])
			pointersInitializesStatements += setPointerField
		}

		if parameter.PitchLines != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "PitchLines")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AxisGrid_Identifiers[parameter.PitchLines])
			pointersInitializesStatements += setPointerField
		}

		if parameter.BeatLines != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "BeatLines")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AxisGrid_Identifiers[parameter.BeatLines])
			pointersInitializesStatements += setPointerField
		}

		if parameter.FirstVoice != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FirstVoice")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BezierGrid_Identifiers[parameter.FirstVoice])
			pointersInitializesStatements += setPointerField
		}

		if parameter.FirstVoiceShiftedRigth != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FirstVoiceShiftedRigth")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BezierGrid_Identifiers[parameter.FirstVoiceShiftedRigth])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SecondVoice != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SecondVoice")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BezierGrid_Identifiers[parameter.SecondVoice])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SecondVoiceShiftedRight != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SecondVoiceShiftedRight")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BezierGrid_Identifiers[parameter.SecondVoiceShiftedRight])
			pointersInitializesStatements += setPointerField
		}

		if parameter.FirstVoiceNotes != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FirstVoiceNotes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CircleGrid_Identifiers[parameter.FirstVoiceNotes])
			pointersInitializesStatements += setPointerField
		}

		if parameter.FirstVoiceNotesShiftedRight != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FirstVoiceNotesShiftedRight")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CircleGrid_Identifiers[parameter.FirstVoiceNotesShiftedRight])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SecondVoiceNotes != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SecondVoiceNotes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CircleGrid_Identifiers[parameter.SecondVoiceNotes])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SecondVoiceNotesShiftedRight != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SecondVoiceNotesShiftedRight")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CircleGrid_Identifiers[parameter.SecondVoiceNotesShiftedRight])
			pointersInitializesStatements += setPointerField
		}

		if parameter.HorizontalAxis != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "HorizontalAxis")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_HorizontalAxis_Identifiers[parameter.HorizontalAxis])
			pointersInitializesStatements += setPointerField
		}

		if parameter.VerticalAxis != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VerticalAxis")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_VerticalAxis_Identifiers[parameter.VerticalAxis])
			pointersInitializesStatements += setPointerField
		}

		if parameter.SpiralOrigin != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralOrigin")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralOrigin_Identifiers[parameter.SpiralOrigin])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(rhombusOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Rhombus instances pointers"
	}
	for idx, rhombus := range rhombusOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Rhombus", idx, rhombus.Name)
		map_Rhombus_Identifiers[rhombus] = id

		// Initialisation of values
		if rhombus.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[rhombus.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(rhombusgridOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of RhombusGrid instances pointers"
	}
	for idx, rhombusgrid := range rhombusgridOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("RhombusGrid", idx, rhombusgrid.Name)
		map_RhombusGrid_Identifiers[rhombusgrid] = id

		// Initialisation of values
		if rhombusgrid.Reference != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Reference")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Rhombus_Identifiers[rhombusgrid.Reference])
			pointersInitializesStatements += setPointerField
		}

		if rhombusgrid.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[rhombusgrid.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

		for _, _rhombus := range rhombusgrid.Rhombuses {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Rhombuses")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Rhombus_Identifiers[_rhombus])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(shapecategoryOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ShapeCategory instances pointers"
	}
	for idx, shapecategory := range shapecategoryOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ShapeCategory", idx, shapecategory.Name)
		map_ShapeCategory_Identifiers[shapecategory] = id

		// Initialisation of values
	}

	if len(spiralbezierOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SpiralBezier instances pointers"
	}
	for idx, spiralbezier := range spiralbezierOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SpiralBezier", idx, spiralbezier.Name)
		map_SpiralBezier_Identifiers[spiralbezier] = id

		// Initialisation of values
		if spiralbezier.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[spiralbezier.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spiralbeziergridOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SpiralBezierGrid instances pointers"
	}
	for idx, spiralbeziergrid := range spiralbeziergridOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SpiralBezierGrid", idx, spiralbeziergrid.Name)
		map_SpiralBezierGrid_Identifiers[spiralbeziergrid] = id

		// Initialisation of values
		if spiralbeziergrid.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[spiralbeziergrid.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

		for _, _spiralbezier := range spiralbeziergrid.SpiralBeziers {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralBeziers")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralBezier_Identifiers[_spiralbezier])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spiralcircleOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SpiralCircle instances pointers"
	}
	for idx, spiralcircle := range spiralcircleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SpiralCircle", idx, spiralcircle.Name)
		map_SpiralCircle_Identifiers[spiralcircle] = id

		// Initialisation of values
		if spiralcircle.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[spiralcircle.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spiralcirclegridOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SpiralCircleGrid instances pointers"
	}
	for idx, spiralcirclegrid := range spiralcirclegridOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SpiralCircleGrid", idx, spiralcirclegrid.Name)
		map_SpiralCircleGrid_Identifiers[spiralcirclegrid] = id

		// Initialisation of values
		if spiralcirclegrid.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[spiralcirclegrid.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

		if spiralcirclegrid.SpiralRhombusGrid != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralRhombusGrid")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralRhombusGrid_Identifiers[spiralcirclegrid.SpiralRhombusGrid])
			pointersInitializesStatements += setPointerField
		}

		for _, _spiralcircle := range spiralcirclegrid.SpiralCircles {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralCircles")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralCircle_Identifiers[_spiralcircle])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spirallineOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SpiralLine instances pointers"
	}
	for idx, spiralline := range spirallineOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SpiralLine", idx, spiralline.Name)
		map_SpiralLine_Identifiers[spiralline] = id

		// Initialisation of values
		if spiralline.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[spiralline.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spirallinegridOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SpiralLineGrid instances pointers"
	}
	for idx, spirallinegrid := range spirallinegridOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SpiralLineGrid", idx, spirallinegrid.Name)
		map_SpiralLineGrid_Identifiers[spirallinegrid] = id

		// Initialisation of values
		if spirallinegrid.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[spirallinegrid.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

		for _, _spiralline := range spirallinegrid.SpiralLines {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralLines")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralLine_Identifiers[_spiralline])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spiraloriginOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SpiralOrigin instances pointers"
	}
	for idx, spiralorigin := range spiraloriginOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SpiralOrigin", idx, spiralorigin.Name)
		map_SpiralOrigin_Identifiers[spiralorigin] = id

		// Initialisation of values
		if spiralorigin.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[spiralorigin.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spiralrhombusOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SpiralRhombus instances pointers"
	}
	for idx, spiralrhombus := range spiralrhombusOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SpiralRhombus", idx, spiralrhombus.Name)
		map_SpiralRhombus_Identifiers[spiralrhombus] = id

		// Initialisation of values
		if spiralrhombus.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[spiralrhombus.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(spiralrhombusgridOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SpiralRhombusGrid instances pointers"
	}
	for idx, spiralrhombusgrid := range spiralrhombusgridOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SpiralRhombusGrid", idx, spiralrhombusgrid.Name)
		map_SpiralRhombusGrid_Identifiers[spiralrhombusgrid] = id

		// Initialisation of values
		if spiralrhombusgrid.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[spiralrhombusgrid.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

		for _, _spiralrhombus := range spiralrhombusgrid.SpiralRhombuses {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SpiralRhombuses")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SpiralRhombus_Identifiers[_spiralrhombus])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(verticalaxisOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of VerticalAxis instances pointers"
	}
	for idx, verticalaxis := range verticalaxisOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("VerticalAxis", idx, verticalaxis.Name)
		map_VerticalAxis_Identifiers[verticalaxis] = id

		// Initialisation of values
		if verticalaxis.ShapeCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ShapeCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ShapeCategory_Identifiers[verticalaxis.ShapeCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	// Local time with timezone
	localTimestamp := stage.commitTimeStamp.Format("2006-01-02 15:04:05.000000 MST")

	// UTC time
	utcTimestamp := stage.commitTimeStamp.UTC().Format("2006-01-02 15:04:05.000000 UTC")
	res = strings.ReplaceAll(res, "{{LocalTimeStamp}}", localTimestamp)
	res = strings.ReplaceAll(res, "{{UTCTimeStamp}}", utcTimestamp)
	res = strings.ReplaceAll(res, "{{CommitId}}", fmt.Sprintf("%.10d", stage.commitId))

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

	if stage.generatesDiff {
		diff := computeDiff(stage.contentWhenParsed, res)
		os.WriteFile(fmt.Sprintf("%s-%.10d-%.10d.delta", name, stage.commitIdWhenParsed, stage.commitId), []byte(diff), os.FileMode(0666))
		diff = ComputeDiff(stage.contentWhenParsed, res)
		os.WriteFile(fmt.Sprintf("%s-%.10d-%.10d.diff", name, stage.commitIdWhenParsed, stage.commitId), []byte(diff), os.FileMode(0666))
	}
	stage.contentWhenParsed = res
	stage.commitIdWhenParsed = stage.commitId

	fmt.Fprintln(file, res)
}

// computeDiff calculates the git-style unified diff between two strings.
func computeDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffToDelta(diffs)
}

// computePrettyDiff calculates the git-style unified diff between two strings.
func computePrettyDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffPrettyHtml(diffs)
}

// applyDiff reconstructs the original string 'a' from the new string 'b' and the diff string 'c'.
func applyDiff(b, c string) (string, error) {
	dmp := diffmatchpatch.New()
	diffs, err := dmp.DiffFromDelta(b, c)
	if err != nil {
		return "", err
	}
	patches := dmp.PatchMake(b, diffs)
	// We are applying the patch in reverse to get from 'b' to 'a'.
	// The library's PatchApply function returns the new string and a slice of booleans indicating the success of each patch application.
	result, _ := dmp.PatchApply(patches, b)
	return result, nil
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
