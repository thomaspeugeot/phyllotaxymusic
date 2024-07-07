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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.StrokeWidth))
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "N")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", circlegrid.N))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "M")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", circlegrid.M))
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", horizontalaxis.StrokeWidth))
		initializerStatements += setValueField

	}

	map_InitialAxis_Identifiers := make(map[*InitialAxis]string)
	_ = map_InitialAxis_Identifiers

	initialaxisOrdered := []*InitialAxis{}
	for initialaxis := range stage.InitialAxiss {
		initialaxisOrdered = append(initialaxisOrdered, initialaxis)
	}
	sort.Slice(initialaxisOrdered[:], func(i, j int) bool {
		return initialaxisOrdered[i].Name < initialaxisOrdered[j].Name
	})
	if len(initialaxisOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, initialaxis := range initialaxisOrdered {

		id = generatesIdentifier("InitialAxis", idx, initialaxis.Name)
		map_InitialAxis_Identifiers[initialaxis] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "InitialAxis")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", initialaxis.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(initialaxis.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDisplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", initialaxis.IsDisplayed))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Angle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", initialaxis.Angle))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Length")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", initialaxis.Length))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", initialaxis.StrokeWidth))
		initializerStatements += setValueField

	}

	map_Line_Identifiers := make(map[*Line]string)
	_ = map_Line_Identifiers

	lineOrdered := []*Line{}
	for line := range stage.Lines {
		lineOrdered = append(lineOrdered, line)
	}
	sort.Slice(lineOrdered[:], func(i, j int) bool {
		return lineOrdered[i].Name < lineOrdered[j].Name
	})
	if len(lineOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, line := range lineOrdered {

		id = generatesIdentifier("Line", idx, line.Name)
		map_Line_Identifiers[line] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Line")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", line.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(line.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X1")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.X1))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y1")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.Y1))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.X2))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.Y2))
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.StrokeWidth))
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "N")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", rhombusgrid.N))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "M")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", rhombusgrid.M))
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", verticalaxis.StrokeWidth))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
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

	for idx, initialaxis := range initialaxisOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("InitialAxis", idx, initialaxis.Name)
		map_InitialAxis_Identifiers[initialaxis] = id

		// Initialisation of values
	}

	for idx, line := range lineOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Line", idx, line.Name)
		map_Line_Identifiers[line] = id

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
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_InitialAxis_Identifiers[parameter.InitialAxis])
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

		res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
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
