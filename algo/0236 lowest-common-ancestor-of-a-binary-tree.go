package algo

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	type TreeNodeWithParent struct {
		*TreeNode
		parent *TreeNodeWithParent
		depth  int
	}

	var (
		p_, q_  *TreeNodeWithParent
		inOrder func(node *TreeNode, parent *TreeNodeWithParent, depth int)
	)

	inOrder = func(node *TreeNode, parent *TreeNodeWithParent, depth int) {
		if node == nil {
			return
		}
		nodeWithParent := &TreeNodeWithParent{
			TreeNode: node,
			parent:   parent,
			depth:    depth,
		}
		inOrder(nodeWithParent.Left, nodeWithParent, depth+1)
		if node == p {
			p_ = nodeWithParent
		} else if node == q {
			q_ = nodeWithParent
		}
		inOrder(nodeWithParent.Right, nodeWithParent, depth+1)
	}

	inOrder(root, nil, 1)

	if p_.depth > q_.depth {
		p_, q_ = q_, p_
	}

	for p_.depth != q_.depth {
		q_ = q_.parent
	}

	for p_ != q_ {
		p_ = p_.parent
		q_ = q_.parent
	}

	return p_.TreeNode
}
