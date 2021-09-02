
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
    N,K,C = gis()
    S = gs()
    earliest = [0] * K
    latest = [0] * K
    last = -10**18; idx = 0
    for i in range(N)  :
        c = S[i]
        if c == 'x' : continue
        if i-last <= C : continue
        earliest[idx] = i; last = i; idx += 1
        if idx == K : break
    last = 10**18; idx = K-1
    for i in range(N-1,-1,-1)  :
        c = S[i]
        if c == 'x' : continue
        if last-i <= C : continue
        latest[idx] = i; last = i; idx -= 1
        if idx < 0 : break
    ansarr = [x for (x,y) in zip(earliest,latest) if x == y]
    ansstr = "\n".join([str(x+1) for x in ansarr])
    if ansstr : print(ansstr)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

