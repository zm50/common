package cache

type Sizable interface {
	Size() int
}

type LRUNode[K comparable, V Sizable] struct {
	key K
	val V
	prev *LRUNode[K, V]
	next *LRUNode[K, V]
}

type LRUCache[K comparable, V Sizable] struct {
	capacity int
	cache map[K]*LRUNode[K, V]
	head *LRUNode[K, V]
	tail *LRUNode[K, V]
}

// 初始化 LRU Cache
func NewLRUCache[K comparable, V Sizable](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		cache: make(map[K]*LRUNode[K, V]),
		head: nil,
		tail: nil,
	}
}

func (l *LRUCache[K, V]) Get(key K) *V {
	if node, ok := l.cache[key]; ok {
		l.moveToHead(node)
		return &(node.val)
	}

	return nil
}

func (l *LRUCache[K, V]) Put(key K, val V) {
    if node, ok := l.cache[key]; ok {
        node.val = val
        l.moveToHead(node) // 将节点移动到链表头部
    } else {
		newNode := &LRUNode[K, V]{
            key:   key,
            val:   val,
        }

		l.addHead(newNode)

		for len(l.cache) > l.capacity {
            tail := l.removeTail() // 删除尾部节点
            delete(l.cache, tail.key) // 在哈希表中移除对应的缓存数据
        }
    }
}

func (l *LRUCache[K, V]) moveToHead(node *LRUNode[K, V]) {
	// 三种场景，node在头部，node在中间，node在尾部
	if node == l.head {
		// node在头部，不需要移动
		return
	}

	if node == l.tail {
		// node在尾部，直接移动到头部
		l.tail = node.prev
		l.tail.next = nil
		l.head.prev = node
		node.prev = nil
		node.next = l.head
		l.head = node
		return
	}

	// node在中间，先将node从原来的位置删除，再插入到头部
	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = nil
	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *LRUCache[K, V]) addHead(node *LRUNode[K, V]) {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.cache[node.key] = node
		node.next = l.head
		l.head.prev = node
		l.head = node
	}

	l.capacity += node.val.Size()
}

func (l *LRUCache[K, V]) removeTail() *LRUNode[K, V] {
	tail := l.tail
	l.tail = tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}

	l.capacity -= tail.val.Size()

	return tail
}
