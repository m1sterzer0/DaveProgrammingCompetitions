
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,W) :
    W.sort()
    q = collections.deque(W)
    ans = 0
    while q :
        heavy = q.pop()
        items = 1
        while q and heavy * items < 50 : q.popleft(); items += 1
        if heavy * items >= 50 : ans += 1
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N = gi()
        W = [gi() for _ in range(N)]
        print(f"Case {ntc} N:{N}",file=sys.stderr)
        ans = solve(N,W)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

