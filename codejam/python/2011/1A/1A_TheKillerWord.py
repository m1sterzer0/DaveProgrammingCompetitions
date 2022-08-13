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

MyState = namedtuple("MyState","clist lidx score")
def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        N,M = gi(),gi(); D = [gs() for _ in range(N) ]; L = [gs() for _ in range(M)]
        bylen = [[] for _ in range(11)]
        pat = [[0]*26 for _ in range(N)]
        for i,d in enumerate(D) :
            bylen[len(d)].append(i)
            for j,c in enumerate(d) :
                pat[i][ord(c)-ord('a')] |= 1 << j
        ansarr = []
        for lstr in L :
            larr = [ord(c)-ord('a') for c in lstr]
            bestidx,bestscore = 0,-1
            st = [MyState(a,0,0) for a in bylen if a]
            while st :
                lst = st.pop()
                if len(lst.clist) == 1 :
                    if lst.score > bestscore or lst.score == bestscore and bestidx > lst.clist[0] :
                        bestidx,bestscore = lst.clist[0],lst.score
                else :
                    d = {}; l = larr[lst.lidx]
                    for cc in lst.clist :
                        p = pat[cc][l]
                        if p in d : d[p].append(cc) 
                        else : d[p] = [cc]
                    for k,v in d.items() :
                        st.append(MyState(v,lst.lidx+1,lst.score + (1 if k==0 and len(d) > 1 else 0)))
            ansarr.append(D[bestidx])
        ans = " ".join(ansarr)
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

