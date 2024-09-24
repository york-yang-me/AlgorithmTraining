package tree

import (
	"strconv"
	"strings"
)

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 递归 序列化二叉树
 * @param root TreeNode类
 * @return TreeNode类
 */
func Serialize(root *TreeNode) string {
	// write code here
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + Serialize(root.Left) + "," + Serialize(root.Right)
}
func Deserialize(s string) *TreeNode {
	// write code here
	list := strings.Split(s, ",")
	var travc func() *TreeNode
	travc = func() *TreeNode {
		val := list[0]
		list = list[1:]
		if val == "#" {
			return nil
		}
		value, _ := strconv.Atoi(val)
		return &TreeNode{Val: value, Left: travc(), Right: travc()}
	}
	return travc()
}
