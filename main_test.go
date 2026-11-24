package main

import "testing"

func TestIsSubtree(t *testing.T) {
	root := tn(3, tn(4, tn(1, nil, nil), tn(2, nil, nil)), tn(5, nil, nil))
	sub := tn(4, tn(1, nil, nil), tn(2, nil, nil))
	if !IsSubtree(root, sub) {
		t.Error("expected subtree found")
	}

	// The candidate matches a prefix but the tree continues below: not a subtree.
	root2 := tn(3, tn(4, tn(1, nil, nil), tn(2, tn(0, nil, nil), nil)), tn(5, nil, nil))
	if IsSubtree(root2, sub) {
		t.Error("prefix match with extra child must fail")
	}

	if !IsSubtree(root, root) {
		t.Error("a tree is its own subtree")
	}
	if !IsSubtree(root, nil) {
		t.Error("empty tree is a subtree of anything")
	}
	if IsSubtree(nil, sub) {
		t.Error("nothing is a subtree of the empty tree")
	}
	// Value exists but shape differs.
	if IsSubtree(root, tn(4, tn(2, nil, nil), tn(1, nil, nil))) {
		t.Error("mirrored candidate must fail")
	}
}
