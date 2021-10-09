import sys
import random
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solveZeroCase(N,dig) :
    if dig <= 0 : return 0
    ans = 0; pv = 10**dig
    for i in range (1,10) :
        if N >= (i+1) * pv :
            ans += 10**dig-9**dig 
        elif N >= i * pv :
            ans += solveZeroCase(N-i*pv,dig-1)  ## Case 1: next digit is also nonzero
            ans += min(10**(dig-1),N-i*pv+1)    ## Case 2: next digit is zero
        else :
            break
    #print(f"DBG: solveZeroCase(N:{N},dig:{dig}) ans:{ans}")
    return ans

def solveNonzeroCase(N,dig,cprod,cache,K) :
    inf = 3*10**18
    if (N,dig,cprod) not in cache :
        if N < (10**(dig+1)-1)//9 :
            cache[(N,dig,cprod)] = 0
        elif dig == 0 :
            ans = 0
            for i in range(1,10) :
                if i <= N and cprod*i <= K : ans += 1
            cache[(N,dig,cprod)] = ans
        else :
            pv = 10**dig
            ans = 0
            for i in range(1,10) :
                if cprod * i > K : break
                if pv * i > N : break
                if pv*(i+1) > N :
                    ans += solveNonzeroCase(N-i*pv,dig-1,cprod*i,cache,K)
                else :
                    ans += solveNonzeroCase(inf,dig-1,cprod*i,cache,K)
            cache[(N,dig,cprod)] = ans
    #print(f"DBG: solveNonzeroCase(N:{N},dig:{dig},cprod:{cprod},K:{K}) result:{cache[(N,dig,cprod)]}")
    return cache[(N,dig,cprod)]

def dumbsolve(N,K) :
    ans = 0
    for x in range(1,N+1) :
        digits = [int(xx) for xx in str(x)]
        if 0 in digits :
            ans += 1
        else :
            curprod = 1
            for d in digits : 
                curprod *= d
                if curprod > K : break
            if curprod <= K : ans += 1
    return ans

def solve(N,K) :
    cache = {}
    ans = 0
    dig = 18
    while 10**dig > N : dig -= 1
    for dd in range(dig+1) :
        ans += solveZeroCase(N,dd)
        ans += solveNonzeroCase(N,dd,1,cache,K)
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,K = gis()
    ans = solve(N,K)
    print(ans)

def check(ntc,Nmin,Nmax,Kmin,Kmax) :
    passed = 0
    for i in range(ntc) :
        N = random.randrange(Nmin,Nmax+1)
        K = random.randrange(Kmin,Kmax+1)
        ans1 = dumbsolve(N,K)
        ans2 = solve(N,K)
        if ans1 != ans2 :
            print(f"ERROR N:{N} K:{K} ans1:{ans1} ans2:{ans2}")
        else :
            passed += 1
    print(f"{passed}/{ntc} passed.")
            
if __name__ == '__main__' :
    main("junk.in")
    #check(10000,1,100,1,500)
    #print(solve(24,999))
    #print(solve(67,1))
    #print(dumbsolve(67,1))
    sys.stdout.flush()

