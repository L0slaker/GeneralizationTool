package tree

import (
	"errors"
	"generalization_tool"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRBTree(t *testing.T) {
	testCases := []struct {
		name    string
		compare generalization_tool.Comparator[int]
		wantV   bool
	}{
		{
			name:    "int",
			compare: compare(),
			wantV:   true,
		},
		{
			name:    "nil",
			compare: nil,
			wantV:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, string](compare())
			assert.Equal(t, tc.wantV, IsRedBlackTree[int, string](redBlackTree.root))
		})
	}
}

func TestRBTree_Add(t *testing.T) {
	RBTreeTest := []struct {
		name string
		want bool
		n    *rbNode[int, string]
	}{
		{
			name: "nil",
			want: true,
			n:    nil,
		},
		{
			name: "node-nil",
			want: true,
			n:    nil,
		},
		{
			name: "root",
			want: true,
			n: &rbNode[int, string]{
				left:  nil,
				right: nil,
				color: Black,
			},
		},
		{
			name: "root-Red",
			want: false,
			n: &rbNode[int, string]{
				left:  nil,
				right: nil,
				color: Red,
			},
		},
		//			 root(黑)
		//			/
		//		   a(黑)
		{
			name: "root with one child",
			want: true,
			n: &rbNode[int, string]{
				left: &rbNode[int, string]{
					left:  nil,
					right: nil,
					color: Red,
				},
				right: nil,
				color: Black,
			},
		},
		//			 root(黑)
		//			/	    \
		//		   a(红)    b(黑)
		{
			name: "root with two child",
			want: false,
			n: &rbNode[int, string]{
				left: &rbNode[int, string]{
					left:  nil,
					right: nil,
					color: Red,
				},
				right: &rbNode[int, string]{
					left:  nil,
					right: nil,
					color: Black,
				},
				color: Black,
			},
		},
		//			 root(黑)
		//			/	     \
		//		   a(黑)      b(黑)
		//		 /  \        /    \
		//      nil  c(红)  d(红)   nil
		//           / \     / \
		//          nil nil nil nil
		{
			name: "not same black node",
			want: true,
			n: &rbNode[int, string]{
				left: &rbNode[int, string]{
					left: nil,
					right: &rbNode[int, string]{
						left:  nil,
						right: nil,
						color: Red,
					},
					color: Black,
				},
				right: &rbNode[int, string]{
					left: &rbNode[int, string]{
						left:  nil,
						right: nil,
						color: Red,
					},
					right: nil,
					color: Black,
				},
				color: Black,
			},
		},
		//		   	     7(黑)
		//			  /	        \
		//		   5(黑)	     10(红)
		//		 /   \            /    \
		//      4(红) 6(红)      9(黑)   12(黑)
		//     / \    / \       /  \   /   \
		//    /   \  /   \     /    \  nil  nil
		//   nil nil nil nil  8(红) 11(红)
		//                   / \    / \
		//                 nil nil nil nil
		{
			name: "root grandson",
			want: true,
			n: &rbNode[int, string]{
				parent: nil,
				key:    7,
				left: &rbNode[int, string]{
					key: 5,
					left: &rbNode[int, string]{
						key:   4,
						color: Red,
					},
					right: &rbNode[int, string]{
						key:   6,
						color: Red,
					},
					color: Black,
				},
				right: &rbNode[int, string]{
					key:   10,
					color: Red,
					left: &rbNode[int, string]{
						key:   9,
						color: Black,
						left: &rbNode[int, string]{
							key:   8,
							color: Red,
						},
					},
					right: &rbNode[int, string]{
						key:   12,
						color: Black,
						left: &rbNode[int, string]{
							key:   11,
							color: Red,
						},
					},
				},
				color: Black,
			},
		},
	}

	for _, rt := range RBTreeTest {
		t.Run(rt.name, func(t *testing.T) {
			res := IsRedBlackTree[int](rt.n)
			assert.Equal(t, rt.want, res)
		})
	}

	testCases := []struct {
		name     string
		key      []int
		want     bool
		wantErr  error
		wantSize int
		wantKey  int
	}{
		{
			name:     "nil",
			key:      nil,
			want:     true,
			wantSize: 0,
		},
		{
			name:     "one element",
			key:      []int{1},
			want:     true,
			wantSize: 1,
		},
		{
			name:     "two element",
			key:      []int{1, 2},
			want:     true,
			wantSize: 2,
			wantKey:  1,
		},
		//		   	   3(黑)
		//		    /	      \
		//		   2(红)	  4(红)
		//		 /   \       /    \
		//     1(黑)  nil   nil   nil
		//     / \
		//   nil nil
		{
			name:     "normal",
			key:      []int{1, 2, 3, 4},
			want:     true,
			wantSize: 4,
			wantKey:  3,
		},
		{
			name:     "same key",
			key:      []int{0, 0, 1, 2, 2, 3},
			want:     true,
			wantSize: 0,
			wantErr:  errors.New("RBTree 不能添加重复节点Key"),
		},
		{
			name:     "disorder",
			key:      []int{1, 2, 0, 3, 5, 4},
			want:     true,
			wantErr:  nil,
			wantSize: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.key); i++ {
				err := rbTree.Add(tc.key[i], i)
				if err != nil {
					assert.Equal(t, tc.wantErr, err)
					return
				}
			}
			res := IsRedBlackTree[int, int](rbTree.root)
			assert.Equal(t, tc.want, res)
			assert.Equal(t, tc.wantSize, rbTree.Size())
		})
	}
}

func TestRBTree_Delete(t *testing.T) {
	testCases := []struct {
		name      string
		deleteKey int
		keys      []int
		want      bool
		size      int
	}{
		{
			name:      "nil",
			deleteKey: 0,
			keys:      nil,
			want:      true,
			size:      0,
		},
		{
			name:      "node-empty",
			deleteKey: 0,
			keys:      []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:      true,
			size:      9,
		},
		{
			name:      "左右非空子节点,删除节点为黑色",
			deleteKey: 11,
			keys:      []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:      true,
			size:      8,
		},
		{
			name:      "左右只有一个非空子节点,删除节点为黑色",
			deleteKey: 11,
			keys:      []int{4, 5, 6, 7, 8, 9, 11, 12},
			want:      true,
			size:      7,
		},
		{
			name:      "左右均为空节点,删除节点为黑色",
			deleteKey: 12,
			keys:      []int{4, 5, 6, 7, 8, 9, 12},
			want:      true,
			size:      6,
		},
		{
			name:      "左右非空子节点,删除节点为红色",
			deleteKey: 5,
			keys:      []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:      true,
			size:      8,
		},
		//此状态无法构造出正确的红黑树
		//{
		//	name: "左右只有一个非空子节点,删除节点为红色",
		//	deleteKey: 5,
		//	keys: []int{4,5,6,7,8,9,11,12},
		//	want: true,
		//},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := rbTree.Add(tc.keys[i], i)
				if err != nil {
					panic(err)
				}
			}

			assert.Equal(t, tc.want, IsRedBlackTree[int](rbTree.root))
			rbTree.Delete(tc.deleteKey)
			assert.Equal(t, tc.want, IsRedBlackTree[int](rbTree.root))
			assert.Equal(t, tc.size, rbTree.Size())
		})
	}
}

func TestRBTree_Find(t *testing.T) {
	testCases := []struct {
		name      string
		target    int
		keys      []int
		wantKey   int
		wantError error
	}{
		{
			name:      "nil",
			target:    0,
			keys:      nil,
			wantError: errors.New("未找到该节点"),
		},
		{
			name:      "node-empty",
			target:    0,
			keys:      []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantError: errors.New("未找到该节点"),
		},
		{
			name:    "find 11",
			target:  11,
			keys:    []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: 11,
		},
		{
			name:    "find 12",
			target:  12,
			keys:    []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: 12,
		},
		{
			name:    "find 7",
			target:  7,
			keys:    []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: 7,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := rbTree.Add(tc.keys[i], tc.keys[i])
				if err != nil {
					panic(err)
				}
			}
			assert.Equal(t, true, IsRedBlackTree[int](rbTree.root))
			found, err := rbTree.Find(tc.target)
			if err != nil {
				assert.Equal(t, tc.wantError, errors.New("未找到该节点"))
			} else {
				assert.Equal(t, tc.target, found)
			}
		})
	}
}

func TestRBTree_KeyValues(t *testing.T) {
	testCases := []struct {
		name       string
		keys       []int
		values     []int
		wantKeys   []int
		wantValues []int
	}{
		{
			name:       "nil",
			keys:       nil,
			values:     nil,
			wantKeys:   []int{},
			wantValues: []int{},
		},
		{
			name:       "normal",
			keys:       []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			values:     []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKeys:   []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantValues: []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := rbTree.Add(tc.keys[i], tc.values[i])
				if err != nil {
					panic(err)
				}
			}
			keys, values := rbTree.KeyValues()
			assert.Equal(t, tc.wantKeys, keys)
			assert.Equal(t, tc.wantValues, values)
		})
	}
}

func TestRBTree_addNode(t *testing.T) {
	testCases := []struct {
		name    string
		key     []int
		want    bool
		wantErr error
	}{
		{
			name: "nil",
			key:  nil,
			want: true,
		},
		{
			name: "normal",
			key:  []int{1, 2, 3, 4},
			want: true,
		},
		{
			name:    "same",
			key:     []int{1, 1, 2, 3, 4},
			want:    true,
			wantErr: errors.New("RBTree 不能添加重复节点Key"),
		},
		{
			name: "disorder",
			key:  []int{1, 2, 0, 3, 5, 4},
			want: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, string](compare())
			for i := 0; i < len(tc.key); i++ {
				err := rbTree.addNode(&rbNode[int, string]{
					key: tc.key[i],
				})
				if err != nil {
					assert.Equal(t, tc.wantErr, err)
				}
			}
			res := IsRedBlackTree[int](rbTree.root)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestRBTree_deleteNode(t *testing.T) {
	testCase := []struct {
		name      string
		deleteKey int
		keys      []int
		want      bool
		wantErr   error
	}{
		{
			name:      "nil",
			deleteKey: 0,
			keys:      nil,
			wantErr:   errors.New("未找到节点"),
		},
		{
			name:      "node-empty",
			deleteKey: 0,
			keys:      []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantErr:   errors.New("未找到节点"),
		},
		{
			name:      "root",
			deleteKey: 2,
			keys:      []int{2, 1},
			want:      true,
		},
		{
			name:      "delete root with one elements in tree",
			deleteKey: 7,
			keys:      []int{7},
			want:      true,
		},
		{
			name:      "delete root with multiple elements in tree",
			deleteKey: 7,
			keys:      []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:      true,
		},
		{
			name:      "删除节点为黑色：自身为右节点，左右为非空子节点",
			deleteKey: 11,
			keys:      []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:      true,
		},
		{
			name:      "删除节点为黑色：自身为右节点，左右只有一个非空子节点",
			deleteKey: 11,
			keys:      []int{2, 3, 4, 5, 6, 7, 8, 9, 11, 12},
			want:      true,
		},
		{
			name:      "删除节点为黑色：自身为右节点，左右均为空节点",
			deleteKey: 6,
			keys:      []int{2, 3, 4, 5, 6, 7, 8, 9, 11, 12},
			want:      true,
		},
		{
			name:      "删除节点为黑色：自身为左节点，左右为非空子节点",
			deleteKey: 3,
			keys:      []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:      true,
		},
		{
			name:      "删除节点为黑色：自身为左节点，左右只有一个非空子节点",
			deleteKey: 3,
			keys:      []int{2, 3, 4, 5, 6, 7, 8, 9, 11, 12},
			want:      true,
		},
		{
			name:      "删除节点为黑色：自身为左节点，左右均为空节点",
			deleteKey: 8,
			keys:      []int{2, 3, 5, 6, 7, 8, 9, 11, 12},
			want:      true,
		},
		{
			name:      "删除节点为黑色：自身为左节点，只有左子节点",
			deleteKey: 3,
			keys:      []int{5, 3, 4, 6, 2},
			want:      true,
		},
		{
			name:      "删除节点为红色：自身为左节点，左右均为空节点,",
			deleteKey: 11,
			keys:      []int{2, 3, 5, 6, 7, 8, 9, 11, 12},
			want:      true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := rbTree.Add(tc.keys[i], i)
				if err != nil {
					panic(err)
				}
			}
			deleteNode := rbTree.findNode(tc.deleteKey)
			if deleteNode == nil {
				assert.Equal(t, tc.wantErr, errors.New("未找到节点"))
			} else {
				rbTree.deleteNode(deleteNode)
				assert.Equal(t, tc.want, IsRedBlackTree[int](rbTree.root))
			}
		})
	}
}

func TestRBTree_findNode(t *testing.T) {
	testCases := []struct {
		name    string
		target  int
		keys    []int
		wantKey int
		wantErr error
	}{
		{
			name:    "nil",
			target:  0,
			keys:    nil,
			wantErr: errors.New("未找到该节点"),
		},
		{
			name:    "node-empty",
			target:  0,
			keys:    []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantErr: errors.New("未找到该节点"),
		},
		{
			name:    "find 11",
			target:  11,
			keys:    []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: 11,
		}, {
			name:    "find 12",
			target:  12,
			keys:    []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: 12,
		}, {
			name:    "find 7",
			target:  7,
			keys:    []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantKey: 7,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := rbTree.Add(tc.keys[i], i)
				if err != nil {
					panic(err)
				}
			}
			assert.Equal(t, true, IsRedBlackTree[int](rbTree.root))
			found := rbTree.findNode(tc.target)
			if found == nil {
				assert.Equal(t, tc.wantErr, errors.New("未找到该节点"))
			} else {
				assert.Equal(t, tc.wantKey, found.key)
			}
		})
	}
}

func TestRBTree_rotateLeft(t *testing.T) {
	testCases := []struct {
		name           string
		keys           []int
		wantParent     int
		wantLeftChild  int
		wantRightChild int
		rotateNode     int
	}{
		{
			name:       "only-root",
			keys:       []int{1},
			wantParent: 1,
			rotateNode: 1,
		},
		{
			name:           "节点有2各子节点，自身是右节点",
			keys:           []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			rotateNode:     9,
			wantParent:     11,
			wantLeftChild:  8,
			wantRightChild: 10,
		},
		{
			name:          "节点右2个子节点，自身是左节点",
			keys:          []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			rotateNode:    5,
			wantParent:    6,
			wantLeftChild: 4,
		},
		{
			name:          "节点有1个子节点",
			keys:          []int{1, 2, 3, 4},
			rotateNode:    2,
			wantParent:    3,
			wantLeftChild: 1,
		},
		{
			name:          "节点没有子节点",
			keys:          []int{1, 2, 3},
			rotateNode:    2,
			wantParent:    3,
			wantLeftChild: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := rbTree.Add(tc.keys[i], i)
				if err != nil {
					panic(err)
				}
			}
			rotateNode := rbTree.findNode(tc.rotateNode)
			rbTree.leftRotate(rotateNode)
			if rotateNode.getParent() != nil {
				assert.Equal(t, tc.wantParent, rotateNode.getParent().key)
				if rotateNode.getLeft() != nil {
					assert.Equal(t, tc.wantLeftChild, rotateNode.getLeft().key)
				}
				if rotateNode.getRight() != nil {
					assert.Equal(t, tc.wantRightChild, rotateNode.getRight().key)
				}
			}
		})
	}
}

func TestRBTree_rotateRight(t *testing.T) {
	testCases := []struct {
		name           string
		keys           []int
		wantParent     int
		wantLeftChild  int
		wantRightChild int
		rotateNode     int
	}{
		{
			name:       "only-root",
			keys:       []int{1},
			wantParent: 1,
			rotateNode: 1,
		},
		{
			name:           "节点有2各子节点，自身是右节点",
			keys:           []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			rotateNode:     9,
			wantParent:     8,
			wantRightChild: 11,
		},
		{
			name:           "节点右2个子节点，自身是左节点",
			keys:           []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
			rotateNode:     5,
			wantParent:     4,
			wantRightChild: 6,
		},
		{
			name:           "节点有1个子节点",
			keys:           []int{4, 5, 3, 2},
			rotateNode:     4,
			wantParent:     3,
			wantRightChild: 5,
		},
		{
			name:           "节点没有子节点",
			keys:           []int{4, 5, 3},
			rotateNode:     4,
			wantParent:     3,
			wantRightChild: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := rbTree.Add(tc.keys[i], i)
				if err != nil {
					panic(err)
				}
			}
			rotateNode := rbTree.findNode(tc.rotateNode)
			rbTree.rightRotate(rotateNode)
			if rotateNode.getParent() != nil {
				assert.Equal(t, tc.wantParent, rotateNode.getParent().key)
				if rotateNode.getLeft() != nil {
					assert.Equal(t, tc.wantLeftChild, rotateNode.getLeft().key)
				}
				if rotateNode.getRight() != nil {
					assert.Equal(t, tc.wantRightChild, rotateNode.getRight().key)
				}
			}
		})
	}
}

func TestRBNode_getColor(t *testing.T) {
	testCases := []struct {
		name      string
		node      *rbNode[int, int]
		wantColor color
	}{
		{
			name:      "nod-nil",
			node:      nil,
			wantColor: Black,
		},
		{
			name:      "new node",
			node:      newRBNode[int, int](1, 1),
			wantColor: Red,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantColor, tc.node.getColor())
		})
	}
}

func TestRBNode_getLeft(t *testing.T) {
	testCases := []struct {
		name     string
		node     *rbNode[int, int]
		wantNode *rbNode[int, int]
	}{
		{
			name:     "nod-nil",
			node:     nil,
			wantNode: nil,
		},
		{
			name:     "new node",
			node:     newRBNode[int](1, 1),
			wantNode: nil,
		},
		{
			name: "new node have left-child",
			node: &rbNode[int, int]{
				key: 2,
				left: &rbNode[int, int]{
					key: 1,
				},
			},
			wantNode: &rbNode[int, int]{
				key: 1,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantNode, tc.node.getLeft())
		})
	}
}

func TestRBNode_getRight(t *testing.T) {
	testCases := []struct {
		name     string
		node     *rbNode[int, int]
		wantNode *rbNode[int, int]
	}{
		{
			name:     "nod-nil",
			node:     nil,
			wantNode: nil,
		},
		{
			name:     "new node",
			node:     newRBNode[int](1, 1),
			wantNode: nil,
		},
		{
			name: "new node have left-child",
			node: &rbNode[int, int]{
				key: 1,
				right: &rbNode[int, int]{
					key: 2,
				},
			},
			wantNode: &rbNode[int, int]{
				key: 2,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantNode, tc.node.getRight())
		})
	}
}

func TestRBNode_getParent(t *testing.T) {
	testCases := []struct {
		name     string
		node     *rbNode[int, int]
		wantNode *rbNode[int, int]
	}{
		{
			name:     "nod-nil",
			node:     nil,
			wantNode: nil,
		},
		{
			name:     "new node",
			node:     newRBNode[int](1, 1),
			wantNode: nil,
		},
		{
			name: "new node have left-child",
			node: &rbNode[int, int]{
				key: 1,
				parent: &rbNode[int, int]{
					key: 2,
				},
			},
			wantNode: &rbNode[int, int]{
				key: 2,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantNode, tc.node.getParent())
		})
	}
}

func TestRBNode_setColor(t *testing.T) {
	testCases := []struct {
		name      string
		node      *rbNode[int, int]
		color     color
		wantColor color
	}{
		{
			name:      "nod-nil",
			node:      nil,
			color:     Red,
			wantColor: Black,
		},
		{
			name:      "new node",
			node:      newRBNode[int, int](1, 1),
			color:     Black,
			wantColor: Black,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.node.setColor(tc.color)
			assert.Equal(t, tc.wantColor, tc.node.getColor())
		})
	}
}

func TestNewRBNode(t *testing.T) {
	testCases := []struct {
		name     string
		key      int
		value    int
		wantNode *rbNode[int, int]
	}{
		{
			name:  "new node",
			key:   1,
			value: 1,
			wantNode: &rbNode[int, int]{
				key:    1,
				value:  1,
				left:   nil,
				right:  nil,
				parent: nil,
				color:  Red,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			node := newRBNode[int, int](tc.key, tc.value)
			assert.Equal(t, tc.wantNode, node)
		})
	}
}

func TestRBNode_getBrother(t *testing.T) {
	testCases := []struct {
		name     string
		keys     []int
		nodeKey  int
		wantNode int
	}{
		{
			name: "nil",
			keys: nil,
		},
		{
			name:    "no-brother",
			nodeKey: 1,
			keys:    []int{1},
		},
		{
			name:    "no-brother",
			nodeKey: 1,
			keys:    []int{1, 2},
		},
		{
			name:     "have brother",
			keys:     []int{1, 2, 3, 4},
			nodeKey:  1,
			wantNode: 3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := redBlackTree.Add(tc.keys[i], i)
				if err != nil {
					panic(err)
				}
			}
			target := redBlackTree.findNode(tc.nodeKey)
			bro := target.getBrother()
			if bro == nil {
				return
			}
			assert.Equal(t, tc.wantNode, bro.key)
		})
	}
}

func TestRBNode_getGrand(t *testing.T) {
	testCases := []struct {
		name     string
		keys     []int
		nodeKey  int
		wantNode int
	}{
		{
			name: "nil",
			keys: nil,
		},
		{
			name:    "no-grandpa",
			nodeKey: 1,
			keys:    []int{1},
		},
		{
			name:    "no-grandpa",
			nodeKey: 1,
			keys:    []int{1, 2},
		},
		{
			name:     "normal",
			keys:     []int{1, 2, 3, 4},
			nodeKey:  4,
			wantNode: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := redBlackTree.Add(tc.keys[i], i)
				if err != nil {
					panic(err)
				}
			}
			tagNode := redBlackTree.findNode(tc.nodeKey)
			brNode := tagNode.getGrand()
			if brNode == nil {
				return
			}
			assert.Equal(t, tc.wantNode, brNode.key)
		})
	}
}

func TestRBNode_getUncle(t *testing.T) {
	testCases := []struct {
		name     string
		keys     []int
		nodeKey  int
		wantNode int
	}{
		{
			name: "nil",
			keys: nil,
		},
		{
			name:    "no-uncle",
			nodeKey: 1,
			keys:    []int{1},
		},
		{
			name:    "no-uncle",
			nodeKey: 1,
			keys:    []int{1, 2},
		},
		{
			name:     "normal",
			keys:     []int{1, 2, 3, 4},
			nodeKey:  4,
			wantNode: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := redBlackTree.Add(tc.keys[i], i)
				if err != nil {
					panic(err)
				}
			}
			tagNode := redBlackTree.findNode(tc.nodeKey)
			brNode := tagNode.getUncle()
			if brNode == nil {
				return
			}
			assert.Equal(t, tc.wantNode, brNode.key)
		})
	}
}

func TestRBNode_set(t *testing.T) {
	testCases := []struct {
		name     string
		node     *rbNode[int, int]
		value    int
		wantNode *rbNode[int, int]
	}{
		{
			name:     "nil",
			node:     nil,
			value:    1,
			wantNode: nil,
		},
		{
			name: "new node",
			node: &rbNode[int, int]{
				key:    1,
				value:  0,
				left:   nil,
				right:  nil,
				parent: nil,
				color:  Red,
			},
			value: 1,
			wantNode: &rbNode[int, int]{
				key:    1,
				value:  1,
				left:   nil,
				right:  nil,
				parent: nil,
				color:  Red,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.node.setNode(tc.value)
			assert.Equal(t, tc.wantNode, tc.node)
		})
	}
}

func TestRBTree_findSuccessor(t *testing.T) {
	testCases := []struct {
		name      string
		keys      []int
		successor int
		wantKey   int
	}{
		{
			name:      "nil successor",
			keys:      nil,
			successor: 8,
		},
		{
			name:      "no successor",
			keys:      []int{2},
			successor: 2,
		},
		{
			name:      "no-right successor",
			keys:      []int{5, 4, 6, 3, 2},
			successor: 4,
			wantKey:   5,
		},
		{
			name:      "right successor",
			keys:      []int{5, 4, 6, 3, 2},
			successor: 3,
			wantKey:   4,
		},
		{
			name:      "right successor",
			keys:      []int{5, 4, 7, 6, 3, 2},
			successor: 5,
			wantKey:   6,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := redBlackTree.Add(tc.keys[i], i)
				if err != nil {
					return
				}
			}
			target := redBlackTree.findNode(tc.successor)
			successor := redBlackTree.findSuccessor(target)
			if successor == nil {
				return
			}
			assert.Equal(t, tc.wantKey, successor.key)
		})
	}
}

func TestRBTree_fixAddLeftBlack(t *testing.T) {
	testCases := []struct {
		name     string
		keys     []int
		addNode  int
		wantNode int
	}{
		{
			name:     "node is right",
			keys:     []int{2, 1, 3},
			addNode:  3,
			wantNode: 2,
		},
		{
			name:     "node is left",
			keys:     []int{2, 1},
			addNode:  1,
			wantNode: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := redBlackTree.Add(tc.keys[i], i)
				if err != nil {
					return
				}
			}
			node := redBlackTree.findNode(tc.addNode)
			x := redBlackTree.fixAddLeftBlack(node)
			assert.Equal(t, tc.wantNode, x.key)
		})
	}
}

func TestRBTree_fixAddRightBlack(t *testing.T) {
	testCases := []struct {
		name    string
		keys    []int
		addNode int
		want    int
	}{
		{
			name:    "node is left",
			keys:    []int{2, 1},
			addNode: 1,
			want:    2,
		},
		{
			name:    "node is right",
			keys:    []int{2, 1, 3},
			addNode: 3,
			want:    3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redBlackTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := redBlackTree.Add(tc.keys[i], i)
				if err != nil {
					return
				}
			}
			node := redBlackTree.findNode(tc.addNode)
			x := redBlackTree.fixAddRightBlack(node)
			assert.Equal(t, tc.want, x.key)
		})

	}
}

func TestRBTree_fixAfterDeleteLeft(t *testing.T) {
	testCases := []struct {
		name      string
		deleteKey int
		keys      []int
		wantNode  int
		wantErr   error
	}{
		{
			name:      "兄弟节点是红色;兄弟节点左子节点左侧是黑色",
			deleteKey: 10,
			keys:      []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantNode:  11,
		},
		{
			name:      "兄弟节点是红色;兄弟节点左侧是黑色",
			deleteKey: 1,
			keys:      []int{2, 1, 3},
			wantNode:  2,
		},
		{
			name:      "兄弟节点是黑色;兄弟节点左侧是黑色",
			deleteKey: 2,
			keys:      []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantNode:  3,
		},
		{
			name:      "兄弟节点是黑色;兄弟节点左侧不是黑色",
			deleteKey: 8,
			keys:      []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantNode:  5,
		},
		{
			name:      "节点左旋之后兄弟节点是红色",
			deleteKey: 21,
			keys:      []int{15, 20, 10, 16, 21, 8, 14, 7},
			wantNode:  15,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := rbTree.Add(tc.keys[i], i)
				if err != nil {
					panic(err)
				}
			}
			deleteNode := rbTree.findNode(tc.deleteKey)
			if deleteNode == nil {
				assert.Equal(t, tc.wantErr, errors.New("未找到该节点"))
			} else {
				fixed := rbTree.fixAfterDeleteLeft(deleteNode)
				assert.Equal(t, tc.wantNode, fixed.key)
			}
		})
	}
}

func TestRBTree_fixAfterDeleteRight(t *testing.T) {
	testCases := []struct {
		name      string
		deleteKey int
		keys      []int
		wantNode  int
		wantErr   error
	}{
		{
			name:      "兄弟节点是红色;兄弟节点左子节点左侧是黑色",
			deleteKey: 12,
			keys:      []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantNode:  11,
		},
		{
			name:      "兄弟节点是红色;兄弟节点左侧是黑色",
			deleteKey: 3,
			keys:      []int{2, 1, 3},
			wantNode:  2,
		},
		{
			name:      "兄弟节点是黑色;兄弟节点左侧是黑色",
			deleteKey: 11,
			keys:      []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			wantNode:  9,
		},
		{
			name:      "兄弟节点是黑色;兄弟节点左侧不是黑色",
			deleteKey: 4,
			keys:      []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 1, 0},
			wantNode:  5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rbTree := NewRBTree[int, int](compare())
			for i := 0; i < len(tc.keys); i++ {
				err := rbTree.Add(tc.keys[i], i)
				if err != nil {
					panic(err)
				}
			}
			deleteNode := rbTree.findNode(tc.deleteKey)
			if deleteNode == nil {
				assert.Equal(t, tc.wantErr, errors.New("未找到该节点"))
			} else {
				fixed := rbTree.fixAfterDeleteRight(deleteNode)
				assert.Equal(t, tc.wantNode, fixed.key)
			}
		})
	}
}

func compare() generalization_tool.Comparator[int] {
	return generalization_tool.ComparatorRealNumber[int]
}

// IsRedBlackTree 检测是否满足红黑树
func IsRedBlackTree[K any, V any](root *rbNode[K, V]) bool {
	//检测节点是否为黑色
	if !root.getColor() {
		return false
	}
	//count 取最左树的黑色节点作为对照
	count := 0
	num := 0
	node := root
	for node != nil {
		if node.getColor() {
			count++
		}
		node = node.getLeft()
	}
	return nodeCheck[K](root, count, num)
}

// nodeCheck 节点检测
// 1.是否有连续的红色节点
// 2.每条路径的黑色节点是否一致
func nodeCheck[K any, V any](n *rbNode[K, V], count int, num int) bool {
	if n == nil {
		return true
	}
	if !n.getColor() && !n.parent.getColor() {
		return false
	}
	if n.getColor() {
		num++
	}
	if n.getLeft() == nil && n.getRight() == nil {
		if num != count {
			return false
		}
	}
	return nodeCheck(n.left, count, num) && nodeCheck(n.right, count, num)
}
