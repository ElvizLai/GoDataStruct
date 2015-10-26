package node
import (
	"elvizlai/utils/queue"
)

type LinkNode struct {
	Item interface{}
	Next *LinkNode
}

//node for binary tree
type BinaryNode struct {
	Item        interface{}
	Left, Right *BinaryNode
}

//node for tree with child
type Node struct {
	Item  interface{}
	Child *Node //指向长子
	Peer  *Node //指向节点链表
}

//DFS父节点处理时机，又可细分前序、中序、后序三种。递归的方式，本身就是一种栈
//         1
//      2      3
//    4   5  6   7
//前序遍历为 1 2 4 5 3 6 7
//中序遍历为 4 2 5 1 6 3 7
//后序遍历为 4 5 2 6 7 3 1
func DFS(root *BinaryNode, doit func(interface{})) {
	if root!=nil {
		doit(root.Item) //前序
		DFS(root.Left, doit)
		//doit(root.Item) //中序
		DFS(root.Right, doit)
		//doit(root.Item) //后续
	}
}


//BFS要依赖队来辅助实现 - 基本思想是判断该节点是否存在子节点，如果存在，就将子节点对应的LinkNode入列
func BFS(root *Node, doit func(interface{})) {
	if root==nil {
		return
	}

	var q = queue.NewChanQueue()
	var fail = false;

	for ; !fail; {
		for kid := root.Child; kid != nil; kid = kid.Peer {
			q.Enqueue(kid) //排队
		}

		doit(root.Item)

		if v := q.Dequeue(); v==nil {
			fail = true
		}else {
			root = v.(*Node)
		}
	}
}