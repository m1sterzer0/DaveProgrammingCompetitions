import sys

sys.setrecursionlimit(10000000)
from collections import deque, namedtuple

## Input crap
infile = sys.stdin
intokens = deque()
def getTokens(): 
    while not intokens:
        for c in infile.readline().rstrip().split() :
            intokens.append(c)    
def gs(): getTokens(); return intokens.popleft()
def gi(): return int(gs())
def gf(): return float(gs())
def gbs(): return [c for c in gs()]
def gis(n): return [gi() for i in range(n)]
def ia(m): return [0] * m
def iai(m,v): return [v] * m
def twodi(n,m,v): return [iai(m,v) for i in range(n)]
def fill2(m) : r = gis(2*m); return r[0::2],r[1::2]
def fill3(m) : r = gis(3*m); return r[0::3],r[1::3],r[2::3]
def fill4(m) : r = gis(4*m); return r[0::4],r[1::4],r[2::4],r[3::4]
MOD = 998244353
##MOD = 1000000007

class fenwicktree :
    def __init__(self,n=1) : self.n = n; self.tot = 0; self.bit = [0] * (n+3)
    def inc(self,idx,val=1) :
        idx += 1
        while idx <= self.n+1 : self.bit[idx] += val;idx += idx & (-idx)
        self.tot += val
    def prefixsum(self,idx) :
        idx += 1; ans = 0
        while idx > 0 : ans += self.bit[idx]; idx -= idx&(-idx)
        return ans
    def clear(self) :
        for i in range(self.n+3) : self.bit[i] = 0
        self.tot = 0
    def dec(self,idx,val=1) : self.inc(idx,-val)
    def incdec(self,left,right,val) : self.inc(left,val); self.dec(right,val)
    def suffixsum(self,idx) : return self.tot - self.prefixsum(idx-1)
    def rangesum(self,left,right)  : return self.prefixsum(right) - self.prefixsum(left-1)




def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N = gi(); A,B = fill2(N); Q = gi()
    qt,qx,qy = [0]*Q, [0]*Q, [0]*Q
    for i in range(Q) :
        qt[i],qx[i] = gi(),gi()
        if qt[i] == 1 or qt[i] == 2 : qx[i] -= 1
        if qt[i] != 3 : qy[i] = gi()
    
    ## Score compression
    rawscores = A + [qy[i] for i in range(Q) if qt[i] == 1]
    scores = sorted(list(set(rawscores)))
    score2idx = {k:v for v,k in enumerate(scores)}

    ## fenwicktree for count and scores
    scorecnt = fenwicktree(len(scores)+5)
    scoresum = fenwicktree(len(scores)+5)
    
    for a,b in zip(A,B) :
        idx = score2idx[a]
        scorecnt.inc(idx,b)
        scoresum.inc(idx,a*b)

    for t,x,y in zip(qt,qx,qy) :
        if t == 1 :
            oldidx = score2idx[A[x]]
            newidx = score2idx[y]
            scorecnt.dec(oldidx,B[x])
            scorecnt.inc(newidx,B[x])
            scoresum.dec(oldidx,A[x]*B[x])
            scoresum.inc(newidx,y*B[x])
            A[x] = y
        elif t == 2 :
            idx = score2idx[A[x]]
            scorecnt.inc(idx,y-B[x])
            scoresum.inc(idx,(y-B[x])*A[x])
            B[x] = y
        else :
            ans = -1
            if scorecnt.suffixsum(0) >= x :
                (l,r) = 0,len(scores)
                while (r-l) > 1 : m = (r+l)>>1; (l,r) = (m,r) if scorecnt.suffixsum(m) >= x else (l,m)
                ans = scoresum.suffixsum(l+1) + (x - scorecnt.suffixsum(l+1)) * scores[l]
            print(ans)


if __name__ == "__main__" :
    main()

