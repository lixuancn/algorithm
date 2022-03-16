//在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
//如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
//我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
//只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//深搜
//变量不能定义在外部，因为跑多个此时用例时会污染
func isCousins(root *TreeNode, x int, y int) bool {
	var xHeight, yHeight int
	var xParent, yParent *TreeNode
	var xFind, yFind bool
	if x == 0 || y == 0 || x == y{
		return false
	}
	var dfsFunc func (cur *TreeNode, x int, y int, parentNode *TreeNode, height int)
	dfsFunc = func (cur *TreeNode, x int, y int, parentNode *TreeNode, height int){
		if cur == nil{
			return
		}
		if cur.Val == x{
			xHeight, xParent, xFind = height, parentNode, true
		}else if cur.Val == y{
			yHeight, yParent, yFind = height, parentNode, true
		}
		if xFind && yFind{
			return
		}
		h := height + 1
		dfsFunc(cur.Left, x, y, cur, h)
		if xFind && yFind{
			return
		}
		dfsFunc(cur.Right, x, y, cur, h)
	}
	dfsFunc(root, x, y, nil, 0)
	return xFind && yFind && xHeight == yHeight && xParent.Val != yParent.Val
}

//广搜
func isCousins_BFS(root *TreeNode, x int, y int) bool {
	if x == 0 || y == 0 || x == y{
		return false
	}
	type Node struct {
		Cur *TreeNode
		ParentValue int
		Height int
	}
	//广度优先
	var xNode, yNode *Node
	var queue []*Node
	queue = append(queue, &Node{
		Cur:       root,
		ParentValue:       0,
		Height:     0,
	})
	for len(queue) > 0{
		current := queue[0]
		queue = queue[1:]
		if current.Cur == nil{
			continue
		}
		if current.Cur.Val == x{
			xNode = current
		}else if current.Cur.Val == y{
			yNode = current
		}
		if xNode != nil && yNode != nil{
			if xNode.Height == yNode.Height && xNode.ParentValue != yNode.ParentValue{
				return true
			}
			return false
		}
		if current.Cur.Left != nil{
			queue = append(queue, &Node{
				Cur:       current.Cur.Left,
				ParentValue: current.Cur.Val,
				Height: current.Height + 1,
			})
		}
		if current.Cur.Right != nil{
			queue = append(queue, &Node{
				Cur:       current.Cur.Right,
				ParentValue: current.Cur.Val,
				Height: current.Height + 1,
			})
		}
	}
	return false
}
