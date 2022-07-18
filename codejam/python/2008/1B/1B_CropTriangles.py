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
        n,A,B,C,D,x0,y0,M = gi(),gi(),gi(),gi(),gi(),gi(),gi(),gi()
        sb = [0] * 9
        x,y = x0,y0
        for _ in range(n) :
            sb[3 * (x%3) + (y%3)] += 1
            x = (A*x+B) % M
            y = (C*y+D) % M
        ans = 0
        for i in range(9) :
            for j in range(i,9) :
                for k in range(j,9) :
                    if ((i//3)+(j//3)+(k//3)) % 3 != 0 : continue
                    if ((i%3)+(j%3)+(k%3)) % 3 != 0 : continue
                    if i == j and j == k :
                        ans += sb[i] * (sb[i]-1) * (sb[i]-2) // 6
                    elif i == j :
                        ans += sb[i] * (sb[i]-1) // 2 * sb[k]
                    elif j == k :
                        ans += sb[i] * sb[j] * (sb[j]-1) // 2
                    else :
                        ans += sb[i] * sb[j] * sb[k]
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

