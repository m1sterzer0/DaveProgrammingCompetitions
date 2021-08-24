
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def fsieve(n) :
    res = [-1] * (n+1)
    for k in range(2,n+1,2) : res[k] = 2
    for i in range(3,n+1,2) :
        if res[i] > 0 : continue
        res[i] = i
        for k in range(i*i,n+1,2*i) : 
            if res[k] == -1 : res[k] = i
    return res

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    A = gis()
    factors = fsieve(100_001)
    pfactors = [False] * (100_001)
    for a in A :
        aa = a
        while aa > 1 :
            p = factors[aa]
            pfactors[p] = True
            while aa % p == 0 : aa //= p
    sb = [False] * (100_001); ans = [1]; sb[1] = True
    for i in range(2,M+1) :
        p = factors[i]
        if not pfactors[p] and sb[i//p] : ans.append(i); sb[i] = True
    ansstr = "\n".join([str(x) for x in ans])
    sys.stdout.write(str(len(ans))+'\n')
    sys.stdout.write(str(ansstr)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

