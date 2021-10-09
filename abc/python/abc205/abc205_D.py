
import sys
import random
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,Q,A,K) :
    QQ = [K[i]<<20|i for i in range(Q)]
    QQ.sort(reverse=True) 
    AA = A.copy()
    AA.sort(reverse=True)
    ans = [0] * Q
    last,candsofar = 0,0
    while QQ :
        qv = QQ[-1] >> 20
        qidx = QQ[-1] & 0xfffff
        if len(AA) == 0 or qv < candsofar + (AA[-1] - last) :
            QQ.pop()
            ans[qidx] = last + (qv - candsofar)
        else :
            aa = AA.pop()
            candsofar += 0 if aa == last else (aa - last - 1)  
            last = aa
    return ans

def test(ntc,Nmin,Nmax,Qmin,Qmax,Vmax) :
    for tc in range(ntc) :
        N = random.randrange(Nmin,Nmax+1)
        Q = random.randrange(Qmin,Qmax+1)
        A = [random.randrange(1,Vmax+1) for i in range(N)]
        K = [random.randrange(1,Vmax+1) for i in range(Q)]
        print(f"Tescase #{tc}")
        ans = solve(N,Q,A,K)
        for a in ans : print(a)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,Q = gis()
    A = gis()
    K = [0] * Q
    for i in range(Q) : K[i] = gi()
    ans = solve(N,Q,A,K)
    for i in range(Q) :
        print(ans[i])


if __name__ == '__main__' :
    main()
    sys.stdout.flush()
    #random.seed(1234)
    #test(10,10**4,10**5,10**4,10**5,10**18)


