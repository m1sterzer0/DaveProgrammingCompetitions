
import sys
import heapq
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,H,V,X,Y) :
    if N > H+V : return -1
    pts = [x << 30 | y for (x,y) in zip(X,Y)]
    pts.sort()
    XX = [x >> 30 for x in pts]
    YY = [x & 0x3fffffff for x in pts]
    best = 10**18
    ## For the Y, calculate the max from the right end
    maxy = [0] * N; m = 0
    for i in range(N-1,-1,-1) :
        m = max(m,YY[i]); maxy[i] = m
    mh = []
    if V >= N : best = maxy[0]

    vcost1 = 0
    for i in range(N) :
        hcost = XX[i]
        heapq.heappush(mh,YY[i])
        if len(mh) > H : vcost1 = max(vcost1,heapq.heappop(mh))
        vcost2 = 0 if i == N-1 else maxy[i+1]
        vcost = max(vcost1,vcost2)
        if N-i-1 <= V : best = min(best,hcost+vcost)
    return best

def getinput(N) :
    X1,X2,A,B,C,D = gis()
    X = [X1,X2]
    for i in range(2,N) : 
        x = (A * X[-2] + B*X[-1] + C) % D + 1
        X.append(x)
    return X 

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,H,V = gis()
        X = getinput(N)
        Y = getinput(N)
        print(f"Case {ntc} N:{N} H:{H} V:{V}",file=sys.stderr)
        ans = solve(N,H,V,X,Y)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

