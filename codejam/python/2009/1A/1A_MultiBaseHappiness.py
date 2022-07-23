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

def dosum(v,b) : 
    x = 0
    while (v > 0) : xx = v%b; x += xx*xx; v //= b
    return x

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    maxval1 = 10000
    ## Encoding -2: Not visited, -1: Current Cycle, 0: Bad, 1: Good
    sb = [[-2] * (maxval1+1) for _ in range(11)]
    buffer = []
    for base in range(2,11) :
        sb[base][0] = 0; sb[base][1] = 1
        for i in range(2,maxval1+1) :
            if sb[base][i] >= 0 : continue
            buffer = buffer[:0]; sb[base][i] = -1; buffer.append(i); curs = i
            while True :
                x = dosum(curs,base)
                if sb[base][x] == -2 : sb[base][x] = -1; buffer.append(x); curs = x; continue
                if sb[base][x] == 0 or sb[base][x] == -1 :
                    for xx in buffer : sb[base][xx] = 0
                    break
                if sb[base][x] == 1 :
                    for xx in buffer : sb[base][xx] = 1
                    break

    ansarr = [-1] * (1<<11)
    maxval2 = 100000000 ## Turns out worst case is 11814485
    buf = []
    for bm in range(4,1<<11,4) :
        buf = buf[:0]
        for i in range(2,11) :
            if bm & (1 << i) != 0 : buf.append(i)
        start = 2
        if len(buf) > 1 :
            for i in range(2,11) :
                if i in buf : start = max(start,ansarr[bm ^ (1<<i)])
        for i in range(start,maxval2+1) :
            good = True
            for j in buf :
                ii = i if i <= maxval1 else dosum(i,j)
                if sb[j][ii] == 0 : good = False; break
            if good : ansarr[bm] = i; break
        if ansarr[bm] == -1 :
            print(f"ERROR buf:{buf}")
            
    T = int(infile.readline().rstrip())
    for tt in range(1,T+1) :
        bases = [int(s) for s in infile.readline().rstrip().split()]
        bm = 0
        for b in bases : bm |= (1 << b)
        print(f"Case #{tt}: {ansarr[bm]}")

if __name__ == "__main__" :
    main()

