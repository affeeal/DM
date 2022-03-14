package main

//import "fmt"

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
	return &tree
}

func ReplaceNode(tree, x, y *AVLTreeNode) {

	if x == tree {

		*tree = *y
		if y != nil {

			y.parent = nil
		}
	} else {

		p := x.parent
		if y != nil {

			*(y.parent) = *p
		}
		if p.left == x {

			*(p.left) = *(y)
		} else {

			*(p.right) = *(y)
		}
	}
}

func RotateLeft(tree, x *AVLTreeNode) {

	y := x.right
	ReplaceNode(tree, x, y)
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

	var node *AVLTreeNode = tree
	for node != nil && node.key != s {

		if compareStrings(s, node.key) < 0 {

			node = node.left
		} else {

			node = node.right
		}
	}

	if node == nil {

		return 0, false
	}
	return node.value, true
}

func (tree *AVLTreeNode) Assign(s string, x int) {

	var newNode AVLTreeNode = InitAVLTreeNode(s, x)
	if tree.key == "" {

		*tree = newNode
	} else {

		var node *AVLTreeNode = tree
		for {

			if compareStrings(s, node.key) < 0 {

				if node.left == nil {

					node.left = &newNode
					newNode.parent = node
					break
				}
				node = node.left
			} else {

				if node.right == nil {

					node.right = &newNode
					newNode.parent = node
					break
				}
				node = node.right
			}
		}
	}

	var a *AVLTreeNode = &newNode
	for {

		var node *AVLTreeNode = a.parent
		if node == nil {

			break
		}
		if a == node.left {

			node.balance--
			if node.balance == 0 {

				break
			}
			if node.balance == -2 {

				if a.balance == 1 {

					RotateLeft(tree, a)
				}
				RotateRight(tree, node)
				break
			}
		} else {

			node.balance++
			if node.balance == 0 {

				break
			}
			if node.balance == 2 {

				if a.balance == -1 {

					RotateRight(tree, a)
				}
				RotateLeft(tree, node)
				break
			}			
		}
		a = node
	}
}