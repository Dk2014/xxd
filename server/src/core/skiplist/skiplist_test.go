package skiplist

import (
	"math/rand"
	"testing"
	"time"
)

var (
	sl   *SL
	data []Int

	benchData1k   []Int
	benchData1kSL *SL

	benchData1m   []Int
	benchData1mSL *SL
)

func init() {
	rand.Seed(time.Now().UnixNano())

	// general test
	for _, val := range rand.Perm(250) {
		data = append(data, Int(val))
	}

	sl = NewSkipList(len(data))

	// benchmark 1k data
	benchData1kSL = NewSkipList(1024)
	for i := 0; i < 1024; i++ {
		data := Int(rand.Int31())
		benchData1k = append(benchData1k, data)
		benchData1kSL.Insert(data)
	}

	// benchmark 1m data
	benchData1mSL = NewSkipList(1024 * 1024)
	for i := 0; i < 1024*1024; i++ {
		data := Int(rand.Int31())
		benchData1m = append(benchData1m, data)
		benchData1mSL.Insert(data)
	}
}

type Int int

func (val Int) Less(data Lesser) bool {
	return val < data.(Int)
}

func (val Int) Equal(data Lesser) bool {
	return val == data.(Int)
}

func TestInsert(t *testing.T) {
	for _, val := range data {
		sl.Insert(val)
	}
}

func TestFind(t *testing.T) {
	for _, val := range data {
		n, i := sl.Find(val)
		if n == nil || n.Data != val {
			t.Errorf("Can't find out %v, expect can", val)
		}
		if n != sl.Index(i) {
			t.Error("Method Find gave out a wrong index")
		}
	}
}

func TestIndex(t *testing.T) {
	pre := sl.Index(0)
	for i := 1; i < len(data); i++ {
		curr := sl.Index(i)
		if pre.Data.(Int) > curr.Data.(Int) {
			t.Error("List is not ordered")
		}
	}
}

func TestIter(t *testing.T) {
	sl.IterAll(func(n *Node) bool {
		dataP := n.Data.(Int)
		sl.Iter(n, func(n *Node) bool {
			if dataP > n.Data.(Int) {
				t.Error("List is not ordered")
			}
			return true
		})
		return true
	})
}

func TestRemove(t *testing.T) {
	for _, val := range data {
		if sl.Remove(val) == nil {
			t.Error("One element isn't in list while removing")
		}
	}
	sl.IterAll(func(n *Node) bool {
		t.Error("Ok, that's embarrassing")
		return true
	})
}

func TestMixed(t *testing.T) {
	testSeed := rand.Perm(250)
	testSeed = testSeed[89 : 89+25]

	for loop := 0; loop < 20; loop++ {

		// first reinitialize the skiplist
		TestInsert(t)
		// and make sure these read-only operation will not change anything
		TestFind(t)
		TestIndex(t)
		TestIter(t)

		for _, val := range testSeed {
			node, ind := sl.Find(Int(val))
			if node == nil {
				t.Error("Can't find those additional elements")
			}
			if sl.Index(ind) != node {
				t.Error("Index found out is not right")
			}

			sl.Remove(Int(val))

			node, _ = sl.Find(Int(val))
			if node != nil {
				t.Error("Didn't remove successfully")
			}
		}
	}
}

func Benchmark1KInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl := NewSkipList(len(benchData1k))
		for _, data := range benchData1k {
			sl.Insert(data)
		}
	}
}

func Benchmark1MInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl := NewSkipList(len(benchData1m))
		for _, data := range benchData1m {
			sl.Insert(data)
		}
	}
}

func Benchmark1KFind(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, data := range benchData1k {
			benchData1kSL.Find(data)
		}
	}
}

func Benchmark1MFind(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, data := range benchData1m {
			benchData1mSL.Find(data)
		}
	}
}

func Benchmark1KIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 1024; i++ {
			benchData1kSL.Index(i)
		}
	}
}

func Benchmark1MIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 1024*1024; i++ {
			benchData1mSL.Index(i)
		}
	}
}

func Benchmark1KRemove(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, data := range benchData1k {
			benchData1kSL.Remove(data)
		}
	}
}

func Benchmark1MRemove(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, data := range benchData1m {
			benchData1mSL.Remove(data)
		}
	}
}
