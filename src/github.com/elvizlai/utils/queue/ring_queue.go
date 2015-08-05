/**
 * Created by elvizlai on 2015/8/5 09:22
 * Copyright © PubCloud
 */
package queue

type ringQueue struct {
	space    []interface{}
	mask     int
	rpt, wpt int //读写游标
}

func NewRingQueue(size int) *ringQueue {
	if size<1||size>0xffff {
		panic("Illegal ringQueue size")
	}
	var sz = 4
	for sz<=size {
		sz*=2
	}

	return &ringQueue{space:make([]interface{}, sz), mask:sz-1}
}

func (rq *ringQueue) Push(key interface{}) (fail bool) {
	var next = (rq.wpt+1)&rq.mask
	if next==rq.rpt { //因为ringQueue的头跟尾视为一个元素，要先探路判断
		return true
	}
	rq.space[rq.wpt] = key //写入数据
	rq.wpt = next//写游标滑动
	return false
}

func (rq *ringQueue)Pop() interface{} {
	if rq.rpt==rq.wpt { //如果读游标跟写游标重合，说明当前array中暂无数据
		return nil
	}

	key := rq.space[rq.rpt]//读取数据
	rq.rpt = (rq.rpt+1)&rq.mask//读游标滑动
	return key
}