
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

def solve(N,M,R) :
    ## Special case 1
    ## Loop through each pair of endpoints
    ## Calculate spare slots
    ## Add (N-2)! * comb((spare slots + N),N)  (Stars and bars baby!)
    if N == 1 : return M
    sparecnt = []
    twosum = 2 * sum(R)
    for i in range(N) :
        for j in range(i+1,N) :
            sp = (M-1) - (twosum - R[i] - R[j])
            sparecnt.append(sp)
    m1,m2 = min(sparecnt),max(sparecnt)
    if m2 < 0 : return 0
    m1 = max(m1,0)
    combnumarr = [0] * (m2-m1+1)
    for i in range(m2-m1+1) :
        if i == 0 :
            c = 1
            for i in range(1,N+1) : c = c * (m1+i) % MOD
            combnumarr[0] = c
        else :
            combnumarr[i] = combnumarr[i-1] * (m1+i+N) % MOD * pow(m1+i,MOD-2,MOD) % MOD
    factN = 1
    for i in range(1,N+1) : factN = factN * i % MOD
    combnumdenom = pow(factN,MOD-2,MOD)
    factNm2 = 1
    for i in range(1,N-2+1) : factNm2 = factNm2 * i % MOD
    ans = 0
    for s in sparecnt :
        if s < 0 : continue
        ans += combnumarr[s-m1]
    ans = ans * 2 % MOD * combnumdenom % MOD * factNm2 % MOD
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M = gis()
        R = []
        for _ in range(N) : R.append(gi())
        print(f"Case {ntc} N:{N} M:{M}",file=sys.stderr)
        ans = solve(N,M,R)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

