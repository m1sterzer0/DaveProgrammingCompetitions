
import sys
import random
from multiprocessing import Pool
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

def solve(N,M,H) :
    ## For each peak, we find the index of a taller peak to the left, and a peak as tall or greater to the right
    lf = [0] * M; rt = [0] * M; st = [(1_000_000_001,-1)]
    for (i,h) in enumerate(H) :
        while st[-1][0] <= h : st.pop()
        lf[i] = st[-1][1]
        st.append((h,i))
    st.clear(); st.append((1_000_000_001,M))
    for i in range(M-1,-1,-1) :
        h = H[i]
        while st[-1][0] < h : st.pop()
        rt[i] = st[-1][1]
        st.append((h,i))

    ## We are looking for three quantities
    ## a) Number of paintings == number of clouds --> Determines amount of gray paint and amount of total paint
    ## b) Total size of clounds --> Determines amount of white paint
    ## c) Counting black paint, which is the trickiest part.  Black paint consists of two parts
    ##    -- the black paint above the clouds.  This can be done with simple algebra.
    ##    -- the black paint above the nonclouds.  Here we assume we will count black paint for all of the paintings and then subtract
    ##       out the columns covered by clouds
    ## d) blue paint can then be calculated as total paint - gray paint - white paint - black paint
    ## Here we process the peaks, can consider clouds over that peak for which that peak is the leftmost tallest peak under the cloud

    ## Tot height in a total range N
    ## (N-0) * 1 + (N-1) * 2 + ... * (N-(N-1)) * N = N*1 + (N-1)*2 + ... + 1 * N
    ## Likely in OEIS: P(0) = 0, P(1) = 1, P(2) = 4, P(3) = 10, P(4) = 20, P(5) = 35
    ## N * (N+1) * (N+2) // 6
    ## Tot width with a options for left side and b options for right side
    ## ab + b * (0 + 1 + 2 + ... + (a-1)) + a * (0 + 1 + 2 + ... + (b-1))
    ## ab + b * (a-1)*a // 2 + a * (b-1) * b // 2
    ## Black paint above clouds(per total width)
    ## 0 * (htrange) + 1 * (htrange-1) + ... + (htrange-1) * 1, same as totheight above
    ## = (htrange-1) * (htrange) * (htrange+1) // 6

    twoinv = pow(2,MOD-2,MOD)
    sixinv = pow(6,MOD-2,MOD)
    white = 0
    totclouds = 0
    blackadder = 0
    blackslopes = [0] * M
    for (i,h) in enumerate(H) :
        lidx = lf[i]+1; ridx = rt[i]-1
        htrange = N-h; a = i-lidx+1; b = ridx-i+1
        numvertintervals =   htrange * (htrange+1) // 2 % MOD # 1+2+...+htrange = htrange*(htrange+1)//2
        numhorizintervals = a*b % MOD
        totalht = htrange * (htrange+1) % MOD * (htrange+2) % MOD * sixinv % MOD
        totalwd = (a * b % MOD + b * (a-1) % MOD * a % MOD * twoinv % MOD + a * (b-1) % MOD * b % MOD * twoinv % MOD) % MOD
        numclouds = numvertintervals * numhorizintervals % MOD
        cloudarea = totalht * totalwd % MOD
        blackpaint = (htrange-1) * htrange % MOD * (htrange+1) % MOD * sixinv % MOD * totalwd % MOD
        white += cloudarea
        totclouds += numclouds
        blackadder += blackpaint
        sl1 = (-b * numvertintervals) % MOD
        sl2 = ((a+b) * numvertintervals) % MOD
        sl3 = (-a * numvertintervals) % MOD
        blackslopes[lidx] = (blackslopes[lidx] + sl1) % MOD
        if i+1 < M :    blackslopes[i+1] = (blackslopes[i+1] + sl2) % MOD
        if ridx+2 < M : blackslopes[ridx+2] = (blackslopes[ridx+2] + sl3) % MOD
    white %= MOD; totclouds %= MOD; blackadder %= MOD
    ## Now we do the gray and black
    gray = (sum(H) % MOD) * totclouds % MOD
    black = 0; blackcnt = totclouds; blackslope = 0
    for i in range(M) :
        blackslope = (blackslope + blackslopes[i]) % MOD
        blackcnt = (blackcnt + blackslope) % MOD
        #print(f"DBG: i:{i} blackslope:{blackslope} blackcnt:{blackcnt}")
        black += (N-H[i]) * blackcnt % MOD
    black = (black % MOD + blackadder) % MOD  
    blue = (N * M % MOD * totclouds % MOD - white - gray - black) % MOD

    ## Now we do the blue
    whiterem = 0 if white == 0 else MOD - white
    grayrem  = 0 if gray == 0 else MOD - gray
    blackrem = 0 if black == 0 else MOD - black
    bluerem  = 0 if blue == 0 else MOD - blue
    return f"{blackrem} {whiterem} {grayrem} {bluerem}"

def getInputs(tt) :
    N,M,K = gis()
    L,S1,A,B = [],[],[],[]
    for _ in range(K) :
        l,s1,a,b = gis()
        L.append(l); S1.append(s1); A.append(a); B.append(b)
    return (tt,N,M,K,L,S1,A,B)

def solvemulti(xxx) :
    (tt,N,M,K,L,S1,A,B) = xxx
    H = []
    for (l,s1,a,b) in zip(L,S1,A,B) :
        H.append(s1)
        for _ in range(1,l) :
            H.append((a*H[-1]+b) % (N-1) + 1)
    print(f"Solving case {tt} (N={N} M={M})...",file=sys.stderr)
    return solve(N,M,H)

def main(infn="") :
    random.seed(8675309)
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = False
    if multi :
        inputs = []
        for tt in range(1,T+1) : inputs.append(getInputs(tt))
        with Pool(processes=32) as pool : outputs = pool.map(solvemulti,inputs)
        for tt,ans in enumerate(outputs,1) : print(f"Case #{tt}: {ans}")
    else :
        for tt in range(1,T+1) : 
            inp = getInputs(tt)
            ans = solvemulti(inp)
            print(f"Case #{tt}: {ans}")

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

