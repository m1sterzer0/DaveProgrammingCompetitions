
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
    N,K = gis()
    A = gis()
    for i in range(N) : A[i] -= 1
    c2i = [-1] * N; i2c = [-1] * N ; c2i[0] = 0; i2c[0] = 0
    v = 0
    for i in range(1,2*N) :
        v = A[v]
        if i == K : print(v+1); break
        if c2i[v] == -1 : c2i[v] = i; i2c[i] = v; continue
        offset = (K-i) % (i - c2i[v])
        endpoint = i2c[c2i[v]+offset]
        print(endpoint+1)
        break

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

