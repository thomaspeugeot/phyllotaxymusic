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
	tree.addNode("Horizontal Axis", parameter.HorizontalAxis, parameter.HorizontalAxis.IsAxisDisplayed)
	tree.addNode("Vertical Axis", parameter.VerticalAxis, parameter.VerticalAxis.IsAxisDisplayed)

	tree.TreeStack.Stage.Commit()
}

func (tree *Tree) addNode(
	name string,
	impl gongtree_models.NodeImplInterface,
	isChecked bool) {
	firstNode := new(gongtree_models.Node).Stage(tree.TreeStack.Stage)
	firstNode.Name = name

	firstNode.IsNodeClickable = true
	firstNode.IsCheckboxDisabled = false

	firstNode.HasCheckboxButton = true
	firstNode.IsChecked = isChecked

	firstNode.Impl = impl

	tree.NodeTree.RootNodes = append(tree.NodeTree.RootNodes, firstNode)
}
