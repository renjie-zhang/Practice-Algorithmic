/*
 * @Descripttion: https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
 * @version: 链表的中序遍历
 * @Author: renjie.zhang
 * @Date: 2019-10-22 15:38:57
 * @LastEditTime: 2019-10-23 19:39:32
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree1 := &TreeNode{Val: 1}
	tree2 := &TreeNode{Val: 2}
	tree3 := &TreeNode{Val: 3}
	tree1.Right = tree2
	tree2.Left = tree3
	//temp := inorderTraversal(tree1)
	temp := inorderTraversal2(tree2)
	fmt.Println(temp)
}

//解法一 使用递归
func inorderTraversal(root *TreeNode) []int {
	var temp []int
	pri(root, &temp)
	return temp
}

func pri(root *TreeNode, t *[]int) {
	if root != nil {
		if root.Left != nil {
			pri(root.Left, t)
		}
		*t = append(*t, root.Val)
		if root.Right != nil {
			pri(root.Right, t)
		}
	}
}

//解法二 莫里斯中序遍历（时间复杂度O(N),空间复杂度（O（1）））
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	//临时切片
	var temp []int
	//将当前结点指向根节点
	current := root
	//前驱结点
	var pre *TreeNode

	for current != nil {
		//当current结点没有左结点
		if current.Left == nil {
			//将结果输出
			temp = append(temp, current.Val)
			current = current.Right
			continue
		}
		//当current结点有左结点
		pre = current.Left

		for pre.Right != nil && pre.Right != current {
			pre = pre.Right
		}

		if pre.Right == nil {
			pre.Right = current
			current = current.Left
			continue
		}
		temp = append(temp, current.Val)

		pre.Right = nil

		current = current.Right
	}
	return temp
}
