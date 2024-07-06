package tree

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"
)

func (tree *Tree) Generate(parameter *phylotaxymusic_models.Parameter) {

	tree.TreeStack.Stage.Reset()

	tree.NodeTree = new(gongtree_models.Tree).Stage(tree.TreeStack.Stage)
	tree.NodeTree.Name = string(phylotaxymusic_models.Sidebar)

	firstNode := new(gongtree_models.Node).Stage(tree.TreeStack.Stage)
	firstNode.Name = "Axis"
	// firstNode.IsWithPreceedingIcon = true
	// firstNode.PreceedingIcon = string(maticons.BUTTON_library_books)

	firstNode.IsNodeClickable = true
	firstNode.IsCheckboxDisabled = false

	firstNode.HasCheckboxButton = true
	firstNode.IsChecked = true

	firstNode.Impl = parameter.HorizontalAxis

	tree.NodeTree.RootNodes = append(tree.NodeTree.RootNodes, firstNode)

	tree.TreeStack.Stage.Commit()
}
