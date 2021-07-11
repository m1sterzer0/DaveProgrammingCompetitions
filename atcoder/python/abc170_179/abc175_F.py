
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

def dopush(mh,d,dist,pre,suf) :
    ll = min(len(pre),len(suf)); x = None
    if len(pre) < len(suf) :
        if not suf.startswith(pre) : return
        residual = suf[len(pre):]
        if residual == residual[::-1] : residual = ""
        x = ("",residual)
    else :
        if not pre.startswith(suf) : return
        residual = pre[len(suf):]
        if residual == residual[::-1] : residual = ""
        x = (residual,"")
    if x in dist : return
    mh.push(d,x) 

def solve(N,S,C) :
    mh = minHeapEnh()
    for i in range(N) :
        if S[i] == S[i][::-1] : 
            mh.push(C[i],("",""))
        else :
            mh.push(C[i],(S[i],""))
    dist = {}
    while not mh.isempty() :
        (d,(pre,suf)) = mh.pop()
        if len(pre) + len(suf) <= 1 : return d
        dist[(pre,suf)] = d
        if suf == "" :
            for i in range(N) :
                dopush(mh,d+C[i],dist,pre,S[i][::-1])
        else :
            for i in range(N) :
                dopush(mh,d+C[i],dist,S[i],suf)
    return -1

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    S = []; C = []
    for i in range(N) : 
        xx = gss()
        S.append(xx[0])
        C.append(int(xx[1]))
    ans = solve(N,S,C)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

