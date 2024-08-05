// generated code - do not edit
package models

import (
	"bufio"
	"errors"
	"go/ast"
	"go/doc/comment"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError
var dummy_time_import time.Time

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

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(stage *StageStruct, pathToFile string) error {

	ReplaceOldDeclarationsInFile(pathToFile)

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	// startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	// log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileFromAst(stage *StageStruct, inFile *ast.File, fset *token.FileSet) error {
	// if there is a meta package import, it is the third import
	if len(inFile.Imports) > 3 {
		log.Fatalln("Too many imports in file", inFile.Name)
	}
	if len(inFile.Imports) == 3 {
		stage.MetaPackageImportAlias = inFile.Imports[2].Name.Name
		stage.MetaPackageImportPath = inFile.Imports[2].Path.Value
	}

	// astCoordinate := "File "
	// log.Println(// astCoordinate)
	for _, decl := range inFile.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl
			// astCoordinate := // astCoordinate + "\tFunction " + funcDecl.Name.Name
			if name := funcDecl.Name; name != nil {
				isOfInterest := strings.Contains(funcDecl.Name.Name, "_")
				if !isOfInterest {
					continue
				}
				// log.Println(// astCoordinate)
			}
			if doc := funcDecl.Doc; doc != nil {
				// astCoordinate := // astCoordinate + "\tDoc"
				for _, comment := range doc.List {
					_ = comment
					// astCoordinate := // astCoordinate + "\tComment: " + comment.Text
					// log.Println(// astCoordinate)
				}
			}
			if body := funcDecl.Body; body != nil {
				// astCoordinate := // astCoordinate + "\tBody: "
				for _, stmt := range body.List {
					switch stmt := stmt.(type) {
					case *ast.ExprStmt:
						exprStmt := stmt
						// astCoordinate := // astCoordinate + "\tExprStmt: "
						switch expr := exprStmt.X.(type) {
						case *ast.CallExpr:
							// astCoordinate := // astCoordinate + "\tCallExpr: "
							callExpr := expr
							switch fun := callExpr.Fun.(type) {
							case *ast.Ident:
								ident := fun
								_ = ident
								// astCoordinate := // astCoordinate + "\tIdent: " + ident.Name
								// log.Println(// astCoordinate)
							}
						}
					case *ast.AssignStmt:
						// Create an ast.CommentMap from the ast.File's comments.
						// This helps keeping the association between comments
						// and AST nodes.
						cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName :=
							UnmarshallGongstructStaging(
								stage, &cmap, assignStmt, astCoordinate)
						_ = instance
						_ = id
						_ = gongstruct
						_ = fieldName
					}
				}
			}
		case *ast.GenDecl:
			genDecl := decl
			// log.Println("\tAST GenDecl: ")
			if doc := genDecl.Doc; doc != nil {
				for _, comment := range doc.List {
					_ = comment
					// log.Println("\tAST Comment: ", comment.Text)
				}
			}
			for _, spec := range genDecl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec
					if path := importSpec.Path; path != nil {
						// log.Println("\t\tAST Path: ", path.Value)
					}
				case *ast.ValueSpec:
					ident := spec.Names[0]
					_ = ident
					if !strings.HasPrefix(ident.Name, "_") {
						continue
					}
					// declaration of a variable without initial value
					if len(spec.Values) == 0 {
						continue
					}
					switch compLit := spec.Values[0].(type) {
					case *ast.CompositeLit:
						var key string
						_ = key
						var value string
						_ = value
						for _, elt := range compLit.Elts {

							// each elt is an expression for struct or for field such as
							// for struct
							//
							//         "dummy.Dummy": &(dummy.Dummy{})
							//
							// or, for field
							//
							//          "dummy.Dummy.Name": (dummy.Dummy{}).Name,
							//
							// first node in the AST is a key value expression
							var ok bool
							var kve *ast.KeyValueExpr
							if kve, ok = elt.(*ast.KeyValueExpr); !ok {
								log.Fatal("Expression should be key value expression" +
									fset.Position(kve.Pos()).String())
							}

							switch bl := kve.Key.(type) {
							case *ast.BasicLit:
								key = bl.Value // "\"dumm.Dummy\"" is the value

								// one remove the ambracing double quotes
								key = strings.TrimPrefix(key, "\"")
								key = strings.TrimSuffix(key, "\"")
							}
							var expressionType GONG__ExpressionType = GONG__STRUCT_INSTANCE
							var docLink GONG__Identifier

							var fieldName string
							var ue *ast.UnaryExpr
							if ue, ok = kve.Value.(*ast.UnaryExpr); !ok {
								expressionType = GONG__FIELD_OR_CONST_VALUE
							}

							var callExpr *ast.CallExpr
							if callExpr, ok = kve.Value.(*ast.CallExpr); ok {

								var se *ast.SelectorExpr
								if se, ok = callExpr.Fun.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector expression" +
										fset.Position(callExpr.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}

								// check the arg type to select wether this is a int or a string enum
								var bl *ast.BasicLit
								if bl, ok = callExpr.Args[0].(*ast.BasicLit); ok {
									switch bl.Kind {
									case token.STRING:
										expressionType = GONG__ENUM_CAST_STRING
									case token.INT:
										expressionType = GONG__ENUM_CAST_INT
									}
								} else {
									log.Fatal("Expression should be a basic lit" +
										fset.Position(se.Pos()).String())
								}

								docLink.Ident = id.Name + "." + se.Sel.Name
								_ = callExpr
							}

							var se2 *ast.SelectorExpr
							switch expressionType {
							case GONG__FIELD_OR_CONST_VALUE:
								if se2, ok = kve.Value.(*ast.SelectorExpr); ok {

									var ident *ast.Ident
									if _, ok = se2.X.(*ast.ParenExpr); ok {
										expressionType = GONG__FIELD_VALUE
										fieldName = se2.Sel.Name
									} else if ident, ok = se2.X.(*ast.Ident); ok {
										expressionType = GONG__IDENTIFIER_CONST
										docLink.Ident = ident.Name + "." + se2.Sel.Name
									} else {
										log.Fatal("Expression should be a selector expression or an ident" +
											fset.Position(kve.Pos()).String())
									}
								} else {

								}
							}

							var pe *ast.ParenExpr
							switch expressionType {
							case GONG__STRUCT_INSTANCE:
								if pe, ok = ue.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							case GONG__FIELD_VALUE:
								if pe, ok = se2.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							}
							switch expressionType {
							case GONG__FIELD_VALUE, GONG__STRUCT_INSTANCE:
								// expect a Composite Litteral with no Element <type>{}
								var cl *ast.CompositeLit
								if cl, ok = pe.X.(*ast.CompositeLit); !ok {
									log.Fatal("Expression should be a composite lit" +
										fset.Position(pe.Pos()).String())
								}

								var se *ast.SelectorExpr
								if se, ok = cl.Type.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector" +
										fset.Position(cl.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}
								docLink.Ident = id.Name + "." + se.Sel.Name
							}

							switch expressionType {
							case GONG__FIELD_VALUE:
								docLink.Ident += "." + fieldName
							}

							// if map_DocLink_Identifier has the same ident, this means
							// that no renaming has occured since the last processing of the
							// file. But it is neccessary to keep it in memory for the
							// marshalling
							if docLink.Ident == key {
								// continue
							}

							// otherwise, one stores the new ident (after renaming) in the
							// renaming map
							docLink.Type = expressionType
							stage.Map_DocLink_Renaming[key] = docLink
						}
					}
				}
			}
		}

	}
	return nil
}

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_Axis = make(map[string]*Axis)
var __gong__map_AxisGrid = make(map[string]*AxisGrid)
var __gong__map_Bezier = make(map[string]*Bezier)
var __gong__map_BezierGrid = make(map[string]*BezierGrid)
var __gong__map_BezierGridStack = make(map[string]*BezierGridStack)
var __gong__map_Circle = make(map[string]*Circle)
var __gong__map_CircleGrid = make(map[string]*CircleGrid)
var __gong__map_HorizontalAxis = make(map[string]*HorizontalAxis)
var __gong__map_Key = make(map[string]*Key)
var __gong__map_Parameter = make(map[string]*Parameter)
var __gong__map_Rhombus = make(map[string]*Rhombus)
var __gong__map_RhombusGrid = make(map[string]*RhombusGrid)
var __gong__map_ShapeCategory = make(map[string]*ShapeCategory)
var __gong__map_VerticalAxis = make(map[string]*VerticalAxis)

// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
// to be recognized as a proper identifier.
// While this was introduced in go 1.19, it is not yet implemented in
// gopls (see [issue](https://github.com/golang/go/issues/57559)
func lookupPackage(name string) (importPath string, ok bool) {
	return name, true
}
func lookupSym(recv, name string) (ok bool) {
	if recv == "" {
		return true
	}
	return false
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(stage *StageStruct, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {

	// used for debug purposes
	astCoordinate := "\tAssignStmt: "

	//
	// First parse all comment groups in the assignement
	// if a comment "//gong:ident [DocLink]" is met and is followed by a string assignement.
	// modify the following AST assignement to assigns the DocLink text to the string value
	//

	// get the comment group of the assigStmt
	commentGroups := (*cmap)[assignStmt]
	// get the the prefix
	var hasGongIdentDirective bool
	var commentText string
	var docLinkText string
	for _, commentGroup := range commentGroups {
		for _, comment := range commentGroup.List {
			if strings.HasPrefix(comment.Text, "//gong:ident") {
				hasGongIdentDirective = true
				commentText = comment.Text
			}
		}
	}
	if hasGongIdentDirective {
		// parser configured to find doclinks
		var docLinkFinder comment.Parser
		docLinkFinder.LookupPackage = lookupPackage
		docLinkFinder.LookupSym = lookupSym
		doc := docLinkFinder.Parse(commentText)

		for _, block := range doc.Content {
			switch paragraph := block.(type) {
			case *comment.Paragraph:
				_ = paragraph
				for _, text := range paragraph.Text {
					switch docLink := text.(type) {
					case *comment.DocLink:
						if docLink.Recv == "" {
							docLinkText = docLink.ImportPath + "." + docLink.Name
						} else {
							docLinkText = docLink.ImportPath + "." + docLink.Recv + "." + docLink.Name
						}

						// we check wether the doc link has been renamed
						// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
						if renamed, ok := (stage.Map_DocLink_Renaming)[docLinkText]; ok {
							docLinkText = renamed.Ident
						}
					}
				}
			}
		}
	}

	for rank, expr := range assignStmt.Lhs {
		if rank > 0 {
			continue
		}
		switch expr := expr.(type) {
		case *ast.Ident:
			// we are on a variable declaration
			ident := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			// log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			// log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		// astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			// astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				// astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					// astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						// astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							// astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								// astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									// astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										_ = ident
										// astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										// log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										// astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										// log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							_ = astCoordinate2
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								// astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									_ = ident
									// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									// log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									// log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers
									case "Axis":
										instanceAxis := (&Axis{Name: instanceName}).Stage(stage)
										instance = any(instanceAxis)
										__gong__map_Axis[identifier] = instanceAxis
									case "AxisGrid":
										instanceAxisGrid := (&AxisGrid{Name: instanceName}).Stage(stage)
										instance = any(instanceAxisGrid)
										__gong__map_AxisGrid[identifier] = instanceAxisGrid
									case "Bezier":
										instanceBezier := (&Bezier{Name: instanceName}).Stage(stage)
										instance = any(instanceBezier)
										__gong__map_Bezier[identifier] = instanceBezier
									case "BezierGrid":
										instanceBezierGrid := (&BezierGrid{Name: instanceName}).Stage(stage)
										instance = any(instanceBezierGrid)
										__gong__map_BezierGrid[identifier] = instanceBezierGrid
									case "BezierGridStack":
										instanceBezierGridStack := (&BezierGridStack{Name: instanceName}).Stage(stage)
										instance = any(instanceBezierGridStack)
										__gong__map_BezierGridStack[identifier] = instanceBezierGridStack
									case "Circle":
										instanceCircle := (&Circle{Name: instanceName}).Stage(stage)
										instance = any(instanceCircle)
										__gong__map_Circle[identifier] = instanceCircle
									case "CircleGrid":
										instanceCircleGrid := (&CircleGrid{Name: instanceName}).Stage(stage)
										instance = any(instanceCircleGrid)
										__gong__map_CircleGrid[identifier] = instanceCircleGrid
									case "HorizontalAxis":
										instanceHorizontalAxis := (&HorizontalAxis{Name: instanceName}).Stage(stage)
										instance = any(instanceHorizontalAxis)
										__gong__map_HorizontalAxis[identifier] = instanceHorizontalAxis
									case "Key":
										instanceKey := (&Key{Name: instanceName}).Stage(stage)
										instance = any(instanceKey)
										__gong__map_Key[identifier] = instanceKey
									case "Parameter":
										instanceParameter := (&Parameter{Name: instanceName}).Stage(stage)
										instance = any(instanceParameter)
										__gong__map_Parameter[identifier] = instanceParameter
									case "Rhombus":
										instanceRhombus := (&Rhombus{Name: instanceName}).Stage(stage)
										instance = any(instanceRhombus)
										__gong__map_Rhombus[identifier] = instanceRhombus
									case "RhombusGrid":
										instanceRhombusGrid := (&RhombusGrid{Name: instanceName}).Stage(stage)
										instance = any(instanceRhombusGrid)
										__gong__map_RhombusGrid[identifier] = instanceRhombusGrid
									case "ShapeCategory":
										instanceShapeCategory := (&ShapeCategory{Name: instanceName}).Stage(stage)
										instance = any(instanceShapeCategory)
										__gong__map_ShapeCategory[identifier] = instanceShapeCategory
									case "VerticalAxis":
										instanceVerticalAxis := (&VerticalAxis{Name: instanceName}).Stage(stage)
										instance = any(instanceVerticalAxis)
										__gong__map_VerticalAxis[identifier] = instanceVerticalAxis
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					// astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					// log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					// astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						// log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Fatalln("gongstructName not found for identifier", identifier)
						}
						switch gongstructName {
						// insertion point for basic lit assignments
						case "Axis":
							switch fieldName {
							// insertion point for date assign code
							}
						case "AxisGrid":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Bezier":
							switch fieldName {
							// insertion point for date assign code
							}
						case "BezierGrid":
							switch fieldName {
							// insertion point for date assign code
							}
						case "BezierGridStack":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Circle":
							switch fieldName {
							// insertion point for date assign code
							}
						case "CircleGrid":
							switch fieldName {
							// insertion point for date assign code
							}
						case "HorizontalAxis":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Key":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Parameter":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Rhombus":
							switch fieldName {
							// insertion point for date assign code
							}
						case "RhombusGrid":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ShapeCategory":
							switch fieldName {
							// insertion point for date assign code
							}
						case "VerticalAxis":
							switch fieldName {
							// insertion point for date assign code
							}
						}
					}
				}
			case *ast.Ident:
				// append function
				ident := fun
				_ = ident
				// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			for _, arg := range callExpr.Args {
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident:
					ident := arg
					_ = ident
					// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
					// log.Println(astCoordinate)
					var ok bool
					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Fatalln("gongstructName not found for identifier", identifier)
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "Axis":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "AxisGrid":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Axiss":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Axis[targetIdentifier]
							__gong__map_AxisGrid[identifier].Axiss =
								append(__gong__map_AxisGrid[identifier].Axiss, target)
						}
					case "Bezier":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "BezierGrid":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Beziers":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Bezier[targetIdentifier]
							__gong__map_BezierGrid[identifier].Beziers =
								append(__gong__map_BezierGrid[identifier].Beziers, target)
						}
					case "BezierGridStack":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "BezierGrids":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_BezierGrid[targetIdentifier]
							__gong__map_BezierGridStack[identifier].BezierGrids =
								append(__gong__map_BezierGridStack[identifier].BezierGrids, target)
						}
					case "Circle":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "CircleGrid":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Circles":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Circle[targetIdentifier]
							__gong__map_CircleGrid[identifier].Circles =
								append(__gong__map_CircleGrid[identifier].Circles, target)
						}
					case "HorizontalAxis":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Key":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Parameter":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Rhombus":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "RhombusGrid":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Rhombuses":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Rhombus[targetIdentifier]
							__gong__map_RhombusGrid[identifier].Rhombuses =
								append(__gong__map_RhombusGrid[identifier].Rhombuses, target)
						}
					case "ShapeCategory":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "VerticalAxis":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					}
				case *ast.SelectorExpr:
					slcExpr := arg
					// astCoordinate := astCoordinate + "\tSelectorExpr"
					switch X := slcExpr.X.(type) {
					case *ast.Ident:
						ident := X
						_ = ident
						// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
						// log.Println(astCoordinate)

					}
					if Sel := slcExpr.Sel; Sel != nil {
						// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
						// log.Println(astCoordinate)
					}
				}
			}
		case *ast.BasicLit, *ast.UnaryExpr:

			var basicLit *ast.BasicLit
			var exprSign = 1.0
			_ = exprSign // in case this is not used

			if bl, ok := expr.(*ast.BasicLit); ok {
				// expression is  for instance ... = 18.000
				basicLit = bl
			} else if ue, ok := expr.(*ast.UnaryExpr); ok {
				// expression is  for instance ... = -18.000
				// we want to extract a *ast.BasicLit from the *ast.UnaryExpr
				basicLit = ue.X.(*ast.BasicLit)
				exprSign = -1
			}

			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			// substitute the RHS part of the assignment if a //gong:ident directive is met
			if hasGongIdentDirective {
				basicLit.Value = "[" + docLinkText + "]"
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "Axis":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Axis[identifier].Name = fielValue
				case "Angle":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Axis[identifier].Angle = exprSign * fielValue
				case "Length":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Axis[identifier].Length = exprSign * fielValue
				case "CenterX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Axis[identifier].CenterX = exprSign * fielValue
				case "CenterY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Axis[identifier].CenterY = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Axis[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Axis[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Axis[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Axis[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Axis[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Axis[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Axis[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Axis[identifier].Transform = fielValue
				}
			case "AxisGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AxisGrid[identifier].Name = fielValue
				}
			case "Bezier":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bezier[identifier].Name = fielValue
				case "StartX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bezier[identifier].StartX = exprSign * fielValue
				case "StartY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bezier[identifier].StartY = exprSign * fielValue
				case "ControlPointStartX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bezier[identifier].ControlPointStartX = exprSign * fielValue
				case "ControlPointStartY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bezier[identifier].ControlPointStartY = exprSign * fielValue
				case "EndX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bezier[identifier].EndX = exprSign * fielValue
				case "EndY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bezier[identifier].EndY = exprSign * fielValue
				case "ControlPointEndX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bezier[identifier].ControlPointEndX = exprSign * fielValue
				case "ControlPointEndY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bezier[identifier].ControlPointEndY = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bezier[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bezier[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bezier[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bezier[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bezier[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bezier[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bezier[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bezier[identifier].Transform = fielValue
				}
			case "BezierGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_BezierGrid[identifier].Name = fielValue
				}
			case "BezierGridStack":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_BezierGridStack[identifier].Name = fielValue
				}
			case "Circle":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Name = fielValue
				case "CenterX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].CenterX = exprSign * fielValue
				case "CenterY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].CenterY = exprSign * fielValue
				case "BespopkeRadius":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].BespopkeRadius = exprSign * fielValue
				case "Pitch":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].Pitch = int(exprSign) * int(fielValue)
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Transform = fielValue
				}
			case "CircleGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CircleGrid[identifier].Name = fielValue
				}
			case "HorizontalAxis":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_HorizontalAxis[identifier].Name = fielValue
				case "AxisHandleBorderLength":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_HorizontalAxis[identifier].AxisHandleBorderLength = exprSign * fielValue
				case "Axis_Length":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_HorizontalAxis[identifier].Axis_Length = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_HorizontalAxis[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_HorizontalAxis[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_HorizontalAxis[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_HorizontalAxis[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_HorizontalAxis[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_HorizontalAxis[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_HorizontalAxis[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_HorizontalAxis[identifier].Transform = fielValue
				}
			case "Key":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Key[identifier].Name = fielValue
				case "Path":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Key[identifier].Path = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Key[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Key[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Key[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Key[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Key[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Key[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Key[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Key[identifier].Transform = fielValue
				}
			case "Parameter":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Parameter[identifier].Name = fielValue
				case "N":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].N = int(exprSign) * int(fielValue)
				case "M":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].M = int(exprSign) * int(fielValue)
				case "Z":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].Z = int(exprSign) * int(fielValue)
				case "InsideAngle":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].InsideAngle = exprSign * fielValue
				case "SideLength":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].SideLength = exprSign * fielValue
				case "StackWidth":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].StackWidth = int(exprSign) * int(fielValue)
				case "NbShitRight":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].NbShitRight = int(exprSign) * int(fielValue)
				case "StackHeight":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].StackHeight = int(exprSign) * int(fielValue)
				case "BezierControlLengthRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].BezierControlLengthRatio = exprSign * fielValue
				case "FkeySizeRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].FkeySizeRatio = exprSign * fielValue
				case "FkeyOriginRelativeX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].FkeyOriginRelativeX = exprSign * fielValue
				case "FkeyOriginRelativeY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].FkeyOriginRelativeY = exprSign * fielValue
				case "PitchHeight":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].PitchHeight = exprSign * fielValue
				case "NbPitchLines":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].NbPitchLines = int(exprSign) * int(fielValue)
				case "MeasureLinesHeightRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].MeasureLinesHeightRatio = exprSign * fielValue
				case "NbMeasureLines":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].NbMeasureLines = int(exprSign) * int(fielValue)
				case "NbMeasureLinesPerCurve":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].NbMeasureLinesPerCurve = int(exprSign) * int(fielValue)
				case "FirstVoiceShiftX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].FirstVoiceShiftX = exprSign * fielValue
				case "FirstVoiceShiftY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].FirstVoiceShiftY = exprSign * fielValue
				case "PitchDifference":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].PitchDifference = int(exprSign) * int(fielValue)
				case "Speed":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].Speed = exprSign * fielValue
				case "Level":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].Level = exprSign * fielValue
				case "OriginX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].OriginX = exprSign * fielValue
				case "OriginY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].OriginY = exprSign * fielValue
				}
			case "Rhombus":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rhombus[identifier].Name = fielValue
				case "CenterX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rhombus[identifier].CenterX = exprSign * fielValue
				case "CenterY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rhombus[identifier].CenterY = exprSign * fielValue
				case "SideLength":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rhombus[identifier].SideLength = exprSign * fielValue
				case "Angle":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rhombus[identifier].Angle = exprSign * fielValue
				case "InsideAngle":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rhombus[identifier].InsideAngle = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rhombus[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rhombus[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rhombus[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rhombus[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rhombus[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rhombus[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rhombus[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rhombus[identifier].Transform = fielValue
				}
			case "RhombusGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RhombusGrid[identifier].Name = fielValue
				}
			case "ShapeCategory":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ShapeCategory[identifier].Name = fielValue
				}
			case "VerticalAxis":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_VerticalAxis[identifier].Name = fielValue
				case "AxisHandleBorderLength":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VerticalAxis[identifier].AxisHandleBorderLength = exprSign * fielValue
				case "Axis_Length":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VerticalAxis[identifier].Axis_Length = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_VerticalAxis[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VerticalAxis[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_VerticalAxis[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VerticalAxis[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VerticalAxis[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_VerticalAxis[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_VerticalAxis[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_VerticalAxis[identifier].Transform = fielValue
				}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			_ = ident
			// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments
			case "Axis":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Axis[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_Axis[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "AxisGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Reference":
					targetIdentifier := ident.Name
					__gong__map_AxisGrid[identifier].Reference = __gong__map_Axis[targetIdentifier]
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AxisGrid[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_AxisGrid[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "Bezier":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bezier[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_Bezier[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "BezierGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Reference":
					targetIdentifier := ident.Name
					__gong__map_BezierGrid[identifier].Reference = __gong__map_Bezier[targetIdentifier]
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_BezierGrid[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_BezierGrid[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "BezierGridStack":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_BezierGridStack[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_BezierGridStack[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "Circle":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_Circle[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				case "HasBespokeRadius":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].HasBespokeRadius = fielValue
				}
			case "CircleGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Reference":
					targetIdentifier := ident.Name
					__gong__map_CircleGrid[identifier].Reference = __gong__map_Circle[targetIdentifier]
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_CircleGrid[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_CircleGrid[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "HorizontalAxis":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_HorizontalAxis[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_HorizontalAxis[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "Key":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Key[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_Key[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "Parameter":
				switch fieldName {
				// insertion point for field dependant code
				case "InitialRhombus":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].InitialRhombus = __gong__map_Rhombus[targetIdentifier]
				case "InitialCircle":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].InitialCircle = __gong__map_Circle[targetIdentifier]
				case "InitialRhombusGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].InitialRhombusGrid = __gong__map_RhombusGrid[targetIdentifier]
				case "InitialCircleGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].InitialCircleGrid = __gong__map_CircleGrid[targetIdentifier]
				case "InitialAxis":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].InitialAxis = __gong__map_Axis[targetIdentifier]
				case "RotatedAxis":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].RotatedAxis = __gong__map_Axis[targetIdentifier]
				case "RotatedRhombus":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].RotatedRhombus = __gong__map_Rhombus[targetIdentifier]
				case "RotatedRhombusGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].RotatedRhombusGrid = __gong__map_RhombusGrid[targetIdentifier]
				case "RotatedCircleGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].RotatedCircleGrid = __gong__map_CircleGrid[targetIdentifier]
				case "NextRhombus":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].NextRhombus = __gong__map_Rhombus[targetIdentifier]
				case "NextCircle":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].NextCircle = __gong__map_Circle[targetIdentifier]
				case "GrowingRhombusGridSeed":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowingRhombusGridSeed = __gong__map_Rhombus[targetIdentifier]
				case "GrowingRhombusGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowingRhombusGrid = __gong__map_RhombusGrid[targetIdentifier]
				case "GrowingCircleGridSeed":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowingCircleGridSeed = __gong__map_Circle[targetIdentifier]
				case "GrowingCircleGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowingCircleGrid = __gong__map_CircleGrid[targetIdentifier]
				case "GrowingCircleGridLeftSeed":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowingCircleGridLeftSeed = __gong__map_Circle[targetIdentifier]
				case "GrowingCircleGridLeft":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowingCircleGridLeft = __gong__map_CircleGrid[targetIdentifier]
				case "ConstructionAxis":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].ConstructionAxis = __gong__map_Axis[targetIdentifier]
				case "ConstructionAxisGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].ConstructionAxisGrid = __gong__map_AxisGrid[targetIdentifier]
				case "ConstructionCircle":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].ConstructionCircle = __gong__map_Circle[targetIdentifier]
				case "ConstructionCircleGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].ConstructionCircleGrid = __gong__map_CircleGrid[targetIdentifier]
				case "GrowthCurveSegment":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowthCurveSegment = __gong__map_Bezier[targetIdentifier]
				case "GrowthCurve":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowthCurve = __gong__map_BezierGrid[targetIdentifier]
				case "GrowthCurveShiftedRightSeed":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowthCurveShiftedRightSeed = __gong__map_Bezier[targetIdentifier]
				case "GrowthCurveShiftedRight":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowthCurveShiftedRight = __gong__map_BezierGrid[targetIdentifier]
				case "GrowthCurveNextSeed":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowthCurveNextSeed = __gong__map_Bezier[targetIdentifier]
				case "GrowthCurveNext":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowthCurveNext = __gong__map_BezierGrid[targetIdentifier]
				case "GrowthCurveNextShiftedRightSeed":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowthCurveNextShiftedRightSeed = __gong__map_Bezier[targetIdentifier]
				case "GrowthCurveNextShiftedRight":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowthCurveNextShiftedRight = __gong__map_BezierGrid[targetIdentifier]
				case "GrowthCurveStack":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowthCurveStack = __gong__map_BezierGridStack[targetIdentifier]
				case "Fkey":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].Fkey = __gong__map_Key[targetIdentifier]
				case "PitchLines":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].PitchLines = __gong__map_AxisGrid[targetIdentifier]
				case "MeasureLines":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].MeasureLines = __gong__map_AxisGrid[targetIdentifier]
				case "FirstVoice":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].FirstVoice = __gong__map_BezierGrid[targetIdentifier]
				case "FirstVoiceShiftRigth":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].FirstVoiceShiftRigth = __gong__map_BezierGrid[targetIdentifier]
				case "SecondVoice":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SecondVoice = __gong__map_BezierGrid[targetIdentifier]
				case "SecondVoiceShiftedRight":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SecondVoiceShiftedRight = __gong__map_BezierGrid[targetIdentifier]
				case "FirstVoiceNotes":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].FirstVoiceNotes = __gong__map_CircleGrid[targetIdentifier]
				case "FirstVoiceNotesShiftedRight":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].FirstVoiceNotesShiftedRight = __gong__map_CircleGrid[targetIdentifier]
				case "SecondVoiceNotes":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SecondVoiceNotes = __gong__map_CircleGrid[targetIdentifier]
				case "SecondVoiceNotesShiftedRight":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SecondVoiceNotesShiftedRight = __gong__map_CircleGrid[targetIdentifier]
				case "IsMinor":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].IsMinor = fielValue
				case "HorizontalAxis":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].HorizontalAxis = __gong__map_HorizontalAxis[targetIdentifier]
				case "VerticalAxis":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].VerticalAxis = __gong__map_VerticalAxis[targetIdentifier]
				}
			case "Rhombus":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rhombus[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_Rhombus[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "RhombusGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Reference":
					targetIdentifier := ident.Name
					__gong__map_RhombusGrid[identifier].Reference = __gong__map_Rhombus[targetIdentifier]
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RhombusGrid[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_RhombusGrid[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "ShapeCategory":
				switch fieldName {
				// insertion point for field dependant code
				case "IsExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ShapeCategory[identifier].IsExpanded = fielValue
				}
			case "VerticalAxis":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_VerticalAxis[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_VerticalAxis[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			}
		case *ast.SelectorExpr:
			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for enums assignments
				case "Axis":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "AxisGrid":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Bezier":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "BezierGrid":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "BezierGridStack":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Circle":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "CircleGrid":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "HorizontalAxis":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Key":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Parameter":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Rhombus":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "RhombusGrid":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ShapeCategory":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "VerticalAxis":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}

// ReplaceOldDeclarationsInFile replaces specific text in a file at the given path.
func ReplaceOldDeclarationsInFile(pathToFile string) error {
	// Open the file for reading.
	file, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// replacing function with Injection
	pattern := regexp.MustCompile(`\b\w*Injection\b`)
	pattern2 := regexp.MustCompile(`\bmap_DocLink_Identifier_\w*\b`)

	// Temporary slice to hold lines from the file.
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Replace the target text with the desired text.
		line := strings.Replace(scanner.Text(), "var ___dummy__Time_stage time.Time", "var _ time.Time", -1)
		line = pattern.ReplaceAllString(line, "_")
		line = pattern2.ReplaceAllString(line, "_")

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Re-open the file for writing.
	file, err = os.Create(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the modified lines back to the file.
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
