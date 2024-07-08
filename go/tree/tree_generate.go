package tree

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"
)

func (tree *Tree) Generate(parameter *phylotaxymusic_models.Parameter) {

	tree.TreeStack.Stage.Reset()

	tree.NodeTree = new(gongtree_models.Tree).Stage(tree.TreeStack.Stage)
	tree.NodeTree.Name = string(phylotaxymusic_models.Sidebar)

	// firstNode.IsWithPreceedingIcon = true
	// firstNode.PreceedingIcon = string(maticons.BUTTON_library_books)
	tree.addNode("Horizontal Axis", parameter.HorizontalAxis, parameter.HorizontalAxis.IsDisplayed)
	tree.addNode("Vertical Axis", parameter.VerticalAxis, parameter.VerticalAxis.IsDisplayed)
	tree.addNode("Initial Rhombus", parameter.InitialRhombus, parameter.InitialRhombus.IsDisplayed)
	tree.addNode("Initial Circle", parameter.InitialCircle, parameter.InitialCircle.IsDisplayed)
	tree.addNode("Initial Rhombus grid", parameter.InitialRhombusGrid, parameter.InitialRhombusGrid.IsDisplayed)
	tree.addNode("Initial Circle grid", parameter.InitialCircleGrid, parameter.InitialCircleGrid.IsDisplayed)
	tree.addNode("Initial Axis", parameter.InitialAxis, parameter.InitialAxis.IsDisplayed)

	tree.TreeStack.Stage.Commit()
}

func (tree *Tree) addNode(
	name string,
	impl gongtree_models.NodeImplInterface,
	isChecked bool) {
	node := new(gongtree_models.Node).Stage(tree.TreeStack.Stage)
	node.Name = name

	node.IsNodeClickable = true
	node.IsCheckboxDisabled = false

	node.HasCheckboxButton = true
	node.IsChecked = isChecked

	node.Impl = impl

	tree.NodeTree.RootNodes = append(tree.NodeTree.RootNodes, node)
}
