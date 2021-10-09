
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def matmul(A,B,C,N,mm) :
    for i in range(N) :
        for j in range(N) :
            res = 0
            for k in range(N) :
                res = (res + A[i][k] * B[k][j]) % mm
            C[i][j] = res

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    H,W = gis()
    mm = 998244353
    full = 2**H-1
    tr = [0] * (2**H)
    for i in range(2**H) :
        tr[i] = [0] * (2**H)
    for prev in range(2**H) :
        startpos = [full ^ prev] ## Bridge holes with horizontal pieces
        ## Now we deal with the double vertical pieces
        for i in range(H-1) :
            adders = []
            for p in startpos :
                if p & (1<<i)     != 0: continue
                if p & (1<<(i+1)) != 0: continue
                adders.append(p | (1<<i) | (1<<(i+1)))
            startpos += adders
        ## Now we deal with the ways to fill in the single pieces
        for p in startpos :
            for i in range(2**H) :
                if i | p == i : tr[prev][i] += 1
    
    tr2   = [0] * (2**H)
    final = [0] * (2**H)
    final2 = [0] * (2**H)
    for i in range(2**H) :
        tr2[i] = [0] * (2**H)
        final[i] = [0] * (2**H)
        final2[i] = [0] * (2**H)
        final[i][i] = 1
        final2[i][i] = 1
    w2 = W
    while w2 > 0 :
        if w2 & 1 == 1 :
            matmul(final,tr,final2,2**H,mm)
            final,final2=final2,final
        w2 >>= 1
        matmul(tr,tr,tr2,2**H,mm)
        tr,tr2 = tr2,tr
    ans = final[full][full]
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

