from os import posix_fadvise
import random
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## Implicitly supports decrease key if we repush a node with a lower dist.
## Heap stores objects as a tuple (d,obj)
# 0 --> 1,2 
# 1 --> 3,4
# 2 --> 5,6
class minHeapEnh :
    vt = []; pos = {}
    def __init__(self) : pass
    def _swap(mh,i,j) :
        (n1,n2) = (mh.vt[i][1],mh.vt[j][1])
        mh.pos[n2],mh.pos[n1] = i,j
        mh.vt[i],mh.vt[j] = mh.vt[j],mh.vt[i]
    def _bubbleup(mh,i) :
        if i == 0 : return
        j = (i-1) >> 1
        if mh.vt[i] < mh.vt[j] : mh._swap(i,j); mh._bubbleup(j)
    def _bubbledown(mh,i) :
        ll = len(mh.vt)
        l = (i<<1) + 1; r = l+1
        res1 = l >= ll or not (mh.vt[i] > mh.vt[l])
        res2 = r >= ll or not (mh.vt[i] > mh.vt[r])
        if res1 and res2 : return
        if res2 or not res1 and not mh.vt[l] > mh.vt[r] :
            mh._swap(i,l); mh._bubbledown(l)
        else :
            mh._swap(i,r); mh._bubbledown(r)
    def push(mh,d,n) :
        if n in mh.pos :
            idx = mh.pos[n]
            n2 = mh.vt[idx]
            if d < n2[0] : mh.vt[idx] = (d,n); mh._bubbleup(idx)
        else :
            mh.vt.append((d,n))
            idx = len(mh.vt)-1
            mh.pos[n] = idx
            mh._bubbleup(idx)
    def pop(mh) :
        ans = mh.vt[0]; del mh.pos[ans[1]]
        n2 = mh.vt.pop()
        if len(mh.vt) >= 1 :
            mh.pos[n2[1]] = 0
            mh.vt[0] = n2
            mh._bubbledown(0)
        return ans
    def isempty(mh) :
        return len(mh.vt) == 0

def checkheap(v) :
    ll = len(v)
    for i in range(ll) :
        l = (i<<1) + 1; r = l+1
        if l < ll and v[i] > v[l] : print(f"ERROR: i:{i} l:{l} v[i]:{v[i]} v[l]:{v[l]}")
        if r < ll and v[i] > v[r] : print(f"ERROR: i:{i} r:{r} v[i]:{v[i]} v[r]:{v[r]}")

if __name__ == '__main__' :
    random.seed(19006492568)
    for i in range(100) :
        for iter in (10,100,1000) :
            mh = minHeapEnh()
            writeChance = random.uniform(0.45,0.55)
            for j in range(iter) :
                if random.random() > writeChance and not mh.isempty() :
                    mh.pop()
                else :
                    mh.push(random.randrange(1000000000),random.randrange(iter))
                print(f"i:{i} iter:{iter} j:{j}")
                checkheap(mh.vt)

