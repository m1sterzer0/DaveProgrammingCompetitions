
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,K) :
    if K <= 3 or N == 2 : return f"0\n1\n1 {N} 1" ## No help
    n = min(K,N)
    edges = []
    wt = K
    edges.append(f"1 {N} {wt}"); wt -= 1
    for i in range(1,n-1) : edges.append(f"{i} {i+1} {wt}"); wt -= 1
    edges.append(f"{n-1} {N} {wt}"); wt -= 1
    gap = (K)*(K-1)//2 - (wt)*(wt+1)//2 - K
    edgestr = "\n".join(edges)
    return f"{gap}\n{len(edges)}\n{edgestr}"

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,K = gis()
        print(f"Case {ntc} N:{N} K:{K}",file=sys.stderr)
        ans = solve(N,K)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

