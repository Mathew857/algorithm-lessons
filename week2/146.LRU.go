/*
请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

 

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4



LRU 的实现关键思想：定义 LRUCache 结构

	type LRUCache struct {
		// 大小
		size 		int
		// 容量
		capacity 	int
		// cache 容器
		cache 		map[int]*DLinkedNode
		// 头部 cache 尾部 cache 指针
		head, tail  *DLinkedNode
	}

其中 cache 实例为一个双链表，size 表示大小，capacity 表示给定容量，head，tail 表示前后驱指针，LRUCache 结构定义之后，LRUCache 的实现即是对 LRUCache 实现增删改查操作：
- 增加节点
	判断节点是否存在：
		如果不存在，则做链表头部节点增加操作，并判断增加后的size 是否大于 capacity，如果大于则移出尾部节点并 size -1
		如果存在，则更新节点 value
- 删除节点
	删除节点本质为 链表节点删除操作，其包含：
		删除中间节点：
		删除尾部节点：
- 修改节点
	判断节点是否存在：
		如果存在，则更新节点 value
		如果不存在，做链表头结点增加操作
- 查询节点
	判断节点是否存在：
		如果存在返回节点 value
		如果不存在返回 -1
同时，因为 LRUCache 本身的特性，即最近最少使用策略，其核心思想是将频繁访问（查询、增加）节点至于链表头部，当 cache size 大于容量 capacity 时候，移出尾部节点

*/
type LRUCache struct {
	// 大小
	size 		int
	// 容量
	capacity 	int
	// cache 容器
	cache 		map[int]*DLinkedNode
	// 头部 cache 尾部 cache 指针
	head, tail  *DLinkedNode
}

// 定义双链表结构
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

// 初始化双链表
func initDLinedNode(key, value int) *DLinkedNode  {
	return &DLinkedNode{
		key: key,
		value, value,
	}
}

// 构造器，构造 LRUCache 实例
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		// cache 实体为一个 hash 表
		cache: map[int]*DLinkedNode{},
		head: initDLinedNode(0, 0)
		tail: initDLinedNode(0, 0)
		capacity: capacity,
	}
	// 将 l 的尾结点赋值给 l 的头结点的下一个节点
	l.head.next = l.tail
	// 将 l 的头结点赋值给 l 的尾结点的前一个节点
	l.tail.prev = l.head
	return l
}

// 添加节点
func (this *LRUCache) Put(key int, value int) {
	// 如果 cache 中不存在这个节点
	if _, ok := this.cache[key]; !ok {
		// 则初始化一个双链表节点
		node := initDLinedNode(key, value)
		// 并将这个节点 node 添加到 cache[key] 中
		this.cache[key] = node 
		// 并将这个节点添加到头部
		this.addToHead(node)
		// cache 长度加一
		this.size++
		// 如果 size + 1 后的值大于 capacity
		if this.size > this.capacity {
			// 移出尾部节点
			removed := this.removeTail()
			// 删除 cache 中key
			delete(this.cache, removed.key)
			// size 减一
			this.size-- 
		}
	}else {
		// 如果 cache 中存在这个节点则将这个节点赋值给 node
		node := this.cache[key]
		// 更新 value
		node.value = value
		// 将这个节点移动到头部节点
		this.moveToHead(node)
	}
}

// 查找节点
func (this *LRUCache) Get(key int) {
	if _, ok := this.cache[key]; !ok {
		return -1
	}	
	// 将 this.cache[key] 的值赋值给 node
	node := this.cache[key]
	// 将 key 这个节点移动到头结点
	this.moveToHead(node)
	// 返回 node 节点的 value
	return node.value
}

// 将节点添加到头部
func (this *LRUCache) addToHead(node *DLinkedNode) {
	// node 前驱指向头部
	node.prev = this.head
	// node 后驱指向头部下一个节点
	node.next = this.head.next
	// 头部节点的下一个节点的前驱结点为 node
	this.head.next.prev = node 
	// 头部节点的下一个节点为 node
	this.head.next = node 
}

// 删除节点
func (this *LRUCache) removeNode(node *DLinkedNode) {
	// 将 node 前驱节点的下一个节点指向 node 的下一个节点
	node.prev.next = node.next
	// node 节点的下一个节点的前驱节点指向 node 的前驱节点
	node.next.prev = node.prev
}

// 移动到头部
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 删除尾部节点 
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev 
	this.removeNode(node)
	return node 
}