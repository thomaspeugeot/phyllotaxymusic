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

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

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
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
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
		return axisOrdered[i].Name < axisOrdered[j].Name
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Angle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.Angle))
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
		return axisgridOrdered[i].Name < axisgridOrdered[j].Name
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
		return bezierOrdered[i].Name < bezierOrdered[j].Name
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
		return beziergridOrdered[i].Name < beziergridOrdered[j].Name
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

	map_Circle_Identifiers := make(map[*Circle]string)
	_ = map_Circle_Identifiers

	circleOrdered := []*Circle{}
	for circle := range stage.Circles {
		circleOrdered = append(circleOrdered, circle)
	}
	sort.Slice(circleOrdered[:], func(i, j int) bool {
		return circleOrdered[i].Name < circleOrdered[j].Name
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

	}

	map_CircleGrid_Identifiers := make(map[*CircleGrid]string)
	_ = map_CircleGrid_Identifiers

	circlegridOrdered := []*CircleGrid{}
	for circlegrid := range stage.CircleGrids {
		circlegridOrdered = append(circlegridOrdered, circlegrid)
	}
	sort.Slice(circlegridOrdered[:], func(i, j int) bool {
		return circlegridOrdered[i].Name < circlegridOrdered[j].Name
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

	map_HorizontalAxis_Identifiers := make(map[*HorizontalAxis]string)
	_ = map_HorizontalAxis_Identifiers

	horizontalaxisOrdered := []*HorizontalAxis{}
	for horizontalaxis := range stage.HorizontalAxiss {
		horizontalaxisOrdered = append(horizontalaxisOrdered, horizontalaxis)
	}
	sort.Slice(horizontalaxisOrdered[:], func(i, j int) bool {
		return horizontalaxisOrdered[i].Name < horizontalaxisOrdered[j].Name
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

	map_Parameter_Identifiers := make(map[*Parameter]string)
	_ = map_Parameter_Identifiers

	parameterOrdered := []*Parameter{}
	for parameter := range stage.Parameters {
		parameterOrdered = append(parameterOrdered, parameter)
	}
	sort.Slice(parameterOrdered[:], func(i, j int) bool {
		return parameterOrdered[i].Name < parameterOrdered[j].Name
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SideLength")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.SideLength))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BezierControlLengthRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", parameter.BezierControlLengthRatio))
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

	}

	map_Rhombus_Identifiers := make(map[*Rhombus]string)
	_ = map_Rhombus_Identifiers

	rhombusOrdered := []*Rhombus{}
	for rhombus := range stage.Rhombuss {
		rhombusOrdered = append(rhombusOrdered, rhombus)
	}
	sort.Slice(rhombusOrdered[:], func(i, j int) bool {
		return rhombusOrdered[i].Name < rhombusOrdered[j].Name
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Angle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.Angle))
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
		return rhombusgridOrdered[i].Name < rhombusgridOrdered[j].Name
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

	map_VerticalAxis_Identifiers := make(map[*VerticalAxis]string)
	_ = map_VerticalAxis_Identifiers

	verticalaxisOrdered := []*VerticalAxis{}
	for verticalaxis := range stage.VerticalAxiss {
		verticalaxisOrdered = append(verticalaxisOrdered, verticalaxis)
	}
	sort.Slice(verticalaxisOrdered[:], func(i, j int) bool {
		return verticalaxisOrdered[i].Name < verticalaxisOrdered[j].Name
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
	for idx, axis := range axisOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Axis", idx, axis.Name)
		map_Axis_Identifiers[axis] = id

		// Initialisation of values
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

		for _, _axis := range axisgrid.Axiss {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Axiss")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Axis_Identifiers[_axis])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, bezier := range bezierOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Bezier", idx, bezier.Name)
		map_Bezier_Identifiers[bezier] = id

		// Initialisation of values
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

		for _, _bezier := range beziergrid.Beziers {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Beziers")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bezier_Identifiers[_bezier])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, circle := range circleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Circle", idx, circle.Name)
		map_Circle_Identifiers[circle] = id

		// Initialisation of values
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

		for _, _circle := range circlegrid.Circles {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Circles")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Circle_Identifiers[_circle])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, horizontalaxis := range horizontalaxisOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("HorizontalAxis", idx, horizontalaxis.Name)
		map_HorizontalAxis_Identifiers[horizontalaxis] = id

		// Initialisation of values
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

		if parameter.GrowthCurveSegment != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GrowthCurveSegment")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bezier_Identifiers[parameter.GrowthCurveSegment])
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

	}

	for idx, rhombus := range rhombusOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Rhombus", idx, rhombus.Name)
		map_Rhombus_Identifiers[rhombus] = id

		// Initialisation of values
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

		for _, _rhombus := range rhombusgrid.Rhombuses {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Rhombuses")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Rhombus_Identifiers[_rhombus])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, verticalaxis := range verticalaxisOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("VerticalAxis", idx, verticalaxis.Name)
		map_VerticalAxis_Identifiers[verticalaxis] = id

		// Initialisation of values
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.StageStruct",
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

	fmt.Fprintln(file, res)
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
