package tree

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	models "github.com/thomaspeugeot/phylotaxymusic/go/models"
)

func (tree *Tree) Generate(parameter *models.Parameter) {

	tree.TreeStack.Stage.Reset()

	tree.NodeTree = new(gongtree_models.Tree).Stage(tree.TreeStack.Stage)
	tree.NodeTree.Name = string(models.Sidebar)

	//
	// collect all shape categories
	//
	map_ShapeCategory_Shapes := make(map[*models.ShapeCategory][]models.Shape, 0)
	for _, shape := range parameter.Shapes {
		sc := shape.GetShapeCategory()
		if sc == nil {
			continue
		}

		if map_ShapeCategory_Shapes[sc] == nil {
			map_ShapeCategory_Shapes[sc] = make([]models.Shape, 0)
		}
		map_ShapeCategory_Shapes[sc] = append(map_ShapeCategory_Shapes[sc], shape)
	}

	for k, shapes := range map_ShapeCategory_Shapes {
		node := new(gongtree_models.Node).Stage(tree.TreeStack.Stage)
		node.Name = k.GetName()

		tree.NodeTree.RootNodes = append(tree.NodeTree.RootNodes, node)
		for _, s := range shapes {
			AddShape(tree, s, node)
		}
	}

	for _, shape := range parameter.Shapes {
		if shape.GetShapeCategory() == nil {
			AddShape(tree, shape, nil)
		}
	}

	tree.TreeStack.Stage.Commit()
}

func AddShape[T models.Shape](tree *Tree, s T, shapeCategoryNode *gongtree_models.Node) {
	tree.addNode(s.GetName(), s, shapeCategoryNode, s.GetIsDisplayed())
}

func (tree *Tree) addNode(
	name string,
	impl gongtree_models.NodeImplInterface,
	sc *gongtree_models.Node,
	isChecked bool) {
	node := new(gongtree_models.Node).Stage(tree.TreeStack.Stage)
	node.Name = name

	node.IsNodeClickable = true
	node.IsCheckboxDisabled = false

	node.HasCheckboxButton = true
	node.IsChecked = isChecked

	node.Impl = impl

	if sc != nil {
		sc.Children = append(sc.Children, node)
	} else {
		tree.NodeTree.RootNodes = append(tree.NodeTree.RootNodes, node)
	}
}
