package orderedmap

type OrderedMap struct {
	keys []interface{}
	m    map[interface{}]interface{}
}

func (omap *OrderedMap) Len() int {
	return len(omap.keys)
}