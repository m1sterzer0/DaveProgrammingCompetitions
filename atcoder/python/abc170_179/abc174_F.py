
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class fenwicktree :
    bit = [0]; n = 0; tot = 0

    def __init__(self,n=1) :
        self.n = n; self.tot = 0
        self.bit = [0] * (n+1)

    def clear(self) :
        for i in range(self.n) : self.bit[i] = 0
        self.tot = 0

    def inc(self,idx,val=1) :
        while idx <= self.n :
            self.bit[idx] += val
            idx += idx & (-idx)
        self.tot += val

    def dec(self,idx,val=1) : self.inc(idx,-val)

    def incdec(self,left,right,val) :
        self.inc(left,val); self.dec(right,val)

    def prefixsum(self,idx) :
        if idx < 1 : return 0
        ans = 0
        while idx > 0 :
            ans += self.bit[idx]
            idx -= idx&(-idx)
        return ans

    def suffixsum(self,idx) : return self.tot - self.prefixsum(idx-1)
    def rangesum(self,left,right)  : return self.prefixsum(right) - self.prefixsum(left-1)


## Process queries in increasing order of right coordinate
## Use a fenwick tree for range sums.
## KEY TRICK: When we encounter a color we have seen before, we decrement the
##     old color position and increase the new color position, only counting the new
##     color at the rightmost position.  This works because we are processing the queries
##     in order of their rightmost coordinate.
## PERF TRICK: Sorting tuples in pypy is a bit slow.  Sorting integers is fast.  We can
##     bitpack our data into an integer, and this helps a ton.

def solve(N,Q,c,l,r) :
    qq = [r[i]<<40 | l[i]<<20 | i for i in range(Q)]
    qq.sort()
    ans = [0] * Q; ft = fenwicktree(N+1); last = [0] * (N+1)
    ptr = 0
    for (xx) in qq :
        rr = xx >> 40
        ll = (xx >> 20) & 0xfffff
        i = xx & 0xfffff
        #print(f"DBG: rr:{rr} ll:{ll} i:{i}")
        while ptr < rr :
            #print(f"DBG: ptr:{ptr} r:{r} c:{c} last:{c}")
            ptr += 1; cc = c[ptr]; lcc = last[cc]
            if lcc > 0 : ft.inc(lcc,-1)
            ft.inc(ptr); last[cc] = ptr
        ans[i] = ft.rangesum(ll,rr)
    return "\n".join(str(x) for x in ans)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,Q = [int(x) for x in infile.readline().split()]
    c = [0] * (N+1)
    c[1:] = [int(x) for x in infile.readline().split()]
    l = [0] * Q
    r = [0] * Q
    lines = infile.readlines()
    for i in range(Q) :
        l[i],r[i] = [int(x) for x in lines[i].split()]
    ans = solve(N,Q,c,l,r)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    #main()
    main()
    sys.stdout.flush()

