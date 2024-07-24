package tree

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"
)

func (tree *Tree) Generate(parameter *phylotaxymusic_models.Parameter) {

	tree.TreeStack.Stage.Reset()

	tree.NodeTree = new(gongtree_models.Tree).Stage(tree.TreeStack.Stage)
	tree.NodeTree.Name = string(phylotaxymusic_models.Sidebar)

	AddShape(tree, parameter.HorizontalAxis)
	AddShape(tree, parameter.VerticalAxis)
	AddShape(tree, parameter.InitialRhombus)
	AddShape(tree, parameter.InitialCircle)
	AddShape(tree, parameter.InitialRhombusGrid)
	AddShape(tree, parameter.InitialCircleGrid)
	AddShape(tree, parameter.InitialAxis)
	AddShape(tree, parameter.RotatedAxis)
	AddShape(tree, parameter.RotatedRhombus)
	AddShape(tree, parameter.RotatedRhombusGrid)
	AddShape(tree, parameter.RotatedCircleGrid)
	AddShape(tree, parameter.NextRhombus)
	AddShape(tree, parameter.NextCircle)
	AddShape(tree, parameter.GrowingRhombusGrid)
	AddShape(tree, parameter.GrowingCircleGrid)
	AddShape(tree, parameter.GrowingCircleGridLeft)
	AddShape(tree, parameter.ConstructionAxis)

	tree.TreeStack.Stage.Commit()
}

type Shape interface {
	gongtree_models.NodeImplInterface
	GetName() string
	GetIsDisplayed() bool
}

func AddShape[T Shape](tree *Tree, s T) {
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
