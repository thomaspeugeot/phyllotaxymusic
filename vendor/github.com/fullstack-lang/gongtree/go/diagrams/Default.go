package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongtree/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Default models.StageStruct
var ___dummy__Time_Default time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Default ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Default map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.Button": &(ref_models.Button{}),

	"ref_models.Button.Icon": (ref_models.Button{}).Icon,

	"ref_models.Button.Name": (ref_models.Button{}).Name,

	"ref_models.Button.SVGIcon": (ref_models.Button{}).SVGIcon,

	"ref_models.FontStyleEnum": ref_models.FontStyleEnum(""),

	"ref_models.ITALIC": ref_models.ITALIC,

	"ref_models.NORMAL": ref_models.NORMAL,

	"ref_models.Node": &(ref_models.Node{}),

	"ref_models.Node.BackgroundColor": (ref_models.Node{}).BackgroundColor,

	"ref_models.Node.Buttons": (ref_models.Node{}).Buttons,

	"ref_models.Node.Children": (ref_models.Node{}).Children,

	"ref_models.Node.FontStyle": (ref_models.Node{}).FontStyle,

	"ref_models.Node.HasCheckboxButton": (ref_models.Node{}).HasCheckboxButton,

	"ref_models.Node.IsCheckboxDisabled": (ref_models.Node{}).IsCheckboxDisabled,

	"ref_models.Node.IsChecked": (ref_models.Node{}).IsChecked,

	"ref_models.Node.IsExpanded": (ref_models.Node{}).IsExpanded,

	"ref_models.Node.IsInEditMode": (ref_models.Node{}).IsInEditMode,

	"ref_models.Node.IsNodeClickable": (ref_models.Node{}).IsNodeClickable,

	"ref_models.Node.IsWithPreceedingIcon": (ref_models.Node{}).IsWithPreceedingIcon,

	"ref_models.Node.Name": (ref_models.Node{}).Name,

	"ref_models.Node.PreceedingIcon": (ref_models.Node{}).PreceedingIcon,

	"ref_models.Node.PreceedingSVGIcon": (ref_models.Node{}).PreceedingSVGIcon,

	"ref_models.SVGIcon": &(ref_models.SVGIcon{}),

	"ref_models.SVGIcon.Name": (ref_models.SVGIcon{}).Name,

	"ref_models.SVGIcon.SVG": (ref_models.SVGIcon{}).SVG,

	"ref_models.Tree": &(ref_models.Tree{}),

	"ref_models.Tree.Name": (ref_models.Tree{}).Name,

	"ref_models.Tree.RootNodes": (ref_models.Tree{}).RootNodes,

	"ref_models.TreeStackDefaultName": ref_models.TreeStackDefaultName,

	"ref_models.TreeStackName": ref_models.TreeStackName(""),
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default"] = DefaultInjection
// }

// DefaultInjection will stage objects of database "Default"
func DefaultInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_BackgroundColor := (&models.Field{Name: `BackgroundColor`}).Stage(stage)
	__Field__000001_FontStyle := (&models.Field{Name: `FontStyle`}).Stage(stage)
	__Field__000002_HasCheckboxButton := (&models.Field{Name: `HasCheckboxButton`}).Stage(stage)
	__Field__000003_Icon := (&models.Field{Name: `Icon`}).Stage(stage)
	__Field__000004_IsCheckboxDisabled := (&models.Field{Name: `IsCheckboxDisabled`}).Stage(stage)
	__Field__000005_IsChecked := (&models.Field{Name: `IsChecked`}).Stage(stage)
	__Field__000006_IsExpanded := (&models.Field{Name: `IsExpanded`}).Stage(stage)
	__Field__000007_IsInEditMode := (&models.Field{Name: `IsInEditMode`}).Stage(stage)
	__Field__000008_IsNodeClickable := (&models.Field{Name: `IsNodeClickable`}).Stage(stage)
	__Field__000009_IsWithPreceedingIcon := (&models.Field{Name: `IsWithPreceedingIcon`}).Stage(stage)
	__Field__000010_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000011_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000012_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000013_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000014_PreceedingIcon := (&models.Field{Name: `PreceedingIcon`}).Stage(stage)
	__Field__000015_SVG := (&models.Field{Name: `SVG`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Default_Button := (&models.GongStructShape{Name: `Default-Button`}).Stage(stage)
	__GongStructShape__000001_Default_Node := (&models.GongStructShape{Name: `Default-Node`}).Stage(stage)
	__GongStructShape__000002_Default_SVGIcon := (&models.GongStructShape{Name: `Default-SVGIcon`}).Stage(stage)
	__GongStructShape__000003_Default_Tree := (&models.GongStructShape{Name: `Default-Tree`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Buttons := (&models.Link{Name: `Buttons`}).Stage(stage)
	__Link__000001_Children := (&models.Link{Name: `Children`}).Stage(stage)
	__Link__000002_PreceedingSVGIcon := (&models.Link{Name: `PreceedingSVGIcon`}).Stage(stage)
	__Link__000003_RootNodes := (&models.Link{Name: `RootNodes`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Default_Button := (&models.Position{Name: `Pos-Default-Button`}).Stage(stage)
	__Position__000001_Pos_Default_Node := (&models.Position{Name: `Pos-Default-Node`}).Stage(stage)
	__Position__000002_Pos_Default_SVGIcon := (&models.Position{Name: `Pos-Default-SVGIcon`}).Stage(stage)
	__Position__000003_Pos_Default_Tree := (&models.Position{Name: `Pos-Default-Tree`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Node and Default-Button`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Node and Default-Node`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_SVGIcon := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Node and Default-SVGIcon`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Tree and Default-Node`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = true

	// Field values setup
	__Field__000000_BackgroundColor.Name = `BackgroundColor`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.BackgroundColor]
	__Field__000000_BackgroundColor.Identifier = `ref_models.Node.BackgroundColor`
	__Field__000000_BackgroundColor.FieldTypeAsString = ``
	__Field__000000_BackgroundColor.Structname = `Node`
	__Field__000000_BackgroundColor.Fieldtypename = `string`

	// Field values setup
	__Field__000001_FontStyle.Name = `FontStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.FontStyle]
	__Field__000001_FontStyle.Identifier = `ref_models.Node.FontStyle`
	__Field__000001_FontStyle.FieldTypeAsString = ``
	__Field__000001_FontStyle.Structname = `Node`
	__Field__000001_FontStyle.Fieldtypename = `FontStyleEnum`

	// Field values setup
	__Field__000002_HasCheckboxButton.Name = `HasCheckboxButton`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.HasCheckboxButton]
	__Field__000002_HasCheckboxButton.Identifier = `ref_models.Node.HasCheckboxButton`
	__Field__000002_HasCheckboxButton.FieldTypeAsString = ``
	__Field__000002_HasCheckboxButton.Structname = `Node`
	__Field__000002_HasCheckboxButton.Fieldtypename = `bool`

	// Field values setup
	__Field__000003_Icon.Name = `Icon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Button.Icon]
	__Field__000003_Icon.Identifier = `ref_models.Button.Icon`
	__Field__000003_Icon.FieldTypeAsString = ``
	__Field__000003_Icon.Structname = `Button`
	__Field__000003_Icon.Fieldtypename = `string`

	// Field values setup
	__Field__000004_IsCheckboxDisabled.Name = `IsCheckboxDisabled`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.IsCheckboxDisabled]
	__Field__000004_IsCheckboxDisabled.Identifier = `ref_models.Node.IsCheckboxDisabled`
	__Field__000004_IsCheckboxDisabled.FieldTypeAsString = ``
	__Field__000004_IsCheckboxDisabled.Structname = `Node`
	__Field__000004_IsCheckboxDisabled.Fieldtypename = `bool`

	// Field values setup
	__Field__000005_IsChecked.Name = `IsChecked`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.IsChecked]
	__Field__000005_IsChecked.Identifier = `ref_models.Node.IsChecked`
	__Field__000005_IsChecked.FieldTypeAsString = ``
	__Field__000005_IsChecked.Structname = `Node`
	__Field__000005_IsChecked.Fieldtypename = `bool`

	// Field values setup
	__Field__000006_IsExpanded.Name = `IsExpanded`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.IsExpanded]
	__Field__000006_IsExpanded.Identifier = `ref_models.Node.IsExpanded`
	__Field__000006_IsExpanded.FieldTypeAsString = ``
	__Field__000006_IsExpanded.Structname = `Node`
	__Field__000006_IsExpanded.Fieldtypename = `bool`

	// Field values setup
	__Field__000007_IsInEditMode.Name = `IsInEditMode`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.IsInEditMode]
	__Field__000007_IsInEditMode.Identifier = `ref_models.Node.IsInEditMode`
	__Field__000007_IsInEditMode.FieldTypeAsString = ``
	__Field__000007_IsInEditMode.Structname = `Node`
	__Field__000007_IsInEditMode.Fieldtypename = `bool`

	// Field values setup
	__Field__000008_IsNodeClickable.Name = `IsNodeClickable`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.IsNodeClickable]
	__Field__000008_IsNodeClickable.Identifier = `ref_models.Node.IsNodeClickable`
	__Field__000008_IsNodeClickable.FieldTypeAsString = ``
	__Field__000008_IsNodeClickable.Structname = `Node`
	__Field__000008_IsNodeClickable.Fieldtypename = `bool`

	// Field values setup
	__Field__000009_IsWithPreceedingIcon.Name = `IsWithPreceedingIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.IsWithPreceedingIcon]
	__Field__000009_IsWithPreceedingIcon.Identifier = `ref_models.Node.IsWithPreceedingIcon`
	__Field__000009_IsWithPreceedingIcon.FieldTypeAsString = ``
	__Field__000009_IsWithPreceedingIcon.Structname = `Node`
	__Field__000009_IsWithPreceedingIcon.Fieldtypename = `bool`

	// Field values setup
	__Field__000010_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVGIcon.Name]
	__Field__000010_Name.Identifier = `ref_models.SVGIcon.Name`
	__Field__000010_Name.FieldTypeAsString = ``
	__Field__000010_Name.Structname = `SVGIcon`
	__Field__000010_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000011_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Tree.Name]
	__Field__000011_Name.Identifier = `ref_models.Tree.Name`
	__Field__000011_Name.FieldTypeAsString = ``
	__Field__000011_Name.Structname = `Tree`
	__Field__000011_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000012_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Button.Name]
	__Field__000012_Name.Identifier = `ref_models.Button.Name`
	__Field__000012_Name.FieldTypeAsString = ``
	__Field__000012_Name.Structname = `Button`
	__Field__000012_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000013_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.Name]
	__Field__000013_Name.Identifier = `ref_models.Node.Name`
	__Field__000013_Name.FieldTypeAsString = ``
	__Field__000013_Name.Structname = `Node`
	__Field__000013_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000014_PreceedingIcon.Name = `PreceedingIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.PreceedingIcon]
	__Field__000014_PreceedingIcon.Identifier = `ref_models.Node.PreceedingIcon`
	__Field__000014_PreceedingIcon.FieldTypeAsString = ``
	__Field__000014_PreceedingIcon.Structname = `Node`
	__Field__000014_PreceedingIcon.Fieldtypename = `string`

	// Field values setup
	__Field__000015_SVG.Name = `SVG`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVGIcon.SVG]
	__Field__000015_SVG.Identifier = `ref_models.SVGIcon.SVG`
	__Field__000015_SVG.FieldTypeAsString = ``
	__Field__000015_SVG.Structname = `SVGIcon`
	__Field__000015_SVG.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_Default_Button.Name = `Default-Button`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Button]
	__GongStructShape__000000_Default_Button.Identifier = `ref_models.Button`
	__GongStructShape__000000_Default_Button.ShowNbInstances = true
	__GongStructShape__000000_Default_Button.NbInstances = 10
	__GongStructShape__000000_Default_Button.Width = 240.000000
	__GongStructShape__000000_Default_Button.Height = 93.000000
	__GongStructShape__000000_Default_Button.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Default_Node.Name = `Default-Node`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__GongStructShape__000001_Default_Node.Identifier = `ref_models.Node`
	__GongStructShape__000001_Default_Node.ShowNbInstances = true
	__GongStructShape__000001_Default_Node.NbInstances = 12
	__GongStructShape__000001_Default_Node.Width = 240.000000
	__GongStructShape__000001_Default_Node.Height = 412.000000
	__GongStructShape__000001_Default_Node.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Default_SVGIcon.Name = `Default-SVGIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVGIcon]
	__GongStructShape__000002_Default_SVGIcon.Identifier = `ref_models.SVGIcon`
	__GongStructShape__000002_Default_SVGIcon.ShowNbInstances = false
	__GongStructShape__000002_Default_SVGIcon.NbInstances = 0
	__GongStructShape__000002_Default_SVGIcon.Width = 240.000000
	__GongStructShape__000002_Default_SVGIcon.Height = 93.000000
	__GongStructShape__000002_Default_SVGIcon.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_Default_Tree.Name = `Default-Tree`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Tree]
	__GongStructShape__000003_Default_Tree.Identifier = `ref_models.Tree`
	__GongStructShape__000003_Default_Tree.ShowNbInstances = true
	__GongStructShape__000003_Default_Tree.NbInstances = 1
	__GongStructShape__000003_Default_Tree.Width = 240.000000
	__GongStructShape__000003_Default_Tree.Height = 78.000000
	__GongStructShape__000003_Default_Tree.IsSelected = false

	// Link values setup
	__Link__000000_Buttons.Name = `Buttons`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.Buttons]
	__Link__000000_Buttons.Identifier = `ref_models.Node.Buttons`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Button]
	__Link__000000_Buttons.Fieldtypename = `ref_models.Button`
	__Link__000000_Buttons.FieldOffsetX = -74.000000
	__Link__000000_Buttons.FieldOffsetY = -19.000000
	__Link__000000_Buttons.TargetMultiplicity = models.MANY
	__Link__000000_Buttons.TargetMultiplicityOffsetX = -36.000000
	__Link__000000_Buttons.TargetMultiplicityOffsetY = 30.000000
	__Link__000000_Buttons.SourceMultiplicity = models.ZERO_ONE
	__Link__000000_Buttons.SourceMultiplicityOffsetX = -42.000000
	__Link__000000_Buttons.SourceMultiplicityOffsetY = -16.000000
	__Link__000000_Buttons.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Buttons.StartRatio = 0.824965
	__Link__000000_Buttons.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Buttons.EndRatio = 0.500000
	__Link__000000_Buttons.CornerOffsetRatio = -0.735764

	// Link values setup
	__Link__000001_Children.Name = `Children`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.Children]
	__Link__000001_Children.Identifier = `ref_models.Node.Children`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__Link__000001_Children.Fieldtypename = `ref_models.Node`
	__Link__000001_Children.FieldOffsetX = -73.000000
	__Link__000001_Children.FieldOffsetY = 22.000000
	__Link__000001_Children.TargetMultiplicity = models.MANY
	__Link__000001_Children.TargetMultiplicityOffsetX = -24.000000
	__Link__000001_Children.TargetMultiplicityOffsetY = -7.000000
	__Link__000001_Children.SourceMultiplicity = models.ZERO_ONE
	__Link__000001_Children.SourceMultiplicityOffsetX = -33.000000
	__Link__000001_Children.SourceMultiplicityOffsetY = 24.000000
	__Link__000001_Children.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Children.StartRatio = 0.451039
	__Link__000001_Children.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Children.EndRatio = 0.676558
	__Link__000001_Children.CornerOffsetRatio = -0.898264

	// Link values setup
	__Link__000002_PreceedingSVGIcon.Name = `PreceedingSVGIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.PreceedingSVGIcon]
	__Link__000002_PreceedingSVGIcon.Identifier = `ref_models.Node.PreceedingSVGIcon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVGIcon]
	__Link__000002_PreceedingSVGIcon.Fieldtypename = `ref_models.SVGIcon`
	__Link__000002_PreceedingSVGIcon.FieldOffsetX = -50.000000
	__Link__000002_PreceedingSVGIcon.FieldOffsetY = -16.000000
	__Link__000002_PreceedingSVGIcon.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_PreceedingSVGIcon.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_PreceedingSVGIcon.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_PreceedingSVGIcon.SourceMultiplicity = models.MANY
	__Link__000002_PreceedingSVGIcon.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_PreceedingSVGIcon.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_PreceedingSVGIcon.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_PreceedingSVGIcon.StartRatio = 0.500000
	__Link__000002_PreceedingSVGIcon.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_PreceedingSVGIcon.EndRatio = 0.500000
	__Link__000002_PreceedingSVGIcon.CornerOffsetRatio = 1.675861

	// Link values setup
	__Link__000003_RootNodes.Name = `RootNodes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Tree.RootNodes]
	__Link__000003_RootNodes.Identifier = `ref_models.Tree.RootNodes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__Link__000003_RootNodes.Fieldtypename = `ref_models.Node`
	__Link__000003_RootNodes.FieldOffsetX = -100.000000
	__Link__000003_RootNodes.FieldOffsetY = -19.000000
	__Link__000003_RootNodes.TargetMultiplicity = models.MANY
	__Link__000003_RootNodes.TargetMultiplicityOffsetX = -28.000000
	__Link__000003_RootNodes.TargetMultiplicityOffsetY = 24.000000
	__Link__000003_RootNodes.SourceMultiplicity = models.ZERO_ONE
	__Link__000003_RootNodes.SourceMultiplicityOffsetX = 30.000000
	__Link__000003_RootNodes.SourceMultiplicityOffsetY = 28.000000
	__Link__000003_RootNodes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_RootNodes.StartRatio = 0.500000
	__Link__000003_RootNodes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_RootNodes.EndRatio = 0.130564
	__Link__000003_RootNodes.CornerOffsetRatio = 1.255903

	// Position values setup
	__Position__000000_Pos_Default_Button.X = 11.000000
	__Position__000000_Pos_Default_Button.Y = 374.000000
	__Position__000000_Pos_Default_Button.Name = `Pos-Default-Button`

	// Position values setup
	__Position__000001_Pos_Default_Node.X = 460.999969
	__Position__000001_Pos_Default_Node.Y = 76.000000
	__Position__000001_Pos_Default_Node.Name = `Pos-Default-Node`

	// Position values setup
	__Position__000002_Pos_Default_SVGIcon.X = 454.999938
	__Position__000002_Pos_Default_SVGIcon.Y = 522.222260
	__Position__000002_Pos_Default_SVGIcon.Name = `Pos-Default-SVGIcon`

	// Position values setup
	__Position__000003_Pos_Default_Tree.X = 15.000000
	__Position__000003_Pos_Default_Tree.Y = 84.000000
	__Position__000003_Pos_Default_Tree.Name = `Pos-Default-Tree`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button.X = 424.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button.Y = 82.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button.Name = `Verticle in class diagram Default in middle between Default-Node and Default-Button`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node.X = 454.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node.Y = 112.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node.Name = `Verticle in class diagram Default in middle between Default-Node and Default-Node`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_SVGIcon.X = 815.999969
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_SVGIcon.Y = 550.611130
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_SVGIcon.Name = `Verticle in class diagram Default in middle between Default-Node and Default-SVGIcon`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node.X = 443.500000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node.Y = 110.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node.Name = `Verticle in class diagram Default in middle between Default-Tree and Default-Node`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Button)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Node)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Tree)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_SVGIcon)
	__GongStructShape__000000_Default_Button.Position = __Position__000000_Pos_Default_Button
	__GongStructShape__000000_Default_Button.Fields = append(__GongStructShape__000000_Default_Button.Fields, __Field__000012_Name)
	__GongStructShape__000000_Default_Button.Fields = append(__GongStructShape__000000_Default_Button.Fields, __Field__000003_Icon)
	__GongStructShape__000001_Default_Node.Position = __Position__000001_Pos_Default_Node
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000013_Name)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000001_FontStyle)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000000_BackgroundColor)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000006_IsExpanded)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000002_HasCheckboxButton)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000005_IsChecked)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000004_IsCheckboxDisabled)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000007_IsInEditMode)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000008_IsNodeClickable)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000009_IsWithPreceedingIcon)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000014_PreceedingIcon)
	__GongStructShape__000001_Default_Node.Links = append(__GongStructShape__000001_Default_Node.Links, __Link__000000_Buttons)
	__GongStructShape__000001_Default_Node.Links = append(__GongStructShape__000001_Default_Node.Links, __Link__000001_Children)
	__GongStructShape__000001_Default_Node.Links = append(__GongStructShape__000001_Default_Node.Links, __Link__000002_PreceedingSVGIcon)
	__GongStructShape__000002_Default_SVGIcon.Position = __Position__000002_Pos_Default_SVGIcon
	__GongStructShape__000002_Default_SVGIcon.Fields = append(__GongStructShape__000002_Default_SVGIcon.Fields, __Field__000010_Name)
	__GongStructShape__000002_Default_SVGIcon.Fields = append(__GongStructShape__000002_Default_SVGIcon.Fields, __Field__000015_SVG)
	__GongStructShape__000003_Default_Tree.Position = __Position__000003_Pos_Default_Tree
	__GongStructShape__000003_Default_Tree.Fields = append(__GongStructShape__000003_Default_Tree.Fields, __Field__000011_Name)
	__GongStructShape__000003_Default_Tree.Links = append(__GongStructShape__000003_Default_Tree.Links, __Link__000003_RootNodes)
	__Link__000000_Buttons.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button
	__Link__000001_Children.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node
	__Link__000002_PreceedingSVGIcon.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_SVGIcon
	__Link__000003_RootNodes.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node
}


