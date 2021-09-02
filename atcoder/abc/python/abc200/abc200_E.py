
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
    N,K = gis()
    dp = [0] * 4
    dp[1] = [0] * (N+1)
    dp[2] = [0] * (2*N+1)
    dp[3] = [0] * (3*N+1)
    cdp = [0] * 4
    cdp[1] = [0] * (N+1)
    cdp[2] = [0] * (2*N+1)
    cdp[3] = [0] * (3*N+1)
    ## Do the base case
    d1 = dp[1]; c1 = cdp[1]
    for i in range(1,N+1) : d1[i] = 1; c1[i] = c1[i-1]+1
    ## Now do the second case
    d2 = dp[2]; c2 = cdp[2]
    for i in range(2,2*N+1) :
        maxv,minv = min(i-1,N),max(i-N,1)
        d2[i] = c1[maxv]-c1[minv-1]
        c2[i] = c2[i-1] + d2[i]
    ## Now for the third case
    d3 = dp[3]; c3 = cdp[3]
    for i in range(3,3*N+1) :
        maxv,minv = min(i-1,2*N),max(i-N,2)
        d3[i] = c2[maxv]-c2[minv-1]
        c3[i] = c3[i-1] + d3[i]

    sum=3; beauty=1; taste=1
    remaining = K
    while True :
        ways = d3[sum]
        if remaining <= ways : break
        sum += 1; remaining -= ways
    while True :
        twosum = sum-beauty
        ways = 0 if twosum < 2 or twosum > 2*N else d2[twosum]
        if remaining <= ways : break
        beauty += 1; remaining -= ways
    while True :
        onesum = sum-beauty-taste
        ways = 0 if onesum < 1 or onesum > N else d1[onesum]
        if remaining <= ways : break
        taste += 1; remaining -= ways
    print(f"{beauty} {taste} {sum-beauty-taste}")

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

