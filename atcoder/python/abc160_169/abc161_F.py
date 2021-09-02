
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## If K divides N, then we repeatedly divide by K and then check if the result equals 1 mod N
## If K doesn't divide N, then we want K iff N % K == 1 <==> (N-1) % K == 0 <==> K divides (N-1)
## Thus, we first collect divisors of (N-1) and N, and then run through the algorithm.
def getDivisors(N) :
    ans = []
    i = 1
    while i*i <= N :
        if N % i == 0 :
            ans.append(i)
            if i*i < N : ans.append(N//i)
        i += 1
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    a = getDivisors(N-1)
    b = getDivisors(N)
    c = list(set(a + b))
    ans = 0
    for k in c :
        if k == 1 : continue
        n = N
        while n % k == 0 : n //= k
        if n % k == 1 : ans += 1
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

