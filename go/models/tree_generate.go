package models

import (
	"cmp"
	"slices"

	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

func (tree *Tree) Generate(parameter *Parameter) {

	tree.TreeStack.Stage.Reset()

	tree.NodeTree = new(gongtree_models.Tree).Stage(tree.TreeStack.Stage)
	tree.NodeTree.Name = string(Sidebar)

	//
	// collect all shape categories
	//
	map_ShapeCategory_Shapes := make(map[*ShapeCategory][]Shape, 0)
	for _, shape := range parameter.Shapes {
		sc := shape.GetShapeCategory()
		if sc == nil {
			continue
		}

		if map_ShapeCategory_Shapes[sc] == nil {
			map_ShapeCategory_Shapes[sc] = make([]Shape, 0)
		}
		map_ShapeCategory_Shapes[sc] = append(map_ShapeCategory_Shapes[sc], shape)
	}

	var shapeCategories []*ShapeCategory
	for k, _ := range map_ShapeCategory_Shapes {
		shapeCategories = append(shapeCategories, k)
	}
	slices.SortFunc(shapeCategories, func(sc1, sc2 *ShapeCategory) int {
		return cmp.Compare(sc1.Name, sc2.Name)
	})

	for _, sc := range shapeCategories {
		shapes := map_ShapeCategory_Shapes[sc]
		node := new(gongtree_models.Node).Stage(tree.TreeStack.Stage)
		node.Name = sc.GetName()

		if sc.IsExpanded {
			node.IsExpanded = true
		} else {
			node.IsExpanded = false
		}

		node.Impl = &shapeCategoryNodeImpl{
			sc:    sc,
			stage: tree.Stage,
		}

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

type shapeCategoryNodeImpl struct {
	sc    *ShapeCategory
	stage *StageStruct
}

func (shapeCategoryNodeImpl *shapeCategoryNodeImpl) OnAfterUpdate(
	stage *gongtree_models.StageStruct, stagedNode, frontNode *gongtree_models.Node,
) {
	if frontNode.IsExpanded == true {
		shapeCategoryNodeImpl.sc.IsExpanded = true
	} else {
		shapeCategoryNodeImpl.sc.IsExpanded = false
	}
	stagedNode.IsExpanded = frontNode.IsExpanded
	shapeCategoryNodeImpl.stage.Commit()
}

func AddShape[T Shape](tree *Tree, s T, shapeCategoryNode *gongtree_models.Node) {
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
