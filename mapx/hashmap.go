package mapx

import "generalization_tool/syncx"

type Hashable interface {
	// Code 返回元素的哈希值
	Code() uint64
	// Equals 比较两元素是否相等
	Equals(key any) bool
}

type node[T Hashable, ValType any] struct {
	key   T
	value ValType
	next  *node[T, ValType]
}

type HashMap[T Hashable, ValType any] struct {
	hashmap  map[uint64]*node[T, ValType]
	nodePool *syncx.Pool[*node[T, ValType]]
}

func NewHashMap[T Hashable, ValType any](size int) *HashMap[T, ValType] {
	return &HashMap[T, ValType]{
		hashmap: make(map[uint64]*node[T, ValType], size),
		nodePool: syncx.NewPool[*node[T, ValType]](func() *node[T, ValType] {
			return &node[T, ValType]{}
		}),
	}
}

func (m *HashMap[T, ValType]) newNode(key T, value ValType) *node[T, ValType] {
	newNode := m.nodePool.Get()
	newNode.value = value
	newNode.key = key
	return newNode
}

func (m *HashMap[T, ValType]) Put(key T, value ValType) error {
	hash := key.Code()
	root, ok := m.hashmap[hash]
	// 池中不存在重复数据，可以新建
	if !ok {
		hash = key.Code()
		newNode := m.newNode(key, value)
		m.hashmap[hash] = newNode
		return nil
	}
	pre := root
	for root != nil {
		// 遍历整个链表，查找是否存在相同的键
		if root.key.Equals(key) {
			//更新value
			root.value = value
			return nil
		}
		pre = root
		root = root.next
	}
	newNode := m.newNode(key, value)
	pre.next = newNode
	return nil
}

func (m *HashMap[T, ValType]) Get(key T) (ValType, bool) {
	hash := key.Code()
	root, ok := m.hashmap[hash]
	var val ValType
	if !ok {
		return val, false
	}
	for root != nil {
		if root.key.Equals(key) {
			return root.value, true
		}
		root = root.next
	}
	return val, false
}

func (m *HashMap[T, ValType]) Delete(key T) (ValType, bool) {
	root, ok := m.hashmap[key.Code()]
	var zero ValType
	if !ok {
		return zero, false
	}
	pre := root
	num := 0
	for root != nil {
		if root.key.Equals(key) {
			// 找到具有相同键的节点
			if num == 0 && root.next == nil {
				// 如果链表只有一个节点，则从哈希表中删除该键
				delete(m.hashmap, key.Code())
			} else if num == 0 && root.next != nil {
				// 如果链表有多个节点且第一个节点就是要删除的节点，则更新哈希表中的根节点
				m.hashmap[key.Code()] = root.next
			} else {
				// 如果节点不是第一个节点，则将前一个节点的next指针指向要删除节点的下一个节点
				pre.next = root.next
			}
			// 记录要删除节点的值，并将节点重置为默认值
			value := root.value
			root.formatting()
			// 将删除的节点放回节点池中以供复用
			m.nodePool.Put(root)
			return value, true
		}
		num++
		pre = root
		root = root.next
	}
	return zero, false
}

// formatting 将节点的键、值和next指针重置
func (n *node[T, ValType]) formatting() {
	var zeroVal ValType
	var zeroT T
	n.key = zeroT
	n.value = zeroVal
	n.next = nil
}

func (m *HashMap[T, ValType]) Keys() []T {
	res := make([]T, 0)
	for _, n := range m.hashmap {
		curNode := n
		for curNode != nil {
			res = append(res, curNode.key)
			curNode = curNode.next
		}
	}
	return res
}

func (m *HashMap[T, ValType]) Values() []ValType {
	res := make([]ValType, 0)
	for _, n := range m.hashmap {
		curNode := n
		for curNode != nil {
			res = append(res, curNode.value)
			curNode = curNode.next
		}
	}
	return res
}
