package ring

// 使用排序列表实现节点管理, nodes记录节点名称扩展名称

type SortArray struct {
	nodes []string
}

func (array *SortArray) Insert(val string) {
	if len(array.nodes) == 0 {
		array.nodes = append(array.nodes, val)
	} else {
		idx := find_first_bigger_value_idx(array.nodes, val)
		nodes := array.nodes[0:idx]
		nodes = append(nodes, val)
		nodes = append(nodes, array.nodes[idx:]...)
		array.nodes = nodes
	}
}

func (array *SortArray) Remove(val string) bool {
	for i := 0; i < len(array.nodes); i++ {
		if array.nodes[i] == val {
			array.nodes = append(array.nodes[:i], array.nodes[i+1:]...)
			return true
		}
	}
	return false
}

func find_first_bigger_value_idx(nodes []string, val string) int {
	left := 0
	right := len(nodes)
	mid := 0
	for {
		if left == right {
			mid = left
			break
		}
		mid := left + (right-left)/2
		if nodes[mid] > val {
			right = mid
		} else if nodes[mid] == val {
			break
		} else {
			left = mid + 1
		}
	}
	return mid
}
