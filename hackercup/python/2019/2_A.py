
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,M,K,A,B,R,C) :
    ans = 'Y' if K == 2 and (R[0] + C[0]) % 2 == (A+B) % 2 and (R[1] + C[1]) % 2 == (A+B) % 2 else 'N'
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M,K = gis()
        A,B = gis()
        R = [0] * K
        C = [0] * K
        for i in range(K) : R[i],C[i] = gis()
        print(f"Case {ntc} N:{N} M:{M}",file=sys.stderr)
        ans = solve(N,M,K,A,B,R,C)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

