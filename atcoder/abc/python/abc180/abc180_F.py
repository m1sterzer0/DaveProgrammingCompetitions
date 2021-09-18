import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def modinv(a,p) : return pow(a,p-2,p)

def solveit(N,M,L,chain,ring,comb) :
    mm = 10**9+7
    dp = [0] * (N+1)
    for i in range(N+1) : dp[i] = [0] * (M+1)
    dp[0][0] = 1
    dp[1][0] = 1
    for n in range(2,N+1) :
        for psize in range(1,min(n,L)+1) :
            n2 = n - psize
            if psize == 1 :
                carr = [(0,1)]
            elif psize == 2 :
                carr = [(1,n-1),(2,n-1)]
            else :
                x = 1 if n == psize else comb[n-1][psize-1]
                carr = [(psize-1,x*chain[psize]%mm),(psize,x*ring[psize]%mm)]
            for (m,ncomb) in carr :
                if m > M : continue
                if n2 == 0 :
                    dp[n][m] = (dp[n][m] + ncomb) % mm
                else :
                    for m2 in range(M-m+1) :
                        dp[n][m+m2] = (dp[n][m+m2] + (ncomb * dp[n2][m2] % mm)) % mm
    return dp[N][M]
            
def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M,L = gis()
    mm = 10**9+7
    chain = [0] * 310; ring = [0] * 310
    chain[3] = 3; ring[3] = 1
    for k in range(4,310) : 
        chain[k] = k * chain[k-1] % mm
        ring[k] = (k-1) * ring[k-1] % mm
    imod = [0] * 310
    imod[1] = 1
    for i in range(2,310) : imod[i] = modinv(i,mm)
    comb = [0] * 310
    for i in range(310) : comb[i] = [0] * 310
    comb[0][0] = 1
    for i in range(1,310) :
        comb[i][0] = comb[i][i] = 1
        for j in range(1,i) :
            comb[i][j] = comb[i][j-1] * (i-j+1) % mm * imod[j] % mm
    if L == 1 : 
        ans = solveit(N,M,L,chain,ring,comb)
    else :
        a1 = solveit(N,M,L-1,chain,ring,comb)
        a2 = solveit(N,M,L,chain,ring,comb)
        ans = (a2 + mm - a1) % mm
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

