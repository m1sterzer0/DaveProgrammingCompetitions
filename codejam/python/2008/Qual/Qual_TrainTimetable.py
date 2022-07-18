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
        ta,NA,NB = gi(),gi(),gi()
        events = []  ## e[0] is time, e[1] is event type, with additions before deletions
        def conv(s) :
            a = [ord(c) for c in s ]
            z = ord('0')
            return 600*(a[0]-z) + 60*(a[1]-z) + 10*(a[3]-z) + (a[4]-z)
        for i in range(NA): 
            t1,t2 = conv(gs()),conv(gs())
            events.append((t1,2))
            events.append((t2+ta,1))
        for i in range(NB) :
            t1,t2 = conv(gs()),conv(gs())
            events.append((t1,3))
            events.append((t2+ta,0))
        events.sort()
        ansa,ansb,availa,availb = 0,0,0,0
        for e in events :
            if e[1] == 0 : availa += 1
            elif e[1] == 1 : availb += 1
            elif e[1] == 2 :
                if availa == 0 :
                    ansa += 1
                else :
                    availa -= 1
            elif e[1] == 3 :
                if availb == 0 :
                    ansb += 1
                else :
                    availb -= 1
        print(f"Case #{tt}: {ansa} {ansb}")

if __name__ == "__main__" :
    main()

