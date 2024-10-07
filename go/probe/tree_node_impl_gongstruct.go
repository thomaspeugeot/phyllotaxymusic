// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/thomaspeugeot/phylotaxymusic/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe *Probe
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	probe *Probe,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.probe = probe
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Axis" {
		fillUpTable[models.Axis](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "AxisGrid" {
		fillUpTable[models.AxisGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Bezier" {
		fillUpTable[models.Bezier](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "BezierGrid" {
		fillUpTable[models.BezierGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "BezierGridStack" {
		fillUpTable[models.BezierGridStack](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Circle" {
		fillUpTable[models.Circle](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CircleGrid" {
		fillUpTable[models.CircleGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "HorizontalAxis" {
		fillUpTable[models.HorizontalAxis](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Key" {
		fillUpTable[models.Key](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "NoteInfo" {
		fillUpTable[models.NoteInfo](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Parameter" {
		fillUpTable[models.Parameter](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Rhombus" {
		fillUpTable[models.Rhombus](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RhombusGrid" {
		fillUpTable[models.RhombusGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ShapeCategory" {
		fillUpTable[models.ShapeCategory](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralBezier" {
		fillUpTable[models.SpiralBezier](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralBezierGrid" {
		fillUpTable[models.SpiralBezierGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralCircle" {
		fillUpTable[models.SpiralCircle](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralCircleGrid" {
		fillUpTable[models.SpiralCircleGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralLine" {
		fillUpTable[models.SpiralLine](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralLineGrid" {
		fillUpTable[models.SpiralLineGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralRhombus" {
		fillUpTable[models.SpiralRhombus](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralRhombusGrid" {
		fillUpTable[models.SpiralRhombusGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "VerticalAxis" {
		fillUpTable[models.VerticalAxis](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
