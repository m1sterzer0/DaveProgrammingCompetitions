
import sys
import collections
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
    N,M = gis()
    A = [0] * M
    B = [0] * M
    for i in range(M) : A[i],B[i] = gis()
    gr = [[] for i in range(N)]
    for (a,b) in zip(A,B) :
        aa = a-1; bb=b-1
        gr[aa].append(bb); gr[bb].append(aa)
    myinf = 10**18
    mm = 10**9+7
    d = [myinf] * N; d[0] = 0
    ways = [0] * N; ways[N-1] = 1
    q = collections.deque()

    q.append(0)
    while q :
        n = q.popleft()
        for c in gr[n] :
            if d[c] == myinf :
                d[c] = d[n]+1
                q.append(c)
    if d[N-1] == myinf :
        ans = 0
    else :
        q.append(N-1)
        while q :
            n = q.popleft()
            for c in gr[n] :
                if d[c] == d[n]-1 :
                    if ways[c] == 0 : q.append(c)
                    ways[c] = (ways[c] + ways[n]) % mm
        ans = ways[0]
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

