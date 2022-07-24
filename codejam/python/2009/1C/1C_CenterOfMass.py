import math
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
        N = gi(); sumx,sumy,sumz,sumvx,sumvy,sumvz = 0,0,0,0,0,0
        for i in range(N) :
            x,y,z,vx,vy,vz = gi(),gi(),gi(),gi(),gi(),gi()
            sumx += x; sumy += y; sumz += z; sumvx += vx; sumvy += vy; sumvz += vz
        if sumvx == 0 and sumvy == 0 and sumvz == 0 :
            tmin = 0.0
        else :
            num   = sumx*sumvx+sumy*sumvy+sumz*sumvz
            denom = -1*(sumvx*sumvx+sumvy*sumvy+sumvz*sumvz)
            tmin = max(0.00,num/denom)
        x = (sumx + tmin * sumvx) / N
        y = (sumy + tmin * sumvy) / N
        z = (sumz + tmin * sumvz) / N
        dmin = math.sqrt(x*x+y*y+z*z)
        print(f"Case #{tt}: {dmin} {tmin}")

if __name__ == "__main__" :
    main()

