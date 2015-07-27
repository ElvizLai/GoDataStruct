package main
import "fmt"
import (
	"github.com/elvizlai/utils/queue"
	"github.com/elvizlai/utils/stack"
	"github.com/elvizlai/utils/cipher"
	"github.com/elvizlai/utils/hashset"
	"github.com/elvizlai/utils/node"
	"github.com/elvizlai/utils/orderedMap"
	"reflect"
)

func main() {
	aqu := queue.NewArrayQueue(0)
	aqu.Enqueue("1")
	aqu.Enqueue("2")
	aqu.Enqueue("3")

	nqu := queue.NewNodeQueue()
	nqu.Enqueue("1")
	nqu.Enqueue("2")
	nqu.Enqueue("3")

	cqu := queue.NewChanQueue()
	cqu.Enqueue("1")
	cqu.Enqueue("2")
	cqu.Enqueue("3")

	ast := stack.NewArrayStack(0)
	ast.Push("1")
	ast.Push("2")
	ast.Push("3")


	nst := stack.NewNodeStack()
	nst.Push("1")
	nst.Push("2")
	nst.Push("3")

	fmt.Println("array queue", aqu.Dequeue(), aqu.Dequeue(), aqu.Dequeue(), aqu.Dequeue())
	fmt.Println("node queue", nqu.Dequeue(), nqu.Dequeue(), nqu.Dequeue(), nqu.Dequeue())
	fmt.Println("chan queue", cqu.Dequeue(), cqu.Dequeue(), cqu.Dequeue(), cqu.Dequeue())

	fmt.Println("array stack", ast.Pop(), ast.Pop(), ast.Pop(), ast.Pop())
	fmt.Println("node stack", nst.Pop(), nst.Pop(), nst.Pop(), nst.Pop())


	cc := cipher.GenEncryptionFunc(func(str string) []byte {
		return []byte(str)
	})

	fmt.Println("cipherFunc", cc("elvizlai"))

	hashSet1 := hashset.NewHashSet()
	hashSet1.Add("elvizlai")
	hashSet1.Add("phone")
	hashSet2 := hashset.NewHashSet()
	hashSet2.Add("elvizlai")

	fmt.Println(hashSet1.Elements())
	fmt.Println(hashSet2.Elements())
	fmt.Println("IsSame:", hashSet1.Same(hashSet2))
	fmt.Println("IsSuperSet:", hashSet1.IsSuperSet(hashSet2))

	fmt.Println("FS")

	n0 := &node.BinaryNode{}
	n0.Item = "1"

	n01 := &node.BinaryNode{}
	n01.Item = "2"

	n02 := &node.BinaryNode{}
	n02.Item = "3"

	n11 := &node.BinaryNode{}
	n11.Item = "4"

	n12 := &node.BinaryNode{}
	n12.Item = "5"

	n21 := &node.BinaryNode{}
	n21.Item = "6"

	n22 := &node.BinaryNode{}
	n22.Item = "7"

	n0.Left = n01
	n0.Right = n02

	n01.Left = n11
	n01.Right = n12

	n02.Left = n21
	n02.Right = n22


	node.DFS(n0, func(item interface{}) {
		fmt.Println(item)
	})

	fmt.Println("BFS")

	m1 := &node.Node{Item:1}

	m2 := &node.Node{Item:2}

	m3 := &node.Node{Item:3}

	m4 := &node.Node{Item:4}

	m5 := &node.Node{Item:5}

	m6 := &node.Node{Item:6}

	m7 := &node.Node{Item:7}

	m1.Child = m2
	m2.Peer = m3
	m2.Child = m4
	m4.Peer = m5
	m3.Child = m6
	m6.Peer = m7


	node.BFS(m1, func(item interface{}) {
		fmt.Println(item)
	})


	fmt.Println("------OrderedMap--------")
	keys := orderedMap.NewKeys(func(e1 interface{}, e2 interface{}) int8 {
		k1, k2 := e1.(string), e2.(string)
		if k1==k2 {
			return 0
		}else if k1>k2 {
			return 1
		}else {
			return -1
		}
	}, reflect.TypeOf(""))


	omap := orderedMap.NewOrderedMap(keys, reflect.TypeOf(""))
	omap.Put("lai", "abc")
	omap.Put("a", "123")
	omap.Put("汉", "elv")
	omap.Put("3", "iz")
	fmt.Println(omap)

}