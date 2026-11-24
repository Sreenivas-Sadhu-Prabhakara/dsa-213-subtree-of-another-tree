package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// IsSubtree reports whether sub appears somewhere in root as a complete
// subtree — same structure AND same leaves, not merely a matching prefix.
func IsSubtree(root, sub *TreeNode) bool {
	if sub == nil {
		return true // the empty tree is a subtree of anything
	}
	if root == nil {
		return false
	}
	return identical(root, sub) || IsSubtree(root.Left, sub) || IsSubtree(root.Right, sub)
}

func identical(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && identical(a.Left, b.Left) && identical(a.Right, b.Right)
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(3, tn(4, tn(1, nil, nil), tn(2, nil, nil)), tn(5, nil, nil))
	sub := tn(4, tn(1, nil, nil), tn(2, nil, nil))
	fmt.Println(IsSubtree(root, sub))
}
