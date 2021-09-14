package rbtreemap

type KEYTYPE int

//START HERE
type RBTREEMULTIMAPnode struct {
	left, right, up int32
	red             bool
	key             KEYTYPE
	count           int
}
type RBTREEMULTIMAP struct {
	lessthan func(a, b KEYTYPE) bool
	tree     []RBTREEMULTIMAPnode
	root     int32
	recycler []int32
	sz       int
	minidx   int32
	maxidx   int32
}
type RBTREEMULTIMAPIterator interface {
	Next() (ok bool)
	Prev() (ok bool)
	Key() KEYTYPE
	Count() int
}
type RBTREEMULTIMAPiter struct {
	cur    int32
	key    KEYTYPE
	count  int
	rbtree *RBTREEMULTIMAP
}

func (i *RBTREEMULTIMAPiter) Key() KEYTYPE { return i.key }
func (i *RBTREEMULTIMAPiter) Count() int   { return i.count }
func (i *RBTREEMULTIMAPiter) Next() bool {
	rbtree := i.rbtree
	v := rbtree.nextidx(i.cur)
	if v == 0 {
		return false
	}
	i.cur, i.key, i.count = v, rbtree.tree[v].key, rbtree.tree[v].count
	return true
}
func (i *RBTREEMULTIMAPiter) Prev() bool {
	rbtree := i.rbtree
	v := rbtree.previdx(i.cur)
	if v == 0 {
		return false
	}
	i.cur, i.key, i.count = v, rbtree.tree[v].key, rbtree.tree[v].count
	return true
}
func NewRBTREEMULTIMAP(lessthan func(a, b KEYTYPE) bool) *RBTREEMULTIMAP {
	q := &RBTREEMULTIMAP{lessthan, make([]RBTREEMULTIMAPnode, 2), int32(0), make([]int32, 0), 0, 0, 0}
	q.tree[0].left, q.tree[0].right, q.tree[0].up, q.tree[0].red = 0, 0, 0, false
	q.recycler = append(q.recycler, 1)
	return q
}
func (q *RBTREEMULTIMAP) Add(k KEYTYPE) {

	// Special case for size 0
	if q.sz == 0 {
		z := q.getNewNodenum()
		tree := q.tree
		q.minidx, q.maxidx, q.sz, q.root = z, z, q.sz+1, z
		tree[z].key, tree[z].count, tree[z].up, tree[z].left, tree[z].right, tree[z].red = k, 1, 0, 0, 0, false
		return
	}

	y, cmp := q.findInsertionPoint(k)
	if cmp == 0 {
		q.tree[y].count++
		q.sz += 1
		return
	}
	z := q.getNewNodenum()
	q.sz += 1
	tree := q.tree
	tree[z].key, tree[z].count, tree[z].up, tree[z].left, tree[z].right, tree[z].red = k, 1, y, 0, 0, true
	if cmp < 0 {
		tree[y].left = z
	} else {
		tree[y].right = z
	}

	// Take care of max/min queries proactively
	if q.sz == 0 || q.lessthan(k, tree[q.minidx].key) {
		q.minidx = z
	}
	if q.sz == 0 || q.lessthan(tree[q.maxidx].key, k) {
		q.maxidx = z
	}

	// Fix-up the RB tree
	var p, g, u int32
	// Loop invariant is that z,p,g are red,red,black respectively at the start of the loop
	for p = tree[z].up; tree[p].red; p = tree[z].up { //While parent is red
		g = tree[p].up // since p is red, it must be real, g might be NULL, since root could have been recolored
		if g == 0 {
			break
		} // p is root, and root will be recolored to black in last step, so we are done.
		if p == tree[g].left {
			u = tree[g].right
			if tree[u].red {
				tree[p].red, tree[u].red, tree[g].red, z = false, false, true, g
				continue
			}
			if z == tree[p].right {
				z = p
				q.rotleft(z)
				p = tree[z].up
			}
			q.rotright(g)
			tree[g].red, tree[p].red = true, false
		} else { // MIRROR CASES
			u = tree[g].left
			if tree[u].red {
				tree[p].red, tree[u].red, tree[g].red, z = false, false, true, g
				continue
			}
			if z == tree[p].left {
				z = p
				q.rotright(z)
				p = tree[z].up
			}
			q.rotleft(g)
			tree[g].red, tree[p].red = true, false
		}
	}
	// Color the root node black at the end
	tree[q.root].red = false
}

func (q *RBTREEMULTIMAP) Delete(k KEYTYPE) bool {
	if q.sz == 0 {
		return false
	}
	z, cmp := q.findInsertionPoint(k)
	if cmp != 0 {
		return false
	} else if q.tree[z].count > 1 {
		q.tree[z].count--
		q.sz--
		return true
	}
	q.sz--
	q.recycler = append(q.recycler, z)

	// Patch up the min/max
	if q.sz > 0 && !q.lessthan(q.tree[q.minidx].key, k) {
		q.minidx = q.nextidx(q.minidx)
	}
	if q.sz > 0 && !q.lessthan(k, q.tree[q.maxidx].key) {
		q.maxidx = q.previdx(q.maxidx)
	}

	// Special case when we are deleting the last element
	if q.sz == 0 {
		q.root = 0
		return true
	}
	// FROM CLRS 3rd Ed pg 324
	tree := q.tree
	var x int32
	y, y_orig_red := z, tree[z].red
	if tree[z].left == 0 {
		x = tree[z].right
		q.rbTransplant(z, x)
	} else if tree[z].right == 0 {
		x = tree[z].left
		q.rbTransplant(z, x)
	} else {
		y = q.findminidx(tree[z].right)
		y_orig_red = tree[y].red
		x = tree[y].right
		if tree[y].up == z {
			tree[x].up = y // Why this step? if x's tree[x].up should already be y
		} else {
			q.rbTransplant(y, x)
			tree[y].right = tree[z].right
			tree[tree[y].right].up = y
		}
		q.rbTransplant(z, y)
		tree[y].left = tree[z].left
		tree[tree[y].left].up = y
		tree[y].red = tree[z].red
	}

	if !y_orig_red {
		for q.root != x && !tree[x].red {
			p := tree[x].up
			if tree[p].left == x {
				s := tree[p].right // s is a sibling of x
				if tree[s].red {
					tree[s].red = false
					tree[p].red = true
					q.rotleft(p)
					s = tree[p].right
				}
				c := tree[s].left
				d := tree[s].right
				if !tree[c].red && !tree[d].red {
					tree[s].red = true
					x = p
				} else {
					if !tree[d].red {
						tree[c].red = false
						tree[s].red = true
						q.rotright(s)
						s = tree[p].right
					}
					tree[s].red = tree[p].red
					tree[p].red = false
					tree[tree[s].right].red = false
					q.rotleft(p)
					x = q.root
				}
			} else {
				s := tree[p].left // s is a sibling of x
				if tree[s].red {
					tree[s].red = false
					tree[p].red = true
					q.rotright(p)
					s = tree[p].left
				}
				c := tree[s].right
				d := tree[s].left
				if !tree[c].red && !tree[d].red {
					tree[s].red = true
					x = p
				} else {
					if !tree[d].red {
						tree[c].red = false
						tree[s].red = true
						q.rotleft(s)
						s = tree[p].left
					}
					tree[s].red = tree[p].red
					tree[p].red = false
					tree[tree[s].left].red = false
					q.rotright(p)
					x = q.root
				}
			}
		}
		tree[x].red = false
	}

	return true
}

//type RBTREEMULTIMAP struct { lessthan func(a,b KEYTYPE) bool;tree []RBTREEMULTIMAPnode; root int32; recycler []int32; sz int; }
func (q *RBTREEMULTIMAP) Clear() {
	q.tree, q.root, q.recycler, q.sz = q.tree[:2], 0, q.recycler[:0], 0
	q.recycler = append(q.recycler, int32(1))
}
func (q *RBTREEMULTIMAP) IsEmpty() bool           { return q.sz == 0 }
func (q *RBTREEMULTIMAP) Contains(k KEYTYPE) bool { _, cmp := q.findInsertionPoint(k); return cmp == 0 }
func (q *RBTREEMULTIMAP) Count(k KEYTYPE) int {
	z, cmp := q.findInsertionPoint(k)
	if cmp != 0 {
		return 0
	}
	return q.tree[z].count
}

func (q *RBTREEMULTIMAP) Len() int { return q.sz }
func (q *RBTREEMULTIMAP) MinKey() (k KEYTYPE) {
	if q.sz == 0 {
		panic("Called MinKey on an empty RBTREEMULTIMAP")
	}
	return q.tree[q.minidx].key
}
func (q *RBTREEMULTIMAP) MaxKey() (k KEYTYPE) {
	if q.sz == 0 {
		panic("Called MaxKey on an empty RBTREEMULTIMAP")
	}
	return q.tree[q.maxidx].key
}
func (q *RBTREEMULTIMAP) LowerBound(k KEYTYPE) (KEYTYPE, bool) {
	var def KEYTYPE
	if q.sz == 0 {
		return def, false
	}
	idx, pos := q.findInsertionPoint(k)
	if pos == 1 {
		idx = q.nextidx(idx)
	}
	if idx <= 0 {
		return def, false
	}
	return q.tree[idx].key, true
}
func (q *RBTREEMULTIMAP) UpperBound(k KEYTYPE) (KEYTYPE, bool) {
	var def KEYTYPE
	if q.sz == 0 {
		return def, false
	}
	idx, pos := q.findInsertionPoint(k)
	if pos != -1 {
		idx = q.nextidx(idx)
	}
	if idx <= 0 {
		return def, false
	}
	return q.tree[idx].key, true
}
func (q *RBTREEMULTIMAP) LowerBoundIter(k KEYTYPE) (RBTREEMULTIMAPIterator, bool) {
	if q.sz == 0 {
		return nil, false
	}
	idx, pos := q.findInsertionPoint(k)
	if pos == 1 {
		idx = q.nextidx(idx)
	}
	if idx <= 0 {
		return nil, false
	}
	return &RBTREEMULTIMAPiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *RBTREEMULTIMAP) UpperBoundIter(k KEYTYPE) (RBTREEMULTIMAPIterator, bool) {
	if q.sz == 0 {
		return nil, false
	}
	idx, pos := q.findInsertionPoint(k)
	if pos != -1 {
		idx = q.nextidx(idx)
	}
	if idx <= 0 {
		return nil, false
	}
	return &RBTREEMULTIMAPiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}

func (q *RBTREEMULTIMAP) FindIter(k KEYTYPE) (RBTREEMULTIMAPIterator, bool) {
	if q.sz == 0 {
		return nil, false
	}
	idx, pos := q.findInsertionPoint(k)
	if pos != 0 {
		return nil, false
	}
	return &RBTREEMULTIMAPiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *RBTREEMULTIMAP) MinIter() (RBTREEMULTIMAPIterator, bool) {
	if q.sz == 0 {
		return nil, false
	}
	idx := q.findminidx(q.root)
	return &RBTREEMULTIMAPiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *RBTREEMULTIMAP) MaxIter() (RBTREEMULTIMAPIterator, bool) {
	if q.sz == 0 {
		return nil, false
	}
	idx := q.findmaxidx(q.root)
	return &RBTREEMULTIMAPiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *RBTREEMULTIMAP) rbTransplant(u, v int32) {
	// Note v may be nil, but we can still set tree[v].up, which we exploit
	tree := q.tree
	if tree[u].up == 0 {
		q.root = v
	} else {
		p := tree[u].up
		if u == tree[p].left {
			tree[p].left = v
		} else {
			tree[p].right = v
		}
	}
	tree[v].up = tree[u].up
}

func (q *RBTREEMULTIMAP) findInsertionPoint(k KEYTYPE) (int32, int8) {
	n, lt, tree := q.root, q.lessthan, q.tree
	for {
		nkey := tree[n].key
		if lt(nkey, k) {
			r := tree[n].right
			if r == 0 {
				return n, 1
			}
			n = r
		} else if lt(k, nkey) {
			l := tree[n].left
			if l == 0 {
				return n, -1
			}
			n = l
		} else {
			return n, 0
		}
	}
}

func (q *RBTREEMULTIMAP) findmaxidx(n1 int32) int32 {
	tree := q.tree
	for {
		xx := tree[n1].right
		if xx == 0 {
			break
		}
		n1 = xx
	}
	return n1
}

func (q *RBTREEMULTIMAP) findminidx(n1 int32) int32 {
	tree := q.tree
	for {
		xx := tree[n1].left
		if xx == 0 {
			break
		}
		n1 = xx
	}
	return n1
}

func (q *RBTREEMULTIMAP) nextidx(cur int32) int32 {
	last := int32(-2)
	tree := q.tree
	rr := tree[cur].right
	if rr > 0 {
		return q.findminidx(rr)
	}
	for {
		last, cur = cur, tree[cur].up
		if cur == 0 || tree[cur].left == last {
			break
		}
	}
	return cur
}

func (q *RBTREEMULTIMAP) previdx(cur int32) int32 {
	last := int32(0)
	tree := q.tree
	ll := tree[cur].left
	if ll > 0 {
		return q.findmaxidx(ll)
	}
	for {
		last, cur = cur, tree[cur].up
		if cur == 0 || tree[cur].right == last {
			break
		}
	}
	return cur
}

func (q *RBTREEMULTIMAP) rotleft(x int32) {
	tree := q.tree
	y := tree[x].right
	p := tree[x].up
	tree[x].right = tree[y].left
	if tree[y].left != 0 {
		tree[tree[y].left].up = x
	}
	tree[y].up = p
	if p == 0 {
		q.root = y
	} else if x == tree[p].left {
		tree[p].left = y
	} else {
		tree[p].right = y
	}
	tree[y].left = x
	tree[x].up = y
}

func (q *RBTREEMULTIMAP) rotright(x int32) {
	tree := q.tree
	y := tree[x].left
	p := tree[x].up
	tree[x].left = tree[y].right
	if tree[y].right != 0 {
		tree[tree[y].right].up = x
	}
	tree[y].up = p
	if p == 0 {
		q.root = y
	} else if x == tree[p].right {
		tree[p].right = y
	} else {
		tree[p].left = y
	}
	tree[y].right = x
	tree[x].up = y
}

func (q *RBTREEMULTIMAP) getNewNodenum() int32 {
	l := len(q.recycler)
	newnode := q.recycler[l-1]
	q.recycler = q.recycler[:l-1]
	if l == 1 {
		q.tree = append(q.tree, RBTREEMULTIMAPnode{})
		q.recycler = append(q.recycler, int32(len(q.tree)-1))
	}
	return newnode
}
