
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## Convolution code leveraged from other transcriptions of atcoder library
MOD = 998244353
IMAG = 911660635
IIMAG = 86583718
rate2 = (0, 911660635, 509520358, 369330050, 332049552, 983190778, 123842337, 238493703, 975955924, 603855026, 856644456, 131300601, 842657263, 730768835, 942482514, 806263778, 151565301, 510815449, 503497456, 743006876, 741047443, 56250497, 867605899, 0)
irate2 = (0, 86583718, 372528824, 373294451, 645684063, 112220581, 692852209, 155456985, 797128860, 90816748, 860285882, 927414960, 354738543, 109331171, 293255632, 535113200, 308540755, 121186627, 608385704, 438932459, 359477183, 824071951, 103369235, 0)
rate3 = (0, 372528824, 337190230, 454590761, 816400692, 578227951, 180142363, 83780245, 6597683, 70046822, 623238099, 183021267, 402682409, 631680428, 344509872, 689220186, 365017329, 774342554, 729444058, 102986190, 128751033, 395565204, 0)
irate3 = (0, 509520358, 929031873, 170256584, 839780419, 282974284, 395914482, 444904435, 72135471, 638914820, 66769500, 771127074, 985925487, 262319669, 262341272, 625870173, 768022760, 859816005, 914661783, 430819711, 272774365, 530924681, 0)
 
def _butterfly(a):
    n = len(a)
    h = (n - 1).bit_length()
    le = 0
    while le < h:
        if h - le == 1:
            p = 1 << (h - le - 1)
            rot = 1
            for s in range(1 << le):
                offset = s << (h - le)
                for i in range(p):
                    l = a[i + offset]
                    r = a[i + offset + p] * rot
                    a[i + offset] = (l + r) % MOD
                    a[i + offset + p] = (l - r) % MOD
                rot *= rate2[(~s & -~s).bit_length()]
                rot %= MOD
            le += 1
        else:
            p = 1 << (h - le - 2)
            rot = 1
            for s in range(1 << le):
                rot2 = rot * rot % MOD
                rot3 = rot2 * rot % MOD
                offset = s << (h - le)
                for i in range(p):
                    a0 = a[i + offset]
                    a1 = a[i + offset + p] * rot
                    a2 = a[i + offset + p * 2] * rot2
                    a3 = a[i + offset + p * 3] * rot3
                    a1na3imag = (a1 - a3) % MOD * IMAG
                    a[i + offset] = (a0 + a2 + a1 + a3) % MOD
                    a[i + offset + p] = (a0 + a2 - a1 - a3) % MOD
                    a[i + offset + p * 2] = (a0 - a2 + a1na3imag) % MOD
                    a[i + offset + p * 3] = (a0 - a2 - a1na3imag) % MOD
                rot *= rate3[(~s & -~s).bit_length()]
                rot %= MOD
            le += 2
 
def _butterflyinv(a):
    n = len(a)
    h = (n - 1).bit_length()
    le = h
    while le:
        if le == 1:
            p = 1 << (h - le)
            irot = 1
            for s in range(1 << (le - 1)):
                offset = s << (h - le + 1)
                for i in range(p):
                    l = a[i + offset]
                    r = a[i + offset + p]
                    a[i + offset] = (l + r) % MOD
                    a[i + offset + p] = (l - r) * irot % MOD
                irot *= irate2[(~s & -~s).bit_length()]
                irot %= MOD
            le -= 1
        else:
            p = 1 << (h - le)
            irot = 1
            for s in range(1 << (le - 2)):
                irot2 = irot * irot % MOD
                irot3 = irot2 * irot % MOD
                offset = s << (h - le + 2)
                for i in range(p):
                    a0 = a[i + offset]
                    a1 = a[i + offset + p]
                    a2 = a[i + offset + p * 2]
                    a3 = a[i + offset + p * 3]
                    a2na3iimag = (a2 - a3) * IIMAG % MOD
                    a[i + offset] = (a0 + a1 + a2 + a3) % MOD
                    a[i + offset + p] = (a0 - a1 + a2na3iimag) * irot % MOD
                    a[i + offset + p * 2] = (a0 + a1 - a2 - a3) * irot2 % MOD
                    a[i + offset + p * 3] = (a0 - a1 - a2na3iimag) * irot3 % MOD
                irot *= irate3[(~s & -~s).bit_length()]
                irot %= MOD
            le -= 2

def convolvefftmod(a,b) :
    finalsz = len(a)+len(b)-1
    z = 1
    while z < finalsz : z *= 2
    la = a.copy()
    for _ in range(z-len(a)) : la.append(0)
    lb = b.copy()
    for _ in range(z-len(b)) : lb.append(0)
    _butterfly(la)
    _butterfly(lb)
    for i in range(z) : la[i] *= lb[i]; la[i] %= MOD
    _butterflyinv(la)
    iz = pow(z,MOD-2,MOD)
    for i in range(z) : la[i] *= iz; la[i] %= MOD
    return la[:finalsz]

def domiddle(N,M,A,B,P,dp,l,m,r) :
    for i in range(M) :
        a,b = A[i],B[i]
        v2 = P[i][0:(r-l+1)]
        for n1,n2 in ((a,b),(b,a)) :
            v1 = dp[n1][l:m+1]
            xx = convolvefftmod(v1,v2)
            for i in range(m+1,r+1) :
                dp[n2][i] += xx[i-l]
                dp[n2][i] %= MOD 

def divide(N,M,A,B,P,dp,l,r) :
    if l == r : return
    m = (r+l)>>1
    divide(N,M,A,B,P,dp,l,m)
    domiddle(N,M,A,B,P,dp,l,m,r)
    divide(N,M,A,B,P,dp,m+1,r)

def solveBrute(N,M,T,A,B,P) :
    dp = [[0]*(T+1) for _ in range(N)]
    dp[0][0] = 1
    for t in range(0,T+1) :
        for i in range(M) :
            a,b = A[i],B[i]
            for tt in range(1,T+1) :
                w = P[i][tt]
                if t+tt > T : continue
                dp[b][t+tt] = (dp[b][t+tt] + w * dp[a][t]) % MOD
                dp[a][t+tt] = (dp[a][t+tt] + w * dp[b][t]) % MOD
    return dp[0][T]

def solve(N,M,T,A,B,P) :
    dp = [[0]*(T+1) for _ in range(N)]
    dp[0][0] = 1
    divide(N,M,A,B,P,dp,0,T)
    return dp[0][T]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M,T = gis()
    A = []; B = []; P = []
    for _ in range(M) :
        a,b = gis(); p = gis(); A.append(a-1); B.append(b-1); P.append([0]+p)
    ans = solve(N,M,T,A,B,P)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

