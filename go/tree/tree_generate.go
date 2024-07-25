package tree

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"
)

func (tree *Tree) Generate(parameter *phylotaxymusic_models.Parameter) {

	tree.TreeStack.Stage.Reset()

	tree.NodeTree = new(gongtree_models.Tree).Stage(tree.TreeStack.Stage)
	tree.NodeTree.Name = string(phylotaxymusic_models.Sidebar)

	for _, shape := range parameter.Shapes {
		AddShape(tree, shape)
	}

	tree.TreeStack.Stage.Commit()
}

func AddShape[T phylotaxymusic_models.Shape](tree *Tree, s T) {
	tree.addNode(s.GetName(), s, s.GetIsDisplayed())
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
