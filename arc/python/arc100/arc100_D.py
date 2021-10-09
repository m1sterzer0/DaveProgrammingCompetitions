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
    N = gi()
    A = gis()
    prefixSums = [0] * N; s = 0
    for i in range(N) :
        s += A[i]
        prefixSums[i] = s
    inf = 10**18
    L1 = [0] * N; L2 = [0] * N; R1 = [0] * N; R2 = [0] * N
    lptr = 0
    for i in range(1,N-2) :
        bestl = prefixSums[lptr]
        bestr = prefixSums[i]-prefixSums[lptr]
        best = abs(bestr-bestl)
        while True :
            if lptr+1 == i : break
            candl = prefixSums[lptr+1]
            candr = prefixSums[i]-prefixSums[lptr+1]
            cand  = abs(candr-candl)
            if cand >= best : break
            lptr += 1
            best,bestl,bestr = cand,candl,candr
        L1[i],L2[i] = bestl,bestr
    rptr = N-2
    for i in range(N-3,0,-1) :
        lsum = prefixSums[i]
        bestl = prefixSums[rptr]-lsum
        bestr = prefixSums[-1]-lsum-bestl
        best = abs(bestr-bestl)
        while True :
            if rptr-1 == i : break
            candl = prefixSums[rptr-1]-lsum
            candr = prefixSums[-1]-lsum-candl
            cand = abs(candr-candl)
            if cand >= best : break
            rptr -= 1
            best,bestl,bestr = cand,candl,candr
        R1[i],R2[i] = bestl,bestr
    best = inf
    for i in range(1,N-2) : 
        best = min(best,max(L1[i],L2[i],R1[i],R2[i]) - min(L1[i],L2[i],R1[i],R2[i]))
    print(best)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

