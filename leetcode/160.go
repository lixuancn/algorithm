type ListNode struct {
    Val int
    Next *ListNode
}

//双指针，可以从hash表O(m+n)，空间O(n)的基础上优化为时间O(n)，空间O(1)
//首先确保A和B都不为空，创建两个指针，分别指向A和B的头，同时向后移动，pointA为空时跳到B的头，pointB为空时跳到A的头。
//当pointA和pointB指向同一个节点时则相交，同时指向null时则不相交。
//证明如下：
//假设AB相交，不相交长度分别是a和b，相交长度为c。那么a+c=m(A的长度), b+c=n(B的长度)
//  如果a=b，则pointA和pointB同时到达相交点。
//  如果a!=b，pointA走完a+c+b，pointB走完b+c+a，则pointA和pointB同时到达相交点。
//假设AB不相交
//  如果m=n，pointA和pointB同时到达各自链表的末尾，同时值为null。
//  如果m!=n，则pointA走过m+n，pointB走过n+m，pointA和pointB同时到达各自链表的末尾，同时值为null。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil{
        return nil
    }
    pointA, pointB := headA, headB
    for pointA != pointB{
        if pointA == nil{
            pointA = headB
        }else{
            pointA = pointA.Next
        }
        if pointB == nil{
            pointB = headA
        }else{
            pointB = pointB.Next
        }
    }
    return pointA
}

//hash表的方法
func getIntersectionNode_hash(headA, headB *ListNode) *ListNode {
    mapA := map[*ListNode]interface{}{}
    for ; headA != nil; headA = headA.Next{
        mapA[headA] = nil
    }
    for ; headB != nil; headB = headB.Next{
        if _, ok := mapA[headB]; ok{
            return headB
        }
    }
    return nil
}

