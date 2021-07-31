
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def dodp(P,H) :
    ansa = [0] * (H+1)
    ansb = [0] * (H+1)
    ans = [0] * (H+1)
    dp = [0] * (H+1)
    lastdp = [0] * (H+1)
    for start in "AB" :
        ansarr = ansa if start == "A" else ansb
        for p in P :
            for i in range(H+1) : dp[i] = 0
            for (k,c) in enumerate(p) :
                dp,lastdp = lastdp,dp
                if c == start :
                    ## Even terms are a min of their current value and the previous value
                    dp[0] = lastdp[0]
                    for i in range(2,H+1,2) : dp[i] = min(lastdp[i],lastdp[i-1])
                    ## Add one to the odd terms
                    for i in range(1,H+1,2) : dp[i] = min(dp[i-1],lastdp[i]+1) ## min guarantees monotonicity
                else :
                    ## Odd terms are a min of their current value and the previous value
                    for i in range(1,H+1,2) : dp[i] = min(lastdp[i],lastdp[i-1])
                    ## Add one to the even terms
                    dp[0] = lastdp[0]+1
                    for i in range(2,H+1,2) : dp[i] = min(dp[i-1],lastdp[i]+1)  ## min guarantees monotonicity
                #print(f"DBG: _k:{k} p:{p} c:{c} dp:{dp} lastdp:{lastdp}")
            for i in range(H+1) : ansarr[i] += dp[i]
    for i in range(H+1) :
        ans[i] = min(ansa[i],ansb[i])
    #print(f"DBG: ans:{ans}")
    return ans

def solve(H,S,K,P,L) :
    cost = dodp(P,H)
    ansarr = []
    for ll in L :
        l,u = -1,len(cost)-1
        while u-l > 1 :
            m = (l+u)>>1
            (l,u) = (l,m) if cost[m] <= ll else (m,u)
        ansarr.append(u)
    ans = " ".join(str(x+1) for x in ansarr)
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        H,S,K = gis()
        P = [ [] for _ in range(S) ]
        for i in range(H) : 
            x = gs()
            for j in range(S) : P[j].append(x[j])
        L = gis()
        print(f"Case {ntc} H:{H} S:{S} K:{K}",file=sys.stderr)
        ans = solve(H,S,K,P,L)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()
