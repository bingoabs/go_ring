package ring

import "log"

// 平衡二叉树 AVL树具有以下性质
// 它是一 棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树
// 注意此处的高度是子树的最深深度，因此出现两个叶子节点差距超过一层是符合要求的
type BalanceTree struct {
	root *TreeNode
}

func (tree *BalanceTree) Add(node uint32) bool {
	log.Println("Balance tree Add function start")
	tree.root = insert_node(tree.root, node)
	log.Println("Balance tree Add function end")
	return true
}

// find the next bigger node value
func (tree *BalanceTree) Find(node uint32) uint32 {
	return search_first_bigger_node(tree.root, node)
}

func (tree *BalanceTree) Remove(node uint32) bool {
	tree.root = remove_node(tree.root, node)
	return true
}

func (tree *BalanceTree) List() []*TreeNode {
	return list_nodes(tree.root)
}

func (tree *BalanceTree) Layout() []string {
	log.Println("Balance Tree layout start")
	str_nodes := []string{}
	nodes := list_nodes(tree.root)
	for i := 0; i < len(nodes); i++ {
		str_nodes = append(str_nodes, nodes[i].ToString())
	}
	log.Println("Balance Tree layout end")
	return str_nodes
}
