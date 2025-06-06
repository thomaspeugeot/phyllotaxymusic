// generated code - do not edit
package probe

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe      *Probe
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
	gongtreeStage *gongtree_models.Stage,
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
	// log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Axis" {
		updateAndCommitTable[models.Axis](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "AxisGrid" {
		updateAndCommitTable[models.AxisGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Bezier" {
		updateAndCommitTable[models.Bezier](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "BezierGrid" {
		updateAndCommitTable[models.BezierGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "BezierGridStack" {
		updateAndCommitTable[models.BezierGridStack](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Chapter" {
		updateAndCommitTable[models.Chapter](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Circle" {
		updateAndCommitTable[models.Circle](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CircleGrid" {
		updateAndCommitTable[models.CircleGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Content" {
		updateAndCommitTable[models.Content](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ExportToMusicxml" {
		updateAndCommitTable[models.ExportToMusicxml](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FrontCurve" {
		updateAndCommitTable[models.FrontCurve](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FrontCurveStack" {
		updateAndCommitTable[models.FrontCurveStack](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "HorizontalAxis" {
		updateAndCommitTable[models.HorizontalAxis](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Key" {
		updateAndCommitTable[models.Key](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Parameter" {
		updateAndCommitTable[models.Parameter](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Rhombus" {
		updateAndCommitTable[models.Rhombus](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RhombusGrid" {
		updateAndCommitTable[models.RhombusGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ShapeCategory" {
		updateAndCommitTable[models.ShapeCategory](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralBezier" {
		updateAndCommitTable[models.SpiralBezier](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralBezierGrid" {
		updateAndCommitTable[models.SpiralBezierGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralCircle" {
		updateAndCommitTable[models.SpiralCircle](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralCircleGrid" {
		updateAndCommitTable[models.SpiralCircleGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralLine" {
		updateAndCommitTable[models.SpiralLine](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralLineGrid" {
		updateAndCommitTable[models.SpiralLineGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralOrigin" {
		updateAndCommitTable[models.SpiralOrigin](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralRhombus" {
		updateAndCommitTable[models.SpiralRhombus](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralRhombusGrid" {
		updateAndCommitTable[models.SpiralRhombusGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "VerticalAxis" {
		updateAndCommitTable[models.VerticalAxis](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
