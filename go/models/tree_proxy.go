package models

import (
	"cmp"
	"log"
	"slices"

	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

// TreeProxy, is for construing the gongtree Tree
type TreeProxy struct {
	PhyllotaxyStage *StageStruct

	gongtreeStage *gongtree_models.StageStruct
	Tree          *gongtree_models.Tree
}

func (tree *TreeProxy) SetGongtreeStage(gongtreeStage *gongtree_models.StageStruct) {
	tree.gongtreeStage = gongtreeStage
}

// UpdateAndCommitTreeStage gathers shapes from the provided Parameter, organizes
// them under their shape categories (if any), and commits the resulting tree
// structure to the gongtreeStage.
//
//   - parameter: The Parameter struct containing a list of shapes to display.
//
// The workflow is as follows:
//  1. Reset the current GONG Tree stage to start from a clean state.
//  2. Create a new Tree entity to represent the sidebar or main tree.
//  3. Group all shapes by their shape categories (if they have one).
//  4. Sort shape categories by their Name.
//  5. For each shape category, create a Node, mark expansion state, and attach child shapes.
//  6. Any shapes without a category are attached at the root level.
//  7. Commit the updated tree to the stage, making it persistent.
func (treeProxy *TreeProxy) UpdateAndCommitTreeStage() {

	// get the only diagram
	parameters := GetGongstructInstancesMap[Parameter](treeProxy.PhyllotaxyStage)
	parameter, ok := (*parameters)["Reference"]
	if !ok {
		log.Fatal("no Reference parameter on stage")
	}

	if len(*parameters) == 0 {
		log.Println("")
		// log.Fatalln("")
	}

	treeProxy.gongtreeStage.Reset()
	treeProxy.Tree = new(gongtree_models.Tree).Stage(treeProxy.gongtreeStage)
	treeProxy.Tree.Name = string(Sidebar)

	//
	// collect all shape categories and creates the map of shapes
	// by shape cateories
	//
	map_ShapeCategory_Shapes := make(map[*ShapeCategory][]ShapeInterface, 0)
	for _, shape := range parameter.Shapes {
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
		node := new(gongtree_models.Node).Stage(treeProxy.gongtreeStage)
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
			parameter,
		)

		treeProxy.Tree.RootNodes = append(treeProxy.Tree.RootNodes, node)
		for _, s := range shapes {
			addShapeNode(treeProxy, s, node, parameter)
		}
	}

	for _, shape := range parameter.Shapes {
		if shape.GetShapeCategory() == nil {
			addShapeNode(treeProxy, shape, nil, parameter)
		}
	}

	treeProxy.gongtreeStage.Commit()
}

func addShapeNode(
	treeProxy *TreeProxy,
	shape ShapeInterface,
	shapeCategoryNode *gongtree_models.Node,
	parameter *Parameter,
) {
	treeProxy.addShapeNode(shape.GetName(), shape, shapeCategoryNode, shape.GetIsDisplayed(), parameter)
}

// addShapeNode is an internal helper method of TreeProxy that creates a GONG Tree node
// with the specified configuration and attaches it to the given parent node.
//
//   - name:  The display name of the node.
//   - impl:  The node implementation, which must implement NodeImplInterface.
//   - sc:    The parent node (or nil if attaching to the Tree root).
//   - isChecked: Whether the node’s checkbox should be initially checked.
func (treeProxy *TreeProxy) addShapeNode(
	name string,
	shape ShapeInterface,
	shapeCategoryNode *gongtree_models.Node,
	isChecked bool,
	parameter *Parameter) {

	node := new(gongtree_models.Node).Stage(treeProxy.gongtreeStage)
	node.Name = name

	node.IsNodeClickable = true
	node.IsCheckboxDisabled = false

	node.HasCheckboxButton = true
	node.IsChecked = isChecked

	node.Impl = NewNodeProxy(node,
		nil,
		shape.GetIsDisplayedPointer(),
		parameter)

	if shapeCategoryNode != nil {
		shapeCategoryNode.Children = append(shapeCategoryNode.Children, node)
	} else {
		treeProxy.Tree.RootNodes = append(treeProxy.Tree.RootNodes, node)
	}
}
