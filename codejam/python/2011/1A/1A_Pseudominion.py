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
Card = namedtuple("Card","idx c s t")
MyState = namedtuple("MyState","t nc tidx c1idx c2idx")
def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        T,C0,C1,C2 = [],[],[],[]
        def placecard(c) :
            if c.t > 0 : T.append(c); return
            if c.c == 0 : C0.append(c); return
            if c.c == 1 : C1.append(c); return
            C2.append(c); return
        N = gi()
        for i in range(N) : c,s,t = gi(),gi(),gi(); cc = Card(i,c,s,t); placecard(cc)
        M = gi()
        for i in range(M) : c,s,t = gi(),gi(),gi(); cc = Card(N+i,c,s,t); placecard(cc)
        ## Presolve the C0 sums
        c0sum = [[0]*81 for _ in range(81)]
        for j in range(81) :
            ca = [x for x in C0 if x.idx < j]
            ca.sort(key=lambda x: x.s)
            ca.reverse()
            s = 0
            for i,c in enumerate(ca) : s += c.s; c0sum[i+1][j] = s
            for i in range(len(ca)+1,81) : c0sum[i][j] = c0sum[len(ca)][j]
        ## Now for the poor man's DFS
        cache = {}
        startst = MyState(1,N,-1,-1,-1)
        stack = [startst]
        while stack :
            st = stack.pop()
            if st.t == 0 : cache[st] = 0; continue
            if st in cache : continue
            ## Check for the states that we need presolved to resolve this state
            if st.tidx+1 < len(T) and T[st.tidx+1].idx < st.nc :
                tidx = st.tidx+1; cc = T[tidx]
                ns = MyState(st.t-1+cc.t,min(N+M,st.nc+cc.c),tidx,st.c1idx,st.c2idx)
                if ns not in cache :
                    stack.append(st)
                    stack.append(ns)
                    continue
                else :
                    cache[st] = cc.s + cache[ns]
            else :
                score = c0sum[min(80,st.t)][st.nc]
                if st.c2idx+1 < len(C2) and C2[st.c2idx+1].idx < st.nc :
                    c2idx = st.c2idx+1; cc = C2[c2idx]
                    ns1 = MyState(st.t,st.nc,st.tidx,st.c1idx,c2idx)
                    ns2 = MyState(st.t-1+cc.t,min(N+M,st.nc+cc.c),st.tidx,st.c1idx,c2idx)
                    if ns1 not in cache or ns2 not in cache :
                        stack.append(st)
                        if ns1 not in cache : stack.append(ns1)
                        if ns2 not in cache : stack.append(ns2)
                        continue
                    else :
                        score = max(score,cache[ns1],cc.s+cache[ns2])
                elif st.c1idx+1 < len(C1) and C1[st.c1idx+1].idx < st.nc :
                    c1idx = st.c1idx+1; cc = C1[c1idx]
                    ns1 = MyState(st.t,st.nc,st.tidx,c1idx,st.c2idx)
                    ns2 = MyState(st.t-1+cc.t,min(N+M,st.nc+cc.c),st.tidx,c1idx,st.c2idx)
                    if ns1 not in cache or ns2 not in cache :
                        stack.append(st)
                        if ns1 not in cache : stack.append(ns1)
                        if ns2 not in cache : stack.append(ns2)
                        continue
                    else :
                        score = max(score,cache[ns1],cc.s+cache[ns2])
                ##print(f"DBG: st:{st} score:{score}")
                cache[st] = score
        ans = cache[startst]
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

