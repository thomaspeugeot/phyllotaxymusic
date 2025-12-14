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
		updateProbeTable[*models.Axis](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "AxisGrid" {
		updateProbeTable[*models.AxisGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Bezier" {
		updateProbeTable[*models.Bezier](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "BezierGrid" {
		updateProbeTable[*models.BezierGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "BezierGridStack" {
		updateProbeTable[*models.BezierGridStack](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Chapter" {
		updateProbeTable[*models.Chapter](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Circle" {
		updateProbeTable[*models.Circle](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CircleGrid" {
		updateProbeTable[*models.CircleGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Content" {
		updateProbeTable[*models.Content](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ExportToMusicxml" {
		updateProbeTable[*models.ExportToMusicxml](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FrontCurve" {
		updateProbeTable[*models.FrontCurve](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FrontCurveStack" {
		updateProbeTable[*models.FrontCurveStack](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "HorizontalAxis" {
		updateProbeTable[*models.HorizontalAxis](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Key" {
		updateProbeTable[*models.Key](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Parameter" {
		updateProbeTable[*models.Parameter](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Rhombus" {
		updateProbeTable[*models.Rhombus](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RhombusGrid" {
		updateProbeTable[*models.RhombusGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ShapeCategory" {
		updateProbeTable[*models.ShapeCategory](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralBezier" {
		updateProbeTable[*models.SpiralBezier](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralBezierGrid" {
		updateProbeTable[*models.SpiralBezierGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralCircle" {
		updateProbeTable[*models.SpiralCircle](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralCircleGrid" {
		updateProbeTable[*models.SpiralCircleGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralLine" {
		updateProbeTable[*models.SpiralLine](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralLineGrid" {
		updateProbeTable[*models.SpiralLineGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralOrigin" {
		updateProbeTable[*models.SpiralOrigin](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralRhombus" {
		updateProbeTable[*models.SpiralRhombus](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SpiralRhombusGrid" {
		updateProbeTable[*models.SpiralRhombusGrid](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "VerticalAxis" {
		updateProbeTable[*models.VerticalAxis](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()
}
