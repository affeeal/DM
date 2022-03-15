package main

var root *AVLTreeNode = nil

type AVLTreeNode struct {

	s string
	x, balance int
	left, right *AVLTreeNode
	parent *AVLTreeNode
}

func InitAVLTree() *AVLTreeNode {

	tree := AVLTreeNode { "", 0, 0, nil, nil, nil}
	root = &tree
	return &tree
}

func ReplaceNode(x, y *AVLTreeNode) {

	if x == root {

		root = y
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

func RotateLeft(x *AVLTreeNode) {

	y := x.right
	ReplaceNode(x, y)
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

func RotateRight(x *AVLTreeNode) {

	y := x.left
	ReplaceNode(x, y)
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

	n := root
	for n != nil && n.s != s {

		if compareStrings(s, n.s) < 0 {

			n = n.left
		} else {

			n = n.right
		}
	}
	if n == nil {

		return 0, false
	}
	return n.x, true
}

func (tree *AVLTreeNode) Assign(s string, x int) {

	n := AVLTreeNode { s, x, 0, nil, nil, nil}
	if root.s == "" {

		root = &n
	} else {

		p := root
		for {

			if compareStrings(s, p.s) < 0 {

				if p.left == nil {

					p.left = &n
					n.parent = p
					break
				}
				p = p.left
			} else {

				if p.right == nil {

					p.right = &n
					n.parent = p
					break
				}
				p = p.right
			}
		}
	}
	pn := &n
	for {

		p := pn.parent
		if p == nil {

			break
		}
		if pn == p.left {

			p.balance--
			if p.balance == 0 {

				break
			}
			if p.balance == -2 {

				if pn.balance == 1 {

					RotateLeft(pn)
				}
				RotateRight(p)
				break
			}
		} else {

			p.balance++
			if p.balance == 0 {

				break
			}
			if p.balance == 2 {

				if pn.balance == -1 {

					RotateRight(pn)
				}
				RotateLeft(p)
				break
			}			
		}
		pn = p
	}
}