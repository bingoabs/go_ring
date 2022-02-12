package ring

// 使用排序列表实现节点管理, nodes记录节点名称扩展名称

type SortArray struct {
	nodes []string
}

func (array *SortArray) Insert(val string) {
	if len(array.nodes) == 0 {
		array.nodes = append(array.nodes, val)
	} else {

	}
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
