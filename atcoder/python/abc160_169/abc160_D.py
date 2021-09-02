
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
    N,X,Y = gis()
    X -= 1
    Y -= 1
    ans = [0] * N
    dist = 0
    for i in range(N) :
        for j in range(i+1,N) :
            if j <= X or i >= Y : dist = j - i
            elif i <= X and j >= Y : dist = X-i + 1 + j-Y
            elif i <= X : dist = X - i + min(j-X,1+Y-j)
            elif j >= Y : dist = j - Y + min(Y-i,i-X+1)
            else : dist = min(j-i,i-X+1+Y-j)
            ans[dist] += 1
    ansstr = "\n".join([str(x) for x in ans[1:]])
    print(ansstr)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

