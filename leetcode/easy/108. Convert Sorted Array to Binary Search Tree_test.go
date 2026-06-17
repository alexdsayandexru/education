package easy

func sortedArrayToBST(nums []int) *TreeNode {
	node := calc(nil, nums)
	return node
}

func calc(node *TreeNode, nums []int) *TreeNode {
	if len(nums) == 0 {
		return node
	}
	if len(nums) == 1 {
		node = add(node, nums[0])
	} else {
		n := len(nums) / 2
		node = add(node, nums[n])

		l := nums[:n]
		r := nums[n+1:]

		calc(node, l)
		calc(node, r)
	}
	return node
}

func add(node *TreeNode, val int) *TreeNode {
	if node == nil {
		node = &TreeNode{
			Val: val,
		}
		return node
	} else {
		if val < node.Val && node.Left != nil {
			add(node.Left, val)
		} else if val >= node.Val && node.Right != nil {
			add(node.Right, val)
		} else if val < node.Val && node.Left == nil {
			node.Left = &TreeNode{
				Val: val,
			}
		} else if val >= node.Val && node.Right == nil {
			node.Right = &TreeNode{
				Val: val,
			}
		}
	}
	return node
}

/*func left(node *TreeNode, buff []int) []int {
	if node.Left != nil {
		node.Left.Val = node.Val + 1
		buff = left(node.Left, buff)
	}

	if node.Right != nil {
		node.Right.Val = node.Val + 1
		buff = left(node.Right, buff)
	}
	buff = append(buff, node.Val)
	return buff
}

func right(node *TreeNode, buff []int) []int {
	if node.Right != nil {
		buff = left(node.Right, buff)
	}
	buff = append(buff, node.Val)
	if node.Left != nil {
		buff = right(node.Left, buff)
	}
	buff = append(buff, -node.Val)
	return buff
}*/
