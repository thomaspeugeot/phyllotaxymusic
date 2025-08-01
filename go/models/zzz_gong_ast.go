// generated code - do not edit
package models

import (
	"bufio"
	"embed"
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
var _ = dummy_strconv_import
var dummy_time_import time.Time
var _ = dummy_time_import

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
func ParseAstFile(stage *Stage, pathToFile string) error {

	ReplaceOldDeclarationsInFile(pathToFile)

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	// Read the file content using os.ReadFile
	content, err := os.ReadFile(fileOfInterest)
	if err != nil {
		return errors.New("Unable to read file " + err.Error())
	}

	// Assign the content to stage.contentWhenParsed
	stage.contentWhenParsed = string(content)

	fset := token.NewFileSet()
	// startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	// log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
// specified by pathToFile within the provided embed.FS directory and
// stages instances declared in the file using the provided Stage.
//
// Parameters:
//
//	stage:      The staging area to populate.
//	directory:  The embedded filesystem containing the file.
//	pathToFile: The path to the Go source file within the embedded filesystem.
//
// Returns:
//
//	An error if reading or parsing the file fails, or if ParseAstFileFromAst fails.
func ParseAstEmbeddedFile(stage *Stage, directory embed.FS, pathToFile string) error {

	// 1. Read the content from the embedded filesystem.
	//    We don't need filepath.Abs as embed.FS uses relative paths.
	//    We also skip ReplaceOldDeclarationsInFile as embedded files are read-only.
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		// Return a specific error if the file can't be read from the embed.FS
		return errors.New(stage.GetName() + "; Unable to read embedded file " + err.Error())
	}

	// 2. Create a FileSet to manage position information.
	fset := token.NewFileSet()

	// 3. Parse the file content.
	//    Instead of passing a filename for the OS to read, we pass the pathToFile
	//    as the identifier and the actual file content (fileContentBytes) as the source.
	//    parser.ParseComments is included to match the original function's behavior.
	//    The type *ast.File is returned by parser.ParseFile.
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		// Return a specific error if parsing fails
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	// 4. Call the common AST processing logic.
	//    Pass the parsed AST (*ast.File), the FileSet, and the stage.
	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileFromAst(stage *Stage, inFile *ast.File, fset *token.FileSet) error {
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
					case *ast.DeclStmt:
						if genDecl, ok := stmt.Decl.(*ast.GenDecl); ok && genDecl.Tok == token.CONST {
							for _, spec := range genDecl.Specs {
								if valueSpec, ok := spec.(*ast.ValueSpec); ok {
									for i, name := range valueSpec.Names {
										if i < len(valueSpec.Values) {
											if basicLit, ok := valueSpec.Values[i].(*ast.BasicLit); ok && basicLit.Kind == token.STRING {
												// Remove quotes from string literal
												value := strings.Trim(basicLit.Value, `"`)

												switch name.Name {
												case "__commitId__":
													if parsedUint, err := strconv.ParseUint(value, 10, 64); err == nil {
														stage.commitId = uint(parsedUint)
														stage.commitIdWhenParsed = stage.commitId
													}
												}
											}
										}
									}
								}
							}
						}
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
						// cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						var cmap ast.CommentMap
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
var __gong__map_Chapter = make(map[string]*Chapter)
var __gong__map_Circle = make(map[string]*Circle)
var __gong__map_CircleGrid = make(map[string]*CircleGrid)
var __gong__map_Content = make(map[string]*Content)
var __gong__map_ExportToMusicxml = make(map[string]*ExportToMusicxml)
var __gong__map_FrontCurve = make(map[string]*FrontCurve)
var __gong__map_FrontCurveStack = make(map[string]*FrontCurveStack)
var __gong__map_HorizontalAxis = make(map[string]*HorizontalAxis)
var __gong__map_Key = make(map[string]*Key)
var __gong__map_Parameter = make(map[string]*Parameter)
var __gong__map_Rhombus = make(map[string]*Rhombus)
var __gong__map_RhombusGrid = make(map[string]*RhombusGrid)
var __gong__map_ShapeCategory = make(map[string]*ShapeCategory)
var __gong__map_SpiralBezier = make(map[string]*SpiralBezier)
var __gong__map_SpiralBezierGrid = make(map[string]*SpiralBezierGrid)
var __gong__map_SpiralCircle = make(map[string]*SpiralCircle)
var __gong__map_SpiralCircleGrid = make(map[string]*SpiralCircleGrid)
var __gong__map_SpiralLine = make(map[string]*SpiralLine)
var __gong__map_SpiralLineGrid = make(map[string]*SpiralLineGrid)
var __gong__map_SpiralOrigin = make(map[string]*SpiralOrigin)
var __gong__map_SpiralRhombus = make(map[string]*SpiralRhombus)
var __gong__map_SpiralRhombusGrid = make(map[string]*SpiralRhombusGrid)
var __gong__map_VerticalAxis = make(map[string]*VerticalAxis)

// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
// to be recognized as a proper identifier.
// While this was introduced in go 1.19, it is not yet implemented in
// gopls (see [issue](https://github.com/golang/go/issues/57559)
func lookupPackage(name string) (importPath string, ok bool) {
	return name, true
}

func lookupSym(recv, name string) bool {
	return recv == ""
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(stage *Stage, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
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
			var basicLit *ast.BasicLit

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
										instanceAxis := new(Axis)
										instanceAxis.Name = instanceName
										instanceAxis.Stage(stage)
										instance = any(instanceAxis)
										__gong__map_Axis[identifier] = instanceAxis
									case "AxisGrid":
										instanceAxisGrid := new(AxisGrid)
										instanceAxisGrid.Name = instanceName
										instanceAxisGrid.Stage(stage)
										instance = any(instanceAxisGrid)
										__gong__map_AxisGrid[identifier] = instanceAxisGrid
									case "Bezier":
										instanceBezier := new(Bezier)
										instanceBezier.Name = instanceName
										instanceBezier.Stage(stage)
										instance = any(instanceBezier)
										__gong__map_Bezier[identifier] = instanceBezier
									case "BezierGrid":
										instanceBezierGrid := new(BezierGrid)
										instanceBezierGrid.Name = instanceName
										instanceBezierGrid.Stage(stage)
										instance = any(instanceBezierGrid)
										__gong__map_BezierGrid[identifier] = instanceBezierGrid
									case "BezierGridStack":
										instanceBezierGridStack := new(BezierGridStack)
										instanceBezierGridStack.Name = instanceName
										instanceBezierGridStack.Stage(stage)
										instance = any(instanceBezierGridStack)
										__gong__map_BezierGridStack[identifier] = instanceBezierGridStack
									case "Chapter":
										instanceChapter := new(Chapter)
										instanceChapter.Name = instanceName
										instanceChapter.Stage(stage)
										instance = any(instanceChapter)
										__gong__map_Chapter[identifier] = instanceChapter
									case "Circle":
										instanceCircle := new(Circle)
										instanceCircle.Name = instanceName
										instanceCircle.Stage(stage)
										instance = any(instanceCircle)
										__gong__map_Circle[identifier] = instanceCircle
									case "CircleGrid":
										instanceCircleGrid := new(CircleGrid)
										instanceCircleGrid.Name = instanceName
										instanceCircleGrid.Stage(stage)
										instance = any(instanceCircleGrid)
										__gong__map_CircleGrid[identifier] = instanceCircleGrid
									case "Content":
										instanceContent := new(Content)
										instanceContent.Name = instanceName
										instanceContent.Stage(stage)
										instance = any(instanceContent)
										__gong__map_Content[identifier] = instanceContent
									case "ExportToMusicxml":
										instanceExportToMusicxml := new(ExportToMusicxml)
										instanceExportToMusicxml.Name = instanceName
										instanceExportToMusicxml.Stage(stage)
										instance = any(instanceExportToMusicxml)
										__gong__map_ExportToMusicxml[identifier] = instanceExportToMusicxml
									case "FrontCurve":
										instanceFrontCurve := new(FrontCurve)
										instanceFrontCurve.Name = instanceName
										instanceFrontCurve.Stage(stage)
										instance = any(instanceFrontCurve)
										__gong__map_FrontCurve[identifier] = instanceFrontCurve
									case "FrontCurveStack":
										instanceFrontCurveStack := new(FrontCurveStack)
										instanceFrontCurveStack.Name = instanceName
										instanceFrontCurveStack.Stage(stage)
										instance = any(instanceFrontCurveStack)
										__gong__map_FrontCurveStack[identifier] = instanceFrontCurveStack
									case "HorizontalAxis":
										instanceHorizontalAxis := new(HorizontalAxis)
										instanceHorizontalAxis.Name = instanceName
										instanceHorizontalAxis.Stage(stage)
										instance = any(instanceHorizontalAxis)
										__gong__map_HorizontalAxis[identifier] = instanceHorizontalAxis
									case "Key":
										instanceKey := new(Key)
										instanceKey.Name = instanceName
										instanceKey.Stage(stage)
										instance = any(instanceKey)
										__gong__map_Key[identifier] = instanceKey
									case "Parameter":
										instanceParameter := new(Parameter)
										instanceParameter.Name = instanceName
										instanceParameter.Stage(stage)
										instance = any(instanceParameter)
										__gong__map_Parameter[identifier] = instanceParameter
									case "Rhombus":
										instanceRhombus := new(Rhombus)
										instanceRhombus.Name = instanceName
										instanceRhombus.Stage(stage)
										instance = any(instanceRhombus)
										__gong__map_Rhombus[identifier] = instanceRhombus
									case "RhombusGrid":
										instanceRhombusGrid := new(RhombusGrid)
										instanceRhombusGrid.Name = instanceName
										instanceRhombusGrid.Stage(stage)
										instance = any(instanceRhombusGrid)
										__gong__map_RhombusGrid[identifier] = instanceRhombusGrid
									case "ShapeCategory":
										instanceShapeCategory := new(ShapeCategory)
										instanceShapeCategory.Name = instanceName
										instanceShapeCategory.Stage(stage)
										instance = any(instanceShapeCategory)
										__gong__map_ShapeCategory[identifier] = instanceShapeCategory
									case "SpiralBezier":
										instanceSpiralBezier := new(SpiralBezier)
										instanceSpiralBezier.Name = instanceName
										instanceSpiralBezier.Stage(stage)
										instance = any(instanceSpiralBezier)
										__gong__map_SpiralBezier[identifier] = instanceSpiralBezier
									case "SpiralBezierGrid":
										instanceSpiralBezierGrid := new(SpiralBezierGrid)
										instanceSpiralBezierGrid.Name = instanceName
										instanceSpiralBezierGrid.Stage(stage)
										instance = any(instanceSpiralBezierGrid)
										__gong__map_SpiralBezierGrid[identifier] = instanceSpiralBezierGrid
									case "SpiralCircle":
										instanceSpiralCircle := new(SpiralCircle)
										instanceSpiralCircle.Name = instanceName
										instanceSpiralCircle.Stage(stage)
										instance = any(instanceSpiralCircle)
										__gong__map_SpiralCircle[identifier] = instanceSpiralCircle
									case "SpiralCircleGrid":
										instanceSpiralCircleGrid := new(SpiralCircleGrid)
										instanceSpiralCircleGrid.Name = instanceName
										instanceSpiralCircleGrid.Stage(stage)
										instance = any(instanceSpiralCircleGrid)
										__gong__map_SpiralCircleGrid[identifier] = instanceSpiralCircleGrid
									case "SpiralLine":
										instanceSpiralLine := new(SpiralLine)
										instanceSpiralLine.Name = instanceName
										instanceSpiralLine.Stage(stage)
										instance = any(instanceSpiralLine)
										__gong__map_SpiralLine[identifier] = instanceSpiralLine
									case "SpiralLineGrid":
										instanceSpiralLineGrid := new(SpiralLineGrid)
										instanceSpiralLineGrid.Name = instanceName
										instanceSpiralLineGrid.Stage(stage)
										instance = any(instanceSpiralLineGrid)
										__gong__map_SpiralLineGrid[identifier] = instanceSpiralLineGrid
									case "SpiralOrigin":
										instanceSpiralOrigin := new(SpiralOrigin)
										instanceSpiralOrigin.Name = instanceName
										instanceSpiralOrigin.Stage(stage)
										instance = any(instanceSpiralOrigin)
										__gong__map_SpiralOrigin[identifier] = instanceSpiralOrigin
									case "SpiralRhombus":
										instanceSpiralRhombus := new(SpiralRhombus)
										instanceSpiralRhombus.Name = instanceName
										instanceSpiralRhombus.Stage(stage)
										instance = any(instanceSpiralRhombus)
										__gong__map_SpiralRhombus[identifier] = instanceSpiralRhombus
									case "SpiralRhombusGrid":
										instanceSpiralRhombusGrid := new(SpiralRhombusGrid)
										instanceSpiralRhombusGrid.Name = instanceName
										instanceSpiralRhombusGrid.Stage(stage)
										instance = any(instanceSpiralRhombusGrid)
										__gong__map_SpiralRhombusGrid[identifier] = instanceSpiralRhombusGrid
									case "VerticalAxis":
										instanceVerticalAxis := new(VerticalAxis)
										instanceVerticalAxis.Name = instanceName
										instanceVerticalAxis.Stage(stage)
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
							log.Println("gongstructName not found for identifier", identifier)
							break
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
						case "Chapter":
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
						case "Content":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ExportToMusicxml":
							switch fieldName {
							// insertion point for date assign code
							}
						case "FrontCurve":
							switch fieldName {
							// insertion point for date assign code
							}
						case "FrontCurveStack":
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
						case "SpiralBezier":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SpiralBezierGrid":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SpiralCircle":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SpiralCircleGrid":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SpiralLine":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SpiralLineGrid":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SpiralOrigin":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SpiralRhombus":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SpiralRhombusGrid":
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
				// pick up the first arg
				if len(callExpr.Args) != 1 {
					break
				}
				arg0 := callExpr.Args[0]

				var se *ast.SelectorExpr
				var ok bool
				if se, ok = arg0.(*ast.SelectorExpr); !ok {
					break
				}

				var seXident *ast.Ident
				if seXident = se.X.(*ast.Ident); !ok {
					break
				}

				basicLit = new(ast.BasicLit)
				// For a "fake" literal, Kind might be set to something like token.STRING or a custom indicator
				basicLit.Kind = token.STRING // Or another appropriate token.Kind
				basicLit.Value = "new(" + seXident.Name + "." + se.Sel.Name + ")"
				// following lines are here to avoid warning "unused write to field..."
				_ = basicLit.Kind
				_ = basicLit.Value
				_ = basicLit
			}
			for argNb, arg := range callExpr.Args {
				_ = argNb
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident, *ast.SelectorExpr:
					var ident *ast.Ident
					var ok bool
					_ = ok
					if ident, ok = arg.(*ast.Ident); !ok {
						// log.Println("we are in the case of new(....)")
					}

					var se *ast.SelectorExpr
					if se, ok = arg.(*ast.SelectorExpr); ok {
						if ident, ok = se.X.(*ast.Ident); !ok {
							// log.Println("we are in the case of append(....)")
						}
					}
					_ = ident

					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Println("gongstructName not found for identifier", identifier)
						break
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
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Axis[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_AxisGrid[identifier]
								instanceWhoseFieldIsAppended.Axiss = append(instanceWhoseFieldIsAppended.Axiss, instanceToAppend)
							}
						}
					case "Bezier":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "BezierGrid":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Beziers":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Bezier[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_BezierGrid[identifier]
								instanceWhoseFieldIsAppended.Beziers = append(instanceWhoseFieldIsAppended.Beziers, instanceToAppend)
							}
						}
					case "BezierGridStack":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "BezierGrids":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_BezierGrid[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_BezierGridStack[identifier]
								instanceWhoseFieldIsAppended.BezierGrids = append(instanceWhoseFieldIsAppended.BezierGrids, instanceToAppend)
							}
						}
					case "Chapter":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Circle":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "CircleGrid":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Circles":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Circle[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_CircleGrid[identifier]
								instanceWhoseFieldIsAppended.Circles = append(instanceWhoseFieldIsAppended.Circles, instanceToAppend)
							}
						}
					case "Content":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Chapters":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Chapter[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Content[identifier]
								instanceWhoseFieldIsAppended.Chapters = append(instanceWhoseFieldIsAppended.Chapters, instanceToAppend)
							}
						}
					case "ExportToMusicxml":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "FrontCurve":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "FrontCurveStack":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "FrontCurves":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_FrontCurve[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_FrontCurveStack[identifier]
								instanceWhoseFieldIsAppended.FrontCurves = append(instanceWhoseFieldIsAppended.FrontCurves, instanceToAppend)
							}
						case "SpiralCircles":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SpiralCircle[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_FrontCurveStack[identifier]
								instanceWhoseFieldIsAppended.SpiralCircles = append(instanceWhoseFieldIsAppended.SpiralCircles, instanceToAppend)
							}
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
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Rhombus[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RhombusGrid[identifier]
								instanceWhoseFieldIsAppended.Rhombuses = append(instanceWhoseFieldIsAppended.Rhombuses, instanceToAppend)
							}
						}
					case "ShapeCategory":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SpiralBezier":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SpiralBezierGrid":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SpiralBeziers":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SpiralBezier[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_SpiralBezierGrid[identifier]
								instanceWhoseFieldIsAppended.SpiralBeziers = append(instanceWhoseFieldIsAppended.SpiralBeziers, instanceToAppend)
							}
						}
					case "SpiralCircle":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SpiralCircleGrid":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SpiralCircles":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SpiralCircle[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_SpiralCircleGrid[identifier]
								instanceWhoseFieldIsAppended.SpiralCircles = append(instanceWhoseFieldIsAppended.SpiralCircles, instanceToAppend)
							}
						}
					case "SpiralLine":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SpiralLineGrid":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SpiralLines":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SpiralLine[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_SpiralLineGrid[identifier]
								instanceWhoseFieldIsAppended.SpiralLines = append(instanceWhoseFieldIsAppended.SpiralLines, instanceToAppend)
							}
						}
					case "SpiralOrigin":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SpiralRhombus":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SpiralRhombusGrid":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SpiralRhombuses":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SpiralRhombus[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_SpiralRhombusGrid[identifier]
								instanceWhoseFieldIsAppended.SpiralRhombuses = append(instanceWhoseFieldIsAppended.SpiralRhombuses, instanceToAppend)
							}
						}
					case "VerticalAxis":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					}
				}
			}
		case *ast.BasicLit, *ast.UnaryExpr, *ast.CompositeLit:

			var basicLit *ast.BasicLit
			var exprSign = 1.0
			_ = exprSign // in case this is not used
			switch v := expr.(type) {
			case *ast.BasicLit:
				// expression is for instance ... = 18.000
				basicLit = v
			case *ast.UnaryExpr:
				// expression is for instance ... = -18.000
				// we want to extract a *ast.BasicLit from the *ast.UnaryExpr
				if bl, ok := v.X.(*ast.BasicLit); ok {
					basicLit = bl
					// Check the operator to set the sign
					if v.Op == token.SUB { // token.SUB is for '-'
						exprSign = -1
					} else if v.Op == token.ADD { // token.ADD is for '+'
						exprSign = 1
					}
				}
			case *ast.CompositeLit:
				var sl *ast.SelectorExpr
				var ident *ast.Ident
				var ok bool

				if sl, ok = v.Type.(*ast.SelectorExpr); !ok {
					break // Exits the switch case
				}

				if ident, ok = sl.X.(*ast.Ident); !ok {
					break // Exits the switch case
				}

				basicLit = new(ast.BasicLit)
				// For a "fake" literal, Kind might be set to something like token.STRING or a custom indicator
				basicLit.Kind = token.STRING // Or another appropriate token.Kind
				basicLit.Value = ident.Name + "." + sl.Sel.Name + "{}"
			}

			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Println("gongstructName not found for identifier", identifier)
				break
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
				case "AngleDegree":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Axis[identifier].AngleDegree = exprSign * fielValue
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
				case "EndX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Axis[identifier].EndX = exprSign * fielValue
				case "EndY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Axis[identifier].EndY = exprSign * fielValue
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
			case "Chapter":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Chapter[identifier].Name = fielValue
				case "MardownContent":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Chapter[identifier].MardownContent = fielValue
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
				case "Pitch":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].Pitch = int(exprSign) * int(fielValue)
				case "BeatNb":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].BeatNb = int(exprSign) * int(fielValue)
				}
			case "CircleGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CircleGrid[identifier].Name = fielValue
				}
			case "Content":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Content[identifier].Name = fielValue
				case "MardownContent":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Content[identifier].MardownContent = fielValue
				case "ContentPath":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Content[identifier].ContentPath = fielValue
				case "OutputPath":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Content[identifier].OutputPath = fielValue
				case "LayoutPath":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Content[identifier].LayoutPath = fielValue
				case "StaticPath":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Content[identifier].StaticPath = fielValue
				}
			case "ExportToMusicxml":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ExportToMusicxml[identifier].Name = fielValue
				}
			case "FrontCurve":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FrontCurve[identifier].Name = fielValue
				case "Path":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FrontCurve[identifier].Path = fielValue
				}
			case "FrontCurveStack":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FrontCurveStack[identifier].Name = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FrontCurveStack[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FrontCurveStack[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FrontCurveStack[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FrontCurveStack[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FrontCurveStack[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FrontCurveStack[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FrontCurveStack[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FrontCurveStack[identifier].Transform = fielValue
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
				case "BackendColor":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Parameter[identifier].BackendColor = fielValue
				case "MinuteColor":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Parameter[identifier].MinuteColor = fielValue
				case "HourColor":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Parameter[identifier].HourColor = fielValue
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
				case "ShiftToNearestCircle":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].ShiftToNearestCircle = int(exprSign) * int(fielValue)
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
				case "SpiralBezierStrength":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].SpiralBezierStrength = exprSign * fielValue
				case "NbInterpolationPoints":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].NbInterpolationPoints = int(exprSign) * int(fielValue)
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
				case "BeatLinesHeightRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].BeatLinesHeightRatio = exprSign * fielValue
				case "NbBeatLines":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].NbBeatLines = int(exprSign) * int(fielValue)
				case "NbOfBeatsInTheme":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].NbOfBeatsInTheme = int(exprSign) * int(fielValue)
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
				case "BeatsPerSecond":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].BeatsPerSecond = exprSign * fielValue
				case "Level":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].Level = exprSign * fielValue
				case "ThemeBinaryEncoding":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].ThemeBinaryEncoding = int(exprSign) * int(fielValue)
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
				case "SpiralOriginX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].SpiralOriginX = exprSign * fielValue
				case "SpiralOriginY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].SpiralOriginY = exprSign * fielValue
				case "OriginCrossWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].OriginCrossWidth = exprSign * fielValue
				case "SpiralRadiusRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].SpiralRadiusRatio = exprSign * fielValue
				case "ActualBeatsTemporalShift":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].ActualBeatsTemporalShift = int(exprSign) * int(fielValue)
				case "PathToStaticFiles":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Parameter[identifier].PathToStaticFiles = fielValue
				case "PathToGeneratedSVG":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Parameter[identifier].PathToGeneratedSVG = fielValue
				case "PathToGeneratedScore":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Parameter[identifier].PathToGeneratedScore = fielValue
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
				case "AngleDegree":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rhombus[identifier].AngleDegree = exprSign * fielValue
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
			case "SpiralBezier":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralBezier[identifier].Name = fielValue
				case "StartX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezier[identifier].StartX = exprSign * fielValue
				case "StartY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezier[identifier].StartY = exprSign * fielValue
				case "ControlPointStartX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezier[identifier].ControlPointStartX = exprSign * fielValue
				case "ControlPointStartY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezier[identifier].ControlPointStartY = exprSign * fielValue
				case "EndX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezier[identifier].EndX = exprSign * fielValue
				case "EndY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezier[identifier].EndY = exprSign * fielValue
				case "ControlPointEndX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezier[identifier].ControlPointEndX = exprSign * fielValue
				case "ControlPointEndY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezier[identifier].ControlPointEndY = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralBezier[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezier[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralBezier[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezier[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezier[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralBezier[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralBezier[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralBezier[identifier].Transform = fielValue
				}
			case "SpiralBezierGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralBezierGrid[identifier].Name = fielValue
				}
			case "SpiralCircle":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralCircle[identifier].Name = fielValue
				case "CenterX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralCircle[identifier].CenterX = exprSign * fielValue
				case "CenterY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralCircle[identifier].CenterY = exprSign * fielValue
				case "BespopkeRadius":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralCircle[identifier].BespopkeRadius = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralCircle[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralCircle[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralCircle[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralCircle[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralCircle[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralCircle[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralCircle[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralCircle[identifier].Transform = fielValue
				case "Pitch":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralCircle[identifier].Pitch = int(exprSign) * int(fielValue)
				case "BeatNb":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralCircle[identifier].BeatNb = int(exprSign) * int(fielValue)
				case "Path":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralCircle[identifier].Path = fielValue
				}
			case "SpiralCircleGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralCircleGrid[identifier].Name = fielValue
				}
			case "SpiralLine":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralLine[identifier].Name = fielValue
				case "StartX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralLine[identifier].StartX = exprSign * fielValue
				case "EndX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralLine[identifier].EndX = exprSign * fielValue
				case "StartY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralLine[identifier].StartY = exprSign * fielValue
				case "EndY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralLine[identifier].EndY = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralLine[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralLine[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralLine[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralLine[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralLine[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralLine[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralLine[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralLine[identifier].Transform = fielValue
				}
			case "SpiralLineGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralLineGrid[identifier].Name = fielValue
				}
			case "SpiralOrigin":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralOrigin[identifier].Name = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralOrigin[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralOrigin[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralOrigin[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralOrigin[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralOrigin[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralOrigin[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralOrigin[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralOrigin[identifier].Transform = fielValue
				}
			case "SpiralRhombus":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralRhombus[identifier].Name = fielValue
				case "X_r0":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombus[identifier].X_r0 = exprSign * fielValue
				case "Y_r0":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombus[identifier].Y_r0 = exprSign * fielValue
				case "X_r1":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombus[identifier].X_r1 = exprSign * fielValue
				case "Y_r1":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombus[identifier].Y_r1 = exprSign * fielValue
				case "X_r2":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombus[identifier].X_r2 = exprSign * fielValue
				case "Y_r2":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombus[identifier].Y_r2 = exprSign * fielValue
				case "X_r3":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombus[identifier].X_r3 = exprSign * fielValue
				case "Y_r3":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombus[identifier].Y_r3 = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralRhombus[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombus[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralRhombus[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombus[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombus[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralRhombus[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralRhombus[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralRhombus[identifier].Transform = fielValue
				}
			case "SpiralRhombusGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SpiralRhombusGrid[identifier].Name = fielValue
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
				log.Println("gongstructName not found for identifier", identifier)
				break
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
			case "Chapter":
				switch fieldName {
				// insertion point for field dependant code
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
				case "ShowName":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].ShowName = fielValue
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
			case "Content":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ExportToMusicxml":
				switch fieldName {
				// insertion point for field dependant code
				case "Parameter":
					targetIdentifier := ident.Name
					__gong__map_ExportToMusicxml[identifier].Parameter = __gong__map_Parameter[targetIdentifier]
				}
			case "FrontCurve":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "FrontCurveStack":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FrontCurveStack[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_FrontCurveStack[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
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
				case "GrowthCurveSeed":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].GrowthCurveSeed = __gong__map_Bezier[targetIdentifier]
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
				case "SpiralRhombusGridSeed":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralRhombusGridSeed = __gong__map_SpiralRhombus[targetIdentifier]
				case "SpiralRhombusGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralRhombusGrid = __gong__map_SpiralRhombusGrid[targetIdentifier]
				case "SpiralCircleSeed":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralCircleSeed = __gong__map_SpiralCircle[targetIdentifier]
				case "SpiralCircleGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralCircleGrid = __gong__map_SpiralCircleGrid[targetIdentifier]
				case "SpiralCircleFullGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralCircleFullGrid = __gong__map_SpiralCircleGrid[targetIdentifier]
				case "SpiralConstructionOuterLineSeed":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralConstructionOuterLineSeed = __gong__map_SpiralLine[targetIdentifier]
				case "SpiralConstructionInnerLineSeed":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralConstructionInnerLineSeed = __gong__map_SpiralLine[targetIdentifier]
				case "SpiralConstructionOuterLineGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralConstructionOuterLineGrid = __gong__map_SpiralLineGrid[targetIdentifier]
				case "SpiralConstructionInnerLineGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralConstructionInnerLineGrid = __gong__map_SpiralLineGrid[targetIdentifier]
				case "SpiralConstructionCircleGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralConstructionCircleGrid = __gong__map_SpiralCircleGrid[targetIdentifier]
				case "SpiralConstructionOuterLineFullGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralConstructionOuterLineFullGrid = __gong__map_SpiralLineGrid[targetIdentifier]
				case "SpiralBezierSeed":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralBezierSeed = __gong__map_SpiralBezier[targetIdentifier]
				case "SpiralBezierGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralBezierGrid = __gong__map_SpiralBezierGrid[targetIdentifier]
				case "SpiralBezierFullGrid":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralBezierFullGrid = __gong__map_SpiralBezierGrid[targetIdentifier]
				case "FrontCurveStack":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].FrontCurveStack = __gong__map_FrontCurveStack[targetIdentifier]
				case "Fkey":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].Fkey = __gong__map_Key[targetIdentifier]
				case "PitchLines":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].PitchLines = __gong__map_AxisGrid[targetIdentifier]
				case "BeatLines":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].BeatLines = __gong__map_AxisGrid[targetIdentifier]
				case "FirstVoice":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].FirstVoice = __gong__map_BezierGrid[targetIdentifier]
				case "FirstVoiceShiftedRigth":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].FirstVoiceShiftedRigth = __gong__map_BezierGrid[targetIdentifier]
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
				case "SpiralOrigin":
					targetIdentifier := ident.Name
					__gong__map_Parameter[identifier].SpiralOrigin = __gong__map_SpiralOrigin[targetIdentifier]
				case "ShowSpiralBezierConstruct":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].ShowSpiralBezierConstruct = fielValue
				case "ShowInterpolationPoints":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Parameter[identifier].ShowInterpolationPoints = fielValue
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
			case "SpiralBezier":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezier[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_SpiralBezier[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "SpiralBezierGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralBezierGrid[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_SpiralBezierGrid[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "SpiralCircle":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralCircle[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_SpiralCircle[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				case "HasBespokeRadius":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralCircle[identifier].HasBespokeRadius = fielValue
				case "ShowName":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralCircle[identifier].ShowName = fielValue
				}
			case "SpiralCircleGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralCircleGrid[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_SpiralCircleGrid[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				case "SpiralRhombusGrid":
					targetIdentifier := ident.Name
					__gong__map_SpiralCircleGrid[identifier].SpiralRhombusGrid = __gong__map_SpiralRhombusGrid[targetIdentifier]
				}
			case "SpiralLine":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralLine[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_SpiralLine[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "SpiralLineGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralLineGrid[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_SpiralLineGrid[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "SpiralOrigin":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralOrigin[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_SpiralOrigin[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "SpiralRhombus":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombus[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_SpiralRhombus[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
				}
			case "SpiralRhombusGrid":
				switch fieldName {
				// insertion point for field dependant code
				case "IsDisplayed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SpiralRhombusGrid[identifier].IsDisplayed = fielValue
				case "ShapeCategory":
					targetIdentifier := ident.Name
					__gong__map_SpiralRhombusGrid[identifier].ShapeCategory = __gong__map_ShapeCategory[targetIdentifier]
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
			var basicLit *ast.BasicLit
			var ident *ast.Ident

			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			case *ast.CompositeLit:
				var ok bool
				var sl *ast.SelectorExpr

				if sl, ok = X.Type.(*ast.SelectorExpr); !ok {
					break // Exits the switch case
				}

				if ident, ok = sl.X.(*ast.Ident); !ok {
					break // Exits the switch case
				}

				basicLit = new(ast.BasicLit)
				// For a "fake" literal, Kind might be set to something like token.STRING or a custom indicator
				basicLit.Kind = token.STRING // Or another appropriate token.Kind
				basicLit.Value = ident.Name + "." + sl.Sel.Name + "{}." + selectorExpr.Sel.Name
			}

			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Println("gongstructName not found for identifier", identifier)
					break
				}

				if basicLit == nil {
					// for the meta field written as ref_models.ENUM_VALUE1
					basicLit = new(ast.BasicLit)
					basicLit.Kind = token.STRING // Or another appropriate token.Kind
					basicLit.Value = selectorExpr.X.(*ast.Ident).Name + "." + Sel.Name
					_ = basicLit.Kind
					_ = basicLit.Value
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for selector expr assignments
				case "Axis":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "AxisGrid":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Bezier":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "BezierGrid":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "BezierGridStack":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Chapter":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Circle":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "CircleGrid":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Content":
					switch fieldName {
					// insertion point for selector expr assign code
					case "Target":
						var val Target
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Content[identifier].Target = Target(val)
					}
				case "ExportToMusicxml":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "FrontCurve":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "FrontCurveStack":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "HorizontalAxis":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Key":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Parameter":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Rhombus":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "RhombusGrid":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ShapeCategory":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SpiralBezier":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SpiralBezierGrid":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SpiralCircle":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SpiralCircleGrid":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SpiralLine":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SpiralLineGrid":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SpiralOrigin":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SpiralRhombus":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SpiralRhombusGrid":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "VerticalAxis":
					switch fieldName {
					// insertion point for selector expr assign code
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
