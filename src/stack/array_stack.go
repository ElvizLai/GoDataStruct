package stack

type arraystack struct {
	item []interface{}
	size int
}

func NewArrayStack(cap int) *arraystack {
	st := make([]interface{}, 0, cap)
	return &arraystack{item:st, size:0}
}

func (this *arraystack)Push(value interface{}) {
	this.item = append(this.item, value)
	this.size++
}

func (this *arraystack)Pop() (interface{}) {
	if this.size==0 {
		return nil
	}
	this.size -= 1
	out := this.item[this.size]
	this.item = this.item[:this.size]
	return out
}

func (this *arraystack)Size() int {
	return this.size
}

func (this *arraystack)IsEmpty()bool{
	return this.size==0
}