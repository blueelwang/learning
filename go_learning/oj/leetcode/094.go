package leetcode





type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}

	stack := []*TreeNode{}
	for ; root != nil; root = root.Left {
		stack = append(stack, root)
	}

	for len(stack) > 0 {
		top := stack[len(stack) - 1]
		stack = stack[0:len(stack) - 1]
		result = append(result, top.Val)
		
		for root = top.Right; root != nil; root = root.Left {
			stack = append(stack, root)
		}
	}
	return result
}

func inorderTraversalRecursively(root *TreeNode, result []int) []int {
	if nil == root {
		return result
	}

	result = inorderTraversalRecursively(root.Left, result)
	result = append(result, root.Val)
	result = inorderTraversalRecursively(root.Right, result)
	return result
}