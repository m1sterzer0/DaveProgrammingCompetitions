
import sys
import bisect
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def calcNimber(l,r,L,R,N,cache) :
    #print(f"DBG: Enter Nimber (l,r):({l},{r})")
    if (l,r) not in cache :
        values = set()
        for i in range(N) :
            if L[i] >= l and R[i] <= r :
                v = 0
                if L[i] > l : v ^= calcNimber(l,L[i]-1,L,R,N,cache)
                if R[i] < r : v ^= calcNimber(R[i]+1,r,L,R,N,cache)
                values.add(v)
        for i in range(1000) :
            if i not in values : 
                cache[(l,r)] = i
                break
    #print(f"DBG: Nimber (l,r):({l},{r}) is {cache[(l,r)]}")
    return cache[(l,r)]

def solvecase() :
    N = gi()
    L = [0] * N
    R = [0] * N
    for i in range(N) : L[i],R[i] = gis()
    for i in range(N) : R[i] -= 1
    cache = {}
    x = calcNimber(1,99,L,R,N,cache)
    ans = "Alice" if x > 0 else "Bob"
    print(ans)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for i in range(T) : 
        solvecase()

if __name__ == '__main__' :
    main("junk.in")
    sys.stdout.flush()

