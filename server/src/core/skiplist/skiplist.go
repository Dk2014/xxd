package skiplist

import (
	"math"
	"math/rand"
)

type Lesser interface {
	Less(Lesser) bool
	Equal(Lesser) bool
}

// Sentinel tail of the skip list and always larger than other
type sentinelTail struct{}

func (val sentinelTail) Less(data Lesser) bool {
	return false
}

func (val sentinelTail) Equal(data Lesser) bool {
	return false
}

// Sentinel head of the skip list and always smaller than other
type sentinelHead struct{}

func (val sentinelHead) Less(data Lesser) bool {
	return true
}

func (val sentinelHead) Equal(data Lesser) bool {
	return false
}

type SL struct {
	head        *Node //keep trace of all "express lane" layers, does NOT keep any useful data
	growFactory float32
	maxHigh     int
}

type Node struct {
	next  []*Node
	width []int
	Data  Lesser
}

func New(growFactory float32, exceptedMax int) *SL {
	maxHigh := int(math.Log2(float64(exceptedMax))) + 1

	// head node
	head := &Node{
		Data:  sentinelHead{},
		next:  make([]*Node, maxHigh, maxHigh),
		width: make([]int, maxHigh, maxHigh),
	}

	// tail node
	tailNode := &Node{
		Data:  sentinelTail{},
		next:  nil,
		width: nil,
	}

	// link tail to head
	for i := 0; i < maxHigh; i++ {
		head.next[i] = tailNode
		head.width[i] = 0
	}

	return &SL{
		maxHigh:     maxHigh,
		growFactory: growFactory,
		head:        head,
	}
}

func NewSkipList(expectMax int) *SL {
	return New(0.25, expectMax)
}

func (sl *SL) Insert(data Lesser) {
	// find node that node.next.data >= data for each level
	lNodes := make([]*Node, sl.maxHigh, sl.maxHigh)
	// and the width from the previous level node
	lWidth := make([]int, sl.maxHigh, sl.maxHigh)

	node := sl.head
	for i := sl.maxHigh - 1; i >= 0; i-- {
		width := 0
		for len(node.next) > 0 && node.next[i].Data.Less(data) {
			width += node.width[i]
			node = node.next[i]
		}
		lWidth[i] = width
		lNodes[i] = node
	}

	// random a high
	high := 1
	for high < sl.maxHigh && rand.Float32() <= sl.growFactory {
		high++
	}

	newNode := &Node{
		Data:  data,
		next:  make([]*Node, high, high),
		width: make([]int, high, high),
	}

	// cut down the bottom level
	newNode.width[0] = 1
	newNode.next[0] = lNodes[0].next[0]
	lNodes[0].next[0] = newNode

	// iterate the rest level
	for i := 1; i < high; i++ {
		// cut down the width of level i
		newNode.width[i] = lNodes[i].width[i] - (lWidth[i-1] + lNodes[i-1].width[i-1]) + 1
		lNodes[i].width[i] = lWidth[i-1] + lNodes[i-1].width[i-1]
		// link in list
		newNode.next[i] = lNodes[i].next[i]
		lNodes[i].next[i] = newNode
	}

	// the rest level
	for i := high; i < sl.maxHigh; i++ {
		lNodes[i].width[i]++
	}
}

func (sl *SL) Remove(data Lesser) *Node {
	// find first node that node.next.data >= data for each level
	lNodes := make([]*Node, sl.maxHigh, sl.maxHigh)

	node := sl.head
	for i := sl.maxHigh - 1; i >= 0; i-- {
		for len(node.next) > 0 && node.next[i].Data.Less(data) {
			node = node.next[i]
		}
		lNodes[i] = node
	}

	// didn't find specified element
	if !lNodes[0].next[0].Data.Equal(data) {
		return nil
	}

	excluded := lNodes[0].next[0]

	// remove node
	for i := 0; i < sl.maxHigh; i++ {
		if lNodes[i].next[i].Data.Equal(data) {
			lNodes[i].width[i] += lNodes[i].next[i].width[i]
			lNodes[i].next[i] = lNodes[i].next[i].next[i]
		}
		lNodes[i].width[i]--
	}

	return excluded
}

func (sl *SL) Find(data Lesser) (*Node, int) {
	width := 0
	node := sl.head
	for i := sl.maxHigh - 1; i >= 0; i-- {
		for len(node.next) > 0 && (node.next[i].Data.Less(data) || node.next[i].Data.Equal(data)) {
			width += node.width[i]
			node = node.next[i]
		}
	}

	if node.Data.Equal(data) {
		return node, width
	}

	return nil, -1
}

func (sl *SL) Index(index int) *Node {
	width := 0
	node := sl.head
	for i := sl.maxHigh - 1; i >= 0; i-- {
		for len(node.next) > 0 && width+node.width[i] <= index {
			width += node.width[i]
			node = node.next[i]
		}
	}

	if width == index && len(node.next) > 0 {
		return node
	}

	return nil
}

func (sl *SL) IterAll(curiousNext func(*Node) bool) {
	sl.Iter(sl.head.next[0], curiousNext)
}

func (sl *SL) Iter(begin *Node, curiousNext func(*Node) bool) {
	for node := begin; !sl.IsTail(node) && curiousNext(node); node = node.next[0] {
	}
}

func (sl *SL) IsTail(tail *Node) bool {
	return len(tail.next) <= 0
}
