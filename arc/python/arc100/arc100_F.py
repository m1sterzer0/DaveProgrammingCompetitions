
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def fastpow(a,b) :
    res = 1
    while b :
        if b & 1 :
            res = res * a % 1_000_000_007
        a = a * a % 1_000_000_007
        b >>= 1
    return res

def iscolorful(A,K) :
    sb = [-1] * (K+1)
    streak = 0
    for (i,a) in enumerate(A) :
        if sb[a]  == -1 :
            streak += 1
        else :
            streak = min(streak+1,i-sb[a])
        sb[a] = i
        if streak == K : return True
    return False

def fact(N) :
    ans = 1
    for i in range(1,N+1) : ans = ans * i % 1_000_000_007
    return ans

def isunique(A,K) :
    prefixlen = 0
    sb = [-1] * (K+1)
    for a in A :
        if sb[a] >= 0 : break
        prefixlen += 1
        sb[a] = 1
    suffixlen = 0
    sb = [-1] * (K+1)
    for a in reversed(A) :
        if sb[a] >= 0 : break
        suffixlen += 1
        sb[a] = 1
    return (prefixlen,suffixlen)

def case1(N,M,K) :
    dp = [0] * K; olddp = [0] * K
    dp2 = [0] * K; olddp2 = [0] * K
    for i in range(N) :
        (dp,olddp,dp2,olddp2) = (olddp,dp,olddp2,dp2)
        dp[0] = 0; dp2[0] = 0; olddp[0] = 1 if i == 0 else 0; olddp2[0] = 0
        for j in range(1,K) :
            dp[j] = olddp[j]
            dp2[j] = olddp2[j]
        for j in range(K-2,0,-1) :
            dp[j] =  (dp[j] + dp[j+1]) % 1_000_000_007
            dp2[j] = (dp2[j] + dp2[j+1]) % 1_000_000_007
        for j in range(K-1) :
            dp[j+1] = (dp[j+1] + olddp[j] * (K-j) % 1_000_000_007) % 1_000_000_007
            dp2[j+1] = (dp2[j+1] + olddp2[j] * (K-j) % 1_000_000_007) % 1_000_000_007
        for j in range(M,K) :
            dp2[j] = (dp2[j] + dp[j]) % 1_000_000_007
        #print(f"i:{i} dp:{dp} dp2:{dp2}")
    return sum(dp2) * fact(K-M) * fastpow(fact(K),1_000_000_005)

def docase2(streak,N,K) :
    dp = [0] * K; olddp = [0] * K
    dp[streak] = 1
    ans = [0] * (N+1)
    ans[0] = 1
    for i in range(N) :
        (dp,olddp) = (olddp,dp)
        for j in range(1,K) :      dp[j]   =  olddp[j]
        for j in range(K-2,0,-1) : dp[j]   =  (dp[j] + dp[j+1]) % 1_000_000_007
        for j in range(K-1) :      dp[j+1] = (dp[j+1] + olddp[j] * (K-j) % 1_000_000_007) % 1_000_000_007
        ans[i+1] = sum(dp)
    return ans

def case2(prefixlen,suffixlen,N,M,K) :
    a = docase2(prefixlen,N,K)
    b = docase2(suffixlen,N,K)
    ans = 0
    for i in range(N+1) :
        j = N - M - i
        if j < 0 : break
        ways = a[i] * b[j] % 1_000_000_007
        ans = (ans + ways) % 1_000_000_007
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,K,M = gis()
    A = gis()
    Atotal = fastpow(K,N-M) * (N-M+1) % 1_000_000_007
    ans = 0
    if iscolorful(A,K) :
        ans = Atotal
    else:
        plen,slen = isunique(A,K)
        if plen == M :
            sub = case1(N,M,K)
            ans = (Atotal + 1_000_000_007 - sub) % 1_000_000_007
        else :
            sub = case2(plen,slen,N,M,K)
            ans = (Atotal + 1_000_000_007 - sub) % 1_000_000_007
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

