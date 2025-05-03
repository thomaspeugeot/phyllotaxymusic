package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type NodeTarget interface {
	OnAfterUpdateNode()
	GetGongtreeStage() *tree.Stage
}

// Generic node creation function
func NewNode(
	target NodeTarget,
	name string,
	isExpandedPointer *bool,
	isCheckedPointer *bool,
) *tree.Node {
	node := new(tree.Node).Stage(target.GetGongtreeStage())
	node.Name = name

	if isExpandedPointer != nil {
		node.IsExpanded = *isExpandedPointer
	}
	if isCheckedPointer != nil {
		node.IsChecked = *isCheckedPointer
	}

	proxy := NewNodeProxy(
		node,
		isExpandedPointer,
		isCheckedPointer,
		target,
	)

	node.Impl = proxy

	return node
}

// NewNodeProxy creates a new proxy for a node
func NewNodeProxy(
	node *tree.Node,
	isExpandedPointer *bool,
	isCheckedPointer *bool,
	target NodeTarget,
) *NodeProxy {
	proxy := new(NodeProxy)
	proxy.node = node
	proxy.isExpandedPointer = isExpandedPointer
	proxy.isCheckedPointer = isCheckedPointer
	proxy.target = target
	return proxy
}

// NodeProxy is a generic proxy for both int and float64
type NodeProxy struct {
	node              *tree.Node
	isExpandedPointer *bool
	isCheckedPointer  *bool
	target            NodeTarget
}

// Updated handles updating values when the node changes
func (proxy *NodeProxy) OnAfterUpdate(treeStage *tree.Stage, stageNode, frontNode *tree.Node) {

	if proxy.isExpandedPointer != nil {
		*proxy.isExpandedPointer = frontNode.IsExpanded
	}
	if proxy.isCheckedPointer != nil {
		*proxy.isCheckedPointer = frontNode.IsChecked
	}
	proxy.target.OnAfterUpdateNode()
}
