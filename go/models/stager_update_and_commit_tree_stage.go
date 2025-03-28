package models

import (
	"cmp"
	"slices"

	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
)

// UpdateAndCommitTreeStage gathers shapes from the provided Stager, organizes
// them under their shape categories (if any), and commits the resulting tree
// structure to the treeStage.
//
//   - stager: The Stager struct containing a list of shapes to display.
//
// The workflow is as follows:
//  1. Reset the current GONG Tree stage to start from a clean state.
//  2. Create a new Tree entity to represent the sidebar or main tree.
//  3. Group all shapes by their shape categories (if they have one).
//  4. Sort shape categories by their Name.
//  5. For each shape category, create a Node, mark expansion state, and attach child shapes.
//  6. Any shapes without a category are attached at the root level.
//  7. Commit the updated tree to the stage, making it persistent.
func (stager *Stager) UpdateAndCommitTreeStage() {

	stager.treeStage.Reset()
	stager.tree = new(gongtree_models.Tree).Stage(stager.treeStage)
	stager.tree.Name = string(Sidebar)

	//
	// collect all shape categories and creates the map of shapes
	// by shape cateories
	//
	map_ShapeCategory_Shapes := make(map[*ShapeCategory][]ShapeInterface, 0)
	for _, shape := range stager.parameter.Shapes {
		shapeCategory := shape.GetShapeCategory()
		if shapeCategory == nil {
			continue
		}

		if map_ShapeCategory_Shapes[shapeCategory] == nil {
			map_ShapeCategory_Shapes[shapeCategory] = make([]ShapeInterface, 0)
		}
		map_ShapeCategory_Shapes[shapeCategory] = append(map_ShapeCategory_Shapes[shapeCategory], shape)
	}

	var shapeCategories []*ShapeCategory
	for k := range map_ShapeCategory_Shapes {
		shapeCategories = append(shapeCategories, k)
	}
	slices.SortFunc(shapeCategories, func(sc1, sc2 *ShapeCategory) int {
		return cmp.Compare(sc1.Name, sc2.Name)
	})

	for _, shapeCategory := range shapeCategories {
		shapes := map_ShapeCategory_Shapes[shapeCategory]
		node := new(gongtree_models.Node).Stage(stager.treeStage)
		node.Name = shapeCategory.GetName()

		if shapeCategory.IsExpanded {
			node.IsExpanded = true
		} else {
			node.IsExpanded = false
		}

		node.Impl = NewNodeProxy(
			node,
			&shapeCategory.IsExpanded,
			nil,
			stager,
		)

		stager.tree.RootNodes = append(stager.tree.RootNodes, node)
		for _, s := range shapes {
			addShapeNode(s, node, stager)
		}
	}

	for _, shape := range stager.parameter.Shapes {
		if shape.GetShapeCategory() == nil {
			addShapeNode(shape, nil, stager)
		}
	}

	stager.treeStage.Commit()
}

func addShapeNode(
	shape ShapeInterface,
	shapeCategoryNode *gongtree_models.Node,
	stager *Stager,
) {
	stager.addShapeNode(shape.GetName(), shape, shapeCategoryNode, shape.GetIsDisplayed())
}

// addShapeNode is an internal helper method of paramater that creates a GONG Tree node
// with the specified configuration and attaches it to the given parent node.
//
//   - name:  The display name of the node.
//   - impl:  The node implementation, which must implement NodeImplInterface.
//   - sc:    The parent node (or nil if attaching to the Tree root).
//   - isChecked: Whether the node’s checkbox should be initially checked.
func (stager *Stager) addShapeNode(
	name string,
	shape ShapeInterface,
	shapeCategoryNode *gongtree_models.Node,
	isChecked bool) {

	node := new(gongtree_models.Node).Stage(stager.treeStage)
	node.Name = name

	node.HasCheckboxButton = true
	node.IsChecked = isChecked

	node.Impl = NewNodeProxy(node,
		nil,
		shape.GetIsDisplayedPointer(),
		stager)

	if shapeCategoryNode != nil {
		shapeCategoryNode.Children = append(shapeCategoryNode.Children, node)
	} else {
		stager.tree.RootNodes = append(stager.tree.RootNodes, node)
	}
}
