package ring

// 平衡二叉树 AVL树具有以下性质
// 它是一 棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树
// 注意此处的高度是子树的最深深度，因此出现两个叶子节点差距超过一层是符合要求的
type BalanceTree struct {
	root *TreeNode
}

func (tree *BalanceTree) Add(node_name string, val int) {
	tree.root = insert_node(tree.root, node_name, val)
}

func (tree *BalanceTree) Find(node_name string, val int) bool {
	if search_node(tree.root, node_name, val) != nil {
		return true
	}
	return false
}

func (tree *BalanceTree) Remove(node_name string, val int) bool {
	tree.root = remove_node(tree.root, node_name, val)
	return true
}

func (tree *BalanceTree) List() []string {
	return list_nodes(tree.root)
}
