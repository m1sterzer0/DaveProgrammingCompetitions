
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

def encnode(x,y,dir,W) : return dir + 4 * y + 4 * W * x
def decnode(enc,W)     : return (enc // (4*W), (enc // 4) % W , enc % 4)

#def solve(H,W,K,x1,y1,x2,y2,bd) :
#    x1 -= 1; y1 -= 1; x2 -= 1; y2 -= 1
#    inf = 10**18
#    dist = [inf] * H * W * 4
#    mh = minHeapEnh()
#    mh.push(0,encnode(x1,y1,0,W))
#    while not mh.isempty() :
#        (d,enc) = mh.pop()
#        (x,y,dir) = decnode(enc,W)
#        #print(f"DBG: x:{x} y:{y} dir:{dir} d:{d}")
#        if x == x2 and y == y2 : return (d + K - 1) // K
#        dist[enc] = d
#        enc1 = encnode(x,y,(dir+1)%4,W)
#        if dist[enc1] == inf : mh.push((d + K - 1) // K * K, enc1)
#        for (xx,yy,nd) in ((x-1,y,0),(x+1,y,1),(x,y-1,2),(x,y+1,3)) :
#            if nd != dir or xx < 0 or xx >= H or yy < 0 or yy >=W or bd[xx][yy] == '@': continue
#            enc2 = encnode(xx,yy,nd,W)
#            if dist[enc2] == inf : mh.push(d+1,enc2)
#    return -1 

from heapq import heappush,heappop

def solve(H,W,K,x1,y1,x2,y2,bd) :
    x1 -= 1; y1 -= 1; x2 -= 1; y2 -= 1
    inf = 10**18
    dist = [inf] * H * W * 4
    mh = []
    heappush(mh,(0,encnode(x1,y1,0,W)))
    while mh :
        (d,enc) = heappop(mh)
        if dist[enc] <= d : continue
        dist[enc] = d
        (x,y,dir) = decnode(enc,W)
        if x == x2 and y == y2 : return (d + K - 1) // K
        enc1 = encnode(x,y,(dir+1)%4,W)
        if dist[enc1] == inf : newd = (d + K - 1) // K * K; heappush(mh,(newd,enc1))
        for (xx,yy,nd) in ((x-1,y,0),(x+1,y,1),(x,y-1,2),(x,y+1,3)) :
            if nd != dir or xx < 0 or xx >= H or yy < 0 or yy >=W or bd[xx][yy] == '@': continue
            enc2 = encnode(xx,yy,nd,W)
            if dist[enc2] == inf : heappush(mh,(d+1,enc2))
    return -1 


def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    H,W,K = gis()
    x1,y1,x2,y2 = gis()
    bd = []
    for i in range(H) : bd.append(gs())
    ans = solve(H,W,K,x1,y1,x2,y2,bd)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

