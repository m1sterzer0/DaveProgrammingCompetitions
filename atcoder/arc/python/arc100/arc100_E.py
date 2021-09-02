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
    bestidx = [i for i in range(2**N)]
    best2idx = [-1 for i in range(2**N)]
    bestval = A.copy()
    best2val = [0 for i in range(2**N)]
    ans = []
    for i in range(1,2**N) :
        bi1,bi2,b1,b2 = i,-1,A[i],0
        for k in range(N) :
            if i | (1<<k) != i : continue
            cidx = i ^ (1<<k)
            for (ci,cv) in ((bestidx[cidx],bestval[cidx]),(best2idx[cidx],best2val[cidx])) :
                if cv > b1 :
                    bi1,bi2,b1,b2 = ci,bi1,cv,b1
                elif cv > b2 and ci != bi1 :
                    bi2,b2 = ci,cv
        bestidx[i],best2idx[i],bestval[i],best2val[i] = bi1,bi2,b1,b2
        ans.append(max(b1+b2, (-1 if len(ans) == 0 else ans[-1])))
    ansstr = "\n".join(str(x) for x in ans)
    print(ansstr)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

