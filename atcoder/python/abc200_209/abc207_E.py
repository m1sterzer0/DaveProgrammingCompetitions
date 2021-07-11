
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    mm = 10**9+7
    suma = [0] * N; rs = 0
    for i in range(N) : rs += A[i]; suma[i] = rs
    dp1 = [1] * N
    dp2 = [0] * N
    ways = [0] * N
    ans = 1
    for i in range(2,N+1) :
        for j in range(i) : ways[j] = 0
        for j in range(N) :
            idx = suma[j] % i
            dp2[j] = (ways[idx]) % mm
            ways[idx] = (ways[idx] + dp1[j]) % mm
        ans = (ans + dp2[N-1]) % mm
        dp1,dp2 = dp2,dp1
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

