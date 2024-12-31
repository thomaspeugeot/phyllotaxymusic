package models

import gongtree_models "github.com/fullstack-lang/gongtree/go/models"

// ShapeCategory organizes shapes into categories
type ShapeCategory struct {
	Name string

	IsExpanded bool
}

type shapeCategoryNodeImpl struct {
	shapeCategory   *ShapeCategory
	phyllotaxyStage *StageStruct
}

func (shapeCategoryNodeImpl *shapeCategoryNodeImpl) OnAfterUpdate(
	stage *gongtree_models.StageStruct, stagedNode, frontNode *gongtree_models.Node,
) {
	if frontNode.IsExpanded == true {
		shapeCategoryNodeImpl.shapeCategory.IsExpanded = true
	} else {
		shapeCategoryNodeImpl.shapeCategory.IsExpanded = false
	}
	stagedNode.IsExpanded = frontNode.IsExpanded
	shapeCategoryNodeImpl.phyllotaxyStage.Commit()
}
