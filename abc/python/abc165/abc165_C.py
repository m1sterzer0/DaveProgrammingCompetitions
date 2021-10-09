
import sys
import itertools
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## Up to 11 slots for increases == 10 bars
## Budget of up to 9 increases
## Total ways is comb(19,9) which is around 90k which is doable 

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M,Q = gis()
    A = []; B = []; C = []; D = []
    for _ in range(Q) :
        a,b,c,d = gis(); A.append(a-1); B.append(b-1); C.append(c); D.append(d)
    pool = [x for x in range(M-1+N)]
    seq = [1] * N
    ans = 0
    for ccc in itertools.combinations(pool,N) :
        seq[0] = 1 + ccc[0]
        for i in range(1,N) : seq[i] = seq[i-1] + (ccc[i] - ccc[i-1] - 1)
        #print(seq)
        lscore = 0
        for (a,b,c,d) in zip(A,B,C,D) :
            if seq[b]-seq[a] == c : lscore += d
        ans = max(ans,lscore)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

