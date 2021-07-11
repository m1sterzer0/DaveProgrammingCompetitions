
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(H,W,K,bd) :
    ans = 0
    blacksquares = []
    for i in range(H) :
        for j in range(W) :
            if bd[i][j] == "#" : blacksquares.append((i,j))
    for hmask in range(2**H) :
        for wmask in range(2**W) :
            lcnt = 0
            for (i,j) in blacksquares :
                if (1 << i) & hmask== 0 and (1 << j) & wmask == 0 : lcnt += 1
            if lcnt == K : ans += 1
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    H,W,K = gis()
    bd = []
    for i in range(H) : bd.append(gs())
    ans = solve(H,W,K,bd)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

