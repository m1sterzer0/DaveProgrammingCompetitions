import sys
from collections import deque

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
INF = 1 << 62
##MOD = 1000000007

def minheaplt (a,b) : return a < b
class minheap :
    def __init__(self,lt=minheaplt) :
        self.lt = lt
        self.h  = []
    def isEmpty(self) : return len(self.h) == 0
    def clear(self) : self.h = self.h[:0]
    def len(self) : return len(self.h)
    def push(self,v) : self.h.append(v); self._siftdown(0,len(self.h)-1)
    def head(self) : return self.h[0]
    def pop(self) : 
        v1 = self.h[0]; l = len(self.h)
        if l == 1 : self.h.pop(); return v1
        l-=1; self.h[0] = self.h[l]; self.h.pop(); self._siftup(0); return v1
    def heapify(self,v) :
        for vv in v : self.h.append(v)
        n = len(self.h)
        for i in range(n//2-1,-1,-1) : self._siftup(i)
    def _siftdown(self,startpos,pos) :
        newitem = self.h[pos]
        while pos > startpos :
            ppos = (pos-1)>>1; p = self.h[ppos]
            if not self.lt(newitem,p) : break
            self.h[pos],pos = p,ppos
        self.h[pos] = newitem
    def _siftup(self,pos) :
        endpos,startpos,newitem,chpos = len(self.h),pos,self.h[pos],2*pos+1
        while chpos < endpos :
            rtpos = chpos+1
            if rtpos < endpos and not self.lt(self.h[chpos],self.h[rtpos]) : chpos = rtpos
            self.h[pos],pos = self.h[chpos],chpos; chpos = 2*pos+1
        self.h[pos] = newitem; self._siftdown(startpos,pos) 

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        N,M = gi(),gi()
        S,W,T = [],[],[]
        for i in range(N) : r = gis(3*M); S.append(r[0::3]); W.append(r[1::3]); T.append(r[2::3])
        darr = [[INF]*(2*M) for _ in range(2*N)]
        mh = minheap()
        mh.push((0,2*N-1,0))
        while not mh.isEmpty() :
            (d,i,j) = mh.pop()
            if darr[i][j] != INF : continue
            darr[i][j] = d
            s,w,t = S[i//2][j//2],W[i//2][j//2],T[i//2][j//2]
            t = t % (s+w) - (s+w)
            nstime = d+1 if (d-t) % (s+w) < s else t + (d-t)//(s+w)*(s+w) + (s+w) + 1
            ewtime = d+1 if (d-t) % (s+w) >= s else t + (d-t)//(s+w)*(s+w) + s + 1
            if i&1 == 0 : ## South side of intersection
                if i-1 >= 0 : mh.push((d+2,i-1,j))
                mh.push((nstime,i+1,j))
            else :  ## North side of intersection
                mh.push((nstime,i-1,j))
                if i+1 < 2*N : mh.push((d+2,i+1,j))
            if j&1 == 0 : ## West side of intersection
                if j-1 >= 0 : mh.push((d+2,i,j-1))
                mh.push((ewtime,i,j+1))
            else : ## East side of intersection
                mh.push((ewtime,i,j-1))
                if j+1 < 2*M : mh.push((d+2,i,j+1))
        print(f"Case #{tt}: {darr[0][2*M-1]}")

if __name__ == "__main__" :
    main()

