
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

def solve(N,M,A,B,Y,H) :
    Z = [0] * N
    for i in range(M) : Z[Y[i]-1] = max(Z[Y[i]-1],H[i])
    ZH = list(set([0,1_000_000_001] + Z)); ZH.sort(); NZH = len(ZH)  ## Coordinate compression on the zombie heights
    z2idx = {}
    for (i,z) in enumerate(ZH) : z2idx[z] = i
    winning = [0] * NZH; losing = [0] * NZH; newwinning = [0] * NZH; newlosing = [0] * NZH; sumwinning = [0] * NZH
    if Z[0] == 0 : winning[0] = 1
    else         : losing[z2idx[Z[0]]] = 1

    ## Now for the DP
    ## DP terms
    ## Z[i] = max zombie height in yard i -- 0 if no zombie
    ## winning[h] = have a safe yard, and tallest fence between that safe yard and current yard is of height h
    ## losing[h]  = no safe yard, and rightmost zombie if of height h

    ## When do we need to do modulus?
    ## -- Have ~6000 states, each of which can be 10**9+7
    ## -- State multiplication factor is around 1,000,000
    ## -- Short answer -- everything cane (barely) fit in an Int64, so we can just clean up modulus after each yard.
    #print(f"DBG: winning:{winning} losing:{losing}")
    for i in range(1,N) :
        for j in range(NZH) : newwinning[j] = 0; newlosing[j] = 0
        s = 0
        for j in range(NZH) : s += winning[j]; sumwinning[j] = s


        ## Do the terms where index of the newwinning/newlosing location doesn't depend on the fence height.
        ## Here I care about the number of fences <= or > the old height
        ## Case 1, we have a zombie in our new yard of height z
        ##    * losing[h <= z] * (all fence heights) --> newlosing[z] 
        ##    * winning[h <= z] * (fence heights <= z) --> newlosing[z]
        ##    * losing[h > z] * (fence heights <= h) --> newlosing[h]
        ##    * losing[h > z] * (fence heights > h)  --> newlosing[z]
        ##    * winning[h > z] * (fence heights <= h) --> newwinning[h]
        ## Case 2, we have no zombine in our new yard
        ##    * losing[h] * (fence heights > h) --> newwinning[0]
        ##    * losing[h] * (fence heights <= h) --> newlosing[h]
        ##    * winning[h] * (fence heights <= h) --> newwinning[h]
        totfences = B[i-1]-A[i-1]+1
        zh = Z[i]
        zidx = z2idx[zh]
        lezfences = 0 if A[i-1] > zh else totfences if B[i-1] <= zh else zh - A[i-1] + 1
        gtzfences = totfences-lezfences
        for j in range(NZH) :
            h = ZH[j]
            lefences = 0 if A[i-1] > h else totfences if B[i-1] <= h else h - A[i-1] + 1
            gtfences = totfences - lefences
            if zh > 0 :
                if j <= zidx :
                    newlosing[zidx] += losing[j] * totfences
                    newlosing[zidx] += winning[j] * lezfences
                else :
                    newlosing[j]    += losing[j] * lefences
                    newlosing[zidx] += losing[j] * gtfences
                    newwinning[j]   += winning[j] * lefences
            else :
                newwinning[0] += losing[j] * gtfences
                newlosing[j]  += losing[j] * lefences
                newwinning[j] += winning[j] * lefences

        ## Now we do the terms where we care about the the fence height for our destination
        ## We need to refactor this code into cumsums. The cases collapse into one simple case
        ##    * if fh > z : sumwinning[fh-1] * segmentwidth --> newwinning[fence height]
        for j in range(NZH) :
            if j <= zidx : continue
            lb,ub = max(ZH[j-1]+1,A[i-1]),min(ZH[j],B[i-1])
            segwidth = ub-lb+1
            if segwidth <= 0 : continue
            newwinning[j] += sumwinning[j-1] * segwidth

        ## Cleanup
        for j in range(NZH) : newwinning[j] %= MOD; newlosing[j] %= MOD
        winning,newwinning = newwinning,winning
        losing,newlosing = newlosing,losing
        #print(f"DBG: i:{i} winning:{winning} losing:{losing}")

    numstates = 1
    for (a,b) in zip(A,B) : numstates = numstates * (b-a+1) % MOD
    totwinning = sum(winning) % MOD
    prob = totwinning * pow(numstates,MOD-2,MOD) % MOD
    return prob

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M = gis()
        A = [0] * (N-1)
        B = [0] * (N-1)
        for i in range(N-1) : A[i],B[i] = gis()
        Y = [0] * M
        H = [0] * M
        for i in range(M) : Y[i],H[i] = gis()
        print(f"Case {ntc} N:{N} M:{M}",file=sys.stderr)
        ans = solve(N,M,A,B,Y,H)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

