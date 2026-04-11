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
	"slices"
	"time"

	"{{ModelsPackageName}}"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time
var _ = slices.Index[[]int, int]

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
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)` /* */

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

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

// ToRawStringLiteral formats a string into safe Go source code,
// using backticks to preserve newlines and readability.
func ToRawStringLiteral(s string) string {
	// Step 1: Replace every backtick with a closing backtick,
	// a double-quoted backtick, and an opening backtick.
	escaped := strings.ReplaceAll(s, "`", "` + \"`\" + `")

	// Step 2: Wrap the entire resulting string in backticks.
	result := "`" + escaped + "`"

	// Step 3: Clean up any empty raw strings (``) at the boundaries
	// just in case your original string started or ended with a backtick.
	result = strings.ReplaceAll(result, "`` + ", "")
	result = strings.ReplaceAll(result, " + ``", "")

	return result
}

// MarshallFile marshall the stage content into a file as an instanciation into a stage
// according to the marshalling policy of the stage.
//
// In GongMarshallingAppendCommit mode, it will append the last commit to the file.
// In other modes, it will rewrite the entire file.
func (stage *Stage) MarshallFile(filename, modelsPackageName, packageName string) {

	if stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {
		contentBytes, err := os.ReadFile(filename)

		// if the file does not exist, marshall the full stage
		if os.IsNotExist(err) {
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr.Error())
			}
			defer file.Close()
			stage.Marshall(file, modelsPackageName, packageName)
			return
		}
		if err != nil {
			log.Fatal(err.Error())
		}

		content := string(contentBytes)

		if stage.isSquashing {
			// we squash: we want to clear the current function body
			// and let the append logic write the squashed commit
			firstBrace := strings.Index(content, "func _(stage *models.Stage) {")
			if firstBrace != -1 {
				firstBrace += len("func _(stage *models.Stage) {")
				content = content[:firstBrace] + "\n}"
			}
		}

		if stage.isApplyingBackwardCommit {
			// we are going backward, we need to remove the last forward commit from the file

			// because commitsBehind has been incremented before the call to this function
			// the index of the commit to remove is len(forwardCommits) - commitsBehind
			commitIndexToRemove := len(stage.forwardCommits) - stage.GetCommitsBehind()

			if commitIndexToRemove < 0 || commitIndexToRemove >= len(stage.forwardCommits) {
				return // Should not happen if history is consistent
			}

			commitToRemove := stage.forwardCommits[commitIndexToRemove]

			lastIndex := strings.LastIndex(content, commitToRemove+"\n")
			if lastIndex != -1 {
				newContent := content[:lastIndex] + content[lastIndex+len(commitToRemove)+1:]
				err = os.WriteFile(filename, []byte(newContent), 0644)
				if err != nil {
					log.Fatal(err.Error())
				}
			} else {
				lastIndex = strings.LastIndex(content, commitToRemove)
				if lastIndex != -1 {
					newContent := content[:lastIndex] + content[lastIndex+len(commitToRemove):]
					err = os.WriteFile(filename, []byte(newContent), 0644)
					if err != nil {
						log.Fatal(err.Error())
					}
				} else {
					// The commit block was not found. This typically happens for the initial
					// commit which is formatted differently (the lines after func _(stage *models.Stage) {).
					// We rewrite the entire file with the current (rewound) stage state to safely remove it.
					file, createErr := os.Create(filename)
					if createErr != nil {
						log.Fatal(createErr.Error())
					}
					defer file.Close()
					stage.Marshall(file, modelsPackageName, packageName)
				}
			}
			return // we are done for the backward case
		}

		if stage.isApplyingForwardCommit {
			// bypass the modified check
		} else if !stage.modified {
			return
		}

		forwardCommits := stage.GetForwardCommits()
		if len(forwardCommits) == 0 {
			return // nothing to do
		}

		activeCommits := len(forwardCommits) - stage.GetCommitsBehind()
		if activeCommits <= 0 {
			return
		}
		forwardCommit := forwardCommits[activeCommits-1]

		// append before the ending brace of the func
		lastBrace := strings.LastIndex(content, "}")
		if lastBrace == -1 {
			// if no brace, something is wrong with the file, so we rewrite it
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr.Error())
			}
			defer file.Close()
			stage.Marshall(file, modelsPackageName, packageName)
			return
		}

		contentBeforeBrace := content[:lastBrace]
		trimmedContentBeforeBrace := strings.TrimSpace(contentBeforeBrace)
		emptyBody := stage.isSquashing ||
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *models.Stage) {") ||
			strings.HasSuffix(trimmedContentBeforeBrace, "// insertion point for setup of pointers")

		// check if the file ends with stage.Commit() before the brace
		if !emptyBody && !strings.HasSuffix(trimmedContentBeforeBrace, "stage.Commit()") {
			contentBeforeBrace = contentBeforeBrace + "\n\tstage.Commit()\n"
		}

		// insert the commit statements before the last brace
		newContent := contentBeforeBrace + forwardCommit + "\n" + content[lastBrace:]

		err = os.WriteFile(filename, []byte(newContent), 0644)
		if err != nil {
			log.Fatal(err.Error())
		}

	} else {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		stage.Marshall(file, modelsPackageName, packageName)
	}
}

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
	var identifiersDecl strings.Builder
	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder

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
		axisi_order, oki := stage.Axis_stagedOrder[axisi]
		axisj_order, okj := stage.Axis_stagedOrder[axisj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return axisi_order < axisj_order
	})
	if len(axisOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, axis := range axisOrdered {

		identifiersDecl.WriteString(axis.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(axis.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(axis.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "AngleDegree"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "Length"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "CenterX"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "CenterY"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "Transform"))
	}

	axisgridOrdered := []*AxisGrid{}
	for axisgrid := range stage.AxisGrids {
		axisgridOrdered = append(axisgridOrdered, axisgrid)
	}
	sort.Slice(axisgridOrdered[:], func(i, j int) bool {
		axisgridi := axisgridOrdered[i]
		axisgridj := axisgridOrdered[j]
		axisgridi_order, oki := stage.AxisGrid_stagedOrder[axisgridi]
		axisgridj_order, okj := stage.AxisGrid_stagedOrder[axisgridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return axisgridi_order < axisgridj_order
	})
	if len(axisgridOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, axisgrid := range axisgridOrdered {

		identifiersDecl.WriteString(axisgrid.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(axisgrid.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(axisgrid.GongMarshallField(stage, "Reference"))
		initializerStatements.WriteString(axisgrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(axisgrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(axisgrid.GongMarshallField(stage, "Axiss"))
	}

	bezierOrdered := []*Bezier{}
	for bezier := range stage.Beziers {
		bezierOrdered = append(bezierOrdered, bezier)
	}
	sort.Slice(bezierOrdered[:], func(i, j int) bool {
		bezieri := bezierOrdered[i]
		bezierj := bezierOrdered[j]
		bezieri_order, oki := stage.Bezier_stagedOrder[bezieri]
		bezierj_order, okj := stage.Bezier_stagedOrder[bezierj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return bezieri_order < bezierj_order
	})
	if len(bezierOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, bezier := range bezierOrdered {

		identifiersDecl.WriteString(bezier.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(bezier.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "ControlPointStartX"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "ControlPointStartY"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "ControlPointEndX"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "ControlPointEndY"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "Transform"))
	}

	beziergridOrdered := []*BezierGrid{}
	for beziergrid := range stage.BezierGrids {
		beziergridOrdered = append(beziergridOrdered, beziergrid)
	}
	sort.Slice(beziergridOrdered[:], func(i, j int) bool {
		beziergridi := beziergridOrdered[i]
		beziergridj := beziergridOrdered[j]
		beziergridi_order, oki := stage.BezierGrid_stagedOrder[beziergridi]
		beziergridj_order, okj := stage.BezierGrid_stagedOrder[beziergridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return beziergridi_order < beziergridj_order
	})
	if len(beziergridOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, beziergrid := range beziergridOrdered {

		identifiersDecl.WriteString(beziergrid.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(beziergrid.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(beziergrid.GongMarshallField(stage, "Reference"))
		initializerStatements.WriteString(beziergrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(beziergrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(beziergrid.GongMarshallField(stage, "Beziers"))
	}

	beziergridstackOrdered := []*BezierGridStack{}
	for beziergridstack := range stage.BezierGridStacks {
		beziergridstackOrdered = append(beziergridstackOrdered, beziergridstack)
	}
	sort.Slice(beziergridstackOrdered[:], func(i, j int) bool {
		beziergridstacki := beziergridstackOrdered[i]
		beziergridstackj := beziergridstackOrdered[j]
		beziergridstacki_order, oki := stage.BezierGridStack_stagedOrder[beziergridstacki]
		beziergridstackj_order, okj := stage.BezierGridStack_stagedOrder[beziergridstackj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return beziergridstacki_order < beziergridstackj_order
	})
	if len(beziergridstackOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, beziergridstack := range beziergridstackOrdered {

		identifiersDecl.WriteString(beziergridstack.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(beziergridstack.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(beziergridstack.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(beziergridstack.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(beziergridstack.GongMarshallField(stage, "BezierGrids"))
	}

	chapterOrdered := []*Chapter{}
	for chapter := range stage.Chapters {
		chapterOrdered = append(chapterOrdered, chapter)
	}
	sort.Slice(chapterOrdered[:], func(i, j int) bool {
		chapteri := chapterOrdered[i]
		chapterj := chapterOrdered[j]
		chapteri_order, oki := stage.Chapter_stagedOrder[chapteri]
		chapterj_order, okj := stage.Chapter_stagedOrder[chapterj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return chapteri_order < chapterj_order
	})
	if len(chapterOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, chapter := range chapterOrdered {

		identifiersDecl.WriteString(chapter.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(chapter.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(chapter.GongMarshallField(stage, "MardownContent"))
	}

	circleOrdered := []*Circle{}
	for circle := range stage.Circles {
		circleOrdered = append(circleOrdered, circle)
	}
	sort.Slice(circleOrdered[:], func(i, j int) bool {
		circlei := circleOrdered[i]
		circlej := circleOrdered[j]
		circlei_order, oki := stage.Circle_stagedOrder[circlei]
		circlej_order, okj := stage.Circle_stagedOrder[circlej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return circlei_order < circlej_order
	})
	if len(circleOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, circle := range circleOrdered {

		identifiersDecl.WriteString(circle.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(circle.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "CenterX"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "CenterY"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "HasBespokeRadius"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "BespopkeRadius"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Transform"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Pitch"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "ShowName"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "BeatNb"))
	}

	circlegridOrdered := []*CircleGrid{}
	for circlegrid := range stage.CircleGrids {
		circlegridOrdered = append(circlegridOrdered, circlegrid)
	}
	sort.Slice(circlegridOrdered[:], func(i, j int) bool {
		circlegridi := circlegridOrdered[i]
		circlegridj := circlegridOrdered[j]
		circlegridi_order, oki := stage.CircleGrid_stagedOrder[circlegridi]
		circlegridj_order, okj := stage.CircleGrid_stagedOrder[circlegridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return circlegridi_order < circlegridj_order
	})
	if len(circlegridOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, circlegrid := range circlegridOrdered {

		identifiersDecl.WriteString(circlegrid.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(circlegrid.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(circlegrid.GongMarshallField(stage, "Reference"))
		initializerStatements.WriteString(circlegrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(circlegrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(circlegrid.GongMarshallField(stage, "Circles"))
	}

	contentOrdered := []*Content{}
	for content := range stage.Contents {
		contentOrdered = append(contentOrdered, content)
	}
	sort.Slice(contentOrdered[:], func(i, j int) bool {
		contenti := contentOrdered[i]
		contentj := contentOrdered[j]
		contenti_order, oki := stage.Content_stagedOrder[contenti]
		contentj_order, okj := stage.Content_stagedOrder[contentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return contenti_order < contentj_order
	})
	if len(contentOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, content := range contentOrdered {

		identifiersDecl.WriteString(content.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(content.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "MardownContent"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "ContentPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "OutputPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "LayoutPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "StaticPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "Target"))
		pointersInitializesStatements.WriteString(content.GongMarshallField(stage, "Chapters"))
	}

	exporttomusicxmlOrdered := []*ExportToMusicxml{}
	for exporttomusicxml := range stage.ExportToMusicxmls {
		exporttomusicxmlOrdered = append(exporttomusicxmlOrdered, exporttomusicxml)
	}
	sort.Slice(exporttomusicxmlOrdered[:], func(i, j int) bool {
		exporttomusicxmli := exporttomusicxmlOrdered[i]
		exporttomusicxmlj := exporttomusicxmlOrdered[j]
		exporttomusicxmli_order, oki := stage.ExportToMusicxml_stagedOrder[exporttomusicxmli]
		exporttomusicxmlj_order, okj := stage.ExportToMusicxml_stagedOrder[exporttomusicxmlj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return exporttomusicxmli_order < exporttomusicxmlj_order
	})
	if len(exporttomusicxmlOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, exporttomusicxml := range exporttomusicxmlOrdered {

		identifiersDecl.WriteString(exporttomusicxml.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(exporttomusicxml.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(exporttomusicxml.GongMarshallField(stage, "Parameter"))
	}

	frontcurveOrdered := []*FrontCurve{}
	for frontcurve := range stage.FrontCurves {
		frontcurveOrdered = append(frontcurveOrdered, frontcurve)
	}
	sort.Slice(frontcurveOrdered[:], func(i, j int) bool {
		frontcurvei := frontcurveOrdered[i]
		frontcurvej := frontcurveOrdered[j]
		frontcurvei_order, oki := stage.FrontCurve_stagedOrder[frontcurvei]
		frontcurvej_order, okj := stage.FrontCurve_stagedOrder[frontcurvej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return frontcurvei_order < frontcurvej_order
	})
	if len(frontcurveOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, frontcurve := range frontcurveOrdered {

		identifiersDecl.WriteString(frontcurve.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(frontcurve.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(frontcurve.GongMarshallField(stage, "Path"))
	}

	frontcurvestackOrdered := []*FrontCurveStack{}
	for frontcurvestack := range stage.FrontCurveStacks {
		frontcurvestackOrdered = append(frontcurvestackOrdered, frontcurvestack)
	}
	sort.Slice(frontcurvestackOrdered[:], func(i, j int) bool {
		frontcurvestacki := frontcurvestackOrdered[i]
		frontcurvestackj := frontcurvestackOrdered[j]
		frontcurvestacki_order, oki := stage.FrontCurveStack_stagedOrder[frontcurvestacki]
		frontcurvestackj_order, okj := stage.FrontCurveStack_stagedOrder[frontcurvestackj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return frontcurvestacki_order < frontcurvestackj_order
	})
	if len(frontcurvestackOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, frontcurvestack := range frontcurvestackOrdered {

		identifiersDecl.WriteString(frontcurvestack.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(frontcurvestack.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(frontcurvestack.GongMarshallField(stage, "FrontCurves"))
		pointersInitializesStatements.WriteString(frontcurvestack.GongMarshallField(stage, "SpiralCircles"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "Transform"))
	}

	horizontalaxisOrdered := []*HorizontalAxis{}
	for horizontalaxis := range stage.HorizontalAxiss {
		horizontalaxisOrdered = append(horizontalaxisOrdered, horizontalaxis)
	}
	sort.Slice(horizontalaxisOrdered[:], func(i, j int) bool {
		horizontalaxisi := horizontalaxisOrdered[i]
		horizontalaxisj := horizontalaxisOrdered[j]
		horizontalaxisi_order, oki := stage.HorizontalAxis_stagedOrder[horizontalaxisi]
		horizontalaxisj_order, okj := stage.HorizontalAxis_stagedOrder[horizontalaxisj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return horizontalaxisi_order < horizontalaxisj_order
	})
	if len(horizontalaxisOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, horizontalaxis := range horizontalaxisOrdered {

		identifiersDecl.WriteString(horizontalaxis.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(horizontalaxis.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "AxisHandleBorderLength"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "Axis_Length"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "Transform"))
	}

	keyOrdered := []*Key{}
	for key := range stage.Keys {
		keyOrdered = append(keyOrdered, key)
	}
	sort.Slice(keyOrdered[:], func(i, j int) bool {
		keyi := keyOrdered[i]
		keyj := keyOrdered[j]
		keyi_order, oki := stage.Key_stagedOrder[keyi]
		keyj_order, okj := stage.Key_stagedOrder[keyj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return keyi_order < keyj_order
	})
	if len(keyOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, key := range keyOrdered {

		identifiersDecl.WriteString(key.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(key.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(key.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "Path"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "Transform"))
	}

	parameterOrdered := []*Parameter{}
	for parameter := range stage.Parameters {
		parameterOrdered = append(parameterOrdered, parameter)
	}
	sort.Slice(parameterOrdered[:], func(i, j int) bool {
		parameteri := parameterOrdered[i]
		parameterj := parameterOrdered[j]
		parameteri_order, oki := stage.Parameter_stagedOrder[parameteri]
		parameterj_order, okj := stage.Parameter_stagedOrder[parameterj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return parameteri_order < parameterj_order
	})
	if len(parameterOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, parameter := range parameterOrdered {

		identifiersDecl.WriteString(parameter.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "BackendColor"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "MinuteColor"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "HourColor"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "N"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "M"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "Z"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "InsideAngle"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "ShiftToNearestCircle"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "SideLength"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "InitialRhombus"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "InitialCircle"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "InitialRhombusGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "InitialCircleGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "InitialAxis"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "RotatedAxis"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "RotatedRhombus"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "RotatedRhombusGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "RotatedCircleGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "NextRhombus"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "NextCircle"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowingRhombusGridSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowingRhombusGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowingCircleGridSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowingCircleGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowingCircleGridLeftSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowingCircleGridLeft"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "ConstructionAxis"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "ConstructionAxisGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "ConstructionCircle"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "ConstructionCircleGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurve"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveShiftedRightSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveShiftedRight"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveNextSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveNext"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveNextShiftedRightSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveNextShiftedRight"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveStack"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "StackWidth"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "NbShitRight"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "StackHeight"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "BezierControlLengthRatio"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralRhombusGridSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralRhombusGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralCircleSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralCircleGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralCircleFullGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralConstructionOuterLineSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralConstructionInnerLineSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralConstructionOuterLineGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralConstructionInnerLineGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralConstructionCircleGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralConstructionOuterLineFullGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralBezierSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralBezierGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralBezierFullGrid"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "SpiralBezierStrength"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "FrontCurveStack"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "NbInterpolationPoints"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "Fkey"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "FkeySizeRatio"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "FkeyOriginRelativeX"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "FkeyOriginRelativeY"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "PitchLines"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "PitchHeight"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "NbPitchLines"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "BeatLines"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "BeatLinesHeightRatio"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "NbBeatLines"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "NbOfBeatsInTheme"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "FirstVoice"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "FirstVoiceShiftedRigth"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "FirstVoiceShiftX"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "FirstVoiceShiftY"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SecondVoice"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SecondVoiceShiftedRight"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "PitchDifference"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "BeatsPerSecond"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "Level"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "FirstVoiceNotes"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "FirstVoiceNotesShiftedRight"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SecondVoiceNotes"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SecondVoiceNotesShiftedRight"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "IsMinor"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "ThemeBinaryEncoding"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "OriginX"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "OriginY"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "HorizontalAxis"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "VerticalAxis"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralOrigin"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "SpiralOriginX"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "SpiralOriginY"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "OriginCrossWidth"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "SpiralRadiusRatio"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "ShowSpiralBezierConstruct"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "ShowInterpolationPoints"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "ActualBeatsTemporalShift"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "PathToStaticFiles"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "PathToGeneratedSVG"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "PathToGeneratedScore"))
	}

	rhombusOrdered := []*Rhombus{}
	for rhombus := range stage.Rhombuss {
		rhombusOrdered = append(rhombusOrdered, rhombus)
	}
	sort.Slice(rhombusOrdered[:], func(i, j int) bool {
		rhombusi := rhombusOrdered[i]
		rhombusj := rhombusOrdered[j]
		rhombusi_order, oki := stage.Rhombus_stagedOrder[rhombusi]
		rhombusj_order, okj := stage.Rhombus_stagedOrder[rhombusj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return rhombusi_order < rhombusj_order
	})
	if len(rhombusOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, rhombus := range rhombusOrdered {

		identifiersDecl.WriteString(rhombus.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(rhombus.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "CenterX"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "CenterY"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "SideLength"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "AngleDegree"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "InsideAngle"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "Transform"))
	}

	rhombusgridOrdered := []*RhombusGrid{}
	for rhombusgrid := range stage.RhombusGrids {
		rhombusgridOrdered = append(rhombusgridOrdered, rhombusgrid)
	}
	sort.Slice(rhombusgridOrdered[:], func(i, j int) bool {
		rhombusgridi := rhombusgridOrdered[i]
		rhombusgridj := rhombusgridOrdered[j]
		rhombusgridi_order, oki := stage.RhombusGrid_stagedOrder[rhombusgridi]
		rhombusgridj_order, okj := stage.RhombusGrid_stagedOrder[rhombusgridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return rhombusgridi_order < rhombusgridj_order
	})
	if len(rhombusgridOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, rhombusgrid := range rhombusgridOrdered {

		identifiersDecl.WriteString(rhombusgrid.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(rhombusgrid.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(rhombusgrid.GongMarshallField(stage, "Reference"))
		initializerStatements.WriteString(rhombusgrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(rhombusgrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(rhombusgrid.GongMarshallField(stage, "Rhombuses"))
	}

	shapecategoryOrdered := []*ShapeCategory{}
	for shapecategory := range stage.ShapeCategorys {
		shapecategoryOrdered = append(shapecategoryOrdered, shapecategory)
	}
	sort.Slice(shapecategoryOrdered[:], func(i, j int) bool {
		shapecategoryi := shapecategoryOrdered[i]
		shapecategoryj := shapecategoryOrdered[j]
		shapecategoryi_order, oki := stage.ShapeCategory_stagedOrder[shapecategoryi]
		shapecategoryj_order, okj := stage.ShapeCategory_stagedOrder[shapecategoryj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return shapecategoryi_order < shapecategoryj_order
	})
	if len(shapecategoryOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, shapecategory := range shapecategoryOrdered {

		identifiersDecl.WriteString(shapecategory.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(shapecategory.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(shapecategory.GongMarshallField(stage, "IsExpanded"))
	}

	spiralbezierOrdered := []*SpiralBezier{}
	for spiralbezier := range stage.SpiralBeziers {
		spiralbezierOrdered = append(spiralbezierOrdered, spiralbezier)
	}
	sort.Slice(spiralbezierOrdered[:], func(i, j int) bool {
		spiralbezieri := spiralbezierOrdered[i]
		spiralbezierj := spiralbezierOrdered[j]
		spiralbezieri_order, oki := stage.SpiralBezier_stagedOrder[spiralbezieri]
		spiralbezierj_order, okj := stage.SpiralBezier_stagedOrder[spiralbezierj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralbezieri_order < spiralbezierj_order
	})
	if len(spiralbezierOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spiralbezier := range spiralbezierOrdered {

		identifiersDecl.WriteString(spiralbezier.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralbezier.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "ControlPointStartX"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "ControlPointStartY"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "ControlPointEndX"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "ControlPointEndY"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "Transform"))
	}

	spiralbeziergridOrdered := []*SpiralBezierGrid{}
	for spiralbeziergrid := range stage.SpiralBezierGrids {
		spiralbeziergridOrdered = append(spiralbeziergridOrdered, spiralbeziergrid)
	}
	sort.Slice(spiralbeziergridOrdered[:], func(i, j int) bool {
		spiralbeziergridi := spiralbeziergridOrdered[i]
		spiralbeziergridj := spiralbeziergridOrdered[j]
		spiralbeziergridi_order, oki := stage.SpiralBezierGrid_stagedOrder[spiralbeziergridi]
		spiralbeziergridj_order, okj := stage.SpiralBezierGrid_stagedOrder[spiralbeziergridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralbeziergridi_order < spiralbeziergridj_order
	})
	if len(spiralbeziergridOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spiralbeziergrid := range spiralbeziergridOrdered {

		identifiersDecl.WriteString(spiralbeziergrid.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralbeziergrid.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralbeziergrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralbeziergrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(spiralbeziergrid.GongMarshallField(stage, "SpiralBeziers"))
	}

	spiralcircleOrdered := []*SpiralCircle{}
	for spiralcircle := range stage.SpiralCircles {
		spiralcircleOrdered = append(spiralcircleOrdered, spiralcircle)
	}
	sort.Slice(spiralcircleOrdered[:], func(i, j int) bool {
		spiralcirclei := spiralcircleOrdered[i]
		spiralcirclej := spiralcircleOrdered[j]
		spiralcirclei_order, oki := stage.SpiralCircle_stagedOrder[spiralcirclei]
		spiralcirclej_order, okj := stage.SpiralCircle_stagedOrder[spiralcirclej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralcirclei_order < spiralcirclej_order
	})
	if len(spiralcircleOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spiralcircle := range spiralcircleOrdered {

		identifiersDecl.WriteString(spiralcircle.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralcircle.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "CenterX"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "CenterY"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "HasBespokeRadius"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "BespopkeRadius"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "Transform"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "Pitch"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "ShowName"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "BeatNb"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "Path"))
	}

	spiralcirclegridOrdered := []*SpiralCircleGrid{}
	for spiralcirclegrid := range stage.SpiralCircleGrids {
		spiralcirclegridOrdered = append(spiralcirclegridOrdered, spiralcirclegrid)
	}
	sort.Slice(spiralcirclegridOrdered[:], func(i, j int) bool {
		spiralcirclegridi := spiralcirclegridOrdered[i]
		spiralcirclegridj := spiralcirclegridOrdered[j]
		spiralcirclegridi_order, oki := stage.SpiralCircleGrid_stagedOrder[spiralcirclegridi]
		spiralcirclegridj_order, okj := stage.SpiralCircleGrid_stagedOrder[spiralcirclegridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralcirclegridi_order < spiralcirclegridj_order
	})
	if len(spiralcirclegridOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spiralcirclegrid := range spiralcirclegridOrdered {

		identifiersDecl.WriteString(spiralcirclegrid.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralcirclegrid.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralcirclegrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralcirclegrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(spiralcirclegrid.GongMarshallField(stage, "SpiralRhombusGrid"))
		pointersInitializesStatements.WriteString(spiralcirclegrid.GongMarshallField(stage, "SpiralCircles"))
	}

	spirallineOrdered := []*SpiralLine{}
	for spiralline := range stage.SpiralLines {
		spirallineOrdered = append(spirallineOrdered, spiralline)
	}
	sort.Slice(spirallineOrdered[:], func(i, j int) bool {
		spirallinei := spirallineOrdered[i]
		spirallinej := spirallineOrdered[j]
		spirallinei_order, oki := stage.SpiralLine_stagedOrder[spirallinei]
		spirallinej_order, okj := stage.SpiralLine_stagedOrder[spirallinej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spirallinei_order < spirallinej_order
	})
	if len(spirallineOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spiralline := range spirallineOrdered {

		identifiersDecl.WriteString(spiralline.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralline.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "Transform"))
	}

	spirallinegridOrdered := []*SpiralLineGrid{}
	for spirallinegrid := range stage.SpiralLineGrids {
		spirallinegridOrdered = append(spirallinegridOrdered, spirallinegrid)
	}
	sort.Slice(spirallinegridOrdered[:], func(i, j int) bool {
		spirallinegridi := spirallinegridOrdered[i]
		spirallinegridj := spirallinegridOrdered[j]
		spirallinegridi_order, oki := stage.SpiralLineGrid_stagedOrder[spirallinegridi]
		spirallinegridj_order, okj := stage.SpiralLineGrid_stagedOrder[spirallinegridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spirallinegridi_order < spirallinegridj_order
	})
	if len(spirallinegridOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spirallinegrid := range spirallinegridOrdered {

		identifiersDecl.WriteString(spirallinegrid.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spirallinegrid.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spirallinegrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spirallinegrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(spirallinegrid.GongMarshallField(stage, "SpiralLines"))
	}

	spiraloriginOrdered := []*SpiralOrigin{}
	for spiralorigin := range stage.SpiralOrigins {
		spiraloriginOrdered = append(spiraloriginOrdered, spiralorigin)
	}
	sort.Slice(spiraloriginOrdered[:], func(i, j int) bool {
		spiralorigini := spiraloriginOrdered[i]
		spiraloriginj := spiraloriginOrdered[j]
		spiralorigini_order, oki := stage.SpiralOrigin_stagedOrder[spiralorigini]
		spiraloriginj_order, okj := stage.SpiralOrigin_stagedOrder[spiraloriginj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralorigini_order < spiraloriginj_order
	})
	if len(spiraloriginOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spiralorigin := range spiraloriginOrdered {

		identifiersDecl.WriteString(spiralorigin.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralorigin.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "Transform"))
	}

	spiralrhombusOrdered := []*SpiralRhombus{}
	for spiralrhombus := range stage.SpiralRhombuss {
		spiralrhombusOrdered = append(spiralrhombusOrdered, spiralrhombus)
	}
	sort.Slice(spiralrhombusOrdered[:], func(i, j int) bool {
		spiralrhombusi := spiralrhombusOrdered[i]
		spiralrhombusj := spiralrhombusOrdered[j]
		spiralrhombusi_order, oki := stage.SpiralRhombus_stagedOrder[spiralrhombusi]
		spiralrhombusj_order, okj := stage.SpiralRhombus_stagedOrder[spiralrhombusj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralrhombusi_order < spiralrhombusj_order
	})
	if len(spiralrhombusOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spiralrhombus := range spiralrhombusOrdered {

		identifiersDecl.WriteString(spiralrhombus.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralrhombus.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "X_r0"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Y_r0"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "X_r1"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Y_r1"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "X_r2"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Y_r2"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "X_r3"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Y_r3"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Transform"))
	}

	spiralrhombusgridOrdered := []*SpiralRhombusGrid{}
	for spiralrhombusgrid := range stage.SpiralRhombusGrids {
		spiralrhombusgridOrdered = append(spiralrhombusgridOrdered, spiralrhombusgrid)
	}
	sort.Slice(spiralrhombusgridOrdered[:], func(i, j int) bool {
		spiralrhombusgridi := spiralrhombusgridOrdered[i]
		spiralrhombusgridj := spiralrhombusgridOrdered[j]
		spiralrhombusgridi_order, oki := stage.SpiralRhombusGrid_stagedOrder[spiralrhombusgridi]
		spiralrhombusgridj_order, okj := stage.SpiralRhombusGrid_stagedOrder[spiralrhombusgridj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spiralrhombusgridi_order < spiralrhombusgridj_order
	})
	if len(spiralrhombusgridOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spiralrhombusgrid := range spiralrhombusgridOrdered {

		identifiersDecl.WriteString(spiralrhombusgrid.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralrhombusgrid.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralrhombusgrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralrhombusgrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(spiralrhombusgrid.GongMarshallField(stage, "SpiralRhombuses"))
	}

	verticalaxisOrdered := []*VerticalAxis{}
	for verticalaxis := range stage.VerticalAxiss {
		verticalaxisOrdered = append(verticalaxisOrdered, verticalaxis)
	}
	sort.Slice(verticalaxisOrdered[:], func(i, j int) bool {
		verticalaxisi := verticalaxisOrdered[i]
		verticalaxisj := verticalaxisOrdered[j]
		verticalaxisi_order, oki := stage.VerticalAxis_stagedOrder[verticalaxisi]
		verticalaxisj_order, okj := stage.VerticalAxis_stagedOrder[verticalaxisj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return verticalaxisi_order < verticalaxisj_order
	})
	if len(verticalaxisOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, verticalaxis := range verticalaxisOrdered {

		identifiersDecl.WriteString(verticalaxis.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(verticalaxis.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "AxisHandleBorderLength"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "Axis_Length"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "Transform"))
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

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl.String())
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements.String())
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements.String())

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
				stage.MetaPackageImportAlias))

		var entries strings.Builder

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
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident))
			case GONG__ENUM_CAST_STRING:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident))
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier))
			case GONG__IDENTIFIER_CONST:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident))
			case GONG__STRUCT_INSTANCE:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident))
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries.String())
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(axis.Name))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(axis.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", axis.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(axis.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(axis.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(axis.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(axis.Transform))

	case "ShapeCategory":
		if axis.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", axis.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", axis.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(axisgrid.Name))
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
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", axisgrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Reference")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "ShapeCategory":
		if axisgrid.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", axisgrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", axisgrid.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", axisgrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Axiss":
		var sb strings.Builder
		for _, _axis := range axisgrid.Axiss {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", axisgrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Axiss")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _axis.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bezier.Name))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bezier.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bezier.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bezier.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bezier.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bezier.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bezier.Transform))

	case "ShapeCategory":
		if bezier.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", bezier.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", bezier.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(beziergrid.Name))
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
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", beziergrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Reference")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "ShapeCategory":
		if beziergrid.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", beziergrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", beziergrid.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", beziergrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Beziers":
		var sb strings.Builder
		for _, _bezier := range beziergrid.Beziers {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", beziergrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Beziers")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _bezier.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(beziergridstack.Name))
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
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", beziergridstack.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "BezierGrids":
		var sb strings.Builder
		for _, _beziergrid := range beziergridstack.BezierGrids {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", beziergridstack.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "BezierGrids")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _beziergrid.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(chapter.Name))
	case "MardownContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", chapter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MardownContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(chapter.MardownContent))

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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(circle.Name))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(circle.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(circle.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(circle.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(circle.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(circle.Transform))
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
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(circlegrid.Name))
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
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", circlegrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Reference")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "ShapeCategory":
		if circlegrid.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", circlegrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", circlegrid.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", circlegrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Circles":
		var sb strings.Builder
		for _, _circle := range circlegrid.Circles {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", circlegrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Circles")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _circle.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.Name))
	case "MardownContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MardownContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.MardownContent))
	case "ContentPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ContentPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.ContentPath))
	case "OutputPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OutputPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.OutputPath))
	case "LayoutPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.LayoutPath))
	case "StaticPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StaticPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.StaticPath))
	case "Target":
		if content.Target.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Target")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+content.Target.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Target")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}

	case "Chapters":
		var sb strings.Builder
		for _, _chapter := range content.Chapters {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", content.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Chapters")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _chapter.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(exporttomusicxml.Name))

	case "Parameter":
		if exporttomusicxml.Parameter != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", exporttomusicxml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Parameter")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", exporttomusicxml.Parameter.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", exporttomusicxml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Parameter")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(frontcurve.Name))
	case "Path":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurve.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Path")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(frontcurve.Path))

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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(frontcurvestack.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", frontcurvestack.IsDisplayed))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(frontcurvestack.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", frontcurvestack.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(frontcurvestack.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(frontcurvestack.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(frontcurvestack.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(frontcurvestack.Transform))

	case "ShapeCategory":
		if frontcurvestack.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", frontcurvestack.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FrontCurves":
		var sb strings.Builder
		for _, _frontcurve := range frontcurvestack.FrontCurves {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "FrontCurves")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _frontcurve.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SpiralCircles":
		var sb strings.Builder
		for _, _spiralcircle := range frontcurvestack.SpiralCircles {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", frontcurvestack.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SpiralCircles")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spiralcircle.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(horizontalaxis.Name))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(horizontalaxis.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", horizontalaxis.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(horizontalaxis.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(horizontalaxis.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(horizontalaxis.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(horizontalaxis.Transform))

	case "ShapeCategory":
		if horizontalaxis.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", horizontalaxis.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", horizontalaxis.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(key.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", key.IsDisplayed))
	case "Path":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Path")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(key.Path))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(key.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", key.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(key.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(key.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(key.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(key.Transform))

	case "ShapeCategory":
		if key.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", key.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", key.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parameter.Name))
	case "BackendColor":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BackendColor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parameter.BackendColor))
	case "MinuteColor":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MinuteColor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parameter.MinuteColor))
	case "HourColor":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HourColor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parameter.HourColor))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parameter.PathToStaticFiles))
	case "PathToGeneratedSVG":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PathToGeneratedSVG")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parameter.PathToGeneratedSVG))
	case "PathToGeneratedScore":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PathToGeneratedScore")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(parameter.PathToGeneratedScore))

	case "InitialRhombus":
		if parameter.InitialRhombus != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialRhombus")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.InitialRhombus.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialRhombus")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "InitialCircle":
		if parameter.InitialCircle != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialCircle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.InitialCircle.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialCircle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "InitialRhombusGrid":
		if parameter.InitialRhombusGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.InitialRhombusGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "InitialCircleGrid":
		if parameter.InitialCircleGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.InitialCircleGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "InitialAxis":
		if parameter.InitialAxis != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.InitialAxis.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "RotatedAxis":
		if parameter.RotatedAxis != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RotatedAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.RotatedAxis.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RotatedAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "RotatedRhombus":
		if parameter.RotatedRhombus != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RotatedRhombus")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.RotatedRhombus.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RotatedRhombus")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "RotatedRhombusGrid":
		if parameter.RotatedRhombusGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RotatedRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.RotatedRhombusGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RotatedRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "RotatedCircleGrid":
		if parameter.RotatedCircleGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RotatedCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.RotatedCircleGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RotatedCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "NextRhombus":
		if parameter.NextRhombus != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NextRhombus")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.NextRhombus.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NextRhombus")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "NextCircle":
		if parameter.NextCircle != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NextCircle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.NextCircle.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NextCircle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowingRhombusGridSeed":
		if parameter.GrowingRhombusGridSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingRhombusGridSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowingRhombusGridSeed.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingRhombusGridSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowingRhombusGrid":
		if parameter.GrowingRhombusGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowingRhombusGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowingCircleGridSeed":
		if parameter.GrowingCircleGridSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingCircleGridSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowingCircleGridSeed.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingCircleGridSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowingCircleGrid":
		if parameter.GrowingCircleGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowingCircleGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowingCircleGridLeftSeed":
		if parameter.GrowingCircleGridLeftSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingCircleGridLeftSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowingCircleGridLeftSeed.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingCircleGridLeftSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowingCircleGridLeft":
		if parameter.GrowingCircleGridLeft != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingCircleGridLeft")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowingCircleGridLeft.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowingCircleGridLeft")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "ConstructionAxis":
		if parameter.ConstructionAxis != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConstructionAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.ConstructionAxis.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConstructionAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "ConstructionAxisGrid":
		if parameter.ConstructionAxisGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConstructionAxisGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.ConstructionAxisGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConstructionAxisGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "ConstructionCircle":
		if parameter.ConstructionCircle != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConstructionCircle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.ConstructionCircle.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConstructionCircle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "ConstructionCircleGrid":
		if parameter.ConstructionCircleGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConstructionCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.ConstructionCircleGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConstructionCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowthCurveSeed":
		if parameter.GrowthCurveSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveSeed.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowthCurve":
		if parameter.GrowthCurve != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurve")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurve.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurve")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowthCurveShiftedRightSeed":
		if parameter.GrowthCurveShiftedRightSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveShiftedRightSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveShiftedRightSeed.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveShiftedRightSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowthCurveShiftedRight":
		if parameter.GrowthCurveShiftedRight != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveShiftedRight.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowthCurveNextSeed":
		if parameter.GrowthCurveNextSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveNextSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveNextSeed.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveNextSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowthCurveNext":
		if parameter.GrowthCurveNext != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveNext")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveNext.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveNext")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowthCurveNextShiftedRightSeed":
		if parameter.GrowthCurveNextShiftedRightSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveNextShiftedRightSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveNextShiftedRightSeed.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveNextShiftedRightSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowthCurveNextShiftedRight":
		if parameter.GrowthCurveNextShiftedRight != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveNextShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveNextShiftedRight.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveNextShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "GrowthCurveStack":
		if parameter.GrowthCurveStack != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveStack")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.GrowthCurveStack.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrowthCurveStack")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralRhombusGridSeed":
		if parameter.SpiralRhombusGridSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralRhombusGridSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralRhombusGridSeed.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralRhombusGridSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralRhombusGrid":
		if parameter.SpiralRhombusGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralRhombusGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralCircleSeed":
		if parameter.SpiralCircleSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralCircleSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralCircleSeed.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralCircleSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralCircleGrid":
		if parameter.SpiralCircleGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralCircleGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralCircleFullGrid":
		if parameter.SpiralCircleFullGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralCircleFullGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralCircleFullGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralCircleFullGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralConstructionOuterLineSeed":
		if parameter.SpiralConstructionOuterLineSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionOuterLineSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralConstructionOuterLineSeed.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionOuterLineSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralConstructionInnerLineSeed":
		if parameter.SpiralConstructionInnerLineSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionInnerLineSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralConstructionInnerLineSeed.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionInnerLineSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralConstructionOuterLineGrid":
		if parameter.SpiralConstructionOuterLineGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionOuterLineGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralConstructionOuterLineGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionOuterLineGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralConstructionInnerLineGrid":
		if parameter.SpiralConstructionInnerLineGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionInnerLineGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralConstructionInnerLineGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionInnerLineGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralConstructionCircleGrid":
		if parameter.SpiralConstructionCircleGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralConstructionCircleGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionCircleGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralConstructionOuterLineFullGrid":
		if parameter.SpiralConstructionOuterLineFullGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionOuterLineFullGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralConstructionOuterLineFullGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralConstructionOuterLineFullGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralBezierSeed":
		if parameter.SpiralBezierSeed != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralBezierSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralBezierSeed.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralBezierSeed")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralBezierGrid":
		if parameter.SpiralBezierGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralBezierGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralBezierGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralBezierGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralBezierFullGrid":
		if parameter.SpiralBezierFullGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralBezierFullGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralBezierFullGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralBezierFullGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FrontCurveStack":
		if parameter.FrontCurveStack != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FrontCurveStack")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.FrontCurveStack.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FrontCurveStack")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Fkey":
		if parameter.Fkey != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Fkey")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.Fkey.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Fkey")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "PitchLines":
		if parameter.PitchLines != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PitchLines")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.PitchLines.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PitchLines")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "BeatLines":
		if parameter.BeatLines != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BeatLines")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.BeatLines.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BeatLines")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FirstVoice":
		if parameter.FirstVoice != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoice")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.FirstVoice.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoice")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FirstVoiceShiftedRigth":
		if parameter.FirstVoiceShiftedRigth != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoiceShiftedRigth")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.FirstVoiceShiftedRigth.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoiceShiftedRigth")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SecondVoice":
		if parameter.SecondVoice != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondVoice")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SecondVoice.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondVoice")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SecondVoiceShiftedRight":
		if parameter.SecondVoiceShiftedRight != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondVoiceShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SecondVoiceShiftedRight.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondVoiceShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FirstVoiceNotes":
		if parameter.FirstVoiceNotes != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoiceNotes")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.FirstVoiceNotes.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoiceNotes")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FirstVoiceNotesShiftedRight":
		if parameter.FirstVoiceNotesShiftedRight != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoiceNotesShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.FirstVoiceNotesShiftedRight.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FirstVoiceNotesShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SecondVoiceNotes":
		if parameter.SecondVoiceNotes != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondVoiceNotes")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SecondVoiceNotes.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondVoiceNotes")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SecondVoiceNotesShiftedRight":
		if parameter.SecondVoiceNotesShiftedRight != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondVoiceNotesShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SecondVoiceNotesShiftedRight.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondVoiceNotesShiftedRight")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "HorizontalAxis":
		if parameter.HorizontalAxis != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HorizontalAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.HorizontalAxis.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HorizontalAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "VerticalAxis":
		if parameter.VerticalAxis != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "VerticalAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.VerticalAxis.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "VerticalAxis")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralOrigin":
		if parameter.SpiralOrigin != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralOrigin")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", parameter.SpiralOrigin.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", parameter.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralOrigin")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rhombus.Name))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rhombus.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rhombus.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rhombus.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rhombus.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rhombus.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rhombus.Transform))

	case "ShapeCategory":
		if rhombus.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", rhombus.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rhombus.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rhombusgrid.Name))
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
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rhombusgrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Reference")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "ShapeCategory":
		if rhombusgrid.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rhombusgrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", rhombusgrid.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rhombusgrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Rhombuses":
		var sb strings.Builder
		for _, _rhombus := range rhombusgrid.Rhombuses {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", rhombusgrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Rhombuses")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _rhombus.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(shapecategory.Name))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralbezier.Name))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralbezier.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralbezier.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralbezier.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralbezier.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralbezier.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralbezier.Transform))

	case "ShapeCategory":
		if spiralbezier.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralbezier.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralbezier.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralbeziergrid.Name))
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
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralbeziergrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralBeziers":
		var sb strings.Builder
		for _, _spiralbezier := range spiralbeziergrid.SpiralBeziers {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", spiralbeziergrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SpiralBeziers")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spiralbezier.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralcircle.Name))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralcircle.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralcircle.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralcircle.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralcircle.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralcircle.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralcircle.Transform))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralcircle.Path))

	case "ShapeCategory":
		if spiralcircle.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralcircle.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralcircle.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralcirclegrid.Name))
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
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralcirclegrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralRhombusGrid":
		if spiralcirclegrid.SpiralRhombusGrid != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralcirclegrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralcirclegrid.SpiralRhombusGrid.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralcirclegrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpiralRhombusGrid")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralCircles":
		var sb strings.Builder
		for _, _spiralcircle := range spiralcirclegrid.SpiralCircles {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", spiralcirclegrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SpiralCircles")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spiralcircle.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralline.Name))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralline.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralline.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralline.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralline.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralline.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralline.Transform))

	case "ShapeCategory":
		if spiralline.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralline.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralline.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spirallinegrid.Name))
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
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spirallinegrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralLines":
		var sb strings.Builder
		for _, _spiralline := range spirallinegrid.SpiralLines {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", spirallinegrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SpiralLines")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spiralline.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralorigin.Name))
	case "IsDisplayed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisplayed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spiralorigin.IsDisplayed))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralorigin.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralorigin.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralorigin.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralorigin.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralorigin.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralorigin.Transform))

	case "ShapeCategory":
		if spiralorigin.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralorigin.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralorigin.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralrhombus.Name))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralrhombus.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", spiralrhombus.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralrhombus.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralrhombus.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralrhombus.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralrhombus.Transform))

	case "ShapeCategory":
		if spiralrhombus.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spiralrhombus.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombus.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spiralrhombusgrid.Name))
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
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spiralrhombusgrid.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SpiralRhombuses":
		var sb strings.Builder
		for _, _spiralrhombus := range spiralrhombusgrid.SpiralRhombuses {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", spiralrhombusgrid.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SpiralRhombuses")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spiralrhombus.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(verticalaxis.Name))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(verticalaxis.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", verticalaxis.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(verticalaxis.Stroke))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(verticalaxis.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(verticalaxis.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(verticalaxis.Transform))

	case "ShapeCategory":
		if verticalaxis.ShapeCategory != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", verticalaxis.ShapeCategory.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", verticalaxis.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShapeCategory")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct VerticalAxis", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (axis *Axis) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(axis.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(axis.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "AngleDegree"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "Length"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "CenterX"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "CenterY"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(axis.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (axisgrid *AxisGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(axisgrid.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(axisgrid.GongMarshallField(stage, "Reference"))
		initializerStatements.WriteString(axisgrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(axisgrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(axisgrid.GongMarshallField(stage, "Axiss"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (bezier *Bezier) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(bezier.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "ControlPointStartX"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "ControlPointStartY"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "ControlPointEndX"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "ControlPointEndY"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(bezier.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (beziergrid *BezierGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(beziergrid.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(beziergrid.GongMarshallField(stage, "Reference"))
		initializerStatements.WriteString(beziergrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(beziergrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(beziergrid.GongMarshallField(stage, "Beziers"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (beziergridstack *BezierGridStack) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(beziergridstack.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(beziergridstack.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(beziergridstack.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(beziergridstack.GongMarshallField(stage, "BezierGrids"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (chapter *Chapter) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(chapter.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(chapter.GongMarshallField(stage, "MardownContent"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circle *Circle) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(circle.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "CenterX"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "CenterY"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "HasBespokeRadius"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "BespopkeRadius"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Transform"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Pitch"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "ShowName"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "BeatNb"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circlegrid *CircleGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circlegrid.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(circlegrid.GongMarshallField(stage, "Reference"))
		initializerStatements.WriteString(circlegrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(circlegrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(circlegrid.GongMarshallField(stage, "Circles"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (content *Content) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(content.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "MardownContent"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "ContentPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "OutputPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "LayoutPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "StaticPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "Target"))
		pointersInitializesStatements.WriteString(content.GongMarshallField(stage, "Chapters"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (exporttomusicxml *ExportToMusicxml) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(exporttomusicxml.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(exporttomusicxml.GongMarshallField(stage, "Parameter"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (frontcurve *FrontCurve) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(frontcurve.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(frontcurve.GongMarshallField(stage, "Path"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (frontcurvestack *FrontCurveStack) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(frontcurvestack.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(frontcurvestack.GongMarshallField(stage, "FrontCurves"))
		pointersInitializesStatements.WriteString(frontcurvestack.GongMarshallField(stage, "SpiralCircles"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(frontcurvestack.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (horizontalaxis *HorizontalAxis) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(horizontalaxis.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "AxisHandleBorderLength"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "Axis_Length"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(horizontalaxis.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (key *Key) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(key.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(key.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "Path"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(key.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (parameter *Parameter) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "BackendColor"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "MinuteColor"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "HourColor"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "N"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "M"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "Z"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "InsideAngle"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "ShiftToNearestCircle"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "SideLength"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "InitialRhombus"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "InitialCircle"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "InitialRhombusGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "InitialCircleGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "InitialAxis"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "RotatedAxis"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "RotatedRhombus"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "RotatedRhombusGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "RotatedCircleGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "NextRhombus"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "NextCircle"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowingRhombusGridSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowingRhombusGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowingCircleGridSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowingCircleGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowingCircleGridLeftSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowingCircleGridLeft"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "ConstructionAxis"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "ConstructionAxisGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "ConstructionCircle"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "ConstructionCircleGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurve"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveShiftedRightSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveShiftedRight"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveNextSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveNext"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveNextShiftedRightSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveNextShiftedRight"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GrowthCurveStack"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "StackWidth"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "NbShitRight"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "StackHeight"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "BezierControlLengthRatio"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralRhombusGridSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralRhombusGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralCircleSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralCircleGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralCircleFullGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralConstructionOuterLineSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralConstructionInnerLineSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralConstructionOuterLineGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralConstructionInnerLineGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralConstructionCircleGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralConstructionOuterLineFullGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralBezierSeed"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralBezierGrid"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralBezierFullGrid"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "SpiralBezierStrength"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "FrontCurveStack"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "NbInterpolationPoints"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "Fkey"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "FkeySizeRatio"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "FkeyOriginRelativeX"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "FkeyOriginRelativeY"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "PitchLines"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "PitchHeight"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "NbPitchLines"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "BeatLines"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "BeatLinesHeightRatio"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "NbBeatLines"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "NbOfBeatsInTheme"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "FirstVoice"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "FirstVoiceShiftedRigth"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "FirstVoiceShiftX"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "FirstVoiceShiftY"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SecondVoice"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SecondVoiceShiftedRight"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "PitchDifference"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "BeatsPerSecond"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "Level"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "FirstVoiceNotes"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "FirstVoiceNotesShiftedRight"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SecondVoiceNotes"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SecondVoiceNotesShiftedRight"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "IsMinor"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "ThemeBinaryEncoding"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "OriginX"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "OriginY"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "HorizontalAxis"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "VerticalAxis"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "SpiralOrigin"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "SpiralOriginX"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "SpiralOriginY"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "OriginCrossWidth"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "SpiralRadiusRatio"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "ShowSpiralBezierConstruct"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "ShowInterpolationPoints"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "ActualBeatsTemporalShift"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "PathToStaticFiles"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "PathToGeneratedSVG"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "PathToGeneratedScore"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rhombus *Rhombus) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(rhombus.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "CenterX"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "CenterY"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "SideLength"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "AngleDegree"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "InsideAngle"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(rhombus.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rhombusgrid *RhombusGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rhombusgrid.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(rhombusgrid.GongMarshallField(stage, "Reference"))
		initializerStatements.WriteString(rhombusgrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(rhombusgrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(rhombusgrid.GongMarshallField(stage, "Rhombuses"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shapecategory *ShapeCategory) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shapecategory.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(shapecategory.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spiralbezier *SpiralBezier) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralbezier.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "ControlPointStartX"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "ControlPointStartY"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "ControlPointEndX"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "ControlPointEndY"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(spiralbezier.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spiralbeziergrid *SpiralBezierGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralbeziergrid.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralbeziergrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralbeziergrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(spiralbeziergrid.GongMarshallField(stage, "SpiralBeziers"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spiralcircle *SpiralCircle) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralcircle.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "CenterX"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "CenterY"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "HasBespokeRadius"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "BespopkeRadius"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "Transform"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "Pitch"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "ShowName"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "BeatNb"))
		initializerStatements.WriteString(spiralcircle.GongMarshallField(stage, "Path"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spiralcirclegrid *SpiralCircleGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralcirclegrid.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralcirclegrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralcirclegrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(spiralcirclegrid.GongMarshallField(stage, "SpiralRhombusGrid"))
		pointersInitializesStatements.WriteString(spiralcirclegrid.GongMarshallField(stage, "SpiralCircles"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spiralline *SpiralLine) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralline.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(spiralline.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spirallinegrid *SpiralLineGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spirallinegrid.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spirallinegrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spirallinegrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(spirallinegrid.GongMarshallField(stage, "SpiralLines"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spiralorigin *SpiralOrigin) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralorigin.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(spiralorigin.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spiralrhombus *SpiralRhombus) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralrhombus.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "X_r0"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Y_r0"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "X_r1"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Y_r1"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "X_r2"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Y_r2"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "X_r3"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Y_r3"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(spiralrhombus.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spiralrhombusgrid *SpiralRhombusGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spiralrhombusgrid.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spiralrhombusgrid.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(spiralrhombusgrid.GongMarshallField(stage, "ShapeCategory"))
		pointersInitializesStatements.WriteString(spiralrhombusgrid.GongMarshallField(stage, "SpiralRhombuses"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (verticalaxis *VerticalAxis) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "IsDisplayed"))
		pointersInitializesStatements.WriteString(verticalaxis.GongMarshallField(stage, "ShapeCategory"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "AxisHandleBorderLength"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "Axis_Length"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(verticalaxis.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
