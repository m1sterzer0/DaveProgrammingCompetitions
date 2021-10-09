
import sys
import random
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,K,A) :
    ## First check if we can make all of the entries equal to A[-1]
    A.sort()
    Amax = A[-1]
    if Amax*N - sum(A) <= K : 
        leftover = K - (Amax*N - sum(A))
        return Amax + leftover//N
    ## Now, we know that the soluton is less than A[-1]
    ## MAIN KEY: calculate the number of increments necessary to make x divide all of the Ai efficiently
    ## Precalculate an array of the number of elements in A <= X and the sum of the elements in A <= x
    cumnum = [0] * (Amax+1)
    cumsum = [0] * (Amax+1)
    for a in A : cumnum[a] += 1; cumsum[a] += a
    for i in range(1,Amax+1) : cumnum[i] += cumnum[i-1]; cumsum[i] += cumsum[i-1]
    def cost(x) :
        res = 0
        for k in range(1_000_000_000) :
            lo,hi = k*x+1,(k+1)*x; hiidx = min(A[-1],hi)
            num = cumnum[hiidx] - cumnum[hi-x]
            ss  = cumsum[hiidx] - cumsum[hi-x]
            res += num*hi - ss
            if cumnum[hiidx] == N : return res
    for x in range(A[-1],-1,-1) :
        if cost(x) <= K : return x


#def solveBrute(N,K,A) :
#    amax = max(A)
#    gcdmax = amax + K // N
#    suma = sum(A)
#    while suma + K < N*gcdmax : gcdmax -= 1
#    for g in range(gcdmax,-1,-1) :
#        k = 0
#        for a in A : 
#            r = a % g
#            if r != 0 : k += (g-r)
#        if k <= K : return g
#
#def test(ntc,Nmin,Nmax,Kmin,Kmax,Amin,Amax) :
#    for tt in range(1,ntc+1) :
#        N = random.randrange(Nmin,Nmax+1)
#        K = random.randrange(Kmin,Kmax+1)
#        A = [ random.randrange(Amin,Amax+1) for _ in range(N) ]
#        ans1 = solveBrute(N,K,A)
#        ans2 = solve(N,K,A)
#        if ans1 != ans2 :
#            print(f"ERROR: tt:{tt} N:{N} K:{K} ans1:{ans1} ans2:{ans2}")

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    ## If Amin + K >= Amax, then we can make all of the elements equal to Amin + K
    ## Otherwise, assume Amin+K < Amax
    ## Brute force ways is to try Amin-K, and when we find someone who can't do it, we find the best they can do and start over.
    N,K = gis()
    A = gis()
    ans = solve(N,K,A)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(1234)
    #test(10,2,200,1,10000,1,10000)
    main()
    sys.stdout.flush()

