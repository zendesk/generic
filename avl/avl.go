package avl

import "github.com/zyedidia/generic"

type Tree[K generic.Orderable[K], V any] struct {
	root *node[K, V]
}

func (t *Tree[K, V]) Add(key K, value V) {
	t.root = t.root.add(key, value)
}

func (t *Tree[K, V]) Remove(key K) {
	t.root.remove(key)
}

func (t *Tree[K, V]) Search(key K) (V, bool) {
	n := t.root.search(key)
	if n == nil {
		var v V
		return v, false
	}
	return n.value, true
}

func (t *Tree[K, V]) Height() int {
	return t.root.getHeight()
}

type node[K generic.Orderable[K], V any] struct {
	key   K
	value V

	height int
	left   *node[K, V]
	right  *node[K, V]
}

func (n *node[K, V]) add(key K, value V) *node[K, V] {
	if n == nil {
		return &node[K, V]{
			key: key,
			value: value,
			height: 1,
			left: nil,
			right: nil,
		}
	}

	if key.Compare(n.key) < 0 {
		n.left = n.left.add(key, value)
	} else if key.Compare(n.key) > 0 {
		n.right = n.right.add(key, value)
	} else {
		n.value = value
	}
	return n.rebalanceTree()
}

func (n *node[K, V]) remove(key K) *node[K, V] {
	if n == nil {
		return nil
	}
	if key.Compare(n.key) < 0 {
		n.left = n.left.remove(key)
	} else if key.Compare(n.key) > 0 {
		n.right = n.right.remove(key)
	} else {
		if n.left != nil && n.right != nil {
			rightMinNode := n.right.findSmallest()
			n.key = rightMinNode.key
			n.value = rightMinNode.value
			n.right = n.right.remove(rightMinNode.key)
		} else if n.left != nil {
			n = n.left
		} else if n.right != nil {
			n = n.right
		} else {
			n = nil
			return n
		}

	}
	return n.rebalanceTree()
}

func (n *node[K, V]) search(key K) *node[K, V] {
	if n == nil {
		return nil
	}
	if key.Compare(n.key) < 0 {
		return n.left.search(key)
	} else if key.Compare(n.key) > 0 {
		return n.right.search(key)
	} else {
		return n
	}
}

func (n *node[K, V]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *node[K, V]) recalculateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

func (n *node[K, V]) rebalanceTree() *node[K, V] {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor <= -2 {
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor >= 2 {
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

func (n *node[K, V]) rotateLeft() *node[K, V] {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *node[K, V]) rotateRight() *node[K, V] {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *node[K, V]) findSmallest() *node[K, V] {
	if n.left != nil {
		return n.left.findSmallest()
	} else {
		return n
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}