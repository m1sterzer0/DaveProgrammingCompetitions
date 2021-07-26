
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,K,V,powers) :
    ## Pass 1: Find the intervals that need help
    ans,i,m,s,mm = 0,N,0,0,10**9+7
    for c in V[::-1] :
        s += 1 if c == "B" else -1
        if s - m > K :
            ans = (ans + powers[i]) % mm
            s -= 2
        m = min(m,s)
        i -= 1
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    powers = [1] * 1_000_001
    mm = 10**9+7
    for i in range(1,1_000_001) :
        powers[i] = 2 * powers[i-1] % mm
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,K = gis()
        V = gs()
        print(f"Case {ntc} N:{N}", file=sys.stderr)
        ans = solve(N,K,V,powers)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

