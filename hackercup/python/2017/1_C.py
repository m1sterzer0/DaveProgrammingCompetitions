
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,M,K,A,B,G,SS,DD) :
    myinf = 100_000_000
    gr =  [[myinf] * N for _ in range(N)]
    for i in range(N) : gr[i][i] = 0
    for (ap1,bp1,g) in zip(A,B,G) :
        a = ap1-1; b = bp1-1; gr[a][b] = min(gr[a][b],g); gr[b][a] = min(gr[b][a],g)

    for k in range(N) :
        for i in range(N) : 
            for j in range(N) :
                gr[i][j] = min(gr[i][j],gr[i][k]+gr[k][j])
    #for i in range(N) :  print(f"DBG: gr[{i}] = {gr[i]}")
    ## dp #number completed, ## number in truck
    ## max possible is no larger than (# cities visited) * (max dist between cities) <= (2*N) * (N-1) * 1000 <= 200*100*1000 = 20_000_000 
    S = [s-1 for s in SS]
    D = [d-1 for d in DD]
    dp0 = [myinf] * (K+1)
    dp1a = [myinf]* (K+1)
    dp1b = [myinf] * (K+1)
    dp2 = [myinf] * (K+1)
    dp0[0] = 0
    dp1a[0] = dp0[0] + gr[0][S[0]]
    if K > 1 : dp2[0] = dp1a[0] + gr[S[0]][S[1]]
    for i in range(1,K+1) :
        dp0[i] = min(dp1a[i-1] + gr[S[i-1]][D[i-1]], myinf if i == 1 else dp1b[i-1] + gr[D[i-2]][D[i-1]])
        if i != K :
            dp1a[i] = dp0[i] + gr[D[i-1]][S[i]]
            dp1b[i] = dp2[i-1] + gr[S[i]][D[i-1]]
            if i != K-1 :
                dp2[i]  = min(dp1a[i] + gr[S[i]][S[i+1]], dp1b[i] + gr[D[i-1]][S[i+1]])
        #print(f"DBG i:{i} dp0:{dp0[i]} dp1a:{dp1a[i]} dp1b:{dp1b[i]} dp2:{dp2[i]}")
    return -1 if dp0[K] >= myinf else dp0[K]
    
def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M,K = gis()
        A,B,G = [],[],[]
        for _ in range(M) : a,b,g = gis(); A.append(a); B.append(b); G.append(g)
        S,D = [],[]
        for _ in range(K) : s,d = gis(); S.append(s); D.append(d)
        print(f"Case {ntc} N:{N} M:{M} K:{K}",file=sys.stderr)
        ans = solve(N,M,K,A,B,G,S,D)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

