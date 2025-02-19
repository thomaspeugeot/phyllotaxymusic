package models

import (
	tree "github.com/fullstack-lang/gongtree/go/models"
)

type Target interface {
	OnAfterUpdateNode()
	GetGongtreeStage() *tree.StageStruct
}

// Generic node creation function
func NewNode(
	target Target,
	name string,
	isExpanderPointer *bool,
) *tree.Node {
	node := new(tree.Node).Stage(target.GetGongtreeStage())
	node.Name = name
	if isExpanderPointer != nil {
		node.IsExpanded = *isExpanderPointer
	}

	proxy := NewNodeProxy(
		node,
		isExpanderPointer,
		target,
	)

	node.Impl = proxy

	return node
}

// NewNodeProxy creates a new proxy for a node
func NewNodeProxy(
	node *tree.Node,
	isExpandedPointer *bool,
	target Target,
) *NodeProxy {
	proxy := new(NodeProxy)
	proxy.node = node
	proxy.isExpandedPointer = isExpandedPointer
	proxy.target = target
	return proxy
}

// NodeProxy is a generic proxy for both int and float64
type NodeProxy struct {
	node              *tree.Node
	isExpandedPointer *bool
	target            Target
}

// Updated handles updating values when the node changes
func (proxy *NodeProxy) OnAfterUpdate(treeStage *tree.StageStruct, stageNode, frontNode *tree.Node) {

	if proxy.isExpandedPointer != nil {
		*proxy.isExpandedPointer = frontNode.IsExpanded
	}
	proxy.target.OnAfterUpdateNode()
}
