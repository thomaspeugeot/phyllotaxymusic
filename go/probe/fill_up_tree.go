package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
)

const SideBarTreeName = "gong"

func fillUpTree(
	probe *Probe,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range probe.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	probe.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: SideBarTreeName}).Stage(probe.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](probe.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", probe.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&tree.Node{Name: name}).Stage(probe.treeStage)

		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}

		switch gongStruct.Name {
		// insertion point
		case "Axis":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Axis](probe.stageOfInterest)
			for _axis := range set {
				nodeInstance := (&tree.Node{Name: _axis.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_axis, "Axis", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "AxisGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.AxisGrid](probe.stageOfInterest)
			for _axisgrid := range set {
				nodeInstance := (&tree.Node{Name: _axisgrid.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_axisgrid, "AxisGrid", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Bezier":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Bezier](probe.stageOfInterest)
			for _bezier := range set {
				nodeInstance := (&tree.Node{Name: _bezier.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_bezier, "Bezier", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "BezierGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.BezierGrid](probe.stageOfInterest)
			for _beziergrid := range set {
				nodeInstance := (&tree.Node{Name: _beziergrid.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_beziergrid, "BezierGrid", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "BezierGridStack":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.BezierGridStack](probe.stageOfInterest)
			for _beziergridstack := range set {
				nodeInstance := (&tree.Node{Name: _beziergridstack.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_beziergridstack, "BezierGridStack", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Circle":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Circle](probe.stageOfInterest)
			for _circle := range set {
				nodeInstance := (&tree.Node{Name: _circle.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_circle, "Circle", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CircleGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CircleGrid](probe.stageOfInterest)
			for _circlegrid := range set {
				nodeInstance := (&tree.Node{Name: _circlegrid.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_circlegrid, "CircleGrid", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ExportToMusicxml":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ExportToMusicxml](probe.stageOfInterest)
			for _exporttomusicxml := range set {
				nodeInstance := (&tree.Node{Name: _exporttomusicxml.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_exporttomusicxml, "ExportToMusicxml", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FrontCurve":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FrontCurve](probe.stageOfInterest)
			for _frontcurve := range set {
				nodeInstance := (&tree.Node{Name: _frontcurve.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_frontcurve, "FrontCurve", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FrontCurveStack":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FrontCurveStack](probe.stageOfInterest)
			for _frontcurvestack := range set {
				nodeInstance := (&tree.Node{Name: _frontcurvestack.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_frontcurvestack, "FrontCurveStack", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "HorizontalAxis":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.HorizontalAxis](probe.stageOfInterest)
			for _horizontalaxis := range set {
				nodeInstance := (&tree.Node{Name: _horizontalaxis.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_horizontalaxis, "HorizontalAxis", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Key":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Key](probe.stageOfInterest)
			for _key := range set {
				nodeInstance := (&tree.Node{Name: _key.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_key, "Key", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Parameter":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Parameter](probe.stageOfInterest)
			for _parameter := range set {
				nodeInstance := (&tree.Node{Name: _parameter.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_parameter, "Parameter", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Rhombus":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Rhombus](probe.stageOfInterest)
			for _rhombus := range set {
				nodeInstance := (&tree.Node{Name: _rhombus.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rhombus, "Rhombus", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RhombusGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RhombusGrid](probe.stageOfInterest)
			for _rhombusgrid := range set {
				nodeInstance := (&tree.Node{Name: _rhombusgrid.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rhombusgrid, "RhombusGrid", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ShapeCategory":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ShapeCategory](probe.stageOfInterest)
			for _shapecategory := range set {
				nodeInstance := (&tree.Node{Name: _shapecategory.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_shapecategory, "ShapeCategory", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SpiralBezier":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SpiralBezier](probe.stageOfInterest)
			for _spiralbezier := range set {
				nodeInstance := (&tree.Node{Name: _spiralbezier.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spiralbezier, "SpiralBezier", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SpiralBezierGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SpiralBezierGrid](probe.stageOfInterest)
			for _spiralbeziergrid := range set {
				nodeInstance := (&tree.Node{Name: _spiralbeziergrid.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spiralbeziergrid, "SpiralBezierGrid", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SpiralCircle":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SpiralCircle](probe.stageOfInterest)
			for _spiralcircle := range set {
				nodeInstance := (&tree.Node{Name: _spiralcircle.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spiralcircle, "SpiralCircle", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SpiralCircleGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SpiralCircleGrid](probe.stageOfInterest)
			for _spiralcirclegrid := range set {
				nodeInstance := (&tree.Node{Name: _spiralcirclegrid.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spiralcirclegrid, "SpiralCircleGrid", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SpiralLine":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SpiralLine](probe.stageOfInterest)
			for _spiralline := range set {
				nodeInstance := (&tree.Node{Name: _spiralline.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spiralline, "SpiralLine", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SpiralLineGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SpiralLineGrid](probe.stageOfInterest)
			for _spirallinegrid := range set {
				nodeInstance := (&tree.Node{Name: _spirallinegrid.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spirallinegrid, "SpiralLineGrid", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SpiralOrigin":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SpiralOrigin](probe.stageOfInterest)
			for _spiralorigin := range set {
				nodeInstance := (&tree.Node{Name: _spiralorigin.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spiralorigin, "SpiralOrigin", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SpiralRhombus":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SpiralRhombus](probe.stageOfInterest)
			for _spiralrhombus := range set {
				nodeInstance := (&tree.Node{Name: _spiralrhombus.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spiralrhombus, "SpiralRhombus", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SpiralRhombusGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SpiralRhombusGrid](probe.stageOfInterest)
			for _spiralrhombusgrid := range set {
				nodeInstance := (&tree.Node{Name: _spiralrhombusgrid.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spiralrhombusgrid, "SpiralRhombusGrid", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "VerticalAxis":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.VerticalAxis](probe.stageOfInterest)
			for _verticalaxis := range set {
				nodeInstance := (&tree.Node{Name: _verticalaxis.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_verticalaxis, "VerticalAxis", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := (&tree.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(probe.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	// Add a refresh button
	nodeRefreshButton := (&tree.Node{Name: ""}).Stage(probe.treeStage)
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := (&tree.Button{
		Name: "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon: string(gongtree_buttons.BUTTON_refresh)}).Stage(probe.treeStage)
	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

	probe.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	probe          *Probe
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	probe *Probe) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.probe = probe
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *tree.StageStruct,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.probe,
	)
}
