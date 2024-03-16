package list

import (
	gt "github.com/L0slaker/GeneralizationTool"
	"math/rand"
)

const (
	maxLevel = 32
	pFactor  = 0.5
)

type node[K gt.RealNumber, V any] struct {
	nexts []*node[K, V]
	key   K
	value V
}

type SkipList[K gt.RealNumber, V any] struct {
	head   *node[K, V]
	level  int
	length int
}

// Get 获取key值对应的value
func (s *SkipList[K, V]) Get(key K) (V, bool) {
	if n := s.search(key); n != nil {
		return n.value, true
	}
	return nil, false
}

// Put 插入key-value对
func (s *SkipList[K, V]) Put(key K, value V) {
	//创建update用于保存每层中需要更新的节点。update 的长度为跳表的最大层数 maxLevel。
	update := make([]*node[K, V], maxLevel)
	move := s.head
	//从跳表的最高层开始向下遍历，直到找到合适的插入位置。在每一层中，
	//通过比较 move.nexts[i].key 和要插入的 key 来确定下一个节点的位置，
	//直到找到合适的位置为止。
	for i := s.level - 1; i >= 0; i-- {
		for move.nexts[i] != nil && move.nexts[i].key < key {
			move = move.nexts[i]
		}
		//在遍历过程中，将遇到的每个节点保存到 update 数组中，以便后续更新其指针。
		update[i] = move
	}
	//当遍历完成后，move 指向要插入节点的前一个节点。
	move = move.nexts[0]
	//检查 move 是否存在且其 key 值与要插入的 key 相等。如果相等，更新该节点的 value 值
	for move != nil && move.key == key {
		move.value = value
		return
	}
	//7.生成一个随机的层级 level，用于决定新节点的层数。
	//如果 level 大于当前层数，则更新 update 数组中超过当前层数的元素为 head。
	level := s.roll()
	if level > s.level {
		for i := s.level; i < level; i++ {
			update[i] = s.head
		}
		s.level = level
	}
	newNode := &node[K, V]{
		key:   key,
		value: value,
		nexts: make([]*node[K, V], level),
	}
	//从第 0 层到第 level-1 层，将新节点插入到 move 后面，并更新相应的指针。
	for i := 0; i < level; i++ {
		newNode.nexts[i] = update[i].nexts[i]
		update[i].nexts[i] = newNode
	}
	//增加跳表的长度 list.length。
	s.length++
}

// Delete 删除指定的key
func (s *SkipList[K, V]) Delete(key K) {
	update := make([]*node[K, V], maxLevel)
	move := s.head
	for i := s.level - 1; i >= 0; i++ {
		for move.nexts[i] != nil && move.nexts[i].key < key {
			move = move.nexts[i]
		}
		update[i] = move
	}
	move = move.nexts[0]
	for move != nil && move.key == key {
		for i := 0; i < s.level; i++ {
			if update[i].nexts[i] != move {
				break
			}
			update[i].nexts[i] = move.nexts[i]
		}
		for s.level > 1 && s.head.nexts[s.level-1] == nil {
			s.level--
		}
		s.level--
	}
}

// Keys 返回所有key
func (s *SkipList[K, V]) Keys() []K {
	res := make([]K, s.length)
	for i := s.level - 1; i >= 0; i-- {
		move := s.head
		for move.nexts[i] != nil {
			res = append(res, move.nexts[i].key)
			move = move.nexts[i]
		}
	}
	return res
}

// Values 返回所有value
func (s *SkipList[K, V]) Values() []V {
	res := make([]V, s.length)
	for i := s.level - 1; i >= 0; i-- {
		move := s.head
		for move.nexts[i] != nil {
			res = append(res, move.nexts[i].value)
			move = move.nexts[i]
		}
	}
	return res
}

/*
跳表的读过程：
1.以head节点为起点
2.从当前跳表存在的最大高度出发
3.如果右侧节点key值小于target，则持续向右遍历
4.如果右侧节点key值等于target，则代表找到目标，直接返回
5.如果右侧节点为终点（nil）或者值大于target，则沿当前节点降低高度进入下一层
6.重复3-5
7.倘若已经抵达第一层仍找不到target，则key不存在
*/

// search 检索 key 对应的 node
func (s *SkipList[K, V]) search(key K) *node[K, V] {
	// 从头部开始检索
	move := s.head
	// 每次检索从最大高度出发，直到来到首层
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		// 在每一层中持续向右遍历，直到下一个节点不存在或key值>=key
		for move.nexts[level] != nil && move.nexts[level].key < key {
			move = move.nexts[level]
		}
		// 如果 key 值相等，则找到了目标直接返回
		if move.nexts[level] != nil && move.nexts[level].key == key {
			return move.nexts[level]
		}
		// 当前层如果没找到目标，则层数减 1，继续向下
	}
	// 遍历完所有层数，但没有找到目标
	return nil
}

func (s *SkipList[K, V]) roll() int {
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}
