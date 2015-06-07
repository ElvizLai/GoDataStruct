package main
import "fmt"
import (
	"queue"
	"stack"
	"cipher"
	"hashset"
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
}