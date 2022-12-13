
# 力扣 66 加一

## 题目描述

给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

### 示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。

### 示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。

### 示例 3：

输入：digits = [0]
输出：[1]
 

### 提示：

1 <= digits.length <= 100
0 <= digits[i] <= 9

## 题解

### 思路
当我们对数组 dights 加一时，我们只需要关注 dights 的末尾出现了多少个9即可：
- 如果 dights 的末尾没有 9 ，例如 [1,2,3]，那么我们直接将末尾的数加一，得到[1,2,4] 并返回即可
- 如果 dights 的末尾有若干个 9，例如 [1,2,3,9,9]，那么我们只需要找出末尾开始的第一个不为 9 的元素，即 3，将该元素加一，得到 [1,2,4,9,9]。随后将末尾的 9 全部置0即返回即可，即 [1,2,4,0,0]
- 如果 dights 的所有元素都是 9 ，例如 [9,9,9,9,9]，那么答案为 [1,0,0,0,0,0]。我们只需要构造一个比 dights 多1 的新数组，将首元素置位1，其余元素置为0即可。

```golang
func plusOne(digits []int) []int {
    // 获取数组长度
    n := len(digits)
    for i := n-1; i >= 0; i-- {
        // 从最后一个元素开始判断是否等于 9
        if digits[i] != 9 {
            // 不等于9的话末尾元素直接加一
            digits[i]++
            for j := i+1; j < n ; j++ {
                digits[j] = 0 
            }
            return digits
        }
    }

    digits = make([]int, n+1)
    digits[0] = 1
    return digits
}
```
### 复杂度分析
- 时间复杂度：O(n)，其中 n 是数组 digits 的长度
- 空间复杂度：O(1)，返回值不计入空间复杂度

# 力扣 21 合并两个有序链表
## 题目描述

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

### 示例 1：

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

### 示例 2：

输入：l1 = [], l2 = []
输出：[]

### 示例 3：

输入：l1 = [], l2 = [0]
输出：[0]
 

### 提示：

两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列

## 题解

### 思路

我们可以按照如下递归地定义两个链表里面的 merge 操作（忽略边界清开，比如空链表等）：

list1[0] + merge(list1[1:], list2)   list1[0] > list2[0]
list2[0] + merge(list1, list2[1:])   otherwise

也就是说，两个链表头部值较小的一个节点与剩下元素的 merge 操作结果合并。

### 算法
我们直接将以上递归过程建模，同时需要考虑边界情况。
如果 l1 或者 l2 一开始就是空链表，那么没有任何操作需要合并，所以我们只需要返回非空链表。否则，我们要判断 l1 和 l2 哪一个链表的头结点的值更小，然后递归地决定下一个添加到结果里的节点。如果两个链表有一个为空，递归结束。

```golang

/*
* Definition for singly-linked list.
* type ListNode struct {
*      Val int
*      Next *ListNode   
* }
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // 如果 list1 等于 list2 ，优先返回 list2
    if list1 == nil {
        return list2
    }

    // 如果 list2 等于空，则返回 list1
    if list2 == nil {
        return list1
    }

    if list1 == list2 {
        return list2
    }

    // 判断 list1 和 list2 哪一个链表的头结点的值更小，然后递归决定下一个添加到结构里面的节点。
    if list1.Val <= list2.Val {
        list1.Next = mergeTwoLists(list1.Next, list2)
        return list1
    } else {
        list2.Next = mergeTwoLists(list1, list2.Next)
        return list2
    }
}

/*
//(1,1):代表第一次进入递归函数，并且从第一个口进入，并且记录进入前链表的状态
merge(1,1): 1->4->5->null, 1->2->3->6->null
    merge(2,2): 4->5->null, 1->2->3->6->null
    	merge(3,2): 4->5->null, 2->3->6->null
    		merge(4,2): 4->5->null, 3->6->null
    			merge(5,1): 4->5->null, 6->null
    				merge(6,1): 5->null, 6->null
    					merge(7): null, 6->null
    					return l2
    				l1.next --- 5->6->null, return l1
    			l1.next --- 4->5->6->null, return l1
    		l2.next --- 3->4->5->6->null, return l2
    	l2.next --- 2->3->4->5->6->null, return l2
    l2.next --- 1->2->3->4->5->6->null, return l2
l1.next --- 1->1->2->3->4->5->6->null, return l1
*/
```

### 复杂度分析

时间复杂度：O(n + m)O(n+m)，其中 nn 和 mm 分别为两个链表的长度。因为每次调用递归都会去掉 l1 或者 l2 的头节点（直到至少有一个链表为空），函数 mergeTwoList 至多只会递归调用每个节点一次。因此，时间复杂度取决于合并后的链表长度，即 O(n+m)O(n+m)。

空间复杂度：O(n + m)O(n+m)，其中 nn 和 mm 分别为两个链表的长度。递归调用 mergeTwoLists 函数时需要消耗栈空间，栈空间的大小取决于递归调用的深度。结束递归调用时 mergeTwoLists 函数最多调用 n+mn+m 次，因此空间复杂度为 O(n+m)O(n+m)。

# 力扣 641. 设计循环双端队列

## 题目描述
设计实现双端队列。

实现 MyCircularDeque 类:

- MyCircularDeque(int k) ：构造函数,双端队列最大为 k 。
- boolean insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true ，否则返回 false 。
- boolean insertLast() ：将一个元素添加到双端队列尾部。如果操作成功返回 true ，否则返回 false 。
- boolean deleteFront() ：从双端队列头部删除一个元素。 如果操作成功返回 true ，否则返回 false 。
- boolean deleteLast() ：从双端队列尾部删除一个元素。如果操作成功返回 true ，否则返回 false 。
- int getFront() )：从双端队列头部获得一个元素。如果双端队列为空，返回 -1 。
- int getRear() ：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1 。
- boolean isEmpty() ：若双端队列为空，则返回 true ，否则返回 false  。
- boolean isFull() ：若双端队列满了，则返回 true ，否则返回 false 。
 

### 示例 1：

输入
["MyCircularDeque", "insertLast", "insertLast", "insertFront", "insertFront", "getRear", "isFull", "deleteLast", "insertFront", "getFront"]
[[3], [1], [2], [3], [4], [], [], [], [4], []]
输出
[null, true, true, true, false, 2, true, true, true, 4]

解释
MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
circularDeque.insertLast(1);			        // 返回 true
circularDeque.insertLast(2);			        // 返回 true
circularDeque.insertFront(3);			        // 返回 true
circularDeque.insertFront(4);			        // 已经满了，返回 false
circularDeque.getRear();  				// 返回 2
circularDeque.isFull();				        // 返回 true
circularDeque.deleteLast();			        // 返回 true
circularDeque.insertFront(4);			        // 返回 true
circularDeque.getFront();				// 返回 4
 
提示：

1 <= k <= 1000
0 <= value <= 1000
insertFront, insertLast, deleteFront, deleteLast, getFront, getRear, isEmpty, isFull  调用次数不大于 2000 次


## 题解

### 思路
使用双向链表模拟双端队列，实现双端队列对首与队尾元素的添加、删除。双向链表实现比较简单，双向链表支持 O(1) 时间复杂度内在指定节点的前后插入新的节点或者删除新的节点。

循环双端队列的属性如下：
- head：队列的头结点
- tail：队列的尾节点
- capacity: 队列的容量大小
- size：队列当前的元素容量

循环双端队列的接口方法如下：
- MyCircularDeque(k int): 初始化队列，同时初始化队列元素数量 size 为0。head,tail 初始化为空。
- insertFront(value int): 队列未满时，在队首头节点 head 之前插入一个新的节点，并更新 haed，并更新 size。
- insertLast(value int): 队列未满时，在队尾节点 tail 之后插入一个新的节点，并更新 tail，并更新 size。
- deleteFront(): 队列不为空时，删除头结点 head，并更新 head 为 head 的后一个节点，并更新 size。
- deleteLast(): 队列不为空时，删除尾节点 tail，并更新 tail 为 tail 的前一个节点，并鞥更新 size。
- getFront(): 返回队首节点指向的值，需要检测队列是否为空。
- getRear(): 返回队尾节点指向的值，需要检测队列是否为空。
- isEmpty(): 检测当前 size 是否为空。
- isFull(): 检测当前 size 是否为 capacity。

```golang
type node struct{
    prev, next   *node
    val          int
}

type MyCircularDeque struct {
    head, tail      *node
    capacity, size  int
}


func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{capacity: k}
}


func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    node := &node{val: value}
    if this.IsEmpty() {
        this.head = node
        this.tail = node
    }else {
        node.next = this.head
        this.head.prev = node
        this.head = node 
    }
    this.size++
    return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    node := &node{val: value}
    if this.IsEmpty() {
        this.head = node
        this.tail = node
    } else {
        this.tail.next = node
        node.prev = this.tail
        this.tail = node
    }
    this.size++
    return true
}


func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    this.head = this.head.next
    if this.head != nil {
        this.head.prev = nil
    }
    this.size--
    return true
}


func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    this.tail = this.tail.prev
    if this.tail != nil {
        this.tail.next = nil 
    }
    this.size--
    return true
}


func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.head.val
}


func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.tail.val
}


func (this *MyCircularDeque) IsEmpty() bool {
    return this.size == 0
}


func (this *MyCircularDeque) IsFull() bool {
    return this.size == this.capacity
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
```

### 复杂度分析

时间复杂度：初始化和每项操作的时间复杂度均为 O(1)O(1)。

空间复杂度：O(k)O(k)，其中 kk 为给定的队列元素数目。



# 力扣 85. 最大矩形

## 题目描述
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

 

示例 1：

### 题解

### 思路

为了计算矩形的最大面积，我们只需要计算每个柱状图中的最大面积，并找到全局最大值。

输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。
示例 2：

输入：matrix = []
输出：0
示例 3：

输入：matrix = [["0"]]
输出：0
示例 4：

输入：matrix = [["1"]]
输出：1
示例 5：

输入：matrix = [["0","0"]]
输出：0


```golang
func maximalRectangle(matrix [][]byte) (ans int) {
    if len(matrix) == 0 {
        return
    }
    m, n := len(matrix), len(matrix[0])
    left := make([][]int, m)
    for i, row := range matrix {
        left[i] = make([]int, n)
        for j, v := range row {
            if v == '0' {
                continue
            }
            if j == 0 {
                left[i][j] = 1
            } else {
                left[i][j] = left[i][j-1] + 1
            }
        }
    }
    for j := 0; j < n; j++ { // 对于每一列，使用基于柱状图的方法
        up := make([]int, m)
        down := make([]int, m)
        stk := []int{}
        for i, l := range left {
            for len(stk) > 0 && left[stk[len(stk)-1]][j] >= l[j] {
                stk = stk[:len(stk)-1]
            }
            up[i] = -1
            if len(stk) > 0 {
                up[i] = stk[len(stk)-1]
            }
            stk = append(stk, i)
        }
        stk = nil
        for i := m - 1; i >= 0; i-- {
            for len(stk) > 0 && left[stk[len(stk)-1]][j] >= left[i][j] {
                stk = stk[:len(stk)-1]
            }
            down[i] = m
            if len(stk) > 0 {
                down[i] = stk[len(stk)-1]
            }
            stk = append(stk, i)
        }
        for i, l := range left {
            height := down[i] - up[i] - 1
            area := height * l[j]
            ans = max(ans, area)
        }
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### 复杂度分析

时间复杂度：O(mn)O(mn)，其中 mm 和 nn 分别是矩阵的行数和列数。计算 left 矩阵需要 O(mn)O(mn) 的时间；对每一列应用柱状图算法需要 O(m)O(m) 的时间，一共需要 O(mn)O(mn) 的时间。

空间复杂度：O(mn)O(mn)，其中 mm 和 nn 分别是矩阵的行数和列数。我们分配了一个与给定矩阵等大的数组，用于存储每个元素的左边连续 11 的数量。
