
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class fenwicktree :
    def __init__(self,n=1) :
        self.n = n
        self.tot = 0
        self.bit = [0] * (n+1)

    def clear(self) :
        for i in range(self.n) : self.bit[i] = 0
        self.tot = 0

    def inc(self,idx,val=1) :
        while idx <= self.n :
            self.bit[idx] += val
            idx += idx & (-idx)
        self.tot += val

    def dec(self,idx,val=1) : self.inc(idx,-val)

    def incdec(self,left,right,val) :
        self.inc(left,val); self.dec(right,val)

    def prefixsum(self,idx) :
        if idx < 1 : return 0
        ans = 0
        while idx > 0 :
            ans += self.bit[idx]
            idx -= idx&(-idx)
        return ans

    def suffixsum(self,idx) : return self.tot - self.prefixsum(idx-1)
    def rangesum(self,left,right)  : return self.prefixsum(right) - self.prefixsum(left-1)

def toohigh(A,N,x) :
    thresh = N * (N+1) // 2
    AA = [1 if a >= x else -1 for a in A]
    psum = 0; prefixsum = [0] * N
    for (i,a) in enumerate(AA) : psum += a; prefixsum[i] = psum
    ft = fenwicktree(2*N+2)
    ft.inc(N+1+0)
    cnt = 0
    for (i,p) in enumerate(prefixsum,start=1) :
        ## We are too high if the number of sums < 0 is > half the threshold
        ## Want p - suffixsum < 0 --> suffixsum > p --> suffixsum >= p+1
        cnt += ft.suffixsum(N+1+p+1)
        ft.inc(N+1+p)
    #print(f"m:{x} cnt:{cnt} thresh:{thresh} A:{A} AA:{AA} toolow:{2*cnt>thresh}")
    return 2*cnt > thresh

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    ## Binary search on the median
    l,u = 0,10**9+1
    while (u-l) > 1 :
        m = (u+l) >> 1
        (l,u) = (l,m) if toohigh(A,N,m) else (m,u)
    print(l)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

