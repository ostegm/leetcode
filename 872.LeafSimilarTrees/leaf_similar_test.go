package main

import (
	"log"
	"testing"
)

/* Test tree with nodes
* 		   3
*     	5  	  1
*    6      9  8
 */

// Leaf nodes
var lnode1 = TreeNode{6, nil, nil}
var lnode2 = TreeNode{9, nil, nil}
var lnode3 = TreeNode{8, nil, nil}

// Mid nodes
var mnode1 = TreeNode{5, &lnode1, nil}
var mnode2 = TreeNode{1, &lnode2, &lnode3}

// Roots
var root = TreeNode{3, &mnode1, &mnode2}
var rootOther = TreeNode{3, &mnode1, nil}

func TestGetLeaves(t *testing.T) {
	var leaves []int
	leaves = getLeaves(&root, leaves)
	log.Printf("Got leaves: %v", leaves)
	expectedLeaves := []int{6, 9, 8}
	if len(leaves) != len(expectedLeaves) {
		t.Errorf("Error, too many leaves in the output, expected 3, got: %v", leaves)
	}
	for i := range leaves {
		if leaves[i] != expectedLeaves[i] {
			t.Errorf("Got %v, wanted %v", leaves, expectedLeaves)
		}
	}
}

func TestLeafSimilarWhenSimilar(t *testing.T) {
	got := leafSimilar(&root, &root)
	if got != true {
		t.Errorf("Roots are similar, should return true")
	}

}

func TestLeafSimilarWhenNotSimilar(t *testing.T) {
	got := leafSimilar(&root, &rootOther)
	if got != false {
		t.Errorf("Roots are not similar, should return false")
	}

}
