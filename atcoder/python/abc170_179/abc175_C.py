
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(X,K,D) :
    m1 = abs(X) // D 
    if m1 > K : return abs(X) - K * D
    if K % 2 == m1 % 2 : return  abs(X) - m1 * D
    return abs(abs(X) - (m1+1)*D)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    X,K,D = gis()
    ans = solve(X,K,D)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

