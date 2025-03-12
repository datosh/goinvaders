package engine

import (
	"slices"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

// Node is the interface all nodes, i.e., objects in the game
// must implement.
type Node interface {
	Name() string
	UUID() uuid.UUID
	Update() error
	Draw(screen *ebiten.Image)
	SetParent(Node)
	GetParent() Node
	AddChild(Node)
	GetChildren() []Node
	AddTag(string)
	RemoveTag(string)
	HasTag(string) bool
}

// BaseNode is the default implementation of the Node interface,
// and helps to reduce the amount of boilerplate code needed
type BaseNode struct {
	name     string
	uuid     uuid.UUID
	tags     []string
	parent   Node
	children []Node
}

func NewNode(name string) *BaseNode {
	return &BaseNode{
		name:     name,
		uuid:     uuid.New(),
		tags:     []string{},
		children: []Node{},
	}
}

func (n *BaseNode) Name() string {
	return n.name
}

func (n *BaseNode) UUID() uuid.UUID {
	return n.uuid
}

func (n *BaseNode) Update() error {
	for _, child := range n.children {
		if err := child.Update(); err != nil {
			return err
		}
	}
	return nil
}

func (n *BaseNode) Draw(screen *ebiten.Image) {
	for _, child := range n.children {
		child.Draw(screen)
	}
}

func (n *BaseNode) SetParent(parent Node) {
	n.parent = parent
}

func (n *BaseNode) GetParent() Node {
	return n.parent
}

// AddChild adds a child node to the current node, which is made its parent.
func (n *BaseNode) AddChild(child Node) {
	child.SetParent(n)
	n.children = append(n.children, child)
}

func (n *BaseNode) GetChildren() []Node {
	return n.children
}

func (n *BaseNode) AddTag(tag string) {
	n.tags = append(n.tags, tag)
}

func (n *BaseNode) RemoveTag(tag string) {
	for i, t := range n.tags {
		if t == tag {
			n.tags = slices.Delete(n.tags, i, i+1)
			return
		}
	}
}

func (n *BaseNode) HasTag(tag string) bool {
	return slices.Contains(n.tags, tag)
}

func VisitAllNodes(node Node, visitor func(Node)) {
	visitor(node)
	for _, child := range node.GetChildren() {
		VisitAllNodes(child, visitor)
	}
}

func GetAllNodesWithTag(node Node, tag string) []Node {
	nodes := []Node{}
	VisitAllNodes(node, func(node Node) {
		if node.HasTag(tag) {
			nodes = append(nodes, node)
		}
	})
	return nodes
}
