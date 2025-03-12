package engine

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Scene struct {
	root Node
}

func NewScene() *Scene {
	return &Scene{
		root: NewNode("root"),
	}
}

func (s *Scene) AddNode(node Node) {
	node.SetParent(s.root)
	s.root.AddChild(node)
}

func (s *Scene) Update() error {
	if err := s.root.Update(); err != nil {
		return err
	}

	colliders := GetAllNodesWithTag(s.root, "Collider")
	for i, collider := range colliders {
		for j := i + 1; j < len(colliders); j++ {
			otherCollider := colliders[j]
			if collider == otherCollider {
				continue
			}

			if circle, ok := collider.(*CircleCollider); ok {
				if otherCircle, ok := otherCollider.(*CircleCollider); ok {
					if circle.CollidesWithCircle(otherCircle) {
						circle.OnCollision.Emit(otherCircle)
						otherCircle.OnCollision.Emit(circle)
					}
				}
				if aabb, ok := otherCollider.(*AABBCollider); ok {
					if circle.CollidesWithAABB(aabb) {
						circle.OnCollision.Emit(aabb)
						aabb.OnCollision.Emit(circle)
					}
				}
			}
			if aabb, ok := collider.(*AABBCollider); ok {
				if otherCircle, ok := otherCollider.(*CircleCollider); ok {
					if aabb.CollidesWithCircle(otherCircle) {
						aabb.OnCollision.Emit(otherCircle)
						otherCircle.OnCollision.Emit(aabb)
					}
				}
				if otherAABB, ok := otherCollider.(*AABBCollider); ok {
					if aabb.CollidesWithAABB(otherAABB) {
						aabb.OnCollision.Emit(otherAABB)
						otherAABB.OnCollision.Emit(aabb)
					}
				}
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		s.PrintNodeTree()
	}

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	s.root.Draw(screen)
}

func (s *Scene) PrintNodeTree() {
	printNodeTree(s.root, 0)
}

func printNodeTree(node Node, indent int) {
	for _, child := range node.GetChildren() {
		fmt.Println(strings.Repeat(" ", indent), child.Name())
		printNodeTree(child, indent+2)
	}
}
