package __4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Node struct {
	treeNode *TreeNode
	parent   *Node
}

type Logic struct {
	maxIdx int
	nodes  map[int][]*Node
}

func newLogic() *Logic {
	return &Logic{
		nodes: make(map[int][]*Node),
	}
}

func (s *Logic) loop(idx int, p *Node) {
	if idx > s.maxIdx {
		s.maxIdx = idx
	}
	s.nodes[idx] = append(s.nodes[idx], p)
	if p.treeNode.Left != nil {
		s.loop(idx+1, &Node{treeNode: p.treeNode.Left, parent: p})
	}
	if p.treeNode.Right != nil {
		s.loop(idx+1, &Node{treeNode: p.treeNode.Right, parent: p})
	}
}

func (s *Logic) result() *TreeNode {
	nodes := s.nodes[s.maxIdx]
	for len(nodes) > 1 {
		m := make(map[*Node]struct{})
		for _, node := range nodes {
			m[node.parent] = struct{}{}
		}

		nodes = make([]*Node, 0, len(m))
		for node := range m {
			nodes = append(nodes, node)
		}
	}
	return nodes[0].treeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	logic := newLogic()
	logic.loop(0, &Node{treeNode: root})
	return logic.result()
}
