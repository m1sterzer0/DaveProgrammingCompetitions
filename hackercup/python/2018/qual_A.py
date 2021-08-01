
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,K,V,A) :
    sb = [False] * N
    idx = (V-1)*K
    for i in range(K) : sb[idx % N] = True; idx += 1
    ans = " ".join([A[i] for i in range(N) if sb[i]])
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,K,V = gis()
        A = []
        for i in range(N) : x = gs(); A.append(x)
        print(f"Case {ntc} N:{N}",file=sys.stderr)
        ans = solve(N,K,V,A)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

