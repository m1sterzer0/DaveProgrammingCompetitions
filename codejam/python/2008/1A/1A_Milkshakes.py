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
##MOD = 1000000007

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        N,M = gi(),gi()
        malted = [-1] * M
        uncnt = [0] * M
        gr = [ [] for _ in range(N) ]
        sb = [0] * N
        mq = deque()
        for i in range(M) :
            t = gi()
            for j in range(t) :
                x,y = gi(),gi()
                if y == 0 :
                    uncnt[i] += 1; gr[x-1].append(i)
                else :
                    malted[i] = x-1
            if uncnt[i] == 0 : mq.append(malted[i])
        good = True
        while mq :
            x = mq.popleft()
            if sb[x] == 1 : continue 
            sb[x] = 1
            for i in gr[x] :
                uncnt[i] -= 1
                if uncnt[i] == 0 :
                    if malted[i] == -1 : good = False; break
                    mq.append(malted[i])
            if not good : break
        if not good :
            print(f"Case #{tt}: IMPOSSIBLE")
        else :
            ans = " ".join([str(x) for x in sb])
            print(f"Case #{tt}: {ans}")


if __name__ == "__main__" :
    main()

