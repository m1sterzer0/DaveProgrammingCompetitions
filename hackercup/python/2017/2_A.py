
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,M,K) :
    myinf = 10**18
    ans = myinf
    for (x,y) in ((N,M),(M,N)) :
        ## Method 1, block off most of a row, and then 1 block on the way back to cause a turnaround
        if K < x and 2*K+3 <= y: ans = min(ans, (x-1)//K + 1)
        ## Method 2, one block under start, wall to right, and one block to cause backtrack
        ## Wall must be 3 blocks tall if K == 1, otherwise 2 blocks is sufficient
        if 2*K+1 <= x and 2*K+3 <= y : ans = min(ans, (5 if K == 1 else 4))
    return -1 if ans == myinf else ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M,K = gis()
        print(f"Case {ntc} N:{N} M:{M} K:{K}",file=sys.stderr)
        ans = solve(N,M,K)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

