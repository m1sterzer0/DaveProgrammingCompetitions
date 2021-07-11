
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def comb(n,r) : 
    r = min(r,n-r)
    if r == 0 : return 1
    ans = 1
    for i in range(1,r+1) : ans *= (n-i+1); ans //= i
    return ans

def solve(A,B,K) :
    ans = []; n = A+B
    while A > 0 or B > 0 :
        if A == 0 : ans.append("b"); B -= 1; continue
        if B == 0 : ans.append("a"); A -= 1; continue
        numwa = comb(B+A-1,A-1)
        if K > numwa : K -= numwa; ans.append("b"); B -= 1
        else         : ans.append("a"); A -= 1
    return "".join(ans)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    A,B,K = gis()
    ans = solve(A,B,K)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

