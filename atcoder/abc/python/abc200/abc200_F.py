
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
    S = gs()
    K = gi()
    ## Want to count total inversions and divide by two
    ## Key trick -- assume string wraps around back to the beginning -- this makes the inversion count even, and since fixing one inversion costs the same
    ## as fixing two -- rounding up to even is exactly what we want to do
    if len(S) == 1 and K == 1 : 
        print(0)
    else :
        mm = 10**9+7
        q = S.count("?")
        totstrings = pow(2,K*q,mm)
        half = pow(2,mm-2,mm)
        totinv = 0
        for i in range(len(S)) :
            c1 = S[i]
            c2 = S[(i+1)%len(S)]
            numinv = totstrings * K % mm * (1 if c1+c2 in ("01","10") else 0 if c1+c2 in ("00","11") else half) % mm
            totinv = (totinv + numinv) % mm
        ans = totinv * half % mm
        sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

