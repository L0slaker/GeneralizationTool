package tree

import (
	"errors"
	"generalization_tool"
)

type color bool

const (
	Red   color = false
	Black color = true
)

var (
	ErrRBTreeSameRBNode     = errors.New("RBTree 不能添加重复节点Key")
	ErrRBTreeNotExistRBNode = errors.New("RBTree 不存在节点Key")
)

type rbNode[K any, V any] struct {
	key    K
	value  V
	color  color
	left   *rbNode[K, V]
	right  *rbNode[K, V]
	parent *rbNode[K, V]
}

func newRBNode[K any, V any](key K, value V) *rbNode[K, V] {
	return &rbNode[K, V]{
		key:    key,
		value:  value,
		color:  Red,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

func (n *rbNode[K, V]) setNode(v V) {
	if n == nil {
		return
	}
	n.value = v
}

type RBTree[K any, V any] struct {
	root    *rbNode[K, V]
	compare generalization_tool.Comparator[K]
	size    int
}

func (rb *RBTree[K, V]) Size() int {
	if rb == nil {
		return 0
	}
	return rb.size
}

func NewRBTree[K any, V any](compare generalization_tool.Comparator[K]) *RBTree[K, V] {
	return &RBTree[K, V]{
		root:    nil,
		compare: compare,
	}
}

// Add 添加节点
func (rb *RBTree[K, V]) Add(key K, value V) error {
	return rb.addNode(newRBNode(key, value))
}

// Delete 删除节点
func (rb *RBTree[K, V]) Delete(key K) (V, bool) {
	if node := rb.findNode(key); node != nil {
		value := node.value
		rb.deleteNode(node)
		return value, true
	}
	var zero V
	return zero, false
}

// Find 查找节点
func (rb *RBTree[K, V]) Find(key K) (V, error) {
	var value V
	if n := rb.findNode(key); n != nil {
		return n.value, nil
	}
	return value, ErrRBTreeNotExistRBNode
}

// Set 设置节点
func (rb *RBTree[K, V]) Set(key K, value V) error {
	if node := rb.findNode(key); node != nil {
		node.setNode(value)
		return nil
	}
	return ErrRBTreeNotExistRBNode
}

func (rb *RBTree[K, V]) KeyValues() ([]K, []V) {
	keys := make([]K, 0, rb.size)
	values := make([]V, 0, rb.size)
	if rb.root == nil {
		return keys, values
	}
	rb.inOrderTraversal(func(n *rbNode[K, V]) {
		keys = append(keys, n.key)
		values = append(values, n.value)
	})
	return keys, values
}

func (rb *RBTree[K, V]) addNode(n *rbNode[K, V]) error {
	var fixNode *rbNode[K, V]
	if rb.root == nil {
		rb.root = newRBNode[K, V](n.key, n.value)
		fixNode = rb.root
	} else {
		temp := rb.root
		cmp := 0
		parent := &rbNode[K, V]{}
		for temp != nil {
			parent = temp
			cmp = rb.compare(n.key, temp.key)
			if cmp < 0 {
				temp = temp.left
			} else if cmp > 0 {
				temp = temp.right
			} else if cmp == 0 {
				return ErrRBTreeSameRBNode
			}
		}
		fixNode = &rbNode[K, V]{
			key:    n.key,
			value:  n.value,
			color:  Red,
			parent: parent,
		}
		if cmp < 0 {
			parent.left = fixNode
		} else {
			parent.right = fixNode
		}
	}
	rb.size++
	rb.fixAfterAdd(fixNode)
	return nil
}

// fixAfterAdd 插入时着色旋转
// fixUncleRed uncle节点为红右
// fixAddLeftBlack uncle节点为黑右
// fixAddRightBlack uncle节点为黑左
func (rb *RBTree[K, V]) fixAfterAdd(n *rbNode[K, V]) {
	n.color = Red
	for n != nil && n != rb.root && n.getParent().getColor() == Red {
		uncle := n.getUncle()
		// 判断uncle节点的颜色
		if uncle.getColor() == Red {
			n = rb.fixUncleRed(n, uncle)
			continue
		}
		if n.getParent() == n.getGrand().getLeft() {
			n = rb.fixAddLeftBlack(n)
			continue
		}
		n = rb.fixAddRightBlack(n)
	}
	rb.root.setColor(Black)
}

// fixUncleRed
// 上溢情况，不需要旋转
// parent、uncle 染成黑色，grand上溢
func (rb *RBTree[K, V]) fixUncleRed(n *rbNode[K, V], u *rbNode[K, V]) *rbNode[K, V] {
	n.getParent().setColor(Black)
	u.setColor(Black)
	n.getGrand().setColor(Red)
	return n.getGrand()
}

// fixAddLeftBlack
// 非上溢情况 && LL或LR情况，LL情况不执行左旋，LR情况双旋
// 将parent染成黑色，grand染成红色
// 如果x为左节点则跳过左旋操作
func (rb *RBTree[K, V]) fixAddLeftBlack(n *rbNode[K, V]) *rbNode[K, V] {
	if n == n.getParent().getRight() {
		n = n.getParent()
		rb.leftRotate(n)
		//rb.leftRotate(n.getParent())
	}
	n.getParent().setColor(Black)
	n.getGrand().setColor(Red)
	rb.rightRotate(n.getGrand())
	return n
}

// fixAddRightBlack
// 非上溢情况 && RR或RL情况，RR情况不执行右旋，RL情况双旋
// 将parent染成黑色，grand染成红色
func (rb *RBTree[K, V]) fixAddRightBlack(n *rbNode[K, V]) *rbNode[K, V] {
	if n == n.getParent().getLeft() {
		n = n.getParent()
		rb.rightRotate(n)
		//rb.rightRotate(n.getParent())
	}
	n.getParent().setColor(Black)
	n.getGrand().setColor(Red)
	rb.leftRotate(n.getGrand())
	return n
}

// deleteNode 红黑树的删除方法
// 1.取出后继节点
// case1：node左右为非空子节点
// case2：node左右只有一个非空子节点；
// case3：node左右均为空节点
// 2.着色旋转
// case1：当删除节点为红色是，直接删除
// case2：当删除节点非空且为黑色时，需要维持平衡树的性质
func (rb *RBTree[K, V]) deleteNode(target *rbNode[K, V]) {
	n := target
	// case1：node左右为非空子节点，取后继节点
	if n.left != nil && n.right != nil {
		s := rb.findSuccessor(n)
		n.key, n.value = s.key, s.value
		n = s
	}
	var replacedNode *rbNode[K, V]
	// case2：node只有一个非空子节点
	if n.left != nil {
		replacedNode = n.left
	} else {
		replacedNode = n.right
	}
	if replacedNode != nil {
		replacedNode.parent = n.parent
		if n.parent == nil {
			rb.root = replacedNode
		} else if n == n.parent.left {
			n.parent.left = replacedNode
		} else {
			n.parent.right = replacedNode
		}
		n.left, n.right, n.parent = nil, nil, nil
		// 替换节点后补齐红黑树的性质
		if n.getColor() {
			rb.fixAfterDelete(replacedNode)
		}
	} else if n.parent == nil {
		// 如果node节点无父节点，说明node为root节点
		rb.root = nil
	} else {
		// case3：node没有子节点
		if n.getColor() {
			rb.fixAfterDelete(n)
		}
		if n.parent != nil {
			if n == n.parent.left {
				n.parent.left = nil
			} else if n == n.parent.right {
				n.parent.right = nil
			}
			n.parent = nil
		}
	}
	rb.size--
}

// findSuccessor 寻找后继节点。后继节点是大于要删除节点的最小节点
// case1: node节点存在右子节点,则右子树的最小节点是node的后继节点
// case2: node节点不存在右子节点,则其第一个为左节点的祖先的父节点为node的后继节点
func (rb *RBTree[K, V]) findSuccessor(n *rbNode[K, V]) *rbNode[K, V] {
	if n == nil {
		return nil
	} else if n.right != nil {
		//case1: node节点存在右子节点,则右子树的最小节点是node的后继节点
		p := n.right
		for p.left != nil {
			p = p.left
		}
		return p
	} else {
		// case2: node节点不存在右子节点,后继节点可能在它的祖先节点中。
		//通过向上遍历节点的父节点，并判断节点 node 是其父节点的左子节点还是右子节点，
		//直到找到第一个满足 node 是其父节点的左子节点的节点
		p := n.parent
		ch := n
		for p != nil && ch == p.right {
			ch = p
			p = p.parent
		}
		return p
	}
}

// fixAfterDelete 删除时着色旋转
// 根据x是节点位置分为fixAfterDeleteLeft,fixAfterDeleteRight两种情况
func (rb *RBTree[K, V]) fixAfterDelete(n *rbNode[K, V]) {
	for n != rb.root && n.getColor() == Black {
		if n == n.parent.getLeft() {
			n = rb.fixAfterDeleteLeft(n)
		} else {
			n = rb.fixAfterDeleteRight(n)
		}
	}
	n.setColor(Black)
}

// fixAfterDeleteLeft 处理n为左子节点时的平衡处理
func (rb *RBTree[K, V]) fixAfterDeleteLeft(n *rbNode[K, V]) *rbNode[K, V] {
	bro := n.getParent().getRight()
	// case1：兄弟节点为红色：兄弟节点染黑，父节点染红，对父节点右旋。然后就成了case2的情况
	if bro.getColor() == Red {
		bro.setColor(Black)
		bro.getParent().setColor(Red)
		rb.leftRotate(n.getParent())
		bro = n.getParent().getRight()
	}
	// case2：兄弟节点为黑
	if bro.getLeft().getColor() == Black && bro.getRight().getColor() == Black {
		// 2.1兄弟节点的子节点都为黑
		// 父节点向下与兄弟节点合并；将兄弟染成红色，父节点染黑即可；
		// 如果父节点为黑，直接将父节点当作被删除的节点处理
		bro.setColor(Red)
		n = n.getParent()
	} else {
		// 2.2兄弟节点的子节点至少有一个为红
		// 如果兄弟节点 bro 的右子节点为黑色，左子节点为红色，
		// 那么将兄弟节点 bro 的颜色设为红色，左子节点设为黑色，
		// 并进行右旋转操作，使得兄弟节点的右子节点变为新的兄弟节点。
		if bro.getRight().getColor() == Black {
			bro.getLeft().setColor(Black)
			bro.setColor(Red)
			rb.rightRotate(bro)
			bro = n.getParent().getRight()
		}
		// 交换兄弟节点 bro 和父节点的颜色，并将父节点设为黑色，
		// 兄弟节点的右子节点设为黑色，并进行左旋转操作，将父节点作为新的节点 n
		bro.setColor(n.getParent().getColor())
		n.getParent().setColor(Black)
		bro.getRight().setColor(Black)
		rb.leftRotate(n.getParent())
		n = rb.root
	}
	return n
}

// fixAfterDeleteRight 处理n为右子节点时的平衡处理
func (rb *RBTree[K, V]) fixAfterDeleteRight(n *rbNode[K, V]) *rbNode[K, V] {
	bro := n.getParent().getLeft()
	// case1：兄弟节点为红
	if bro.getColor() == Red {
		bro.setColor(Black)
		n.getParent().setColor(Red)
		rb.rightRotate(n.getParent())
		bro = n.getBrother()
	}
	// case2：兄弟节点为黑
	if bro.getLeft().getColor() == Black && bro.getRight().getColor() == Black {
		// 2.1兄弟节点没有红色子节点
		bro.setColor(Red)
		n = n.getParent()
	} else {
		// 2.2兄弟节点至少有1个红色子节点
		if bro.getLeft().getColor() == Black {
			bro.getRight().setColor(Black)
			bro.setColor(Red)
			rb.leftRotate(bro)
			bro = n.getParent().getLeft()
		}
		// 交换兄弟节点 bro 和父节点的颜色，并将父节点设为黑色，
		// 兄弟节点的右子节点设为黑色，并进行左旋转操作，将父节点作为新的节点 n
		bro.setColor(n.getParent().getColor())
		n.getParent().setColor(Black)
		bro.getLeft().setColor(Black)
		rb.rightRotate(n.getParent())
		n = rb.root
	}
	return n
}

func (rb *RBTree[K, V]) findNode(key K) *rbNode[K, V] {
	n := rb.root
	for n != nil {
		cmp := rb.compare(key, n.key)
		if cmp < 0 {
			n = n.left
		} else if cmp > 0 {
			n = n.right
		} else {
			return n
		}
	}
	return nil
}

func (rb *RBTree[K, V]) inOrderTraversal(visit func(node *rbNode[K, V])) {
	//1.创建一个栈 stack 来模拟递归调用栈，并初始化为空。
	stackArea := make([]*rbNode[K, V], 0, rb.size)
	//2.初始化当前节点为红黑树的根节点 rb.root。
	cur := rb.root
	//3.进入一个循环，循环条件为当前节点不为 nil，或者栈不为空（还有节点需要处理）。
	for cur != nil || len(stackArea) > 0 {
		for cur != nil {
			stackArea = append(stackArea, cur)
			cur = cur.left
		}
		cur = stackArea[len(stackArea)-1]
		stackArea = stackArea[:len(stackArea)-1]
		visit(cur)
		cur = cur.right
	}
}

func (rb *RBTree[K, V]) leftRotate(n *rbNode[K, V]) {
	if n == nil || n.getRight() == nil {
		return
	}
	r := n.right

	n.right = r.left
	if r.left != nil {
		r.left.parent = n
	}

	r.parent = n.parent
	if n.parent == nil {
		rb.root = r
	} else if n.parent.left == n {
		n.parent.left = r
	} else {
		n.parent.right = r
	}

	r.left = n
	n.parent = r
}

func (rb *RBTree[K, V]) rightRotate(n *rbNode[K, V]) {
	if n == nil || n.getLeft() == nil {
		return
	}
	l := n.left
	n.left = l.right
	if l.right != nil {
		l.right.parent = n
	}

	l.parent = n.parent
	if n.parent == nil {
		rb.root = l
	} else if n.parent.right == n {
		n.parent.right = l
	} else {
		n.parent.left = l
	}

	l.right = n
	n.parent = l
}

func (n *rbNode[K, V]) getColor() color {
	if n == nil {
		return Black
	}
	return n.color
}

func (n *rbNode[K, V]) setColor(color color) {
	if n == nil {
		return
	}
	n.color = color
}

func (n *rbNode[K, V]) getLeft() *rbNode[K, V] {
	if n == nil {
		return nil
	}
	return n.left
}

func (n *rbNode[K, V]) getRight() *rbNode[K, V] {
	if n == nil {
		return nil
	}
	return n.right
}

func (n *rbNode[K, V]) getParent() *rbNode[K, V] {
	if n == nil {
		return nil
	}
	return n.parent
}

func (n *rbNode[K, V]) getUncle() *rbNode[K, V] {
	if n == nil {
		return nil
	}
	return n.getParent().getBrother()
}

func (n *rbNode[K, V]) getBrother() *rbNode[K, V] {
	if n == nil {
		return nil
	}
	if n == n.getParent().getLeft() {
		return n.getParent().getRight()
	}
	return n.getParent().getLeft()
}

func (n *rbNode[K, V]) getGrand() *rbNode[K, V] {
	if n == nil {
		return nil
	}
	return n.getParent().getParent()
}
