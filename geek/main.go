package main

var a = 12

func main() {
	t1 := new(TreeNode)
	t1.Val = 2
	t1.Left = nil
	t1.Right = nil

	t2 := new(TreeNode)
	t2.Val = 3
	t2.Left = nil
	t2.Right = nil

	t := new(TreeNode)
	t.Val = 1
	t.Left = t1
	t.Right = t2
	ints := preorderTraversal(t)
	print(len(ints))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (list []int) {
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	return list
}
