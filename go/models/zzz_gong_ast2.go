// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

var _time__dummyDeclaration2 time.Duration
var _ = _time__dummyDeclaration2

// swagger:ignore
type GONG__ExpressionType string

const (
	GONG__STRUCT_INSTANCE      GONG__ExpressionType = "STRUCT_INSTANCE"
	GONG__FIELD_OR_CONST_VALUE GONG__ExpressionType = "FIELD_OR_CONST_VALUE"
	GONG__FIELD_VALUE          GONG__ExpressionType = "FIELD_VALUE"
	GONG__ENUM_CAST_INT        GONG__ExpressionType = "ENUM_CAST_INT"
	GONG__ENUM_CAST_STRING     GONG__ExpressionType = "ENUM_CAST_STRING"
	GONG__IDENTIFIER_CONST     GONG__ExpressionType = "IDENTIFIER_CONST"
)

// ------------------------------------------------------------------------------------------------
// STATIC AST PARSING LOGIC
// ------------------------------------------------------------------------------------------------

// ModelUnmarshaller abstracts the logic for setting fields on a staged instance
type ModelUnmarshaller interface {
	// Initialize creates the struct, stages it, and returns the pointer as 'any'
	Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error)

	// UnmarshallField sets a field's value based on the AST expression
	UnmarshallField(stage *Stage, instance GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error
}

func (stage *Stage) UnmarshallFile(pathToFile string, preserveOrder bool) error {
	return ParseAstFile(stage, pathToFile, preserveOrder)
}

// ParseAstFile Parse pathToFile and stages all instances declared in the file
func ParseAstFile(stage *Stage, pathToFile string, preserveOrder bool) error {
	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
func ParseAstEmbeddedFile(stage *Stage, directory embed.FS, pathToFile string) error {
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		return errors.New(stage.GetName() + "; Unable to read embedded file " + err.Error())
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset, false)
}

// GongParseAstString parses the Go source code from a string
func GongParseAstString(stage *Stage, blob string, preserveOrder bool) error {
	fileString := "package main\nfunc _() {\n" + blob + "\n}"
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", fileString, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances using the Unmarshaller registry
func ParseAstFileFromAst(stage *Stage, inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {

	// 1. Remove Global Variables: Use a local map to track variable names to instances
	identifierMap := make(map[string]GongstructIF)

	for _, instance := range stage.GetInstances() {
		identifierMap[instance.GongGetIdentifier(stage)] = instance
	}

	// 2. Visitor Pattern: Traverse the AST
	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var typeName string
					var instanceName string

					// Inspect RHS to find the Struct Type and Name
					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								typeName = selExpr.Sel.Name
								// Attempt to find Name field in literal
								for _, elt := range compLit.Elts {
									if kv, ok := elt.(*ast.KeyValueExpr); ok {
										if k, ok := kv.Key.(*ast.Ident); ok && k.Name == "Name" {
											if v, ok := kv.Value.(*ast.BasicLit); ok {
												instanceName = strings.Trim(v.Value, "\"`")
											}
										}
									}
								}
								return false
							}
						}
						return true
					})

					// Dispatch to specific Unmarshaller
					if typeName != "" {
						if unmarshaller, exists := stage.GongUnmarshallers[typeName]; exists {
							instance, err := unmarshaller.Initialize(stage, ident.Name, instanceName, preserveOrder)
							if err == nil {
								identifierMap[ident.Name] = instance
							}
						}
					}
				}
				return false
			}

			// CASE 2: Assignment ( = )
			if node.Tok == token.ASSIGN {
				if selExpr, ok := node.Lhs[0].(*ast.SelectorExpr); ok {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if instance, exists := identifierMap[ident.Name]; exists {
							typeName := instance.GongGetGongstructName()
							if unmarshaller, exists := stage.GongUnmarshallers[typeName]; exists {
								// 3. Strategy Pattern: Delegate to Handler
								unmarshaller.UnmarshallField(stage, instance, selExpr.Sel.Name, node.Rhs[0], identifierMap)
							}
						}
					}
				}
			}
		case *ast.ExprStmt:
			if call, ok := node.X.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					if sel.Sel.Name == "Unstage" {
						if ident, ok := sel.X.(*ast.Ident); ok {
							if instance, ok := identifierMap[ident.Name]; ok {
								instance.UnstageVoid(stage)
							}
						}
					}
					if sel.Sel.Name == "Commit" {
						if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "stage" {
							if stage.IsInDeltaMode() && stage.navigationMode != GongNavigationModeNavigating {
								stage.Commit()
							} else {
								stage.ComputeInstancesNb()
								stage.ComputeReferenceAndOrders()
								if stage.OnInitCommitCallback != nil {
									stage.OnInitCommitCallback.BeforeCommit(stage)
								}
								if stage.OnInitCommitFromBackCallback != nil {
									stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
								}
								// 1. Run all Before Commit hooks
								for _, hook := range stage.beforeCommitHooks {
									hook(stage)
								}

								// 2. Run all After Commit hooks
								for _, hook := range stage.afterCommitHooks {
									hook(stage)
								}
							}
						}
					}
				}
			}
		}
		return true
	})

	return nil
}

// --- Generic Helpers for Unmarshallers ---

func GongExtractString(expr ast.Expr) string {
	switch e := expr.(type) {

	// Case 1: It's a standard string literal
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			if unquoted, err := strconv.Unquote(e.Value); err == nil {
				return unquoted
			}
		}

	// Case 2: It's a concatenated string
	case *ast.BinaryExpr:
		// We only care if they are being added together
		if e.Op == token.ADD {
			// Recursively extract the left and right sides
			left := GongExtractString(e.X)
			right := GongExtractString(e.Y)

			// Join them back together
			return left + right
		}
	}

	return ""
}

func GongExtractInt(expr ast.Expr) int {
	if bl, ok := expr.(*ast.BasicLit); ok {
		val, _ := strconv.Atoi(bl.Value)
		return val
	}
	if ue, ok := expr.(*ast.UnaryExpr); ok && ue.Op == token.SUB {
		if bl, ok := ue.X.(*ast.BasicLit); ok {
			val, _ := strconv.Atoi(bl.Value)
			return -val
		}
	}
	return 0
}

// ExtractMiddleUint takes a formatted string and returns the extracted integer.
func ExtractMiddleUint(input string) (uint, error) {
	// Compile the Regex Pattern
	re := regexp.MustCompile(`__.*?__(\d+)_.*`)

	// Find the matches
	matches := re.FindStringSubmatch(input)

	// Validate that we found the pattern and the capture group
	if len(matches) < 2 {
		return 0, fmt.Errorf("pattern not found in string: %s", input)
	}

	// matches[0] is the whole string, matches[1] is the capture group (\d+)
	numberStr := matches[1]

	// Convert string to integer (handles leading zeros automatically)
	result, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert %s to int: %v", numberStr, err)
	}

	return uint(result), nil
}

func GongExtractFloat(expr ast.Expr) float64 {
	if bl, ok := expr.(*ast.BasicLit); ok {
		val, _ := strconv.ParseFloat(bl.Value, 64)
		return val
	}
	if ue, ok := expr.(*ast.UnaryExpr); ok && ue.Op == token.SUB {
		if bl, ok := ue.X.(*ast.BasicLit); ok {
			val, _ := strconv.ParseFloat(bl.Value, 64)
			return -val
		}
	}
	return 0.0
}

func GongExtractBool(expr ast.Expr) bool {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name == "true"
	}
	return false
}

func GongExtractExpr(expr ast.Expr) any {
	switch v := expr.(type) {
	case *ast.BasicLit:
		return v.Value
	case *ast.CompositeLit:
		// Reconstruct "Package.Struct{}"
		if sel, ok := v.Type.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok {
				return id.Name + "." + sel.Sel.Name + "{}"
			}
		}
	case *ast.SelectorExpr:
		// Reconstruct "Package.Struct{}.Field"
		// X is likely a CompositeLit (Package.Struct{})
		if cl, ok := v.X.(*ast.CompositeLit); ok {
			if sel, ok := cl.Type.(*ast.SelectorExpr); ok {
				if id, ok := sel.X.(*ast.Ident); ok {
					return id.Name + "." + sel.Sel.Name + "{}." + v.Sel.Name
				}
			}
		}
		// Reconstruct "Package.Identifier"
		if id, ok := v.X.(*ast.Ident); ok {
			return id.Name + "." + v.Sel.Name
		}
	case *ast.CallExpr:
		// Reconstruct "new(Package.Struct)"
		if fun, ok := v.Fun.(*ast.Ident); ok && fun.Name == "new" {
			if len(v.Args) == 1 {
				if sel, ok := v.Args[0].(*ast.SelectorExpr); ok {
					if id, ok := sel.X.(*ast.Ident); ok {
						return "new(" + id.Name + "." + sel.Sel.Name + ")"
					}
				}
			}
		}
	}
	return ""
}

// GongUnmarshallSliceOfPointers handles append, slices.Delete, and slices.Insert for slice fields
func GongUnmarshallSliceOfPointers[T PointerToGongstruct](
	slice *[]T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) (err error) {

	if call, ok := valueExpr.(*ast.CallExpr); ok {
		funcName := ""
		var isSlices bool

		if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "slices" {
				isSlices = true
				funcName = sel.Sel.Name
			}
		} else if ident, ok := call.Fun.(*ast.Ident); ok {
			funcName = ident.Name
		}

		if isSlices {
			if funcName == "Delete" && len(call.Args) == 3 {
				start := GongExtractInt(call.Args[1])
				end := GongExtractInt(call.Args[2])
				if end > len(*slice) {
					// Handle error: log warning, resize slice, or return error
					// For example:
					return fmt.Errorf("index out of bounds: %d for len %d", end, len(*slice))
				}
				*slice = slices.Delete(*slice, start, end)
			} else if funcName == "Insert" && len(call.Args) == 3 {
				index := GongExtractInt(call.Args[1])
				if index > len(*slice) {
					return fmt.Errorf("index out of bounds: %d for len %d", index, len(*slice))
				}
				if ident, ok := call.Args[2].(*ast.Ident); ok {
					if val, ok := identifierMap[ident.Name]; ok {
						*slice = slices.Insert(*slice, index, val.(T))
					} else {
						log.Println("Ast2 Insert Unkown identifier", ident.Name)
					}
				}
			}
		} else if funcName == "append" {
			if len(call.Args) >= 2 {
				if ident, ok := call.Args[len(call.Args)-1].(*ast.Ident); ok {
					if val, ok := identifierMap[ident.Name]; ok {
						*slice = append(*slice, val.(T))
					} else {
						log.Println("Ast2 append Unkown identifier", ident.Name)
					}
				}
			}
		}
	}
	return nil
}

// GongUnmarshallPointer handles assignment of a single pointer field
func GongUnmarshallPointer[T PointerToGongstruct](
	ptr *T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) {

	if ident, ok := valueExpr.(*ast.Ident); ok {
		if ident.Name == "nil" {
			var zero T
			*ptr = zero
			return
		}
		if val, ok := identifierMap[ident.Name]; ok {
			*ptr = val.(T)
		}
	}
}

// GongUnmarshallEnum handles assignment of enum fields (via SelectorExpr or String fallback)
func GongUnmarshallEnum[T interface{ FromCodeString(string) error }](
	ptr T,
	valueExpr ast.Expr) {

	// Case 1: Standard Enum usage (models.EnumType_Value)
	if sel, ok := valueExpr.(*ast.SelectorExpr); ok {
		if err := ptr.FromCodeString(sel.Sel.Name); err != nil {
			log.Printf("UnmarshallEnum: Error parsing code string '%s': %v", sel.Sel.Name, err)
		}
	} else {
		// Case 2: Fallback to string literal
		valStr := GongExtractString(valueExpr)
		if valStr != "" {
			if err := ptr.FromCodeString(valStr); err != nil {
				log.Printf("UnmarshallEnum: Error parsing string literal '%s': %v", valStr, err)
			}
		}
	}
}

// insertion point per named struct
type AxisUnmarshaller struct{}

func (u *AxisUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Axis)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *AxisUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Axis)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "AngleDegree":
		instance.AngleDegree = GongExtractFloat(valueExpr)
	case "Length":
		instance.Length = GongExtractFloat(valueExpr)
	case "CenterX":
		instance.CenterX = GongExtractFloat(valueExpr)
	case "CenterY":
		instance.CenterY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type AxisGridUnmarshaller struct{}

func (u *AxisGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(AxisGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *AxisGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AxisGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Reference":
		GongUnmarshallPointer(&instance.Reference, valueExpr, identifierMap)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "Axiss":
		GongUnmarshallSliceOfPointers(&instance.Axiss, valueExpr, identifierMap)
	}
	return nil
}

type BezierUnmarshaller struct{}

func (u *BezierUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Bezier)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *BezierUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bezier)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "ControlPointStartX":
		instance.ControlPointStartX = GongExtractFloat(valueExpr)
	case "ControlPointStartY":
		instance.ControlPointStartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "ControlPointEndX":
		instance.ControlPointEndX = GongExtractFloat(valueExpr)
	case "ControlPointEndY":
		instance.ControlPointEndY = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type BezierGridUnmarshaller struct{}

func (u *BezierGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(BezierGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *BezierGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BezierGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Reference":
		GongUnmarshallPointer(&instance.Reference, valueExpr, identifierMap)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "Beziers":
		GongUnmarshallSliceOfPointers(&instance.Beziers, valueExpr, identifierMap)
	}
	return nil
}

type BezierGridStackUnmarshaller struct{}

func (u *BezierGridStackUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(BezierGridStack)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *BezierGridStackUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BezierGridStack)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "BezierGrids":
		GongUnmarshallSliceOfPointers(&instance.BezierGrids, valueExpr, identifierMap)
	}
	return nil
}

type ChapterUnmarshaller struct{}

func (u *ChapterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Chapter)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ChapterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Chapter)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MardownContent":
		instance.MardownContent = GongExtractString(valueExpr)
	}
	return nil
}

type CircleUnmarshaller struct{}

func (u *CircleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Circle)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *CircleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Circle)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "CenterX":
		instance.CenterX = GongExtractFloat(valueExpr)
	case "CenterY":
		instance.CenterY = GongExtractFloat(valueExpr)
	case "HasBespokeRadius":
		instance.HasBespokeRadius = GongExtractBool(valueExpr)
	case "BespopkeRadius":
		instance.BespopkeRadius = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Pitch":
		instance.Pitch = GongExtractInt(valueExpr)
	case "ShowName":
		instance.ShowName = GongExtractBool(valueExpr)
	case "BeatNb":
		instance.BeatNb = GongExtractInt(valueExpr)
	}
	return nil
}

type CircleGridUnmarshaller struct{}

func (u *CircleGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(CircleGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *CircleGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CircleGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Reference":
		GongUnmarshallPointer(&instance.Reference, valueExpr, identifierMap)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "Circles":
		GongUnmarshallSliceOfPointers(&instance.Circles, valueExpr, identifierMap)
	}
	return nil
}

type ContentUnmarshaller struct{}

func (u *ContentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Content)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ContentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Content)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MardownContent":
		instance.MardownContent = GongExtractString(valueExpr)
	case "ContentPath":
		instance.ContentPath = GongExtractString(valueExpr)
	case "OutputPath":
		instance.OutputPath = GongExtractString(valueExpr)
	case "LayoutPath":
		instance.LayoutPath = GongExtractString(valueExpr)
	case "StaticPath":
		instance.StaticPath = GongExtractString(valueExpr)
	case "Target":
		GongUnmarshallEnum(&instance.Target, valueExpr)
	case "Chapters":
		GongUnmarshallSliceOfPointers(&instance.Chapters, valueExpr, identifierMap)
	}
	return nil
}

type ExportToMusicxmlUnmarshaller struct{}

func (u *ExportToMusicxmlUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ExportToMusicxml)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ExportToMusicxmlUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ExportToMusicxml)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Parameter":
		GongUnmarshallPointer(&instance.Parameter, valueExpr, identifierMap)
	}
	return nil
}

type FrontCurveUnmarshaller struct{}

func (u *FrontCurveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(FrontCurve)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *FrontCurveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FrontCurve)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Path":
		instance.Path = GongExtractString(valueExpr)
	}
	return nil
}

type FrontCurveStackUnmarshaller struct{}

func (u *FrontCurveStackUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(FrontCurveStack)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *FrontCurveStackUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FrontCurveStack)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "FrontCurves":
		GongUnmarshallSliceOfPointers(&instance.FrontCurves, valueExpr, identifierMap)
	case "SpiralCircles":
		GongUnmarshallSliceOfPointers(&instance.SpiralCircles, valueExpr, identifierMap)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type HorizontalAxisUnmarshaller struct{}

func (u *HorizontalAxisUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(HorizontalAxis)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *HorizontalAxisUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HorizontalAxis)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "AxisHandleBorderLength":
		instance.AxisHandleBorderLength = GongExtractFloat(valueExpr)
	case "Axis_Length":
		instance.Axis_Length = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type KeyUnmarshaller struct{}

func (u *KeyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Key)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *KeyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Key)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "Path":
		instance.Path = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type ParameterUnmarshaller struct{}

func (u *ParameterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Parameter)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ParameterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Parameter)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BackendColor":
		instance.BackendColor = GongExtractString(valueExpr)
	case "MinuteColor":
		instance.MinuteColor = GongExtractString(valueExpr)
	case "HourColor":
		instance.HourColor = GongExtractString(valueExpr)
	case "N":
		instance.N = GongExtractInt(valueExpr)
	case "M":
		instance.M = GongExtractInt(valueExpr)
	case "Z":
		instance.Z = GongExtractInt(valueExpr)
	case "InsideAngle":
		instance.InsideAngle = GongExtractFloat(valueExpr)
	case "ShiftToNearestCircle":
		instance.ShiftToNearestCircle = GongExtractInt(valueExpr)
	case "SideLength":
		instance.SideLength = GongExtractFloat(valueExpr)
	case "InitialRhombus":
		GongUnmarshallPointer(&instance.InitialRhombus, valueExpr, identifierMap)
	case "InitialCircle":
		GongUnmarshallPointer(&instance.InitialCircle, valueExpr, identifierMap)
	case "InitialRhombusGrid":
		GongUnmarshallPointer(&instance.InitialRhombusGrid, valueExpr, identifierMap)
	case "InitialCircleGrid":
		GongUnmarshallPointer(&instance.InitialCircleGrid, valueExpr, identifierMap)
	case "InitialAxis":
		GongUnmarshallPointer(&instance.InitialAxis, valueExpr, identifierMap)
	case "RotatedAxis":
		GongUnmarshallPointer(&instance.RotatedAxis, valueExpr, identifierMap)
	case "RotatedRhombus":
		GongUnmarshallPointer(&instance.RotatedRhombus, valueExpr, identifierMap)
	case "RotatedRhombusGrid":
		GongUnmarshallPointer(&instance.RotatedRhombusGrid, valueExpr, identifierMap)
	case "RotatedCircleGrid":
		GongUnmarshallPointer(&instance.RotatedCircleGrid, valueExpr, identifierMap)
	case "NextRhombus":
		GongUnmarshallPointer(&instance.NextRhombus, valueExpr, identifierMap)
	case "NextCircle":
		GongUnmarshallPointer(&instance.NextCircle, valueExpr, identifierMap)
	case "GrowingRhombusGridSeed":
		GongUnmarshallPointer(&instance.GrowingRhombusGridSeed, valueExpr, identifierMap)
	case "GrowingRhombusGrid":
		GongUnmarshallPointer(&instance.GrowingRhombusGrid, valueExpr, identifierMap)
	case "GrowingCircleGridSeed":
		GongUnmarshallPointer(&instance.GrowingCircleGridSeed, valueExpr, identifierMap)
	case "GrowingCircleGrid":
		GongUnmarshallPointer(&instance.GrowingCircleGrid, valueExpr, identifierMap)
	case "GrowingCircleGridLeftSeed":
		GongUnmarshallPointer(&instance.GrowingCircleGridLeftSeed, valueExpr, identifierMap)
	case "GrowingCircleGridLeft":
		GongUnmarshallPointer(&instance.GrowingCircleGridLeft, valueExpr, identifierMap)
	case "ConstructionAxis":
		GongUnmarshallPointer(&instance.ConstructionAxis, valueExpr, identifierMap)
	case "ConstructionAxisGrid":
		GongUnmarshallPointer(&instance.ConstructionAxisGrid, valueExpr, identifierMap)
	case "ConstructionCircle":
		GongUnmarshallPointer(&instance.ConstructionCircle, valueExpr, identifierMap)
	case "ConstructionCircleGrid":
		GongUnmarshallPointer(&instance.ConstructionCircleGrid, valueExpr, identifierMap)
	case "GrowthCurveSeed":
		GongUnmarshallPointer(&instance.GrowthCurveSeed, valueExpr, identifierMap)
	case "GrowthCurve":
		GongUnmarshallPointer(&instance.GrowthCurve, valueExpr, identifierMap)
	case "GrowthCurveShiftedRightSeed":
		GongUnmarshallPointer(&instance.GrowthCurveShiftedRightSeed, valueExpr, identifierMap)
	case "GrowthCurveShiftedRight":
		GongUnmarshallPointer(&instance.GrowthCurveShiftedRight, valueExpr, identifierMap)
	case "GrowthCurveNextSeed":
		GongUnmarshallPointer(&instance.GrowthCurveNextSeed, valueExpr, identifierMap)
	case "GrowthCurveNext":
		GongUnmarshallPointer(&instance.GrowthCurveNext, valueExpr, identifierMap)
	case "GrowthCurveNextShiftedRightSeed":
		GongUnmarshallPointer(&instance.GrowthCurveNextShiftedRightSeed, valueExpr, identifierMap)
	case "GrowthCurveNextShiftedRight":
		GongUnmarshallPointer(&instance.GrowthCurveNextShiftedRight, valueExpr, identifierMap)
	case "GrowthCurveStack":
		GongUnmarshallPointer(&instance.GrowthCurveStack, valueExpr, identifierMap)
	case "StackWidth":
		instance.StackWidth = GongExtractInt(valueExpr)
	case "NbShitRight":
		instance.NbShitRight = GongExtractInt(valueExpr)
	case "StackHeight":
		instance.StackHeight = GongExtractInt(valueExpr)
	case "BezierControlLengthRatio":
		instance.BezierControlLengthRatio = GongExtractFloat(valueExpr)
	case "SpiralRhombusGridSeed":
		GongUnmarshallPointer(&instance.SpiralRhombusGridSeed, valueExpr, identifierMap)
	case "SpiralRhombusGrid":
		GongUnmarshallPointer(&instance.SpiralRhombusGrid, valueExpr, identifierMap)
	case "SpiralCircleSeed":
		GongUnmarshallPointer(&instance.SpiralCircleSeed, valueExpr, identifierMap)
	case "SpiralCircleGrid":
		GongUnmarshallPointer(&instance.SpiralCircleGrid, valueExpr, identifierMap)
	case "SpiralCircleFullGrid":
		GongUnmarshallPointer(&instance.SpiralCircleFullGrid, valueExpr, identifierMap)
	case "SpiralConstructionOuterLineSeed":
		GongUnmarshallPointer(&instance.SpiralConstructionOuterLineSeed, valueExpr, identifierMap)
	case "SpiralConstructionInnerLineSeed":
		GongUnmarshallPointer(&instance.SpiralConstructionInnerLineSeed, valueExpr, identifierMap)
	case "SpiralConstructionOuterLineGrid":
		GongUnmarshallPointer(&instance.SpiralConstructionOuterLineGrid, valueExpr, identifierMap)
	case "SpiralConstructionInnerLineGrid":
		GongUnmarshallPointer(&instance.SpiralConstructionInnerLineGrid, valueExpr, identifierMap)
	case "SpiralConstructionCircleGrid":
		GongUnmarshallPointer(&instance.SpiralConstructionCircleGrid, valueExpr, identifierMap)
	case "SpiralConstructionOuterLineFullGrid":
		GongUnmarshallPointer(&instance.SpiralConstructionOuterLineFullGrid, valueExpr, identifierMap)
	case "SpiralBezierSeed":
		GongUnmarshallPointer(&instance.SpiralBezierSeed, valueExpr, identifierMap)
	case "SpiralBezierGrid":
		GongUnmarshallPointer(&instance.SpiralBezierGrid, valueExpr, identifierMap)
	case "SpiralBezierFullGrid":
		GongUnmarshallPointer(&instance.SpiralBezierFullGrid, valueExpr, identifierMap)
	case "SpiralBezierStrength":
		instance.SpiralBezierStrength = GongExtractFloat(valueExpr)
	case "FrontCurveStack":
		GongUnmarshallPointer(&instance.FrontCurveStack, valueExpr, identifierMap)
	case "NbInterpolationPoints":
		instance.NbInterpolationPoints = GongExtractInt(valueExpr)
	case "Fkey":
		GongUnmarshallPointer(&instance.Fkey, valueExpr, identifierMap)
	case "FkeySizeRatio":
		instance.FkeySizeRatio = GongExtractFloat(valueExpr)
	case "FkeyOriginRelativeX":
		instance.FkeyOriginRelativeX = GongExtractFloat(valueExpr)
	case "FkeyOriginRelativeY":
		instance.FkeyOriginRelativeY = GongExtractFloat(valueExpr)
	case "PitchLines":
		GongUnmarshallPointer(&instance.PitchLines, valueExpr, identifierMap)
	case "PitchHeight":
		instance.PitchHeight = GongExtractFloat(valueExpr)
	case "NbPitchLines":
		instance.NbPitchLines = GongExtractInt(valueExpr)
	case "BeatLines":
		GongUnmarshallPointer(&instance.BeatLines, valueExpr, identifierMap)
	case "BeatLinesHeightRatio":
		instance.BeatLinesHeightRatio = GongExtractFloat(valueExpr)
	case "NbBeatLines":
		instance.NbBeatLines = GongExtractInt(valueExpr)
	case "NbOfBeatsInTheme":
		instance.NbOfBeatsInTheme = GongExtractInt(valueExpr)
	case "FirstVoice":
		GongUnmarshallPointer(&instance.FirstVoice, valueExpr, identifierMap)
	case "FirstVoiceShiftedRigth":
		GongUnmarshallPointer(&instance.FirstVoiceShiftedRigth, valueExpr, identifierMap)
	case "FirstVoiceShiftX":
		instance.FirstVoiceShiftX = GongExtractFloat(valueExpr)
	case "FirstVoiceShiftY":
		instance.FirstVoiceShiftY = GongExtractFloat(valueExpr)
	case "SecondVoice":
		GongUnmarshallPointer(&instance.SecondVoice, valueExpr, identifierMap)
	case "SecondVoiceShiftedRight":
		GongUnmarshallPointer(&instance.SecondVoiceShiftedRight, valueExpr, identifierMap)
	case "PitchDifference":
		instance.PitchDifference = GongExtractInt(valueExpr)
	case "BeatsPerSecond":
		instance.BeatsPerSecond = GongExtractFloat(valueExpr)
	case "Level":
		instance.Level = GongExtractFloat(valueExpr)
	case "FirstVoiceNotes":
		GongUnmarshallPointer(&instance.FirstVoiceNotes, valueExpr, identifierMap)
	case "FirstVoiceNotesShiftedRight":
		GongUnmarshallPointer(&instance.FirstVoiceNotesShiftedRight, valueExpr, identifierMap)
	case "SecondVoiceNotes":
		GongUnmarshallPointer(&instance.SecondVoiceNotes, valueExpr, identifierMap)
	case "SecondVoiceNotesShiftedRight":
		GongUnmarshallPointer(&instance.SecondVoiceNotesShiftedRight, valueExpr, identifierMap)
	case "IsMinor":
		instance.IsMinor = GongExtractBool(valueExpr)
	case "ThemeBinaryEncoding":
		instance.ThemeBinaryEncoding = GongExtractInt(valueExpr)
	case "OriginX":
		instance.OriginX = GongExtractFloat(valueExpr)
	case "OriginY":
		instance.OriginY = GongExtractFloat(valueExpr)
	case "HorizontalAxis":
		GongUnmarshallPointer(&instance.HorizontalAxis, valueExpr, identifierMap)
	case "VerticalAxis":
		GongUnmarshallPointer(&instance.VerticalAxis, valueExpr, identifierMap)
	case "SpiralOrigin":
		GongUnmarshallPointer(&instance.SpiralOrigin, valueExpr, identifierMap)
	case "SpiralOriginX":
		instance.SpiralOriginX = GongExtractFloat(valueExpr)
	case "SpiralOriginY":
		instance.SpiralOriginY = GongExtractFloat(valueExpr)
	case "OriginCrossWidth":
		instance.OriginCrossWidth = GongExtractFloat(valueExpr)
	case "SpiralRadiusRatio":
		instance.SpiralRadiusRatio = GongExtractFloat(valueExpr)
	case "ShowSpiralBezierConstruct":
		instance.ShowSpiralBezierConstruct = GongExtractBool(valueExpr)
	case "ShowInterpolationPoints":
		instance.ShowInterpolationPoints = GongExtractBool(valueExpr)
	case "ActualBeatsTemporalShift":
		instance.ActualBeatsTemporalShift = GongExtractInt(valueExpr)
	case "PathToStaticFiles":
		instance.PathToStaticFiles = GongExtractString(valueExpr)
	case "PathToGeneratedSVG":
		instance.PathToGeneratedSVG = GongExtractString(valueExpr)
	case "PathToGeneratedScore":
		instance.PathToGeneratedScore = GongExtractString(valueExpr)
	}
	return nil
}

type RhombusUnmarshaller struct{}

func (u *RhombusUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Rhombus)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *RhombusUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Rhombus)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "CenterX":
		instance.CenterX = GongExtractFloat(valueExpr)
	case "CenterY":
		instance.CenterY = GongExtractFloat(valueExpr)
	case "SideLength":
		instance.SideLength = GongExtractFloat(valueExpr)
	case "AngleDegree":
		instance.AngleDegree = GongExtractFloat(valueExpr)
	case "InsideAngle":
		instance.InsideAngle = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type RhombusGridUnmarshaller struct{}

func (u *RhombusGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(RhombusGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *RhombusGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RhombusGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Reference":
		GongUnmarshallPointer(&instance.Reference, valueExpr, identifierMap)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "Rhombuses":
		GongUnmarshallSliceOfPointers(&instance.Rhombuses, valueExpr, identifierMap)
	}
	return nil
}

type ShapeCategoryUnmarshaller struct{}

func (u *ShapeCategoryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ShapeCategory)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ShapeCategoryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShapeCategory)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type SpiralBezierUnmarshaller struct{}

func (u *SpiralBezierUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SpiralBezier)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SpiralBezierUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpiralBezier)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "ControlPointStartX":
		instance.ControlPointStartX = GongExtractFloat(valueExpr)
	case "ControlPointStartY":
		instance.ControlPointStartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "ControlPointEndX":
		instance.ControlPointEndX = GongExtractFloat(valueExpr)
	case "ControlPointEndY":
		instance.ControlPointEndY = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type SpiralBezierGridUnmarshaller struct{}

func (u *SpiralBezierGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SpiralBezierGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SpiralBezierGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpiralBezierGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "SpiralBeziers":
		GongUnmarshallSliceOfPointers(&instance.SpiralBeziers, valueExpr, identifierMap)
	}
	return nil
}

type SpiralCircleUnmarshaller struct{}

func (u *SpiralCircleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SpiralCircle)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SpiralCircleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpiralCircle)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "CenterX":
		instance.CenterX = GongExtractFloat(valueExpr)
	case "CenterY":
		instance.CenterY = GongExtractFloat(valueExpr)
	case "HasBespokeRadius":
		instance.HasBespokeRadius = GongExtractBool(valueExpr)
	case "BespopkeRadius":
		instance.BespopkeRadius = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Pitch":
		instance.Pitch = GongExtractInt(valueExpr)
	case "ShowName":
		instance.ShowName = GongExtractBool(valueExpr)
	case "BeatNb":
		instance.BeatNb = GongExtractInt(valueExpr)
	case "Path":
		instance.Path = GongExtractString(valueExpr)
	}
	return nil
}

type SpiralCircleGridUnmarshaller struct{}

func (u *SpiralCircleGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SpiralCircleGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SpiralCircleGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpiralCircleGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "SpiralRhombusGrid":
		GongUnmarshallPointer(&instance.SpiralRhombusGrid, valueExpr, identifierMap)
	case "SpiralCircles":
		GongUnmarshallSliceOfPointers(&instance.SpiralCircles, valueExpr, identifierMap)
	}
	return nil
}

type SpiralLineUnmarshaller struct{}

func (u *SpiralLineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SpiralLine)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SpiralLineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpiralLine)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type SpiralLineGridUnmarshaller struct{}

func (u *SpiralLineGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SpiralLineGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SpiralLineGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpiralLineGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "SpiralLines":
		GongUnmarshallSliceOfPointers(&instance.SpiralLines, valueExpr, identifierMap)
	}
	return nil
}

type SpiralOriginUnmarshaller struct{}

func (u *SpiralOriginUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SpiralOrigin)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SpiralOriginUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpiralOrigin)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type SpiralRhombusUnmarshaller struct{}

func (u *SpiralRhombusUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SpiralRhombus)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SpiralRhombusUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpiralRhombus)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "X_r0":
		instance.X_r0 = GongExtractFloat(valueExpr)
	case "Y_r0":
		instance.Y_r0 = GongExtractFloat(valueExpr)
	case "X_r1":
		instance.X_r1 = GongExtractFloat(valueExpr)
	case "Y_r1":
		instance.Y_r1 = GongExtractFloat(valueExpr)
	case "X_r2":
		instance.X_r2 = GongExtractFloat(valueExpr)
	case "Y_r2":
		instance.Y_r2 = GongExtractFloat(valueExpr)
	case "X_r3":
		instance.X_r3 = GongExtractFloat(valueExpr)
	case "Y_r3":
		instance.Y_r3 = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type SpiralRhombusGridUnmarshaller struct{}

func (u *SpiralRhombusGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SpiralRhombusGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SpiralRhombusGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpiralRhombusGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "SpiralRhombuses":
		GongUnmarshallSliceOfPointers(&instance.SpiralRhombuses, valueExpr, identifierMap)
	}
	return nil
}

type VerticalAxisUnmarshaller struct{}

func (u *VerticalAxisUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(VerticalAxis)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *VerticalAxisUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VerticalAxis)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDisplayed":
		instance.IsDisplayed = GongExtractBool(valueExpr)
	case "ShapeCategory":
		GongUnmarshallPointer(&instance.ShapeCategory, valueExpr, identifierMap)
	case "AxisHandleBorderLength":
		instance.AxisHandleBorderLength = GongExtractFloat(valueExpr)
	case "Axis_Length":
		instance.Axis_Length = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}
