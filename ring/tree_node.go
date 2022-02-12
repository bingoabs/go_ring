package ring

// 平衡二叉树 AVL树具有以下性质
// 它是一 棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树
// 注意此处的高度是子树的最深深度，因此出现两个叶子节点差距超过一层是符合要求的

type TreeNode struct {
	node_name    string
	val          int
	left         *TreeNode
	right        *TreeNode
	left_height  int
	right_height int
}

// 旋转操作分为以下四种情况
// 1）左左  2）左右  3）右左  4）右右
// 对于 “左左”和 “右右”的情况，只要进行 一次单旋转就可以使其恢复平衡
// 对于 “左右”和 “右左”则要进行一次 双旋转
// 对于左右和右左这两种情况，单旋转不能使它达到一个平衡状态，要经过两次旋转。双旋转是针对于这两种情况的解决方案
// 同样的，这样两种情况也是对称的，只要解决了左右这种情况，右左就很好办了
// 为使树恢复平衡，我们需要进行两步，第一步，把k1作为根，进行一次右右旋转，旋转之后就变成了左左情况，
// 所以第二步再进行一次左左旋转，最后得到了一棵以k2为根的平衡二叉树树
// 即 *左*右*结构执行*右*左*旋转

// 左旋转就是将左子节点提到根节点，而原先根节点成为右子节点

func insert_node(root *TreeNode, node_name string, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			node_name:    node_name,
			val:          val,
			left_height:  0,
			right_height: 0,
		}
	}
	if root.val < val {
		root.left = insert_node(root.left, node_name, val)
		root.left_height++
	} else {
		root.right = insert_node(root.right, node_name, val)
		root.right_height++
	}
	root = rebalance(root)
	return root
}

func search_node(root *TreeNode, node_name string, val int) *TreeNode {
	for {
		if root == nil {
			return nil
		}
		if root.val > val {
			root = root.left
		} else if root.val == val {
			// 基本不可能出现相同的哈希值，因为节点数量不应该多到会产生hash冲突的程度
			// if root.node_name != node_name {
			// 	panic("Rarely error, the hash value for different nodes is same!")
			// }
			return root
		} else {
			root = root.right
		}
	}
}

func remove_node(root *TreeNode, node_name string, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.val > val {
		remove_node(root.left, node_name, val)
		root = rebalance(root)
	} else if root.val == val {
		root = generate_root(root.left, root.right)
	} else {
		remove_node(root.right, node_name, val)
		root = rebalance(root)
	}
	return root
}

func rebalance(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	switch root.right_height - root.left_height {
	case 2: // 右子树不平衡
		// 检查子节点是否是右左结构，如果是，那么需要先执行一次左旋转,再执行一次右旋转
		// 否则只需要一次右旋转

		// 检查右子树的左子树是否高度较高，如果是，那么是右左结构，需要进行左右旋转
		if root.right.left_height > root.right.right_height {
			root.right = single_left_rotate(root.right)
		}
		root = single_right_rotate(root)
	case -2:
		// 左右结构，两次旋转，先右旋，然后左旋
		if root.left.right_height > root.left.left_height {
			root.left = single_right_rotate(root.left)
		} // 左左结构，一次左旋
		root = single_left_rotate(root)
	}
	return root
}
func get_smallest_node(node *TreeNode) *TreeNode {
	for {
		if node.right != nil {
			node = node.right
		} else {
			return node
		}
	}
}
func generate_root(left *TreeNode, right *TreeNode) *TreeNode {
	if left != nil && right != nil {
		smallest_node_in_right := get_smallest_node(right)
		root := &TreeNode{
			node_name:    smallest_node_in_right.node_name,
			val:          smallest_node_in_right.val,
			left_height:  0,
			right_height: 0,
		}
		right = remove_node(right, root.node_name, root.val)
		root.left = left
		root.right = right
		root = update_height(root)
		return rebalance(root)
	} else {
		var root *TreeNode = nil
		if left == nil {
			root = right
		} else if right == nil {
			root = left
		}
		return root
	}
}

func update_height(node *TreeNode) *TreeNode {
	if node.left == nil {
		node.left_height = 0
	} else {
		node.left_height = Max(node.left.left_height, node.left.right_height) + 1
	}
	if node.right == nil {
		node.right_height = 0
	} else {
		node.right_height = Max(node.right.left_height, node.right.right_height) + 1
	}
	return node
}

// 左旋即左子节点成为新的根节点，而原有根节点成为新的根节点的右节点
func single_left_rotate(root *TreeNode) *TreeNode {
	new_root := root.left
	root.left = new_root.right
	new_root.right = root

	update_height(root)
	update_height(new_root)
	return new_root
}

func single_right_rotate(root *TreeNode) *TreeNode {
	new_root := root.right
	root.right = new_root.left
	new_root.left = root

	update_height(root)
	update_height(new_root)
	return new_root
}

func is_balance(root *TreeNode) bool {
	return get_height(root) >= 0
}

// 如果子树平衡，返回子树高度，否则返回 -1
func get_height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left_height := get_height(root.left)
	right_height := get_height(root.right)
	if left_height == -1 || right_height == -1 || Abs(left_height-right_height) > 1 {
		return -1
	}
	return Max(left_height, right_height) + 1
}

func list_nodes(root *TreeNode) []string {
	// TODO golang queue
	queue
	type TreeNode struct {
		node_name    string
		val          int
		left         *TreeNode
		right        *TreeNode
		left_height  int
		right_height int
	}
}