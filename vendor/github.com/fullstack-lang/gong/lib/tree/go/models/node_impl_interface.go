package models

type NodeImplInterface interface {

	// OnAfterUpdate function is called each time a node is modified
	OnAfterUpdate(stage *Stage, stagedNode, frontNode *Node)
}
