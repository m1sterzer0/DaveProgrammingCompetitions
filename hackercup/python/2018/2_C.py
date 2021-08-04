
import sys
import random
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

def solvebrute(N,S,E,X,Y) :
    ans = 0
    for i in range(4**N) :
        dir = [ (i >> (2*j)) & 0x3 for j in range(N)]
        ## 0 is N, 1 is E, 2 is S, 3 is W
        blocked = False
        for j in range(N) :
            for k in range(N) :
                if Y[j] < Y[k] < S and X[k] > X[j] and dir[j] == 0 and dir[k] == 3 : blocked = True; break
                if Y[j] < Y[k] < E and X[k] < X[j] and dir[j] == 0 and dir[k] == 1 : blocked = True; break
                if Y[j] > Y[k] > S and X[k] > X[j] and dir[j] == 2 and dir[k] == 3 : blocked = True; break
                if Y[j] > Y[k] > E and X[k] < X[j] and dir[j] == 2 and dir[k] == 1 : blocked = True; break
            if blocked : break
        if blocked : ans += 1
    return ans


def solve(N,S,E,X,Y) :
    ## Mirror if needed to get S < E
    if S > E : 
        S,E = E,S
        X = [1_000_000_000 - x for x in X]

    ll = [(Y[i],i) for i in range(N)]
    ll.sort()
    sortedlasers = [x[1] for x in ll]
    minx = min(X)

    ## State vector consists of 4 8 bit indices
    ## -- Index of leftmost  laser pointing up below start
    ## -- Index of rightmost laser pointing up below end
    ## -- Index of rightmost laser pointing left above start
    ## -- Index of leftmost  laser pointing right above end
    ## Only need two at any one time, making this O(N^3) w/ O(N^2) storage

    np1 = (N+1)
    dp = [0] * (np1)**2
    nxtdp = [0] * (np1)**2
    dp[np1*N+N] = 1
    ## Phase 1 -- Below start
    for idx in sortedlasers :
        x,y = X[idx],Y[idx]
        if y > S : break
        for i in range(np1*np1) : nxtdp[i] = 0
        for lu in range(np1) :
            for ru in range(np1) :
                st = np1*lu+ru
                if dp[st] == 0 : continue
                nxtdp[st] = (nxtdp[st] + dp[st]) % MOD                           ## Down
                if lu == N or X[lu] > x : nxtdp[st] = (nxtdp[st] + dp[st]) % MOD ## Left
                if ru == N or X[ru] < x : nxtdp[st] = (nxtdp[st] + dp[st]) % MOD ## Right
                ## up
                newlu = idx if lu == N or X[lu] > x else lu
                newru = idx if ru == N or X[ru] < x else ru
                newst = np1*newlu+newru
                nxtdp[newst] = (nxtdp[newst]+dp[st]) % MOD
        dp,nxtdp = nxtdp,dp
        #print(f"DBG idx:{idx} x:{x-minx+1} y:{y} S:{S} E:{E} sumdp:{sum(dp)}")

    ## Cleanup 1
    for i in range(np1*np1) : nxtdp[i] = 0
    for lu in range(np1) :
        for ru in range(np1) :
            st = np1*lu+ru
            newst = np1*N+ru
            nxtdp[newst] = (nxtdp[newst] + dp[st]) % MOD
    dp,nxtdp = nxtdp,dp

    ## Phase 2 -- between start and end
    for idx in sortedlasers :
        x,y = X[idx],Y[idx]
        if y < S : continue
        if y > E : break
        for i in range(np1*np1) : nxtdp[i] = 0
        for rl in range(np1) :
            for ru in range(np1) :
                st = np1*rl+ru
                if dp[st] == 0 : continue
                if rl == N or X[rl] < x : nxtdp[st] = (nxtdp[st] + dp[st]) % MOD ## down
                if ru == N or X[ru] < x : nxtdp[st] = (nxtdp[st] + dp[st]) % MOD ## right
                newrl = idx if rl == N or X[rl] < x else rl
                newru = idx if ru == N or X[ru] < x else ru
                newst1 = np1*rl+newru
                newst2 = np1*newrl+ru
                nxtdp[newst1] = (nxtdp[newst1]+dp[st]) % MOD  ## up
                nxtdp[newst2] = (nxtdp[newst2]+dp[st]) % MOD  ## left
        dp,nxtdp = nxtdp,dp
        #print(f"DBG idx:{idx} x:{x-minx+1} y:{y} S:{S} E:{E} sumdp:{sum(dp)}")

    ## Cleanup 2
    for i in range(np1*np1) : nxtdp[i] = 0
    for rl in range(np1) :
        for ru in range(np1) :
            st = np1*rl+ru
            newst = np1*rl+N
            nxtdp[newst] = (nxtdp[newst] + dp[st]) % MOD
    dp,nxtdp = nxtdp,dp
    
    ## Phase 3 -- above end
    for idx in sortedlasers :
        x,y = X[idx],Y[idx]
        if y < E : continue
        for i in range(np1*np1) : nxtdp[i] = 0
        for rl in range(np1) :
            for lr in range(np1) :
                st = np1*rl+lr
                if dp[st] == 0 : continue
                nxtdp[st] = (nxtdp[st] + dp[st]) % MOD                                                        ## up is safe
                if (rl == N or X[rl] < x) and (lr == N or X[lr] > x) : nxtdp[st] = (nxtdp[st] + dp[st]) % MOD ## down
                newrl = idx if rl == N or X[rl] < x else rl
                newlr = idx if lr == N or X[lr] > x else lr
                newst1 = np1*rl+newlr
                newst2 = np1*newrl+lr
                nxtdp[newst1] = (nxtdp[newst1]+dp[st]) % MOD  ## Right
                nxtdp[newst2] = (nxtdp[newst2]+dp[st]) % MOD  ## Left
        dp,nxtdp = nxtdp,dp
        #print(f"DBG idx:{idx} x:{x-minx+1} y:{y} S:{S} E:{E} sumdp:{sum(dp)}")

    ans = (pow(4,N,MOD) - sum(dp)) % MOD
    return ans

def test(ntc,Nmin,Nmax,check=True) :
    numpass = 0
    for tt in range(1,ntc+1) :
        N = random.randrange(Nmin,Nmax+1)
        Y = [x for x in range(1,N+3)]; random.shuffle(Y)
        X = [x for x in range(1,N+1)]; random.shuffle(X)
        S = Y.pop(); E = Y.pop()
        #print(f"DBG: N:{N} S:{S} E:{E} X:{X} Y:{Y}")
        ans = solve(N,S,E,X,Y)
        if not check :
            print(f"tt:{tt} N:{N} ans:{ans}")
        else :
            ans2 = solvebrute(N,S,E,X,Y)
            if ans == ans2 : numpass += 1
            else :
                print(f"ERROR: tt:{tt} N:{N} S:{S} E:{E} X:{X} Y:{Y} ans:{ans} ans2:{ans2}")
                ans = solve(N,S,E,X,Y)
                ans2 = solvebrute(N,S,E,X,Y)
    if check :
        print(f"{numpass}/{ntc} passed")

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,S,E = gis()
        X = [0] * N; Y = [0] * N
        for i in range(N) : X[i],Y[i] = gis()
        print(f"Case {ntc} N:{N}",file=sys.stderr)
        ans = solve(N,S,E,X,Y)
        print(ans)

if __name__ == '__main__' :
    random.seed(8675309)
    #test(10,1,7)
    #test(100,1,7)
    #test(1000,1,7)
    #test(10000,1,7)
    main()
    sys.stdout.flush()

