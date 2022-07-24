import sys
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

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        W,Q = gi(),gi(); bd = [gs() for _ in range(W)]; QQ = [gi() for _ in range(Q)]
        midx = namedtuple("midx","i j v")
        mval = namedtuple("mval","l s")
        ansarr = [ mval(0,"") for _ in range(251) ]
        numleft = 0
        for q in QQ :
            if ansarr[q].l == 0 :
                numleft += 1; ansarr[q] = mval(-1,"")
        que = deque(); rnd = 1; lkup = {}
        for i in range(W) :
            for j in range(W) :
                if bd[i][j] in "-+" : continue
                v = int(bd[i][j])
                if ansarr[v].l == -1 : numleft -= 1; ansarr[v] = mval(1,bd[i][j])
                lkup[midx(i,j,v)] = mval(1,bd[i][j])
                que.append(midx(i,j,v))
        deltas = [(0,-1),(0,1),(-1,0),(1,0)]
        while que :
            m = que.popleft()
            mv = lkup[m]
            if rnd == mv.l :
                if numleft == 0 : break
                rnd += 1
            for (d1,d2) in deltas :
                i1,j1 = m.i+d1,m.j+d2
                if i1 < 0 or i1 >= W or j1 < 0 or j1 >= W : continue
                sgn = 1 if bd[i1][j1] == '+' else -1
                for (d3,d4) in deltas :
                    i2,j2 = i1+d3,j1+d4
                    if i2 < 0 or i2 >= W or j2 < 0 or j2 >= W : continue
                    v = m.v+sgn*int(bd[i2][j2])
                    key = midx(i2,j2,v)
                    val = mval(1000000,"")
                    if key in lkup : val = lkup[key]
                    if val.l < rnd : continue
                    val2 = mval(rnd,mv.s + (bd[i1][j1] + bd[i2][j2]))
                    if val.l == val2.l and val.s <= val2.s : continue
                    if key not in lkup : que.append(key)
                    lkup[key] = val2
                    if key.v >= 0 and key.v <= 250 and (ansarr[key.v].l == -1 or ansarr[key.v].l == rnd and val2.s < ansarr[key.v].s) :
                        if (ansarr[key.v].l == -1) : numleft -= 1
                        ansarr[key.v] = val2
        print(f"Case #{tt}:")
        for q in QQ :
            print(ansarr[q].s)

if __name__ == "__main__" :
    main()

