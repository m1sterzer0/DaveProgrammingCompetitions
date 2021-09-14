package rbtreemap

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func lt(a,b KEYTYPE) bool { return a < b }
func cmp(a,b KEYTYPE) int { return int(a)-int(b) }
func eq(a,b KEYTYPE)  bool { return a == b }

type map1 interface { 
	Len() int
	IsEmpty() bool
	Add(KEYTYPE,VALTYPE)
	MinKey() KEYTYPE
	MaxKey() KEYTYPE
	Clear()
	Contains(KEYTYPE) bool
	Delete(KEYTYPE) bool
	Lookup(KEYTYPE) (VALTYPE,bool)
}

type map1iter interface {
	Next() bool
	Prev() bool
	Key() KEYTYPE
	Value() VALTYPE
}

func basicMapTest(m map1, t *testing.T) {
	for i:=0;i<2;i++ {
		if m.Len() != 0 { t.Error(fmt.Sprintf("i:%v ERROR 1",i)); return }
		if !m.IsEmpty() { t.Error(fmt.Sprintf("i:%v ERROR 2",i)); return }
		m.Add(1, 100)
		m.Add(2, 200)
		m.Add(3, 300)
		if m.Len() != 3 { t.Error(fmt.Sprintf("i:%v ERROR 3",i)); return }
		if m.IsEmpty() { t.Error(fmt.Sprintf("i:%v ERROR 4",i)); return }
		if m.MinKey() != 1 { t.Error(fmt.Sprintf("i:%v ERROR 5",i)); return }
		if m.MaxKey() != 3 { t.Error(fmt.Sprintf("i:%v ERROR 6",i)); return }
		v,ok := m.Lookup(3); if !ok || v != 300 { t.Error(fmt.Sprintf("i:%v ERROR 7",i)); return }
		v,ok = m.Lookup(4); if ok { t.Error(fmt.Sprintf("i:%v ERROR 8",i)); return }
		if !m.Contains(3) { t.Error(fmt.Sprintf("i:%v ERROR 9",i)); return}
		ok = m.Delete(3)
		if !ok { t.Error(fmt.Sprintf("i:%v ERROR 10",i)); return }
		if m.Len() != 2 { t.Error(fmt.Sprintf("i:%v ERROR 11",i)); return }
		if m.MaxKey() != 2 { t.Error(fmt.Sprintf("i:%v ERROR 12",i)); return }
		m.Clear()
	}
}

func randMapTest(m map1, t *testing.T, fcheck1 func(*testing.T) bool, fcheck2 func(choosemin bool) map1iter, kmax, vmax, seed, nrounds, nsessions, niter int, dbg bool) {
	ref := make(map[int]int)
	rand.Seed(int64(seed))
	exp := m
	for round:=0;round<nrounds;round++ {
		if dbg { fmt.Printf("INFO: Starting round %v\n",round) }
		for session:=0;session<nsessions;session++ {
			addprob := 0.05 + 0.95 * rand.Float64()
			for iter:=0;iter<niter;iter++ {
				ap := rand.Float64()
				if ap < addprob {
					k := rand.Intn(kmax)
					v := rand.Intn(vmax)
					ref[k] = v
					exp.Add(KEYTYPE(k),VALTYPE(v))
				} else {
					k := rand.Intn(kmax)
					_,ok := ref[k]
					if ok { 
						//fmt.Printf("Real Deleting k:%v iter:%v\n",k,iter)
						delete(ref,k)
						ok2 := exp.Delete(KEYTYPE(k)) && !exp.Contains(KEYTYPE(k))
						if !ok2 { t.Error(fmt.Sprintf("round:%v session:%v iter:%v Delete error 1",round,session,iter)); return }
					} else {
						//fmt.Printf("Fake Deleting k:%v iter:%v\n",k,iter)
						ok2 := exp.Delete(KEYTYPE(k)) || exp.Contains(KEYTYPE(k))
						if ok2 { t.Error(fmt.Sprintf("round:%v session:%v iter:%v Delete error 2",round,session,iter)); return }
					}
				}
				reflen := len(ref)
				explen := exp.Len()
				//fmt.Printf("DBG: iter:%v reflen:%v explen:%v\n",iter,len(ref),exp.Len())
				if reflen != explen { t.Error(fmt.Sprintf("round:%v session:%v iter:%v Length mismatch reflen:%v explen:%v",round,session,iter,reflen,explen)); return  }
				if !fcheck1(t) { t.Error(fmt.Sprintf("round:%v session:%v iter:%v Structural check fail",round,session,iter)); return  }
			}

			if len(ref) > 0 {
				if session % 2 == 0 { 
					myminiter := fcheck2(true)
					if !checkMinIter(&ref,myminiter,t) { t.Error(fmt.Sprintf("round:%v session:%v Iterator check failed",round,session)); return  }
				} else {
					mymaxiter := fcheck2(false)
					if !checkMaxIter(&ref,mymaxiter,t) { t.Error(fmt.Sprintf("round:%v session:%v Iterator check failed",round,session)); return  }
				}
			}
		}
		exp.Clear()
		ref = make(map[int]int)
	}
}

func checkMinIter(ref *map[int]int, m map1iter, t *testing.T) bool {
	refkeys := make([]int,0)
	for k := range *ref { refkeys = append(refkeys,k) }
	sort.Slice( refkeys, func(i,j int)bool{return refkeys[i]<refkeys[j]} )
	ok := true
	for _,k := range refkeys {
		if int(m.Key()) != k { t.Error("Iterator key mismatch"); return false }
		if int(m.Value()) != (*ref)[k] { t.Error("Iterator value mismatch"); return false }
		ok = ok && m.Next()
	}
	if ok { t.Error("Iterator should have stopped"); return false }
	return true
}

func checkMaxIter(ref *map[int]int, m map1iter, t *testing.T) bool {
	refkeys := make([]int,0)
	for k := range *ref { refkeys = append(refkeys,k) }
	sort.Slice( refkeys, func(i,j int)bool{return refkeys[i]<refkeys[j]} )
	ok := true
	for i:=len(refkeys)-1;i>=0;i-- {
		k := refkeys[i]
		if int(m.Key()) != k { t.Error("Iterator key mismatch"); return false }
		if int(m.Value()) != (*ref)[k] { t.Error("Iterator value mismatch"); return false }
		ok = ok && m.Prev()
	}
	if ok { t.Error("Iterator should have stopped"); return false }
	return true
}

func TestBasicRbtreemap(t *testing.T) {
	m := NewRBTREEMAP(lt)
	basicMapTest(m,t)
}


func genRbtreemapChecker(exp *RBTREEMAP) (func (*testing.T) bool,func(bool) map1iter) {
	f := func (t *testing.T) bool {
		if exp.sz == 0 && exp.root != 0 { t.Error("BAD EMPTY ROOT"); return false }
		if exp.sz != 0 && exp.root == 0 { t.Error("BAD NON-EMPTY ROOT"); return false  }
		if !checkRedBlackProperties(exp) { t.Error("RED-BLACK FAIL"); return false  }
		if exp.tree[0].red { t.Error("NIL NODE should never be RED"); return false  }
		if exp.sz > 0 && len(exp.tree) <= int(exp.minidx) { t.Error("minidx out of range"); return false}
		if exp.sz > 0 && len(exp.tree) <= int(exp.maxidx) { t.Error("maxidx out of range"); return false}
		return true
	}

	f2 := func (minflag bool) map1iter {
		if minflag {
			f,_ := exp.MinIter(); return f
		} else { 
			f,_ := exp.MaxIter(); return f
		}
	}
	return f,f2
}

func checkRedBlackProperties(exp *RBTREEMAP) bool {
	if exp.sz == 0 { return true }
	if exp.tree[exp.root].red { return false } // Root should be black
	bcnt := 0; targ := -1
	var dfs func(n int32, lastred bool) bool
	dfs = func(n int32, lastred bool) bool {
		if exp.tree[n].red && lastred { return false }
		if !exp.tree[n].red { bcnt++ }
		if exp.tree[n].left == 0 { 
			if targ == -1 { targ = bcnt }
			if bcnt != targ { return false }
		} else {
			res := dfs(exp.tree[n].left,exp.tree[n].red)
			if !res { return false }
		}
		if exp.tree[n].right == 0 { 
			if targ == -1 { targ = bcnt }
			if bcnt != targ { return false }
		} else {
			res := dfs(exp.tree[n].right,exp.tree[n].red)
			if !res { return false }
		}
		if !exp.tree[n].red { bcnt-- }
		return true
	}
	return dfs(exp.root,false)
}


func TestRandRbtreemap(t *testing.T) {
	exp := NewRBTREEMAP(lt)
	f1,f2 := genRbtreemapChecker(exp)
	randMapTest(exp,t,f1,f2,1000,1000,8675309,5,20,10_000,false)
}

type benchmarkmap interface { Add(KEYTYPE,VALTYPE); Delete(KEYTYPE) bool; Lookup(KEYTYPE) (VALTYPE,bool); Clear() }

func domapInsertBenchmark(niter int,randomize bool,mymap benchmarkmap,b *testing.B) {
	rand.Seed(8675309)
	insertarr := make([]int,niter)
	for i := 0; i < niter; i++ { insertarr[i] = i }
	if randomize { 
		rand.Shuffle(len(insertarr), func (i,j int) { insertarr[i],insertarr[j] = insertarr[j],insertarr[i] })
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		mymap.Clear()
		for _,x := range(insertarr) { mymap.Add(KEYTYPE(x),VALTYPE(x)) }
	}
}

func domapLookupBenchmark(niter int,randomize bool,mymap benchmarkmap,b *testing.B) {
	rand.Seed(8675309)
	insertarr := make([]int,niter)
	for i := 0; i < niter; i++ { insertarr[i] = i }
	rand.Shuffle(len(insertarr), func (i,j int) { insertarr[i],insertarr[j] = insertarr[j],insertarr[i] })
	mymap.Clear()
	for _,x := range(insertarr) { mymap.Add(KEYTYPE(x),VALTYPE(x)) }
	if randomize {
		rand.Shuffle(len(insertarr), func (i,j int) { insertarr[i],insertarr[j] = insertarr[j],insertarr[i] })
	} else {
		sort.Slice(insertarr, func (i,j int) bool { return insertarr[i] < insertarr[j] })
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		for _,x := range(insertarr) { mymap.Lookup(KEYTYPE(x)) }
	}
}

func domapInsDelBenchmark(niter int,randomize bool,mymap benchmarkmap,b *testing.B) {
	rand.Seed(8675309)
	insertarr := make([]int,niter)
	deletearr := make([]int,niter)
	for i := 0; i < niter; i++ { insertarr[i] = i; deletearr[i] = i }
	if randomize { 
		rand.Shuffle(len(insertarr), func (i,j int) { insertarr[i],insertarr[j] = insertarr[j],insertarr[i] })
		rand.Shuffle(len(deletearr), func (i,j int) { deletearr[i],deletearr[j] = deletearr[j],deletearr[i] })
	}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		mymap.Clear()
		for _,x := range(insertarr) { mymap.Add(KEYTYPE(x),VALTYPE(x)) }
		for _,x := range(deletearr) { mymap.Delete(KEYTYPE(x)) }
	}
}

func BenchmarkRbtreemapSequentialInserts100k(b *testing.B) { domapInsertBenchmark(100_000,false,NewRBTREEMAP(lt),b) }
func BenchmarkRbtreemapSequentialInserts1M(b *testing.B) { domapInsertBenchmark(1_000_000,false,NewRBTREEMAP(lt),b) }
func BenchmarkRbtreemapRandomInserts100k(b *testing.B) { domapInsertBenchmark(100_000,true,NewRBTREEMAP(lt),b) }
func BenchmarkRbtreemapRandomInserts1M(b *testing.B) { domapInsertBenchmark(1_000_000,true,NewRBTREEMAP(lt),b) }
func BenchmarkRbtreemapSequentialLookup100k(b *testing.B) { domapLookupBenchmark(100_000,false,NewRBTREEMAP(lt),b) }
func BenchmarkRbtreemapSequentialLookup1M(b *testing.B) { domapLookupBenchmark(1_000_000,false,NewRBTREEMAP(lt),b) }
func BenchmarkRbtreemapRandomLookup100k(b *testing.B) { domapLookupBenchmark(100_000,true,NewRBTREEMAP(lt),b) }
func BenchmarkRbtreemapRandomLookup1M(b *testing.B) { domapLookupBenchmark(1_000_000,true,NewRBTREEMAP(lt),b) }
func BenchmarkRbtreemapSequentialInsDel100k(b *testing.B) { domapInsDelBenchmark(100_000,false,NewRBTREEMAP(lt),b) }
func BenchmarkRbtreemapSequentialInsDel1M(b *testing.B) { domapInsDelBenchmark(1_000_000,false,NewRBTREEMAP(lt),b) }
func BenchmarkRbtreemapRandomInsDel100k(b *testing.B) { domapInsDelBenchmark(100_000,true,NewRBTREEMAP(lt),b) }
func BenchmarkRbtreemapRandomInsDel1M(b *testing.B) { domapInsDelBenchmark(1_000_000,true,NewRBTREEMAP(lt),b) }


