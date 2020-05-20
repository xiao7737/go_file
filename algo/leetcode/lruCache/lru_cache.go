package lruCache

//双向链表+map   实现lru
type LruCacheNode struct {
	key   int
	value int
	prev  *LruCacheNode
	next  *LruCacheNode
}

type LruCache struct {
	md       map[int]*LruCacheNode
	len      int //当前容量
	capacity int //最大容量
	head     *LruCacheNode
	tail     *LruCacheNode
}

func Create(capacity int) LruCache {
	lc := LruCache{
		md:       make(map[int]*LruCacheNode),
		capacity: capacity,
		len:      0,
	}
	lc.head = &LruCacheNode{}
	lc.tail = &LruCacheNode{}
	lc.head.prev = nil
	lc.head.next = lc.tail
	lc.tail.prev = lc.head
	lc.tail.next = nil
	return lc
}

func (lc *LruCache) Get(key int) int {
	//键不存在
	if _, exist := lc.md[key]; !exist {
		return -1
	}
	//存在，将该节点从当前位置删除，移动到链表头部，然后返回该值
	lc.MoveToHead(lc.md[key])
	return lc.md[key].value
}

func (lc *LruCache) Put(key int, value int) {
	//判断key是否存在，存在则重新赋值，移动到最前面
	if _, exist := lc.md[key]; exist {
		lc.md[key].value = value
		lc.MoveToHead(lc.md[key])
	} else {
		//不存在
		//插入前检查容量
		if lc.len > lc.capacity {
			panic("err")
		}
		if lc.len == lc.capacity {
			lc.CutTail()
			lc.len--
		}
		//插入到头部
		node := &LruCacheNode{
			key:   key,
			value: value,
		}
		lc.AddToHead(node)
		lc.md[key] = node
		lc.len++
	}
}

func (lc *LruCache) MoveToHead(node *LruCacheNode) {
	lc.RemoveNode(node)
	lc.AddToHead(node)
}

//将该节点添加到链表最前面
func (lc *LruCache) AddToHead(node *LruCacheNode) {
	node.prev = lc.head
	node.next = lc.head.next
	lc.head.next.prev = node
	lc.head.next = node
}

//淘汰末尾节点
func (lc *LruCache) CutTail() {
	if lc.len > 0 {
		//删除map中元素
		delete(lc.md, lc.tail.prev.key)
		//去除尾巴
		lc.RemoveNode(lc.tail.prev)
	}
}

//删除该节点
func (lc *LruCache) RemoveNode(node *LruCacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}
