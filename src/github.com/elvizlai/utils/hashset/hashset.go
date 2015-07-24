package hashset
import (
	"bytes"
	"fmt"
)

type set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Elements() []interface{}
	String() string
}

type hashSet struct {
	m map[interface{}]bool
}

func NewHashSet() *hashSet {
	return &hashSet{m:make(map[interface{}]bool)}
}

//添加元素
func (set *hashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

//删除元素
func (set *hashSet) Remove(e interface{}) {
	delete(set.m, e)
}

//清除所有元素
func (set *hashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

//判断与其他HashSet类型值是否相等
func (set *hashSet) Contains(e interface{}) bool {
	return set.m[e]
}

func (set *hashSet) Len() int {
	return len(set.m)
}

func (set *hashSet) Same(other set) bool {
	if other==nil {
		return false
	}

	if set.Len()!=other.Len() {
		return false
	}

	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

//生产快照
func (set *hashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range set.m {
		if actualLen<initialLen {
			snapshot[actualLen]=key
		}else {
			snapshot=append(snapshot, key)
		}
		actualLen++
	}
	if actualLen<initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

//获取自身的字符串表示形式
func (set *hashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		}else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

//a是否为b超集即b中的所有元素在a中都能找到
func (set *hashSet) IsSuperSet(other set) bool {
	if other==nil {
		return false
	}
	oneLen := set.Len()
	otherLen := other.Len()
	if oneLen==0||oneLen==otherLen {
		return false
	}
	if oneLen>0&&otherLen==0 {
		return true
	}
	for _, v := range other.Elements() {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}