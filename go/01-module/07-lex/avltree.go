package main

//import "fmt"

var relevantTree *AVLTreeNode = nil

type AVLTreeNode struct {

	key string
	value, balance int
	left, right *AVLTreeNode
	parent *AVLTreeNode
}

func InitAVLTreeNode(s string, x int) AVLTreeNode {

	return AVLTreeNode { s, x, 0, nil, nil, nil}
}

func InitAVLTree() *AVLTreeNode {

	var tree AVLTreeNode = InitAVLTreeNode("", 0)
	relevantTree = &tree
	return &tree
}

func ReplaceNode(tree, x, y *AVLTreeNode) {

	if x == tree {

		relevantTree = y
		if y != nil {

			y.parent = nil
		}
	} else {

		p := x.parent
		if y != nil {

			y.parent = p
		}
		if p.left == x {

			p.left = y
		} else {

			p.right = y
		}
	}
}

func RotateLeft(tree, x *AVLTreeNode) {

	y := x.right
	ReplaceNode(tree, x, y)
	tree = relevantTree
	b := y.left
	if b != nil {

		b.parent = x
	}
	x.right = b
	x.parent = y
	y.left = x

	x.balance--
	if y.balance > 0 {

		x.balance -= y.balance
	}
	y.balance--
	if x.balance < 0 {

		y.balance += x.balance
	}
}

func RotateRight(tree, x *AVLTreeNode) {

	y := x.left
	ReplaceNode(tree, x, y)
	tree = relevantTree
	b := y.right
	if b != nil {

		b.parent = x
	}
	x.left = b
	x.parent = y
	y.right = x

	x.balance++
	if y.balance < 0 {

		x.balance -= y.balance
	}
	y.balance++
	if x.balance > 0 {

		y.balance += x.balance
	}
}

func (tree *AVLTreeNode) Lookup(s string) (x int, exists bool) {

	var n *AVLTreeNode = relevantTree
	for n != nil && n.key != s {

		if compareStrings(s, n.key) < 0 {

			n = n.left
		} else {

			n = n.right
		}
	}

	if n == nil {

		return 0, false
	}
	return n.value, true
}

func (tree *AVLTreeNode) Assign(s string, x int) {

	var a AVLTreeNode = InitAVLTreeNode(s, x)
	if relevantTree.key == "" {

		relevantTree = &a
	} else {

		var p *AVLTreeNode = relevantTree
		for {

			if compareStrings(s, p.key) < 0 {

				if p.left == nil {

					p.left = &a
					a.parent = p
					break
				}
				p = p.left
			} else {

				if p.right == nil {

					p.right = &a
					a.parent = p
					break
				}
				p = p.right
			}
		}
	}

	var pa *AVLTreeNode = &a
	for {

		var p *AVLTreeNode = pa.parent
		if p == nil {

			break
		}
		if pa == p.left {

			p.balance--
			if p.balance == 0 {

				break
			}
			if p.balance == -2 {

				if pa.balance == 1 {

					RotateLeft(relevantTree, pa)
				}
				RotateRight(relevantTree, p)
				break
			}
		} else {

			p.balance++
			if p.balance == 0 {

				break
			}
			if p.balance == 2 {

				if pa.balance == -1 {

					RotateRight(relevantTree, pa)
				}
				RotateLeft(relevantTree, p)
				break
			}			
		}
		pa = p
	}
}